import React from "react";
import './alumnos.css'

export const Tarjetas = ({ alumnos }) => {
    return (
        <>
            {alumnos.map((alumno, index) => (
                <div key={index} className="card">
                    <div className="card-header">
                        <img src={alumno.Imagen || "#"} alt="Imagen Ex-Alumno" className="profile-img" />
                    </div>
                    <div className="card-body">
                        <p className="name">{alumno.Nombre}</p>
                        <a href={`mailto:${alumno.Email}`} className="mail">{alumno.Email}</a>
                        <p className="job">{alumno.Trabajo} | {alumno.Area}</p>
                    </div>
                    <div className="card-footer">
                        <p className="count"><span>Egreso:</span> {new Date(alumno.Egreso).toLocaleDateString()} | <span>Contacto:</span> {alumno.Contacto}</p>
                    </div>
                </div>
            ))}
        </>
    );
};
