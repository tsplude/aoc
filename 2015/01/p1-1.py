import sys

input_file = sys.argv[1]
instructions = open(input_file, 'r').read()
print instructions.count('(') - instructions.count(')')
