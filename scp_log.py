import subprocess


def ScpToRemote():
    netid = "siyuanc3"
    for i in range(1, 11):
        hostname = "fa22-cs425-80%02d.cs.illinois.edu" % (i,)
        filename = "MP1DemoDataFA22/MP1 Demo Data FA22/vm%d.log" % i
        print("On " + hostname + "\n")
        sshProcess = subprocess.Popen(["scp",
                                filename,
                                "{}@{}:mp1/vm{}.log".format(netid, hostname, i)]).communicate()
    

if __name__  == "__main__":
    # clone = 'if cd ~/mp1; then git pull; else git clone git@gitlab.engr.illinois.edu:siyuan-ruiqi/mp1.git ~/mp1; fi'
    # save_to_known_hosts = "ssh-keyscan gitlab.engr.illinois.edu >> ~/.ssh/known_hosts"
    ScpToRemote()
    # RunCommandEachMachine(cmd = save_to_known_hosts)

