
Refactorizando a caso de uso

En las anteriores lecciones vimos cómo aplicar el principio de inversión de dependencia mediante la inyección de la misma para desacoplar nuestro controlador (o handler) HTTP de la capa de persistencia y de su implementación específica (MySQL).


Sin embargo, si creémos que nuestra aplicación puede seguir creciendo, y que potencialmente:


Nuestros casos de usos podrán ser ejecutados desde diferentes puntos de entrada (interfaz de línea de comandos, event handlers, etc).
Y cada uno de ellos podrá llevar asociada una lógica de aplicación (diferentes acciones como determinadas validaciones a nivel de aplicación, publicación de eventos, etc) transversales a todos los puntos de entrada.

Entonces, lo mejor será que tratemos de centralizar dicho código en una pieza independiente, que posteriormente podrá ser inyectada en cualquiera de los puntos de entrada, sin necesidad de tener que duplicar dicho código.


Nomenclatura y estructura de carpetas

Con la refactorización mencionada, ahora se nos abre la puerta a la tercera y última de las habituales capas que comúnmente forman la Arquitectura Hexagonal: la capa de aplicación.


A nivel conceptual estará situada entre nuestra infraestructura y nuestro dominio, lo cuál lo representaremos en nuestra estructura de carpetas añadiendo una nueva carpeta en el nivel dónde hasta ahora solo teníamos nuestro dominio. De este modo nos quedaría:


raíz: dominio.
carpetas a primer nivel (excepto platform): aplicación.
carpetas a segundo nivel (dentro de platform): infraestructura.

Además, la separación de esta nueva capa (es decir, las carpetas que la conformarán), será a nivel de grupo de operaciones (creating, fetching).
De este modo, podremos crear los correspodientes servicios de aplicación o casas de uso (p. ej: creación de curso: creating.CourseService), dejando clara la intención de cada uno de ellos.


Referencias adicionales

Como venimos diciendo, las adaptaciones de Arquitectura Hexagonal a Go y la correspondiente estructura de carpetas que se propone en este curso es algo que tiene nuestro toque personalizado pero que no nos hemos sacado de la manga, sinó que ha sido el resultado de horas de estudio de las diferentes propuestas de la comunidad y de contrastación con las mismas.


En este caso específico, una de las referencias de la comunidad más populares es el repositorio GoDDD en el que Marcus Olsson refactorizó a Go la aplicación Java que desarrolló Citerus en estrecha colaboración con Eric Evans, basada en los ejemplos de su libro. En el mismo podemos encontrar un patrón muy similar a la hora de organizar la capa de aplicación (tracking, booking, etc), pero sin esa separación física con la capa de infraestructura (platform).


