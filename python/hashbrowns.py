import argparse
import hashlib
import itertools


def parse_args():
    parser = argparse.ArgumentParser()
    parser.add_argument('-u', '--username', default='williamjackson')
    parser.add_argument('-l', '--length', type=int, default=1)
    parser.add_argument('-s', '--start', default='')
    parser.add_argument('-b', '--best', default='z'*64)
    return parser.parse_args()


def main():
    args = parse_args()
    base64_chars = '+/0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz'
    best = args.best
    for i, x in enumerate(itertools.permutations(base64_chars, args.length)):
        suffix = ''.join(x)
        value = f'{args.username}/{suffix}'
        if i % 500000 == 0:
            print(value, end='\r')
        if suffix > args.start:
            sha = hashlib.sha256(value.encode()).hexdigest()
            if sha < best:
                best = sha
                print(value, sha)


if __name__ == '__main__':
    try:
        main()
    except KeyboardInterrupt:
        print('\n')
