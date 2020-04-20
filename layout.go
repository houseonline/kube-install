package main

import (
    "os"
)


func installYML(softdir string) {
    install_file, err := os.Create(softdir+"/workflow/k8scluster-install.yml") //新建inventory配置文件
    checkErr(err)
    defer install_file.Close() //main函数结束前， 关闭文件
    install_file.WriteString("- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/0.base/genfile\n")
    install_file.WriteString("- remote_user: root\n  hosts: node\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/0.base/kernel\n")
    install_file.WriteString("- remote_user: root\n  hosts: master,node,nginx\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/0.base/all\n")
    install_file.WriteString("- remote_user: root\n  hosts: node\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/8.action/delnode\n")
    install_file.WriteString("- remote_user: root\n  hosts: node\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/0.base/docker\n")
    install_file.WriteString("- remote_user: root\n  hosts: master1\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/1.cfssl/copycfssl\n")
    install_file.WriteString("- remote_user: root\n  hosts: master1\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/1.cfssl/createssl\n")
    install_file.WriteString("- remote_user: root\n  hosts: etcd\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/2.etcd\n")
    install_file.WriteString("- remote_user: root\n  hosts: master1\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/3.network/etcd_network\n")
    install_file.WriteString("- remote_user: root\n  hosts: node\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/3.network/flanneld\n")
    install_file.WriteString("- remote_user: root\n  hosts: nginx\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/4.kube-nginx/nginx\n")
    install_file.WriteString("- remote_user: root\n  hosts: master,node\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/5.kube-master/kubectl\n")
    install_file.WriteString("- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/5.kube-master/apiserver\n")
    install_file.WriteString("- remote_user: root\n  hosts: nginx\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/4.kube-nginx/keepalived\n")
    install_file.WriteString("- remote_user: root\n  hosts: master1\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/5.kube-master/api-rbac\n")
    install_file.WriteString("- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/5.kube-master/controller-manager\n")
    install_file.WriteString("- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/5.kube-master/scheduler\n")
    install_file.WriteString("- remote_user: root\n  hosts: node\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/6.kube-node/kubelet\n")
    install_file.WriteString("- remote_user: root\n  hosts: master1\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/6.kube-node/approve-csr\n")
    install_file.WriteString("- remote_user: root\n  hosts: node\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/6.kube-node/kube-proxy\n")
    install_file.WriteString("- remote_user: root\n  hosts: master1\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/7.kube-addons\n")
    install_file.WriteString("- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/5.kube-master/admintoken\n")
    install_file.WriteString("- remote_user: root\n  hosts: nginx\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/4.kube-nginx/alter\n")
    install_file.WriteString("- remote_user: root\n  hosts: node\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/9.finish/install\n")
}

func onemasterinstallYML(softdir string) {
    onemasterinstall_file, err := os.Create(softdir+"/workflow/k8scluster-onemasterinstall.yml") //新建inventory配置文件
    checkErr(err)
    defer onemasterinstall_file.Close() //main函数结束前， 关闭文件
    onemasterinstall_file.WriteString("- remote_user: root\n  hosts: node\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/0.base/kernel\n")
    onemasterinstall_file.WriteString("- remote_user: root\n  hosts: master1,node\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/0.base/all\n")
    onemasterinstall_file.WriteString("- remote_user: root\n  hosts: node\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/8.action/delnode\n")
    onemasterinstall_file.WriteString("- remote_user: root\n  hosts: node\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/0.base/docker\n")
    onemasterinstall_file.WriteString("- remote_user: root\n  hosts: master1\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/1.cfssl/copycfssl\n")
    onemasterinstall_file.WriteString("- remote_user: root\n  hosts: master1\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/1.cfssl/createssl\n")
    onemasterinstall_file.WriteString("- remote_user: root\n  hosts: etcd\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/2.etcd\n")
    onemasterinstall_file.WriteString("- remote_user: root\n  hosts: master1\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/3.network/etcd_network\n")
    onemasterinstall_file.WriteString("- remote_user: root\n  hosts: node\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/3.network/flanneld\n")
    onemasterinstall_file.WriteString("- remote_user: root\n  hosts: master1,node\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/5.kube-master/kubectl\n")
    onemasterinstall_file.WriteString("- remote_user: root\n  hosts: master1\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/5.kube-master/apiserver\n")
    onemasterinstall_file.WriteString("- remote_user: root\n  hosts: master1\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/5.kube-master/api-rbac\n")
    onemasterinstall_file.WriteString("- remote_user: root\n  hosts: master1\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/5.kube-master/controller-manager\n")
    onemasterinstall_file.WriteString("- remote_user: root\n  hosts: master1\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/5.kube-master/scheduler\n")
    onemasterinstall_file.WriteString("- remote_user: root\n  hosts: node\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/6.kube-node/kubelet\n")
    onemasterinstall_file.WriteString("- remote_user: root\n  hosts: master1\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/6.kube-node/approve-csr\n")
    onemasterinstall_file.WriteString("- remote_user: root\n  hosts: node\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/6.kube-node/kube-proxy\n")
    onemasterinstall_file.WriteString("- remote_user: root\n  hosts: master1\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/7.kube-addons\n")
    onemasterinstall_file.WriteString("- remote_user: root\n  hosts: master1\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/5.kube-master/admintoken\n")
    onemasterinstall_file.WriteString("- remote_user: root\n  hosts: node\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/9.finish/install\n")
}

func addnodeYML(softdir string) {
    addnode_file, err := os.Create(softdir+"/workflow/k8scluster-addnode.yml") //新建inventory配置文件
    checkErr(err)
    defer addnode_file.Close() //main函数结束前， 关闭文件
    addnode_file.WriteString("- remote_user: root\n  hosts: addnode\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/0.base/kernel\n")
    addnode_file.WriteString("- remote_user: root\n  hosts: addnode\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/0.base/all\n")
    addnode_file.WriteString("- remote_user: root\n  hosts: addnode\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/8.action/delnode\n")
    addnode_file.WriteString("- remote_user: root\n  hosts: addnode\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/0.base/docker\n")
    addnode_file.WriteString("- remote_user: root\n  hosts: addnode\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/3.network/flanneld\n")
    addnode_file.WriteString("- remote_user: root\n  hosts: addnode\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/6.kube-node/kubelet\n")
    addnode_file.WriteString("- remote_user: root\n  hosts: addnode\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/6.kube-node/kube-proxy\n")
    addnode_file.WriteString("- remote_user: root\n  hosts: addnode\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/9.finish/addnode\n")
}

func delnodeYML(softdir string) {
    delnode_file, err := os.Create(softdir+"/workflow/k8scluster-delnode.yml") //新建inventory配置文件
    checkErr(err)
    defer delnode_file.Close() //main函数结束前， 关闭文件
    delnode_file.WriteString("- remote_user: root\n  hosts: delnode\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/8.action/delnode\n")
}

func rebuildmasterYML(softdir string) {
    rebuildmaster_file, err := os.Create(softdir+"/workflow/k8scluster-rebuildmaster.yml") //新建inventory配置文件
    checkErr(err)
    defer rebuildmaster_file.Close() //main函数结束前， 关闭文件
    rebuildmaster_file.WriteString("- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/0.base/genfile\n")
    rebuildmaster_file.WriteString("- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/0.base/all\n")
    rebuildmaster_file.WriteString("- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/1.cfssl/copycfssl\n")
    rebuildmaster_file.WriteString("- remote_user: root\n  hosts: etcd\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/2.etcd\n")
    rebuildmaster_file.WriteString("- remote_user: root\n  hosts: nginx\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/4.kube-nginx/nginx\n")
    rebuildmaster_file.WriteString("- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/5.kube-master/kubectl\n")
    rebuildmaster_file.WriteString("- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/5.kube-master/apiserver\n")
    rebuildmaster_file.WriteString("- remote_user: root\n  hosts: nginx\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/4.kube-nginx/keepalived\n")
    rebuildmaster_file.WriteString("- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/5.kube-master/controller-manager\n")
    rebuildmaster_file.WriteString("- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/5.kube-master/scheduler\n")
    rebuildmaster_file.WriteString("- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/5.kube-master/admintoken\n")
    rebuildmaster_file.WriteString("- remote_user: root\n  hosts: nginx\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/4.kube-nginx/alter\n")
}

func delmasterYML(softdir string) {
    delmaster_file, err := os.Create(softdir+"/workflow/k8scluster-delmaster.yml") //新建inventory配置文件
    checkErr(err)
    defer delmaster_file.Close() //main函数结束前， 关闭文件
    delmaster_file.WriteString("- remote_user: root\n  hosts: delmaster\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/8.action/delmaster\n")
}
