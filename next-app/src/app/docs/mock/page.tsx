'use client'

import { useState } from "react";

import Editor from 'react-simple-code-editor';
import { highlight, languages } from 'prismjs';
import 'prismjs/themes/prism-tomorrow.css'; //Example style, you can use another

export default function Custom() {
    const [responseBody, setResponseBody] = useState('{\n "foo": "Fooooooo",\n "hello": [123, true] \n}');
    const [methodName, setmethodName] = useState("GET");
    const [statusCode, setstatusCode] = useState(200);
    const [data, setData] = useState(null);
    const [tooManyError, setTooManyError] = useState(false);
    const [expiredError, setExpiredError] = useState(false);
    const [editMode, setEditMode] = useState(false);


    const handleMethodNameChange = (event: any) => {
        setmethodName(event.target.value);
    };

    const handleStatusCodeChange = (event: any) => {
        setstatusCode(event.target.value);
    };

    const handleClick = async () => {
        if (!editMode) {
            // console.log(JSON.stringify({
            //     method: methodName,
            //     response_code: statusCode,
            //     response_body: responseBody
            // }))
            try {
                JSON.parse(responseBody);
            } catch (e) {
                alert("Check the JSON Format");
                return
            }
            try {
                const response = await fetch('/custom', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify({
                        method: methodName,
                        response_code: Number(statusCode),
                        response_body: responseBody
                    })
                });

                if (!response.ok) {
                    if (response.status == 429) {
                        setTooManyError(true);
                    }
                    throw new Error(`HTTP error! status: ${response.status}`);
                }

                const jsonData = await response.json();
                // console.log(jsonData)
                setData(jsonData.hash);
                setEditMode(true);
            } catch (error) {
                console.log('Error fetching data:', error);
            }
        } else {
            // console.log(JSON.stringify({
            //     method: methodName,
            //     response_code: statusCode,
            //     response_body: responseBody
            // }))
            try {
                JSON.parse(responseBody);
            } catch (e) {
                alert("Check the JSON Format");
                return
            }
            try {
                const response = await fetch('/editcustom/' + data, {
                    method: 'PUT',
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify({
                        method: methodName,
                        response_code: Number(statusCode),
                        response_body: responseBody
                    })
                });

                if (!response.ok) {
                    if (response.status == 404) {
                        setExpiredError(true);
                    }
                    throw new Error(`HTTP error! status: ${response.status}`);
                }

                alert("Edited!")

            } catch (error) {
                console.error('Error fetching data:', error);
            }
        }

    };

    return (
        <div>
            <h1 className="text-3xl mb-5 text-white">Custom Endpoint Maker</h1>
            <p className="text-white text-lg">Here you can create custom endpoints in case you need a dummy endpoint to test specific data. The endpoint expires after 1hr so you must to create a new one.</p>
            <div className="grid grid-cols-2 gap-4 mt-16 mr-8 mb-8">
                <div>
                    <label htmlFor="response" className="block font-medium text-white">
                        Request Method:
                    </label>
                    <select
                        id="response"
                        name="response"
                        className="mt-1 block w-full h-8 border-gray-600 shadow-sm focus:border-indigo-500 focus:ring-indigo-500   text-black text-center"
                        onChange={handleMethodNameChange}
                        value={methodName}
                    >
                        <option value="GET">GET</option>
                        <option value="POST">POST</option>
                        <option value="PUT">PUT</option>
                        <option value="PATCH">PATCH</option>
                        <option value="DELETE">DELETE</option>
                    </select>
                </div>
                <div>
                    <label htmlFor="status" className="block font-medium text-white">
                        Status Code:
                    </label>
                    <select
                        id="status"
                        name="status"
                        className="mt-1 block w-full h-8 border-gray-600 shadow-sm focus:border-indigo-500 focus:ring-indigo-500   text-black text-center"
                        onChange={handleStatusCodeChange}
                        value={statusCode}
                    >
                        <option value="200">200</option>
                        <option value="400">400</option>
                        <option value="401">401</option>
                        <option value="403">403</option>
                        <option value="404">404</option>
                        <option value="500">500</option>
                        <option value="503">503</option>
                    </select>
                </div>
                <div className="col-span-2">
                    <label htmlFor="body" className="block font-medium text-white">
                        Response Body:
                    </label>
                    <Editor
                        value={responseBody}
                        onValueChange={code => setResponseBody(code)}
                        highlight={code => highlight(code, languages.javascript, 'json')}
                        padding={12}
                        style={{
                            fontFamily: '"Fira code", "Fira Mono", monospace',
                            fontSize: 18,
                            height: "250px",
                            backgroundColor: "#282c34"

                        }}
                    />
                </div>
            </div>
            <div className="grid grid-cols-2 gap-4 mt-4 mr-8 mb-8">
                <button
                    type="submit"
                    className={editMode ? "text-center px-4 py-2 bg-blue-500 font-semibold text-white" : "text-center px-4 py-2 bg-primary font-semibold text-white"}
                    onClick={handleClick}
                >
                    {
                        editMode ? "Edit Endpoint" : "Create Endpoint"
                    }
                </button>
                <div>

                    {
                        data ?
                            <span className="text-2xl">
                                <span>{"-> "}</span>
                                < code className="text-red-500 hover:underline mb-2 md:mb-0 w-64 flex-1">
                                    <a target="_blank" href={"/custom/" + data}>
                                        fooapi.com/custom/{data}
                                    </a>
                                </code>
                            </span>
                            : false

                    }
                    {
                        tooManyError ?
                            <h1>Too many attempts to create an endpoint, try 1 hour later. You can still using the last one created</h1>
                            : false
                    }
                    {
                        expiredError ?
                            <h1>Your custom endpoint expired. Reload and create a new one</h1>
                            : false
                    }
                </div>
            </div>
        </div >
    );
}