"use client";
import { ArrowLeft2, ArrowRight2 } from "iconsax-react";
import { FC } from "react";

export interface PaginationProps {
  total: number;
  perPage: number;
  page: number;
  onPreviousClick?: () => void;
  onNextClick?: () => void;
  onPageClick?: (page: number) => void;
}

const Pagination: FC<PaginationProps> = ({
  total = 1,
  perPage = 1,
  page = 1,
  onPreviousClick = () => { },
  onNextClick = () => { },
  onPageClick = () => { },
}) => {
  const totalPage = Math.ceil(total / perPage);
  const maxPageNumbers = 5;

  const getPageNumbers = () => {
    const half = Math.floor(maxPageNumbers / 2);
    let start = Math.max(page - half, 1);
    const end = Math.min(start + maxPageNumbers - 1, totalPage);

    if (end - start < maxPageNumbers - 1) {
      start = Math.max(end - maxPageNumbers + 1, 1);
    }

    const pages = [];
    for (let i = start; i <= end; i++) {
      pages.push(i);
    }
    return pages;
  };

  const pages = getPageNumbers();

  return (
    <div className="flex flex-col items-center mt-[30px] lg:mt-[50px]">
      <div className="flex flex-col md:flex-row w-full justify-between items-center text-center space-x-[10px] space-y-[10px] md:space-y-[0]">
        <div className="inline-block">
          <button
            disabled={page === 1}
            onClick={onPreviousClick}
            className={`${page === 1
              ? "cursor-not-allowed bg-gray-100 text-gray-200"
              : "bg-[#FEE] text-red-500"
              } px-3 space-x-2 flex flex-row items-center h-[36px] leading-[32px] rounded-full text-center`}
          >
            <ArrowLeft2
              className={`${page === 1 ? "text-gray-200" : "text-red-500"}`}
            />
            <span>Sebelumnya</span>
          </button>
        </div>

        <div className="inline-block">
          <div className="flex flex-wrap justify-center space-x-2 ">
            {pages.map((pageNumber) => (
              <button
                key={pageNumber}
                onClick={() => onPageClick(pageNumber)}
                className={`${pageNumber === page
                  ? "bg-red-500 text-white"
                  : "bg-gray-200 text-gray-700"
                  } px-3 py-1 rounded-full`}
              >
                {pageNumber}
              </button>
            ))}
          </div>
        </div>

        <div className="inline-block">
          <button
            disabled={page >= totalPage}
            onClick={onNextClick}
            className={`${page >= totalPage
              ? "cursor-not-allowed bg-gray-100 text-gray-200"
              : "bg-[#FEE] text-red-500"
              } px-3 space-x-2 flex flex-row items-center h-[36px] leading-[32px] rounded-full text-center`}
          >
            <span>Selanjutnya</span>
            <ArrowRight2
              className={`${page >= totalPage ? "text-gray-200" : "text-red-500"
                }`}
            />
          </button>
        </div>
      </div>


    </div>
  );
};

export default Pagination;
