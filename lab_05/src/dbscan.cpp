#include <algorithm>
#include <iostream>
#include "dbscan.h"

void DBSCAN::run()
{

    // РЕЗУЛЬТАТЫ НОРМАЛЬНОГО АЛГОСА
    // time: 20791491
    // points: 212
    // clusters: 5
    // 92 30 30 30 30 
    
    // считаем достижимости
    prl_calc_mtx(0, m_points.size());

    // поиск ядерных точек
    prl_calc_kernel(0, m_points.size());

    // временный массив классифицированных точек
    vector<int> tmp;

    // поиск граничных точек
    prl_calc_border(0, unclassified.size(), tmp);

    // удаление граничных точек из неклассифицированных
    for (int &index: tmp)
        erase_vector(unclassified, index);
    tmp.clear();

    // помечаем оставшие точки шумом
    prl_calc_noise(0, unclassified.size(), tmp);

    // удаление шумовых точек из неклассифицированных
    for (int &index: tmp)
        erase_vector(unclassified, index);

    // инициализация некластиризованных точек (ядерные)
    prl_unclustered_init(0, kernel.size());

    // кластеризация
    while (unclustered.size() != 0)
    {
        vector<int> cluster;
        vector<Check> checker;
        vector<Check> tmp_checker;
        int index = unclustered[0];
        cluster.push_back(index);

        // добавить все ядровые точки достижимые из текущей
        // добавить все граниченые точки достижимые из ядровых кластера
        // пройтись по добавленным и добавить точки, которые достижимы из них и не шумовые
        // проходиться по таким точкам, пока они не закончатся
        prl_clustering_init(0, m_points.size(), index, cluster, checker);

        while (checker.size() != 0)
        {
            prl_clustering_step(0, checker.size(), cluster, checker, tmp_checker);

            checker.swap(tmp_checker);
            tmp_checker.clear();
        }

        for (int &index: cluster)
            erase_vector(unclustered, index);

        clusters.push_back(cluster);
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
    for (int i = 0; i < clusters.size(); i++)
        for (int j = 0; j < clusters[i].size(); j++)
            m_points[clusters[i][j]].clusterID = i;

    for (auto &cluster: clusters)
        clusters_size.push_back(cluster.size());

    sort();

    points_size = m_points.size();
    cluster_size = clusters_size.size();
    noise_size = noise.size();
}

template <typename T>
void DBSCAN::erase_vector(vector<T> &from, T &what)
{
    from.erase(
        remove(
            from.begin(), from.end(), what
        ), from.end()
    );
}

double DBSCAN::distance(Point p1, Point p2)
{
    return pow(p1.x - p2.x, 2) + pow(p1.y - p2.y, 2);
}

void DBSCAN::prl_calc_mtx(int start, int end)
{
    for (int i = start; i < end; i++)
        for (int j = 0; j < m_points.size(); j++)
            if (i != j && distance(m_points[i], m_points[j]) <= m_epsilon)
                mtx[i][j] = true;
}

void DBSCAN::prl_calc_kernel(int start, int end)
{
    for (int i = start; i < end; i++)
    {
        int neighbors_count = 0;

        for (int j = 0; j < m_points.size(); j++)
            neighbors_count += int(mtx[i][j]);

        if (neighbors_count + 1 >= m_minPoints)
        {
            kernel.push_back(i);
            unclassified.erase(
                remove(
                    unclassified.begin(), unclassified.end(), i
                ), unclassified.end()
            );
        }
    }
}

void DBSCAN::prl_calc_border(int start, int end, vector<int> &tmp)
{
    for (int i = 0; i < unclassified.size(); i++)
    {
        int index = unclassified[i];

        for (int j = 0; j < m_points.size(); j++)
            if (
                mtx[index][j] && 
                find(
                    unclassified.begin(), unclassified.end(), j
                ) == unclassified.end()
            )
            {
                border.push_back(index);
                tmp.push_back(index);
            }
    }
}

void DBSCAN::prl_calc_noise(int start, int end, vector<int> &tmp)
{
    for (int i = start; i < end; i++)
    {
        int index = unclassified[i];

        noise.push_back(index);
        tmp.push_back(index);
    }
}

void DBSCAN::prl_unclustered_init(int start, int end)
{
    for (int i = start; i < end; i++)
        unclustered.push_back(kernel[i]);
}

void DBSCAN::prl_clustering_init(int start, int end, int index, vector<int> &cluster, vector<Check> &checker)
{
    for (int i = start; i < end; i++)
        if (
            mtx[index][i] &&
            find(
                kernel.begin(), kernel.end(), i
            ) != kernel.end()
        )
        {
            cluster.push_back(i);
            checker.push_back(Check{i, true});
        }
}

void DBSCAN::prl_clustering_step(int start, int end, vector<int> &cluster, vector<Check> &checker, vector<Check> &tmp_checker)
{
    for (int i = 0; i < checker.size(); i++)
        for (int j = 0; j < m_points.size(); j++)
        {
            bool is_reachable = mtx[checker[i].index][j];
            bool in_cluster = find(cluster.begin(), cluster.end(), j) != cluster.end();
            bool is_border = find(border.begin(), border.end(), j) != border.end();
            bool is_kernel = find(kernel.begin(), kernel.end(), j) != kernel.end();

            if (is_reachable && !in_cluster)
            {
                if (is_border && checker[i].is_kernel)
                    tmp_checker.push_back(Check{j, false});
                else
                    tmp_checker.push_back(Check{j, true});

                cluster.push_back(j);
            }
        }
}