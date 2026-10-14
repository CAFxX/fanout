package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8 }
type GenS2 struct { F0 uint8; F1 int }
func genFn0(p0 uint8) int64 {
    var v7 uint8 = (p0 + (p0 * p0))
    var v8 float32 = -2.25
    switch "hello" {
    case "a":
        if false {
            for fi9 := 0; fi9 < 19; fi9++ {
                v10 := GenS1{F0: true, F1: 1}
                var v11 uint = 21
                _ = fi9
                _ = v10
                _ = v11
            }
            fmt.Println(v8, v8, v8)
            switch ("foo bar" + "hello") {
            case "\n\t":
                v12 := []int{-149, 1}
                _ = v12
            case "hello":
                for fi13 := 0; fi13 < 1; fi13++ {
                    var v14 int64 = 1
                    _ = fi13
                    _ = v14
                }
            }
        } else {
            v8 = v8
        }
    case "\n\t":
        fmt.Println(v8, p0)
    default:
        v15 := GenS1{F0: (v8 <= 3.14159), F1: v7}
        _ = v15
    }
    if false {
        v16 := []int{438}
        _ = v16
    }
    _ = p0
    _ = v7
    _ = v8
    return 273
}
func genFn1(p0 int16, p1 uint8) int32 {
    fmt.Println(p1, p0)
    v17 := GenS1{F0: true, F1: p1}
    var v18 int32 = -331
    _ = p0
    _ = p1
    _ = v17
    _ = v18
    return 1
}

func main() {
    v1 := []int{-716, 128}
    var v2 uint = 4294967295
    v3 := []string{}
    var v4 int16 = 1000
    var v5 int32 = 2147483647
    var v6 bool = (!true)
    if false {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        v6 = (((v2 % ((v2) | 1)) == v2) && (758 <= 4294967295))
    } else {
        v5 = v5
    }
    v19 := GenS1{F0: ((0.1 + 1.5) > 0.0), F1: 127}
    if true {
        switch 128 {
        case 1:
            var v20 uint = (v2 ^ (^(v2)))
            _ = v20
        case 42:
            var v21 int8 = -51
            _ = v21
        case 32768:
            fmt.Println(v6, v2, v4)
        default:
            var v22 uint32 = 42
            _ = v22
        }
    }
    var v23 uint64 = 18446744073709551615
    v24 := GenS2{F0: 42, F1: int(v4)}
    var v25 uint64 = v23
    for fi26 := 0; fi26 < 1; fi26++ {
        switch fi26 {
        case -2147483648:
            if (255 == 255) {
                v27 := GenS1{F0: (!(!true)), F1: 2}
                var v28 int64 = -1
                fmt.Println(v4)
                _ = v27
                _ = v28
            } else {
                v29 := GenS2{F0: 42, F1: (-(fi26))}
                fmt.Println(v6, v6)
                var v30 uint32 = 0
                _ = v29
                _ = v30
            }
        }
        _ = fi26
    }
    var v31 uint32 = 1
    v3 = append(v3, ("" + ("hello" + "foo bar")))
    v1 = append(v1, 127)
    v23 = (v25 & uint64(v31))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
