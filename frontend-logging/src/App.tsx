import React, { useEffect, useState } from "react";

const App: React.FC = () => {
    const [logs, setLogs] = useState<{ id: number; message: string }[]>([]);

    useEffect(() => {
        // @ts-ignore
        const fetchLogs = async () => {
            const response = await fetch("http://localhost:8082/logs");
            if (response.ok) {
                const data = await response.json();
                setLogs(data);
            } else {
                alert("Ошибка загрузки логов!");
            }
        };

        fetchLogs();
    }, []);

    return (
        <div>
            <h1>Логи</h1>
            <ul>
                {logs.map((log) => (
                    <li key={log.id}>{log.message}</li>
                ))}
            </ul>
        </div>
    );
};

export default App;
