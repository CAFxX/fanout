package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8; F2 string }
type GenS2 struct { F0 bool; F1 bool; F2 uint8 }
func genFn0() uint32 {
    if (3.14159 >= float64((-2.25 / 1.5))) {
        var v4 string = "\n\t"
        var v5 uint16 = 65535
        var v6 int64 = 42
        _ = v4
        _ = v5
        _ = v6
    }
    if true {
        v7 := []int{}
        var v8 int32 = -204
        _ = v7
        _ = v8
    }
    v9 := []int{}
    switch ("a" + ("\n\t" + "foo bar")) {
    case "\n\t":
        var v10 uint16 = 42
        _ = v10
    case "":
        _ = len(v9)
    default:
        v11 := GenS1{F0: -27, F1: 255, F2: (("foo bar" + "hello") + "")}
        _ = v11
    }
    _ = v9
    return 2147483647
}
func genFn1(p0 int32) bool {
    var v12 bool = (-2.25 != 0.0)
    var v13 float32 = 3.14159
    var v14 uint16 = 723
    switch 42 {
    case 242:
        v14 = (v14 ^ (v14 * 0))
    case 389:
        v15 := GenS2{F0: false, F1: v12, F2: 128}
        _ = v15
    default:
        v16 := []int{}
        _ = v16
    }
    switch -985 {
    case 545:
        var v17 uint8 = 127
        _ = v17
    }
    _ = p0
    _ = v12
    _ = v13
    _ = v14
    return v12
}

func main() {
    var v1 uint8 = 1
    v2 := []uint32{2147483648, 793, 0}
    var v3 uint64 = 42
    v18 := GenS1{F0: 1000, F1: v1, F2: "\n\t"}
    v19 := GenS2{F0: (657 < 2), F1: true, F2: v1}
    var v20 uint16 = 763
    var v21 int16 = 0
    var v22 int16 = (v21 / (((v21 + v21)) | 1))
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v23 uint8 = v1
    var v24 int64 = 224
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
}
