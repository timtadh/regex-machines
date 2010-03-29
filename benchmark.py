
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


#x_axis1 = list()
#y_axis1 = list()
#for i in range(1, 30):
    #x_axis1.append(i)
    #y_axis1.append(testclp(i))

#x_axis2 = list()
#y_axis2 = list()
#for i in range(1, 20):
    #x_axis2.append(i)
    #y_axis2.append(testpyclp(i))

#x_axis1 = list()
#y_axis1 = list()
#for i in range(2, 7):
    #x_axis1.append(i)
    #y_axis1.append(testclp(i))
    ##print i

#x_axis2 = list()
#y_axis2 = list()
#for i in range(2, 7):
    #x_axis2.append(i)
    #y_axis2.append(testpyclp(i))
    ##print i


#x_axis1 = list()
#y_axis1 = list()
#for i in range(2, 6):
    #x_axis1.append(i)
    #y_axis1.append(testbest(i))
    #print i

#x_axis2 = list()
#y_axis2 = list()
#for i in range(2, 6):
    #x_axis2.append(i)
    #y_axis2.append(testpybest(i))
    #print i

##p.title('CDF of Wget Ratios')
p.plot(x[:20], r, 'r-')
p.plot(x[:20], b, 'b-')
p.plot(x[:22], py, 'k-')
p.plot(x, t, 'g-')
p.legend(('recursive', 'backtracking', 'python', 'thompson'))
p.xlabel('Time to match a?^na^n  against a^n')
p.show()