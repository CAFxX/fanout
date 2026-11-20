package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8; F2 uint16 }
type GenS2 struct { F0 int; F1 int8; F2 uint }
type GenS3 struct { F0 int32 }
func genFn0(p0 int16, p1 uint64, p2 uint8) string {
    var v5 int8 = 42
    v6 := GenS1{F0: 255, F1: (v5 / ((v5) | 1)), F2: 765}
    v7 := []int{128, 2147483647, 0, 384}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return ("\n\t" + "foo bar")
}
func genFn1(p0 int64, p1 bool, p2 uint64) uint32 {
    for fi8 := 0; fi8 < 5; fi8++ {
        v9 := GenS1{F0: 800, F1: 100, F2: 186}
        _ = fi8
        _ = v9
    }
    v10 := []int{-802}
    v11 := GenS2{F0: 42, F1: 1, F2: 255}
    if len(v10) >= 2 {
        v10 = v10[1:len(v10)]
    }
    for fi12 := 0; fi12 < 3; fi12++ {
        v13 := GenS2{F0: fi12, F1: 57, F2: 2}
        _ = fi12
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    return 42
}

func main() {
    var v1 int32 = -814
    var v2 uint16 = 1
    var v3 int16 = 1000
    v4 := []string{(("" + "\n\t") + ("foo bar" + "hello")), "", ("" + "\n\t")}
    _ = len(v4)
    v1 += (v1 + (-(v1)))
    var v14 int = -9223372036854775808
    fmt.Println(v14, v1)
    for fi15 := 0; fi15 < 13; fi15++ {
        v14 = (v14 - fi15)
        for fi16 := 0; fi16 < 8; fi16++ {
            var v17 int8 = 0
            _ = fi16
            _ = v17
        }
        v2 = (v2 / ((v2) | 1))
        _ = fi15
    }
    fmt.Println(v3, v2)
    v18 := GenS2{F0: (v14 * v14), F1: 92, F2: 0}
    var v19 uint8 = 128
    if true {
        var v20 uint8 = v19
        v21 := GenS1{F0: v2, F1: int8(v2), F2: 263}
        _ = v20
        _ = v21
    }
    v22 := GenS3{F0: (v1 + v1)}
    var v23 float64 = 1e10
    v23 = ((float64(3.14159) * 1.5) - float64((3.14159 / (-2.25 * float32(v23)))))
    v3 = (v3 + 1000)
    fmt.Println(v19)
    v14 -= v14
    var v24 float64 = 0.0
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}
