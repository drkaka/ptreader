package ptreader

import "fmt"

func (r *Reader) readMeta() error {
	if _, err := r.readVersion(); err != nil {
		return err
	}

	if _, err := r.readMagic(); err != nil {
		return err
	}

	if _, _, err := r.readTags(); err != nil {
		return err
	}

	if _, err := r.readMinfilter(); err != nil {
		return err
	}

	if _, err := r.readFoldLevel(); err != nil {
		return err
	}

	if _, err := r.readPrefs(); err != nil {
		return err
	}

	return nil
}

func (r *Reader) readVersion() (int, error) {
	version, err := r.readInt()
	if err != nil {
		return 0, err
	}
	if version < 10 {
		return 0, fmt.Errorf("db version not supported")
	}
	return version, nil
}

func (r *Reader) readMagic() (int, error) {
	magic, err := r.readInt()
	if err != nil {
		return 0, err
	} else if magic != 1347700294 {
		return 0, fmt.Errorf("db is wrong")
	}
	return magic, err
}

// return tag names, colors, error
func (r *Reader) readTags() ([]string, []int, error) {
	num, err := r.readInt()
	if err != nil {
		return nil, nil, err
	}

	var tags []string
	var colors []int
	oneTag := make([]byte, 32)
	for i := 0; i < num; i++ {
		n, err := r.r.Read(oneTag)
		if n != 32 {
			return nil, nil, fmt.Errorf("read tag wrong")
		}
		tags = append(tags, string(oneTag))

		cl, err := r.readInt()
		if err != nil {
			return nil, nil, err
		}
		colors = append(colors, cl)
	}
	return tags, colors, nil
}

func (r *Reader) readMinfilter() (int, error) {
	minfilter, err := r.readInt()
	if err != nil {
		return 0, err
	}
	return minfilter, err
}

func (r *Reader) readFoldLevel() (int, error) {
	foldLevel, err := r.readInt()
	if err != nil {
		return 0, err
	}
	return foldLevel, err
}

func (r *Reader) readPrefs() ([6]int, error) {
	var prefs [6]int

	for i := 0; i < 6; i++ {
		pref, err := r.readInt()
		if err != nil {
			return prefs, err
		}
		prefs[i] = pref
	}

	return prefs, nil
}
