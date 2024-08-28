import React from "react";
import "./navbar.css";

export const Navbar = () => {
    return (
        <header className="header">
            <a href="#" className="logo">Logo</a>
            <nav className="navbar">
                <a href="/">Inicio</a>
                <a href="/">Egresados</a>
                <a href="/">Eventos</a>
                <a href="/">Trabajos</a>
            </nav>
        </header>
    )
}