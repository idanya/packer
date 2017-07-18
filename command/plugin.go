//
// This file is automatically generated by scripts/generate-plugins.go -- Do not edit!
//

package command

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/idanya/packer/packer"
	"github.com/idanya/packer/packer/plugin"

	alicloudecsbuilder "github.com/idanya/packer/builder/alicloud/ecs"
	amazonchrootbuilder "github.com/idanya/packer/builder/amazon/chroot"
	amazonebsbuilder "github.com/idanya/packer/builder/amazon/ebs"
	amazonebssurrogatebuilder "github.com/idanya/packer/builder/amazon/ebssurrogate"
	amazonebsvolumebuilder "github.com/idanya/packer/builder/amazon/ebsvolume"
	amazoninstancebuilder "github.com/idanya/packer/builder/amazon/instance"
	azurearmbuilder "github.com/idanya/packer/builder/azure/arm"
	cloudstackbuilder "github.com/idanya/packer/builder/cloudstack"
	digitaloceanbuilder "github.com/idanya/packer/builder/digitalocean"
	dockerbuilder "github.com/idanya/packer/builder/docker"
	filebuilder "github.com/idanya/packer/builder/file"
	googlecomputebuilder "github.com/idanya/packer/builder/googlecompute"
	hypervisobuilder "github.com/idanya/packer/builder/hyperv/iso"
	nullbuilder "github.com/idanya/packer/builder/null"
	oneandonebuilder "github.com/idanya/packer/builder/oneandone"
	openstackbuilder "github.com/idanya/packer/builder/openstack"
	parallelsisobuilder "github.com/idanya/packer/builder/parallels/iso"
	parallelspvmbuilder "github.com/idanya/packer/builder/parallels/pvm"
	profitbricksbuilder "github.com/idanya/packer/builder/profitbricks"
	qemubuilder "github.com/idanya/packer/builder/qemu"
	tritonbuilder "github.com/idanya/packer/builder/triton"
	virtualboxisobuilder "github.com/idanya/packer/builder/virtualbox/iso"
	virtualboxovfbuilder "github.com/idanya/packer/builder/virtualbox/ovf"
	vmwareisobuilder "github.com/idanya/packer/builder/vmware/iso"
	vmwarevmxbuilder "github.com/idanya/packer/builder/vmware/vmx"
	alicloudimportpostprocessor "github.com/idanya/packer/post-processor/alicloud-import"
	amazonimportpostprocessor "github.com/idanya/packer/post-processor/amazon-import"
	artificepostprocessor "github.com/idanya/packer/post-processor/artifice"
	atlaspostprocessor "github.com/idanya/packer/post-processor/atlas"
	checksumpostprocessor "github.com/idanya/packer/post-processor/checksum"
	compresspostprocessor "github.com/idanya/packer/post-processor/compress"
	dockerimportpostprocessor "github.com/idanya/packer/post-processor/docker-import"
	dockerpushpostprocessor "github.com/idanya/packer/post-processor/docker-push"
	dockersavepostprocessor "github.com/idanya/packer/post-processor/docker-save"
	dockertagpostprocessor "github.com/idanya/packer/post-processor/docker-tag"
	googlecomputeexportpostprocessor "github.com/idanya/packer/post-processor/googlecompute-export"
	manifestpostprocessor "github.com/idanya/packer/post-processor/manifest"
	shelllocalpostprocessor "github.com/idanya/packer/post-processor/shell-local"
	vagrantpostprocessor "github.com/idanya/packer/post-processor/vagrant"
	vagrantcloudpostprocessor "github.com/idanya/packer/post-processor/vagrant-cloud"
	vspherepostprocessor "github.com/idanya/packer/post-processor/vsphere"
	ansibleprovisioner "github.com/idanya/packer/provisioner/ansible"
	ansiblelocalprovisioner "github.com/idanya/packer/provisioner/ansible-local"
	chefclientprovisioner "github.com/idanya/packer/provisioner/chef-client"
	chefsoloprovisioner "github.com/idanya/packer/provisioner/chef-solo"
	convergeprovisioner "github.com/idanya/packer/provisioner/converge"
	fileprovisioner "github.com/idanya/packer/provisioner/file"
	powershellprovisioner "github.com/idanya/packer/provisioner/powershell"
	puppetmasterlessprovisioner "github.com/idanya/packer/provisioner/puppet-masterless"
	puppetserverprovisioner "github.com/idanya/packer/provisioner/puppet-server"
	saltmasterlessprovisioner "github.com/idanya/packer/provisioner/salt-masterless"
	shellprovisioner "github.com/idanya/packer/provisioner/shell"
	shelllocalprovisioner "github.com/idanya/packer/provisioner/shell-local"
	windowsrestartprovisioner "github.com/idanya/packer/provisioner/windows-restart"
	windowsshellprovisioner "github.com/idanya/packer/provisioner/windows-shell"
)

type PluginCommand struct {
	Meta
}

var Builders = map[string]packer.Builder{
	"alicloud-ecs":        new(alicloudecsbuilder.Builder),
	"amazon-chroot":       new(amazonchrootbuilder.Builder),
	"amazon-ebs":          new(amazonebsbuilder.Builder),
	"amazon-ebssurrogate": new(amazonebssurrogatebuilder.Builder),
	"amazon-ebsvolume":    new(amazonebsvolumebuilder.Builder),
	"amazon-instance":     new(amazoninstancebuilder.Builder),
	"azure-arm":           new(azurearmbuilder.Builder),
	"cloudstack":          new(cloudstackbuilder.Builder),
	"digitalocean":        new(digitaloceanbuilder.Builder),
	"docker":              new(dockerbuilder.Builder),
	"file":                new(filebuilder.Builder),
	"googlecompute":       new(googlecomputebuilder.Builder),
	"hyperv-iso":          new(hypervisobuilder.Builder),
	"null":                new(nullbuilder.Builder),
	"oneandone":           new(oneandonebuilder.Builder),
	"openstack":           new(openstackbuilder.Builder),
	"parallels-iso":       new(parallelsisobuilder.Builder),
	"parallels-pvm":       new(parallelspvmbuilder.Builder),
	"profitbricks":        new(profitbricksbuilder.Builder),
	"qemu":                new(qemubuilder.Builder),
	"triton":              new(tritonbuilder.Builder),
	"virtualbox-iso":      new(virtualboxisobuilder.Builder),
	"virtualbox-ovf":      new(virtualboxovfbuilder.Builder),
	"vmware-iso":          new(vmwareisobuilder.Builder),
	"vmware-vmx":          new(vmwarevmxbuilder.Builder),
}

var Provisioners = map[string]packer.Provisioner{
	"ansible":           new(ansibleprovisioner.Provisioner),
	"ansible-local":     new(ansiblelocalprovisioner.Provisioner),
	"chef-client":       new(chefclientprovisioner.Provisioner),
	"chef-solo":         new(chefsoloprovisioner.Provisioner),
	"converge":          new(convergeprovisioner.Provisioner),
	"file":              new(fileprovisioner.Provisioner),
	"powershell":        new(powershellprovisioner.Provisioner),
	"puppet-masterless": new(puppetmasterlessprovisioner.Provisioner),
	"puppet-server":     new(puppetserverprovisioner.Provisioner),
	"salt-masterless":   new(saltmasterlessprovisioner.Provisioner),
	"shell":             new(shellprovisioner.Provisioner),
	"shell-local":       new(shelllocalprovisioner.Provisioner),
	"windows-restart":   new(windowsrestartprovisioner.Provisioner),
	"windows-shell":     new(windowsshellprovisioner.Provisioner),
}

var PostProcessors = map[string]packer.PostProcessor{
	"alicloud-import":      new(alicloudimportpostprocessor.PostProcessor),
	"amazon-import":        new(amazonimportpostprocessor.PostProcessor),
	"artifice":             new(artificepostprocessor.PostProcessor),
	"atlas":                new(atlaspostprocessor.PostProcessor),
	"checksum":             new(checksumpostprocessor.PostProcessor),
	"compress":             new(compresspostprocessor.PostProcessor),
	"docker-import":        new(dockerimportpostprocessor.PostProcessor),
	"docker-push":          new(dockerpushpostprocessor.PostProcessor),
	"docker-save":          new(dockersavepostprocessor.PostProcessor),
	"docker-tag":           new(dockertagpostprocessor.PostProcessor),
	"googlecompute-export": new(googlecomputeexportpostprocessor.PostProcessor),
	"manifest":             new(manifestpostprocessor.PostProcessor),
	"shell-local":          new(shelllocalpostprocessor.PostProcessor),
	"vagrant":              new(vagrantpostprocessor.PostProcessor),
	"vagrant-cloud":        new(vagrantcloudpostprocessor.PostProcessor),
	"vsphere":              new(vspherepostprocessor.PostProcessor),
}

var pluginRegexp = regexp.MustCompile("packer-(builder|post-processor|provisioner)-(.+)")

func (c *PluginCommand) Run(args []string) int {
	// This is an internal call (users should not call this directly) so we're
	// not going to do much input validation. If there's a problem we'll often
	// just crash. Error handling should be added to facilitate debugging.
	log.Printf("args: %#v", args)
	if len(args) != 1 {
		c.Ui.Error("Wrong number of args")
		return 1
	}

	// Plugin will match something like "packer-builder-amazon-ebs"
	parts := pluginRegexp.FindStringSubmatch(args[0])
	if len(parts) != 3 {
		c.Ui.Error(fmt.Sprintf("Error parsing plugin argument [DEBUG]: %#v", parts))
		return 1
	}
	pluginType := parts[1] // capture group 1 (builder|post-processor|provisioner)
	pluginName := parts[2] // capture group 2 (.+)

	server, err := plugin.Server()
	if err != nil {
		c.Ui.Error(fmt.Sprintf("Error starting plugin server: %s", err))
		return 1
	}

	switch pluginType {
	case "builder":
		builder, found := Builders[pluginName]
		if !found {
			c.Ui.Error(fmt.Sprintf("Could not load builder: %s", pluginName))
			return 1
		}
		server.RegisterBuilder(builder)
	case "provisioner":
		provisioner, found := Provisioners[pluginName]
		if !found {
			c.Ui.Error(fmt.Sprintf("Could not load provisioner: %s", pluginName))
			return 1
		}
		server.RegisterProvisioner(provisioner)
	case "post-processor":
		postProcessor, found := PostProcessors[pluginName]
		if !found {
			c.Ui.Error(fmt.Sprintf("Could not load post-processor: %s", pluginName))
			return 1
		}
		server.RegisterPostProcessor(postProcessor)
	}

	server.Serve()

	return 0
}

func (*PluginCommand) Help() string {
	helpText := `
Usage: packer plugin PLUGIN

  Runs an internally-compiled version of a plugin from the packer binary.

  NOTE: this is an internal command and you should not call it yourself.
`

	return strings.TrimSpace(helpText)
}

func (c *PluginCommand) Synopsis() string {
	return "internal plugin command"
}
