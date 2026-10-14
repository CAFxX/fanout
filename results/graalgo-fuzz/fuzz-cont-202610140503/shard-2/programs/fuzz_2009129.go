package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint; F2 int32 }
type GenS2 struct { F0 bool; F1 uint64 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int16; F1 bool; F2 int }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 string }
func genFn0(p0 string, p1 uint32, p2 string) int8 {
    v6 := GenS1{F0: 1, F1: 42, F2: -196}
    fmt.Println(p0)
    fmt.Println(p2, p0, p1)
    v7 := GenS1{F0: 4294967295, F1: 786, F2: 1}
    switch 806 {
    case 127:
        var v8 float64 = 1e10
        _ = v8
    case -992:
        for fi9 := 0; fi9 < 15; fi9++ {
            v10 := GenS1{F0: uint64(p1), F1: 255, F2: 42}
            _ = fi9
            _ = v10
        }
    case -2147483648:
        v11 := GenS2{F0: (415 == 0), F1: 1}
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 100
}
func genFn1(p0 bool) uint8 {
    p0 = p0
    p0 = p0
    _ = p0
    return 0
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int8 = 127
    var v3 int16 = -1
    var v4 float32 = (1e10 - 3.14159)
    var v5 uint = 42
    delete(v1, "foo bar")
    v4 = float32((1.5 / float64((v4 - 0.1))))
    for fi12 := 0; fi12 < 16; fi12++ {
        fmt.Println(v5, v4)
        v1["hello"] = 2147483648
        _ = fi12
    }
    switch ((("hello" + "") + "hello") + "a") {
    case "foo bar":
        var v13 float64 = -2.25
        _ = v13
    case "hello":
        for fi14 := 0; fi14 < 17; fi14++ {
            fmt.Println(fi14, v3, fi14)
            for fi15 := 0; fi15 < 3; fi15++ {
                var v16 uint16 = 1
                _ = fi15
                _ = v16
            }
            _ = fi14
        }
    }
    for fi17 := 0; fi17 < 13; fi17++ {
        for fi18 := 0; fi18 < 2; fi18++ {
            v19 := GenS3{F0: 810}
            if (false || (0.1 == 3.14159)) {
                v20 := v1[("foo bar" + ("\n\t" + ("" + ("\n\t" + "a"))))]
                fmt.Println(v2)
                if (0 > 0) {
                    v21 := GenS4{F0: v3, F1: true, F2: (fi18 | (^(fi18)))}
                    _ = v21
                }
                _ = v20
            } else {
                switch ("hello" + "a") {
                case "\n\t":
                    if ((1 >= 65535) && false) {
                        switch fi17 {
                        case 32768:
                            v22 := GenS5{F0: (("\n\t" + "foo bar") + "")}
                            _ = v22
                        default:
                            v23 := GenS6{F0: 9223372036854775807}
                            _ = v23
                        }
                        var v24 float32 = v4
                        _ = v24
                    } else {
                        v2 += v2
                        var v25 int64 = -9223372036854775808
                        _ = v25
                    }
                case "a":
                    var v26 int8 = -128
                    _ = v26
                default:
                    v27 := GenS7{F0: "foo bar"}
                    _ = v27
                }
            }
            _ = fi18
            _ = v19
        }
        v28 := GenS5{F0: ""}
        _ = fi17
        _ = v28
    }
    v29 := []int{128, -2147483648, -1}
    for fi30 := 0; fi30 < 19; fi30++ {
        delete(v1, "\n\t")
        _ = fi30
    }
    if len(v29) >= 2 {
        v29 = v29[1:len(v29)]
    }
    v2 += -89
    var v31 uint32 = 683
    var v32 uint = v5
    var v33 uint8 = 255
    for fi34 := 0; fi34 < 7; fi34++ {
        var v35 int = (fi34 | (fi34 * (fi34 - int(v31))))
        _ = fi34
        _ = v35
    }
    delete(v1, ("a" + "a"))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
}
