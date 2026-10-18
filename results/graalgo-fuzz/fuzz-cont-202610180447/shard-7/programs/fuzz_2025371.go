package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32; F2 bool }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint16 }
func genFn0(p0 int) int32 {
    v6 := GenS1{F0: 4294967295, F1: 1, F2: (!true)}
    var v7 uint16 = 361
    for fi8 := 0; fi8 < 7; fi8++ {
        v9 := GenS1{F0: 1, F1: 552, F2: (false && (true && true))}
        _ = fi8
        _ = v9
    }
    v10 := GenS2{F0: 87}
    _ = p0
    _ = v10
    _ = v6
    _ = v7
    return -577
}
func genFn1(p0 uint16, p1 int) int16 {
    p0 = (p0 / (((^(p0))) | 1))
    var v11 uint16 = (p0 ^ p0)
    p1 -= p1
    _ = p0
    _ = p1
    _ = v11
    return 140
}

func main() {
    var v1 uint8 = 128
    v2 := make(map[uint32]int)
    var v3 uint = 521
    v4 := make(map[string]string)
    var v5 int = 785
    v3 -= v3
    var v12 int8 = 0
    v13 := GenS3{F0: 628}
    switch "hello" {
    case "foo bar":
        v1 = v1
    case "\n\t":
        if false {
            if true {
                v2[4294967295] = (-(v5))
                v14 := v2[0]
                if (!(((("hello" + "\n\t") + "foo bar") > "") || (!(-961 >= 1)))) {
                    var v15 uint64 = 2
                    var v16 int = -120
                    v4[("" + ("foo bar" + "foo bar"))] = ""
                    _ = v15
                    _ = v16
                } else {
                    v3 += 2
                    fmt.Println(v14)
                }
                _ = v14
            }
            v12 -= v12
        }
    default:
        var v17 string = ("" + ("hello" + ""))
        _ = v17
    }
    switch v5 {
    case -1:
        var v18 uint8 = v1
        _ = v18
    }
    var v19 uint64 = 18446744073709551615
    var v20 int32 = genFn0(v5)
    v5 = int(v20)
    fmt.Println(v12)
    var v21 int64 = 162
    v22 := []int{}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
