package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string; F2 uint8 }
func genFn0() uint {
    if false {
        if (732 >= 2147483648) {
            var v5 uint8 = 128
            _ = v5
        }
    } else {
        var v6 float32 = 0.1
        var v7 int16 = 0
        for fi8 := 0; fi8 < 11; fi8++ {
            v6 = (v6 / float32((3.14159 * 1e10)))
            _ = fi8
        }
        _ = v6
        _ = v7
    }
    var v9 string = "\n\t"
    var v10 uint16 = 32767
    _ = v10
    _ = v9
    return 255
}
func genFn1(p0 uint32, p1 int32, p2 bool) uint16 {
    for fi11 := 0; fi11 < 14; fi11++ {
        var v12 int32 = (p1 / ((p1) | 1))
        for fi13 := 0; fi13 < 16; fi13++ {
            v14 := []int{fi11}
            _ = fi13
            _ = v14
        }
        _ = fi11
        _ = v12
    }
    p2 = (0 != 9223372036854775807)
    switch 949 {
    case -226:
        v15 := []int{}
        _ = v15
    case -399:
        p1 += p1
    case -2147483648:
        if true {
            if false {
                v16 := GenS1{F0: -1, F1: "foo bar", F2: 127}
                var v17 string = (("a" + "a") + "\n\t")
                _ = v16
                _ = v17
            } else {
                var v18 int = 4294967295
                _ = v18
            }
        }
    default:
        v19 := []int{}
        _ = v19
    }
    if true {
        p2 = p2
        p2 = ((("\n\t" > "a") && true) || false)
        fmt.Println(p0, p1, p2)
    }
    for fi20 := 0; fi20 < 5; fi20++ {
        v21 := []int{(fi20 % ((-1) | 1)), fi20, (fi20 | int(p1))}
        switch ("a" + "") {
        case "foo bar":
            if (-2 > -402) {
                var v22 uint = uint(p0)
                _ = v22
            }
        case "":
            var v23 uint16 = 2
            _ = v23
        case "\n\t":
            var v24 float64 = (-2.25 * 0.0)
            _ = v24
        }
        p2 = (!(!(1 > 220)))
        _ = fi20
        _ = v21
    }
    _ = p0
    _ = p1
    _ = p2
    return 255
}

func main() {
    var v1 int32 = 42
    var v2 uint = 260
    var v3 uint8 = 164
    v4 := []int{}
    if len(v4) > 0 {
        v4[0] = -1
    }
    v3 = (v3 / ((2) | 1))
    switch 9223372036854775807 {
    case 32768:
        v1 = v1
    case 247:
        fmt.Println(v1, v2)
    default:
        var v25 int32 = (v1 | (-(v1)))
        _ = v25
    }
    var v26 string = ""
    v26 = v26
    v1 += (v1 - v1)
    v27 := GenS1{F0: 1, F1: v26, F2: (v3 << ((v2) & 63))}
    for fi28 := 0; fi28 < 9; fi28++ {
        if false {
            var v29 int32 = 1
            _ = v29
        }
        v30 := GenS1{F0: 9223372036854775807, F1: "hello", F2: v3}
        _ = len(v4)
        _ = fi28
        _ = v30
    }
    v2 += (v2 / ((65535) | 1))
    v26 = (v26 + (v26 + v26))
    v1 = v1
    var v31 uint8 = v3
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v31
    _ = v4
}
