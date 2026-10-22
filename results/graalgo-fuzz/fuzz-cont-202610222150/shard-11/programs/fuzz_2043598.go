package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
func genFn0(p0 uint32) int {
    var v5 float32 = (float32(3.14159) + 3.14159)
    v6 := GenS1{F0: -2147483648}
    if (!true) {
        var v7 uint16 = 42
        _ = v7
    }
    v8 := GenS1{F0: 32768}
    _ = p0
    _ = v5
    _ = v6
    _ = v8
    return -2147483648
}

func main() {
    var v1 string = "a"
    var v2 int64 = 42
    var v3 uint32 = 739
    v4 := make(map[uint32]uint32)
    v9 := []int{0, 32768, 32767}
    v4[v3] = v3
    switch 2147483647 {
    case 32768:
        var v10 uint = 238
        _ = v10
    default:
        for fi11 := 0; fi11 < 3; fi11++ {
            for fi12 := 0; fi12 < 8; fi12++ {
                v1 = ("foo bar" + (v1 + v1))
                _ = fi12
            }
            var v13 int = (fi11 % (((fi11 - (fi11 << ((2) & 63)))) | 1))
            _ = fi11
            _ = v13
        }
    }
    if (!false) {
        v14 := GenS1{F0: 449}
        _ = v14
    } else {
        var v15 bool = true
        _ = v15
    }
    v2 -= (v2 / ((1) | 1))
    switch 554 {
    case 32768:
        var v16 int = -2147483648
        _ = v16
    case 575:
        delete(v4, v3)
    }
    v2 += (-(v2))
    for fi17 := 0; fi17 < 17; fi17++ {
        v3 = v3
        var v18 uint = 65535
        _ = fi17
        _ = v18
    }
    for fi19 := 0; fi19 < 14; fi19++ {
        v20 := GenS1{F0: genFn0(v3)}
        _ = len(v9)
        _ = fi19
        _ = v20
    }
    if ((42 >= 42) || true) {
        switch "\n\t" {
        case "foo bar":
            fmt.Println(v2)
        default:
            var v21 int64 = v2
            _ = v21
        }
    }
    v9 = append(v9, 72)
    v2 += (v2 / ((v2) | 1))
    if ((0.0 - 0.0) < 3.14159) {
        switch 128 {
        case 32768:
            for fi22 := 0; fi22 < 3; fi22++ {
                var v23 uint32 = 2147483647
                var v24 uint32 = (v23 / ((2) | 1))
                if true {
                    v24 += (v23 & (v3 % ((v23) | 1)))
                }
                _ = fi22
                _ = v23
                _ = v24
            }
        case 4294967295:
            v1 = ((v1 + (v1 + "")) + v1)
        default:
            var v25 float64 = 1e10
            _ = v25
        }
        for fi26 := 0; fi26 < 2; fi26++ {
            fmt.Println(fi26, v3, v3)
            _ = fi26
        }
        var v27 uint64 = 0
        _ = v27
    }
    _ = len(v4)
    v4[v3] = (v3 & (^(v3)))
    var v28 int8 = -2
    fmt.Println(v1, v28, v2)
    if true {
        for fi29 := 0; fi29 < 10; fi29++ {
            v30 := GenS1{F0: -679}
            var v31 uint16 = 42
            _ = fi29
            _ = v30
            _ = v31
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v9
}
