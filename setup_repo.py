import subprocess


def RunCommandEachMachine(cmd):
    netid = "siyuanc3"
    for i in range(1, 11):
        hostname = "fa22-cs425-80%02d.cs.illinois.edu" % (i,)
        print("On " + hostname + "\n")
        sshProcess = subprocess.Popen(['ssh',
                                "{}@{}".format(netid, hostname),
                                cmd]).communicate()
    

if __name__  == "__main__":
    clone = 'if cd ~/mp1; then git pull; else git clone git@gitlab.engr.illinois.edu:siyuan-ruiqi/mp1.git ~/mp1; fi'
    save_to_known_hosts = "ssh-keyscan gitlab.engr.illinois.edu >> ~/.ssh/known_hosts"
    RunCommandEachMachine(cmd = clone)
    # RunCommandEachMachine(cmd = save_to_known_hosts)

