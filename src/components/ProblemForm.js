import React, {Component} from 'react';


const ProductForm = () => {
    return (
        <div className="container mx-auto text-black">
            <div className="post_area bg-grey-lighter p-10">
                <div>Create Post</div>

                <div className="post__box bg-white p-10">
                    <form class="w-full max-w-xs">
                        <div class="md:flex md:items-center mb-6">
                            <div class="md:w-1/3">
                                <label class="block text-grey font-bold md:text-right mb-1 md:mb-0 pr-4" for="inline-full-name">
                                    Title
                                </label>
                            </div>
                            <div class="md:w-2/3">
                                <input class="bg-grey-lighter appearance-none border-2 border-grey-lighter rounded w-full py-2 px-4 text-grey-darker leading-tight focus:outline-none focus:bg-white focus:border-purple" id="inline-full-name" type="text" value="Jane Doe" />
                            </div>
                        </div>
                        
                        <div class="md:flex md:items-center mb-6">
                            <div class="md:w-1/3">
                                <label class="block text-grey font-bold md:text-right mb-1 md:mb-0 pr-4" for="inline-username">
                                    Detail
                                </label>
                            </div>
                            <div class="md:w-2/3">
                                <textarea class="bg-grey-lighter appearance-none border-2 border-grey-lighter rounded w-full py-2 px-4 text-grey-darker leading-tight focus:outline-none focus:bg-white focus:border-purple" id="inline-username" type="password" placeholder="******************"></textarea>
                            </div>
                        </div>
                    </form>
                </div>
            </div>
        </div>
    )
}

export default ProductForm;