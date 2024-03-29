### Strategy pattern

### Patron de diseño de comportamiento que permite cambiar el comportamiento de una aplicación en tiempo de ejecución, donde el objeto original, tiene una referencia a un strategy object y este delega su comportamiento a la strategy linkeada y en order que el objeto va cambiando o el contexto cambia, otros objetos pueden cambiar la estrategia linkeada.

### Algunos ejemplos: Supongamos que tenemos un cliente que consume 3 tipos de bases de datos, y queremos instanciar 3 tipos de conecciones distintas, postgres, mysql y mongo. el patron strategy permite que con solo una interfaz esto sea posible.
 
### ejemplo de implementación: Supongamos que estamos construyendo una memoria cache con un tamaño limitado y cuando sea que esta alcance su maximo tamaño, queremos liberar espacio, esto se puede lograr mediante algoritmos como
### LRU: Least Recently Used - Remover la entrada que ha sido menos utilizada recientemente. 
### FIFO: First In First Out . Remvoer la entrada que ha sido creada primero.
### LFU: Least Frequently Used - Remover la entrada que se utiliza con menos frecuencia.

### el problema está en que queremos desacoplar nuestra clase CACHE de los algoritmos que pueden cambiar en tiempo de ejecución. Y También la clase CACHE no deberia cambiar cuando un nuevo algoritmo es agregado. El patron strategy propone que cada algoritmo tieen que tener su propia clase. y cada clase debe seguir la misma interfaz.