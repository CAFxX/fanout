package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int8; F1 uint16; F2 uint }

func main() {
    var v1 int16 = 42
    var v2 float32 = (1.5 * (3.14159 - -2.25))
    v3 := make(map[string]int)
    var v4 bool = (!true)
    v5 := make(map[string]uint32)
    fmt.Println(v4, v2)
    var v6 bool = (!((4294967295 <= 2) && (2147483648 <= 43)))
    for fi7 := 0; fi7 < 13; fi7++ {
        var v8 int32 = 0
        var v9 uint = 1
        var v10 uint = v9
        _ = fi7
        _ = v10
        _ = v8
        _ = v9
    }
    v11 := []int{32767, -939, 32767}
    if (v4 && (187 >= 0)) {
        v12 := GenS1{F0: 4294967295}
        for fi13 := 0; fi13 < 8; fi13++ {
            v2 -= (v2 / (v2 - ((((-2.25 + 3.14159) - v2) / v2) * float32(-2.25))))
            fmt.Println(v6, v4, v1)
            _ = fi13
        }
        _ = v12
    } else {
        fmt.Println(v1)
    }
    if (1 < 100) {
        v6 = (255 == 65535)
        v1 += (v1 << ((18446744073709551615) & 63))
    } else {
        var v14 int16 = -789
        _ = v14
    }
    switch "a" {
    case "a":
        var v15 float32 = v2
        _ = v15
    default:
        var v16 int64 = 0
        _ = v16
    }
    v6 = ((!(588 >= 42)) || false)
    if true {
        fmt.Println(v1, v1, v1)
    } else {
        for fi17 := 0; fi17 < 7; fi17++ {
            fmt.Println(v1)
            _ = fi17
        }
        v18 := GenS2{F0: 42, F1: 0, F2: 269}
        _ = v18
    }
    fmt.Println(v1, v1)
    var v19 int64 = 280
    if (v1 == (v1 | (v1 + v1))) {
        fmt.Println(v6)
    }
    for fi20 := 0; fi20 < 17; fi20++ {
        v11 = append(v11, 497)
        _ = fi20
    }
    v21 := GenS1{F0: 127}
    fmt.Println(v19, v6)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v19
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
