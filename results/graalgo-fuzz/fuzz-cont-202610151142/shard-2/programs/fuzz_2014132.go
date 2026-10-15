package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 bool }
type GenS2 struct { F0 int32; F1 uint }
func genFn0(p0 int16, p1 string, p2 int) int {
    p0 = (^(p0))
    if (!((2 >= 42) && ((p0 * p0) < 32767))) {
        fmt.Println(p1)
        p1 = ("foo bar" + "\n\t")
    }
    p2 = p2
    v6 := GenS1{F0: 516, F1: p0, F2: true}
    p2 += 42
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return -396
}
func genFn1(p0 bool, p1 string, p2 int8) int64 {
    var v7 string = "foo bar"
    v8 := []int{32767, 2147483647, genFn0(-32768, "foo bar", 42), 4294967295}
    for fi9 := 0; fi9 < 5; fi9++ {
        var v10 string = v7
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 0
}

func main() {
    var v1 int16 = -32768
    var v2 uint16 = 1
    v3 := []uint32{2147483648, 2, 2147483647}
    v4 := make(map[uint32]uint32)
    var v5 int8 = 52
    for fi11 := 0; fi11 < 3; fi11++ {
        fmt.Println(v2, v5)
        switch ("" + "\n\t") {
        case "":
            v12 := GenS1{F0: genFn1((42 >= 42), "", -8), F1: v1, F2: false}
            _ = v12
        default:
            v5 += 55
        }
        fmt.Println(v5, v1)
        _ = fi11
    }
    v5 = 78
    if false {
        for fi13 := 0; fi13 < 8; fi13++ {
            var v14 uint64 = 4294967295
            fmt.Println(fi13)
            _ = fi13
            _ = v14
        }
        v5 = int8(v1)
        for fi15 := 0; fi15 < 3; fi15++ {
            v16 := GenS1{F0: 0, F1: v1, F2: ((v2 << ((65535) & 63)) == v2)}
            _ = fi15
            _ = v16
        }
    }
    v17 := GenS2{F0: -2147483648, F1: 42}
    for fi18 := 0; fi18 < 18; fi18++ {
        v5 = 100
        v19 := GenS2{F0: -1, F1: 0}
        _ = fi18
        _ = v19
    }
    delete(v4, 774)
    if true {
        switch "\n\t" {
        case "":
            v1 -= (v1 ^ int16(v2))
        case "hello":
            v1 = (v1 + v1)
        case "foo bar":
            v1 = (-(v1))
        }
        var v20 bool = false
        _ = v20
    } else {
        v2 = 2
    }
    for fi21 := 0; fi21 < 13; fi21++ {
        var v22 int16 = v1
        v2 = 2
        _ = fi21
        _ = v22
    }
    var v23 int16 = v1
    v23 -= (^(v1))
    for fi24 := 0; fi24 < 12; fi24++ {
        v25 := GenS1{F0: 215, F1: 540, F2: true}
        _ = fi24
        _ = v25
    }
    v26 := v4[42]
    v27 := GenS1{F0: genFn1((2 >= 127), (("foo bar" + ("foo bar" + "\n\t")) + "a"), v5), F1: int16(v26), F2: (int16(v26) > v1)}
    var v28 int16 = v1
    var v29 uint64 = 18446744073709551615
    fmt.Println(v2, v2)
    for fi30 := 0; fi30 < 17; fi30++ {
        var v31 string = "\n\t"
        _ = fi30
        _ = v31
    }
    switch 32767 {
    case 563:
        for fi32 := 0; fi32 < 11; fi32++ {
            switch (fi32 & int(v2)) {
            case -200:
                v29 = v29
            case 0:
                v26 += v26
            case 4294967295:
                v4[(v26 / (((^(v26))) | 1))] = 42
            default:
                if true {
                    var v33 int64 = 0
                    _ = v33
                } else {
                    v5 += v5
                    var v34 uint16 = v2
                    _ = v34
                }
            }
            v35 := GenS1{F0: 0, F1: (-(v1)), F2: true}
            _ = fi32
            _ = v35
        }
    case 128:
        switch "hello" {
        case "a":
            fmt.Println(v1, v29, v26)
        case "hello":
            var v36 float64 = 0.0
            _ = v36
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v23
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
