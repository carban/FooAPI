import ServicePage from "@/app/components/docs/servicePage";
import data from "./data.json"
import Link from "next/link";
import { Metadata } from "next";

export const metadata: Metadata = {
   title: "Cities"
}

export default function Dummy() {
   return <div>
      <ServicePage data={data} />
      <code className="text-red-500 text-base hover:underline mb-2 md:mb-0 w-64 flex-1">
         <Link href="/docs#geo"> /docs#geo</Link>
      </code>
   </div>
}