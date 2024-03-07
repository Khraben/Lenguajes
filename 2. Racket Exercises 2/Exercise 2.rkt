;;Ejercicio 2
;;determina si la subcadena contiene el argumento
(define (filtrarSubcadena subcadena argumento)
  (not
   (equal? #f (regexp-match (regexp (regexp-quote argumento)) subcadena))))

;;funcion principal
(define (filtrarPorArgumento lista argumento)
  (filter
   (lambda (cadena) (filtrarSubcadena cadena argumento))
   lista))

(display "\nTenemos la lista: ")
(display "('la casa', 'el perro' , 'pintando la cerca')")
(display " y vamos a filtrar por el argumento: la \n Resultado: \n")
(display (filtrarPorArgumento '("la casa" "el perro" "pintando la cerca") "la"))


