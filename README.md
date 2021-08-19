# image-host
To use this small image host, you need to add a config file named "config.yaml" to program folder.
欢迎使用这个无比简单傻瓜，总共不到500行代码的小型图床。在启动它之前，您需要在程序目录添加一个名为config.yaml的配置文件。

Config Format Example:
config.yaml配置格式示例:

```yaml
user: catalpam # 用户名
secret: yoursecret # Password/Secret 密码
dir: /Users/catalpam/imageHost/ # 图床图片保存目录

# image slim 
# 图像压缩
Slim: 
  enable: true
  compression_threshold: 500000  #The minimum value of amount of pixels to slim.

Watermark:
  # word watermark 
  # 文字水印
  word: 
    enable: true
    content: Catalpam's Blog
    
  # Image Watermark Function Has Yet to Be Completed 
  # 图像水仍在开发中，请先关闭此功能。
  image:
    enable: false
    path: /Users/catalpam/Documents/image/wartermark.jpeg
```
The image processing include watermarking and image compression are powered by bimg which use libvips via C Bindings.
This Packsge(bimg) does not support watermarking for images with alpha channels, so currently all PNG images are converted to JPEG before watermarked. This problem may be fixed later by replacing third-party libraries or using Python-Opencv directly.
图像处理目前拥有的功能为水印和图像压缩，使用bimg驱动，bimg使用C语言libvips库，可提供高速、低内存占用的的图像处理。
然而非常可惜的是这个包(bimg)不支持带有alpha通道的图像水印，所以为了简化操作，目前所有PNG图像在水印之前都被转换为JPEG。这个问题可能会在将来通过直接与lipvips混编或直接使用Python-Opencv来解决（也有可能就这样鸽掉，又不是不能用）。

之后测试后会放出release版本，有
