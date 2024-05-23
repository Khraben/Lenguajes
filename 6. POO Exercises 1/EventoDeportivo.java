public class EventoDeportivo extends Evento {
    private String lugar;

    public EventoDeportivo(String descripcion, String fecha, String lugar) {
        super(descripcion, fecha);
        this.lugar = lugar;
    }

    @Override
    public String toString() {
        return "EventoDeportivo [Lugar=" + lugar + ", " + super.toString() + "]";
    }
}
