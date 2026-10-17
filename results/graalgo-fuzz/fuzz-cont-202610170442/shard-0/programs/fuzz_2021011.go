package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int; F2 uint8 }
type GenS2 struct { F0 bool }
func genFn0(p0 int64, p1 uint8) int64 {
    if true {
        if (1.5 >= float32(0.0)) {
            fmt.Println(p1)
        }
        v4 := GenS1{F0: 2, F1: 127, F2: 2}
        for fi5 := 0; fi5 < 20; fi5++ {
            switch fi5 {
            case 32768:
                p1 += (p1 >> ((4294967295) & 63))
            case -9223372036854775808:
                var v6 int = 127
                _ = v6
            }
            v7 := []int{-2147483648, (fi5 * (fi5 + -175)), -893}
            _ = fi5
            _ = v7
        }
        _ = v4
    }
    p1 += p1
    if false {
        p0 -= p0
    } else {
        fmt.Println(p1)
    }
    if false {
        fmt.Println(p1, p0)
    } else {
        var v8 int = 9223372036854775807
        _ = v8
    }
    _ = p0
    _ = p1
    return 0
}

func main() {
    var v1 string = ""
    var v2 int32 = 2147483647
    var v3 int16 = -319
    v3 += 0
    v9 := []int{-9223372036854775808, 1, 128}
    v1 = (v1 + ((v1 + (v1 + v1)) + ""))
    v9 = append(v9, 127)
    switch ("foo bar" + v1) {
    case "hello":
        v9 = append(v9, -158)
    case "a":
        var v10 string = v1
        _ = v10
    default:
        if len(v9) > 0 {
            v9[0] = 32767
        }
    }
    var v11 int = 294
    v12 := GenS2{F0: (1 >= 1)}
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v9
}
