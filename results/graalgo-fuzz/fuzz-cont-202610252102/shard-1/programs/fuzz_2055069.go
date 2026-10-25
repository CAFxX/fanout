package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 uint64 }
type GenS2 struct { F0 uint8; F1 uint64; F2 string }
type GenS3 struct { F0 int8; F1 uint8; F2 uint8 }

func main() {
    v1 := []string{}
    var v2 int16 = 32767
    var v3 uint32 = 2
    var v4 uint32 = v3
    for fi5 := 0; fi5 < 10; fi5++ {
        v6 := GenS1{F0: 182, F1: (("" + "hello") + ""), F2: 945}
        var v7 int8 = 1
        var v8 uint = 65535
        _ = fi5
        _ = v6
        _ = v7
        _ = v8
    }
    var v9 int32 = 1
    fmt.Println(v2)
    for fi10 := 0; fi10 < 20; fi10++ {
        var v11 int8 = -109
        v4 = v4
        _ = fi10
        _ = v11
    }
    v12 := GenS2{F0: 1, F1: 0, F2: "foo bar"}
    v3 -= (^(v3))
    switch 9223372036854775807 {
    case 271:
        v9 = v9
    case -1:
        var v13 int8 = 127
        _ = v13
    default:
        if ((v2 | v2) > v2) {
            v9 -= v9
            fmt.Println(v2)
            var v14 int32 = -139
            _ = v14
        }
    }
    v15 := GenS1{F0: v9, F1: "a", F2: 172}
    fmt.Println(v4)
    v2 = 42
    var v16 bool = ((!(998 != 42)) || true)
    var v17 bool = v16
    if (!v17) {
        fmt.Println(v3)
        var v18 float64 = (3.14159 * ((1e10 / 1.0) / float64((-2.25 / 1.5))))
        switch -773 {
        case -2147483648:
            var v19 int8 = 35
            _ = v19
        case 32767:
            for fi20 := 0; fi20 < 4; fi20++ {
                v2 = 312
                _ = fi20
            }
        default:
            v21 := GenS3{F0: -9, F1: 42, F2: 51}
            _ = v21
        }
        _ = v18
    } else {
        v22 := GenS1{F0: 1, F1: "\n\t", F2: 18446744073709551615}
        v23 := GenS2{F0: 128, F1: 18446744073709551615, F2: "\n\t"}
        _ = v22
        _ = v23
    }
    for fi24 := 0; fi24 < 4; fi24++ {
        if (v9 > 1) {
            fmt.Println(fi24, v4, v17)
        }
        v25 := GenS3{F0: -23, F1: 226, F2: 127}
        var v26 string = ("\n\t" + ("" + "foo bar"))
        _ = fi24
        _ = v25
        _ = v26
    }
    var v27 uint = 255
    var v28 uint64 = 42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v9
}
