import React, { SetStateAction, useState } from "react";
import { Button, Card, Input } from "../ui";
import type { Todo } from "@/App";

type Props = {
  Created: string;
  Todo: string;
  ID: number;
  setData: React.Dispatch<SetStateAction<Todo[]>>;
};

export default function TodoCard({ Created, Todo, ID, setData }: Props) {
  const [edit, setEdit] = useState<boolean>(false);
  const [input, setInput] = useState<string>("");

  const handleDelete = (id: number) => {
    fetch(`http://localhost:4000/delete?id=${id}`, {
      method: "DELETE",
      mode: "cors",
    })
      .then(() => {
        setData((prevState) => prevState.filter((todo) => todo.ID != id));
      })
      .catch((err) => console.log(err));
  };

  const handleEdit = (id: number) => {
    fetch(`http://localhost:4000/edit?id=${id}`, {
      method: "PATCH",
      mode: "cors",
      body: JSON.stringify(input),
    })
      .then(() => {
        setData((prevState) => {
          return prevState.map((todo) =>
            todo.ID == ID ? { ID, Todo: input, Created: todo.Created } : todo
          );
        });
      })
      .catch((err) => console.log(err));
    setEdit(false);
  };

  return (
    <Card
      key={Created}
      className="flex justify-between items-center gap-x-4 p-4"
    >
      {edit ? (
        <>
          <Input onChange={(e) => setInput(e.target.value)} value={input} />
          <div className="flex gap-x-2">
            <Button onClick={() => handleEdit(ID)}>Submit</Button>
            <Button onClick={() => setEdit(false)}>Cancel</Button>
          </div>
        </>
      ) : (
        <>
          <span>{Todo}</span>
          <div className="flex gap-x-2">
            <Button onClick={() => setEdit(true)}>Edit</Button>
            <Button onClick={() => handleDelete(ID)}>Delete</Button>
          </div>
        </>
      )}
    </Card>
  );
}
