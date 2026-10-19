package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 int32 }
type GenS2 struct { F0 int; F1 uint16 }
func genFn0() string {
    var v4 uint32 = 2
    v5 := []int{127, 32767}
    if (18446744073709551615 <= 364) {
        var v6 uint8 = 1
        switch 646 {
        case 127:
            var v7 float32 = -2.25
            _ = v7
        default:
            fmt.Println(v4)
        }
        switch "" {
        case "":
            v4 += v4
        case "hello":
            if ((((115 > 0) || false) || (42 != -9223372036854775808)) && false) {
                v8 := GenS1{F0: v6, F1: 770, F2: -2147483648}
                for fi9 := 0; fi9 < 6; fi9++ {
                    v4 -= (v4 + (^(v4)))
                    _ = fi9
                }
                v4 -= (v4 / ((42) | 1))
                _ = v8
            }
        case "\n\t":
            v4 += 65535
        default:
            var v10 uint16 = 0
            _ = v10
        }
        _ = v6
    }
    v11 := GenS1{F0: 0, F1: 1, F2: -44}
    _ = v11
    _ = v4
    _ = v5
    return "foo bar"
}

func main() {
    var v1 float32 = float32(3.14159)
    var v2 uint64 = 249
    var v3 uint16 = 213
    var v12 uint8 = 23
    for fi13 := 0; fi13 < 11; fi13++ {
        var v14 int32 = -239
        _ = fi13
        _ = v14
    }
    var v15 int64 = 42
    for fi16 := 0; fi16 < 19; fi16++ {
        if (-32768 <= 42) {
            v2 = (v2 & v2)
            v17 := GenS2{F0: 32768, F1: 667}
            _ = v17
        }
        v18 := []int{}
        v2 += 99
        _ = fi16
        _ = v18
    }
    for fi19 := 0; fi19 < 18; fi19++ {
        v1 += v1
        _ = fi19
    }
    for fi20 := 0; fi20 < 11; fi20++ {
        var v21 int32 = 0
        _ = fi20
        _ = v21
    }
    v15 -= v15
    fmt.Println(v2, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v15
    _ = v2
    _ = v3
}
