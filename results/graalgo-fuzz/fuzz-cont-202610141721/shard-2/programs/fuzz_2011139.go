package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint16; F1 string; F2 int8 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint16 }
func genFn0(p0 int16, p1 uint) int {
    p0 += -32768
    if true {
        p1 = (p1 | p1)
    }
    var v6 int32 = 1
    v7 := []int{}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 962
}

func main() {
    var v1 float64 = (3.14159 * 0.0)
    v2 := make(map[int]int)
    v3 := make(map[uint32]uint32)
    var v4 bool = (0 != 2)
    var v5 int = 1
    if (32767 <= -32768) {
        v8 := GenS1{F0: 106}
        for fi9 := 0; fi9 < 15; fi9++ {
            fmt.Println(v4, v5, v5)
            _ = fi9
        }
        var v10 uint32 = 42
        _ = v10
        _ = v8
    }
    delete(v3, 2147483648)
    if v4 {
        v2[(v5 >> ((980) & 63))] = v5
    }
    var v11 string = "hello"
    var v12 int32 = -452
    var v13 int64 = 0
    for fi14 := 0; fi14 < 18; fi14++ {
        v15 := []int{(v5 ^ (-(fi14))), 0, fi14}
        delete(v3, 504)
        _ = fi14
        _ = v15
    }
    v13 = 42
    fmt.Println(v4, v4, v5)
    v16 := GenS2{F0: 42, F1: v11, F2: 0}
    var v17 uint64 = 1
    if (((!false) && v4) || false) {
        v18 := GenS3{F0: (32767 != 0)}
        _ = v18
    }
    for fi19 := 0; fi19 < 3; fi19++ {
        var v20 uint = 153
        _ = fi19
        _ = v20
    }
    var v21 int8 = -78
    v13 += -9223372036854775808
    if (true && v4) {
        v22 := GenS4{F0: 65535}
        _ = v22
    } else {
        v4 = true
        v11 = ""
        var v23 int32 = (v12 / ((v12) | 1))
        _ = v23
    }
    var v24 uint8 = 127
    fmt.Println(v5, v4, v5)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
