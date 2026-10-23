package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int8; F1 int16; F2 int64 }
func genFn0() string {
    for fi5 := 0; fi5 < 11; fi5++ {
        var v6 uint8 = 0
        _ = fi5
        _ = v6
    }
    v7 := GenS1{F0: 127, F1: -32768}
    v8 := GenS1{F0: -119, F1: 173}
    var v9 uint16 = 255
    _ = v7
    _ = v8
    _ = v9
    return "foo bar"
}
func genFn1(p0 uint16, p1 uint8, p2 uint16) uint {
    switch ("\n\t" + "hello") {
    case "":
        var v10 uint32 = 585
        _ = v10
    case "hello":
        switch ("a" + "hello") {
        case "foo bar":
            var v11 int8 = 0
            _ = v11
        default:
            var v12 uint16 = p2
            _ = v12
        }
    default:
        p0 -= p0
    }
    v13 := []int{4294967295, -204, int(p1), 32767}
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return 42
}

func main() {
    v1 := []int{-490, 1, 0}
    var v2 int8 = -82
    var v3 string = ("hello" + "\n\t")
    var v4 uint8 = 142
    fmt.Println(v2, v2)
    fmt.Println(v3)
    v1 = append(v1, -9223372036854775808)
    var v14 string = (("hello" + (v3 + (v3 + "\n\t"))) + (v3 + (v3 + v3)))
    for fi15 := 0; fi15 < 15; fi15++ {
        var v16 float64 = 0.0
        v17 := GenS1{F0: v2, F1: 1000}
        v18 := GenS2{F0: (v4 | 255)}
        _ = fi15
        _ = v16
        _ = v17
        _ = v18
    }
    v1 = append(v1, 959)
    v4 = v4
    for fi19 := 0; fi19 < 5; fi19++ {
        v20 := GenS2{F0: v4}
        v1 = append(v1, fi19)
        _ = fi19
        _ = v20
    }
    var v21 int64 = -416
    if false {
        if ((0.0 - float32(((-2.25 / 1.0) - (0.0 - 1e10)))) > (0.1 - -2.25)) {
            v2 += v2
        }
        var v22 int32 = 42
        for fi23 := 0; fi23 < 11; fi23++ {
            fmt.Println(v4, v14)
            v24 := GenS1{F0: v2, F1: 1}
            if false {
                v25 := GenS3{F0: -69, F1: 1, F2: (v21 & v21)}
                _ = v25
            }
            _ = fi23
            _ = v24
        }
        _ = v22
    } else {
        v26 := GenS3{F0: 88, F1: 0, F2: v21}
        var v27 int16 = -32768
        _ = v26
        _ = v27
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v2
    _ = v21
    _ = v3
    _ = v4
}
