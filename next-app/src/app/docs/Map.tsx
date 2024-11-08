"use client"

// https://medium.com/@tomisinabiodun/displaying-a-leaflet-map-in-nextjs-85f86fccc10c
import { MapContainer, Marker, Popup, TileLayer } from "react-leaflet"
import "leaflet/dist/leaflet.css"
import "leaflet-defaulticon-compatibility"
import "leaflet-defaulticon-compatibility/dist/leaflet-defaulticon-compatibility.css"

export default function Map(props: any) {
  const { position, zoom } = props

  return <MapContainer className="map" center={position} zoom={zoom} scrollWheelZoom={false}>
    <TileLayer
      attribution='&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
      url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
    />
    <Marker position={[23.43, 90.24]}>
      <Popup>
          <b>Country: </b> Dakha <br /> <b>City: </b> Bangladesh
      </Popup>
    </Marker>
    <Marker position={[50.5, 4.2]}>
      <Popup>
      <b>Country: </b> Belgium <br /> <b>City: </b> Brussels
      </Popup>
    </Marker>
    <Marker position={[12.22, -1.31]}>
      <Popup>
      <b>Country: </b> Burkina Faso <br /> <b>City: </b> Ouagadougou
      </Popup>
    </Marker>
  </MapContainer>
}