package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 uint8 }
type GenS2 struct { F0 int8; F1 int32 }
func genFn0(p0 int, p1 uint64) int16 {
    p1 -= p1
    p1 = 2
    var v4 float64 = float64((float32((0.0 + 3.14159)) + (0.1 / 3.14159)))
    if true {
        switch p0 {
        case 32768:
            var v5 int32 = 344
            _ = v5
        case 4294967295:
            fmt.Println(p0)
        case -2147483648:
            for fi6 := 0; fi6 < 6; fi6++ {
                v7 := []int{529, p0}
                _ = fi6
                _ = v7
            }
        default:
            var v8 uint32 = 42
            _ = v8
        }
        v9 := GenS1{F0: 1, F1: 4, F2: 0}
        _ = v9
    } else {
        if false {
            v10 := GenS2{F0: 115, F1: 42}
            if (-2147483648 == -2147483648) {
                v11 := GenS1{F0: 0, F1: 186, F2: 24}
                for fi12 := 0; fi12 < 7; fi12++ {
                    fmt.Println(p1)
                    var v13 string = (("\n\t" + ("hello" + "foo bar")) + ("a" + "a"))
                    _ = fi12
                    _ = v13
                }
                switch 2147483647 {
                case 128:
                    var v14 uint8 = 166
                    _ = v14
                default:
                    v15 := []int{(p0 + (p0 % (((p0 >> ((886) & 63))) | 1))), (p0 ^ 1), p0, p0}
                    _ = v15
                }
                _ = v11
            }
            v4 -= (v4 / v4)
            _ = v10
        } else {
            v4 += (float64((0.0 + (0.1 * (3.14159 - 1e10)))) / v4)
            v16 := []int{}
            v4 -= float64(float32(v4))
            _ = v16
        }
        var v17 bool = false
        _ = v17
    }
    var v18 bool = ((-(p0)) < (p0 / ((p0) | 1)))
    _ = p0
    _ = p1
    _ = v18
    _ = v4
    return 1
}
func genFn1(p0 string) uint16 {
    v19 := []int{}
    var v20 float32 = (1.5 * 1.5)
    var v21 int64 = -1
    _ = p0
    _ = v19
    _ = v20
    _ = v21
    return 42
}

func main() {
    v1 := []uint32{0, 1, 949}
    var v2 int = -1
    v3 := make(map[uint32]uint32)
    var v22 uint8 = 1
    fmt.Println(v2)
    var v23 string = "hello"
    var v24 uint16 = 42
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v25 := GenS2{F0: 125, F1: -1}
    var v26 int64 = 376
    v27 := GenS2{F0: int8(v24), F1: -957}
    var v28 int8 = 1
    if (genFn1(("a" + v23)) != genFn1(v23)) {
        v26 -= (-(v26))
    }
    var v29 int = (v2 << ((621) & 63))
    fmt.Println(v22)
    var v30 uint8 = (^(v22))
    v31 := GenS2{F0: v28, F1: 0}
    v26 = -297
    var v32 int = v29
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
}
