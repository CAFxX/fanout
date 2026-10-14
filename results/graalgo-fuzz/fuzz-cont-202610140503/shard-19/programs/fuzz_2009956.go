package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 string; F1 int16; F2 uint }
type GenS4 struct { F0 int64; F1 uint16; F2 bool }
func genFn0(p0 uint, p1 string, p2 int64) int {
    v7 := GenS1{F0: 9223372036854775807}
    var v8 string = (p1 + "\n\t")
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 1
}
func genFn1() uint8 {
    if false {
        if (!(162 <= 1)) {
            var v9 float64 = 0.1
            for fi10 := 0; fi10 < 7; fi10++ {
                switch fi10 {
                case 32768:
                    v11 := GenS1{F0: fi10}
                    _ = v11
                case -2147483648:
                    var v12 float32 = ((float32(v9) * float32(v9)) - (1.5 * 1.5))
                    _ = v12
                }
                _ = fi10
            }
            _ = v9
        }
    } else {
        var v13 int = 2147483647
        for fi14 := 0; fi14 < 16; fi14++ {
            v15 := GenS2{F0: true}
            _ = fi14
            _ = v15
        }
        v13 -= (v13 + 127)
        _ = v13
    }
    var v16 int32 = 229
    v17 := []int{973, -1}
    switch (("foo bar" + ("hello" + "hello")) + "a") {
    case "a":
        var v18 uint16 = 2
        _ = v18
    }
    _ = v16
    _ = v17
    return 2
}

func main() {
    v1 := []uint32{}
    v2 := []int{2147483647, 32768, 507, 15}
    v3 := []uint32{65535, 1}
    var v4 uint = 4294967295
    v5 := []uint32{768, 2, 42, 0}
    var v6 uint8 = 2
    v4 += (v4 * v4)
    _ = len(v3)
    switch "\n\t" {
    case "\n\t":
        v6 += v6
    default:
        var v19 uint8 = genFn1()
        _ = v19
    }
    v4 = 4294967295
    var v20 uint8 = (v6 + (v6 / ((v6) | 1)))
    switch genFn0(462, ("" + ""), -413) {
    case -2147483648:
        switch 42 {
        case 127:
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        }
    case 419:
        v4 += (v4 / ((250) | 1))
    case -1:
        var v21 int16 = 42
        _ = v21
    default:
        for fi22 := 0; fi22 < 20; fi22++ {
            if ((v20 << ((65535) & 63)) < (v20 * v6)) {
                v23 := GenS2{F0: ((0.0 / float32(float64((((0.0 * float32(-2.25)) / 1.0) - (1e10 / 1.0))))) <= 3.14159)}
                fmt.Println(v4, v4)
                v4 -= v4
                _ = v23
            }
            _ = fi22
        }
    }
    v24 := GenS3{F0: (("hello" + "foo bar") + ("\n\t" + "a")), F1: 0, F2: (^(v4))}
    fmt.Println(v20, v4, v4)
    v25 := GenS4{F0: 42, F1: uint16(v6), F2: true}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
