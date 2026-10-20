package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64; F2 uint8 }
func genFn0(p0 string, p1 uint8) int32 {
    switch 9223372036854775807 {
    case -1:
        for fi5 := 0; fi5 < 8; fi5++ {
            var v6 float64 = 0.0
            var v7 string = p0
            _ = fi5
            _ = v6
            _ = v7
        }
    }
    v8 := []int{9223372036854775807, 58, 2147483647, 2147483647}
    p0 = ("" + "\n\t")
    _ = p0
    _ = p1
    _ = v8
    return -1
}

func main() {
    var v1 uint32 = 0
    v2 := make(map[int]bool)
    v3 := []string{"a"}
    var v4 uint = 4294967295
    switch 32767 {
    case 32768:
        fmt.Println(v4, v1)
    case 32767:
        if false {
            fmt.Println(v1)
            var v9 bool = false
            v9 = v9
            _ = v9
        }
    default:
        if len(v3) > 0 {
            v3[0] = ""
        }
    }
    for fi10 := 0; fi10 < 5; fi10++ {
        var v11 uint16 = 255
        for fi12 := 0; fi12 < 12; fi12++ {
            fmt.Println(fi12)
            _ = fi12
        }
        _ = fi10
        _ = v11
    }
    v13 := GenS1{F0: (32767 > -32768), F1: 18446744073709551615, F2: 0}
    fmt.Println(v1, v1)
    if false {
        switch 9223372036854775807 {
        case 2147483647:
            v4 += 2
        default:
            v1 -= (^(v1))
        }
    } else {
        var v14 int8 = -1
        _ = v14
    }
    var v15 int16 = 1
    v15 = int16(v4)
    for fi16 := 0; fi16 < 15; fi16++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi16
    }
    var v17 string = "a"
    v17 = v17
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v15
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
