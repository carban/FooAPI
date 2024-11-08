import { RiMapPinLine, RiBracesFill, RiPaintBrushLine, RiRobot2Line } from 'react-icons/ri'
import { GrGraphQl } from 'react-icons/gr';
import ServiceCard from './serviceCard';

function featuredServices(): object[] {
    return [
        {
            name: "Dummy Data",
            desc: "Here you can find dummy data of different topics",
            icon: <RiBracesFill className="w-8 h-8 text-primary" />,
            path: "/docs#dummy"
        },
        {
            name: "GraphQL",
            desc: "Here you can consume dummy data using GraphQL queries",
            icon: <GrGraphQl className="w-8 h-8 text-primary" />,
            path: "/docs#ql"
        },
        {
            name: "GeoJSON",
            desc: "Here you can find a feature collection of cities and (soon) layers of countries",
            icon: <RiMapPinLine className="w-8 h-8 text-primary" />,
            path: "/docs#geo"
        },
        {
            name: "Custom Mock",
            desc: "Here you can create an endpoint, custom payload, custom response",
            icon: <RiPaintBrushLine className="w-8 h-8 text-primary" />,
            path: "/docs#custom"
        },
        {
            name: "AI",
            desc: "AI data generated by your schema or data types",
            icon: <RiRobot2Line className="w-8 h-8 text-primary" />,
            path: "/docs#ai"
        },
    ];
}

export default function Featured() {
    const featured: any[] = featuredServices();
    return <div>
        <h2>Featured Services</h2>
        <div className="grid lg:grid-cols-5 md:grid-cols-3 sm:grid-cols-1 gap-5 py-5">
            {
                featured.map((f, index) => (
                    <ServiceCard key={index} service={f} />
                ))
            }
        </div>
    </div>
}