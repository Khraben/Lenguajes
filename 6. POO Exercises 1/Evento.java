public class Evento {
    private String descripcion;
    private String fecha;

    public Evento(String descripcion, String fecha) {
        this.descripcion = descripcion;
        this.fecha = fecha;
    }

    @Override
    public String toString() {
        return "Evento [Descripción=" + descripcion + ", Fecha=" + fecha + "]";
    }
}
