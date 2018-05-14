import React, { Component } from 'react';
import Modal from 'react-bootstrap4-modal';

export default class Product extends Component {
    constructor(props) {
        super(props);
        this.state = {
            name: '',
            label: '',
            price: '',
            description: ''
        };

        

        this.handleNameChange = this.handleNameChange.bind(this);
        this.handleLabelChange = this.handleLabelChange.bind(this);
        this.handlePriceChange = this.handlePriceChange.bind(this);
        this.handleDescChange = this.handleDescChange.bind(this);

        this.updateProduct = this.updateProduct.bind(this)
    }

    componentDidMount() {
        const { selectedProduct } = this.props;
        // console.log('selectedProduct : ', selectedProduct)
        if (selectedProduct) {
            this.setState({ ...selectedProduct })
        }
    }
    handleNameChange(e) {
        this.setState({ name: e.target.value});
    }

    handlePriceChange(e) {
        this.setState({price: e.target.value});
    }

    handleLabelChange(e) {
        this.setState({label: e.target.value});
    }

    handleDescChange(e) {
        this.setState({description: e.target.value});
    }


    updateProduct() {
        const { saveProduct } = this.props;
        saveProduct({...this.state})
    }
    render() {
        const { 
            selectedProduct,
            showProduct,
            showHideProductHandler,
        } = this.props;

        
        let addOrEdit ='Add'
        if (selectedProduct) {
            addOrEdit = 'Edit'
        }

        return (
            <Modal visible={showProduct} onClickBackdrop={showHideProductHandler}>
                <div className="modal-content">
                    <div className="modal-header">
                    <h5 className="modal-title" id="exampleModalLabel">{addOrEdit} Product </h5>
                    <button type="button" className="close" onClick={showHideProductHandler} data-dismiss="modal" aria-label="Close">
                        <span aria-hidden="true">&times;</span>
                    </button>
                    </div>
                    <div className="modal-body">
                        <form>
                        <div className="form-group">
                                <label htmlFor="productName">Name</label>
                                <input type="text" required={true} value={this.state.name} onChange={this.handleNameChange}  className="form-control" id="productName" placeholder="Enter Product name" />
                            </div>
                            <div className="form-group">
                                <label htmlFor="productLabel">Label</label>
                                <input type="text" value={this.state.label} onChange={this.handleLabelChange}  className="form-control" id="productLabel" placeholder="Enter Product label" />
                            </div>
                            <div className="form-group">
                                <label htmlFor="productPrice">Price</label>
                                <input type="number" value={this.state.price} onChange={this.handlePriceChange}  className="form-control" id="productPrice" placeholder="Enter Product price" />
                            </div>
                            <div className="form-group">
                                <label htmlFor="productDescription">Description</label>
                                <input type="text" value={this.state.description} onChange={this.handleDescChange}  className="form-control" id="productDescription" placeholder="Enter Product description" />
                            </div>
                        </form>
                    </div>
                    <div className="modal-footer">
                    <button onClick={showHideProductHandler} type="button" className="btn btn-secondary" data-dismiss="modal">Close</button>
                    <button type="button" onClick={this.updateProduct} className="btn btn-primary">Save changes</button>
                    </div>
                </div>
            </Modal>
        )
    }
}
