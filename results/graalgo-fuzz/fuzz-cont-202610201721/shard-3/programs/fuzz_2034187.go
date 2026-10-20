package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint16; F1 uint32 }
func genFn0(p0 bool) int64 {
    p0 = p0
    p0 = (p0 && (1e10 == (float64((-2.25 / (-2.25 / 3.14159))) + 0.0)))
    switch "foo bar" {
    case "foo bar":
        if p0 {
            p0 = false
            v4 := GenS1{F0: 65535}
            if ((!(797 == 0)) && true) {
                for fi5 := 0; fi5 < 1; fi5++ {
                    fmt.Println(p0, fi5)
                    _ = fi5
                }
                v6 := []int{}
                _ = v6
            }
            _ = v4
        } else {
            var v7 int16 = 694
            v7 = 1000
            _ = v7
        }
    case "\n\t":
        v8 := []int{127, 673}
        _ = v8
    case "hello":
        var v9 bool = (false || (p0 && (-2147483648 == -930)))
        _ = v9
    }
    var v10 uint16 = 42
    _ = p0
    _ = v10
    return 415
}
func genFn1(p0 int64) uint16 {
    var v11 int16 = -32768
    fmt.Println(p0, p0)
    _ = p0
    _ = v11
    return 339
}

func main() {
    var v1 uint = 910
    var v2 int16 = -86
    var v3 uint16 = 2
    v12 := []int{}
    for fi13 := 0; fi13 < 3; fi13++ {
        v1 -= v1
        for fi14 := 0; fi14 < 13; fi14++ {
            if len(v12) > 0 {
                v12[0] = fi13
            }
            var v15 uint32 = 2
            _ = fi14
            _ = v15
        }
        v2 = v2
        _ = fi13
    }
    for fi16 := 0; fi16 < 5; fi16++ {
        for fi17 := 0; fi17 < 1; fi17++ {
            v18 := GenS1{F0: genFn1(0)}
            fmt.Println(fi16)
            _ = fi17
            _ = v18
        }
        fmt.Println(v2)
        if len(v12) >= 2 {
            v12 = v12[1:len(v12)]
        }
        _ = fi16
    }
    var v19 float64 = ((-2.25 / 1.5) - ((-2.25 * 3.14159) * 0.1))
    v20 := GenS2{F0: genFn1(0), F1: 2147483648}
    var v21 int = 42
    switch 21 {
    case 32767:
        v3 += 0
    case 0:
        fmt.Println(v3)
    case 950:
        v22 := GenS1{F0: v3}
        _ = v22
    default:
        fmt.Println(v3, v1, v1)
    }
    v23 := GenS2{F0: (v3 | (v3 >> (((v1 << (((v1 | uint(v3))) & 63))) & 63))), F1: 527}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v23
    _ = v3
}
