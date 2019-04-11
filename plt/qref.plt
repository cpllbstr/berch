#!/usr/bin/gnuplot
set style line 1 lt 1 lw 2 pt 1 linecolor rgb "red"
set style line 2 lt 1 lw 2 pt 1 linecolor rgb "green"
set style line 3 lt 1 lw 2 pt 1 linecolor rgb "blue"
set grid xtics lc rgb '#555555' lw 1 lt 0
set grid ytics lc rgb '#555555' lw 1 lt 0
#############
set term svg enhanced background rgb 'white'
set output "./img/z1.2/qref.svg"
set ylabel "Доля отказов"
set xlabel "Длина очереди"
plot "./dat/qref.dat"  using 1:2 smooth csplines title  'refuses' w l ls 1
##############
set term svg enhanced background rgb 'white'
set output "./img/z1.2/kcpref.svg"
set ylabel "Загрузка операторов"
set xlabel "Число операторов"
plot "./dat/qref.dat" using 3:4 smooth csplines title 'Коэффициент загрузки' w l ls 3
#############
set term svg enhanced background rgb 'white'
set output "./img/z1.2/matozh.svg"
set autoscale
set ylabel "Матожидание длины очереди"
set xlabel "Длина очереди"
plot "./dat/qref.dat" using 1:5 smooth csplines title 'Матожидание длины очереди' w l ls 3
#############
set term svg enhanced background rgb 'white'
set output "./img/z1.2/queucoeff.svg"
set autoscale
set ylabel "Коэффициент занятости"
set xlabel "Длина очереди"
plot "./dat/qref.dat" using 1:6 smooth csplines title 'Коэффициент занятости очереди' w l ls 3
#############
set term svg enhanced background rgb 'white'
set output "./img/z1.2/qwtime.svg"
set autoscale
set ylabel "Среднее время ожидания в очереди"
set xlabel "Длина очереди"
plot "./dat/qref.dat" using 1:7 smooth csplines title 'Время' w l ls 3