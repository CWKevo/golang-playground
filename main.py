from os import name as dist
from pathlib import Path
from hashlib import sha256



def main():
    newline: bytes = b"\r\n" if dist == "nt" else b"\n"

    to_strip: set[bytes] = {newline, b"\t", b" ", b";"}
    files = Path(".").iterdir()

    for file in files:
        if file.is_dir(): continue

        clean = file.read_bytes()

        for strip in to_strip:
            clean = clean.replace(strip, b"")


        print(file.name, sha256(clean, usedforsecurity=False).hexdigest())



if __name__ == '__main__':
    main()
