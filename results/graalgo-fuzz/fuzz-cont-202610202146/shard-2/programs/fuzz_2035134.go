package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint; F1 uint64 }
type GenS3 struct { F0 int; F1 uint16 }
type GenS4 struct { F0 uint64; F1 uint32 }
type GenS5 struct { F0 int16; F1 uint16; F2 int32 }
func genFn0() uint {
    v7 := GenS1{F0: 2147483647}
    var v8 uint = 1
    _ = v7
    _ = v8
    return (v8 | v8)
}
func genFn1(p0 uint16) uint {
    p0 = 42
    for fi9 := 0; fi9 < 8; fi9++ {
        v10 := GenS1{F0: -1}
        v11 := GenS1{F0: 1}
        _ = fi9
        _ = v10
        _ = v11
    }
    v12 := GenS2{F0: 2, F1: 0}
    _ = p0
    _ = v12
    return 129
}

func main() {
    var v1 uint64 = 320
    var v2 uint64 = (v1 + v1)
    v3 := []int{128, 42}
    var v4 int64 = -350
    v5 := []string{"", ("\n\t" + (("a" + (("" + ("foo bar" + "a")) + "a")) + (("hello" + "\n\t") + "a"))), (("a" + ("hello" + "foo bar")) + "a")}
    v6 := make(map[int]bool)
    switch 0 {
    case 127:
        var v13 int16 = 42
        _ = v13
    case -111:
        var v14 float32 = (-2.25 * (float32((-2.25 - float64(1.5))) - 1.5))
        _ = v14
    default:
        v4 += v4
    }
    fmt.Println(v4, v1)
    v1 += 42
    fmt.Println(v4, v2)
    v15 := GenS1{F0: -2147483648}
    v16 := GenS3{F0: 727, F1: uint16(v1)}
    v4 = v4
    v17 := GenS4{F0: v1, F1: 0}
    var v18 uint16 = 230
    v19 := GenS5{F0: 404, F1: (v18 - 295), F2: -1}
    v20 := GenS1{F0: 42}
    var v21 uint = 18446744073709551615
    fmt.Println(v1, v18, v2)
    for fi22 := 0; fi22 < 6; fi22++ {
        if ((v18 + v18) != 0) {
            v6[fi22] = (false || ((fi22 + fi22) >= fi22))
        } else {
            for fi23 := 0; fi23 < 10; fi23++ {
                var v24 uint32 = 42
                _ = fi23
                _ = v24
            }
            var v25 int = fi22
            v26 := GenS3{F0: int(v21), F1: 32767}
            _ = v25
            _ = v26
        }
        _ = fi22
    }
    v4 -= (v4 % ((948) | 1))
    for fi27 := 0; fi27 < 18; fi27++ {
        for fi28 := 0; fi28 < 13; fi28++ {
            var v29 int64 = (v4 * v4)
            var v30 uint32 = 2147483647
            _ = fi28
            _ = v29
            _ = v30
        }
        var v31 int8 = 0
        var v32 uint32 = 2147483647
        _ = fi27
        _ = v31
        _ = v32
    }
    var v33 int32 = -2147483648
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
