#include <iostream>
#include <math.h>  
using namespace std;
FILE *scr,*data; 

int N = 6;
int N_max = 11;
int Ts = 229;
int Tc = 43;

double lamb = (double)1/Tc;
double mu = (double)1/Ts;

int channels[11];
double queue_length[11];
double load_coef[11];
double wait_time[11];


long int fact(int x)
{
    if(x < 0) // если пользователь ввел отрицательное число
        return 0; // возвращаем ноль
    if (x == 0) // если пользователь ввел ноль,
        return 1; // возвращаем факториал от нуля
    else // Во всех остальных случаях
        return x * fact(x - 1); // делаем рекурсию.
}

int main()
{
	double* P;
	int n = N;
	for(int i = 0; i < (N_max - N + 1); i++)
	{
		channels[i] = n;
		queue_length[i] = 0;
		load_coef[i] = 0;
		wait_time[i] = 0;
		n = n + 1;
	}

	n = N;
	int i;
	double sum;

	data = fopen("data.txt","w");

	while(n < (N_max + 1))
	{
		P = new double[n+1];
		i = 0;
		sum = 0;

		double Pn = (double)pow(lamb,n)/(fact(n)*pow(mu,n));
    	double inf_part = (double)lamb/(n*mu - lamb);

    	//находим P0
    	while(i < n+1)
    	{
    		sum  = (double)sum + (double)pow(lamb,i)/(fact(i)*pow(mu,i));
    		i += 1;
    	}
    	P[0] = (double)1/(sum + Pn*inf_part);

    	//находим предельные вероятности состояний системы
    	i = 1;
    	while(i < n+1)
    	{
        	P[i] = (double)pow(lamb,i)/(fact(i)*pow(mu,i))*P[0];
        	i += 1;
        }
        
        double a = (double)lamb/(n*mu);

        queue_length[n-N] = (double)(P[n]*a)/pow(1-a,2);
        wait_time[n-N] = (double)queue_length[n-N]/lamb;
        load_coef[n-N] = (double)lamb/(mu*n);

        fprintf(data,"%d\t%lf\t%lf\t%lf\n",channels[n-N],queue_length[n-N],wait_time[n-N],load_coef[n-N]);

        cout << queue_length[n-N] << endl;


		n += 1;
		delete[] P;
	}

	scr = fopen("script.dat","w");
	fprintf(scr, "set grid ytics mytics\n");
	fprintf(scr, "set grid\n");
	fprintf(scr,"plot 'data.txt' using 1:2 smooth bezier with line title 'Мат. ожидание длины очереди'\n");
	fprintf(scr, "pause -1\n");
	fprintf(scr,"plot 'data.txt' using 1:3 smooth bezier with line title 'Мат. ожидание времени пребывания клиентов в очереди'\n");
	fprintf(scr, "pause -1\n");
	fprintf(scr,"plot 'data.txt' using 1:4 smooth bezier with line title 'Коэффициент загрузки операторов'\n");
	fprintf(scr, "pause -1\n");
}