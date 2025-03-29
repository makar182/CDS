import React, { useState } from "react";

const App: React.FC = () => {
    const [firstName, setFirstName] = useState("");
    const [lastName, setLastName] = useState("");
    const [middleName, setMiddleName] = useState("");

    // @ts-ignore
    const handleSubmit = async () => {
        const response = await fetch("http://localhost:8080/submit", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({ first_name: firstName, last_name: lastName, middle_name: middleName }),
        });

        if (response.ok) {
            alert("Данные отправлены!");
        } else {
            alert("Ошибка отправки данных!");
        }
    };

    return (
        <div>
            <h1>Введите ваши данные</h1>
            <input placeholder="Имя" value={firstName} onChange={(e) => setFirstName(e.target.value)} />
            <input placeholder="Фамилия" value={lastName} onChange={(e) => setLastName(e.target.value)} />
            <input placeholder="Отчество" value={middleName} onChange={(e) => setMiddleName(e.target.value)} />
            <button onClick={handleSubmit}>Отправить</button>
        </div>
    );
};

export default App;
