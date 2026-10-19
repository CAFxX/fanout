package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8; F2 int16 }

func main() {
    var v1 int = 0
    v2 := make(map[string]int)
    var v3 uint16 = 1
    v4 := []int{v1}
    var v5 int64 = 1
    var v6 uint = 255
    v5 = v5
    v5 = 67
    switch 9223372036854775807 {
    case 4294967295:
        var v7 int64 = (-(v5))
        _ = v7
    case -9223372036854775808:
        var v8 float64 = 1e10
        _ = v8
    case 513:
        v5 += v5
    default:
        for fi9 := 0; fi9 < 1; fi9++ {
            for fi10 := 0; fi10 < 8; fi10++ {
                var v11 int64 = (v5 + 378)
                _ = fi10
                _ = v11
            }
            v6 += v6
            var v12 uint16 = v3
            _ = fi9
            _ = v12
        }
    }
    if (9223372036854775807 > (v5 & (^(v5)))) {
        var v13 int16 = 42
        if ("a" <= "hello") {
            v13 = v13
            for fi14 := 0; fi14 < 20; fi14++ {
                var v15 uint64 = 868
                var v16 int32 = -2147483648
                for fi17 := 0; fi17 < 3; fi17++ {
                    v2["hello"] = (fi17 & 42)
                    _ = fi17
                }
                _ = fi14
                _ = v15
                _ = v16
            }
        }
        var v18 float64 = ((0.1 * 3.14159) / 3.14159)
        _ = v13
        _ = v18
    } else {
        var v19 string = (("a" + "a") + ("foo bar" + (("\n\t" + "foo bar") + "foo bar")))
        fmt.Println(v3, v5, v1)
        for fi20 := 0; fi20 < 17; fi20++ {
            v6 = v6
            _ = fi20
        }
        _ = v19
    }
    for fi21 := 0; fi21 < 14; fi21++ {
        var v22 int8 = -28
        v23 := GenS1{F0: "\n\t", F1: (v22 << ((v6) & 63)), F2: 0}
        v1 += (v1 - 4294967295)
        _ = fi21
        _ = v22
        _ = v23
    }
    v1 = (v1 * int(v6))
    if ((v6 * (v6 * (v6 - v6))) < (^(v6))) {
        switch (v1 >> ((94) & 63)) {
        case 32768:
            v6 = (v6 & (v6 & 255))
        case 127:
            _ = len(v2)
        default:
            if len(v4) > 0 {
                v4[0] = (v1 % ((32768) | 1))
            }
        }
    }
    _ = len(v4)
    var v24 bool = (false || (4294967295 < 65535))
    v25 := GenS1{F0: ("" + "\n\t"), F1: -128, F2: -519}
    for fi26 := 0; fi26 < 14; fi26++ {
        v27 := GenS1{F0: "foo bar", F1: 100, F2: -32768}
        fmt.Println(v3, v3)
        v1 = (fi26 - 128)
        _ = fi26
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
