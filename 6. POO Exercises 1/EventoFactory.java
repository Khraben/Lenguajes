public class EventoFactory implements AbstractFactory<Evento> {
 @Override
 public Evento crear(String tipo) {
  switch (tipo) {
   case "Deportivo":
    return new EventoDeportivo("Descripción", "Fecha", "Lugar");
   case "Cultural":
    return new EventoCultural("Descripción", "Fecha", "Organizador");
   default:
    return new Evento("Descripción", "Fecha");
  }
 }
}
