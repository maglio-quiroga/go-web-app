import {Calendar , dayjsLocalizer} from 'react-big-calendar';
import 'react-big-calendar/lib/css/react-big-calendar.css';
import dayjs from 'dayjs';

export const Calendario = () => {
    const loc = dayjsLocalizer(dayjs)
    return (
        <div style={
            {height : "95vh" , width: "98vw" , padding : "30px"}
        }>
            <Calendar
                localizer={loc}
            />
        </div>
    )

};