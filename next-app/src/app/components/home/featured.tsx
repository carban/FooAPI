import { RiMapPinLine, RiBracesFill, RiPaintBrushLine, RiRobot2Line } from 'react-icons/ri'
import { GrGraphQl } from 'react-icons/gr';
import ServiceCard from './serviceCard';

function featuredServices(): object[] {
    return [
        {
            name: "Dummy Data",
            desc: "lorem",
            icon: <RiBracesFill className="w-8 h-8 text-primary" />,
            path: "/docs"
        },
        {
            name: "GraphQL",
            desc: "Go to this step by step guideline process on how to certify for your weekly benefits:",
            icon: <GrGraphQl className="w-8 h-8 text-primary" />,
            path: "/docs"
        },
        {
            name: "GeoJSON",
            desc: "Go to this step by step guideline process on how to certify for your weekly benefits:",
            icon: <RiMapPinLine className="w-8 h-8 text-primary" />,
            path: "/docs"
        },
        {
            name: "Custom Mock",
            desc: "Go to this step by step guideline process on how to certify for your weekly benefits:",
            icon: <RiPaintBrushLine className="w-8 h-8 text-primary" />,
            path: "/docs"
        },
        {
            name: "AI",
            desc: "Go to this step by step guideline process on how to certify for your weekly benefits:",
            icon: <RiRobot2Line className="w-8 h-8 text-primary" />,
            path: "/docs"
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