import Link from "next/link";

export default function Foot() {
    return <div className="w-full h-10 flex items-center justify-around bg-gray-900">
        <Link href="/">
            <p className="text-lg font-semibold text-primary">
                <b>F004P1.C0M</b>
            </p>
        </Link>
    </div>
}