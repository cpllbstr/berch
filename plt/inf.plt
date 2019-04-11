#!/usr/bin/gnuplot
set style line 1 lt 1 lw 2 pt 1 linecolor rgb "red"
set style line 2 lt 1 lw 2 pt 1 linecolor rgb "green"
set style line 3 lt 1 lw 2 pt 1 linecolor rgb "blue"
set grid xtics lc rgb '#555555' lw 1 lt 0
set grid ytics lc rgb '#555555' lw 1 lt 0
#############
set term svg enhanced background rgb 'white'
set output "./img/z1.3/qinf.svg"
set ylabel "Матожидание"
set xlabel "Число операторов"
plot "./dat/inf.dat"  using 1:2 smooth csplines title  'Матожидание очереди' w l ls 2
#############
set term svg enhanced background rgb 'white'
set output "./img/z1.3/ninf.svg"
set ylabel "Коэфф загрузки операторов"
set xlabel "Число операторов"
plot "./dat/inf.dat"  using 1:3 smooth csplines title  'Коэфф загрузки' w l ls 2
