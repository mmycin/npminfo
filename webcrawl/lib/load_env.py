from lib.imports import *
from dotenv import load_dotenv
import os

load_dotenv()
URL: Final[str] = str(os.getenv("URL"))