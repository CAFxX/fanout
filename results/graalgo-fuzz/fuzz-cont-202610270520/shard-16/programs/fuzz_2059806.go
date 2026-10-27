package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint16; F1 int64 }
type GenS3 struct { F0 string; F1 uint8 }
type GenS4 struct { F0 uint8; F1 uint8 }
func genFn0(p0 int16, p1 int32) int64 {
    fmt.Println(p0, p0)
    p1 = (p1 ^ 0)
    p0 = p0
    v7 := GenS1{F0: 0}
    v8 := []int{42}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 42
}

func main() {
    var v1 int8 = 111
    v2 := []int{}
    v3 := make(map[uint32]bool)
    var v4 float64 = 1.5
    v5 := []int{-50}
    v6 := make(map[string]uint32)
    if false {
        var v9 uint64 = 0
        _ = v9
    } else {
        v4 = (v4 / v4)
        switch 2147483647 {
        case 0:
            fmt.Println(v4, v1)
        default:
            v10 := GenS2{F0: 991, F1: 359}
            _ = v10
        }
        switch ("foo bar" + ("a" + ("" + ("foo bar" + ("foo bar" + "foo bar"))))) {
        case "\n\t":
            for fi11 := 0; fi11 < 5; fi11++ {
                var v12 uint8 = 225
                delete(v6, "hello")
                _ = fi11
                _ = v12
            }
        case "foo bar":
            var v13 int64 = 9223372036854775807
            _ = v13
        case "hello":
            var v14 uint8 = 2
            _ = v14
        }
    }
    var v15 float32 = (((0.1 * 1.5) / ((((1.5 - (float32(3.14159) * 1.5)) * 0.0) + ((float32((1e10 / v4)) - ((-2.25 - 0.0) + (1e10 - (((float32(v4) - 1.5) / 1e10) / 1.5)))) / float32(float64(3.14159)))) * (1.5 + (0.1 / 1e10)))) + 1.5)
    _ = len(v3)
    switch -2147483648 {
    case 42:
        var v16 int = 2147483647
        _ = v16
    case -405:
        v1 += v1
    case 441:
        if ((v15 / float32((v4 - 0.1))) < (1.5 * v15)) {
            for fi17 := 0; fi17 < 4; fi17++ {
                v18 := GenS2{F0: 1, F1: 463}
                v19 := GenS2{F0: 65535, F1: 9223372036854775807}
                for fi20 := 0; fi20 < 1; fi20++ {
                    _ = len(v5)
                    var v21 uint = 42
                    switch (fi17 / ((fi20) | 1)) {
                    case -9223372036854775808:
                        v21 -= 255
                    case -477:
                        v4 -= 1e10
                    case 4294967295:
                        v4 = ((1e10 * -2.25) * (v4 - (v4 + ((v4 - (v4 / ((float64(v15) + v4) - v4))) + v4))))
                    default:
                        delete(v3, 0)
                    }
                    _ = fi20
                    _ = v21
                }
                _ = fi17
                _ = v18
                _ = v19
            }
            for fi22 := 0; fi22 < 7; fi22++ {
                if true {
                    var v23 int = fi22
                    var v24 uint16 = 752
                    _ = v23
                    _ = v24
                }
                _ = fi22
            }
            v1 -= 127
        }
    default:
        v25 := GenS3{F0: ("\n\t" + ("foo bar" + (("foo bar" + ("foo bar" + "hello")) + "\n\t"))), F1: 127}
        _ = v25
    }
    _ = len(v6)
    v1 = v1
    var v26 int64 = genFn0(-408, 2147483647)
    for fi27 := 0; fi27 < 10; fi27++ {
        v26 = (v26 / ((v26) | 1))
        v28 := GenS4{F0: 177, F1: 0}
        if (false || (!true)) {
            var v29 uint32 = 169
            _ = v29
        }
        _ = fi27
        _ = v28
    }
    fmt.Println(v26)
    var v30 bool = false
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v15
    _ = v2
    _ = v26
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
