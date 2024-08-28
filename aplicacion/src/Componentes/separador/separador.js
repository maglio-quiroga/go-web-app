import React from "react";
import './separador.css';

export const Separador = (props) => {
    return (
        <div className="separador">
            <div className="contenido">
                {props.nombre}
            </div>
        </div>
    );
}