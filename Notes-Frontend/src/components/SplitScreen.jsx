import React from "react";
import AddNote from "./AddNote";
import ShowNote from "./ShowNote";

const SplitScreen = () => {
  
  const notes = [
    { id: 1, title: "Title 1", content: "Content for Note 1" },
    { id: 2, title: "Title 2", content: "Content for Note 2" },
    { id: 3, title: "Title 3", content: "Content for Note 3" },
    { id: 4, title: "Title 4", content: "Content for Note 4" },
    { id: 5, title: "Title 5", content: "Content for Note 5" },
    { id: 1, title: "Title 1", content: "Content for Note 1" },
    { id: 2, title: "Title 2", content: "Content for Note 2" },
    { id: 3, title: "Title 3", content: "Content for Note 3" },
    { id: 4, title: "Title 4", content: "Content for Note 4" },
    { id: 5, title: "Title 5", content: "Content for Note 5" }, 
  ];

  return (
    <div className="bg-gray-600 h-[100vh]"> 
      <div className="bg-gray-600 h-full w-full flex"> 
        <div className="bg-gray-600 w-full sm:w-[30%] p-6 h-full overflow-hidden">
          <AddNote />
        </div>

        <div className="bg-gray-600 w-full sm:w-[70%] p-6 overflow-y-auto h-full flex flex-wrap gap-4">
          
          {notes.map((note) => (
            <div key={note.id} className="w-full sm:w-[calc(50%-1rem)]">
              <ShowNote title={note.title} content={note.content} />
            </div>
          ))}
        </div>
      </div>
    </div>
  );
};

export default SplitScreen;
