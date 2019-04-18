#!/usr/bin/gnuplot
set style line 2 lt 1 lw 2 pt 1 linecolor rgb "#e28cb3"
set grid xtics lc rgb '#555555' lw 1 lt 0
set grid ytics lc rgb '#555555' lw 1 lt 0
#############
set terminal pdfcairo  enhanced color font 'Helvetica,10'
set output "./tex/img/z2/nozh.pdf"
set ylabel "Матожидание"
set xlabel "Число наладчиков"
plot "./dat/zad2.dat"  using 1:2 smooth csplines title  'Ожидающие наладки станки' w l ls 2
#############
set terminal pdfcairo  enhanced color font 'Helvetica,10'
set output "./tex/img/z2/nprost.pdf"
set ylabel "Матожидание"
set xlabel "Число наладчиков"
plot "./dat/zad2.dat"  using 1:3 smooth csplines title  'Простаивающие станки' w l ls 2
#############
set terminal pdfcairo  enhanced color font 'Helvetica,10'
set output "./tex/img/z2/mz.pdf"
set ylabel "Матожидание"
set xlabel "Число наладчиков"
plot "./dat/zad2.dat"  using 1:4 smooth csplines title  'Занятые наладчики' w l ls 2
#############
set terminal pdfcairo  enhanced color font 'Helvetica,10'
set output "./tex/img/z2/mzdn.pdf"
set ylabel "Коэффициент занятости"
set xlabel "Число наладчиков"
plot "./dat/zad2.dat"  using 1:5 smooth csplines title  'Занятые наладчики' w l ls 2