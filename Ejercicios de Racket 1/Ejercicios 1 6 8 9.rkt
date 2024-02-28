;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;
;;;;;;;;;;;;;;;Ejercicios 1 de Racket;;;;;;;;;;;;;;;;
;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;

;;Ejercicio 1:
(define (int Cap I N)
  (if (= N 0)
      Cap                                 ;;entra al if
      (int (+ Cap (* Cap I)) I (- N 1))   ;;else
      )
  )

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
;;Ejercicio 9:
(define (eliminar_elemento E L)
  (apply append
         (map (lambda (x)
                (if (= x E)
                    '()         ;;entra al if
                    (list x))   ;;else
                )
              L)
         )
  )