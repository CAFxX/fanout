package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint; F2 uint64 }
type GenS2 struct { F0 uint32; F1 int }
func genFn0(p0 int32, p1 int32, p2 uint16) int16 {
    var v5 uint16 = (p2 - (p2 & p2))
    for fi6 := 0; fi6 < 12; fi6++ {
        for fi7 := 0; fi7 < 1; fi7++ {
            switch fi7 {
            case -264:
                var v8 int32 = -197
                _ = v8
            case -1:
                if (2 <= 2147483648) {
                    var v9 uint16 = p2
                    fmt.Println(v9, fi6, v5)
                    _ = v9
                } else {
                    var v10 string = ("a" + (("foo bar" + "foo bar") + "\n\t"))
                    fmt.Println(v5, v10, v10)
                    p1 -= 1
                    _ = v10
                }
            case -749:
                v11 := GenS1{F0: p1, F1: 2, F2: 4294967295}
                _ = v11
            }
            _ = fi7
        }
        v12 := GenS2{F0: 2, F1: fi6}
        _ = fi6
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return -32768
}
func genFn1(p0 uint, p1 int8, p2 int8) uint16 {
    var v13 int16 = genFn0(-2147483648, -2147483648, 927)
    var v14 uint64 = 4294967295
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v14
    return 42
}

func main() {
    var v1 int16 = -1
    v2 := make(map[string]bool)
    v3 := []string{"foo bar"}
    var v4 uint8 = 42
    v4 = v4
    switch "a" {
    case "":
        v1 = genFn0(1, -1, 0)
    }
    var v15 bool = true
    var v16 float64 = (((0.0 + float64(-2.25)) * 0.0) + 0.1)
    v16 += (-2.25 + (float64(1e10) - 1e10))
    var v17 int16 = genFn0(42, 186, 65535)
    if v15 {
        var v18 float32 = 3.14159
        _ = v18
    } else {
        delete(v2, "hello")
    }
    if (-240 == 0) {
        var v19 int = -333
        var v20 int32 = -2147483648
        v15 = false
        _ = v19
        _ = v20
    } else {
        if v15 {
            var v21 float64 = ((3.14159 - float64(-2.25)) + 3.14159)
            _ = v21
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
