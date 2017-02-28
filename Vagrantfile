# -*- mode: ruby -*-
# vi: set ft=ruby :

unless Vagrant.has_plugin?('nugrant')
  warn "[\e[1m\e[31mERROR\e[0m]: Please run: vagrant plugin install nugrant"
  exit -1
end
unless Vagrant.has_plugin?('vagrant-hostmanager')
  warn "[\e[1m\e[31mERROR\e[0m]: Please run: vagrant plugin install vagrant-hostmanager"
  exit -1
end
unless Vagrant.has_plugin?('vagrant-docker-compose')
  warn "[\e[1m\e[31mERROR\e[0m]: Please run: vagrant plugin install vagrant-docker-compose"
  exit -1
end
unless Vagrant.has_plugin?('vagrant-guest_ansible')
  warn "[\e[1m\e[31mERROR\e[0m]: Please run: vagrant plugin install vagrant-guest_ansible"
  exit -1
end

# Network configuration
BRIDGE_NETWORK = '13.13.13.13'
BRIDGE_NETMASK = '255.255.0.0'

# Environment configuration
def setup_defaults()
  {
      :name => 'SF Movies - dev environment',
      :cpus => 1,
      :memory => 4096,
      :customize => [:modifyvm, :id,
                     '--nicpromisc2', 'allow-all',
                     '--groups', '/Dev environment'],
      :gui => false,
      :box => 'ubuntu/trusty64',
      :provision => {
          :playbook => 'ops/provision/application-host.yml',
          :docker => '/vagrant',
          :compose => '/vagrant/ops/docker-compose.yml'
      }
  }
end

Vagrant.configure(2) do |config|
  # Load configurations
  config.user.defaults = setup_defaults

  # Enable ssh forward agent
  config.ssh.forward_agent = true

  # Enable hostmanager
  config.hostmanager.enabled = true
  config.hostmanager.manage_host = true

  # Create SF Movies application servers
  # ------------------------------------------------------------------------------------------------
  config.vm.define "sfmovies", {:primary => true} do |server|

    server.vm.hostname = :sfmovies
    server.vm.box = config.user.box

    # Network
    server.vm.network :private_network, :ip => BRIDGE_NETWORK, :netmask => BRIDGE_NETMASK

    vagrant_root = File.dirname(__FILE__)
    server.vm.synced_folder vagrant_root, vagrant_root

    # VM configuration
    server.vm.provider :virtualbox do |virtualbox|
      virtualbox.name = config.user.name
      virtualbox.memory = config.user.memory
      virtualbox.cpus = config.user.cpus
      virtualbox.gui = config.user.gui
      virtualbox.customize config.user.customize
    end

    server.vm.provision :hostmanager

    # The following line terminates all ssh connections. Therefore Vagrant will be forced to reconnect.
    # That's a workaround to have the docker command in the PATH and add Vagrant to the docker group by logging in/out
    server.vm.provision :logout, type: :shell, run: :always do |shell|
      shell.inline = "ps aux | grep 'sshd:' | awk '{print end $2}' | xargs kill"
    end

    # Provision of dev environment
    provisioner = Vagrant::Util::Platform.windows? ? :guest_ansible : :ansible
    server.vm.provision :host, type: provisioner, run: :once do |ansible|
      ansible.playbook = config.user.provision.playbook
    end

    server.vm.provision :build, type: :docker do |docker|
      docker.build_image config.user.provision.docker, args: '--tag sfmovies-core'
    end

    # Setup deploy in to dev environment
    server.vm.provision :deploy, type: :docker_compose, run: :always do |compose|
      compose.yml = config.user.provision.compose
      compose.rebuild = true
    end
  end
end
