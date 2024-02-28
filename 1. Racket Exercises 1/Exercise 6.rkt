;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;
;;;;;;;;;;;;;;;Ejercicios 1 de Racket;;;;;;;;;;;;;;;;
;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;

;;Ejercicio 6:
(define (merge lista1 lista2)
  (cond
    ((null? lista1)
     lista2)
    ((null? lista2)
     lista1)
    ((< (car lista1) (car lista2))
     (cons (car lista1)(merge (cdr lista1) lista2)))
    (else
     (cons (car lista2) (merge lista1 (cdr lista2)))
     )
    )
  )