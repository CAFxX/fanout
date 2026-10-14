package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int; F2 string }
type GenS2 struct { F0 uint; F1 string }
type GenS3 struct { F0 uint16 }
func genFn0() int {
    var v5 uint32 = 65535
    for fi6 := 0; fi6 < 2; fi6++ {
        for fi7 := 0; fi7 < 12; fi7++ {
            v5 -= v5
            v5 = 1
            _ = fi7
        }
        var v8 int8 = 100
        _ = fi6
        _ = v8
    }
    _ = v5
    return 950
}
func genFn1(p0 int32, p1 uint) uint8 {
    var v9 uint64 = 18446744073709551615
    fmt.Println(v9, p1, v9)
    _ = p0
    _ = p1
    _ = v9
    return 2
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[int]int)
    v3 := []int{}
    var v4 int8 = 0
    var v10 uint = 1
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi11 := 0; fi11 < 1; fi11++ {
        if ((v4 / ((90) | 1)) > v4) {
            var v12 uint = (v10 | 42)
            v10 = (v10 & v10)
            for fi13 := 0; fi13 < 17; fi13++ {
                delete(v1, (fi11 | (fi13 ^ fi11)))
                v14 := GenS1{F0: -1, F1: 12, F2: "\n\t"}
                _ = fi13
                _ = v14
            }
            _ = v12
        }
        _ = fi11
    }
    var v15 uint8 = 160
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v15 -= genFn1(2147483647, v10)
    v16 := GenS1{F0: 42, F1: 0, F2: ""}
    var v17 string = "\n\t"
    v10 = (v10 * 958)
    var v18 uint32 = 1
    if false {
        if (!true) {
            if false {
                if ((1.5 >= (0.0 / 1.0)) || false) {
                    v19 := GenS1{F0: -32768, F1: 127, F2: v17}
                    _ = v19
                }
            }
            v15 += genFn1(-903, 567)
            if true {
                v20 := GenS2{F0: (v10 - 18446744073709551615), F1: ("" + v17)}
                v17 = "hello"
                _ = v20
            } else {
                v15 -= 128
                if len(v3) > 0 {
                    v3[0] = 32768
                }
            }
        }
        v21 := GenS2{F0: (v10 / ((1) | 1)), F1: v17}
        if ((("" + v17) + "foo bar") != "") {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
        _ = v21
    } else {
        v10 = (v10 * v10)
        fmt.Println(v17, v15, v10)
        switch -775 {
        case 2147483647:
            v22 := GenS3{F0: 42}
            _ = v22
        default:
            var v23 bool = (!true)
            _ = v23
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
