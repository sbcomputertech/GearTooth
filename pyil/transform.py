"""
- pyil.transform: a script ran by the geartooth compiler to post-process generated python code
"""

import python_minifier, sys, ratio, zlib, base64

il_file = open(sys.argv[1], "r")
minified = python_minifier.minify(il_file.read())
il_file.close()

exp_code = ratio.run(minified)

b64 = base64.b64encode(exp_code.encode(), b"%&")
zl = zlib.compress(b64, 9)

out_file = open(sys.argv[2], "wb")
out_file.write(zl)
out_file.close()
