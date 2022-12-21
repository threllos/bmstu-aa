#pragma once

#include <vector>
#include <cmath>

#define UNCLASSIFIED -1
#define CORE_POINT 1
#define BORDER_POINT 2
#define NOISE -2

using namespace std;

typedef struct Point_
{
    float x, y, z;
    int clusterID;
} Point;


class DBSCAN {
public:    
    DBSCAN(unsigned int minPts, float eps, vector<Point> points)
    {
        m_minPoints = minPts;
        m_epsilon = eps;
        m_points = points;
        m_pointSize = points.size();
    }

    vector<Point> m_points;

    int points_size;
    int cluster_size;
    vector<int> clusters_size;

    void run();
    void sort();
    void preparation();

protected:
    vector<int> calculate_cluster(Point point);
    int expand_cluster(Point point, int clusterID);
    double calculate_distance(Point p1, Point p2);
    
private:    
    unsigned int m_pointSize;
    unsigned int m_minPoints;
    float m_epsilon;
};
