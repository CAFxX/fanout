package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16 }
type GenS2 struct { F0 uint16; F1 int16; F2 bool }
type GenS3 struct { F0 uint16; F1 bool }
type GenS4 struct { F0 int32; F1 int; F2 int64 }
func genFn0(p0 int8, p1 int32, p2 int32) uint {
    switch 128 {
    case -2147483648:
        if false {
            var v5 int32 = p2
            v6 := []int{0, 0}
            _ = v5
            _ = v6
        }
    case 32768:
        p1 += (p2 * p1)
    case 128:
        var v7 int64 = 483
        _ = v7
    default:
        p1 = (p2 - (p2 % ((p1) | 1)))
    }
    p0 = (p0 | int8(p1))
    _ = p0
    _ = p1
    _ = p2
    return 327
}
func genFn1(p0 int) uint64 {
    v8 := []int{32768}
    for fi9 := 0; fi9 < 15; fi9++ {
        if len(v8) > 0 {
            v8[0] = (^(p0))
        }
        _ = fi9
    }
    if true {
        if true {
            p0 -= p0
        } else {
            fmt.Println(p0)
            fmt.Println(p0)
            p0 -= p0
        }
        switch 509 {
        case -819:
            if true {
                var v10 uint16 = 65535
                if true {
                    v11 := GenS1{F0: 1, F1: 42}
                    _ = v11
                }
                fmt.Println(v10)
                _ = v10
            }
        default:
            v12 := GenS1{F0: 752, F1: -275}
            _ = v12
        }
        var v13 int16 = 789
        _ = v13
    } else {
        p0 += p0
    }
    _ = p0
    _ = v8
    return 218
}

func main() {
    v1 := make(map[string]string)
    v2 := []uint32{42, 2147483648, 683, 4294967295}
    var v3 bool = false
    v4 := []string{"", ("hello" + ("foo bar" + "")), "foo bar"}
    var v14 int32 = -875
    v14 += v14
    v15 := v1[("\n\t" + "foo bar")]
    var v16 uint64 = 18446744073709551615
    for fi17 := 0; fi17 < 11; fi17++ {
        v18 := GenS2{F0: 65535, F1: -32768, F2: (!(0 == 2))}
        v3 = v3
        fmt.Println(v15, v16)
        _ = fi17
        _ = v18
    }
    for fi19 := 0; fi19 < 3; fi19++ {
        switch ((((((v15 + v15) + v15) + (((v15 + v15) + ((((v15 + ((v15 + "foo bar") + "a")) + v15) + "\n\t") + v15)) + "a")) + v15) + (v15 + ((v15 + v15) + "a"))) + (v15 + (v15 + "\n\t"))) {
        case "foo bar":
            for fi20 := 0; fi20 < 2; fi20++ {
                _ = len(v4)
                fmt.Println(fi19)
                _ = fi20
            }
        case "":
            if true {
                v16 = (v16 + (v16 >> ((870) & 63)))
            } else {
                v16 -= 1
            }
        }
        var v21 float64 = -2.25
        for fi22 := 0; fi22 < 14; fi22++ {
            v23 := GenS3{F0: 32767, F1: true}
            fmt.Println(fi22, v15)
            var v24 int8 = -84
            _ = fi22
            _ = v23
            _ = v24
        }
        _ = fi19
        _ = v21
    }
    v14 = (v14 + (^(v14)))
    switch 1 {
    case -399:
        fmt.Println(v14, v15, v14)
    default:
        fmt.Println(v3)
    }
    if (!(true && v3)) {
        _ = len(v2)
        v16 -= v16
        v15 = (v15 + v15)
    }
    fmt.Println(v14, v16, v3)
    v25 := GenS4{F0: v14, F1: 4294967295, F2: 42}
    var v26 uint64 = v16
    var v27 float64 = 1.5
    v16 = 2
    for fi28 := 0; fi28 < 8; fi28++ {
        for fi29 := 0; fi29 < 14; fi29++ {
            var v30 float32 = 1e10
            _ = fi29
            _ = v30
        }
        v26 = 0
        _ = fi28
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
}
