package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64 }
type GenS2 struct { F0 int8; F1 string; F2 uint64 }

func main() {
    v1 := make(map[string]uint32)
    v2 := []uint32{420, 65535, 4294967295, 2147483648}
    v3 := []string{"a", "a", "hello", ("hello" + ("" + ((("foo bar" + "") + "") + "a")))}
    var v4 int = 32767
    var v5 uint = 460
    var v6 int64 = 9223372036854775807
    v7 := v1["foo bar"]
    if (2 != 42) {
        var v8 int = v4
        for fi9 := 0; fi9 < 6; fi9++ {
            v1[""] = v7
            var v10 int64 = v6
            fmt.Println(fi9, v5, v8)
            _ = fi9
            _ = v10
        }
        _ = v8
    } else {
        switch -81 {
        case -594:
            v4 += v4
        default:
            v7 = (v7 >> (((v5 / ((298) | 1))) & 63))
        }
    }
    var v11 float32 = 1e10
    for fi12 := 0; fi12 < 2; fi12++ {
        var v13 float32 = v11
        _ = fi12
        _ = v13
    }
    if len(v2) > 0 {
        v2[0] = (v7 * (^(v7)))
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if false {
        fmt.Println(v4, v5)
        switch "\n\t" {
        case "hello":
            switch (v4 | v4) {
            case 128:
                var v14 int32 = 0
                _ = v14
            case -526:
                var v15 float32 = 1.5
                _ = v15
            default:
                for fi16 := 0; fi16 < 15; fi16++ {
                    v17 := GenS1{F0: v7, F1: 18446744073709551615}
                    _ = fi16
                    _ = v17
                }
            }
        case "foo bar":
            v6 = (v6 - (v6 * (v6 >> ((v5) & 63))))
        case "":
            _ = len(v2)
        }
        switch -2147483648 {
        case -2147483648:
            if false {
                var v18 uint8 = 86
                delete(v1, "")
                _ = v18
            } else {
                v11 = 1.5
                v5 -= v5
            }
        case -315:
            v19 := GenS2{F0: 1, F1: (("a" + (("" + ("hello" + ("foo bar" + ("hello" + "\n\t")))) + "")) + ("\n\t" + ("a" + "foo bar"))), F2: 1}
            _ = v19
        }
    }
    v5 = 4294967295
    for fi20 := 0; fi20 < 6; fi20++ {
        v21 := GenS1{F0: (v7 - v7), F1: 0}
        if len(v2) > 0 {
            v2[0] = v7
        }
        if true {
            v1[("a" + ("a" + ""))] = 4294967295
            fmt.Println(v5, v5, fi20)
        } else {
            var v22 int64 = 9223372036854775807
            var v23 float64 = ((((((float64(0.1) + -2.25) + 0.1) / (((-2.25 - 0.0) / 0.1) - (3.14159 / float64((v11 + ((v11 - (v11 + v11)) + (v11 / 1.0))))))) + 1.5) - (1e10 * (((1.5 + ((float64((float32(-2.25) * v11)) / 3.14159) + 1e10)) / (0.1 * ((1.5 / 0.1) + 1e10))) + float64((0.0 / float32(0.1)))))) + float64(0.1))
            _ = v22
            _ = v23
        }
        _ = fi20
        _ = v21
    }
    for fi24 := 0; fi24 < 6; fi24++ {
        for fi25 := 0; fi25 < 3; fi25++ {
            v26 := GenS1{F0: v7, F1: 0}
            v6 = 926
            var v27 string = "a"
            _ = fi25
            _ = v26
            _ = v27
        }
        _ = fi24
    }
    fmt.Println(v5, v7, v5)
    v4 = (v4 << (((v5 | 0)) & 63))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
