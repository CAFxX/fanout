package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool; F2 uint }
type GenS2 struct { F0 uint32; F1 string; F2 uint }
type GenS3 struct { F0 uint16; F1 bool; F2 bool }
func genFn0(p0 uint32) int16 {
    v4 := []int{871, 646}
    p0 -= 728
    _ = p0
    _ = v4
    return -32768
}

func main() {
    var v1 int16 = 642
    var v2 string = "a"
    var v3 int32 = -1
    v5 := []int{32768, 32768}
    if true {
        var v6 int32 = v3
        v6 += 1
        _ = v6
    }
    v7 := GenS1{F0: true, F1: (((1 <= 338) || false) || (1e10 != 1e10)), F2: 0}
    fmt.Println(v3, v3)
    var v8 uint16 = 42
    v8 += v8
    var v9 int16 = v1
    if true {
        var v10 bool = true
        var v11 int = 42
        var v12 uint8 = 112
        _ = v10
        _ = v11
        _ = v12
    }
    if false {
        v13 := GenS2{F0: 1, F1: (v2 + v2), F2: uint(v8)}
        v3 = v3
        _ = v13
    }
    _ = len(v5)
    var v14 uint64 = 225
    switch 9223372036854775807 {
    case 1:
        v5 = append(v5, 0)
    case -1:
        var v15 float32 = 3.14159
        _ = v15
    case 956:
        if ((-397 == 42) && true) {
            if (v2 != v2) {
                var v16 int64 = -9223372036854775808
                for fi17 := 0; fi17 < 1; fi17++ {
                    if true {
                        fmt.Println(v2, v3, v14)
                    }
                    _ = fi17
                }
                _ = v16
            }
            v5 = append(v5, 2147483647)
        } else {
            v18 := GenS2{F0: 1, F1: ("hello" + v2), F2: 639}
            _ = v18
        }
    default:
        var v19 int16 = v1
        _ = v19
    }
    if (1 <= (v3 / ((v3) | 1))) {
        switch 1 {
        case 4294967295:
            var v20 bool = true
            _ = v20
        }
    } else {
        v1 = v1
    }
    v9 = v9
    v14 = v14
    switch -476 {
    case -472:
        var v21 int = -983
        _ = v21
    case -235:
        for fi22 := 0; fi22 < 16; fi22++ {
            v23 := GenS3{F0: v8, F1: true, F2: true}
            _ = fi22
            _ = v23
        }
    default:
        v9 = (v1 - -1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
