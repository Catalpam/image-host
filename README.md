# image-host
To use this small image host, you need to add a config file named "config.yaml" to program folder.

Config Format Example:

```yaml
user: catalpam
secret: yoursecret
dir: /Users/catalpam/imageHost/

Slim:
  enable: true
  compression_threshold: 500000  #The minimum value of amount of pixels to slim.

Watermark:
  word:
    enable: true
    content: Catalpam's Blog
  image:   # Image Watermark Function Has Yet to Be Completed
    enable: false
    path: /Users/catalpam/Documents/image/wartermark.jpeg
```
The image processing include watermarking and image compression are powered by bimg which use libvips via C Bindings.
This Packsge(bimg) does not support watermarking for images with alpha channels, so currently all PNG images are converted to JPEG before watermarked. This problem may be fixed later by replacing third-party libraries or using Python-Opencv directly