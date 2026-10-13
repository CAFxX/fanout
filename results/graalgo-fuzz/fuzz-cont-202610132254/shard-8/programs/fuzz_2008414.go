package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 bool }

func main() {
    var v1 uint32 = 2147483648
    var v2 uint8 = 245
    var v3 int32 = 1
    var v4 uint = 609
    var v5 int16 = 0
    var v6 uint8 = 1
    switch int(v1) {
    case -178:
        fmt.Println(v1)
    case 128:
        if (uint32(v2) != v1) {
            v5 -= v5
            if ((v1 * v1) == 2) {
                switch -701 {
                case 4294967295:
                    v7 := GenS1{F0: 464}
                    _ = v7
                }
                v8 := []int{224, -1, 32768, 128}
                v4 = (v4 ^ (^(v4)))
                _ = v8
            } else {
                v4 = (v4 - 0)
            }
            v1 += 2
        } else {
            if (0 == v5) {
                var v9 int32 = v3
                _ = v9
            }
        }
    default:
        var v10 uint = 4294967295
        _ = v10
    }
    v1 -= (v1 % (((v1 << (((v4 ^ uint(v1))) & 63))) | 1))
    var v11 int64 = 42
    v12 := []int{470, int(v11), -369, 9223372036854775807}
    switch "foo bar" {
    case "a":
        var v13 uint32 = 0
        _ = v13
    default:
        _ = len(v12)
    }
    fmt.Println(v4, v2, v4)
    v4 = v4
    if len(v12) >= 2 {
        v12 = v12[1:len(v12)]
    }
    var v14 string = "hello"
    v5 = (v5 | 32767)
    var v15 float64 = 1.5
    var v16 int16 = v5
    var v17 bool = ("foo bar" < ((v14 + v14) + (v14 + "hello")))
    v15 -= v15
    v18 := GenS2{F0: (float64(0.0) != 0.0)}
    _ = len(v12)
    for fi19 := 0; fi19 < 2; fi19++ {
        var v20 int64 = v11
        _ = fi19
        _ = v20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
