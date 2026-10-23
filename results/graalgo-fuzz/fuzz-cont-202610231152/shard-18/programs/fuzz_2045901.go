package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int; F2 int32 }

func main() {
    var v1 float64 = 0.1
    v2 := []int{-907, -2147483648, -454, -216}
    var v3 uint = 0
    var v4 uint = (^(v3))
    v5 := GenS1{F0: 18446744073709551615, F1: 127, F2: int32(v4)}
    if ("" >= "a") {
        switch 128 {
        case 128:
            fmt.Println(v3, v4)
        case -280:
            v4 += (v4 - 587)
        default:
            v6 := GenS1{F0: 4294967295, F1: -574, F2: -1}
            _ = v6
        }
        fmt.Println(v3)
    }
    for fi7 := 0; fi7 < 16; fi7++ {
        if false {
            v8 := GenS1{F0: 4294967295, F1: fi7, F2: 2147483647}
            v3 += (v3 << ((65535) & 63))
            var v9 uint16 = 226
            _ = v8
            _ = v9
        } else {
            if (((v1 * v1) - 0.0) > v1) {
                if true {
                    v10 := GenS1{F0: 42, F1: 9223372036854775807, F2: -573}
                    _ = v10
                }
                v2 = append(v2, (^(fi7)))
                var v11 int32 = 0
                _ = v11
            } else {
                if (4294967295 > 2147483647) {
                    var v12 uint32 = 1
                    _ = v12
                }
            }
        }
        v13 := GenS1{F0: 454, F1: (fi7 ^ (-(fi7))), F2: 0}
        _ = fi7
        _ = v13
    }
    v4 -= v4
    fmt.Println(v1, v3, v4)
    v1 -= v1
    var v14 bool = (!true)
    fmt.Println(v1)
    v4 = 548
    v2 = append(v2, -203)
    v4 += v3
    if true {
        if len(v2) > 0 {
            v2[0] = -2147483648
        }
        v14 = (v3 != 4294967295)
        var v15 int = -844
        _ = v15
    } else {
        switch (("\n\t" + ("" + "foo bar")) + ("hello" + "a")) {
        case "foo bar":
            v16 := GenS1{F0: 1, F1: 1, F2: 682}
            _ = v16
        case "\n\t":
            for fi17 := 0; fi17 < 7; fi17++ {
                v18 := GenS1{F0: 42, F1: (fi17 % (((fi17 >> ((930) & 63))) | 1)), F2: -1}
                if v14 {
                    v4 -= (v3 ^ (v3 | v3))
                }
                var v19 int = fi17
                _ = fi17
                _ = v18
                _ = v19
            }
        case "hello":
            var v20 int16 = -32768
            _ = v20
        }
        var v21 float64 = (float64((-2.25 * float32(3.14159))) / float64((1.5 / 3.14159)))
        _ = v21
    }
    var v22 uint64 = 421
    v23 := GenS1{F0: 2, F1: -42, F2: int32(v3)}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
