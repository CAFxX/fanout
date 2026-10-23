package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16; F2 int8 }
func genFn0(p0 uint64, p1 int16, p2 int32) int16 {
    var v7 bool = true
    v8 := []int{351}
    switch 6 {
    case 621:
        if (!v7) {
            if v7 {
                for fi9 := 0; fi9 < 13; fi9++ {
                    var v10 int = 4294967295
                    fmt.Println(p0, p1, v10)
                    _ = fi9
                    _ = v10
                }
            } else {
                v7 = (-109 <= -566)
            }
            p1 = p1
        }
    case 0:
        if (0 > -1) {
            v7 = false
            var v11 int16 = -32768
            fmt.Println(p0, p1, p1)
            _ = v11
        }
    case -2147483648:
        switch ((("hello" + "\n\t") + "a") + (("\n\t" + "") + "a")) {
        case "\n\t":
            var v12 uint32 = 2147483648
            _ = v12
        case "hello":
            var v13 uint64 = p0
            _ = v13
        default:
            var v14 int16 = p1
            _ = v14
        }
    default:
        _ = len(v8)
    }
    for fi15 := 0; fi15 < 18; fi15++ {
        var v16 float64 = 0.0
        _ = fi15
        _ = v16
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return p1
}
func genFn1(p0 uint16) string {
    if (-892 == 9223372036854775807) {
        v17 := GenS1{F0: true, F1: (p0 >> ((2) & 63)), F2: -128}
        _ = v17
    }
    switch 32768 {
    case 1:
        var v18 int32 = -7
        _ = v18
    default:
        for fi19 := 0; fi19 < 16; fi19++ {
            v20 := GenS1{F0: true, F1: 42, F2: -11}
            p0 -= 2
            _ = fi19
            _ = v20
        }
    }
    var v21 int8 = 127
    var v22 uint32 = 2147483647
    _ = p0
    _ = v21
    _ = v22
    return ("a" + ("" + ""))
}

func main() {
    v1 := []string{"", "hello"}
    v2 := make(map[int]int)
    v3 := []string{(("foo bar" + ("hello" + (("foo bar" + "hello") + "\n\t"))) + ""), "\n\t"}
    var v4 bool = (32767 <= 0)
    v5 := make(map[string]string)
    v6 := []uint32{}
    v6 = append(v6, 4294967295)
    var v23 int = 260
    v24 := GenS1{F0: v4, F1: 65535, F2: -33}
    var v25 int64 = -329
    fmt.Println(v23)
    fmt.Println(v23, v4)
    var v26 uint16 = 432
    v26 -= v26
    var v27 uint32 = 892
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
