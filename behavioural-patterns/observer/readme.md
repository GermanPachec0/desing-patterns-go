### Observer pattern es un patron de comportamiento que nos permite definir un mecanismo de suscripción para notificar a multiples objetos sobre cualquier evento que están observando.

### Ejemplo: En un e-commerce los productos pueden estar fuera de stock cada cierto tiempo. hay clientes que quieren saber que producto en particular está fuera de stock. Para esto existen 3 soluciones. 
### 1 - el cliente checkea la disponibilidad del producto cada x cantidad de tiempo
### 2 - el e-commerce notifica todos los nuevos productos que estan en stock. 
### 3 - El cliente se suscribe a un porducto en particular el cual está interesado cuando el producto está en stock. Y también muchos clientes pueden estar suscritos al mismo producto.

### Subject, instancia que publica cuando un evento ocurre.
### Observer, Suscribe a un evento y se lo notifica cuando ocurre.

