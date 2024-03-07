;Facturas 'Quemadas' de la forma nombreProducto, cant, p/u
(define primerFactura '(("Azúcar" 3 1000)
                        ("Café" 2 2800))
  )

(define segundaFactura '(("Arroz" 2 1800)
                         ("Azúcar" 1 1000)
                         ("Leche" 5 1200))
  )

;;Ejercicio1
;;devuelve el monto total del impuesto en la factura
(define (impuestoTotalFactura factura umbral)
  (cond
    ((null? factura)
     0)
    ((> (caddar factura) umbral)
     (+ (* (caddar factura)(cadar factura) 0.13)
        (impuestoTotalFactura (cdr factura) umbral)))
    (else
     (+ 0 (impuestoTotalFactura (cdr factura) umbral)))
    )
  )
;;devuelve el monto total de la factura sin ningun impuesto
(define (subtotalFactura factura)
  (if (null? factura)
      0                                            ;;if
      (+ (* (caddar factura) (cadar factura))      ;;else
         (subtotalFactura (cdr factura)))))

;;CASOS DE PRUEBA
;;1
(display "\nPara la primerFactura con un umbral de 2000 tenemos que: \n El Subtotal es de: ")
(display (subtotalFactura primerFactura))
(display "\n El impuesto es de: ")
(display (impuestoTotalFactura primerFactura 2000))

;;2
(display "\nPara la segundaFactura con un umbral de 1100 tenemos que: \n El Subtotal es de: ")
(display (subtotalFactura segundaFactura))
(display "\n El impuesto es de: ")
(display (impuestoTotalFactura segundaFactura 1100))

;;3
(display "\nPara la segundaFactura con un umbral de 1500 tenemos que: \n El Subtotal es de: ")
(display (subtotalFactura segundaFactura))
(display "\n El impuesto es de: ")
(display (impuestoTotalFactura segundaFactura 1500))