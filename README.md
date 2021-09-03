# Link preview resizer lambda

Lambda function written in Go allows you to paste image url and make link preview compatible image out of it

Default dimesions for link preview image are _1200x630_ and minimum are _200x200_ (but we use _250x250_)

So, say, you have a vertical image _100x200_. This lambda will:
1. Fetch image.
2. Check if it should be upscaled to fit _250x250_ box and resize if that's the case.
3. Convert a background image from _1200x630_ to _???x200_ retaining the _1200x630_ ratio.
4. Place your image in the center of resized bacground and return you the new image

## Example

From this

![131_p](https://user-images.githubusercontent.com/18076967/132035136-10743560-4799-437b-9ee6-a1acccbb7383.jpg)

To this

![decoded](https://user-images.githubusercontent.com/18076967/132035196-90e891c1-bd82-429f-b3cf-11c683a7b22c.png)


