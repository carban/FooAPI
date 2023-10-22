"use client"

import React from "react";
import EndpointContent from './endpointContent';
import { SlArrowDown, SlArrowUp } from "react-icons/sl";
import Link from "next/link";

export default function Endpoint({ status, path, desc, payload, response }: { status: string, path: string, desc: string, payload: Object | null, response: Object | null }) {

    let colors: any = {
        "GET": "bg-blue-400",
        "PUT": "bg-yellow-400",
        "PATCH": "bg-yellow-400",
        "POST": "bg-green-400",
        "DELETE": "bg-red-400",
    }

    let color: string = colors[status];

    const [toggle, setToggle] = React.useState(false);

    function handleToggle() {
        setToggle(!toggle);
    }

    return <div>
        <div className="flex flex-col xl:flex-row lg:flex-row md:flex-col items-center p-2 lg:justify-between mt-3 bg-gray-800">
            <h1 className={`${color} font-bold text-sm text-center rounded-3xl w-24 text-gray-900 lg:mr-14`}>{status}</h1>
            <code className="lg:w-60 hover:underline">
                <Link href={path} >{path}</Link>
            </code>
            <h3 className="text-sm">
                {desc}
            </h3>
            <button onClick={handleToggle}>
                {
                    toggle ? <SlArrowUp className="w-6 h-6" /> : <SlArrowDown className="w-6 h-6" />
                }
            </button>
        </div>
        {
            toggle ? <EndpointContent payload={payload} response={response} /> : false
        }
    </div>
}