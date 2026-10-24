package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool; F2 uint64 }
func genFn0(p0 int16, p1 uint64, p2 int) uint64 {
    v6 := []int{}
    fmt.Println(p0)
    var v7 int8 = -1
    var v8 uint16 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return p1
}

func main() {
    var v1 uint64 = 2
    var v2 bool = (2147483647 <= 2)
    v3 := []string{"a"}
    var v4 int = 127
    var v5 uint32 = 536
    v4 = 714
    var v9 bool = false
    for fi10 := 0; fi10 < 3; fi10++ {
        v9 = v2
        _ = fi10
    }
    for fi11 := 0; fi11 < 2; fi11++ {
        v12 := GenS1{F0: 2, F1: (!(0 == v1)), F2: v1}
        if false {
            var v13 float64 = (1.5 + 0.1)
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = v13
        }
        _ = fi11
        _ = v12
    }
    v2 = (((v2 || v9) && v9) || (v2 && v2))
    var v14 uint = 691
    var v15 float64 = 3.14159
    var v16 int8 = 125
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
