import { useState, useEffect } from "react";
import { Button } from "./components/ui/Button";
import { Input } from "./components/ui/input";
import { Label } from "@radix-ui/react-label";

function App() {
  const [data, setData] = useState<string[]>([]);
  const [input, setInput] = useState("");

  useEffect(() => {
    fetch("http://localhost:4001/submit", {
      method: "POST",
      mode: "cors",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(data.at(-1)),
    })
      .then((res) => console.log(res))
      .catch((err) => console.log(err));
  }, [data]);

  const handleTodo = (e: React.MouseEvent<HTMLButtonElement>) => {
    e.preventDefault();
    setData((previous) => [...previous, input]);
  };

  return (
    <>
      <form className="flex flex-col gap-y-4 max-w-7xl mx-auto pt-12">
        <Label htmlFor="todo" className="">
          Add todo
        </Label>
        <Input
          type="text"
          id="todo"
          value={input}
          onChange={(e) => setInput(e.target.value)}
        />
        <Button onClick={handleTodo} className="max-w-sm w-full mx-auto">
          Add todo
        </Button>
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
