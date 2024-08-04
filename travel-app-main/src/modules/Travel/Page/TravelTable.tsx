import { ArrowDown, ArrowUp, Refresh2, SearchNormal1 } from "iconsax-react";
import { useEffect, useState } from "react";
import { useForm } from "react-hook-form";
import { Link } from "react-router-dom";
import { useAppDispatch, useAppSelector } from "../../../app/hooks";
import { getsTravel } from "../Services/TravelSlice";
import Pagination from "./Pagination";

type SearchForm = {
    q: string;
};

const TravelTable = () => {
    const [isArrowDown, setIsArrowDown] = useState(true);
    const [page, setPage] = useState(1)
    const [perPage, setPerPage] = useState(5)
    const [search, setSearch] = useState("")
    const dispatch = useAppDispatch();

    const { list, isError, isFetching } = useAppSelector(
        (state) => state.travel
    );

    const {
        register,
        handleSubmit,
        reset,
        formState: { errors },
    } = useForm<SearchForm>({
        // resolver: yupResolver(validationSchema),
    });

    const toggleArrow = () => {
        setIsArrowDown(!isArrowDown);

        dispatch(getsTravel({
            filter: {
                search: search,
            },
            pagination: {
                page: page,
                perPage: perPage,
            },
            sort: {
                field: "travelled_countries_code",
                order: !isArrowDown ? "ASC" : "DESC"
            }
        }));
    };

    useEffect(() => {
        async function _getsTravel() {
            await dispatch(getsTravel({
                filter: {
                    search: search,
                },
                pagination: {
                    page: page,
                    perPage: perPage
                },
                sort: {
                    field: "travelled_countries_code",
                    order: !isArrowDown ? "ASC" : "DESC"
                }
            }));
        }

        // async function _getsDataForm() {
        //     await dispatch(getsDataForm({}));
        // }

        // _getsDataForm();
        _getsTravel();

        if (isError) {
            // 
        }

    }, [isError, page, perPage]);

    const doSearch = (data: SearchForm) => {
        setSearch(data.q)
        dispatch(getsTravel({
            filter: {
                search: data.q,
            },
            pagination: {
                page: page,
                perPage: perPage,
            }
        }));
    }

    const doReset = () => {
        setSearch("")
        reset()
        dispatch(getsTravel({
            filter: {
                search: "",
            },
            pagination: {
                page: page,
                perPage: perPage,
            }
        }));
    }

    return (
        <>
            <div className="w-full">
                <div className="bg-white rounded-2xl py-4 md:py-7 px-4 md:px-8 xl:px-10">
                    <div className='w-full flex flex-col md:flex-row justify-between items-center'>
                        <div className="w-full p-2 flex flex-row space-x-10">
                            <span> Total Data : {list.GETS_TRAVEL.recordsTotal}</span>
                            <span> Page : {list.GETS_TRAVEL.currentPage}</span>
                            <span> Per Page : {list.GETS_TRAVEL.limit}</span>
                        </div>
                        <div className="w-full flex flex-col md:flex-row justify-end md:space-x-4">

                            <select
                                className="mb-2 h-14 bg-[#F2F2F8] w-full md:w-20 rounded-full p-3 placeholder:text-[#4C4C4C] focus:outline-none"
                                onChange={(v) => {
                                    setPerPage(parseInt(v.target.value))
                                }}
                            >
                                <option selected value={5} >5</option>
                                <option value={10}>10</option>
                                <option value={20}>20</option>
                                <option value={30}>30</option>
                                <option value={40}>40</option>
                                <option value={50}>50</option>
                            </select>

                            <form
                                onSubmit={handleSubmit(doSearch)}
                                className="flex flex-col w-full md:flex-row justify-end mb-5 md:space-x-2 space-y-2 md:space-y-0"
                            >
                                <label className="w-full">
                                    <input
                                        className="bg-[#F2F2F8] w-full rounded-full py-[14px] px-[30px] placeholder:text-[#4C4C4C] focus:outline-none"
                                        placeholder="Country Name"
                                        type="text"
                                        id="q"
                                        {...register("q")}
                                    />
                                </label>
                                <div className="justify-end flex flex-row h-12">
                                    <button
                                        type="submit"
                                        className=" bg-red-200 rounded-l-full px-3 text-red-500 hover:bg-red-300 flex items-center"
                                    >
                                        <SearchNormal1 className='mr-1' />
                                        <span>Cari</span>
                                    </button>
                                    <button
                                        onClick={doReset}
                                        type="button"
                                        className=" bg-gray-200 rounded-r-full px-3 text-gray-500 hover:bg-gray-300 flex items-center"
                                    >
                                        <Refresh2 className='mr-1' />
                                    </button>
                                </div>

                                <span className='text-red-500 ml-8'>
                                    {errors.q && errors.q?.message}
                                </span>
                            </form>
                        </div>

                    </div>

                    <div className="mt-7 overflow-x-auto">
                        <table className="w-full whitespace-nowrap">
                            <thead>
                                <tr className="mb-3 bg-primary-500 font-semibold text-white focus:outline-none h-16 ">
                                    <td className="rounded-l-full pl-5">
                                        <div className="flex flex-row gap-4 items-center">
                                            <div onClick={toggleArrow} className="flex justify-center items-center rounded-full cursor-pointer hover:bg-slate-600 p-2">
                                                {isArrowDown ? <ArrowUp /> : <ArrowDown />}
                                            </div>
                                            <p className="text-md leading-none">Country</p>
                                        </div>
                                    </td>
                                    <td className="rounded-r-full">
                                        <div className="flex items-center">
                                            <p className="text-base font-medium leading-none">Total</p>
                                        </div>
                                    </td>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    !isFetching ?
                                        list.GETS_TRAVEL.data && list.GETS_TRAVEL.data.map((value, index) => {
                                            return (
                                                <tr key={`-${index}`} className="mb-3 focus:outline-none h-16 border-b-2 border-gray-100">
                                                    <td className="flex flex-row items-center mt-6">
                                                        <div className="ml-7 flex items-center">
                                                            <p className="text-md font-bold leading-none text-gray-700 ml-2">
                                                                {index + 1}
                                                            </p>
                                                        </div>
                                                        <div className="ml-6 md:ml-16 flex items-center">
                                                            <p className="text-md font-bold leading-none text-gray-700 ml-2">
                                                                {value.label}
                                                            </p>
                                                        </div>
                                                    </td>
                                                    <td className="">
                                                        <div className="flex items-center">
                                                            <p className="text-md font-semibold leading-none text-gray-700">
                                                                {value.total}
                                                            </p>
                                                        </div>
                                                    </td>
                                                </tr>
                                            )
                                        })
                                        :
                                        <>
                                            {
                                                [1, 2].map(value => {
                                                    return (
                                                        <tr key={`loadding-${value}`} className="mb-3 focus:outline-none h-16 border border-gray-100 rounded">
                                                            <td colSpan={5} className="">
                                                                <div className="container mt-2 mx-auto">
                                                                    <div className="bg-[#f5f5f5] rounded-[20px]">
                                                                        <div className="rounded-md p-4 w-full mx-auto">
                                                                            <div className="animate-pulse flex space-x-4">
                                                                                <div className="flex-1 space-y-6 py-1">
                                                                                    <div className="space-y-3">
                                                                                        <div className="grid grid-cols-3 gap-4">
                                                                                            <div className="h-2 bg-slate-500 rounded col-span-2"></div>
                                                                                            <div className="h-2 bg-slate-500 rounded col-span-1"></div>
                                                                                        </div>
                                                                                        <div className="h-2 bg-slate-500 rounded"></div>
                                                                                    </div>
                                                                                </div>
                                                                            </div>
                                                                        </div>
                                                                    </div>
                                                                </div>
                                                            </td>
                                                        </tr>
                                                    )
                                                })
                                            }
                                        </>
                                }

                                {
                                    list && !list.GETS_TRAVEL.recordsTotal && <>
                                        <tr key={`loadding`} className="mb-3 focus:outline-none h-16 border border-gray-100 rounded">
                                            <td colSpan={5} className="">
                                                <div className="container lg:max-w-[1710px] mx-auto text-center">
                                                    <div className="bg-[#ffffff] rounded-[20px] container mx-auto py-10 px-[12px] 2xl:px-0">

                                                        <h3 className="text-black text-[20px] md:text-[22px] lg:text-[30px] font-bold mt-[40px] mb-[15px] leading-[1.3]">
                                                            Oops! data tidak ditemukan
                                                        </h3>

                                                        <Link
                                                            to="/"
                                                            className="bg-black text-white text-[14px] font-medium inline-block uppercase rounded-full py-[15px] px-[38px] transition duration-500 ease-in-out hover:bg-[#EF4335]"
                                                        >
                                                            Refresh <Refresh2
                                                                className="inline-block relative -top-[2px] ml-3"
                                                                size={20}
                                                            />
                                                        </Link>
                                                    </div>
                                                </div>
                                            </td>
                                        </tr>
                                    </>
                                }
                            </tbody>
                        </table>
                    </div>

                    <Pagination
                        total={list.GETS_TRAVEL.recordsTotal}
                        page={page}
                        perPage={perPage}
                        onNextClick={() => {
                            setPage(page + 1)
                        }}
                        onPreviousClick={() => {
                            setPage(page - 1)
                        }}
                        onPageClick={(clickedPageNumber) => {
                            setPage(clickedPageNumber)
                        }}
                    />

                </div>
            </div >
        </>
    )
}

export default TravelTable