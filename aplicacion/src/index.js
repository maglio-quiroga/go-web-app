import React from 'react'
import ReactDom from 'react-dom/client'
import {Navbar} from './Componentes/navbar/navbar'
import {GaleriaImagenes} from './Componentes/carousel/carousel'
import {Separador} from './Componentes/separador/separador'
import {Tarjeta} from './Componentes/tarjetas/alumnos'
import {Calendario} from './Componentes/calendario/calendario'
import {Formulario} from './Componentes/formulario/form'
import {Footer} from './Componentes/footer/footer'
import "./index.css"

const info = async () => {
    const response = await fetch("http://localhost:8080/inicio")
    const data = await response.json()
    return data
}

info().then(data => console.log(data));

const root = ReactDom.createRoot(document.getElementById("root"))
root.render(
    <>
        <Navbar/>
        <GaleriaImagenes/>
        <Separador
            nombre = "Egresados Mas Recientes"
        />
        <div className='card-container'>
        <Tarjeta></Tarjeta>
        <Tarjeta></Tarjeta>
        <Tarjeta></Tarjeta>
        </div>
        <Separador
            nombre = "Proximos Eventos"
        />
        <Calendario></Calendario>
        <Separador
            nombre = "Formulario Ex-Alumnos"
        />
        <Formulario></Formulario>
        <div className='fot'>
            <Footer></Footer>
        </div>
    </>
)