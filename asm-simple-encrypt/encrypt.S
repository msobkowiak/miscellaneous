/*
 * Author: Monika Sobkowiak <monika.sobkowiak@gmail.com>
 * 
 * Exercise made for Computer Architecture class 
 *
 * The function encrypt add 1 to the ascii code of all the characters of
 * the phrase which are not space or 0. 
 *
 * Phrase is passes to Assembly as a pointer defined in C
 *
 * To compile write make -f Makefile main
 */
.text
    .global str1 		# original phrase
    .global str2		# changed phrase

    .global encrypt		# function in Assembly invoked in C

encrypt:
    pushl %ebp		# save previous stack frame 
    movl %esp, %ebp	# new stack frame
    pushl %ebx	
    pushl %ecx

    # move strints to register
    movl str1, %ebx
    movl str2, %ecx
    # initialize a counter with 0 to count the number of changed chars
    movl $0, %edx

cycle:
    movb (%ebx), %al

    # check end of string
    cmpb $0, %al
    je end

    # check if is 0 or space in that case don't change anything 		
    cmpb $48, %al 
    je increment
    cmpb $32, %al
    je increment

    # change the char
    addb $1, %al

    # increment the counter of number of changed chars
    addl $1, %edx

# increment current position being looked
increment:
    movb %al, (%ecx)
    addl $1, %ebx
    addl $1, %ecx
    # go back to the start of cycle
    jmp cycle

# restore the stack and registers
end:	
    movl %edx, %eax
    popl %ecx
    popl %ebx		
    movl %ebp, %esp			
    popl %ebp

    # return the counter of number of changed chars
    ret





