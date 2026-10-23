package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8; F2 string }
func genFn0(p0 uint64, p1 int8, p2 int64) uint16 {
    switch 32768 {
    case 959:
        for fi4 := 0; fi4 < 13; fi4++ {
            fmt.Println(p0, p0, p1)
            p0 = p0
            _ = fi4
        }
    default:
        fmt.Println(p0)
    }
    p0 -= 725
    var v5 uint8 = 27
    for fi6 := 0; fi6 < 18; fi6++ {
        var v7 bool = true
        _ = fi6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 32767
}

func main() {
    var v1 int8 = 36
    var v2 uint32 = 4294967295
    var v3 bool = (0.0 < (0.1 * -2.25))
    fmt.Println(v2, v2)
    v3 = true
    v2 = (v2 << ((841) & 63))
    v8 := []int{-466, 997, -580}
    var v9 int = 1
    if (2 == 18446744073709551615) {
        v10 := GenS1{F0: true, F1: 127, F2: ((("hello" + "a") + "a") + "a")}
        fmt.Println(v3)
        fmt.Println(v3, v9)
        _ = v10
    } else {
        var v11 int = (v9 / ((v9) | 1))
        switch "hello" {
        case "a":
            if v3 {
                v1 += 1
            } else {
                if len(v8) >= 2 {
                    v8 = v8[1:len(v8)]
                }
                var v12 int16 = -21
                _ = v12
            }
        default:
            fmt.Println(v3, v1)
        }
        _ = v11
    }
    var v13 uint8 = 255
    fmt.Println(v13, v2, v13)
    v13 = v13
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
