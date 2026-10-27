package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint64 }

func main() {
    v1 := make(map[int]int)
    v2 := []int{490, -786}
    v3 := make(map[uint32]bool)
    var v4 int = 2147483647
    var v5 uint = 423
    var v6 bool = (!((false && ((1.5 > 0.0) && (true || (0.0 != 0.1)))) || false))
    v6 = true
    v5 = 255
    for fi7 := 0; fi7 < 9; fi7++ {
        for fi8 := 0; fi8 < 17; fi8++ {
            v9 := GenS1{F0: (v5 << (((^(v5))) & 63))}
            v10 := GenS1{F0: 1}
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = fi7
    }
    fmt.Println(v4, v4)
    var v11 float32 = (-2.25 + 1e10)
    _ = len(v2)
    var v12 string = ""
    v12 = ""
    v4 += v4
    for fi13 := 0; fi13 < 7; fi13++ {
        v12 = (v12 + (v12 + ("hello" + "\n\t")))
        _ = fi13
    }
    for fi14 := 0; fi14 < 11; fi14++ {
        switch -758 {
        case -1:
            for fi15 := 0; fi15 < 7; fi15++ {
                _ = len(v3)
                v12 = ((("foo bar" + "") + (v12 + "\n\t")) + ("" + (v12 + v12)))
                _ = fi15
            }
        case 0:
            for fi16 := 0; fi16 < 7; fi16++ {
                v17 := GenS2{F0: 641}
                _ = fi16
                _ = v17
            }
        }
        v18 := GenS1{F0: v5}
        if len(v2) > 0 {
            v2[0] = (fi14 % (((-(v4))) | 1))
        }
        _ = fi14
        _ = v18
    }
    v19 := v1[v4]
    var v20 uint = 68
    v5 = 885
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
