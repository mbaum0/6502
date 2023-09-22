import sys
from enum import Enum
import struct
import binascii

LDA_IMM = 0xA9
STA_ABS = 0x8D
BRK = 0x0

def parse_param_to_bytes(param):
    num = param.strip('#$')
    num = int(num, 16)
    first_byte = (num & 0xFF00) >> 8
    second_byte = num & 0x00FF
    return num, first_byte, second_byte

def assemble_lda(params):
    if len(params) == 1:
        param = params[0]
        num, first, second = parse_param_to_bytes(param)
        if "#" in param:
            # immediate
            return struct.pack("BB", LDA_IMM, second)
    return struct.pack("")

def assemble_sta(params):
    if len(params) == 1:
        param = params[0]
        num, first, second = parse_param_to_bytes(param)
        if num > 255:
            # absolute address
            return struct.pack("BBB", STA_ABS, second, first)

def assemble_brk():
    return struct.pack("B", BRK)


def extract_opcodes_and_parameters(file_path):
    opcodes_and_parameters = []

    with open(file_path, 'r') as file:
        for line in file:
            parts = line.strip().split()
            op_code = parts[0]
            parameters = parts[1].split(',') if len(parts) > 1 else []
            opcodes_and_parameters.append((op_code, parameters))

    return opcodes_and_parameters

def assemble(op, params):
    blob = bytearray()
    if op == "LDA":
        blob.extend(assemble_lda(params))
    elif op == "STA":
        blob.extend(assemble_sta(params))
    elif op == "BRK":
        blob.extend(assemble_brk())
    return blob
        

if __name__ == "__main__":
    if len(sys.argv) != 2:
        print("Usage: python script.py <input_file>")
        sys.exit(1)

    input_file_path = sys.argv[1]

    try:
        opcodes_and_parameters = extract_opcodes_and_parameters(input_file_path)
        blob = bytearray()
        for op_code, parameters in opcodes_and_parameters:
            res = assemble(op_code, parameters)
            blob.extend(res)
            hex_string = ""
            if res is not None:
                hex_string = binascii.hexlify(res, bytes_per_sep=1, sep=" ").decode('utf-8').upper()
            #print(f"Op Code: {op_code}, Parameters: {parameters}\t{hex_string}")
            print(f"{op_code:<15} {len(parameters):<10} {hex_string:<10}")
            with open("out.bin", "wb") as file:
                file.write(blob)
    except FileNotFoundError:
        print(f"File not found: {input_file_path}")
        sys.exit(1)
