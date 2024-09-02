import React from 'react';
import { Calendar, dayjsLocalizer } from 'react-big-calendar';
import 'react-big-calendar/lib/css/react-big-calendar.css';
import dayjs from 'dayjs';

export const Calendario = ({ eventos }) => {
    const loc = dayjsLocalizer(dayjs);
    const eventosFormateados = eventos.map(evento => ({
        title: evento.Nombre,
        start: new Date(evento.Inicio),
        end: new Date(evento.Final),
        allDay: false,
        resource: evento.Ubicacion
    }));

    return (
        <div style={{ height: "95vh", width: "98vw", padding: "30px" }}>
            <Calendar
                localizer={loc}
                events={eventosFormateados}
                startAccessor="start"
                endAccessor="end"
            />
        </div>
    );
};
