package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64 }
type GenS2 struct { F0 int; F1 int8; F2 uint32 }
type GenS3 struct { F0 int64; F1 bool; F2 string }
type GenS4 struct { F0 uint8 }

func main() {
    var v1 int = 574
    v2 := make(map[string]bool)
    var v3 uint8 = 254
    v4 := []int{2147483647, v1, v1, (v1 - v1)}
    for fi5 := 0; fi5 < 9; fi5++ {
        for fi6 := 0; fi6 < 16; fi6++ {
            fmt.Println(fi6, fi6, fi5)
            var v7 int32 = 661
            var v8 bool = false
            _ = fi6
            _ = v7
            _ = v8
        }
        _ = fi5
    }
    switch "a" {
    case "\n\t":
        v3 = (v3 % ((v3) | 1))
    case "a":
        for fi9 := 0; fi9 < 3; fi9++ {
            var v10 int16 = -32768
            v3 = 47
            for fi11 := 0; fi11 < 16; fi11++ {
                v12 := GenS1{F0: 32767, F1: 18446744073709551615}
                var v13 uint16 = 32767
                _ = fi11
                _ = v12
                _ = v13
            }
            _ = fi9
            _ = v10
        }
    case "hello":
        for fi14 := 0; fi14 < 15; fi14++ {
            v15 := GenS1{F0: 32767, F1: 2}
            delete(v2, "foo bar")
            switch "\n\t" {
            case "hello":
                v3 += v3
            case "foo bar":
                v3 -= (v3 ^ v3)
            case "a":
                if false {
                    v1 -= -1
                    var v16 int = fi14
                    _ = len(v2)
                    _ = v16
                } else {
                    fmt.Println(v1)
                    v17 := GenS2{F0: v1, F1: 1, F2: 2147483647}
                    _ = v17
                }
            default:
                var v18 int8 = 1
                _ = v18
            }
            _ = fi14
            _ = v15
        }
    }
    fmt.Println(v1, v1)
    for fi19 := 0; fi19 < 18; fi19++ {
        for fi20 := 0; fi20 < 4; fi20++ {
            var v21 int = (^(v1))
            v22 := GenS3{F0: 9223372036854775807, F1: ((fi19 | (-(v21))) != 4294967295), F2: ("hello" + ("" + "foo bar"))}
            _ = fi20
            _ = v21
            _ = v22
        }
        fmt.Println(v3, fi19, v1)
        var v23 uint8 = 1
        _ = fi19
        _ = v23
    }
    var v24 int32 = 727
    v25 := GenS3{F0: 42, F1: true, F2: ("foo bar" + "\n\t")}
    v26 := GenS4{F0: (v3 % ((v3) | 1))}
    v3 += (v3 / ((v3) | 1))
    v3 = v3
    if (true || true) {
        var v27 float64 = 0.1
        fmt.Println(v24)
        fmt.Println(v27, v24)
        _ = v27
    }
    for fi28 := 0; fi28 < 16; fi28++ {
        var v29 string = ("a" + ("a" + "\n\t"))
        var v30 int = (v1 ^ (fi28 ^ 42))
        _ = fi28
        _ = v29
        _ = v30
    }
    var v31 uint64 = 4294967295
    var v32 uint64 = v31
    for fi33 := 0; fi33 < 4; fi33++ {
        v32 = (v32 ^ (v31 - v31))
        v24 = (v24 - v24)
        _ = fi33
    }
    var v34 float64 = ((0.0 * (0.1 - ((1.5 - 1.5) * (1.5 / ((float64(1.5) + (-2.25 * 0.1)) - 0.0))))) + ((1.5 * 0.1) - -2.25))
    v35 := GenS1{F0: 32767, F1: (^(v31))}
    switch ("" + ("\n\t" + (("hello" + "foo bar") + ""))) {
    case "hello":
        var v36 uint64 = 18446744073709551615
        _ = v36
    case "\n\t":
        v37 := GenS1{F0: 1, F1: (v32 % ((v31) | 1))}
        _ = v37
    default:
        v38 := GenS3{F0: -295, F1: false, F2: ""}
        _ = v38
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v31
    _ = v32
    _ = v34
    _ = v35
    _ = v4
}
