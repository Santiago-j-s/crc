# Trabajo Final de Redes: Análisis de Capacidad de Detección de CRC

### Bienvenido

El presente manual de usuario pretende ser una guia para aprender a utilizar este software.

Esta aplicación fue realizada como trabajo final para la asignatura Redes y Transmision de Datos correspondiente a la carrera
de Licenciatura de Informática dictada en la Universidad de la Patagonia San Juan Bosco (UNPSJB).

## Instalación

Es necesario tener instalado go y seteada la variable de entorno GOPATH

https://golang.org/doc/install

```
go get "github.com/Santiago-j-s/crc"
```

Dentro de la carpeta crc

```
cd server
go build server.go
```

## Iniciar WebApp

* Iniciar el ejecutable que se encuentra en la carpeta server

* Ingresar a http://localhost:6080

**La aplicación permite calcular el CRC de un mensaje a partir de un polinomio dado, y también permite analizar la capacidad de detección de un polinomio**

## Cálculo de CRC

![][CRC]

**En la pestaña CRC:**

* Ingresar un número binario cuya cantidad de bits sea múltiplo de 8 en el campo Mensaje

* Ingresar un número binario de 8 bits en el campo Polinomio

* Click en Calcular

Bajo Resultado se puede ver el CRC calculado

## Análisis de Capacidad de Detección

![][Analisis]

**En la pestaña Análisis:**

* Ingresar un número binario de 8 bits en el campo Polinomio

* Click en Calcular

### Interpretar Resultados

Para cada cantidad de bits alterados se detecta cuantos cambios posibles pasarán inadvertidos por el polinomio ingresado.

Esto significa que, por ejemplo, en la imágen mostrada si al transmitir un mensaje junto a su CRC de acuerdo al polinomio ingresado se alteran dos bits, de entre todas las alteraciones posibles hay 9 que pasarán inadvertidas.

Un buen polinomio se distingue por la carácteristica de que detecta todas las alteraciones posibles de 3 bits o menos.

Sin importar el polinomio, los errores de un bit siempre son detectados, por eso no se muestran en la sección de resultados.

**Ejemplos de buenos polinomios para CRC-8:**

* 0xCF 11001111
* 0x4d 01001101
* 0x1d 00011101

(Obtenidos del zoo de polinomios del Philip Koopman https://users.ece.cmu.edu/~koopman/crc/crc8.html)

## TO DO

* Calcular todas las combinaciones posibles para cada una de las cantidades de bits alterados para poder ofrecer probabilidad de detección por cantidad de bits alterados


[CRC]: ./docs/CRC.png
[Analisis]: ./docs/Análisis.png