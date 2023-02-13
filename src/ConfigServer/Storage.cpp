
#include "Storage.h"
#include "util/tc_thread_rwlock.h"
#include "TCodec.h"
#include <string>

static int extractPodSeq(const std::string& sPodName, const std::string& sGenerateName)
{
    try
    {
        assert(sPodName.size() > sGenerateName.size());
        auto sPodSeq = sPodName.substr(sGenerateName.size());
        return std::stoi(sPodSeq, nullptr, 10);
    }
    catch (std::exception& exception)
    {
        return -1;
    }
}

class StorageImp
{
public:
    static StorageImp& instance()
    {
        static StorageImp imp;
        return imp;
    }

    StorageImp(const StorageImp&) = delete;

    ~StorageImp() = default;

private:
    StorageImp() = default;

public:
    tars::TC_ThreadRWLocker mutex_;
    std::unordered_map<std::string, int> seqMap_;
    std::unordered_map<std::string, int> cacheSeqMap_;
};

void Storage::onPodAdded(const boost::json::value& v, K8SWatchEventDrive drive)
{
    try
    {
        VAR_FROM_JSON(std::string, generate, v.at_pointer("/metadata/generateName"));
        if (generate.empty())
        {
            return;
        }

        VAR_FROM_JSON(std::string, name, v.at_pointer("/metadata/name"));
        if (name.empty())
        {
            return;
        }

        int seq = extractPodSeq(name, generate);
        if (seq == -1)
        {
            return;
        }

        auto domain = name + "." + generate.substr(0, generate.size() - 1);

        VAR_FROM_JSON(std::string, ip, v.at_pointer("/status/podIP"));
        if (drive == K8SWatchEventDrive::List)
        {
            StorageImp::instance().cacheSeqMap_[name] = seq;
            StorageImp::instance().cacheSeqMap_[domain] = seq;
            if (!ip.empty())
            {
                StorageImp::instance().cacheSeqMap_[ip] = seq;
            }
        }
        else if (drive == K8SWatchEventDrive::Watch)
        {
            StorageImp::instance().mutex_.writeLock();
            StorageImp::instance().seqMap_[name] = seq;
            StorageImp::instance().seqMap_[domain] = seq;
            if (!ip.empty())
            {
                StorageImp::instance().seqMap_[ip] = seq;
            }
            StorageImp::instance().mutex_.unWriteLock();
        }
    }
    catch (...)
    {
    }
}

void Storage::onPodModified(const boost::json::value& v)
{
    onPodAdded(v, K8SWatchEventDrive::Watch);
}

void Storage::onPodDelete(const boost::json::value& v)
{
    try
    {
        VAR_FROM_JSON(std::string, generate, v.at_pointer("/metadata/generateName"));
        if (generate.empty())
        {
            return;
        }

        VAR_FROM_JSON(std::string, name, v.at_pointer("/metadata/name"));
        if (name.empty())
        {
            return;
        }

        auto domain = name + "." + generate.substr(0, generate.size() - 1);

        VAR_FROM_JSON(std::string, ip, v.at_pointer("/status/podIP"));

        StorageImp::instance().mutex_.writeLock();
        StorageImp::instance().seqMap_.erase(name);
        StorageImp::instance().seqMap_.erase(domain);
        if (!ip.empty())
        {
            StorageImp::instance().seqMap_.erase(ip);
        }
        StorageImp::instance().mutex_.unWriteLock();
    }
    catch (...)
    {
    }
}

void Storage::prePodList()
{
    StorageImp::instance().cacheSeqMap_.clear();
}

void Storage::postPodList()
{
    {
        StorageImp::instance().mutex_.writeLock();
        std::swap(StorageImp::instance().cacheSeqMap_, StorageImp::instance().seqMap_);
        StorageImp::instance().mutex_.unWriteLock();
    }
    StorageImp::instance().cacheSeqMap_.clear();
}

void Storage::getSeqMap(const std::function<void(const std::unordered_map<std::string, int>&)>& f)
{
    StorageImp::instance().mutex_.readLock();
    f(StorageImp::instance().seqMap_);
    StorageImp::instance().mutex_.unReadLock();
}
