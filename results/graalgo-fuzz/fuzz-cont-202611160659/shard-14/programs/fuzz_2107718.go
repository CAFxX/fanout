package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 int }
type GenS2 struct { F0 string }
func genFn0() uint32 {
    switch -359 {
    case -2147483648:
        v4 := GenS1{F0: 91, F1: 125, F2: -1}
        _ = v4
    }
    var v5 uint16 = 2
    fmt.Println(v5)
    _ = v5
    return 208
}

func main() {
    var v1 int32 = -2147483648
    v2 := []uint32{65535}
    var v3 uint32 = 229
    fmt.Println(v3, v3)
    _ = len(v2)
    for fi6 := 0; fi6 < 3; fi6++ {
        fmt.Println(v3, v3)
        for fi7 := 0; fi7 < 14; fi7++ {
            var v8 int16 = 32767
            v3 = v3
            if false {
                if true {
                    v8 = (^(v8))
                    var v9 bool = (false || (9223372036854775807 > -9223372036854775808))
                    v8 = (v8 + int16(v3))
                    _ = v9
                } else {
                    var v10 float32 = 1e10
                    _ = v10
                }
            } else {
                v11 := GenS2{F0: "\n\t"}
                _ = v11
            }
            _ = fi7
            _ = v8
        }
        _ = fi6
    }
    v12 := GenS2{F0: "a"}
    switch "a" {
    case "\n\t":
        if len(v2) > 0 {
            v2[0] = (v3 & (v3 << ((255) & 63)))
        }
    case "hello":
        v3 = 461
    default:
        v3 -= (v3 - v3)
    }
    v1 += -1
    switch -299 {
    case 209:
        var v13 uint32 = genFn0()
        _ = v13
    case -9223372036854775808:
        v14 := GenS1{F0: 1, F1: 75, F2: 637}
        _ = v14
    case 32768:
        fmt.Println(v3, v1)
    }
    _ = len(v2)
    v2 = append(v2, v3)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
}
