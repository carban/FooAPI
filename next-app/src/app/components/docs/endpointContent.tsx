import React from "react";
import SyntaxHighlighter from 'react-syntax-highlighter';
import { atomOneDark } from 'react-syntax-highlighter/dist/esm/styles/hljs';
import Image from 'next/image'
import img1 from '../../../../public/1.png'
import imgMaker from '../../../../public/foooo.png'

export default function EndpointContent({ queries, code, payload, response, img }: { queries: string | null, code: string, payload: object | null, response: object | null, img: string | null }) {
    let imgToShow;
    if (img == "1") {
        imgToShow = img1;
    } else {
        imgToShow = imgMaker
    }
    return <div>
        {
            (payload || response) || img ? <div className="items-center rounded-b-2xl p-5 lg:justify-around mb-3 bg-gray-800 space-y-8">
                <div className="items-center w-full">
                    {
                        queries ? <div className="mb-2 mt-2">
                            <p className="text-gray-400 mb-2 mt-2">Queries Available</p>
                            <SyntaxHighlighter language="json" style={atomOneDark} wrapLines={true}>
                                {queries}
                            </SyntaxHighlighter>
                        </div> : true
                    }
                    <div className="mb-2 mt-2">
                        <p className="text-gray-400 mb-2 mt-2">Usage Example</p>
                        <SyntaxHighlighter language="javascript" style={atomOneDark} wrapLines={true}>
                            {code}
                        </SyntaxHighlighter>
                    </div>
                    <div className="mb-2 mt-2">
                        <p className="text-gray-400 mb-2 mt-2">Payload Example</p>
                        <SyntaxHighlighter language="json" style={atomOneDark} wrapLines={true}>
                            {payload ? JSON.stringify(payload, null, 4) : "EMPTY"}
                        </SyntaxHighlighter>
                    </div>
                    <div className="mb-2 mt-2">
                        <p className="text-gray-400 mb-2 mt-2">Response Example</p>
                        {
                            !img ?
                                <SyntaxHighlighter language="json" style={atomOneDark} wrapLines={true}>
                                    {response ? JSON.stringify(response, null, 4) : "EMPTY"}
                                </SyntaxHighlighter>
                                : <Image src={imgToShow} alt={""} width={468} height={312}></Image>
                        }
                    </div>
                </div>
            </div> : true
        }
    </div>
}