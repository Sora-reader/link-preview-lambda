# Link preview resizer lambda

Lambda function written in **Go** allows you to paste image url and make link preview compatible image out of it

Default maximum and minimum dimesions are _1500x1500_ and _200x200_. We use _1200x630_ and _250x250_

So, say, you have a vertical image _100x200_. This lambda will:

1. Fetch image.
2. Check if it should be resized to fit minimum _250x250_ and _1200x630_ maximum box and resize if that's the case.
3. Copy image, resize it to _1200x630_ ratio and blur.
4. Place your image in the center of resized background and return you the new image.

## Example

From this

![thumbnail](https://user-images.githubusercontent.com/18076967/132035136-10743560-4799-437b-9ee6-a1acccbb7383.jpg)

To this

![decoded](https://user-images.githubusercontent.com/18076967/132100144-3eb46f0a-4a56-4750-92dd-1234c6f8e928.png)
