package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32 }

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[uint32]string)
    var v3 int32 = 2147483647
    var v4 int64 = -994
    v5 := []uint32{2147483648}
    v3 = (v3 * v3)
    var v6 string = ((("" + "a") + "") + "\n\t")
    v7 := GenS1{F0: 215, F1: 851}
    for fi8 := 0; fi8 < 2; fi8++ {
        v4 += 42
        if (12 < int8(v3)) {
            if (false || false) {
                v3 -= v3
            }
        } else {
            switch v6 {
            case "foo bar":
                var v9 int = 128
                _ = v9
            default:
                var v10 uint = 4294967295
                _ = v10
            }
        }
        _ = fi8
    }
    var v11 int32 = (^(v3))
    if (-2147483648 == 605) {
        v3 -= (v11 << ((18446744073709551615) & 63))
    }
    if false {
        if true {
            v4 = -150
            var v12 int32 = (^(v11))
            if true {
                var v13 uint32 = 900
                if len(v5) > 0 {
                    v5[0] = v13
                }
                fmt.Println(v3, v4, v6)
                _ = v13
            } else {
                v11 = (v3 / ((v12) | 1))
            }
            _ = v12
        }
    }
    v14 := v2[2147483647]
    v15 := GenS1{F0: 2, F1: 731}
    if len(v5) > 0 {
        v5[0] = 2147483647
    }
    v6 = "a"
    for fi16 := 0; fi16 < 18; fi16++ {
        var v17 uint8 = 238
        _ = fi16
        _ = v17
    }
    v11 -= v3
    var v18 uint64 = 4294967295
    var v19 uint32 = 2
    _ = len(v2)
    v4 = int64(v19)
    _ = len(v1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v14
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
