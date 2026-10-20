package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8; F2 uint32 }

func main() {
    v1 := make(map[uint32]string)
    var v2 float32 = 0.1
    v3 := []string{("a" + ("\n\t" + "a"))}
    v4 := make(map[int]bool)
    var v5 uint16 = 1
    var v6 int = 0
    v6 = 83
    var v7 uint32 = 42
    v7 += (v7 % (((v7 - 65535)) | 1))
    for fi8 := 0; fi8 < 7; fi8++ {
        var v9 bool = false
        var v10 int32 = 0
        for fi11 := 0; fi11 < 10; fi11++ {
            v9 = false
            var v12 uint8 = 128
            var v13 float64 = (0.1 + (1e10 / 1.0))
            _ = fi11
            _ = v12
            _ = v13
        }
        _ = fi8
        _ = v10
        _ = v9
    }
    if len(v3) > 0 {
        v3[0] = ("a" + (("foo bar" + (("foo bar" + (("\n\t" + "foo bar") + ("a" + ("foo bar" + "")))) + "")) + "hello"))
    }
    if (128 >= 0) {
        var v14 int8 = 0
        _ = len(v3)
        _ = v14
    }
    v5 = 740
    var v15 int = (-(v6))
    if (-788 >= 1000) {
        v16 := GenS1{F0: 1, F1: 255, F2: v7}
        _ = v16
    }
    var v17 float64 = 3.14159
    switch "foo bar" {
    case "hello":
        _ = len(v3)
    default:
        if true {
            v18 := GenS1{F0: -1, F1: 128, F2: 14}
            _ = v18
        } else {
            v6 -= -9223372036854775808
            v19 := GenS1{F0: 1, F1: 0, F2: (v7 % ((v7) | 1))}
            _ = v19
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
