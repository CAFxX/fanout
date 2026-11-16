package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64 }
func genFn0(p0 int8, p1 uint16, p2 string) int16 {
    fmt.Println(p2, p0)
    var v5 uint16 = 310
    v6 := []int{-2147483648, 32768, -1, 9223372036854775807}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return -720
}

func main() {
    var v1 float32 = -2.25
    var v2 uint8 = 255
    var v3 string = "a"
    v4 := []int{-9223372036854775808}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    for fi7 := 0; fi7 < 4; fi7++ {
        v3 = "hello"
        _ = fi7
    }
    var v8 uint64 = 745
    var v9 int8 = -37
    switch int(v9) {
    case 0:
        var v10 int8 = (v9 >> ((431) & 63))
        _ = v10
    case 32767:
        var v11 string = "hello"
        _ = v11
    case -193:
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    default:
        for fi12 := 0; fi12 < 18; fi12++ {
            v1 += 1e10
            if (-2.25 <= (3.14159 / (1e10 + 0.1))) {
                v1 = v1
                var v13 float32 = (v1 - v1)
                v3 = (v3 + "a")
                _ = v13
            }
            v8 += 1
            _ = fi12
        }
    }
    var v14 int = 32767
    v15 := GenS1{F0: -1, F1: v8}
    for fi16 := 0; fi16 < 13; fi16++ {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi16
    }
    _ = len(v4)
    v8 = (v8 / (((v8 >> ((450) & 63))) | 1))
    var v17 int = v14
    for fi18 := 0; fi18 < 8; fi18++ {
        fmt.Println(v1)
        v19 := GenS1{F0: int32(v2), F1: v8}
        _ = fi18
        _ = v19
    }
    v3 = ((v3 + v3) + (v3 + "hello"))
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v15
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
