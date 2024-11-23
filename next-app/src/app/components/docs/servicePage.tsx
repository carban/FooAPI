import Endpoint from "@/app/components/docs/endpoint";

export default function ServicePage({ data }: { data: { title: string, desc: string, endpoints: any[] } }) {
    const { title, desc, endpoints } = data;
    return <div>
        <h1 className="text-3xl mb-5 text-white">{title}</h1>
        <p className="text-white">{desc}</p>
        <div className="border-b-2 border-gray-700 m-5"></div>
        {
            endpoints.map((e, index) => (
                <Endpoint key={index} status={e.status} path={e.path} desc={e.desc} queries={e.queries} code={e.code} payload={e.payload} response={e.response} img={e.img} />
            ))
        }
    </div>
}