import { useState, useEffect } from "react";
import { Button, Card, Input, Label } from "./components/ui";

type Todo = {
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

  const handleTodo = (e: React.MouseEvent<HTMLButtonElement>) => {
    e.preventDefault();
    if (input == "") return;
    postData();
  };

  const handleDelete = (id: number) => {
    fetch(`http://localhost:4000/delete?id=${id}`, {
      method: "DELETE",
      mode: "cors",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(input),
    })
      .then(() => {
        setData((prevState) => prevState.filter((todo) => todo.ID != id));
      })
      .catch((err) => console.log(err));
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
        <Button onClick={handleTodo} className="max-w-sm w-full mx-auto">
          Add todo
        </Button>
      </form>
      <div className="mt-12 max-w-7xl mx-auto px-12">
        {data.map((todo) => {
          return (
            <Card
              key={todo.Created}
              className="flex justify-between items-center gap-x-4 p-4"
            >
              <span>{todo.Todo}</span>
              <Button onClick={() => handleDelete(todo.ID)}>Delete</Button>
            </Card>
          );
        })}
      </div>
    </>
  );
}

export default App;
