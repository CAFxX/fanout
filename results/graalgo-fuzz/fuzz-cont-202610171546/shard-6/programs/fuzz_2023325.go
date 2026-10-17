package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string }
type GenS2 struct { F0 uint; F1 int8; F2 int8 }
func genFn0(p0 string, p1 int8, p2 uint64) uint64 {
    switch (p0 + p0) {
    case "\n\t":
        p1 += p1
    case "a":
        var v7 int16 = -214
        _ = v7
    case "foo bar":
        v8 := GenS1{F0: 4294967295, F1: p0}
        _ = v8
    default:
        var v9 uint16 = 32767
        _ = v9
    }
    v10 := GenS2{F0: 255, F1: p1, F2: (p1 - -1)}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return p2
}
func genFn1(p0 int16, p1 uint16) int {
    p0 = p0
    p0 = -550
    p0 -= (-(p0))
    var v11 uint32 = 2
    if (v11 != 42) {
        var v12 int64 = 0
        _ = v12
    } else {
        p0 += (p0 ^ 0)
        p0 += (^(p0))
        if false {
            var v13 float32 = ((-2.25 + -2.25) + 0.1)
            var v14 uint16 = p1
            var v15 float32 = v13
            _ = v13
            _ = v14
            _ = v15
        }
    }
    _ = p0
    _ = p1
    _ = v11
    return 127
}

func main() {
    var v1 uint16 = 1
    v2 := []string{("" + ("foo bar" + "")), ("\n\t" + ("hello" + "hello")), (("a" + ("" + ("\n\t" + "foo bar"))) + ((("a" + "\n\t") + ("hello" + ((("a" + "") + "hello") + "foo bar"))) + ("hello" + "")))}
    v3 := []uint32{110}
    var v4 uint32 = 1
    v5 := []uint32{(v4 ^ 0), v4, (v4 - (v4 << ((177) & 63)))}
    var v6 uint8 = 1
    if (((-2.25 + -2.25) - 0.0) >= 0.1) {
        for fi16 := 0; fi16 < 3; fi16++ {
            if (false && (-1 > 42)) {
                fmt.Println(v1)
            }
            var v17 uint16 = (v1 * (v1 * (v1 + v1)))
            v17 = v17
            _ = fi16
            _ = v17
        }
        var v18 int = genFn1(-359, v1)
        _ = v18
    }
    if len(v3) > 0 {
        v3[0] = (v4 / (((^(v4))) | 1))
    }
    v5 = append(v5, 57)
    v4 += 22
    fmt.Println(v1, v1, v4)
    _ = len(v2)
    v1 += (v1 << ((1) & 63))
    var v19 string = "a"
    if len(v2) > 0 {
        v2[0] = (v19 + (v19 + "\n\t"))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
