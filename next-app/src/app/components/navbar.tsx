"use client"

import React from "react";
import Link from "next/link"

const Logo = () => <div>
    <p className="self-center text-3xl font-semibold text-white">Foo
        <b className="text-primary">Api</b>
    </p>
</div>

export default function Navbar() {

    const [toggle, setToggle] = React.useState(false);

    function handleToggle() {
        setToggle(!toggle);
    }

    return <nav className='bg-gray-900'>
        <div className="max-w-full flex flex-wrap items-center justify-between mx-8 p-4">
            <Link href="/">
                <Logo />
            </Link>
            <button onClick={handleToggle} data-collapse-toggle="navbar-default" type="button" className="inline-flex items-center p-2 w-10 h-10 justify-center text-sm text-gray-500 rounded-lg md:hidden hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-gray-200 text-gray-400 hover:bg-gray-700 focus:ring-gray-600" aria-controls="navbar-default" aria-expanded="false">
                <span className="sr-only">Open main menu</span>
                <svg className="w-5 h-5" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 17 14">
                    <path stroke="currentColor" d="M1 1h15M1 7h15M1 13h15" />
                </svg>
            </button>
            <div className={`${toggle ? "" : "hidden"} w-full md:block md:w-auto`} id="navbar-default">
                <ul className="font-medium flex flex-col md:flex-row p-4 md:p-0 mt-4 border border-gray-100 rounded-lg bg-gray-50 md:space-x-8 md:mt-0 md:border-0 bg-gray-800 md:bg-gray-900 border-gray-700">
                    <li>
                        <Link href="/" className="block py-2 pl-3 pr-4 text-white rounded md:bg-transparent md:p-0 text-primary " aria-current="page">Home</Link>
                    </li>
                    <li>
                        <Link href="/docs" className="block py-2 pl-3 pr-4 text-gray-900 rounded hover:bg-gray-100 md:hover:bg-transparent md:border-0 md:hover:text-blue-700 md:p-0 text-white md:hover:text-blue-500 hover:bg-gray-800 hover:text-white md:hover:bg-transparent">Docs</Link>
                    </li>
                    <li>
                        <Link href="/about" className="block py-2 pl-3 pr-4 text-gray-900 rounded hover:bg-gray-100 md:hover:bg-transparent md:border-0 md:hover:text-blue-700 md:p-0 text-white md:hover:text-blue-500 hover:bg-gray-800 hover:text-white md:hover:bg-transparent">About</Link>
                    </li>
                    <li>
                        <a href="https://www.buymeacoffee.com/carban" target="_blank"><p className="md:text-primary">Buy me a coffe!</p></a>
                    </li>
                </ul>
            </div>
        </div>
    </nav>
}