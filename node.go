package ptreader

import (
	"io"
	"time"
)

// Moment that happened
type Moment struct {
	Unixtime        int64
	Activeseconds   int
	Semiidleseconds int
	Key             int
	Lmb             int //left mouse button
	Rmb             int //right mouse button
	Scrollwheel     int
}

// Node record
type Node struct {
	Name     string
	TagIndex int
	Hidden   byte
	Moments  []Moment
	SubNodes []Node
}

func (r *Reader) readMoment() (Moment, error) {
	var mnt Moment

	date, err := r.readUint16()
	if err != nil {
		return mnt, err
	}

	firstMins, err := r.readUint16()
	if err != nil {
		return mnt, err
	}

	year := int(2000 + (date >> 9))
	month := int(date >> 5 & 0xF)
	day := int(date & 0x1F)
	hour := int(firstMins / 60)
	minute := int(firstMins - 60*(firstMins/60))

	mnt.Unixtime = time.Date(year, time.Month(month), day, hour, minute, 0, 0, time.Local).Unix()

	mnt.Activeseconds, err = r.readInt()
	if err != nil {
		return mnt, err
	}

	mnt.Semiidleseconds, err = r.readInt()
	if err != nil {
		return mnt, err
	}

	mnt.Key, err = r.readInt()
	if err != nil {
		return mnt, err
	}

	mnt.Lmb, err = r.readInt()
	if err != nil {
		return mnt, err
	}

	mnt.Rmb, err = r.readInt()
	if err != nil {
		return mnt, err
	}

	mnt.Scrollwheel, err = r.readInt()
	if err != nil {
		return mnt, err
	}
	return mnt, nil
}

func (r *Reader) readNode() (Node, error) {
	var n Node
	var err error
	n.Name, err = r.readString()
	if err != nil {
		return n, err
	}

	n.TagIndex, err = r.readInt()
	if err != nil {
		return n, err
	}

	n.Hidden, err = r.readByte()
	if err != nil {
		return n, err
	}

	numDays, err := r.readInt()
	if err != nil {
		return n, err
	}

	var monents []Moment
	for i := 0; i < numDays; i++ {
		mnt, err := r.readMoment()
		if err != nil {
			return n, err
		}
		monents = append(monents, mnt)
	}
	n.Moments = monents

	numChildren, err := r.readInt()
	if err != nil && err == io.EOF {
		return n, nil
	}

	var subNodes []Node
	for i := 0; i < numChildren; i++ {
		one, err := r.readNode()
		if err != nil {
			return one, err
		}
		subNodes = append(subNodes, one)
	}
	n.SubNodes = subNodes

	return n, nil
}

// GetNodeTime to get total seconds on that node
func GetNodeTime(n *Node) int64 {
	seconds := int64(0)

	mnts := n.Moments
	for i := 0; i < len(mnts); i++ {
		seconds += int64(mnts[i].Activeseconds)
	}

	for i := 0; i < len(n.SubNodes); i++ {
		seconds += GetNodeTime(&n.SubNodes[i])
	}
	return seconds
}
