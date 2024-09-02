import React from "react";
import './alumnos.css'

export const Tarjeta = () => {
    return (
        <div className="card">
        <div className="card-header">
            <img src="#" alt="Imagen Ex-Alumno" class="profile-img"></img>
        </div>
        <div className="card-body">
            <p class="name">Nombre</p>
            <a href="#" class="mail">Email</a>
            <p class="job">Trabajo | Area</p>
        </div>
        <div className="card-footer">
            <p class="count"><span>Egreso</span> Fecha | <span>Contacto</span> Numero</p>
        </div>
    </div>
    );
}