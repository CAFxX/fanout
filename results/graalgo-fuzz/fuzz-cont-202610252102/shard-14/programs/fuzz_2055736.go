package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64; F2 string }
type GenS2 struct { F0 string; F1 int32; F2 uint16 }
type GenS3 struct { F0 int64; F1 int8 }
func genFn0(p0 bool, p1 int64, p2 bool) bool {
    p0 = false
    v5 := GenS1{F0: -711, F1: 9223372036854775807, F2: "foo bar"}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return ((!p2) || p0)
}

func main() {
    var v1 int = 2147483647
    v2 := make(map[int]bool)
    var v3 int16 = -1
    var v4 int8 = 0
    v3 -= 308
    var v6 string = ""
    for fi7 := 0; fi7 < 7; fi7++ {
        v4 -= (v4 | (-(v4)))
        delete(v2, (-(fi7)))
        _ = fi7
    }
    v3 = v3
    v8 := GenS1{F0: -1, F1: -9223372036854775808, F2: v6}
    var v9 float32 = (3.14159 / 0.1)
    for fi10 := 0; fi10 < 11; fi10++ {
        var v11 int64 = 1
        v12 := GenS2{F0: v6, F1: -523, F2: 2}
        _ = fi10
        _ = v11
        _ = v12
    }
    fmt.Println(v3, v4, v1)
    if false {
        v1 += 42
        for fi13 := 0; fi13 < 5; fi13++ {
            v6 = v6
            var v14 float32 = ((3.14159 / v9) * ((v9 + v9) - ((1.5 / v9) * v9)))
            _ = fi13
            _ = v14
        }
        v15 := GenS3{F0: 909, F1: v4}
        _ = v15
    } else {
        if true {
            for fi16 := 0; fi16 < 18; fi16++ {
                v17 := GenS2{F0: "", F1: -1, F2: 42}
                delete(v2, 32768)
                _ = fi16
                _ = v17
            }
            switch int(v3) {
            case 42:
                v18 := GenS2{F0: "", F1: 1, F2: 42}
                _ = v18
            case 414:
                var v19 float32 = (1e10 - v9)
                _ = v19
            default:
                for fi20 := 0; fi20 < 1; fi20++ {
                    switch 867 {
                    case 42:
                        v9 += (3.14159 / 3.14159)
                    }
                    _ = fi20
                }
            }
        }
        for fi21 := 0; fi21 < 8; fi21++ {
            v4 += (v4 & (v4 + (^(v4))))
            var v22 float64 = ((((1e10 + 0.0) - 3.14159) + 0.0) / 1.5)
            _ = fi21
            _ = v22
        }
    }
    for fi23 := 0; fi23 < 17; fi23++ {
        delete(v2, 727)
        _ = fi23
    }
    v2[4294967295] = true
    v24 := []int{}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v6
    _ = v8
    _ = v9
}
