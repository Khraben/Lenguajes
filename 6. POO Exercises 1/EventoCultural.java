public class EventoCultural extends Evento {
    private String organizador;

    public EventoCultural(String descripcion, String fecha, String organizador) {
        super(descripcion, fecha);
        this.organizador = organizador;
    }

    @Override
    public String toString() {
        return "EventoCultural [Organizador=" + organizador + ", " + super.toString() + "]";
    }
}
