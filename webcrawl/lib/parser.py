from lib.imports import Response, requests, Final
from lib.load_env import URL

Res: Response = requests.get(URL)
docs: Final[str] = Res.text