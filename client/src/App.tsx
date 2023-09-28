import { useState } from "react";

function App() {
  const [data, setData] = useState<string[]>([]);
  const [input, setInput] = useState("");

  const handlePost = (e: React.MouseEvent<HTMLButtonElement>) => {
    e.preventDefault();
    fetch("http://localhost:4000/submit", {
      method: "POST",
      mode: "cors",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(data),
    })
      .then((res) => console.log(res))
      .catch((err) => console.log(err));
  };

  const handleTodo = (e: React.MouseEvent<HTMLButtonElement>) => {
    e.preventDefault();
    setData((previous) => [...previous, input]);
  };

  return (
    <>
      <form>
        <label htmlFor="todo">Add todo</label>
        <input
          type="text"
          id="todo"
          value={input}
          onChange={(e) => setInput(e.target.value)}
        />
        <button onClick={handleTodo}>Add todo</button>
        <button onClick={handlePost}>Submit</button>
      </form>
      <div>
        {data.map((todo, i) => (
          <p key={i}>{todo}</p>
        ))}
      </div>
    </>
  );
}

export default App;
