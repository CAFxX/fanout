package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint }
type GenS2 struct { F0 int; F1 uint32; F2 uint32 }
type GenS3 struct { F0 uint64; F1 uint16 }

func main() {
    var v1 float64 = 1e10
    var v2 string = "\n\t"
    v3 := []string{(v2 + ("\n\t" + "foo bar")), "hello", ((v2 + ("foo bar" + (("hello" + ("\n\t" + (v2 + (("\n\t" + v2) + (v2 + "foo bar"))))) + ("" + v2)))) + v2), ("foo bar" + "")}
    v4 := []string{}
    var v5 uint64 = 18446744073709551615
    var v6 string = "foo bar"
    if true {
        var v7 int32 = -578
        var v8 float32 = 3.14159
        var v9 int8 = -128
        _ = v7
        _ = v8
        _ = v9
    }
    fmt.Println(v5, v6)
    if len(v3) > 0 {
        v3[0] = v2
    }
    v10 := GenS1{F0: -123, F1: 31}
    v11 := GenS1{F0: 1, F1: 1}
    v12 := GenS2{F0: 1, F1: 4294967295, F2: 4294967295}
    v13 := GenS3{F0: (v5 - (v5 | v5)), F1: 255}
    for fi14 := 0; fi14 < 5; fi14++ {
        var v15 float32 = 1.5
        var v16 uint32 = 2147483647
        _ = fi14
        _ = v15
        _ = v16
    }
    v17 := GenS3{F0: (v5 ^ v5), F1: uint16(v5)}
    var v18 uint = 65535
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
