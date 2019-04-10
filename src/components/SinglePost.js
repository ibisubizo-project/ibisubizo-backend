import React from 'react';


const SinglePost = () => {
    return (
        <div className="flex border-b border-solid border-grey-light">
            <div className="w-10/10 p-3 pl-0">
                <div className="flex justify-between">
                    <div>
                        <span className="font-bold"><a href="#" className="text-black">Ofonime Francis</a></span>
                        <span className="text-grey-dark">7 Apr 2019</span>
                    </div>
                    
                    <div>
                        <a href="#" className="text-grey-dark hover:text-teal"><i className="fa fa-chevron-down"></i></a>
                    </div>
                </div>

                <div>
                    <div className="mb-4">
                        <p className="mb-6">💥 Check out this Slack clone built with <a href="#" className="text-teal">@tailwindcss</a> using no custom CSS and just the default configuration:</p>
                        <p className="mb-4"><a href="#" className="text-teal">https://codepen.io/adamwathan/pen/JOQWVa...</a></p>
                        <p className="mb-6">(based on some work <a href="#" className="text-teal">@Killgt</a> started for <a href="#" className="text-teal">tailwindcomponents.com</a> !)</p>
                        <p><a href="#"><img src="https://s3-us-west-2.amazonaws.com/s.cdpn.io/195612/tt_tweet2.jpg" alt="tweet image" className="border border-solid border-grey-light rounded-sm" /></a></p>
                    </div>
                    <div className="pb-2">
                        <span className="mr-8"><a href="#" className="text-grey-dark no-underline hover:no-underline hover:text-blue-light">
                            <i className="fa fa-comment fa-lg mr-2"></i> 19</a>
                        </span>
                        <span className="mr-8"><a href="#" className="text-grey-dark no-underline hover:no-underline hover:text-green">
                            <i className="fa fa-heart fa-lg mr-2"></i> 56</a>
                        </span>
                        <span className="mr-8"><a href="#" className="text-grey-dark no-underline hover:no-underline hover:text-red">
                            <i className="fa fa-share fa-lg mr-2"></i> 247</a>
                        </span> 
                    </div>
                </div>
            </div>
        </div>

    )
}

export default SinglePost