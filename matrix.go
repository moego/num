package number

import (
	"strconv"

)

type Matrix [][]int
type Matrixf [][]float64

func RandMatrix(lenx,leny int,duemin,duemax int)(m Matrix){
	for x:=0;x<lenx;x++{
		m=append(m,make([]int,leny))
		for y:=0;y<leny;y++{
			m[x][y]=RandInt(duemin,duemax)
		}
	}
	return
}
func RandMatrixf(lenx,leny int,duemin,duemax float64)(m Matrixf){
	for x:=0;x<lenx;x++{
		m=append(m,make([]float64,leny))
		for y:=0;y<leny;y++{
			m[x][y]=RandFloat(duemin,duemax)
		}
	}
	return
}
func (matrix Matrix)Size()(x,y int)  {
	x=len(matrix)
	for _,line:=range matrix{
		if(len(line)>y){
			y=len(line)
		}
	}
	return
}
func (matrixf Matrixf)Size()(x,y int)  {
	x=len(matrixf)
	for _,line:=range matrixf{
		if(len(line)>y){
			y=len(line)
		}
	}
	return
}
func (matrixf Matrixf)String()(string)  {

	s:="\n"
	for x:=0;x<len( matrixf);x++{
		for y:=0;y<len(matrixf[x])-1;y++{
			s+= strconv.FormatFloat(matrixf[x][y],'f',-1,64)+"\t"
		}
		s+=strconv.FormatFloat(matrixf[x][len(matrixf[x])-1],'f',-1,64)+"\n"
	}
	return s

}

func (matrix Matrix)String()(string)  {

	s:="\n"
	for x:=0;x<len(matrix);x++{
		for y:=0;y<len(matrix[x])-1;y++{
			s+=strconv.Itoa(matrix[x][y])+"\t"
		}
		s+=strconv.Itoa(matrix[x][len(matrix[x])-1])+"\n"
	}
	return s

}

func NewMatrix(datas... []int)(*Matrix,error){
	result:=make(Matrix,len(datas))
	for index,data:=range datas{
		result[index]= data
	}
	return &result,nil
}
func NewMatrixf(datas... []float64)(*Matrixf,error){
	result:=make(Matrixf,len(datas))
	for index,data:=range datas{
		result[index]= data
	}
	return &result,nil
}