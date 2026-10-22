package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 int8 }
func genFn0(p0 int, p1 uint64, p2 string) string {
    var v5 int8 = 127
    switch p2 {
    case "a":
        var v6 int8 = -128
        _ = v6
    case "":
        var v7 int64 = 1
        _ = v7
    case "foo bar":
        p1 += p1
    default:
        if ((!((float64(0.1) - 0.1) == 3.14159)) || true) {
            p1 = (p1 ^ 0)
            if true {
                var v8 int = p0
                v9 := GenS1{F0: (^(v8)), F1: (v5 % ((0) | 1)), F2: v5}
                _ = v8
                _ = v9
            }
        } else {
            p2 = p2
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return ("" + p2)
}
func genFn1() int8 {
    for fi10 := 0; fi10 < 3; fi10++ {
        v11 := []int{}
        if (!((((1.5 - float64(-2.25)) * float64(-2.25)) / ((float64(1e10) - 1.5) + 0.1)) != (float64(0.1) + -2.25))) {
            var v12 uint8 = 0
            v12 -= (^(v12))
            v12 -= 2
            _ = v12
        } else {
            for fi13 := 0; fi13 < 3; fi13++ {
                fmt.Println(fi10, fi13)
                _ = fi13
            }
            for fi14 := 0; fi14 < 13; fi14++ {
                fmt.Println(fi10)
                v15 := GenS1{F0: fi14, F1: 42, F2: 42}
                _ = fi14
                _ = v15
            }
            var v16 float32 = (1e10 + float32((0.0 / 0.1)))
            _ = v16
        }
        _ = fi10
        _ = v11
    }
    if (2147483647 < 2) {
        if true {
            var v17 uint16 = 1
            for fi18 := 0; fi18 < 5; fi18++ {
                var v19 float64 = 1.5
                var v20 uint16 = v17
                v21 := []int{fi18, fi18, 9223372036854775807, (fi18 | (fi18 << ((272) & 63)))}
                _ = fi18
                _ = v19
                _ = v20
                _ = v21
            }
            var v22 uint64 = 42
            _ = v17
            _ = v22
        } else {
            var v23 bool = (0 <= 255)
            v23 = (!false)
            var v24 float32 = (0.1 / 1e10)
            _ = v23
            _ = v24
        }
        var v25 uint64 = 18446744073709551615
        _ = v25
    } else {
        var v26 uint64 = 4294967295
        v27 := []int{32767, 0}
        _ = v26
        _ = v27
    }
    return 67
}

func main() {
    v1 := make(map[string]int)
    v2 := []uint32{}
    v3 := make(map[int]string)
    var v4 uint = 255
    _ = len(v1)
    fmt.Println(v4)
    fmt.Println(v4)
    v28 := GenS1{F0: 877, F1: 1, F2: genFn1()}
    v4 += v4
    var v29 uint8 = 175
    _ = len(v1)
    for fi30 := 0; fi30 < 14; fi30++ {
        v4 -= (^(v4))
        switch "\n\t" {
        case "\n\t":
            v4 = (^(v4))
        case "a":
            v29 += 1
        default:
            if false {
                delete(v1, "")
                v2 = append(v2, 42)
            }
        }
        _ = fi30
    }
    v31 := GenS1{F0: 127, F1: -128, F2: 1}
    if true {
        var v32 uint32 = 2147483648
        _ = v32
    }
    var v33 uint32 = 2147483647
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v31
    _ = v33
    _ = v4
}
