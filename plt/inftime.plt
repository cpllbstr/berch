#!/usr/bin/gnuplot
set style line 2 lt 1 lw 2 pt 1 linecolor rgb "#7a2084"
set grid xtics lc rgb '#555555' lw 1 lt 0
set grid ytics lc rgb '#555555' lw 1 lt 0
#############
set term svg enhanced background rgb 'white'
set output "./img/z1.4/qinftime.svg"
set ylabel "Матожидание"
set xlabel "Число операторов"
plot "./dat/inftime.dat"  using 1:3 smooth csplines title  'Матожидание очереди' w l ls 2
#############
set term svg enhanced background rgb 'white'
set output "./img/z1.4/ninftime.svg"
set ylabel "Коэфф загрузки операторов"
set xlabel "Число операторов"
plot "./dat/inftime.dat"  using 1:2 smooth csplines title  'Коэфф загрузки' w l ls 2
#############
set term svg enhanced background rgb 'white'
set output "./img/z1.4/tinftimetime.svg"
set ylabel "Время ожидания"
set xlabel "Число операторов"
plot "./dat/inftime.dat"  using 1:4 smooth csplines title  'Среднее время ожидания' w l ls 2