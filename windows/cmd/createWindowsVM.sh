#!  /bin/sh

#   CGE VM parameters
export GCE_project=curso-docker-completo

export GCE_zone=southamerica-east1-a
export GCE_vmInstanceName=docker-builder
export GCE_vmMachineType=e2-standard-8
export GCE_vmImageProject=windows-cloud
export GCE_vmImageFamily=windows-2019-core
export GCE_vmDiskSpace=256
export GCE_vmDiskType=pd-standard

gcloud config set project ${GCE_project}

gcloud config set compute/zone ${GCE_zone}

gcloud compute instances create ${GCE_vmInstanceName} \
    --machine-type=${GCE_vmMachineType} \
    --enable-display-device \
    --can-ip-forward \
    --tags=https-server \
    --image-project=${GCE_vmImageProject} \
    --image-family=${GCE_vmImageFamily} \
    --boot-disk-size=${GCE_vmDiskSpace} \
    --boot-disk-type=${GCE_vmDiskType}

gcloud beta compute \
    --project ${GCE_project} \
    reset-windows-password ${GCE_vmInstanceName} \
    --zone ${GCE_zone}
