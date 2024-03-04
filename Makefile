NAME=enrollment
WINEXE=$(NAME).exe
LINUXEXE=$(NAME).linux
MACEXE=$(NAME).mac
EXE=$(WINEXE) $(LINUXEXE) $(MACEXE)
GIT_ZIP=src-$(shell date +"%Y-%m-%d").zip
PACK_ZIP=pack-$(shell date +"%Y-%m-%d").zip
PDF_SRC=README.md
PDF_DST=README.pdf

CONF=config.toml
TESTFILE=test.png

all: build archive

build:
	GOOS=linux GOARCH=amd64 go build -o $(LINUXEXE) .
	GOOS=windows GOARCH=amd64 go build -o $(WINEXE) .
	GOOS=darwin GOARCH=amd64 go build -o $(MACEXE) .


zip: archive 
	rm -f $(PACK_ZIP)
	zip -r $(PACK_ZIP) $(TESTFILE) $(EXE) $(GIT_ZIP) $(PDF_SRC) $(CONF)
	make clean

archive: build
	git archive -o $(GIT_ZIP) HEAD

clean:
	rm -f $(EXE) $(GIT_ZIP) $(PDF_DST)

# pdf: $(PDF_SRC)
# 	pandoc --pdf-engine=xelatex -V CJKmainfont="Songti TC" -V mainfont="OperatorMono Nerd Font" $(PDF_SRC) -o $(PDF_DST)

.PHONY: all build zip clean pdf
