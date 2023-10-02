import { useState, useEffect } from "react";
import { Button, Input, Label } from "./components/ui";
import { TodoCard } from "./components/reusables";

export type Todo = {
  ID: number;
  Todo: string;
  Created: string;
};

function App() {
  const [data, setData] = useState<Todo[]>([]);
  const [input, setInput] = useState("");

  useEffect(() => {
    fetch("http://localhost:4000/get-all")
      .then((res) => res.json())
      .then((res) => setData(res))
      .catch((err) => console.log(err));
  }, []);

  const postData = () => {
    fetch("http://localhost:4000/submit", {
      method: "POST",
      mode: "cors",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(input),
    })
      .then((res) => res.json())
      .then((res) => {
        console.log(res);
        return setData((previous) => {
          return [
            ...previous,
            { ID: res, Todo: input, Created: String(Date.now()) },
          ];
        });
      })
      .catch((err) => console.log(err));
  };

  const handleSubmit = (e: React.MouseEvent<HTMLButtonElement>) => {
    e.preventDefault();
    if (input == "") return;
    postData();
  };

  return (
    <>
      <form className="flex flex-col gap-y-4 max-w-7xl mx-auto pt-12 px-12">
        <Label htmlFor="todo" className="">
          Add todo
        </Label>
        <Input
          type="text"
          id="todo"
          value={input}
          onChange={(e) => setInput(e.target.value)}
        />
        <Button onClick={handleSubmit} className="max-w-sm w-full mx-auto">
          Add todo
        </Button>
      </form>
      <div className="mt-12 max-w-7xl mx-auto px-12 flex flex-col gap-y-2">
        {data.map((todo) => {
          return (
            <TodoCard
              setData={setData}
              ID={todo.ID}
              Created={todo.Created}
              Todo={todo.Todo}
              key={todo.Created}
            />
          );
        })}
      </div>
    </>
  );
}

export default App;
