import React, { SetStateAction } from "react";
import { Button, Card } from "../ui";
import type { Todo } from "@/App";

type Props = {
  Created: string;
  Todo: string;
  ID: number;
  setData: React.Dispatch<SetStateAction<Todo[]>>;
};

export default function TodoCard({ Created, Todo, ID, setData }: Props) {
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
    })
      .then(() => {
        setData((prevState) => prevState.filter((todo) => todo.ID != id));
      })
      .catch((err) => console.log(err));
  };
  return (
    <Card
      key={Created}
      className="flex justify-between items-center gap-x-4 p-4"
    >
      <span>{Todo}</span>
      <div className="flex gap-x-2">
        <Button onClick={() => handleEdit(ID)}>Edit</Button>
        <Button onClick={() => handleDelete(ID)}>Delete</Button>
      </div>
    </Card>
  );
}
