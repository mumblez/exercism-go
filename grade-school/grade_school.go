package school

import "sort"

// Grade ...
type Grade struct {
	grade    int
	students []string
}

type gds []Grade

// School ...
type School struct {
	Grades gds
}

// New ...
func New() *School {
	return &School{Grades: []Grade{}}
}

func (s *School) checkExists(stu string, g int) (gradeExists, studentExists bool) {
	for _, gs := range s.Grades {
		if gs.grade == g {
			gradeExists = true
			for _, student := range gs.students {
				if student == stu {
					studentExists = true
					break
				}
			}
		}
	}
	return
}

// Add ...
func (s *School) Add(stu string, g int) {
	gradeExists, studentExists := s.checkExists(stu, g)
	if gradeExists && studentExists {
		return
	}
	if !gradeExists && !studentExists {
		s.Grades = append(s.Grades, Grade{grade: g, students: []string{stu}})
		return
	}
	if gradeExists && !studentExists {
		for k, gs := range s.Grades {
			if gs.grade == g {
				s.Grades[k].students = append(s.Grades[k].students, stu)
				return
			}
		}
	}
}

// Grade ...
func (s *School) Grade(g int) []string {
	gradeExist, _ := s.checkExists("bob", g)
	if !gradeExist {
		return nil
	}
	for _, gs := range s.Grades {
		if gs.grade == g {
			sort.Strings(gs.students)
			return gs.students
		}
	}
	return nil
}

// === sorting functions ===
func (g gds) Len() int {
	return len(g)
}

func (g gds) Swap(i, j int) {
	g[i], g[j] = g[j], g[i]
}

func (g gds) Less(i, j int) bool {
	return g[i].grade < g[j].grade
}

// === end sorting functions ===

// Enrollment ...
func (s *School) Enrollment() []Grade {
	// sort students
	for _, gs := range s.Grades {
		sort.Strings(gs.students)
	}
	// sort grades (custom sort)
	sort.Sort(gds(s.Grades))
	return s.Grades
}
