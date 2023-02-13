#pragma once

#include "TCodec.h"
#include "TServant.h"
#include "K8SParams.h"

struct TServer
{
    bool hostIPC;
    bool autoRelease;
    int replicas;
    int maxReplicas;
    int minReplicas;
    int asyncThread;
    std::string resourceName;
    std::string appName;
    std::string serverName;
    std::string templateName;
    std::string profileContent;
    std::string runningHost;
    std::vector <TServant> servants;
};

const static auto FixedEnvs = boost::json::array({
        {
                { "name", "Namespace" },
                { "valueFrom", boost::json::object{
                        { "fieldRef", {
                                { "fieldPath", "metadata.namespace" },
                        },
                        },
                }
                }
        },

        {
                { "name", "PodName" },
                { "valueFrom", boost::json::object{
                        { "fieldRef", {
                                { "fieldPath", "metadata.name" },
                        },
                        },
                },
                },
        },

        {
                { "name", "PodIP" },
                { "valueFrom", boost::json::object{
                        { "fieldRef", {
                                { "fieldPath", "status.podIP" },
                        },
                        },
                },
                },
        },

        {

                { "name", "ServerApp" },
                { "valueFrom", boost::json::object{
                        { "fieldRef", {
                                { "fieldPath", "metadata.labels['tars.io/ServerApp']" },
                                { "apiVersion", "v1" }
                        },
                        },
                },
                },
        }});

const static auto FixedMounts = boost::json::array({
        {
                { "mountPath", "/usr/local/app/tars/app_log" },
                { "name", "host-log-dir" },
                { "subPathExpr", "$(Namespace)/$(PodName)" },
                {
                        "source", boost::json::object{
                        {
                                "hostPath", {
                                { "path", "/usr/local/app/tars/app_log" },
                                { "type", "DirectoryOrCreate" },
                        }, }
                }
                }
        }
});


ENCODE_STRUCT_TO_JSON(TServer, s, j)
{
    j = {
            { "apiVersion", API_GROUP_VERSION },
            { "kind",       "TServer" },
            {
              "metadata",   {
                                    { "namespace", K8SParams::Namespace() },
                                    { "name",   s.resourceName },
                                    { "annotation", boost::json::object{}}
                            }
            },
            {
              "spec",       {
                                    { "app",       s.appName },
                                    { "server", s.serverName },
                                    { "subType",    "tars" },
                                    { "tars", boost::json::object{
                                            { "template",    s.templateName },
                                            { "profile",     s.profileContent },
                                            { "asyncThread", s.asyncThread },
                                            { "servants",    boost::json::value_from(s.servants) },
                                    }},
                                    { "k8s", boost::json::object{
                                            { "env",       FixedEnvs },
                                            { "mounts",    FixedMounts },
                                            { "replicase", s.replicas },
                                    }
                                    }
                            }
            }
    };
    if (s.maxReplicas != 0)
    {
        j.at_pointer("/metadata/annotation").get_object().insert(
                {{ "tars.io/MaxReplicas", std::to_string(s.maxReplicas) }});
    }
    if (s.autoRelease)
    {
        j.at_pointer("/metadata/annotation").get_object().insert(
                {{ "tars.io/AutoRelease", "true" }});
    }
    if (s.hostIPC)
    {
        j.at_pointer("/spec/k8s").get_object().insert({{ "hostIPC", true }});
    }

    if (!s.runningHost.empty())
    {
    }
}
