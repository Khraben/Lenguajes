public class ContactoFamiliar extends Contacto {
    private String relacion;

    public ContactoFamiliar(Persona persona, String relacion) {
        super(persona);
        this.relacion = relacion;
    }

    @Override
    public String toString() {
        return "ContactoFamiliar [Relacion=" + relacion + ", Persona=" + super.toString() + "]";
    }
}
