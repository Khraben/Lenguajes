/*
 * CON RESPECTO A LA ELECCION ENTRE EAGLE O LAZY SINGLETON TENEMOS LO SIGUIENTE:
   -  Eagle Singleton es una implementación de Singleton que crea la instancia de la clase en el momento de la carga de la clase
   -  Lazy Singleton crea la instancia de la clase solo cuando se llama al método de obtención de la instancia.
   :. En este caso, se eligió Lazy Singleton porque no se necesita crear la instancia de la clase Agenda hasta que se llame al método agregarElemento.
 * CON RESPECTO A EL LUGAR DONDE SE USA ABSTRACT FACTORY TENEMOS LO SIGUIENTE:
   -  Se una en la creación de los objetos de tipo Contacto y Evento, que son creados por las fábricas ContactoFactory y EventoFactory respectivamente.
   :. Se eligieron estas clases dado que Abstract Factory porque permite crear familias de objetos relacionados sin especificar sus clases concretas.
*/

public class Main {
 public static void main(String[] args) {
  // Crear fabrica de contactos y eventos
  AbstractFactory<Contacto> contactoFactory = new ContactoFactory();
  AbstractFactory<Evento> eventoFactory = new EventoFactory();

  // Crear contactos
  Contacto contacto1 = contactoFactory.crear("Familiar");
  Contacto contacto2 = contactoFactory.crear("Empresarial");

  // Crear eventos
  Evento evento1 = eventoFactory.crear("Deportivo");
  Evento evento2 = eventoFactory.crear("Cultural");

  // Obtener instancia de la agenda (Singleton Lazy)
  Agenda agenda = Agenda.getInstancia();

  // Agregar elementos a la agenda
  agenda.agregarElemento(contacto1);
  agenda.agregarElemento(contacto2);
  agenda.agregarElemento(evento1);
  agenda.agregarElemento(evento2);

  // Imprimir todos los elementos de la agenda
  System.out.println("Todos los elementos de la agenda:");
  System.out.println(agenda);

  // Filtrar y mostrar eventos y contactos
  System.out.println("\nEventos:");
  agenda.filtrarEventos().forEach(System.out::println);

  System.out.println("\nContactos:");
  agenda.filtrarContactos().forEach(System.out::println);
 }
}
