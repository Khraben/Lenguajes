import java.util.ArrayList;
import java.util.List;
import java.util.stream.Collectors;

//LAZY SINGLETON
public class Agenda {
  private List<Object> elementos;

  private Agenda() {
    this.elementos = new ArrayList<>();
  }

  private static class AgendaHolder {
    private static final Agenda instancia = new Agenda();
  }

  public static Agenda getInstancia() {
    return AgendaHolder.instancia;
  }

  public void agregarElemento(Object elemento) {
    elementos.add(elemento);
  }

  public void eliminarElemento(Object elemento) {
    elementos.remove(elemento);
  }

  public List<Object> getElementos() {
    return elementos;
  }

  public List<Evento> filtrarEventos() {
    return elementos.stream()
        .filter(e -> e instanceof Evento)
        .map(e -> (Evento) e)
        .collect(Collectors.toList());
  }

  public List<Contacto> filtrarContactos() {
    return elementos.stream()
        .filter(e -> e instanceof Contacto)
        .map(e -> (Contacto) e)
        .collect(Collectors.toList());
  }

  @Override
  public String toString() {
    return "Agenda [Elementos=" + elementos + "]";
  }
}