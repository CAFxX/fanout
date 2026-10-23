package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
func genFn0(p0 uint8, p1 int16) int {
    fmt.Println(p1, p0)
    var v6 int16 = (-(p1))
    switch "a" {
    case "hello":
        v7 := GenS1{F0: ("a" + "")}
        _ = v7
    default:
        fmt.Println(p0, v6, v6)
    }
    v8 := []int{0, 277}
    _ = p0
    _ = p1
    _ = v6
    _ = v8
    return -781
}
func genFn1(p0 uint) uint32 {
    v9 := []int{genFn0(159, 1), 128, 4294967295, -2147483648}
    p0 += p0
    var v10 int64 = -673
    fmt.Println(p0, v10)
    _ = p0
    _ = v10
    _ = v9
    return 972
}

func main() {
    v1 := []int{32767}
    v2 := make(map[string]int)
    var v3 int16 = 1
    v4 := make(map[string]bool)
    v5 := make(map[int]int)
    v11 := GenS1{F0: "hello"}
    for fi12 := 0; fi12 < 16; fi12++ {
        var v13 int16 = int16(fi12)
        var v14 int64 = 325
        _ = fi12
        _ = v13
        _ = v14
    }
    if true {
        fmt.Println(v3)
    }
    var v15 int64 = 1
    var v16 int8 = 127
    if (65535 <= 1) {
        v3 -= v3
    }
    fmt.Println(v16, v3)
    var v17 uint8 = 28
    var v18 float32 = (0.0 * 0.1)
    fmt.Println(v18, v16, v18)
    for fi19 := 0; fi19 < 1; fi19++ {
        v15 = (v15 >> ((481) & 63))
        for fi20 := 0; fi20 < 8; fi20++ {
            v17 = 127
            v17 += 128
            _ = fi20
        }
        _ = fi19
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
