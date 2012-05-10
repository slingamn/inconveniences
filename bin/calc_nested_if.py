#!/usr/bin/python

"""
Computes an Libreoffice Calc formula that turns raw scores into letter grades,
via a monstrous nested "IF" expression.

Give it one argument, the name of the input cell (e.g., "Z8"), and stdin like this:

96 A+
88 A
81 A-
74 B+
67 B
60 B-
53 C+
47 C
40 C-
35 D
0  F

While I'm at it, here's one more tip about automating the grading process in Calc:
to drop the two lowest quiz scores, fill in 0 for all untaken quizzes, then do:
SUM(C8:O8) - SMALL(C8:O8, 1) - SMALL(C8:O8, 2)
"""

import sys

def make_formula(lines, cell):
    cutoffs = []
    for line in lines:
        minimum, grade = line.rstrip().split()
        minimum = int(minimum)
        cutoffs.append((minimum, grade))

    # lowest bucket first
    cutoffs.reverse()
    # verify original order
    assert cutoffs == sorted(cutoffs), 'Cutoffs in wrong order, double-check.'
    assert cutoffs[0][0] == 0, 'Bottom cutoff must be 0.'
    assert cutoffs[-1][0] < 100, 'Top cutoff must be <100.'

    # start from the bottom bucket
    cutoffs.sort()
    # stupid edge case: if the grade is not >= 0
    formula = '"XXX"'
    for minimum, grade in cutoffs:
        formula = 'IF(%s >= %d; "%s"; %s)' % (cell, minimum, grade, formula)

    return formula

if __name__ == '__main__':
    if len(sys.argv) < 2:
        print >>sys.stderr, "Requires one argument, the name of the input cell."
        sys.exit(1)
    cell = sys.argv[1]
    print make_formula(sys.stdin, cell)
