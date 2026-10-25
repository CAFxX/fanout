package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16 }
func genFn0(p0 uint8) uint64 {
    switch "a" {
    case "foo bar":
        var v5 int32 = 1
        _ = v5
    case "":
        var v6 float64 = -2.25
        _ = v6
    case "hello":
        p0 = 1
    default:
        for fi7 := 0; fi7 < 16; fi7++ {
            p0 = (^(p0))
            p0 = 42
            p0 = (p0 & p0)
            _ = fi7
        }
    }
    var v8 uint8 = p0
    v8 += (^(p0))
    if (32767 != 973) {
        for fi9 := 0; fi9 < 10; fi9++ {
            switch (fi9 | (-(fi9))) {
            case 4294967295:
                for fi10 := 0; fi10 < 17; fi10++ {
                    v11 := []int{fi10, fi10, -9223372036854775808, fi10}
                    _ = fi10
                    _ = v11
                }
            }
            p0 = 42
            _ = fi9
        }
    } else {
        for fi12 := 0; fi12 < 18; fi12++ {
            v13 := GenS1{F0: 606, F1: 32767}
            _ = fi12
            _ = v13
        }
    }
    var v14 float64 = 0.1
    _ = p0
    _ = v14
    _ = v8
    return 42
}
func genFn1(p0 bool, p1 int16, p2 int) uint16 {
    v15 := GenS1{F0: 42, F1: -1}
    v16 := []int{p2}
    var v17 int64 = 310
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    _ = v16
    _ = v17
    return 1
}

func main() {
    var v1 uint32 = 2147483647
    v2 := make(map[uint32]string)
    v3 := []int{}
    var v4 uint16 = 255
    v4 += v4
    v1 = v1
    if false {
        for fi18 := 0; fi18 < 9; fi18++ {
            var v19 string = ("foo bar" + ("a" + ((("a" + "foo bar") + (("" + "") + "hello")) + (("" + "\n\t") + ("hello" + ("hello" + "a"))))))
            if len(v3) > 0 {
                v3[0] = -9223372036854775808
            }
            _ = fi18
            _ = v19
        }
        v20 := GenS1{F0: 42, F1: 0}
        for fi21 := 0; fi21 < 2; fi21++ {
            v4 += v4
            var v22 string = "hello"
            v1 -= (v1 % ((449) | 1))
            _ = fi21
            _ = v22
        }
        _ = v20
    }
    var v23 int8 = 0
    var v24 int8 = (v23 ^ int8(v4))
    v3 = append(v3, -9223372036854775808)
    v4 = v4
    v1 = (v1 - (v1 * v1))
    v24 -= v23
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}
