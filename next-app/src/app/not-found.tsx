import Link from "next/link";
import Navbar from "./components/navbar";
import { TbError404 } from "react-icons/tb";

export default function NotFound() {
    return <div>
        <Navbar />
        <div className="container mx-auto px-5 py-10 text-center mb-64">
            <h1 className="mb-10 text-3xl font-semibold tracking-tight text-white">Foops! there was a problem</h1>
            <center>
                <TbError404 className="w-32 h-32 text-primary" />
            </center>
            <p className="text-white text-lg">
                We could not find the page you were looking for.
            </p>
            <p className="text-white text-lg">
                Go back to <Link href="/">Main Page</Link>
            </p>
        </div>
    </div >
}