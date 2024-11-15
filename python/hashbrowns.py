import argparse
import hashlib

ALPHABET = '+/0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz'
CHAR_IDX = {char: idx for idx, char in enumerate(ALPHABET)}

def get_next_string(s: str = None) -> str:
    if not s:
        return '+'

    chars = list(s)
    i = len(chars) - 1
    while i >= 0 and chars[i] == 'z':
        chars[i] = '+'
        i -= 1
    if i < 0:
        return '+' + ''.join(chars)

    current_index = CHAR_IDX[chars[i]]
    chars[i] = ALPHABET[current_index + 1]
    return ''.join(chars)


def parse_args():
    parser = argparse.ArgumentParser()
    parser.add_argument('-u', '--username', default='williamjackson')
    parser.add_argument('-l', '--length', type=int, default=1)
    parser.add_argument('-s', '--start', default='+')
    parser.add_argument('-b', '--best', default='f'*64)
    return parser.parse_args()


def format_sha(sha: str) -> str:
    return f'{sha[0:8]} {sha[8:16]} {sha[16:24]} {sha[24:32]} {sha[32:40]} {sha[40:48]} {sha[48:56]} {sha[56:]}'


def main():
    args = parse_args()
    best = args.best
    i = 0
    suffix = args.start
    while True:
        i += 1
        value = f'{args.username}/{suffix}'
        if i % 500000 == 0:
            print(value, end='\r')
        sha = hashlib.sha256(value.encode()).hexdigest()
        if sha < best:
            best = sha
            print(value, format_sha(sha))
        suffix = get_next_string(suffix)


if __name__ == '__main__':
    try:
        main()
    except KeyboardInterrupt:
        print('\n')
