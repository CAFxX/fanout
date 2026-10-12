package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32 }
func genFn0(p0 int64, p1 int8, p2 int32) uint16 {
    if (-829 >= -794) {
        var v7 uint32 = 2147483648
        _ = v7
    }
    v8 := []int{-2147483648, 899, 0, 90}
    p1 += (p1 >> ((567) & 63))
    for fi9 := 0; fi9 < 1; fi9++ {
        if true {
            var v10 float64 = 3.14159
            if true {
                for fi11 := 0; fi11 < 9; fi11++ {
                    var v12 int8 = p1
                    _ = fi11
                    _ = v12
                }
                for fi13 := 0; fi13 < 13; fi13++ {
                    if len(v8) >= 2 {
                        v8 = v8[1:len(v8)]
                    }
                    var v14 int64 = p0
                    var v15 float64 = (float64(float32(v10)) * 1.5)
                    _ = fi13
                    _ = v14
                    _ = v15
                }
                fmt.Println(p0)
            }
            _ = v10
        } else {
            p1 -= (p1 | (p1 & p1))
            var v16 uint64 = 2
            switch fi9 {
            case 2147483647:
                v17 := GenS1{F0: 2, F1: p2}
                _ = v17
            case 9223372036854775807:
                if true {
                    v18 := GenS1{F0: 65535, F1: p2}
                    fmt.Println(fi9, v16, p2)
                    var v19 int16 = -32768
                    _ = v18
                    _ = v19
                }
            }
            _ = v16
        }
        fmt.Println(p1, p0)
        _ = fi9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 480
}

func main() {
    v1 := []int{42, 0, 42, 32768}
    v2 := []int{}
    v3 := []int{}
    var v4 string = ""
    var v5 uint16 = 42
    v6 := []int{9223372036854775807, 1}
    var v20 float64 = (1e10 * float64((0.0 + -2.25)))
    fmt.Println(v5, v20)
    _ = len(v3)
    v5 += 0
    if (!true) {
        var v21 int64 = -785
        v22 := GenS1{F0: v5, F1: 2147483647}
        _ = v21
        _ = v22
    }
    if (2147483647 != 100) {
        switch v4 {
        case "":
            v20 += float64(3.14159)
        case "foo bar":
            if (2 >= 244) {
                for fi23 := 0; fi23 < 12; fi23++ {
                    v20 -= float64((0.1 * 0.0))
                    _ = fi23
                }
            }
        case "a":
            var v24 uint8 = 1
            _ = v24
        default:
            for fi25 := 0; fi25 < 18; fi25++ {
                var v26 int16 = -1
                var v27 uint = 2
                var v28 string = "hello"
                _ = fi25
                _ = v26
                _ = v27
                _ = v28
            }
        }
        fmt.Println(v4, v20, v5)
        v5 += (v5 * v5)
    }
    _ = len(v6)
    var v29 float32 = 0.1
    switch 1 {
    case 531:
        v30 := GenS1{F0: (^(v5)), F1: 303}
        _ = v30
    case 2147483647:
        fmt.Println(v20, v5)
    default:
        fmt.Println(v20, v20)
    }
    var v31 string = "\n\t"
    v29 += v29
    if ((false && (true || true)) || (!(false || (!true)))) {
        var v32 float32 = (1e10 - 0.1)
        _ = v32
    } else {
        _ = len(v2)
        for fi33 := 0; fi33 < 1; fi33++ {
            var v34 int32 = 0
            _ = fi33
            _ = v34
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v29
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
