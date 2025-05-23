# lists of all output files that can be made
ALL_PDF = $(patsubst %.tex,%.pdf,$(wildcard *.tex))
ALL_DVI = $(patsubst %.tex,%.dvi,$(wildcard *.tex))
ALL_PS = $(patsubst %.tex,%.ps,$(wildcard *.tex))


# makes everything that can be made
# default is pdf
all : $(ALL_PDF)

pdf : $(ALL_PDF)
dvi : $(ALL_DVI)
ps : $(ALL_PS)



%.pdf: %.tex
	latexmk -lualatex $*.tex

%.dvi : %.tex
	latexmk -dvi $*.tex 

%.ps : %.tex
	latexmk -ps $*.tex

clean:
	rm pollux.pdf
	latexmk -c
# cleans everything but the .tex and .pdf (or eventually .dvi or .ps)

mrproper: clean
	latexmk -C
# cleans everything but the .tex

clean_% :
	latexmk -c $*.tex

mrproper_% : clean_%
	rm -rf $*.pdf
	rm -rf $*.dvi
	rm -rf $*.ps



help :
	cat README.md


.PHONY: all clean mrproper help pdf dvi ps
