import './App.css';
import { GetCpuDetails } from "../wailsjs/go/main/App";
import { useEffect, useState } from 'react';

function App() {
    const [cpuDetails, setCpuDetails] = useState<string>("");

    useEffect(() => {
        GetCpuDetails().then(details => setCpuDetails(details));
    }, []);

    return (
        <div id="App">
            <p>{ cpuDetails }</p>
        </div>
    )
}

export default App
