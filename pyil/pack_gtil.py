import os

def discover_gtil(dir: str):
    files = os.listdir(dir)
    il_files = []
    for file in files:
        if(file.endswith(".gtil")):
            il_files.append(file)
    return il_files