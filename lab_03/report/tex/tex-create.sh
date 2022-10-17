pdflatex.exe -synctex=1 -interaction=nonstopmode "report".tex
bibtex report.aux
pdflatex.exe -synctex=1 -interaction=nonstopmode "report".tex