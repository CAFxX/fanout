package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 bool; F1 uint32 }
type GenS3 struct { F0 int8 }
func genFn0(p0 uint8) uint8 {
    fmt.Println(p0)
    for fi4 := 0; fi4 < 19; fi4++ {
        v5 := GenS1{F0: 77}
        fmt.Println(fi4, fi4)
        _ = fi4
        _ = v5
    }
    var v6 int64 = -670
    _ = p0
    _ = v6
    return 255
}
func genFn1(p0 int32, p1 int64, p2 int) uint16 {
    if false {
        v7 := GenS2{F0: (p1 != p1), F1: 2}
        p0 -= (p0 & 2147483647)
        _ = v7
    }
    var v8 uint16 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 65535
}

func main() {
    var v1 int64 = -1
    v2 := make(map[string]uint32)
    var v3 float64 = 0.1
    var v9 uint8 = 75
    fmt.Println(v1, v3)
    v10 := GenS3{F0: 1}
    var v11 float64 = float64((0.1 - 3.14159))
    var v12 int8 = 127
    var v13 int8 = 1
    var v14 uint = 999
    if false {
        var v15 uint8 = v9
        v3 -= v3
        _ = v15
    }
    v16 := v2[("" + ("hello" + (("" + "a") + ("" + ""))))]
    var v17 float64 = v11
    fmt.Println(v3, v13)
    for fi18 := 0; fi18 < 12; fi18++ {
        switch "hello" {
        case "a":
            v1 = (-(v1))
        case "":
            fmt.Println(v12, v9)
        case "hello":
            v1 -= 0
        default:
            var v19 uint64 = 813
            _ = v19
        }
        _ = fi18
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v9
}
