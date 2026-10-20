package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string }
func genFn0() uint {
    v4 := []int{-964, 1, 127, 9223372036854775807}
    for fi5 := 0; fi5 < 6; fi5++ {
        v6 := GenS1{F0: 2, F1: "\n\t"}
        _ = fi5
        _ = v6
    }
    var v7 uint8 = 255
    _ = v4
    _ = v7
    return 42
}
func genFn1(p0 int8, p1 int32, p2 int) string {
    if false {
        p0 -= p0
        var v8 uint64 = 141
        _ = v8
    }
    p1 = p1
    _ = p0
    _ = p1
    _ = p2
    return "hello"
}

func main() {
    var v1 float32 = 1e10
    var v2 float32 = 1.5
    var v3 int64 = 0
    for fi9 := 0; fi9 < 18; fi9++ {
        v10 := []int{453, fi9, (fi9 * (-(fi9)))}
        _ = fi9
        _ = v10
    }
    if true {
        if true {
            fmt.Println(v1, v2, v2)
            fmt.Println(v2, v1)
        } else {
            if true {
                switch "" {
                case "a":
                    v1 = v1
                case "\n\t":
                    var v11 float64 = float64(v2)
                    _ = v11
                case "":
                    if true {
                        var v12 int32 = -1
                        _ = v12
                    }
                default:
                    v13 := []int{}
                    _ = v13
                }
                var v14 int16 = 42
                fmt.Println(v2)
                _ = v14
            }
        }
        var v15 int8 = 1
        for fi16 := 0; fi16 < 1; fi16++ {
            var v17 bool = ((255 <= 65535) || true)
            var v18 uint8 = 127
            _ = fi16
            _ = v17
            _ = v18
        }
        _ = v15
    } else {
        v1 = (1.5 - -2.25)
        v3 -= v3
    }
    v19 := []int{int(v3), -2147483648, 2147483647}
    if (713 <= 4294967295) {
        v3 = (v3 | -477)
    } else {
        v20 := GenS1{F0: 367, F1: ("a" + "\n\t")}
        var v21 int8 = -37
        _ = v20
        _ = v21
    }
    v22 := GenS1{F0: 4294967295, F1: ("" + ("" + "\n\t"))}
    fmt.Println(v1, v2)
    fmt.Println(v2)
    v23 := GenS1{F0: 2, F1: "foo bar"}
    v24 := GenS1{F0: 18446744073709551615, F1: "\n\t"}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
}
