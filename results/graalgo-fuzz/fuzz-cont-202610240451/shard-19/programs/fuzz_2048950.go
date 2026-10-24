package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32 }
type GenS2 struct { F0 uint; F1 int32 }

func main() {
    var v1 bool = (!(2147483647 >= 0))
    v2 := []uint32{2147483648}
    var v3 int = 0
    var v4 uint = 4294967295
    switch ("hello" + ("hello" + "\n\t")) {
    case "a":
        v3 -= (v3 - -116)
    case "\n\t":
        var v5 int = v3
        _ = v5
    default:
        v4 += (v4 >> (((v4 / (((v4 + 0)) | 1))) & 63))
    }
    v4 = v4
    for fi6 := 0; fi6 < 9; fi6++ {
        if (-128 <= 100) {
            v3 += -752
        }
        for fi7 := 0; fi7 < 1; fi7++ {
            _ = len(v2)
            v4 += v4
            _ = len(v2)
            _ = fi7
        }
        _ = fi6
    }
    if (!((false && true) && false)) {
        for fi8 := 0; fi8 < 3; fi8++ {
            if (2 == 2147483647) {
                fmt.Println(v4, v3, v3)
            }
            if v1 {
                var v9 uint64 = 4294967295
                fmt.Println(v3, v9, v9)
                var v10 uint8 = 128
                _ = v10
                _ = v9
            } else {
                if len(v2) > 0 {
                    v2[0] = 65535
                }
            }
            v11 := GenS1{F0: 32767, F1: -737}
            _ = fi8
            _ = v11
        }
        v12 := GenS2{F0: v4, F1: -454}
        _ = v12
    }
    var v13 uint8 = 135
    for fi14 := 0; fi14 < 1; fi14++ {
        var v15 float32 = (0.0 * 0.0)
        var v16 uint32 = 4294967295
        v2 = append(v2, v16)
        _ = fi14
        _ = v15
        _ = v16
    }
    var v17 uint = (v4 - (v4 ^ 65535))
    var v18 int8 = 100
    fmt.Println(v17)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
