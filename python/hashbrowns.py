import argparse
import hashlib
import itertools


def parse_args():
    parser = argparse.ArgumentParser()
    parser.add_argument('-u', '--username', default='williamjackson')
    parser.add_argument('-l', '--length', type=int, default=1)
    parser.add_argument('-s', '--start', default='')
    parser.add_argument('-b', '--best', default='f'*64)
    return parser.parse_args()


def format_sha(sha: str) -> str:
    return sha[0:8] + ' ' + sha[8:16] + ' ' + sha[16:24] + ' ' + sha[24:32] + ' ' + sha[32:40] + ' ' + sha[40:48] + ' ' + sha[48:56] + ' ' + sha[56:]


def main():
    args = parse_args()
    base64_chars = '+/0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz'
    best = args.best
    for i, x in enumerate(itertools.product(base64_chars, repeat=args.length)):
        suffix = ''.join(x)
        value = f'{args.username}/{suffix}'
        if i % 500000 == 0:
            print(value, end='\r')
        if suffix > args.start:
            sha = hashlib.sha256(value.encode()).hexdigest()
            if sha < best:
                best = sha
                print(value, format_sha(sha))


if __name__ == '__main__':
    try:
        main()
    except KeyboardInterrupt:
        print('\n')
