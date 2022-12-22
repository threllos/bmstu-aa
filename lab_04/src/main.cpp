#include <stdio.h>
#include <iostream>
#include "dbscan.h"

#define MINIMUM_POINTS 4    // minimum number of cluster
#define EPSILON (0.75 * 0.75)           // distance for clustering, metre^2

#define get_time(time) __asm__ __volatile__("rdtsc" :"=A"(time))

void read_data(vector<Point> &points)
{
    float x, y;
    int points_len;
    FILE *stream = fopen ("benchmark_hepta.dat", "r");

    fscanf(stream, "%d\n", &points_len);

    for (int i = 0; i < points_len; i++)
    {
        fscanf(stream, "%f,%f\n", &x, &y);
        Point tmp_point{x, y, UNCLASSIFIED};
        points.push_back(tmp_point);
    }

    fclose(stream);
}

void print_result(vector<Point> &points)
{
    printf("Количество точек: %ld\n", points.size());
    printf("Таблица точек с индификатором кластера:\n");
    printf(
        "-------------------------------\n"
        " x      y      z     cluster_id\n"
        "-------------------------------\n"
    );

    for (auto &point: points)
    {
        printf("%5.2lf  ", point.x);
        printf("%5.2lf  ", point.y);
        printf("%d\n", point.clusterID);
    }

    printf("-------------------------------\n");
}

void save_result(vector<Point> &points)
{
    FILE *stream = fopen ("save_data.dat", "w");

    fprintf(stream, "%ld\n", points.size());

    for (auto &point: points)
    {
        fprintf(stream, "%lf,", point.x);
        fprintf(stream, "%lf,", point.y);
        fprintf(stream, "%d\n", point.clusterID);
    }

    fclose(stream);
}

int main()
{   
    size_t start, end;
    vector<Point> points;

    read_data(points);

    DBSCAN ds(MINIMUM_POINTS, EPSILON, points);

    get_time(start);
    ds.run();
    get_time(end);

    ds.preparation();
    
    printf("time: %zu\n", end - start);
    printf("points: %d\n", ds.points_size);
    printf("noise: %d\n", ds.noise_size);
    printf("clusters: %d\n", ds.cluster_size);
    for (auto &size: ds.clusters_size)
        printf("%d ", size);
    printf("\n");

    save_result(ds.m_points);
    // print_result(ds.m_points); 

    return 0;
}
