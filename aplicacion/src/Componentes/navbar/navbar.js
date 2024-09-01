import React , {useState} from "react";
import "./navbar.css";

export const Navbar = () => {
    const [estado , cambiarEstado] = useState(false)
    const accionMenu = () => {
        cambiarEstado(!estado)
    }
    return (
        <header className="header">
            <a href="#" className="logo">Logo</a>
            <nav className={`navbar ${estado ? "activo" : ""}`}>
                <a href="/">Inicio</a>
                <a href="/">Egresados</a>
                <a href="/">Eventos</a>
                <a href="/">Trabajos</a>
            </nav>
            <button className="c-hamburguesa" onClick={ accionMenu }>
                <i className="bi bi-list i-hamburguesa"></i>
            </button>
        </header>
    )
}