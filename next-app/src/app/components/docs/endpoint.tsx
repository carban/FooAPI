"use client"

import React from "react";
import EndpointContent from './endpointContent';
import { SlArrowDown, SlArrowUp } from "react-icons/sl";
import Link from "next/link";

export default function Endpoint({ status, path, desc, queries, payload, response, img}: { status: string, path: string, queries: string, desc: string, payload: Object | null, response: Object | null, img: string | null}) {

    let colors: any = {
        "GET": "bg-blue-400",
        "PUT": "bg-yellow-400",
        "PATCH": "bg-yellow-400",
        "POST": "bg-green-400",
        "DELETE": "bg-red-400",
        "GRAPHQL": "bg-pink-400"
    }

    let color: string = colors[status];

    const [toggle, setToggle] = React.useState(false);

    function handleToggle() {
        setToggle(!toggle);
    }

    return (
        <div className="p-2 mt-3 bg-gray-800">
            <div className="flex flex-col md:flex-row items-start md:items-center justify-between md:space-x-4 flex-wrap">
                {/* Method */}
                <h1 className={`${color} font-bold text-sm text-center rounded-3xl text-gray-900 mb-2 md:mb-0 min-w-[80px] md:min-w-[90px] w-full md:w-auto`}>
                    {status}
                </h1>

                {/* Endpoint Path */}
                <code className="text-white hover:underline mb-2 md:mb-0 w-64 flex-1">
                    <Link href={path}>{path}</Link>
                </code>

                {/* Description */}
                <h3 onClick={handleToggle} className="text-sm text-white flex-1">
                    {desc}
                </h3>

                {/* Expand/Collapse Button */}
                <button onClick={handleToggle} className="ml-0 md:ml-2 w-full md:w-auto">
                    {toggle ? (
                        <SlArrowUp className="w-6 h-6" />
                    ) : (
                        <SlArrowDown className="w-6 h-6" />
                    )}
                </button>
            </div>
            {toggle && <EndpointContent queries={queries} payload={payload} response={response} img={img}/>}
        </div>
    );
}