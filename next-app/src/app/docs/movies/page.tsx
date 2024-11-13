import ServicePage from "@/app/components/docs/servicePage";
import data from "./data.json"
import { Metadata } from "next";

export const metadata: Metadata = {
   title: "Movies"
}

export default function Dummy() {
   return <ServicePage data={data} />
}