package Controllers

import (
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zxc10110/mvc_09/Models"
)

func CountDigits(number int) int {
	if number < 10 {
		return 1
	} else {
		return 1 + CountDigits(number/10)
	}
}

func firsttheeDigits(number int) int {
	firsttheeDigits := number / 1000000
	return firsttheeDigits
}

func CreateStudent(c *gin.Context) {
	var student Models.Student
	c.BindJSON(&student)
	tDigit := firsttheeDigits(student.StudentId)

	//studentid < 8
	if CountDigits(student.StudentId) < 8 || CountDigits(student.StudentId) > 8 {
		c.JSON(http.StatusBadRequest, "รหัสนักศึกษาต้องมี 8 หลัก")
		return
	} else {
		//check first three digits
		if tDigit < 60 || tDigit > 63 {
			c.JSON(http.StatusBadRequest, "รหัสนักศึกษาต้องนำหน้าด้วย 60-63 เท่านั้น")
			return
		} else {

			//check year == studentId
			if tDigit == 60 && student.Year != 4 {
				c.JSON(http.StatusBadRequest, "รหัสนักศึกษาและช้ันปีต้องสอดคล้องกัน")
				return
			} else if tDigit == 61 && student.Year != 3 {
				c.JSON(http.StatusBadRequest, "รหัสนักศึกษาและช้ันปีต้องสอดคล้องกัน")
				return
			} else if tDigit == 62 && student.Year != 2 {
				c.JSON(http.StatusBadRequest, "รหัสนักศึกษาและช้ันปีต้องสอดคล้องกัน")
				return
			} else if tDigit == 63 && student.Year != 1 {
				c.JSON(http.StatusBadRequest, "รหัสนักศึกษาและช้ันปีต้องสอดคล้องกัน")
				return
			} else {
				//create
				err := Models.CreateStudent(&student)
				if err != nil {
					c.JSON(http.StatusNotFound, "ไม่พบข้อมูล")
					return
				}
				c.JSON(http.StatusOK, student)
				return
			}
		}

	}

}

func GetStudent(c *gin.Context) {
	var student []Models.Student
	err := Models.GetAllStudent(&student)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, student)
}

func GetSport(c *gin.Context) {
	var sport []Models.Sport
	err := Models.GetSport(&sport)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, sport)
}

func SportStart(c *gin.Context) {
	//count student
	count, e := Models.CountStudent()
	if e != nil {
		return
	}

	//random
	var sport Models.Student
	sport, er := Models.RandomSport()
	if er != nil {
		return
	}

	sportNumber := rand.Intn(count.Count / 2)
	result := Models.Sport{
		SportNo: sportNumber,
		Student: sport,
	}

	//count sport > 2
	countS, err := Models.CountSport()
	if err != nil {
		return
	}
	if countS.Count > 2 {
		c.JSON(http.StatusBadRequest, "ครบคู่แล้ว")
		return
	} else {
		//getSport
		//getS, errrr := Models.GetSport()
		errr := Models.CreateSport(&result)
		if errr != nil {
			return
		}

		c.JSON(http.StatusOK, result)
		return
	}
}
