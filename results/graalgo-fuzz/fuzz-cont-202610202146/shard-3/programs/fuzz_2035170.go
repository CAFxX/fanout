package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
func genFn0(p0 int8, p1 string) uint {
    v7 := GenS1{F0: 663, F1: 128}
    for fi8 := 0; fi8 < 18; fi8++ {
        fmt.Println(fi8, fi8, p1)
        _ = fi8
    }
    fmt.Println(p1, p1)
    p0 = -128
    _ = p0
    _ = p1
    _ = v7
    return 65535
}

func main() {
    var v1 float32 = -2.25
    v2 := make(map[uint32]uint32)
    v3 := make(map[int]uint32)
    var v4 uint8 = 42
    var v5 float32 = (3.14159 - (1.5 + v1))
    v6 := make(map[string]string)
    switch 127 {
    case 9223372036854775807:
        if false {
            var v9 uint16 = 2
            fmt.Println(v4, v9)
            v2[1] = 65535
            _ = v9
        } else {
            v4 -= 1
            v10 := GenS1{F0: 83, F1: v4}
            _ = v10
        }
    case 0:
        delete(v3, 127)
    case -2147483648:
        v6["foo bar"] = "hello"
    default:
        switch -9223372036854775808 {
        case -1:
            v11 := []int{1, 128, 2147483647, -99}
            _ = v11
        case 1:
            if false {
                var v12 uint64 = 493
                v1 += -2.25
                _ = v12
            } else {
                var v13 int32 = 842
                var v14 int = 406
                switch "\n\t" {
                case "":
                    fmt.Println(v5)
                case "foo bar":
                    if true {
                        v15 := GenS1{F0: 0, F1: 1}
                        v16 := v3[v14]
                        _ = v15
                        _ = v16
                    } else {
                        v17 := GenS1{F0: 18446744073709551615, F1: (v4 | 196)}
                        var v18 uint = 2
                        v18 -= genFn0(100, "")
                        _ = v17
                        _ = v18
                    }
                }
                _ = v13
                _ = v14
            }
        default:
            switch 0 {
            case -2147483648:
                fmt.Println(v5, v4, v1)
            case 2147483647:
                var v19 int16 = 917
                _ = v19
            }
        }
    }
    fmt.Println(v1, v4)
    delete(v2, 4294967295)
    v4 += 127
    switch ((("hello" + ("\n\t" + "hello")) + ("foo bar" + (("\n\t" + ("\n\t" + ("" + ("" + "\n\t")))) + ""))) + "\n\t") {
    case "foo bar":
        var v20 int = 42
        _ = v20
    case "":
        v4 += v4
    case "\n\t":
        _ = len(v2)
    }
    v21 := v3[127]
    var v22 float64 = float64(((float32(1.5) * (0.0 * (v1 + (float32(0.0) * v1)))) * 0.1))
    v22 = 1.5
    v4 = v4
    v23 := GenS1{F0: 42, F1: (v4 / ((v4) | 1))}
    v24 := []int{}
    var v25 uint16 = 42
    v25 = v25
    var v26 int16 = -1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
