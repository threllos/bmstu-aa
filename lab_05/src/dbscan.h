#pragma once

#include <vector>
#include <cmath>

#define UNCLASSIFIED -1

using namespace std;

typedef struct Point_
{
    float x, y;
    int clusterID;
} Point;

typedef struct Check_
{
    int index;
    bool is_kernel;
} Check;


class DBSCAN {
public:    
    DBSCAN(unsigned int minPts, float eps, vector<Point> points)
    {
        m_minPoints = minPts;
        m_epsilon = eps;
        m_points = points;

        for (int i = 0; i < m_points.size(); i++)
            mtx.push_back(vector<bool>(m_points.size()));

        for (int i = 0; i < m_points.size(); i++)
            unclassified.push_back(i);
    }

    vector<Point> m_points;
    vector<vector<int>> clusters;
    vector<vector<bool>> mtx;
    vector<int> kernel, border, noise;

    int points_size;
    int cluster_size;
    int noise_size;
    vector<int> clusters_size;

    void run();
    void sort();
    void preparation();

protected:
    void prl_calc_mtx(int start, int end);
    void prl_calc_kernel(int start, int end);
    void prl_calc_border(int start, int end, vector<int> &tmp);
    void prl_calc_noise(int start, int end, vector<int> &tmp);
    void prl_unclustered_init(int start, int end);
    void prl_clustering_init(int start, int end, int index, vector<int> &cluster, vector<Check> &checker);
    void prl_clustering_step(int start, int end, vector<int> &cluster, vector<Check> &checker, vector<Check> &tmp_checker);

    template <typename T>
    void erase_vector(vector<T> &from, T &what);
    double distance(Point p1, Point p2);
    
private:
    vector<int> unclustered;
    vector<int> unclassified;   
    int m_minPoints;
    float m_epsilon;
};
