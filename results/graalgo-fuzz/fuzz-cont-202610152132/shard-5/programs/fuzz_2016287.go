package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64; F2 bool }
type GenS2 struct { F0 int16; F1 uint; F2 string }
func genFn0(p0 uint64, p1 int8, p2 int8) int8 {
    var v6 float64 = 0.1
    for fi7 := 0; fi7 < 16; fi7++ {
        var v8 string = "a"
        v9 := GenS1{F0: 18446744073709551615, F1: 485, F2: (true || false)}
        _ = fi7
        _ = v8
        _ = v9
    }
    p0 -= p0
    p1 = -1
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return -66
}
func genFn1(p0 bool, p1 int32) uint16 {
    for fi10 := 0; fi10 < 17; fi10++ {
        var v11 uint = 18446744073709551615
        v11 += (v11 >> ((v11) & 63))
        _ = fi10
        _ = v11
    }
    fmt.Println(p1)
    v12 := GenS1{F0: 1, F1: 1, F2: p0}
    if (p0 || (2 == 135)) {
        v13 := GenS1{F0: 0, F1: 18446744073709551615, F2: (float32(-2.25) > (((-2.25 - 1e10) - 1e10) + 1.5))}
        _ = v13
    } else {
        v14 := []int{0, 1, 4294967295}
        p0 = (!false)
        _ = v14
    }
    _ = p0
    _ = p1
    _ = v12
    return 233
}

func main() {
    var v1 int = 4294967295
    var v2 int32 = -920
    var v3 string = "foo bar"
    v4 := make(map[uint32]bool)
    var v5 uint = 2
    v15 := []int{-2147483648, v1}
    var v16 uint32 = 65535
    fmt.Println(v1, v5)
    for fi17 := 0; fi17 < 13; fi17++ {
        var v18 int64 = 0
        _ = fi17
        _ = v18
    }
    for fi19 := 0; fi19 < 4; fi19++ {
        var v20 float64 = -2.25
        if len(v15) >= 2 {
            v15 = v15[1:len(v15)]
        }
        _ = fi19
        _ = v20
    }
    switch v1 {
    case -9223372036854775808:
        for fi21 := 0; fi21 < 5; fi21++ {
            fmt.Println(v1)
            var v22 uint8 = 2
            _ = fi21
            _ = v22
        }
    }
    v5 -= (^(v5))
    var v23 uint64 = 2
    v4[v16] = (false && (!(!(v16 < v16))))
    v24 := GenS1{F0: (v5 | (v5 % ((4294967295) | 1))), F1: (^(v23)), F2: true}
    v25 := GenS2{F0: 1, F1: (v5 << (((^(v5))) & 63)), F2: "\n\t"}
    v3 = v3
    if ((!true) && true) {
        for fi26 := 0; fi26 < 3; fi26++ {
            v23 -= 1
            fmt.Println(v2, fi26, fi26)
            _ = fi26
        }
        if (0.0 == (-2.25 / 1e10)) {
            var v27 float32 = (0.0 - ((3.14159 * 3.14159) - 0.0))
            v1 += (-(v1))
            _ = v27
        } else {
            var v28 float32 = 0.1
            fmt.Println(v23)
            switch v3 {
            case "foo bar":
                v1 += v1
            }
            _ = v28
        }
    }
    if (((!(0 != 0)) || true) && (!true)) {
        v3 = (v3 + "foo bar")
    }
    for fi29 := 0; fi29 < 19; fi29++ {
        v2 -= v2
        _ = fi29
    }
    var v30 int64 = -1
    for fi31 := 0; fi31 < 16; fi31++ {
        v32 := GenS2{F0: 473, F1: 0, F2: "\n\t"}
        _ = fi31
        _ = v32
    }
    v5 -= v5
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
