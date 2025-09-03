# SWYW-AUTH - Microservicio de Autenticación

## Descripción

SWYW-AUTH es un microservicio desarrollado en **Go** que maneja la autenticación y autorización del ecosistema SWYW. Este servicio está diseñado para proporcionar alta disponibilidad aprovechando las capacidades de concurrencia nativas de Go.

### Funcionalidades

- **Registro de usuarios** - Creación de nuevas cuentas
- **Login/Autenticación** - Validación de credenciales y generación de tokens
- **Recuperación de contraseña** - *(Funcionalidad en desarrollo futuro)*

### Ventajas de Go para Auth

- **Concurrencia nativa**: Manejo eficiente de múltiples requests simultáneos
- **Alta disponibilidad**: Goroutines permiten respuesta rápida bajo carga
- **Performance**: Compilación nativa para máximo rendimiento
- **Seguridad**: Tipado fuerte y manejo seguro de memoria

## Arquitectura

```
SWYW-AUTH/
├── src/               # Código fuente del microservicio
├── main.go           # Punto de entrada de la aplicación
├── Dockerfile        # Imagen Docker del servicio
├── go.mod            # Dependencias del proyecto
├── go.sum            # Verificación de dependencias
└── README.md         # Este archivo
```

## Requisitos Previos

- **Docker**: Versión 20.10 o superior
- **Red Docker**: Red `swyw` debe estar creada (ver README principal)

```bash
# Si no has creado la red, ejecuta:
docker network create swyw
```

## Configuración y Despliegue

### Paso 1: Levantar Base de Datos PostgreSQL

Antes de ejecutar el servicio de autenticación, es **obligatorio** levantar la base de datos PostgreSQL:

```bash
docker run -d --name s-postgres --network swyw \
  -e POSTGRES_PASSWORD=1234 \
  -p 5432:5432 postgres:13.22-trixie
```

**Parámetros importantes:**
- `--name s-postgres`: Nombre del contenedor (requerido para conectividad)
- `--network swyw`: Conecta a la red personalizada del proyecto
- `-e POSTGRES_PASSWORD=1234`: Contraseña de la BD (recuerda este valor)
- `-p 5432:5432`: Puerto expuesto para conexiones externas

### Paso 2: Configurar Estructura de Base de Datos

Una vez que PostgreSQL esté corriendo, debemos crear la tabla de usuarios:

```bash
# Acceder al contenedor de PostgreSQL
docker exec -it s-postgres bash

# Conectar a PostgreSQL como usuario postgres
psql -U postgres

# Crear la tabla de usuarios
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(150) UNIQUE NOT NULL,
    create_at TIMESTAMP DEFAULT NOW(),
    pass TEXT NOT NULL
);

# Salir de PostgreSQL
\q

# Salir del contenedor
exit
```

### Paso 3: Construir Imagen del Servicio Auth

```bash
# Navegar al directorio del servicio si te saliste en el paso anterior.
cd SWYW-AUTH

# Construir la imagen Docker
docker build -t auth-service .
```

### Paso 4: Ejecutar el Servicio de Autenticación

Tienes **dos opciones** para pasar las variables de entorno:

#### Opción A: Usando archivo .env (Recomendado)

Crear archivo `.env` en el directorio en caso tal no exista `SWYW-AUTH`:

```bash
# Archivo .env
DB_USER=postgres
DB_PASSWORD=1234
DB_HOST=s-postgres
DB_NAME=postgres
DB_PORT=5432
```

Ejecutar el contenedor:

```bash
docker run -d --name swyw-auth \
  -p 4000:4000 \
  --env-file .env \
  --network swyw \
  auth-service
```

#### Opción B: Variables de entorno directas

```bash
docker run -d --name swyw-auth \
  -p 4000:4000 \
  -e DB_USER=postgres \
  -e DB_PASSWORD=1234 \
  -e DB_HOST=s-postgres \
  -e DB_NAME=postgres \
  -e DB_PORT=5432 \
  --network swyw \
  auth-service
```

## Variables de Entorno

| Variable | Descripción | Valor por Defecto |
|----------|-------------|-------------------|
| `DB_USER` | Usuario de PostgreSQL | `postgres` |
| `DB_PASSWORD` | Contraseña de la BD | `1234` |
| `DB_HOST` | Host del contenedor de BD | `s-postgres` |
| `DB_NAME` | Nombre de la base de datos | `postgres` |
| `DB_PORT` | Puerto de PostgreSQL | `5432` |

## Verificación del Despliegue

### 1. Verificar que los contenedores estén corriendo

```bash
# Ver contenedores activos
docker ps

# Deberías ver algo similar a:
# CONTAINER ID   IMAGE           COMMAND                  CREATED         STATUS         PORTS                    NAMES
# xxxxxxxxx      auth-service    "./auth-service"         2 minutes ago   Up 2 minutes   0.0.0.0:4000->4000/tcp   swyw-auth
# xxxxxxxxx      postgres:13.22-trixie   "docker-entrypoint.s..."   5 minutes ago   Up 5 minutes   0.0.0.0:5432->5432/tcp   s-postgres
```

### 2. Verificar conectividad de red

```bash
# Verificar que ambos contenedores estén en la red swyw
docker network inspect swyw
```

### 3. Revisar logs del servicio

```bash
# Ver logs del servicio de autenticación
docker logs swyw-auth

# Ver logs en tiempo real
docker logs -f swyw-auth
```

### 4. Probar el servicio

```bash
# Ejemplo de health check (si está implementado)
curl http://localhost:4000/health

# Ejemplo de endpoint de registro (ajustar según tu API)
curl -X POST http://localhost:4000/register \
  -H "Content-Type: application/json" \
  -d '{"name":"Test User","email":"test@example.com","password":"testpass"}'
```

## Troubleshooting

### Problema: Servicio no puede conectar a la BD

**Solución**: Verificar que:
1. El contenedor `s-postgres` esté corriendo
2. Las variables de entorno sean correctas
3. Ambos contenedores estén en la red `swyw`

### Problema: Puerto 4000 ya está en uso

```bash
# Ver qué proceso usa el puerto
sudo lsof -i :4000

# Usar un puerto diferente
docker run -p 4001:4000 --env-file .env --network swyw auth-service
```

### Problema: Tabla de usuarios no existe

Repetir el Paso 2 para crear la estructura de BD.

## Limpieza

Para limpiar los contenedores cuando termines las pruebas:

```bash
# Detener contenedores
docker stop swyw-auth s-postgres

# Remover contenedores
docker rm swyw-auth s-postgres

# Remover imagen (opcional)
docker rmi auth-service
```

## Desarrollo

### Ejecutar en modo desarrollo

```bash
# Instalar dependencias
go mod download

# Ejecutar aplicación
go run main.go
```

### Reconstruir imagen tras cambios

```bash
docker build -t auth-service . --no-cache
```

---

**SWYW-AUTH** - Microservicio de autenticación desarrollado con Go para máxima concurrencia y disponibilidad.
