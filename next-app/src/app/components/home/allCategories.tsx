import { RiBracesFill, RiUser3Line, RiListOrdered2, RiShoppingCart2Line, RiAlignLeft } from 'react-icons/ri'
import { LiaCommentSolid } from 'react-icons/lia'
import { BiWorld, BiSolidCity, BiPhotoAlbum, BiCameraMovie, BiMusic } from 'react-icons/bi'
import ServiceCard from './serviceCard';

function allServices(): object[] {
    return [
        {
            name: "Users",
            icon: <RiUser3Line className="w-8 h-8 text-emerald-500" />,
            path: "/docs/users"
        },
        {
            name: "Todos",
            icon: <RiListOrdered2 className="w-8 h-8 text-emerald-500" />,
            path: "/docs/todos"
        },
        {
            name: "Products",
            icon: <RiShoppingCart2Line className="w-8 h-8 text-emerald-500" />,
            path: "/docs/products"
        },
        {
            name: "Posts",
            icon: <RiAlignLeft className="w-8 h-8 text-yellow-400" />,
            path: "/docs/posts"
        },
        {
            name: "Comments",
            icon: <LiaCommentSolid className="w-8 h-8 text-yellow-400" />,
            path: "/docs/comments"
        },
        {
            name: "Countries",
            icon: <BiWorld className="w-8 h-8 text-indigo-500" />,
            path: "/docs/countries"
        },
        {
            name: "Cities",
            icon: <BiSolidCity className="w-8 h-8 text-indigo-500" />,
            path: "/docs/cities"
        },
        {
            name: "Photos",
            icon: <BiPhotoAlbum className="w-8 h-8 text-pink-300" />,
            path: "/docs/photos"
        },
        {
            name: "Movies",
            icon: <BiCameraMovie className="w-8 h-8 text-orange-400" />,
            path: "/docs/movies"
        },
        {
            name: "Songs",
            icon: <BiMusic className="w-8 h-8 text-orange-400" />,
            path: "/docs/songs"
        },
    ];
}

export default function AllCategories() {
    const services: any[] = allServices();
    return <div>
        <h2>All Categories</h2>
        <div className="grid lg:grid-cols-5 md:grid-cols-3 sm:grid-cols-2 gap-5 py-6">
            {
                services.map((s, index) => (
                    <ServiceCard key={index} service={s} />
                ))
            }
        </div>
    </div>
}