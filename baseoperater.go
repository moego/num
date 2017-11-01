package number

func MulC(m Matrix,b int)(mr Matrix){
	x,y:=m.Size()
	for i:=0;i<x;i++{
		mr=append(mr,make([]int,y))
		for j:=0;j<y;j++ {
			mr[i][j]=b*m[i][j]
		}
	}
	return
}
func Mul(a,b Matrix)(mr Matrix){

	x,y1:=a.Size()
	x1,y:=b.Size()

	if(y!=x||x1!=y1){
		panic("matrix size error")
	}

	for i:=0;i<x;i++{
		mr=append(mr,make([]int,y))
		for j:=0;j<y;j++ {
			mr[i][j]=0
			for k:=0; k<x; k++{
				if(k<x1&&k<y1){
					s:=a[i][k]*b[k][j];
					mr[i][j]+=s
				}

			}
		}
	}
	return
}

func Sum(ms... Matrix)(m Matrix){
	if(len(ms)==0){
		return
	}else if(len(ms)==1){
		return ms[0]
	}else {
		l:=len(ms)
		x,y:=ms[0].Size()
		for ll:=0;ll<l;ll++{
			x1,y1:=ms[ll].Size()
			if(x1!=x||y1!=y){
				panic("matrix size error")
			}
		}

		for i:=0;i<x;i++{
			m=append(m,make([]int,y))
			for j:=0;j<y;j++ {
				m[i][j]=0
				for ll:=0;ll<l;ll++{
					m[i][j]+= ms[ll][i][j]
				}
			}
		}
	}
	return
}

