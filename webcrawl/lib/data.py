from lib.imports import *
from lib.parser import docs

soup = BeautifulSoup(docs, "html.parser")

data: Data = []

names: List[str] = []
descriptions: List[str] = []
links: List[str] = []
reg: List = soup.find_all(class_="box box-pkg rounded p-3")

for item in reg:
    names.append(str(item.h3.text))
    descriptions.append(str(item.p.text))
    links.append(str(item.ul.li.a.get("href")))
    
for i in range(len(names)):
    data.append({
        names[i]: [descriptions[i], links[i]]
    })
    
data = data