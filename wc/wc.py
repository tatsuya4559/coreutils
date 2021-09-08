#!/usr/bin/env python
import sys
import os


def eprint(*args, **kwargs):
    print(file=sys.stderr, *args, **kwargs)


def main():
    nlines = 0
    nwords = 0
    nchars = 0
    if len(sys.argv) < 2:
        for _ in sys.stdin:
            nlines += 1
        print(nlines)
        exit(0)

    for filepath in sys.argv[1:]:
        nl = 0
        nw = 0
        nc = 0
        if not os.path.exists(filepath):
            eprint(f'wc: {filepath}: open: No such file or directory')
            continue
        with open(filepath, 'r') as fp:
            for line in fp:
                nl += 1
                nw += len(line.split())
                nc += len(line)
        print(f'{nl} {nw} {nc} {filepath}')
        nlines += nl
        nwords += nw
        nchars += nc
    if len(sys.argv) > 2:
        print(f'{nlines} {nwords} {nchars} total')


if __name__ == '__main__':
    main()
