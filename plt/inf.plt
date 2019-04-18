#!/usr/bin/gnuplot
set style line 1 lt 1 lw 2 pt 1 linecolor rgb "red"
set style line 2 lt 1 lw 2 pt 1 linecolor rgb "green"
set style line 3 lt 1 lw 2 pt 1 linecolor rgb "blue"
set grid xtics lc rgb '#555555' lw 1 lt 0
set grid ytics lc rgb '#555555' lw 1 lt 0
set encoding utf8
#############
set terminal pdfcairo  enhanced color font 'Arial,10' 
set output "./tex/img/z1.3/qinf.pdf"
set ylabel "Матожидание"
set xlabel "Число операторов"
plot "./dat/inf.dat"  using 1:2 smooth csplines title  'Матожидание очереди' w l ls 2
#############
set terminal pdfcairo  enhanced color font 'Helvetica,10'
set output "./tex/img/z1.3/ninf.pdf"
set ylabel "Коэфф загрузки операторов"
set xlabel "Число операторов"
plot "./dat/inf.dat"  using 1:3 smooth csplines title  'Коэфф загрузки' w l ls 2
#############
set terminal pdfcairo  enhanced color font 'Helvetica,10'
set output "./tex/img/z1.3/tinf.pdf"
set ylabel "Время ожидания"
set xlabel "Число операторов"
plot "./dat/inf.dat"  using 1:4 smooth csplines title  'Среднее время ожидания' w l ls 2