package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool; F2 int64 }
type GenS2 struct { F0 uint16; F1 uint16; F2 int }
type GenS3 struct { F0 uint; F1 uint; F2 int8 }
type GenS4 struct { F0 uint16 }
func genFn0(p0 bool) uint {
    switch 127 {
    case -481:
        switch 0 {
        case -1:
            v5 := GenS1{F0: 77, F1: p0, F2: 608}
            _ = v5
        case -2147483648:
            for fi6 := 0; fi6 < 6; fi6++ {
                v7 := GenS2{F0: 361, F1: 255, F2: -1}
                switch "hello" {
                case "hello":
                    v8 := []int{fi6, fi6, (fi6 ^ 32768), fi6}
                    _ = v8
                }
                p0 = true
                _ = fi6
                _ = v7
            }
        case 32768:
            p0 = false
        default:
            p0 = p0
        }
    case 127:
        var v9 string = ""
        _ = v9
    default:
        var v10 bool = (!(false || true))
        _ = v10
    }
    fmt.Println(p0)
    _ = p0
    return 636
}

func main() {
    var v1 uint64 = 0
    v2 := make(map[int]bool)
    var v3 bool = (true && false)
    v4 := make(map[string]string)
    v1 = v1
    fmt.Println(v3, v1)
    v11 := v4["foo bar"]
    var v12 string = ("hello" + ((v11 + (v11 + "hello")) + (v11 + (("hello" + "hello") + v11))))
    for fi13 := 0; fi13 < 5; fi13++ {
        v14 := GenS1{F0: 0, F1: false, F2: 922}
        v15 := GenS3{F0: 42, F1: 285, F2: -128}
        _ = fi13
        _ = v14
        _ = v15
    }
    if false {
        v16 := GenS1{F0: 1, F1: v3, F2: -9223372036854775808}
        _ = v16
    }
    if false {
        var v17 uint64 = v1
        var v18 uint32 = 265
        var v19 uint8 = 42
        _ = v17
        _ = v18
        _ = v19
    } else {
        var v20 uint32 = 929
        var v21 float64 = (((float64(float32((3.14159 + (-2.25 + 0.0)))) + (-2.25 / ((((1.5 - ((1.5 - (float64(float32(-2.25)) - float64(-2.25))) + 0.1)) / 1.5) + -2.25) + 0.1))) / 1.0) * (-2.25 + 1.5))
        _ = v20
        _ = v21
    }
    v22 := GenS3{F0: 18446744073709551615, F1: 263, F2: 89}
    v23 := GenS4{F0: 2}
    v24 := []int{169}
    var v25 float64 = 0.0
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
}
