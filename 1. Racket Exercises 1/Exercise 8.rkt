;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;
;;;;;;;;;;;;;;;Ejercicios 1 de Racket;;;;;;;;;;;;;;;;
;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;

;;Ejercicio 8:
(define (sub-conjunto? lista1 lista2)
  (cond
    ((null? lista1)
     #t)
    ((null? lista2)
     #f)
    ((= (car lista1) (car lista2))
     (sub-conjunto? (cdr lista1) lista2))
    (else
     (sub-conjunto? lista1 (cdr lista2))
     )
    )
  )