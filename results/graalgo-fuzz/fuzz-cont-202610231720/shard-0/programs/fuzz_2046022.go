package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 string; F1 uint32; F2 string }
func genFn0(p0 int) uint16 {
    p0 += p0
    var v4 float32 = (0.0 / ((3.14159 * 1.5) - (1.5 - (((float32(((0.0 - 0.1) - (1e10 / -2.25))) - 0.0) * -2.25) * 0.0))))
    v4 -= (1.5 - v4)
    _ = p0
    _ = v4
    return 42
}

func main() {
    v1 := []string{}
    var v2 uint8 = 127
    var v3 int = 215
    switch ("hello" + ("a" + "\n\t")) {
    case "a":
        fmt.Println(v2)
    default:
        v2 = 1
    }
    if (false || (true && ((v2 ^ v2) == 97))) {
        v5 := GenS1{F0: 1}
        _ = v5
    }
    fmt.Println(v2)
    if ((v3 >> ((2) & 63)) >= v3) {
        if (65535 <= 65535) {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            var v6 int64 = -9223372036854775808
            fmt.Println(v3, v3, v2)
            _ = v6
        } else {
            var v7 uint32 = 151
            _ = v7
        }
        v2 -= (v2 * (v2 * 171))
    } else {
        v3 -= (^(v3))
        switch ("a" + "a") {
        case "":
            if len(v1) > 0 {
                v1[0] = "foo bar"
            }
        case "foo bar":
            var v8 uint = 698
            _ = v8
        case "\n\t":
            v1 = append(v1, ("\n\t" + "hello"))
        default:
            var v9 uint64 = 2
            _ = v9
        }
        var v10 string = ("foo bar" + "foo bar")
        _ = v10
    }
    v2 -= v2
    switch 32768 {
    case 127:
        var v11 uint64 = 2
        _ = v11
    }
    v12 := GenS2{F0: (("a" + ("" + "foo bar")) + "a"), F1: 42, F2: ""}
    for fi13 := 0; fi13 < 13; fi13++ {
        v1 = append(v1, ("a" + "foo bar"))
        var v14 bool = (295 < 42)
        _ = fi13
        _ = v14
    }
    _ = len(v1)
    if (42 != -9223372036854775808) {
        v3 += v3
    } else {
        var v15 uint32 = 4294967295
        _ = v15
    }
    for fi16 := 0; fi16 < 16; fi16++ {
        var v17 int32 = 411
        var v18 int32 = 1
        _ = fi16
        _ = v17
        _ = v18
    }
    if (v3 != (v3 >> ((18446744073709551615) & 63))) {
        v19 := GenS1{F0: 0}
        _ = v19
    }
    v3 = 153
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
}
