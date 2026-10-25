package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
type GenS2 struct { F0 string }
func genFn0() uint {
    var v4 int = -9223372036854775808
    if true {
        v4 = v4
        for fi5 := 0; fi5 < 20; fi5++ {
            switch "" {
            case "hello":
                for fi6 := 0; fi6 < 4; fi6++ {
                    fmt.Println(fi5, fi5)
                    var v7 int64 = 1
                    _ = fi6
                    _ = v7
                }
            default:
                v8 := []int{-220, -1}
                _ = v8
            }
            _ = fi5
        }
    }
    _ = v4
    return 2
}
func genFn1(p0 int64) uint8 {
    p0 -= (p0 * (p0 * (p0 + 303)))
    p0 = (p0 * (p0 << ((65535) & 63)))
    v9 := []int{-56, 2147483647}
    _ = p0
    _ = v9
    return 255
}

func main() {
    var v1 string = (("hello" + "") + "foo bar")
    var v2 float64 = (0.1 / 1e10)
    var v3 int64 = 738
    switch v1 {
    case "\n\t":
        var v10 int16 = -32768
        _ = v10
    case "a":
        v11 := []int{-1, 9223372036854775807}
        _ = v11
    default:
        v2 -= (v2 + ((v2 * (float64(3.14159) + ((((v2 / v2) * v2) - v2) - v2))) - 0.1))
    }
    v12 := GenS1{F0: 1, F1: 744}
    v3 = 42
    var v13 string = (v1 + ((v1 + (v1 + (v1 + (v1 + (v1 + (v1 + v1)))))) + (v1 + (v1 + v1))))
    v3 -= v3
    if true {
        var v14 float64 = (v2 * ((v2 * (v2 / (v2 * (0.1 + (v2 - v2))))) + float64(float32(v2))))
        _ = v14
    } else {
        var v15 int16 = -615
        for fi16 := 0; fi16 < 10; fi16++ {
            var v17 string = (("foo bar" + v13) + (v13 + (v1 + ("a" + ("\n\t" + v13)))))
            for fi18 := 0; fi18 < 17; fi18++ {
                fmt.Println(v3, v3, v13)
                v3 -= 9223372036854775807
                var v19 int64 = v3
                _ = fi18
                _ = v19
            }
            v20 := []int{32768, fi16, fi16}
            _ = fi16
            _ = v17
            _ = v20
        }
        _ = v15
    }
    v21 := []int{0}
    switch v1 {
    case "":
        var v22 int8 = -67
        _ = v22
    case "hello":
        var v23 int = 42
        _ = v23
    case "a":
        v3 -= v3
    }
    v13 = (v13 + "\n\t")
    var v24 float64 = float64((0.1 / (((0.1 - 0.0) - -2.25) / (float32(v2) - (1.5 + (0.0 / 3.14159))))))
    v25 := GenS1{F0: -32768, F1: 255}
    var v26 int8 = 42
    switch "\n\t" {
    case "\n\t":
        for fi27 := 0; fi27 < 16; fi27++ {
            switch v13 {
            case "a":
                var v28 int64 = v3
                _ = v28
            }
            v3 += 140
            _ = fi27
        }
    default:
        for fi29 := 0; fi29 < 13; fi29++ {
            switch fi29 {
            case -1:
                v3 -= v3
            case 0:
                var v30 string = "a"
                _ = v30
            case 255:
                for fi31 := 0; fi31 < 7; fi31++ {
                    if len(v21) >= 2 {
                        v21 = v21[1:len(v21)]
                    }
                    v32 := GenS2{F0: v1}
                    v24 -= (-2.25 + 3.14159)
                    _ = fi31
                    _ = v32
                }
            default:
                if len(v21) >= 2 {
                    v21 = v21[1:len(v21)]
                }
            }
            v24 += (v24 + v2)
            v24 += 0.1
            _ = fi29
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v21
    _ = v24
    _ = v25
    _ = v26
    _ = v3
}
