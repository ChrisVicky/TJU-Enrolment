NAME=enrollment
WINEXE=$(NAME).exe
EXE=$(WINEXE)
CONF=config.toml
GIT_ZIP=src-$(shell date +"%Y-%m-%d").zip
PACK_ZIP=pack-$(shell date +"%Y-%m-%d").zip
PDF_SRC=README.md
PDF_DST=README.pdf

all: build archive

build:
	GOOS=windows GOARCH=amd64 go build -o $(EXE) .

zip: archive pdf
	rm -f $(PACK_ZIP)
	zip -r $(PACK_ZIP) $(EXE) $(GIT_ZIP) $(PDF_DST) $(CONF)
	make clean

archive: build
	git archive -o $(GIT_ZIP) HEAD

clean:
	rm -f $(EXE) $(GIT_ZIP) $(PDF_DST)

pdf: $(PDF_SRC)
	pandoc --pdf-engine=xelatex -V CJKmainfont="Songti TC" -V mainfont="OperatorMono Nerd Font" $(PDF_SRC) -o $(PDF_DST)

.PHONY: all build zip clean pdf
