"use client"

import Link from "next/link"
import Navbar from "../navbar"
import { usePathname } from 'next/navigation'
import { RiAlignLeft, RiListOrdered2, RiPaintBrushLine, RiShoppingCart2Line, RiUser3Line } from "react-icons/ri";
import { LiaCommentSolid } from "react-icons/lia";
import { BiCameraMovie, BiMusic, BiPhotoAlbum, BiSolidCity, BiWorld } from "react-icons/bi";

function allServices(): object[] {
    return [
        {
            name: "Users",
            icon: <RiUser3Line className="w-6 h-6 text-emerald-500" />,
            path: "/docs/users"
        },
        {
            name: "Todos",
            icon: <RiListOrdered2 className="w-6 h-6 text-emerald-500" />,
            path: "/docs/todos"
        },
        {
            name: "Products",
            icon: <RiShoppingCart2Line className="w-6 h-6 text-emerald-500" />,
            path: "/docs/products"
        },
        {
            name: "Posts",
            icon: <RiAlignLeft className="w-6 h-6 text-yellow-400" />,
            path: "/docs/posts"
        },
        {
            name: "Comments",
            icon: <LiaCommentSolid className="w-6 h-6 text-yellow-400" />,
            path: "/docs/comments"
        },
        {
            name: "Photos",
            icon: <BiPhotoAlbum className="w-6 h-6 text-pink-300" />,
            path: "/docs/photos"
        },
        {
            name: "Movies",
            icon: <BiCameraMovie className="w-6 h-6 text-orange-400" />,
            path: "/docs/movies"
        },
        {
            name: "Songs",
            icon: <BiMusic className="w-6 h-6 text-orange-400" />,
            path: "/docs/songs"
        },
        {
            name: "Countries",
            icon: <BiWorld className="w-6 h-6 text-indigo-500" />,
            path: "/docs/countries"
        },
        {
            name: "Cities",
            icon: <BiSolidCity className="w-6 h-6 text-indigo-500" />,
            path: "/docs/cities"
        },
        {
            name: "Custom",
            icon: <RiPaintBrushLine className="w-6 h-6 text-primary" />,
            path: "/docs/mock"
        },
    ];
}

export default function VerticalNavbar({ children }: any) {
    const pathname = usePathname();
    const services: any[] = allServices();
    return <main className="min-h-screen w-full bg-gray-900" x-data="layout">
        {/* <header 
        </header> */}
        <Navbar />
        <div className="flex">
            <aside className="flex flex-col w-60 mx-5 py-4 border-r rtl:border-r-0 rtl:border-l bg-gray-900 border-gray-700" x-show="asideOpen">
                {
                    services.map((s, index) => (
                        <Link key={index} href={s.path} className={`flex items-center space-x-1 rounded-md px-2 py-3 text-white ${pathname === s.path ? 'bg-gray-800' : 'hover:bg-gray-800'}`}>
                            {s.icon}
                            <span>{s.name}</span>
                        </Link>
                    ))
                }
            </aside>
            <div className="container max-w-full lg:px-15 md:px-10">
                {children}
            </div>
        </div>
    </main>
}