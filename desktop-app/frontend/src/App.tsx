import './App.css';
import { GetCpuDetails, GetCpuUsage, GetRamDetails, GetGpuDetails } from "../wailsjs/go/main/App";
import { useEffect, useState } from 'react';

function App() {
    const [cpuDetails, setCpuDetails] = useState<string>("");

    useEffect(() => { // probably not how we want to do this in the future, should use something like react-query
        GetCpuDetails().then(details => setCpuDetails(details));
    }, []);

    const [cpuUsage, setCpuUsage] = useState<string>("");

    useEffect(() => {
        const interval = setInterval(() => {
            GetCpuUsage().then(usage => setCpuUsage(usage));
        }
        , 1000);
        return () => clearInterval(interval);
    }, []);

    const [ramDetails, setRamDetails] = useState<string>("");

    useEffect(() => {
        const interval = setInterval(() => {
            GetRamDetails().then(details => setRamDetails(details));
        }
        , 1000);
        return () => clearInterval(interval);
    }, []);

    const [gpuDetails, setGpuDetails] = useState<string>("");

    useEffect(() => {
        const interval = setInterval(() => {
            GetGpuDetails().then(details => setGpuDetails(details));
        }
        , 1000);
        return () => clearInterval(interval);
    }, []);

    return (
        <div id="App">
            <p>{ cpuDetails }</p>
            <p>{ cpuUsage }</p>
            <p>{ ramDetails }</p>
            <p>{ gpuDetails }</p>
        </div>
    )
}

export default App
