public class Contacto {
    private Persona persona;

    public Contacto(Persona persona) {
        this.persona = persona;
    }

    @Override
    public String toString() {
        return "Contacto [Persona=" + persona + "]";
    }
}
