package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8; F2 uint16 }
type GenS2 struct { F0 uint64; F1 string; F2 uint64 }

func main() {
    v1 := make(map[uint32]string)
    v2 := []string{"a", ((("foo bar" + ("foo bar" + "hello")) + "\n\t") + "\n\t")}
    var v3 uint32 = 953
    var v4 int16 = -1
    var v5 float64 = (-2.25 - ((float64(0.0) / (((0.1 / -2.25) - (1e10 / 1.0)) - -2.25)) + (float64(1e10) * float64(1e10))))
    var v6 uint32 = (v3 * v3)
    fmt.Println(v6)
    var v7 uint8 = 42
    var v8 uint = 42
    for fi9 := 0; fi9 < 6; fi9++ {
        v10 := GenS1{F0: (("" + ("foo bar" + "\n\t")) + ("" + (("" + "") + "foo bar"))), F1: (v7 >> ((0) & 63)), F2: 935}
        _ = fi9
        _ = v10
    }
    var v11 int8 = 2
    var v12 int8 = (^(v11))
    var v13 uint8 = (v7 - v7)
    if (v8 <= 189) {
        switch "a" {
        case "a":
            if ((v13 * 2) <= (^(v7))) {
                v7 = v7
                fmt.Println(v8, v11, v13)
                if ((true || false) && (0 < 19)) {
                    v14 := GenS1{F0: "foo bar", F1: v7, F2: 255}
                    _ = v14
                } else {
                    for fi15 := 0; fi15 < 12; fi15++ {
                        var v16 uint32 = v3
                        for fi17 := 0; fi17 < 1; fi17++ {
                            var v18 uint64 = 42
                            v13 = 82
                            _ = fi17
                            _ = v18
                        }
                        var v19 int32 = 1
                        _ = fi15
                        _ = v16
                        _ = v19
                    }
                    var v20 float32 = -2.25
                    _ = v20
                }
            } else {
                var v21 int8 = (v12 % ((v12) | 1))
                _ = len(v2)
                v2 = append(v2, "\n\t")
                _ = v21
            }
        default:
            v2 = append(v2, "")
        }
        var v22 uint16 = 255
        v23 := GenS1{F0: "hello", F1: v13, F2: v22}
        _ = v22
        _ = v23
    } else {
        v24 := GenS1{F0: "\n\t", F1: (v13 - v7), F2: 0}
        if (65535 == 65535) {
            v5 += 0.1
            for fi25 := 0; fi25 < 3; fi25++ {
                v6 = (v6 & uint32(v13))
                v5 = v5
                _ = fi25
            }
            v6 = v3
        }
        _ = v24
    }
    v26 := GenS2{F0: 331, F1: "foo bar", F2: uint64(v6)}
    if (928 >= -454) {
        v27 := v1[v3]
        _ = v27
    } else {
        var v28 uint32 = v3
        var v29 float32 = 1.5
        fmt.Println(v7)
        _ = v28
        _ = v29
    }
    v6 = v6
    v3 -= (v6 - (v6 >> ((579) & 63)))
    for fi30 := 0; fi30 < 20; fi30++ {
        v13 -= 169
        var v31 int64 = -9223372036854775808
        _ = fi30
        _ = v31
    }
    fmt.Println(v12, v6)
    switch "foo bar" {
    case "a":
        switch ("hello" + (("hello" + "hello") + (("" + ((("a" + "foo bar") + "") + (((("\n\t" + ("foo bar" + ("hello" + "\n\t"))) + "\n\t") + "foo bar") + "hello"))) + ""))) {
        case "foo bar":
            var v32 int64 = 415
            _ = v32
        default:
            var v33 bool = true
            _ = v33
        }
    case "\n\t":
        v3 -= 2147483647
    }
    v11 = (v11 / ((v11) | 1))
    fmt.Println(v8, v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
