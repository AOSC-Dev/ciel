package main

import (
	"bytes"
	"ciel-driver"
)

func dpkgPackageFiles(c *ciel.Container) map[string]bool {
	stdout := new(bytes.Buffer)
	if exitCode := c.CommandRaw(ciel.ShellPath, nil, stdout, nil, "-l", "-c",
		`dpkg-query --listfiles $(dpkg-query --show --showformat=\$"{Package}\n")`,
	); exitCode != 0 {
		return nil
	}
	hashmap := make(map[string]bool, 100000)
	dataset := bytes.Split(stdout.Bytes(), []byte{'\n'})
	root := c.Fs.TargetDir()
	for _, record := range dataset {
		if len(record) == 0 {
			continue
		}
		path := string(record)
		evalSymlinksCleanCache()
		path = evalSymlinks(root, path, true)
		hashmap[path] = true
	}
	return hashmap
}