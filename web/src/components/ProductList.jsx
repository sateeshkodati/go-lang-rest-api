
import React from 'react';

import Product from './Product';

function ProductList(props) {
    const { 
        products,
        showProduct,
        selectedProduct,
        showHideProductHandler,
        editProduct,
        deleteProduct,
        saveProduct,
    } = props;

    return (
        <div>
            <table className="table">
                <thead className="thead-light">
                <tr>
                    <th scope="col">Name</th>
                    <th scope="col">Label</th>
                    <th scope="col">Price</th>
                    <th scope="col">Description</th>
                    <th scope="col">&nbsp;</th>
                    <th scope="col">&nbsp;</th>
                </tr>
                </thead>
                <tbody>
                { products && products.map(p => 
                    (<tr key={p.name}>
                    <td>{p.name}</td>
                    <td>{p.label}</td>
                    <td>{p.price}</td>
                    <td>{p.description}</td>
                    <td><button type="button" onClick={() => editProduct({...p})} className="btn btn-outline-primary">View/ Edit</button></td>
                    <td><button type="button" onClick={() => deleteProduct(p.name)} className="btn btn-outline-danger">Delete</button></td>
                </tr>)
                )}
                </tbody>
            </table>
            
            { showProduct && 
            <Product
            showProduct={showProduct}
            selectedProduct={selectedProduct}
            showHideProductHandler={showHideProductHandler} 
            saveProduct={saveProduct}
        />
            }
            
        </div>
    )
}

export default ProductList;
