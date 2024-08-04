import TravelTable from "./TravelTable";

const Travel = () => {

    return (
        <>
            <div className="gradient-bg">
                <div className="container mx-auto pt-10 pb-10">
                    <div
                        key={`title`}
                        // className={`bg-[#d8d8d8] rounded-[20px] ${!value.Image && "py-[50px] md:py-[90px] lg:py-[60px] xl:py-[90px] px-[30px] md:px-[90px] lg:px-[60px] xl:px-[110px]"}`}
                        className={`bg-[#d8d8d8] rounded-[20px] py-[50px] md:py-[50px] lg:py-[50px] xl:py-[50px] px-[30px] md:px-[90px] lg:px-[60px] xl:px-[110px]`}

                        data-aos="fade-up"
                        data-aos-delay="100"
                        data-aos-duration="600"
                        data-aos-once="true"
                    >
                        <div className="grid items-center gap-[25px] grid-cols-1 sm:grid-cols-1 md:grid-cols-1 lg:grid-cols-2 xl:grid-cols-2 2xl:grid-cols-2">
                            <div>
                                <h2 className="text-black font-semibold text-[25px] md:text-[30px] lg:text-[32px] xl:text-[36px] leading-[1.2] max-w-[480px]">
                                    Statistic Travelled Country
                                </h2>
                            </div>

                            <div className="lg:text-end lg:max-w-[412px] lg:ml-auto">
                                <p className="mb-[20px]">

                                </p>
                            </div>
                        </div>
                    </div>
                </div>

                <div className="container mx-auto">
                    <div
                        data-aos="fade-up"
                        data-aos-delay="100"
                        data-aos-duration="600"
                        data-aos-once="true"
                    >
                        <TravelTable />
                    </div>
                </div>
            </div>

        </>
    )
}

export default Travel