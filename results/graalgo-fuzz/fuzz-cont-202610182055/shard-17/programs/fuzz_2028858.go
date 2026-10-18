package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint8; F1 uint; F2 uint32 }
type GenS3 struct { F0 uint32; F1 string }
func genFn0(p0 int16) int {
    var v6 uint64 = 18446744073709551615
    var v7 int64 = -9223372036854775808
    _ = p0
    _ = v6
    _ = v7
    return 4294967295
}
func genFn1() uint16 {
    switch ("a" + "\n\t") {
    case "":
        var v8 uint32 = 1
        _ = v8
    case "a":
        if (-403 >= -150) {
            var v9 uint16 = 739
            _ = v9
        } else {
            if false {
                var v10 int8 = -128
                _ = v10
            }
            v11 := []int{}
            v11 = append(v11, 127)
            _ = v11
        }
    default:
        var v12 uint8 = 2
        _ = v12
    }
    var v13 uint8 = 96
    switch ("a" + "a") {
    case "a":
        if (1000 <= 0) {
            var v14 float32 = 0.0
            var v15 int16 = -876
            v16 := GenS1{F0: false}
            _ = v14
            _ = v15
            _ = v16
        }
    case "foo bar":
        v17 := GenS2{F0: 50, F1: 0, F2: 2147483648}
        _ = v17
    }
    v18 := []int{}
    fmt.Println(v13)
    _ = v13
    _ = v18
    return 1
}

func main() {
    var v1 uint16 = 32767
    v2 := []string{}
    var v3 int8 = 0
    v4 := []string{"a", "", "foo bar"}
    var v5 float32 = (1.5 / -2.25)
    v5 = v5
    for fi19 := 0; fi19 < 16; fi19++ {
        v20 := GenS2{F0: 0, F1: 65535, F2: 922}
        _ = fi19
        _ = v20
    }
    v21 := GenS2{F0: 1, F1: 2, F2: 1}
    var v22 uint16 = v1
    for fi23 := 0; fi23 < 4; fi23++ {
        var v24 uint32 = 942
        fmt.Println(fi23, v1, v22)
        _ = fi23
        _ = v24
    }
    v22 -= genFn1()
    v3 -= (v3 * (v3 - -128))
    var v25 int = 128
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v26 := GenS3{F0: 1, F1: ("\n\t" + "foo bar")}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
