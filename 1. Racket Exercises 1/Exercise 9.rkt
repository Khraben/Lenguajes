;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;
;;;;;;;;;;;;;;;Ejercicios 1 de Racket;;;;;;;;;;;;;;;;
;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;

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