package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32; F2 int }

func main() {
    var v1 string = "\n\t"
    var v2 float32 = 0.0
    var v3 int16 = 1
    fmt.Println(v3, v1)
    v1 = "foo bar"
    for fi4 := 0; fi4 < 8; fi4++ {
        for fi5 := 0; fi5 < 11; fi5++ {
            var v6 uint16 = 854
            var v7 bool = (!(569 > -980))
            _ = fi5
            _ = v6
            _ = v7
        }
        v2 -= -2.25
        v3 -= v3
        _ = fi4
    }
    var v8 float64 = 0.1
    if (!(!(false || false))) {
        v2 -= v2
        v2 = v2
        var v9 int32 = 2147483647
        _ = v9
    } else {
        v10 := GenS1{F0: (v3 ^ (v3 & 1)), F1: 42, F2: -136}
        _ = v10
    }
    fmt.Println(v1, v1)
    if (!false) {
        if (true || true) {
            fmt.Println(v8, v3)
            var v11 int64 = 9223372036854775807
            _ = v11
        } else {
            v8 -= 1e10
        }
        v12 := GenS1{F0: (v3 / ((42) | 1)), F1: 699, F2: 32767}
        _ = v12
    } else {
        var v13 float64 = -2.25
        _ = v13
    }
    v8 -= (v8 + float64(0.0))
    fmt.Println(v3, v2, v2)
    v14 := GenS1{F0: (^(v3)), F1: 2, F2: 32768}
    v1 = (("a" + ("foo bar" + v1)) + (v1 + v1))
    fmt.Println(v3, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v8
}
