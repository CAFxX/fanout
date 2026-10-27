package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 uint32 }
func genFn0(p0 uint32) int64 {
    var v7 uint16 = 32767
    v7 += (^(v7))
    var v8 bool = (0 == -3)
    fmt.Println(v7, p0, v8)
    v9 := []int{530}
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return -453
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[int]bool)
    var v3 uint32 = 2147483647
    var v4 uint = 18446744073709551615
    v5 := make(map[string]int)
    var v6 uint16 = 2
    var v10 int64 = 0
    if false {
        if (-32768 != 1) {
            var v11 int64 = v10
            _ = v11
        }
    } else {
        v12 := GenS1{F0: -689, F1: 108, F2: (v3 ^ v3)}
        v5[("a" + (("a" + "\n\t") + ""))] = -1
        var v13 uint16 = (v6 << ((255) & 63))
        _ = v12
        _ = v13
    }
    v10 -= 1
    var v14 int16 = -32768
    var v15 uint32 = (^(v3))
    switch 440 {
    case -9223372036854775808:
        var v16 uint32 = v3
        _ = v16
    }
    v15 = v3
    switch "a" {
    case "":
        v14 = v14
    default:
        v17 := GenS1{F0: (v14 & 42), F1: 127, F2: v15}
        _ = v17
    }
    if true {
        for fi18 := 0; fi18 < 18; fi18++ {
            v4 += 2
            v3 = uint32(v6)
            _ = fi18
        }
    } else {
        if false {
            v4 = v4
        } else {
            v4 = (v4 / ((441) | 1))
        }
        fmt.Println(v15)
        if (!false) {
            for fi19 := 0; fi19 < 16; fi19++ {
                var v20 uint32 = 1
                var v21 float64 = (3.14159 + 3.14159)
                _ = fi19
                _ = v20
                _ = v21
            }
        } else {
            switch "a" {
            case "a":
                v6 = v6
            case "":
                var v22 float64 = 1.5
                _ = v22
            }
            v23 := GenS1{F0: -1, F1: 42, F2: 1}
            _ = v23
        }
    }
    var v24 uint64 = 18446744073709551615
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
