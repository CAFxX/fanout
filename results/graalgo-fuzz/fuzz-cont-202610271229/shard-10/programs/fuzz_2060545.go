package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool }

func main() {
    var v1 int64 = -9223372036854775808
    var v2 uint = 846
    var v3 int32 = 2147483647
    if ((v3 % ((v3) | 1)) <= (v3 & (v3 / (((v3 * 2147483647)) | 1)))) {
        switch ("a" + ("foo bar" + "")) {
        case "":
            v2 -= (v2 << ((255) & 63))
        case "\n\t":
            var v4 uint64 = 42
            _ = v4
        }
        for fi5 := 0; fi5 < 16; fi5++ {
            var v6 uint32 = 1
            v7 := []int{}
            _ = fi5
            _ = v6
            _ = v7
        }
        v3 += (v3 / (((v3 * 260)) | 1))
    }
    v8 := GenS1{F0: 2, F1: (!(!(2147483647 >= 42)))}
    v3 += (^(v3))
    var v9 uint8 = 1
    if (((0.0 >= (0.0 * -2.25)) || ((false && false) && true)) || ((v1 < v1) && true)) {
        for fi10 := 0; fi10 < 7; fi10++ {
            v11 := []int{-9223372036854775808, 42, fi10, fi10}
            var v12 uint64 = 827
            fmt.Println(v9, fi10, v3)
            _ = fi10
            _ = v11
            _ = v12
        }
    }
    var v13 int64 = 1
    for fi14 := 0; fi14 < 11; fi14++ {
        var v15 int = fi14
        for fi16 := 0; fi16 < 4; fi16++ {
            v9 = v9
            _ = fi16
        }
        _ = fi14
        _ = v15
    }
    var v17 int64 = (v1 >> ((v2) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v17
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
