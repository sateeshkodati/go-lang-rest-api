import React, { Component } from 'react';

import './App.css';

import * as api from './api/products';

import ProductList from './components/ProductList';

class App extends Component {
  constructor(props) {
    super(props);
    this.state = {
      selectedProduct: null,
      selectedProductName: '',
      products: null,
      showProduct: false,
    };

    this.showHideProductHandler = this.showHideProductHandler.bind(this)
    
    this.addProduct = this.addProduct.bind(this);
    this.editProduct = this.editProduct.bind(this)
    this.deleteProduct = this.deleteProduct.bind(this);
    this.saveProduct = this.saveProduct.bind(this);

  }
  async componentDidMount() {
    try {
      const data = await api.getAll();
      this.setState({ products: data });  
    } catch (error) {
      //show error page
    }
  }

  showHideProductHandler() {
    const { showProduct } = this.state;
    this.setState({ showProduct: !showProduct });
  }

  addProduct() {
    this.setState({ selectedProduct: null });
    this.showHideProductHandler()
  }
  editProduct(product) {
    // const {products} = this.state;
    // const p = products.filter(p => p.name === product.name)[0];
    
    this.setState({ selectedProduct: product, selectedProductName: product.name });  
    // this.setState({selectedProduct})

    console.log(this.state)
    this.showHideProductHandler()
  }
  
  async saveProduct(product) {
    if (!product.name) {
      alert('name is required');
      return;
    }
    
    if (!product.price) {
      alert('name is required');
      return;
    }
    try {
      product.price = parseFloat(product.price)  
    } catch (error) {
      alert('price should be float value')
      return
    }
    
    if (product.price === NaN) {
      product.price='';
      alert('price should be float value')
      return
    }
    
    const {  products, selectedProduct } = this.state;
    
    if (selectedProduct) {
      products.map((p, i) => {
        if (p.name === product.name) {
          products[i] = product;
        }
      });
    } else {
      products.push(product)
    }
    
    if (!selectedProduct) {
      const { data } = await api.create(product);
    } else {
      const result = await api.update(product.name, product);
    }

    this.setState({ products: products });

    this.showHideProductHandler();
    this.setState({ selectedproduct: null });
  }

  async deleteProduct(productName) {
    const products = this.state.products.filter(p => p.name !== productName);
    const { data } = await api.remove(productName);
    this.setState({ products })
  }

  render() {
    const { products, showProduct, selectedProduct } = this.state;
    
    return (
      <div className="App">
        <header className="App-header">
          <h1 className="App-title">Go lang - Products web app</h1>
        </header>
        <div className="App-body">
          <div style={{paddingBottom:20}}>
            <button type="button" onClick={this.addProduct} className="btn btn-outline-primary" data-toggle="modal" data-target="#exampleModal">Add Product</button>
          </div>
          <ProductList 
            selectedProduct={selectedProduct}
            products={products}
            showProduct={showProduct}
            showHideProductHandler={this.showHideProductHandler} 
            editProduct={this.editProduct}
            saveProduct={this.saveProduct}
            deleteProduct={this.deleteProduct}
          />
        </div>
      </div>
    );
  }
}

export default App;
