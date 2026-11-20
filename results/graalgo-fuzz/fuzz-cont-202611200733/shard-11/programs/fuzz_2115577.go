package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }
type GenS2 struct { F0 uint16; F1 int64 }
type GenS3 struct { F0 uint64; F1 string }
func genFn0() uint64 {
    v7 := []int{4294967295, 0, 32767, 332}
    if (-394 > 940) {
        v7 = append(v7, -142)
        v7 = append(v7, 0)
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
    } else {
        switch -578 {
        case 569:
            for fi8 := 0; fi8 < 12; fi8++ {
                var v9 uint32 = 920
                _ = fi8
                _ = v9
            }
        default:
            v7 = append(v7, 1)
        }
        var v10 uint64 = 1
        for fi11 := 0; fi11 < 17; fi11++ {
            if len(v7) > 0 {
                v7[0] = (fi11 - fi11)
            }
            v10 -= 1
            _ = fi11
        }
        _ = v10
    }
    switch "" {
    case "":
        var v12 int = 32768
        _ = v12
    case "hello":
        v13 := GenS1{F0: 0, F1: 643}
        _ = v13
    case "a":
        for fi14 := 0; fi14 < 16; fi14++ {
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
            _ = fi14
        }
    default:
        v15 := GenS2{F0: 1, F1: 281}
        _ = v15
    }
    var v16 bool = (0 < -32)
    v16 = (!false)
    _ = v16
    _ = v7
    return 539
}
func genFn1(p0 uint, p1 uint16) uint64 {
    var v17 int8 = 127
    var v18 float64 = 1.5
    _ = p0
    _ = p1
    _ = v17
    _ = v18
    return 42
}

func main() {
    v1 := make(map[string]string)
    var v2 uint32 = 42
    v3 := []string{}
    v4 := []uint32{65535, (v2 % (((^(v2))) | 1)), v2}
    var v5 int32 = -1
    var v6 float32 = float32((1e10 + (float64(1e10) + 0.1)))
    var v19 uint = 2
    for fi20 := 0; fi20 < 18; fi20++ {
        v21 := GenS1{F0: 127, F1: 1}
        v22 := GenS1{F0: (^(fi20)), F1: 2}
        v4 = append(v4, (^(v2)))
        _ = fi20
        _ = v21
        _ = v22
    }
    var v23 float32 = (v6 + ((0.0 * 1e10) - 0.0))
    var v24 uint64 = 42
    var v25 float32 = v23
    v1[("a" + (("hello" + ("a" + "")) + ("\n\t" + "hello")))] = ("a" + "")
    v24 += (v24 - uint64(v2))
    if true {
        var v26 int64 = 0
        _ = v26
    }
    if true {
        var v27 int = -347
        if (false && false) {
            delete(v1, "\n\t")
            var v28 int32 = (^(v5))
            _ = v28
        } else {
            var v29 bool = ((461 >= -1) && true)
            fmt.Println(v6, v27)
            _ = v29
        }
        var v30 int64 = 618
        _ = v27
        _ = v30
    }
    v31 := GenS3{F0: (v24 % ((genFn0()) | 1)), F1: "a"}
    var v32 uint64 = (v24 ^ 0)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
