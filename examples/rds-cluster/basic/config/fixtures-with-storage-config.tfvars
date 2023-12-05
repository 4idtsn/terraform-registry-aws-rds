aws_region = "us-east-1"
is_enabled = true

cluster_config = [
  {
    cluster_identifier = "test-cluster-1"
    engine_mode        = "provisioned"
  }
]

cluster_backup_config = [
  {
    cluster_identifier  = "test-cluster-1"
    skip_final_snapshot = true
  },
]

cluster_storage_config = [
  {
    cluster_identifier = "test-cluster-1"
    allocated_storage  = 100
    storage_type       = "io1"
  },
]
