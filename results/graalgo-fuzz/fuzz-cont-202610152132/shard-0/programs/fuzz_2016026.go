package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 int16 }
type GenS2 struct { F0 int16; F1 uint64 }
func genFn0(p0 int64) uint {
    var v7 float64 = (1e10 - 1e10)
    var v8 uint16 = 255
    var v9 string = "foo bar"
    v10 := []int{-9223372036854775808, 128, 32767}
    switch -159 {
    case 109:
        var v11 int64 = (p0 * -668)
        _ = v11
    case -1:
        if false {
            if (45 != 127) {
                fmt.Println(v8)
            } else {
                v7 = v7
                v10 = append(v10, 4294967295)
            }
        }
    }
    _ = p0
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 2
}

func main() {
    v1 := []int{32767, 127}
    var v2 int16 = 0
    v3 := make(map[string]uint32)
    v4 := []int{-2147483648, 127, -582}
    v5 := []string{"hello", ("hello" + ("hello" + "\n\t")), (("\n\t" + ("" + "hello")) + "")}
    v6 := []int{-409, -1, -9223372036854775808, 582}
    v2 = (^(v2))
    switch "\n\t" {
    case "\n\t":
        var v12 float32 = ((-2.25 / ((1.5 - 0.1) - float32((1e10 * float64((0.1 - 1e10)))))) * (1.5 - 0.1))
        _ = v12
    case "":
        v13 := GenS1{F0: 38, F1: 1, F2: 0}
        _ = v13
    case "foo bar":
        fmt.Println(v2)
    }
    fmt.Println(v2)
    var v14 int8 = -128
    if true {
        var v15 uint64 = 204
        v16 := GenS2{F0: int16(v14), F1: 821}
        var v17 int = -135
        _ = v15
        _ = v16
        _ = v17
    }
    if len(v1) > 0 {
        v1[0] = -1
    }
    v18 := GenS2{F0: v2, F1: 161}
    v19 := GenS2{F0: (v2 / (((v2 ^ 32767)) | 1)), F1: 0}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
