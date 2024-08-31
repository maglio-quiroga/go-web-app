import React from 'react';
import './form.css';


export const Formulario = () => {
    return (
        <div className='container'>
            <form method='POST' action='#' encType='multipart/form-data' className='form'>
                <label htmlFor='nombre'>Nombre</label>
                <input type='text' name='nombre' id='nombre' placeholder='Nombre y Apellido'></input>
                <label htmlFor='email'>Email</label>
                <input type='email' name='email' id='email' placeholder='Email Institucional'></input>
                <label htmlFor='egreso'>Fecha de Egreso</label>
                <input type='date' name='egreso' id='egreso'></input>
                <label htmlFor='trabajo'>Trabajo</label>
                <input type='text' name='trabajo' id='trabajo' placeholder='Trabajo Actual'></input>
                <label htmlFor='contacto'>Contacto</label>
                <input type='text' name='contacto' id='contacto' placeholder='Numero Telefonico'></input>
                <label htmlFor='archivo' className='file-label'>Imagen de Perfil</label>
                <input type='file' name='archivo' id='archivo' className='file-input'></input>
                <label htmlFor='area'>Area de Interes</label>
                <select id="area" name="area">
                <option value="opcion1">Ingenieria de Software</option>
                <option value="opcion2">TI/TIC/TICAR</option>
                <option value="opcion3">Inteligencia de Maquinas</option>
                </select>
                <label htmlFor='desc'>Descripcion</label>
                <textarea name='desc'placeholder='Agrega una breve descripcion tuya'></textarea>
                <button type='submit' className='submit-button'>Enviar Peticion</button>
            </form>
        </div>
    )
};