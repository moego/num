package number



func Sum(m1,m2 Matrix)(m Matrix){
	x1,y1:=m1.Size()
	x,y:=m2.Size()
	if(x1==x||y1==y){
		for i:=0;i<len(m1);i++{
			m=append(m,make([]int,y))
			for j:=0;j<y;j++ {
				m[i][j]=m1[i][j]+m2[i][j]
			}
		}
	}else {
		panic("matrix size error")
	}
	return
}
