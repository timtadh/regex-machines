
import time
import re
import pylab as p
import sys

def test(n):
    n = int(n)
    r = re.compile('a?'*n+'a'*n)
    t = 'a'*n
    s = time.time()
    r.match(t)
    e = time.time()
    return e-s

x = list()
r = list()
b = list()
t = list()
py = list()
for line in sys.stdin.readlines():
    f = [float(f) for f in line[:-1].split(',')]
    x.append(f[0])
    if f[0] <= 20:
        r.append(f[1])
        b.append(f[2])
    if f[0] <= 22:
        py.append(test(f[0]))
    t.append(f[3])

p.plot(x[:20], r, 'r-')
p.plot(x[:20], b, 'b-')
p.plot(x[:22], py, 'k-')
p.plot(x, t, 'g-')
p.legend(('recursive', 'backtracking', 'python', 'thompson'))
p.xlabel('Time to match a?^na^n  against a^n')
p.show()
