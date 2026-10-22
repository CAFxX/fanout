package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16; F2 int32 }

func main() {
    var v1 int8 = -64
    v2 := make(map[string]string)
    var v3 string = "foo bar"
    v4 := []uint32{4294967295, 65535}
    var v5 uint16 = 438
    v6 := GenS1{F0: 4294967295, F1: 610, F2: 866}
    if false {
        if (!false) {
            v1 = (v1 + (v1 % ((1) | 1)))
        } else {
            for fi7 := 0; fi7 < 16; fi7++ {
                fmt.Println(v1, v3, v3)
                v8 := GenS1{F0: 4294967295, F1: (v5 << ((18446744073709551615) & 63)), F2: 268}
                _ = fi7
                _ = v8
            }
        }
        var v9 uint16 = 40
        _ = v9
    } else {
        v3 = ""
        v3 = v3
        var v10 float32 = 3.14159
        _ = v10
    }
    switch v3 {
    case "a":
        if (!false) {
            var v11 uint16 = 42
            v3 = ("" + "hello")
            _ = v11
        } else {
            for fi12 := 0; fi12 < 20; fi12++ {
                v1 = 127
                v1 -= (v1 & (v1 ^ v1))
                _ = fi12
            }
            for fi13 := 0; fi13 < 4; fi13++ {
                _ = len(v4)
                _ = fi13
            }
            var v14 uint8 = uint8(v5)
            _ = v14
        }
    case "\n\t":
        var v15 float32 = float32(0.0)
        _ = v15
    case "hello":
        v3 = v3
    }
    if true {
        var v16 uint16 = 1
        _ = v16
    } else {
        v2[v3] = (v3 + (("" + v3) + (v3 + v3)))
    }
    var v17 int64 = -9223372036854775808
    for fi18 := 0; fi18 < 20; fi18++ {
        var v19 float64 = 0.0
        var v20 string = "\n\t"
        fmt.Println(v5, v5)
        _ = fi18
        _ = v19
        _ = v20
    }
    v21 := GenS1{F0: 439, F1: 1, F2: -2147483648}
    v1 -= v1
    fmt.Println(v1, v5, v3)
    _ = len(v2)
    _ = len(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
