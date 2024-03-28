package merkledag

func Hash2File(store KVStore, hash []byte, path string, hp HashPool) []byte {
	treeData, _ := store.Get(hash)
	if treeData == nil {
		return nil
	}

	entries, err := parseTree(treeData)
	if err != nil {
		return nil
	}

	var fileHash []byte
	for _, entry := range entries {
		if entry.Path == path {
			fileHash = entry.Hash
			break
		}
	}

	if fileHash == nil {
		return nil
	}

	fileContent, _ := store.Get(fileHash)
	if fileContent == nil {
		return nil
	}

	return fileContent
}

func parseTree(data []byte) ([]TreeEntry, error) {

	return []TreeEntry{}, nil
}

type TreeEntry struct {
	Path string
	Hash []byte
}

