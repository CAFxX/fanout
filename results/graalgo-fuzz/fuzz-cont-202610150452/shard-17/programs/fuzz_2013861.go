package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16 }
type GenS2 struct { F0 int64 }
func genFn0(p0 uint16) int64 {
    v6 := GenS1{F0: -564, F1: 42}
    v7 := []int{32768, -408, 127}
    switch 317 {
    case 9223372036854775807:
        p0 += (p0 + p0)
    case -635:
        var v8 bool = (584 == 1)
        _ = v8
    case -1:
        _ = len(v7)
    default:
        switch "a" {
        case "\n\t":
            p0 = (p0 % ((1) | 1))
        default:
            v9 := GenS1{F0: 481, F1: 1}
            _ = v9
        }
    }
    fmt.Println(p0)
    var v10 int16 = -299
    _ = p0
    _ = v10
    _ = v6
    _ = v7
    return -1
}

func main() {
    var v1 bool = (true || false)
    var v2 uint16 = 65535
    v3 := []uint32{4294967295}
    v4 := make(map[int]uint32)
    var v5 int32 = -2147483648
    var v11 bool = ((v2 % ((37) | 1)) > v2)
    var v12 uint16 = 331
    var v13 string = ("hello" + "\n\t")
    v14 := GenS2{F0: -9223372036854775808}
    if (4294967295 == 243) {
        var v15 float64 = 1.5
        _ = v15
    } else {
        _ = len(v4)
    }
    if v11 {
        fmt.Println(v12)
        switch (((v13 + "") + "a") + v13) {
        case "":
            if len(v3) > 0 {
                v3[0] = 65535
            }
        case "a":
            v11 = true
        case "\n\t":
            fmt.Println(v5)
        default:
            _ = len(v4)
        }
    }
    fmt.Println(v13, v1)
    v12 = v12
    v13 = "foo bar"
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
