package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16 }
type GenS2 struct { F0 int16; F1 uint8 }
type GenS3 struct { F0 uint8 }
func genFn0(p0 uint8, p1 uint64, p2 int) int32 {
    p1 = p1
    var v7 int16 = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 2147483647
}
func genFn1(p0 int32) uint64 {
    var v8 float64 = 0.0
    var v9 uint = 42
    if true {
        var v10 int8 = int8(v9)
        switch ((("" + "hello") + "foo bar") + "hello") {
        case "\n\t":
            var v11 int16 = 0
            _ = v11
        case "hello":
            fmt.Println(v9, v10)
        case "a":
            v9 -= 18446744073709551615
        default:
            var v12 int32 = (p0 << ((v9) & 63))
            _ = v12
        }
        for fi13 := 0; fi13 < 17; fi13++ {
            v14 := []int{}
            _ = fi13
            _ = v14
        }
        _ = v10
    }
    _ = p0
    _ = v8
    _ = v9
    return 42
}

func main() {
    var v1 float32 = float32(-2.25)
    var v2 uint64 = 1
    var v3 int = -692
    var v4 string = ("a" + ("" + (("hello" + "foo bar") + ("\n\t" + "a"))))
    var v5 float64 = float64(v1)
    var v6 uint64 = (v2 - (v2 << ((1) & 63)))
    if ((!(!false)) && (42 <= 129)) {
        var v15 int8 = -1
        var v16 int16 = 333
        _ = v15
        _ = v16
    }
    v17 := []int{(v3 & (-(v3))), (v3 + v3), (v3 - -759)}
    var v18 int32 = 673
    for fi19 := 0; fi19 < 11; fi19++ {
        v20 := GenS1{F0: -25, F1: -1}
        fmt.Println(v4, v5)
        _ = fi19
        _ = v20
    }
    v21 := GenS1{F0: 55, F1: 976}
    v6 -= genFn1((v18 + (v18 & v18)))
    for fi22 := 0; fi22 < 20; fi22++ {
        var v23 int64 = 1
        _ = fi22
        _ = v23
    }
    for fi24 := 0; fi24 < 5; fi24++ {
        fmt.Println(v4)
        for fi25 := 0; fi25 < 11; fi25++ {
            v26 := GenS2{F0: -1, F1: 127}
            _ = fi25
            _ = v26
        }
        _ = fi24
    }
    v27 := GenS2{F0: -1, F1: 127}
    v28 := GenS3{F0: 128}
    fmt.Println(v4)
    v18 -= v18
    switch v3 {
    case -1:
        v18 += (v18 << ((42) & 63))
    }
    v5 = v5
    v17 = append(v17, v3)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
