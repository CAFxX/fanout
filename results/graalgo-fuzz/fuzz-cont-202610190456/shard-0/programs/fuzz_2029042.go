package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int }
func genFn0(p0 uint8, p1 int32, p2 int32) bool {
    var v4 int16 = 1000
    if false {
        p1 = p1
        var v5 int8 = -125
        var v6 int8 = int8(p1)
        _ = v5
        _ = v6
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return true
}
func genFn1(p0 uint16, p1 bool) int {
    for fi7 := 0; fi7 < 1; fi7++ {
        var v8 uint16 = (^(p0))
        _ = fi7
        _ = v8
    }
    var v9 uint32 = 375
    for fi10 := 0; fi10 < 10; fi10++ {
        p0 = (p0 ^ 255)
        v11 := GenS1{F0: 63, F1: 710}
        _ = fi10
        _ = v11
    }
    v12 := GenS1{F0: -2147483648, F1: 332}
    v13 := []int{}
    _ = p0
    _ = p1
    _ = v12
    _ = v13
    _ = v9
    return 2147483647
}

func main() {
    var v1 float32 = 1.5
    var v2 float32 = v1
    v3 := []string{"", ""}
    fmt.Println(v2, v2)
    if true {
        if false {
            if (false && false) {
                var v14 string = ("hello" + ("hello" + "\n\t"))
                v14 = v14
                v14 = (v14 + (v14 + (v14 + v14)))
                _ = v14
            }
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        var v15 bool = false
        _ = v15
    }
    if ((32767 == 32767) || ((1 != 42) && (543 > 1))) {
        fmt.Println(v1, v2)
        fmt.Println(v1, v1)
        v2 = 0.1
    } else {
        var v16 uint16 = 558
        _ = v16
    }
    if len(v3) > 0 {
        v3[0] = "\n\t"
    }
    fmt.Println(v2, v2)
    var v17 uint64 = 18446744073709551615
    if true {
        v18 := GenS1{F0: -1, F1: 513}
        v1 = 1e10
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = v18
    }
    v1 -= (((v2 + float32((0.0 + -2.25))) / v2) * (((0.1 * 1e10) / v2) + ((((v2 + (v2 + (v1 - 1e10))) * v2) / v1) * (v1 + v2))))
    var v19 uint64 = 206
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v17
    _ = v19
    _ = v2
    _ = v3
}
