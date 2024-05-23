public class Persona {
    private String nombre;
    private String apellido;
    private String telefono;
    private String email;

    public Persona(String nombre, String apellido, String telefono, String email) {
        this.nombre = nombre;
        this.apellido = apellido;
        this.telefono = telefono;
        this.email = email;
    }

    @Override
    public String toString() {
        return "Persona [Nombre=" + nombre + ", Apellido=" + apellido + ", Teléfono=" + telefono + ", Email=" + email
                + "]";
    }
}
