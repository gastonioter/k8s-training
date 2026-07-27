import React from "react";

export interface Song {
  id: number;
  author: string;
  title: string;
}

interface SongCardProps {
  song: Song;
}

export const SongCard: React.FC<SongCardProps> = ({ song }) => {
  return (
    <div className="flex items-center justify-between p-4 bg-white rounded-xl shadow-sm border border-slate-100 hover:shadow-md transition-shadow duration-200">
      <div className="flex items-center gap-4">
        {/* Number badge */}
        <div className="flex items-center justify-center w-10 h-10 rounded-lg bg-indigo-50 text-indigo-600 font-bold text-sm">
          #{song.id}
        </div>

        {/* Song info */}
        <div>
          <h3 className="font-semibold text-slate-800 text-base">
            {song.title}
          </h3>
          <p className="text-sm text-slate-500">{song.author}</p>
        </div>
      </div>

      {/* Audio indicator icon */}
      <div className="text-slate-400 hover:text-indigo-600 transition-colors cursor-pointer">
        <svg
          xmlns="http://www.w3.org/2000/svg"
          fill="none"
          viewBox="0 0 24 24"
          strokeWidth={1.5}
          stroke="currentColor"
          className="w-6 h-6"
        >
          <path
            strokeLinecap="round"
            strokeLinejoin="round"
            d="M5.25 5.653c0-.856.917-1.398 1.667-.986l11.54 6.347a1.125 1.125 0 0 1 0 1.972l-11.54 6.347a1.125 1.125 0 0 1-1.667-.986V5.653Z"
          />
        </svg>
      </div>
    </div>
  );
};
