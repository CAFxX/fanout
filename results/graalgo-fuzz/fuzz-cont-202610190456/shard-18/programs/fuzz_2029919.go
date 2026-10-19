package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint64; F1 uint }
type GenS3 struct { F0 uint8; F1 string }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int16; F1 uint; F2 int16 }

func main() {
    v1 := make(map[string]uint32)
    var v2 float64 = 0.1
    var v3 float32 = (1.5 / 1e10)
    var v4 int8 = 100
    v5 := []int{-2147483648, -1}
    _ = len(v5)
    v6 := GenS1{F0: (("foo bar" + "\n\t") + (("a" + "foo bar") + ("" + ("hello" + ("a" + "")))))}
    _ = len(v1)
    if (!true) {
        v3 = (v3 / v3)
    } else {
        if (!false) {
            v7 := GenS1{F0: ("a" + "")}
            _ = v7
        } else {
            for fi8 := 0; fi8 < 16; fi8++ {
                for fi9 := 0; fi9 < 16; fi9++ {
                    switch ((("" + "\n\t") + "foo bar") + "a") {
                    case "hello":
                        v2 -= v2
                    }
                    switch (fi9 + (fi9 >> ((0) & 63))) {
                    case 37:
                        v2 -= ((((float64(float32(float64(v3))) + ((v2 + (float64(v3) + 1.5)) + float64((1e10 - float32(3.14159))))) * v2) - v2) - v2)
                    default:
                        v10 := v1[""]
                        _ = v10
                    }
                    var v11 int64 = 0
                    _ = fi9
                    _ = v11
                }
                _ = fi8
            }
        }
    }
    var v12 float32 = 1.5
    v1["hello"] = 42
    var v13 uint = 542
    v14 := GenS2{F0: uint64(v13), F1: v13}
    if false {
        var v15 uint16 = 979
        _ = v15
    } else {
        if (!(!((("" + "a") + (("foo bar" + (("a" + (("foo bar" + "a") + "hello")) + "")) + "\n\t")) < "a"))) {
            var v16 uint = (v13 << ((2) & 63))
            var v17 uint = (v13 / ((18446744073709551615) | 1))
            v13 += (v16 & v13)
            _ = v16
            _ = v17
        } else {
            var v18 uint64 = 4294967295
            for fi19 := 0; fi19 < 4; fi19++ {
                var v20 int16 = 32767
                _ = fi19
                _ = v20
            }
            for fi21 := 0; fi21 < 12; fi21++ {
                v22 := GenS1{F0: (("\n\t" + "hello") + "\n\t")}
                v23 := GenS3{F0: 1, F1: ("\n\t" + "")}
                var v24 int32 = -1
                _ = fi21
                _ = v22
                _ = v23
                _ = v24
            }
            _ = v18
        }
        for fi25 := 0; fi25 < 15; fi25++ {
            var v26 int8 = v4
            _ = fi25
            _ = v26
        }
        v27 := GenS4{F0: 133}
        _ = v27
    }
    v28 := GenS4{F0: 32768}
    v13 -= (v13 * v13)
    v29 := GenS5{F0: 0}
    switch -2147483648 {
    case -232:
        switch 4294967295 {
        case -817:
            delete(v1, "foo bar")
        case 0:
            var v30 int8 = -1
            _ = v30
        default:
            switch "" {
            case "hello":
                fmt.Println(v12, v3)
            default:
                v31 := GenS6{F0: 32767, F1: (v13 | (^(v13))), F2: -1}
                _ = v31
            }
        }
    case -893:
        v2 = float64((v3 + (float32(v2) + (v12 - 3.14159))))
    case -1:
        v2 += (0.0 - float64((v3 * 0.0)))
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
