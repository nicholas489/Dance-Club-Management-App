package util

// Member structure
type Member struct {
	FullName    string
	PhoneNumber string
	Address     string
	Attended    int
	Paid        int
}

// sort that does it by attends
type ByAttended []Member

// needed for the sort.Interface
func (a ByAttended) Len() int           { return len(a) }
func (a ByAttended) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAttended) Less(i, j int) bool { return a[i].Attended < a[j].Attended }

// sort by who paid the least to most
type ByPaid []Member

// needed for the sort.Interface
func (a ByPaid) Len() int           { return len(a) }
func (a ByPaid) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByPaid) Less(i, j int) bool { return a[i].Paid < a[j].Paid }
