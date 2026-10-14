package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int; F1 string; F2 uint64 }
type GenS3 struct { F0 int16; F1 int; F2 bool }
func genFn0(p0 uint16, p1 bool, p2 int) int8 {
    fmt.Println(p0, p1)
    p0 = 32767
    p0 += p0
    var v4 float32 = (-2.25 + (1.5 + -2.25))
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return -1
}
func genFn1() int {
    var v5 string = "a"
    var v6 bool = false
    v6 = (42 >= 127)
    _ = v5
    _ = v6
    return -2147483648
}

func main() {
    var v1 uint64 = 2
    var v2 uint64 = v1
    v3 := make(map[uint32]uint32)
    var v7 string = ""
    if (!true) {
        for fi8 := 0; fi8 < 18; fi8++ {
            var v9 float64 = -2.25
            v1 += 18446744073709551615
            v7 = (("\n\t" + v7) + v7)
            _ = fi8
            _ = v9
        }
    } else {
        switch v7 {
        case "a":
            v2 = 746
        case "foo bar":
            fmt.Println(v1)
        default:
            v10 := GenS1{F0: 255}
            _ = v10
        }
        if true {
            var v11 int = 2147483647
            switch 359 {
            case 127:
                v11 -= (v11 << ((255) & 63))
            case 1:
                v12 := v3[125]
                _ = v12
            default:
                _ = len(v3)
            }
            delete(v3, 2147483648)
            _ = v11
        }
    }
    var v13 uint16 = 1
    v1 = v1
    for fi14 := 0; fi14 < 17; fi14++ {
        v15 := GenS2{F0: 0, F1: "\n\t", F2: 1}
        _ = fi14
        _ = v15
    }
    _ = len(v3)
    fmt.Println(v7)
    v13 = (v13 & v13)
    var v16 int8 = int8(v13)
    if false {
        v17 := GenS3{F0: -32768, F1: -9223372036854775808, F2: true}
        fmt.Println(v16, v2, v16)
        if true {
            var v18 uint64 = (v2 % ((2) | 1))
            _ = v18
        } else {
            var v19 int16 = 0
            var v20 uint16 = v13
            _ = len(v3)
            _ = v19
            _ = v20
        }
        _ = v17
    }
    var v21 uint16 = 65535
    var v22 string = "hello"
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v16
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v7
}
