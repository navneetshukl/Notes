import React from "react";

const ShowNote = () => {
  const title = "Sample Title";
  const content =
    "This is the content of the note. It can be a longer piece of text that describes something important.";

  return (
    <div className="max-w-[25rem] bg-white p-6 rounded-lg shadow-md border border-gray-200 min-h-[15rem]">
      <h2 className="text-2xl font-semibold text-gray-900 mb-4">{title}</h2>
      <p className="text-gray-700 text-sm">{content}</p>
    </div>
  );
};

export default ShowNote;
