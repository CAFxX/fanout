package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint; F2 int64 }
type GenS2 struct { F0 int16; F1 int16; F2 int }
func genFn0(p0 uint8, p1 string) int64 {
    v4 := GenS1{F0: 1, F1: 18446744073709551615, F2: -52}
    for fi5 := 0; fi5 < 19; fi5++ {
        for fi6 := 0; fi6 < 17; fi6++ {
            v7 := []int{127}
            _ = fi6
            _ = v7
        }
        v8 := []int{-469, fi5}
        _ = len(v8)
        _ = fi5
        _ = v8
    }
    v9 := GenS1{F0: -480, F1: 1, F2: 42}
    fmt.Println(p0, p1)
    _ = p0
    _ = p1
    _ = v4
    _ = v9
    return -1
}

func main() {
    var v1 int = -659
    v2 := []int{}
    var v3 uint16 = 0
    v1 += v1
    switch (("" + ("hello" + ("foo bar" + (("hello" + "\n\t") + (("hello" + "foo bar") + "foo bar"))))) + "hello") {
    case "foo bar":
        var v10 uint16 = 593
        _ = v10
    case "\n\t":
        var v11 int16 = 42
        _ = v11
    }
    v12 := GenS1{F0: 32767, F1: 0, F2: 0}
    var v13 int64 = -9223372036854775808
    var v14 string = "a"
    for fi15 := 0; fi15 < 6; fi15++ {
        v1 = (fi15 ^ fi15)
        v14 = v14
        _ = fi15
    }
    fmt.Println(v14, v1)
    v16 := GenS2{F0: -1, F1: 1, F2: (v1 / (((v1 + (v1 - v1))) | 1))}
    v2 = append(v2, (v1 | (v1 / ((v1) | 1))))
    for fi17 := 0; fi17 < 14; fi17++ {
        var v18 int = fi17
        var v19 int32 = 0
        for fi20 := 0; fi20 < 8; fi20++ {
            fmt.Println(fi17, v13, v18)
            _ = fi20
        }
        _ = fi17
        _ = v18
        _ = v19
    }
    v1 = (v1 / (((v1 + (v1 * v1))) | 1))
    _ = len(v2)
    var v21 float64 = 0.0
    var v22 int16 = int16(v3)
    v23 := GenS2{F0: (v22 | -857), F1: (v22 ^ v22), F2: (v1 ^ -9223372036854775808)}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v16
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
}
