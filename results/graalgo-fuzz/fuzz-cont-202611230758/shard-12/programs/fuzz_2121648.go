package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32 }
type GenS2 struct { F0 uint64; F1 int64; F2 int32 }
func genFn0(p0 bool, p1 string) uint {
    for fi7 := 0; fi7 < 20; fi7++ {
        var v8 float64 = float64(float32(3.14159))
        v9 := []int{fi7, (fi7 % ((fi7) | 1))}
        _ = fi7
        _ = v8
        _ = v9
    }
    p1 = (("" + "foo bar") + p1)
    _ = p0
    _ = p1
    return 4294967295
}

func main() {
    var v1 uint16 = 65535
    v2 := []int{}
    v3 := []string{(("a" + "hello") + ("" + "")), "\n\t", ("\n\t" + "\n\t"), ("a" + "\n\t")}
    v4 := make(map[string]int)
    v5 := make(map[uint32]int)
    v6 := make(map[int]bool)
    var v10 float64 = 0.1
    v2 = append(v2, 128)
    for fi11 := 0; fi11 < 1; fi11++ {
        var v12 uint8 = 255
        _ = fi11
        _ = v12
    }
    var v13 bool = false
    for fi14 := 0; fi14 < 11; fi14++ {
        v10 = (v10 - 0.1)
        var v15 int16 = 1000
        for fi16 := 0; fi16 < 16; fi16++ {
            v13 = v13
            fmt.Println(v10, v10)
            _ = fi16
        }
        _ = fi14
        _ = v15
    }
    switch -438 {
    case 262:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    case -1:
        v17 := GenS1{F0: 100, F1: 77}
        _ = v17
    case 2147483647:
        var v18 int = 127
        _ = v18
    default:
        var v19 float64 = (0.0 - (v10 * (1e10 - float64(-2.25))))
        _ = v19
    }
    v20 := GenS1{F0: 127, F1: 2147483647}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    _ = len(v5)
    if (!(v1 <= 540)) {
        v13 = v13
    } else {
        var v21 float32 = 0.1
        v21 -= v21
        fmt.Println(v1)
        _ = v21
    }
    v10 = v10
    switch 32768 {
    case 263:
        v22 := GenS1{F0: 95, F1: 0}
        _ = v22
    case -9223372036854775808:
        v10 -= (v10 * v10)
    default:
        v23 := GenS2{F0: 4294967295, F1: -716, F2: 2147483647}
        _ = v23
    }
    _ = len(v6)
    fmt.Println(v1, v10)
    v10 = 1.5
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v13
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
