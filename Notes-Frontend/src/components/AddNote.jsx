import React from "react";

const AddNote = () => {
  return (
    
    <form className="space-y-6 flex flex-col items-center">
      <div>
        <input
          type="text"
          id="title"
          className="bg-white border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-[20rem] p-2.5"
          placeholder="Title"
          required
        />
      </div>

      <div>
        <textarea
          id="content"
          rows="10"
          className="block p-4 w-[20rem] text-sm text-gray-900 bg-white rounded-lg border border-gray-300 focus:ring-blue-500 focus:border-blue-500 resize-none"
          placeholder="Content"
        ></textarea>
      </div>

      <div>
        <button className="px-8 py-4 text-medium font-semibold text-white bg-blue-500 rounded-lg hover:bg-blue-600 focus:outline-none focus:ring-2 focus:ring-blue-400 w-[30rem] sm:w-auto">
          Add Note
        </button>
      </div>
    </form>
    
  );
};

export default AddNote;
