import { useState } from "react";
import { Button } from "./components/ui/Button";
import { Input } from "./components/ui/input";
import { Label } from "@radix-ui/react-label";

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
      <form className="flex flex-col gap-y-4 max-w-7xl mx-auto pt-12">
        <Label htmlFor="todo" className="mx-auto">
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
        <Button onClick={handlePost} className="max-w-sm w-full mx-auto">
          Submit
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
