## Patrón state

### el patron state es un patron de comportamiento que permite a un objeto cambiar su comportamiento cuando su estado interno cambia. 

### ejemplo: Para el ejemplo, se tiene una maquina expendedora que puede pasar por 4 estados diferentes: hasItem, noItem, itemRequested, hasMoney. Y hay cuatro tipo de acciones que puede ejecutar la maquina expendedora: Select the item, Add the item, Insert money, Dispense item. 

### para nuestro ejemplo este patron de diseño debería ser utilizado cuando una request llega y el objeto necesita cambiar su estado actual. en el ejemplo de la maquina expendora puede pasar a diferentes estado, por ejemplo: si la maquia está en itemRequested, una vez que la "insert money" ocurre, la maquina se mueve al estado hasMoney.

### depende del estado actual, la maquina se puede comportar de diferentes maneras para la misma request. por ejemplo si el usuario quiere un item y la maquina está con el estado hasItem se lo dará pero se rejecteara la request sí está en el estado noItem