package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
func genFn0(p0 int8, p1 string, p2 uint16) string {
    p0 += (p0 + p0)
    var v6 int64 = 0
    if true {
        if false {
            var v7 float64 = 1.5
            v8 := []int{-163, 1}
            _ = v7
            _ = v8
        }
    } else {
        p0 = (p0 % ((p0) | 1))
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return p1
}

func main() {
    var v1 float32 = (0.0 * -2.25)
    v2 := []string{"hello"}
    v3 := []string{}
    var v4 int32 = 1
    v5 := []uint32{}
    fmt.Println(v1, v4)
    var v9 string = "a"
    var v10 int16 = -825
    switch -9223372036854775808 {
    case 32767:
        _ = len(v3)
    }
    var v11 int16 = 42
    for fi12 := 0; fi12 < 19; fi12++ {
        if true {
            for fi13 := 0; fi13 < 14; fi13++ {
                var v14 uint8 = 73
                _ = fi13
                _ = v14
            }
            v15 := GenS1{F0: 255}
            var v16 uint32 = 2147483647
            _ = v15
            _ = v16
        }
        v11 += (v11 & (v10 | (v11 - (v11 & (^(v10))))))
        _ = fi12
    }
    for fi17 := 0; fi17 < 6; fi17++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi17
    }
    var v18 string = (v9 + (v9 + v9))
    v1 -= (v1 * -2.25)
    fmt.Println(v4, v11)
    var v19 string = ("foo bar" + v18)
    for fi20 := 0; fi20 < 3; fi20++ {
        for fi21 := 0; fi21 < 6; fi21++ {
            var v22 uint = 255
            _ = fi21
            _ = v22
        }
        v23 := GenS1{F0: 95}
        switch fi20 {
        case 128:
            _ = len(v2)
        case 9223372036854775807:
            var v24 uint = 18446744073709551615
            _ = v24
        case -222:
            var v25 uint = 4294967295
            _ = v25
        default:
            for fi26 := 0; fi26 < 6; fi26++ {
                fmt.Println(v19, v10, v19)
                var v27 int = (fi26 << ((4294967295) & 63))
                _ = fi26
                _ = v27
            }
        }
        _ = fi20
        _ = v23
    }
    var v28 uint32 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v18
    _ = v19
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
