#include <algorithm>
#include "dbscan.h"

void DBSCAN::run()
{
    int clusterID = 1;
    vector<Point>::iterator iter;

    for (auto &point: m_points)
    {
        if (point.clusterID == UNCLASSIFIED && expand_cluster(point, clusterID) == 0)
            clusterID++;
    }
}

void DBSCAN::sort()
{
    std::sort(
        m_points.begin(),
        m_points.end(),
        [](Point p1, Point p2) {return (p1.clusterID < p2.clusterID);}
    );
}

void DBSCAN::preparation()
{
    int size = 1;
    int clusterID = m_points[0].clusterID;

    sort();

    for (int i = 1; i < m_points.size(); i++)
    {
        if (m_points[i].clusterID != clusterID)
        {
            clusters_size.push_back(size);
            clusterID = m_points[i].clusterID;
            size = 0;
        }

        size++;
    }

    clusters_size.push_back(size);
    points_size = m_points.size();
    cluster_size = clusters_size.size();
}

int DBSCAN::expand_cluster(Point point, int clusterID)
{    
    vector<int> clusterSeeds = calculate_cluster(point);

    if (clusterSeeds.size() < m_minPoints)
    {
        point.clusterID = NOISE;
        return -1;
    }

    int index = 0, indexCorePoint = 0;

    for (auto &seed: clusterSeeds)
    {
        m_points[seed].clusterID = clusterID;

        if (
            m_points[seed].x == point.x &&
            m_points[seed].y == point.y &&
            m_points[seed].z == point.z
        )
            indexCorePoint = index;
        
        index++;
    }
    
    clusterSeeds.erase(clusterSeeds.begin() + indexCorePoint);
    vector<int>::size_type n = clusterSeeds.size();
    
    for(int i = 0; i < n; i++)
    {
        vector<int> clusterNeighors = calculate_cluster(m_points[clusterSeeds[i]]);

        if (clusterNeighors.size() >= m_minPoints)
        {
            for (auto &neighor: clusterNeighors)
            {
                if (
                    m_points[neighor].clusterID == UNCLASSIFIED ||
                    m_points[neighor].clusterID == NOISE
                )
                {
                    clusterSeeds.push_back(neighor);
                    n = clusterSeeds.size();
                }

                m_points[neighor].clusterID = clusterID;
            }
        }
    }

    return 0;
}

vector<int> DBSCAN::calculate_cluster(Point point)
{
    int index = 0;
    vector<Point>::iterator iter;
    vector<int> clusterIndex;

    for (auto &target_point: m_points)
    {
        if (calculate_distance(point, target_point) <= m_epsilon)
            clusterIndex.push_back(index);
        
        index++;
    }

    return clusterIndex;
}

double DBSCAN::calculate_distance(Point p1, Point p2)
{
    return pow(p1.x - p2.x, 2) + pow(p1.y - p2.y, 2) + pow(p1.z - p2.z, 2);
}


