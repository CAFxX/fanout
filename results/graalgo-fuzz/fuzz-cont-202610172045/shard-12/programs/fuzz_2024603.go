package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 string }
type GenS2 struct { F0 int32; F1 int64; F2 int16 }

func main() {
    var v1 uint = 4294967295
    v2 := []string{"hello", "\n\t", "hello", ("\n\t" + "hello")}
    v3 := []string{"a", ("a" + ((("\n\t" + "") + ("\n\t" + ("\n\t" + (((("foo bar" + "foo bar") + "\n\t") + "hello") + "")))) + ("a" + "foo bar"))), ("" + "")}
    var v4 uint64 = 42
    var v5 string = ""
    v6 := make(map[string]bool)
    for fi7 := 0; fi7 < 11; fi7++ {
        fmt.Println(v5)
        if (65535 < 255) {
            var v8 uint8 = 255
            _ = v8
        } else {
            v9 := GenS1{F0: 42, F1: 2147483647, F2: (((v5 + "") + v5) + "a")}
            for fi10 := 0; fi10 < 15; fi10++ {
                var v11 uint64 = 0
                if false {
                    v12 := GenS2{F0: -1, F1: 0, F2: 368}
                    _ = v12
                } else {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    fmt.Println(v11, fi10)
                }
                _ = fi10
                _ = v11
            }
            if len(v2) > 0 {
                v2[0] = "hello"
            }
            _ = v9
        }
        _ = fi7
    }
    fmt.Println(v5)
    var v13 int32 = 2147483647
    var v14 uint16 = 173
    v15 := GenS2{F0: (v13 - 704), F1: -89, F2: -835}
    var v16 uint8 = 218
    switch -483 {
    case 32767:
        delete(v6, "foo bar")
    case 210:
        var v17 float64 = 1.5
        _ = v17
    default:
        if (42 != v16) {
            var v18 float64 = 0.1
            var v19 uint = (v1 >> ((1) & 63))
            for fi20 := 0; fi20 < 14; fi20++ {
                var v21 float32 = (1.5 - 0.1)
                var v22 float32 = float32(v18)
                v1 = (v19 * v1)
                _ = fi20
                _ = v21
                _ = v22
            }
            _ = v18
            _ = v19
        }
    }
    v16 = 42
    var v23 uint = v1
    for fi24 := 0; fi24 < 19; fi24++ {
        if len(v2) > 0 {
            v2[0] = ((v5 + v5) + (v5 + ("" + (v5 + "hello"))))
        }
        if ((v14 / ((v14) | 1)) > (v14 * (v14 - 309))) {
            v4 = (v4 - v4)
        }
        v25 := GenS2{F0: v13, F1: 445, F2: int16(v14)}
        _ = fi24
        _ = v25
    }
    switch 334 {
    case -2147483648:
        v26 := v6[v5]
        _ = v26
    case 42:
        v27 := GenS2{F0: (v13 | v13), F1: 1, F2: -1}
        _ = v27
    default:
        fmt.Println(v16)
    }
    var v28 int32 = (^(v13))
    v4 -= v4
    switch "" {
    case "foo bar":
        var v29 uint8 = (v16 + 42)
        _ = v29
    default:
        v1 = v23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v23
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
