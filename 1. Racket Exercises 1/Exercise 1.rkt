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