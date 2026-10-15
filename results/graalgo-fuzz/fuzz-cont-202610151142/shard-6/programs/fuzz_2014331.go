package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }
type GenS2 struct { F0 int8 }
func genFn0() uint32 {
    var v7 int64 = -1
    var v8 uint16 = 255
    var v9 uint16 = v8
    var v10 int64 = v7
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 1
}
func genFn1(p0 int32, p1 uint) int8 {
    v11 := []int{747, 0, 128, 696}
    var v12 int8 = 100
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    return 127
}

func main() {
    var v1 uint8 = 113
    var v2 uint16 = 65535
    v3 := []string{"", "a", "", ((("hello" + (("hello" + "a") + "hello")) + "hello") + "")}
    var v4 int32 = 1
    var v5 uint = 73
    var v6 int = 1
    v1 = v1
    if (!false) {
        var v13 int32 = (v4 + v4)
        var v14 uint = v5
        if (true && false) {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            v15 := GenS1{F0: (v6 % ((v6) | 1)), F1: (^(v5))}
            v16 := GenS2{F0: genFn1((v13 & -66), v5)}
            _ = v15
            _ = v16
        }
        _ = v13
        _ = v14
    } else {
        v17 := GenS2{F0: 100}
        var v18 float64 = 3.14159
        v5 += 65535
        _ = v17
        _ = v18
    }
    v4 -= 2147483647
    for fi19 := 0; fi19 < 3; fi19++ {
        for fi20 := 0; fi20 < 17; fi20++ {
            v2 += 406
            var v21 uint64 = 1
            _ = fi20
            _ = v21
        }
        _ = fi19
    }
    v6 += 127
    if (false || ("" > "a")) {
        if true {
            if (!(1 < 42)) {
                _ = len(v3)
                for fi22 := 0; fi22 < 8; fi22++ {
                    fmt.Println(fi22, v6)
                    _ = fi22
                }
                var v23 uint64 = 18446744073709551615
                _ = v23
            } else {
                v4 -= (v4 & v4)
                for fi24 := 0; fi24 < 9; fi24++ {
                    v2 = v2
                    _ = fi24
                }
                var v25 string = "hello"
                _ = v25
            }
        }
    }
    for fi26 := 0; fi26 < 8; fi26++ {
        var v27 float64 = float64(0.1)
        _ = fi26
        _ = v27
    }
    for fi28 := 0; fi28 < 7; fi28++ {
        fmt.Println(v6, fi28)
        v2 -= v2
        _ = fi28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
