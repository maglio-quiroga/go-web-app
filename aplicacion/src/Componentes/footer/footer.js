import React from 'react';
import './footer.css';

export const Footer = () => {
    const logoUrl = '#';
    return (
            <>
            <div className='img'>
                <img src={logoUrl} alt='Logo Uda'></img>
            </div>
            <div className='links'>
                <h3>Enlaces</h3>
                <a href='#'>Intranet Alumnos</a>
                <a href='#'>Intranet Académicos</a>
                <a href='#'>Moodle</a>
                <a href='#'>Instagram</a>
            </div>
            <div className='contact'>
                <h3>Contactos</h3>
                <pre>Ubícanos en
                Copiapó, Av. Copayapu 485</pre>
                <p>52 2 255555</p>
                <p>anakarina.pena@uda.cl</p>
            </div>
            </>
    )
};