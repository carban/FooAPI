export default {
    title: "This is the Photos data collection",
    desc: "Lorem ipsum, dolor sit amet consectetur adipisicing elit. Fuga, nam error nihil facilis dolorum consequuntur facere aut porro totam, labore blanditiis voluptatum tempora odio! Molestiae hic quasi harum eius dolore.",
    endpoints: [
        {
            status: "GET",
            path: "/api/mit/edu",
            payload: null,
            response: {},
            desc: "Lorem ipsum dolor sit amet consectetur adipisicing elit. Commodi corrupti dolor"
        },
        {
            status: "GET",
            path: "/api/mit/edu/:id",
            payload: {},
            response: null,
            desc: "Lorem ipsum dolor sit amet consectetur adipisicing elit. Commodi corrupti dolor"
        },
        {
            status: "PUT",
            path: "/api/mit/edu/:id",
            payload: [
                {
                    hola: 2
                },
                '99 more elements...'],
            response: {},
            desc: "Lorem ipsum dolor sit amet consectetur adipisicing elit. Commodi corrupti dolor"
        },
        {
            status: "PATCH",
            path: "/api/mit/edu/:id",
            payload: {},
            response: { hola: 2 },
            desc: "Lorem ipsum dolor sit amet consectetur adipisicing elit. Commodi corrupti dolor"
        },
        {
            status: "POST",
            path: "/api/mit/edu",
            payload: {},
            response: {},
            desc: "Lorem ipsum dolor sit amet consectetur adipisicing elit. Commodi corrupti dolor"
        },
        {
            status: "DELETE",
            path: "/api/mit/edu/:id",
            payload: {},
            response: {},
            desc: "Lorem ipsum dolor sit amet consectetur adipisicing elit. Commodi corrupti dolor"
        },
    ]
}