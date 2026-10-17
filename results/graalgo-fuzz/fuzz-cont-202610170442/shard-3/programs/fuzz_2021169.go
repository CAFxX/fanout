package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32 }
type GenS2 struct { F0 int64; F1 uint8; F2 string }

func main() {
    var v1 uint16 = 989
    v2 := make(map[uint32]uint32)
    v3 := make(map[int]uint32)
    var v4 float32 = 0.1
    fmt.Println(v1, v4)
    if false {
        var v5 int32 = -643
        _ = v5
    }
    var v6 int64 = 967
    switch ("" + (("\n\t" + ("foo bar" + ("a" + ("foo bar" + "\n\t")))) + "foo bar")) {
    case "foo bar":
        var v7 uint32 = 42
        _ = v7
    case "\n\t":
        if ((-128 < 1) && (473 >= 0)) {
            switch ("a" + "a") {
            case "foo bar":
                var v8 uint16 = v1
                _ = v8
            case "":
                var v9 uint = 0
                _ = v9
            case "\n\t":
                v10 := GenS1{F0: (("foo bar" + "a") + ("" + "a")), F1: 42}
                _ = v10
            default:
                var v11 float32 = ((float32(1e10) * (v4 * float32(-2.25))) / v4)
                _ = v11
            }
            fmt.Println(v1)
        }
    default:
        v12 := GenS2{F0: 251, F1: 90, F2: "foo bar"}
        _ = v12
    }
    v1 += (v1 / (((^(v1))) | 1))
    fmt.Println(v4)
    v13 := GenS2{F0: -538, F1: 2, F2: ""}
    for fi14 := 0; fi14 < 7; fi14++ {
        var v15 int16 = 0
        _ = fi14
        _ = v15
    }
    fmt.Println(v6)
    fmt.Println(v1)
    v6 += v6
    v4 = 1.5
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v6
}
