NAME=enrolment
WINEXE=$(NAME).exe
LINUXEXE=$(NAME).linux
MACEXEINTEL=$(NAME).intel.mac
MACEXEM=$(NAME).m.mac
EXE=$(WINEXE) $(LINUXEXE) $(MACEXEM) $(MACEXEINTEL)
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
	GOOS=darwin GOARCH=amd64 go build -o $(MACEXEM) .
	GOOS=darwin GOARCH=arm64 go build -o $(MACEXEINTEL) .

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
