package main
import (
       "strconv"
       "strings"
       "bufio"
       "fmt"
       "os"
)

func Swap( sli []int, index int ){
     temp := sli[ index ];
     sli[ index ] = sli[ index + 1 ];
     sli[ index + 1 ] = temp;
     }

func BubbleSort( sli []int ) {
     sliLen := len( sli )
     dontBreak := false;
     for i:=0; i< sliLen; i++ {
     	 for j:=0; j< sliLen-i-1; j++ {
	     if sli[ j ] > sli[ j+1 ] {
	     	Swap( sli, j );
		dontBreak = true;
	     	}
	     }
	     if !dontBreak {
	     	break;
		}
	}
     for i:=0; i<len( sli ); i++ {
     	 fmt.Printf( "%d ",sli[ i ] )
	 }
	 fmt.Println();
     }	 
func main() {
     scanner := bufio.NewScanner( os.Stdin );
     fmt.Printf( "Enter integers to sort( max 10 int ):" )
     scanner.Scan();
     intStr := scanner.Text();
     arTemp := strings.Split(intStr, " ");
     arTempLen := len( arTemp );
     arTempInt := make([]int, arTempLen);
     for itr:=0; itr<arTempLen; itr++ {
     	 arTempInt[itr],_ = strconv.Atoi(arTemp[itr]);
     }
     BubbleSort( arTempInt );
}

     	 