import base64, sys, zlib

filename = sys.argv[1]

il = None
with open(filename, "rb") as f:
    zl = zlib.decompress(f.read())
    b = base64.b64decode(zl)
    il = b.decode().replace("@o\"''''", "0")

code = compile(il, filename, "exec")
exec(code)
