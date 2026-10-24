package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int }
func genFn0() uint8 {
    var v6 string = ""
    v6 = v6
    for fi7 := 0; fi7 < 8; fi7++ {
        v8 := []int{(fi7 ^ fi7), (fi7 & (fi7 << ((767) & 63))), (fi7 >> ((255) & 63)), 128}
        _ = fi7
        _ = v8
    }
    v6 = v6
    _ = v6
    return 2
}
func genFn1(p0 uint16, p1 uint64) uint64 {
    for fi9 := 0; fi9 < 15; fi9++ {
        v10 := []int{fi9, fi9, 246, 60}
        if len(v10) >= 2 {
            v10 = v10[1:len(v10)]
        }
        if ((fi9 << ((866) & 63)) == 1) {
            p1 = (^(p1))
            p0 = p0
            p0 -= p0
        }
        _ = fi9
        _ = v10
    }
    v11 := GenS1{F0: 42, F1: -2147483648}
    _ = p0
    _ = p1
    _ = v11
    return p1
}

func main() {
    var v1 int16 = 32767
    v2 := make(map[string]uint32)
    var v3 uint16 = 65535
    var v4 uint8 = 2
    var v5 int16 = (v1 / ((v1) | 1))
    for fi12 := 0; fi12 < 4; fi12++ {
        var v13 uint32 = 4294967295
        _ = fi12
        _ = v13
    }
    switch 9223372036854775807 {
    case 32767:
        switch 2147483647 {
        case -197:
            v14 := []int{}
            _ = v14
        case -825:
            var v15 float64 = 1.5
            _ = v15
        case 128:
            for fi16 := 0; fi16 < 18; fi16++ {
                v17 := []int{(fi16 & (fi16 | -617)), fi16, fi16}
                _ = fi16
                _ = v17
            }
        default:
            var v18 uint64 = 641
            _ = v18
        }
    }
    fmt.Println(v3)
    v5 += (v1 >> ((4294967295) & 63))
    v5 += -1
    var v19 string = (("" + "\n\t") + ((("hello" + "foo bar") + "") + "hello"))
    _ = len(v2)
    v20 := GenS1{F0: 127, F1: 42}
    v5 -= 32767
    var v21 float64 = 0.1
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
