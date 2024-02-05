package main

import (
  "bytes"
  "encoding/json"
  "fmt"
  "io"
  "io/ioutil"
  "net/http"

  "checkout/models"
)

func main() {
  // URL de la API de Mercado Pago para crear preferencias de checkout.
  url := "https://api.mercadopago.com/checkout/preferences"
  method := "POST"

  // Crear una estructura de preferencia con los detalles del producto,
  // las URLs de redirección y la URL de notificación.
  preference := models.CheckoutPreference{
    Items: []models.ItemModel{
      {
        Title:     "Producto de ejemplo",
        Quantity:  1,
        UnitPrice: 20000,
      },
    },
    BackUrls: models.UrlsModel{
      Success: "https://tu-sitio.com/success",
      Failure: "https://tu-sitio.com/failure",
      Pending: "https://tu-sitio.com/pending",
    },
    NotificationURL:   "https://tu-sitio.com/notif",
    ExternalReference: "1234",
  }

  // Convertir la estructura de preferencia a formato JSON para el cuerpo de la solicitud.
  bodyPref, err := json.Marshal(preference)
  if err != nil {
    panic(err.Error())
  }

  // Crear un cliente HTTP y una solicitud POST con el cuerpo JSON.
  client := &http.Client{}
  req, err := http.NewRequest(method, url, bytes.NewBuffer(bodyPref))
  if err != nil {
    fmt.Println(err)
    return
  }

  // Añadir headers necesarios, incluyendo la autorización con el token de acceso.
  req.Header.Add("Authorization", "Bearer APP_USR-MERCADO-PAGO-TOKEN")
  req.Header.Add("Content-Type", "application/json")

  // Realizar la solicitud y obtener la respuesta.
  res, err := client.Do(req)
  if err != nil {
    fmt.Println(err)
    return
  }

  defer func(Body io.ReadCloser) {
    _ = Body.Close()
  }(res.Body)

  // Leer el cuerpo de la respuesta.
  body, err := ioutil.ReadAll(res.Body)
  if err != nil {
    fmt.Println(err)
    return
  }

  // Convertir el cuerpo de la respuesta desde JSON a la estructura de datos adecuada.
  data, err := models.UnmarshalMakePayment(body)
  if err != nil {
    fmt.Println(err)
  }

  // Imprimir el punto de inicio del checkout (URL para redirigir al usuario para el pago).
  fmt.Println(data.InitPoint)
}
