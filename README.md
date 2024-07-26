# Simulador de Procesamiento de Pedidos Concurrente en Go

¡Bienvenido al proyecto del Simulador de Procesamiento de Pedidos Concurrente en Go! Este repositorio acompaña al post "NIVEL 4: CREEMOS UN SIMULADOR DE PROCESAMIENTO DE PEDIDOS CONCURRENTE EN GOLANG" de nuestro blog gognition.pro donde aprendimos a crear un simulador de pedidos utilizando concurrencia en Go.

## 📋 Descripción

Este proyecto implementa un simulador de procesamiento de pedidos concurrente usando Go. Simula la generación de pedidos, verificación de inventario, procesamiento de pagos y preparación de envíos, todo de manera concurrente. Introduce conceptos avanzados de programación en Go como gorrutinas, canales, y manejo seguro de concurrencia.

## 🚀 Comenzando

Para utilizar este simulador, asegúrate de tener Go instalado en tu sistema.

### Prerrequisitos

- Go (versión 1.15 o superior recomendada)

### Instalación

1. Haz un fork de este repositorio haciendo clic en el botón "Fork" en la parte superior derecha de esta página.

2. Clona tu fork a tu máquina local:

```bash
git clone https://github.com/TU_USUARIO/gognition-nivel4-simulador-pedidos.git
```

## 💻️ Uso

Para ejecutar el simulador de pedidos:
```bash
go run main.go
```

El programa simulará el procesamiento de pedidos durante 30 segundos, mostrando mensajes como:
```bash
Inventario verificado para el pedido 1
Pago procesado para el pedido 1
Envío preparado para el pedido 1
Inventario insuficiente para el pedido 2
Pago procesado para el pedido 3
...
```

## 🏗️ Compilación (Build)

Para compilar el proyecto y crear un ejecutable, puedes usar los siguientes comandos dependiendo de tu sistema operativo objetivo:
```bash
go build -o simulador-pedidos
```

Para Windows (desde Linux o macOS):
```bash
GOOS=windows GOARCH=amd64 go build -o simulador-pedidos.exe
```

Para macOS (desde Windows o Linux):
```bash
GOOS=darwin GOARCH=amd64 go build -o simulador-pedidos-mac
```

Para Linux (desde Windows o macOS):
```bash
GOOS=linux GOARCH=amd64 go build -o simulador-pedidos-linux
```

Después de la compilación, puedes ejecutar el programa usando:
```
./simulador-pedidos  # En Linux o macOS
simulador-pedidos.exe  # En Windows
```

## Visita gognition.pro https://www.gognition.pro/