package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8 }
type GenS2 struct { F0 uint; F1 uint64 }
func genFn0(p0 uint) uint8 {
    var v4 bool = (!true)
    v5 := GenS1{F0: -651, F1: 87}
    _ = p0
    _ = v4
    _ = v5
    return 59
}

func main() {
    v1 := []string{"", "hello"}
    var v2 uint32 = 1
    v3 := make(map[string]uint32)
    var v6 uint8 = genFn0(42)
    if false {
        var v7 int64 = 0
        if len(v1) > 0 {
            v1[0] = ("\n\t" + "")
        }
        if ("foo bar" < "") {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            v2 -= (^(v2))
            v8 := GenS1{F0: 1, F1: v6}
            _ = v8
        }
        _ = v7
    }
    switch "\n\t" {
    case "":
        for fi9 := 0; fi9 < 18; fi9++ {
            fmt.Println(v6, v2)
            if (255 <= (v6 / (((^(v6))) | 1))) {
                _ = len(v1)
            } else {
                _ = len(v1)
            }
            var v10 float64 = (0.0 * ((1e10 / 1.0) + (float64(float32((1e10 + 0.0))) - 0.1)))
            _ = fi9
            _ = v10
        }
    default:
        v6 = (^(v6))
    }
    var v11 uint = 4294967295
    switch (("foo bar" + "foo bar") + "foo bar") {
    case "hello":
        v11 = 4294967295
    }
    if false {
        var v12 uint16 = 32767
        _ = v12
    } else {
        var v13 uint8 = v6
        for fi14 := 0; fi14 < 1; fi14++ {
            var v15 int = fi14
            v13 = (v13 * v6)
            _ = fi14
            _ = v15
        }
        _ = v13
    }
    v16 := GenS2{F0: 255, F1: 2}
    var v17 uint16 = 255
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v6
}
