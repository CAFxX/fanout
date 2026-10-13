package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string }
type GenS2 struct { F0 uint8 }
func genFn0(p0 int8) int {
    var v7 uint8 = 255
    v8 := GenS1{F0: 2, F1: ""}
    for fi9 := 0; fi9 < 1; fi9++ {
        v10 := GenS1{F0: 65535, F1: ((((("\n\t" + "foo bar") + "foo bar") + "") + "") + ("\n\t" + ("\n\t" + "a")))}
        v11 := []int{(fi9 * fi9), (fi9 & (-(fi9)))}
        v12 := GenS1{F0: 2147483648, F1: "\n\t"}
        _ = fi9
        _ = v10
        _ = v11
        _ = v12
    }
    switch "a" {
    case "":
        var v13 int = 32767
        _ = v13
    }
    _ = p0
    _ = v7
    _ = v8
    return 32767
}

func main() {
    var v1 string = ""
    v2 := []int{-9223372036854775808}
    v3 := make(map[string]uint32)
    var v4 int64 = 1
    var v5 uint8 = 128
    v6 := []string{(("hello" + "a") + v1), v1, v1}
    for fi14 := 0; fi14 < 2; fi14++ {
        var v15 float32 = 0.1
        fmt.Println(fi14, v4, fi14)
        _ = fi14
        _ = v15
    }
    for fi16 := 0; fi16 < 1; fi16++ {
        var v17 uint32 = 2147483647
        for fi18 := 0; fi18 < 4; fi18++ {
            v19 := v3[v1]
            _ = fi18
            _ = v19
        }
        var v20 int8 = 127
        _ = fi16
        _ = v17
        _ = v20
    }
    v21 := GenS2{F0: v5}
    v1 = v1
    var v22 int32 = 694
    var v23 uint16 = 1
    v24 := GenS2{F0: 212}
    switch 127 {
    case -131:
        for fi25 := 0; fi25 < 1; fi25++ {
            for fi26 := 0; fi26 < 15; fi26++ {
                v27 := v3[""]
                _ = fi26
                _ = v27
            }
            v5 = (v5 | v5)
            if (0 <= -128) {
                var v28 bool = (false && (((true || true) || true) || ((-(v4)) >= 1)))
                _ = v28
            }
            _ = fi25
        }
    case 14:
        var v29 float32 = (0.1 - 1e10)
        _ = v29
    }
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
