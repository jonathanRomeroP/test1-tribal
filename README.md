# API REST (Golang, Docker, Docker-compose) 😁.

#  <font color='red'>Requisitos</font>
* Instale la última versión estable de [Docker](https://docs.docker.com/install/linux/docker-ce/ubuntu/#install-docker-ce-1)
* Instale [Docker Compose](https://docs.docker.com/compose/install/#install-compose)

#  <font color='red'>Instalación</font>
* Explorar la raíz del repositorio
* Construye las imágenes
    - `docker-compose build`
* Inicie los containers
    - `docker-compose up -d`

Después de iniciar los contenedores, puede probar la Api en (proxy):
```url
http://localhost/api/
```

#  <font color='red'>Pasos Para Consumir La Api</font>

* Registrarse
    -   ```url 
        http://localhost/user/signup
        ```
    -   ```
        {
            "Name": "Jonathan",
            "Email": "jhoropertuz@gmail.com",
            "Password": "12345"
        }
    ```

* Login (tendras el token necesario para consultar las canciones)
    -   ```url
        http://localhost/user/login
        ```
    -   ```
            {
                "Email": "jhoropertuz@gmail.com",
                "Password": "12345"
            }
        ```

* Filtrar Canciones
    - 🔴 Es necesario el token que la api te da al logiarte
    - 🔴 Debes pasa por el Header dicho Token con la  key "token"
        ![Image text](https://github.com/jonathanRomeroP/test1-tribal/blob/devProyect/public/img/token.png)
    -  El filtrado se realiza con el siguiente endpoin
    ```url 
    http://localhost/api/song/filter/?name=string&artist=string
    ```

    - Puedes pasar por query params
        - name (nombre de la cancion)
        - artist (artista musical)
        - album (album del artista)

#  <font color='red'>Rutas Activas</font>

![Image text](https://github.com/jonathanRomeroP/test1-tribal/blob/devProyect/public/img/rutas.png)