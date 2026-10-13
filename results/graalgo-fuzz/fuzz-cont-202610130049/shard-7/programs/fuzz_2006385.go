package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }
type GenS2 struct { F0 uint }

func main() {
    var v1 int16 = -32768
    var v2 int64 = 9223372036854775807
    var v3 uint = 115
    v4 := []string{"\n\t"}
    var v5 string = (("a" + (("\n\t" + "foo bar") + ("\n\t" + "hello"))) + "\n\t")
    v3 = v3
    if len(v4) > 0 {
        v4[0] = v5
    }
    if (1 <= v3) {
        var v6 int16 = v1
        if (true && true) {
            var v7 uint32 = 2147483648
            v1 = (-(v1))
            var v8 string = "hello"
            _ = v7
            _ = v8
        } else {
            v9 := GenS1{F0: -97, F1: 2147483647}
            _ = v9
        }
        _ = v6
    }
    var v10 uint32 = 237
    var v11 uint8 = 229
    fmt.Println(v5, v2, v3)
    var v12 int64 = -9223372036854775808
    v10 -= (v10 % ((375) | 1))
    var v13 int16 = 840
    v14 := GenS1{F0: 32767, F1: -19}
    for fi15 := 0; fi15 < 9; fi15++ {
        v10 -= v10
        v16 := GenS1{F0: (fi15 ^ -2147483648), F1: -1}
        switch (^(fi15)) {
        case 4294967295:
            fmt.Println(v12, fi15, v11)
        }
        _ = fi15
        _ = v16
    }
    fmt.Println(v2, v5, v1)
    if (false && false) {
        fmt.Println(v12, v11)
    } else {
        v10 = v10
        v13 = v13
        v13 -= (v1 ^ v13)
    }
    for fi17 := 0; fi17 < 1; fi17++ {
        v18 := GenS2{F0: (v3 | (v3 * 255))}
        if len(v4) > 0 {
            v4[0] = v5
        }
        _ = fi17
        _ = v18
    }
    var v19 float32 = float32((3.14159 + float64((((1e10 / 1.5) / 1e10) * 1e10))))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
