import Link from "next/link";

export default function Foot() {
    return <div className="w-full h-10 flex items-center justify-around">
        <Link href="/">
            <p className="text-lg font-semibold text-primary">f4k3
                <b>81t5.10</b>
            </p>
        </Link>
    </div>
}