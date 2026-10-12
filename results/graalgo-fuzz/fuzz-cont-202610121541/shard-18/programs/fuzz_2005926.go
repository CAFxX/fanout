package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint16; F1 string }
type GenS3 struct { F0 uint; F1 uint; F2 int64 }

func main() {
    var v1 int8 = 42
    var v2 int8 = (^(v1))
    v3 := make(map[int]string)
    v4 := []int{955, 764, -2147483648, -2147483648}
    v5 := make(map[int]string)
    v6 := []string{(((("a" + (("a" + ("hello" + "a")) + (("a" + ("foo bar" + "hello")) + "\n\t"))) + ("" + (("a" + "foo bar") + "a"))) + "") + "a"), "foo bar", "", "\n\t"}
    v7 := GenS1{F0: false}
    if ((0.1 - 1.5) < 1e10) {
        var v8 bool = true
        for fi9 := 0; fi9 < 1; fi9++ {
            switch ("hello" + ("foo bar" + "\n\t")) {
            case "\n\t":
                v8 = v8
            }
            var v10 int = (-(fi9))
            _ = fi9
            _ = v10
        }
        v8 = v8
        _ = v8
    } else {
        v1 -= (v1 * 42)
        if false {
            if (((1.5 + 0.0) == (0.0 / -2.25)) || true) {
                v2 = (v2 / ((v1) | 1))
                var v11 uint8 = 255
                v12 := GenS1{F0: false}
                _ = v11
                _ = v12
            }
            var v13 uint8 = 244
            for fi14 := 0; fi14 < 7; fi14++ {
                var v15 uint8 = 128
                _ = fi14
                _ = v15
            }
            _ = v13
        }
        var v16 bool = true
        _ = v16
    }
    for fi17 := 0; fi17 < 13; fi17++ {
        if ("a" != "") {
            fmt.Println(fi17, v2)
            var v18 float64 = 0.1
            _ = v18
        }
        fmt.Println(fi17, fi17)
        v19 := GenS2{F0: 2, F1: ("\n\t" + "a")}
        _ = fi17
        _ = v19
    }
    v20 := GenS2{F0: 65535, F1: ("\n\t" + "\n\t")}
    var v21 uint8 = 73
    switch "foo bar" {
    case "foo bar":
        var v22 uint64 = 1
        _ = v22
    case "":
        _ = len(v3)
    case "\n\t":
        switch "\n\t" {
        case "a":
            switch 32768 {
            case 9223372036854775807:
                var v23 int = 32768
                _ = v23
            default:
                v24 := GenS3{F0: 65535, F1: 65535, F2: 1}
                _ = v24
            }
        default:
            if ((0 >= 0) || true) {
                v2 = (v1 + v2)
                for fi25 := 0; fi25 < 7; fi25++ {
                    _ = len(v6)
                    if (((-2.25 * ((3.14159 + 1.5) + 1e10)) * (1.5 * float32(-2.25))) == (((0.1 * 3.14159) - ((0.0 * (3.14159 / 3.14159)) - (1e10 * (0.0 / ((1.5 + (3.14159 + 1e10)) / 0.1))))) - (((((1.5 * 1.5) / 1e10) / 1.5) * (3.14159 + 0.1)) + ((((((((1.5 - 0.0) - float32(0.1)) - 0.1) - 1e10) / 0.1) - float32(float64(float32(float64((0.1 / -2.25)))))) * (3.14159 - (3.14159 - (3.14159 * (1.5 - -2.25))))) + -2.25)))) {
                        v6 = append(v6, "")
                        var v26 uint16 = 2
                        _ = v26
                    }
                    _ = fi25
                }
                if len(v6) > 0 {
                    v6[0] = "hello"
                }
            } else {
                if (false || true) {
                    _ = len(v3)
                    fmt.Println(v1, v21)
                    if true {
                        v2 = v2
                    }
                } else {
                    v2 = -1
                    fmt.Println(v1)
                }
            }
        }
    default:
        var v27 float64 = 1e10
        _ = v27
    }
    var v28 uint16 = 349
    delete(v3, 4294967295)
    for fi29 := 0; fi29 < 2; fi29++ {
        var v30 int8 = 1
        _ = fi29
        _ = v30
    }
    var v31 int = 42
    _ = len(v5)
    if true {
        if true {
            fmt.Println(v1, v21, v1)
            for fi32 := 0; fi32 < 3; fi32++ {
                v33 := GenS2{F0: v28, F1: ""}
                _ = fi32
                _ = v33
            }
            v21 = 2
        }
        v5[v31] = ""
        var v34 int = (v31 & 0)
        _ = v34
    }
    fmt.Println(v31)
    v2 += v1
    v35 := GenS3{F0: 2, F1: uint(v28), F2: 926}
    switch v31 {
    case -2147483648:
        v28 = (^(v28))
    case -9223372036854775808:
        v36 := GenS1{F0: true}
        _ = v36
    default:
        v37 := v5[(v31 & v31)]
        _ = v37
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v28
    _ = v3
    _ = v31
    _ = v35
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
