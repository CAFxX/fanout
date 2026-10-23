package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32; F2 uint32 }
type GenS2 struct { F0 uint; F1 bool; F2 int16 }
type GenS3 struct { F0 uint8 }
func genFn0(p0 int64, p1 int8, p2 uint) int {
    fmt.Println(p1, p0)
    if (1 > 0) {
        p2 += p2
    } else {
        p2 = 255
        v6 := GenS1{F0: 65535, F1: -249, F2: 256}
        _ = v6
    }
    for fi7 := 0; fi7 < 17; fi7++ {
        var v8 float64 = 3.14159
        var v9 uint = p2
        _ = fi7
        _ = v8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    return -9223372036854775808
}
func genFn1(p0 uint) uint8 {
    v10 := []int{169, -32, -175}
    var v11 bool = (-128 <= 1)
    if len(v10) > 0 {
        v10[0] = 4294967295
    }
    _ = p0
    _ = v10
    _ = v11
    return 28
}

func main() {
    var v1 int = -9223372036854775808
    var v2 float32 = 0.1
    var v3 bool = (42 <= -34)
    var v4 int16 = 498
    var v5 uint = 4294967295
    v5 = v5
    v12 := []int{-1}
    switch ("hello" + "\n\t") {
    case "hello":
        switch (v1 - 1) {
        case 9223372036854775807:
            for fi13 := 0; fi13 < 16; fi13++ {
                switch "" {
                case "\n\t":
                    v4 = 1
                case "foo bar":
                    var v14 float32 = float32(3.14159)
                    _ = v14
                default:
                    for fi15 := 0; fi15 < 10; fi15++ {
                        v16 := GenS2{F0: 42, F1: (((578 < 952) && (65535 == 629)) && (((("" + "a") + "hello") + "") <= "hello")), F2: (v4 ^ (-(v4)))}
                        _ = fi15
                        _ = v16
                    }
                }
                v2 = v2
                _ = fi13
            }
        }
    case "foo bar":
        v17 := GenS3{F0: 127}
        _ = v17
    case "":
        fmt.Println(v4, v4)
    }
    v1 = v1
    var v18 bool = ((v4 / ((1000) | 1)) >= v4)
    var v19 int32 = 1
    var v20 int32 = int32(v1)
    if v18 {
        v4 += (-(v4))
    }
    v1 += 404
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
