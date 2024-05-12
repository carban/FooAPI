import React from "react";
import SyntaxHighlighter from 'react-syntax-highlighter';
import { atomOneDark } from 'react-syntax-highlighter/dist/esm/styles/hljs';

export default function EndpointContent({ payload, response }: { payload: object | null, response: object | null }) {
    return <div>
        <div className="items-center rounded-b-2xl p-5 lg:justify-around mb-3 bg-gray-800 space-y-8">
            <div className="items-center w-full">
                <p className="text-gray-400 mb-2">Payload</p>
                <SyntaxHighlighter language="json" style={atomOneDark} wrapLines={true} className="w-full mt-2">
                    {payload ? JSON.stringify(payload, null, 4) : "EMPTY"}
                </SyntaxHighlighter>
            </div>
            <div className="items-center w-full">
                <p className="text-gray-400 mb-2">Response</p>
                <SyntaxHighlighter language="json" style={atomOneDark} wrapLines={true} className="w-full mt-2">
                    {response ? JSON.stringify(response, null, 4) : "EMPTY"}
                </SyntaxHighlighter>
            </div>
        </div>
    </div>
}