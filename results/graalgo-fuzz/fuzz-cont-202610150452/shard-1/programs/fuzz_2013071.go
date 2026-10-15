package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64; F2 uint32 }
func genFn0(p0 int16, p1 uint16) int16 {
    for fi4 := 0; fi4 < 1; fi4++ {
        fmt.Println(p1, p0)
        fmt.Println(p0, fi4, fi4)
        _ = fi4
    }
    for fi5 := 0; fi5 < 3; fi5++ {
        v6 := GenS1{F0: true, F1: 0, F2: 757}
        v7 := GenS1{F0: true, F1: 1, F2: 2147483647}
        _ = fi5
        _ = v6
        _ = v7
    }
    p1 = p1
    var v8 float64 = 1e10
    if false {
        var v9 string = ("\n\t" + ("" + ""))
        fmt.Println(p0, p0, p0)
        _ = v9
    } else {
        var v10 float64 = 3.14159
        _ = v10
    }
    _ = p0
    _ = p1
    _ = v8
    return 210
}

func main() {
    var v1 bool = (255 == 1)
    var v2 int = -976
    var v3 float64 = 0.1
    v11 := GenS1{F0: (!v1), F1: 2, F2: 369}
    var v12 uint = 1
    if v1 {
        v12 += 0
    } else {
        v3 += 3.14159
        if true {
            var v13 uint64 = 467
            if (-983 > 960) {
                v14 := []int{}
                v12 += (v12 - (v12 + (v12 & (v12 & v12))))
                _ = v14
            }
            _ = v13
        }
    }
    v15 := []int{127, (v2 % ((v2) | 1))}
    v16 := GenS1{F0: false, F1: 42, F2: 0}
    fmt.Println(v12, v12, v3)
    fmt.Println(v1, v2)
    _ = len(v15)
    var v17 int16 = 32767
    for fi18 := 0; fi18 < 3; fi18++ {
        v3 -= float64(-2.25)
        _ = fi18
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
}
