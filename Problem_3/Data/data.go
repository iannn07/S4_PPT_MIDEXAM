package Data

type Data struct {
	Nama   string
	NIM    int
	Alamat string
}

var students = []Data{
	{Nama: "Ian", NIM: 2501983105, Alamat: "BINUS"},
	{Nama: "Budi", NIM: 251983106, Alamat: "MID EXAM"},
	{Nama: "Dharmawan", NIM: 1234567890, Alamat: "PPT"},
}

func SingleData(name string) []Data {
	for _, student := range students {
		if student.Nama == name {
			return []Data{student}
		}
	}
	return []Data{}
}

func AllData() []Data {
	return students
}
