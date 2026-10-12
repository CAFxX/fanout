package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16; F2 uint }
type GenS2 struct { F0 int64; F1 uint8; F2 int32 }
func genFn0(p0 string, p1 uint32, p2 uint64) uint16 {
    var v4 int8 = 100
    p2 -= (p2 % (((p2 * p2)) | 1))
    v5 := []int{32767, 1000, 571, 127}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 32767
}

func main() {
    var v1 int8 = 127
    var v2 int32 = 2147483647
    v3 := []int{32768, 2147483647}
    v6 := GenS1{F0: 127, F1: 42, F2: 255}
    var v7 string = ("\n\t" + "\n\t")
    fmt.Println(v1, v2, v1)
    switch (v7 + v7) {
    case "foo bar":
        v7 = v7
    case "":
        var v8 int32 = v2
        _ = v8
    }
    if (128 != 0) {
        fmt.Println(v2)
        if (1.5 >= 0.0) {
            v1 -= v1
            if true {
                v9 := GenS1{F0: -9223372036854775808, F1: 32767, F2: 255}
                v10 := GenS1{F0: 548, F1: -74, F2: 481}
                _ = v10
                _ = v9
            }
        }
        var v11 string = "foo bar"
        _ = v11
    } else {
        if (genFn0(((("" + (v7 + v7)) + (v7 + v7)) + ((v7 + (v7 + v7)) + ("foo bar" + ("foo bar" + v7)))), 65, 0) <= genFn0("foo bar", 25, 18446744073709551615)) {
            var v12 int16 = -881
            v1 = (v1 / (((^(v1))) | 1))
            _ = v12
        } else {
            fmt.Println(v1, v1)
            var v13 uint32 = 627
            _ = v13
        }
    }
    v14 := GenS1{F0: 2147483647, F1: -1, F2: 18446744073709551615}
    var v15 int8 = v1
    fmt.Println(v1, v15)
    v16 := GenS2{F0: -9223372036854775808, F1: 127, F2: (^(v2))}
    v1 = v15
    var v17 int8 = -128
    var v18 int8 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v6
    _ = v7
}
