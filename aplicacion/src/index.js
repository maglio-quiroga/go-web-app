import React from 'react'
import ReactDom from 'react-dom/client'
import {Navbar} from './Componentes/navbar/navbar'
import "./index.css"

const root = ReactDom.createRoot(document.getElementById("root"))
root.render(
    <>
        <Navbar/>
    </>
)