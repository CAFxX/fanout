package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool }
type GenS2 struct { F0 uint16; F1 string; F2 uint }
type GenS3 struct { F0 uint32; F1 int }
func genFn0(p0 uint16, p1 int32, p2 uint) bool {
    if (521 == 42) {
        switch "" {
        case "":
            var v6 int8 = -108
            _ = v6
        case "a":
            p2 -= (p2 + 845)
        case "hello":
            p2 = p2
        }
        var v7 float64 = 3.14159
        v8 := []int{9223372036854775807, -687, 1, -274}
        _ = v7
        _ = v8
    }
    fmt.Println(p2, p2, p0)
    p2 += 2
    for fi9 := 0; fi9 < 12; fi9++ {
        p2 = (p2 * 2)
        if true {
            for fi10 := 0; fi10 < 10; fi10++ {
                p2 += p2
                _ = fi10
            }
            p0 = p0
            p2 -= p2
        }
        _ = fi9
    }
    for fi11 := 0; fi11 < 4; fi11++ {
        switch "a" {
        case "":
            var v12 string = "\n\t"
            _ = v12
        case "foo bar":
            for fi13 := 0; fi13 < 16; fi13++ {
                switch 128 {
                case 42:
                    p0 += (p0 % ((1) | 1))
                case 9223372036854775807:
                    var v14 float64 = 1e10
                    _ = v14
                default:
                    var v15 int = fi13
                    _ = v15
                }
                _ = fi13
            }
        }
        v16 := GenS1{F0: -1, F1: (0 != -1)}
        p1 = (p1 | 2147483647)
        _ = fi11
        _ = v16
    }
    _ = p0
    _ = p1
    _ = p2
    return (41 <= p0)
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 float32 = 0.0
    v3 := []int{127, -2147483648, -322}
    v4 := make(map[uint32]int)
    var v5 float32 = v2
    v17 := GenS1{F0: 100, F1: (true && ((1.5 * 3.14159) <= (3.14159 * ((-2.25 / (((3.14159 + ((3.14159 * 0.0) * 1.5)) * 1e10) / 1.0)) * 1e10))))}
    fmt.Println(v5)
    for fi18 := 0; fi18 < 5; fi18++ {
        for fi19 := 0; fi19 < 20; fi19++ {
            v2 -= v5
            _ = fi19
        }
        var v20 float64 = 1.5
        _ = fi18
        _ = v20
    }
    fmt.Println(v2)
    if (519 >= 42) {
        for fi21 := 0; fi21 < 20; fi21++ {
            delete(v1, 958)
            v22 := GenS2{F0: 65535, F1: "a", F2: 1}
            _ = fi21
            _ = v22
        }
    }
    v3 = append(v3, 127)
    var v23 string = (("a" + "a") + "")
    v5 = (((float32(float64(1.5)) - 0.0) - (1e10 - v5)) - (float32(-2.25) - 3.14159))
    switch v23 {
    case "\n\t":
        v23 = "hello"
    case "":
        v24 := GenS1{F0: 127, F1: true}
        _ = v24
    case "hello":
        var v25 int = 1
        _ = v25
    default:
        v26 := GenS3{F0: 65535, F1: -83}
        _ = v26
    }
    var v27 float32 = (v2 / v5)
    v2 = v2
    delete(v1, 102)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v23
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
