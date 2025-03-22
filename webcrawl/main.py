from lib.exportcsv import export
from lib.printer import *
from lib.data import data

def main(debug = True) -> None:
    if debug is True:
        export(data, "libraries.csv")
    else:
        print("Built for production")
    
if __name__ == "__main__":
    main()