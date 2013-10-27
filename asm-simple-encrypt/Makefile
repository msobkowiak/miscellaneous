# variables
INCLUDES = encrypt.h
SOURCES = main.c encrypt.S
OBJFILES = main.o encrypt.o
EXEC = main

# suffix rules
.SUFFIXES : .c .o
.SUFFIXES : .S .o

# transform .c on .o 
.c.o:
	gcc -g -ansi -Wall -c $<  

# transform .S on .o
.S.o:
	gcc -c -g $<  
		
${OBJFILES}: ${SOURCES} ${INCLUDES}	

# make executable file 
${EXEC}: ${OBJFILES}
	gcc -ansi -lm -Wall -g -o ${EXEC} ${OBJFILES}

run:  ${EXEC}
	./${EXEC}

clean: 
	rm ${OBJFILES} ${EXEC} 
	

