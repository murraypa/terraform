resource_group_name = "myResourceGroup"
location            = "eastus"
cluster_name        = "myAKSCluster"
kubernetes_version  = "1.20.7"
node_count          = 3
node_vm_size        = "Standard_D2_v2"
node_pool_names     = ["pool1", "pool2"]
os_disk_size_gb     = 128
vnet_subnet_id      = "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/myVnet/subnets/mySubnet"
cluster {
  name = "my-aks-cluster"
  location = "westus2"

  # Create a primary node pool
  node_pool {
    name       = "np-primary"
    vm_size   = "Standard_D4s_v3"
    count     = 3
    max_pods = 100
    min_count = 2
    type        = "VirtualMachineScaleSets"
  }

  # Create a secondary node pool
  node_pool {
    name       = "np-secondary"
    vm_size   = "Standard_D4s_v3"
    count     = 2
    max_pods = 50
    min_count = 2
    type        = "VirtualMachineScaleSets"
  }
}
