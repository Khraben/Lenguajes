public class ContactoFactory implements AbstractFactory<Contacto> {
 @Override
 public Contacto crear(String tipo) {
  Persona persona = new Persona("Nombre", "Apellido", "Telefono", "Email");
  switch (tipo) {
   case "Familiar":
    return new ContactoFamiliar(persona, "Relación");
   case "Empresarial":
    return new ContactoEmpresarial(persona, "Empresa", "Cargo");
   default:
    return new Contacto(persona);
  }
 }
}
