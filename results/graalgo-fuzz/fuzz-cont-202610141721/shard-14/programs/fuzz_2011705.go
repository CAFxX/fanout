package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 bool; F1 uint32 }
func genFn0(p0 uint32, p1 bool, p2 int8) uint8 {
    var v4 uint = 727
    var v5 float32 = ((0.1 - 1e10) + float32((float64(1e10) / 1.0)))
    var v6 int8 = (p2 - p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    return 55
}
func genFn1(p0 uint16, p1 int16, p2 bool) int64 {
    v7 := GenS1{F0: 2147483647}
    v8 := []int{915, -9223372036854775808}
    for fi9 := 0; fi9 < 1; fi9++ {
        if p2 {
            v8 = append(v8, (fi9 >> ((42) & 63)))
        }
        p0 = 42
        _ = fi9
    }
    var v10 int32 = -1
    p2 = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v7
    _ = v8
    return 42
}

func main() {
    v1 := make(map[int]uint32)
    var v2 string = ""
    var v3 uint = 974
    var v11 int = 0
    if (77 < 42) {
        var v12 string = (v2 + (v2 + ((v2 + "\n\t") + "foo bar")))
        v13 := GenS1{F0: -1}
        v14 := GenS1{F0: 1}
        _ = v12
        _ = v13
        _ = v14
    }
    var v15 float64 = (1.5 / 0.1)
    if ((int64(v11) > genFn1(32767, 126, false)) || true) {
        v16 := GenS1{F0: -751}
        for fi17 := 0; fi17 < 14; fi17++ {
            var v18 float32 = 0.0
            fmt.Println(v18)
            _ = fi17
            _ = v18
        }
        _ = v16
    }
    v19 := GenS2{F0: (("a" + v2) == v2), F1: 42}
    v15 += v15
    switch "hello" {
    case "hello":
        var v20 float32 = 0.1
        _ = v20
    case "foo bar":
        switch v2 {
        case "foo bar":
            v3 += v3
        }
    default:
        if (!true) {
            for fi21 := 0; fi21 < 1; fi21++ {
                var v22 uint32 = 483
                _ = fi21
                _ = v22
            }
        }
    }
    v15 -= ((-2.25 + float64(1.5)) - (v15 - v15))
    var v23 uint16 = 2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v15
    _ = v19
    _ = v2
    _ = v23
    _ = v3
}
