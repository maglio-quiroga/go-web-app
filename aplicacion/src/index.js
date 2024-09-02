import React, { useState, useEffect } from 'react'
import ReactDom from 'react-dom/client'
import { Navbar } from './Componentes/navbar/navbar'
import { GaleriaImagenes } from './Componentes/carousel/carousel'
import { Separador } from './Componentes/separador/separador'
import { Tarjetas } from './Componentes/tarjetas/alumnos'
import { Calendario } from './Componentes/calendario/calendario'
import { Formulario } from './Componentes/formulario/form'
import { Footer } from './Componentes/footer/footer'
import "./index.css"

const App = () => {
    const [alumnos, setAlumnos] = useState([])

    useEffect(() => {
        const fetchData = async () => {
            const response = await fetch("http://localhost:8080/inicio")
            const data = await response.json()
            setAlumnos(data.Alumnos)
        }

        fetchData()
    }, [])

    const [eventos , setEventos] = useState([])

    useEffect(
        () => {
            const fetchData = async () => {
                const response = await fetch("http://localhost:8080/inicio")
                const data = await response.json()
                setEventos(data.Eventos)
            }
            fetchData()
        }, []
    )

    return (
        <>
            <Navbar/>
            <GaleriaImagenes/>
            <Separador nombre="Egresados Mas Recientes" />
            <div className='card-container'>
                <Tarjetas alumnos={alumnos} />
            </div>
            <Separador nombre="Proximos Eventos" />
            <Calendario
                eventos={eventos}
            />
            <Separador nombre="Formulario Ex-Alumnos" />
            <Formulario/>
            <div className='fot'>
                <Footer/>
            </div>
        </>
    )
}

const root = ReactDom.createRoot(document.getElementById("root"))
root.render(<App />)
