package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int }
type GenS2 struct { F0 int }
func genFn0(p0 uint16, p1 int, p2 int) string {
    p1 -= -1
    if false {
        v4 := []int{(p1 * p1), (p1 % ((9223372036854775807) | 1)), (p2 ^ (-(p1))), 61}
        _ = v4
    } else {
        if (!(42 < 2)) {
            var v5 int = (p1 ^ (p1 * p1))
            v6 := GenS1{F0: true, F1: p1}
            _ = v5
            _ = v6
        }
        v7 := []int{p2, (^(p1))}
        _ = v7
    }
    p2 = p2
    var v8 uint8 = 124
    for fi9 := 0; fi9 < 12; fi9++ {
        v10 := GenS1{F0: false, F1: p2}
        v11 := GenS2{F0: -2147483648}
        for fi12 := 0; fi12 < 14; fi12++ {
            p2 += fi12
            v13 := []int{p2, 127}
            var v14 int64 = 1
            v15 := GenS2{F0: (p1 * (p2 << ((0) & 63)))}
            _ = fi12
            _ = v13
            _ = v14
            _ = v15
        }
        _ = fi9
        _ = v10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return "foo bar"
}
func genFn1(p0 int64, p1 int8) uint16 {
    switch 251 {
    case 32768:
        p0 = 1
    default:
        for fi16 := 0; fi16 < 19; fi16++ {
            v17 := GenS2{F0: (fi16 << ((18446744073709551615) & 63))}
            _ = fi16
            _ = v17
        }
    }
    switch 259 {
    case -1:
        for fi18 := 0; fi18 < 16; fi18++ {
            p1 = 100
            for fi19 := 0; fi19 < 6; fi19++ {
                var v20 string = ""
                _ = fi19
                _ = v20
            }
            var v21 int8 = (p1 / ((p1) | 1))
            _ = fi18
            _ = v21
        }
    case -180:
        p0 = (p0 % (((-(p0))) | 1))
    }
    _ = p0
    _ = p1
    return 246
}

func main() {
    var v1 int16 = 985
    v2 := []int{42, 734, 32767, 9223372036854775807}
    v3 := make(map[uint32]string)
    v1 -= v1
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if len(v2) > 0 {
        v2[0] = 2147483647
    }
    if true {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    } else {
        v1 = (v1 + 1)
    }
    var v22 int64 = 0
    var v23 uint = 18446744073709551615
    v22 = (v22 - (-(v22)))
    fmt.Println(v22, v22, v1)
    for fi24 := 0; fi24 < 11; fi24++ {
        var v25 int32 = -841
        _ = fi24
        _ = v25
    }
    var v26 uint16 = uint16(v23)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v26
    _ = v3
}
