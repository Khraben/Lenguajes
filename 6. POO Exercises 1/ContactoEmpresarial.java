public class ContactoEmpresarial extends Contacto {
    private String empresa;
    private String cargo;

    public ContactoEmpresarial(Persona persona, String empresa, String cargo) {
        super(persona);
        this.empresa = empresa;
        this.cargo = cargo;
    }

    @Override
    public String toString() {
        return "ContactoEmpresarial [Empresa=" + empresa + ", Cargo=" + cargo + ", Persona=" + super.toString() + "]";
    }
}
