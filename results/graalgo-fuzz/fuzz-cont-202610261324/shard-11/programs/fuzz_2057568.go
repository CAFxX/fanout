package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 bool; F1 bool }
func genFn0() uint {
    switch "foo bar" {
    case "foo bar":
        v5 := []int{32767, 9223372036854775807, 42}
        _ = v5
    case "":
        v6 := []int{58, 9223372036854775807, 127, 327}
        _ = v6
    default:
        var v7 int16 = -1
        _ = v7
    }
    if (1 < 0) {
        var v8 uint16 = 32767
        var v9 int64 = 42
        _ = v8
        _ = v9
    } else {
        var v10 uint8 = 255
        v11 := GenS1{F0: -17, F1: 991}
        _ = v10
        _ = v11
    }
    if true {
        if false {
            v12 := []int{}
            var v13 int32 = 221
            v13 -= (v13 - (v13 + (v13 % ((-2147483648) | 1))))
            _ = v12
            _ = v13
        }
        var v14 uint16 = 0
        for fi15 := 0; fi15 < 17; fi15++ {
            fmt.Println(v14, fi15)
            _ = fi15
        }
        _ = v14
    }
    for fi16 := 0; fi16 < 15; fi16++ {
        var v17 float32 = 1e10
        v18 := []int{fi16}
        if len(v18) > 0 {
            v18[0] = fi16
        }
        _ = fi16
        _ = v17
        _ = v18
    }
    return 18446744073709551615
}
func genFn1(p0 int, p1 uint64) int {
    if true {
        p0 = 0
    }
    fmt.Println(p0)
    v19 := GenS1{F0: p0, F1: 2}
    p0 = 127
    _ = p0
    _ = p1
    _ = v19
    return 209
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []string{(("hello" + "\n\t") + (("hello" + "foo bar") + ("" + "hello"))), ""}
    var v3 float32 = 3.14159
    var v4 uint16 = 32767
    switch ("foo bar" + "\n\t") {
    case "a":
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    v3 = 0.1
    v3 -= (((float32(-2.25) + -2.25) * (3.14159 + 1.5)) - v3)
    v20 := GenS1{F0: 128, F1: 42}
    v1[2] = 2
    for fi21 := 0; fi21 < 20; fi21++ {
        v3 = v3
        _ = fi21
    }
    v22 := v1[uint32(v4)]
    v4 = 32767
    var v23 bool = false
    v24 := GenS1{F0: 9223372036854775807, F1: 155}
    if (!(!false)) {
        v1[(v22 | 1)] = 1
        var v25 uint = 4294967295
        switch genFn1(int(v25), 1) {
        case 32767:
            var v26 float64 = 1e10
            _ = v26
        case 42:
            fmt.Println(v25, v25)
        }
        _ = v25
    } else {
        _ = len(v1)
    }
    switch "foo bar" {
    case "\n\t":
        v22 += (v22 * (v22 / ((v22) | 1)))
    default:
        for fi27 := 0; fi27 < 16; fi27++ {
            if v23 {
                v28 := GenS2{F0: int16(fi27)}
                _ = v28
            }
            _ = fi27
        }
    }
    var v29 string = ""
    var v30 int8 = -62
    v31 := GenS3{F0: (2 != 127), F1: true}
    if v23 {
        for fi32 := 0; fi32 < 1; fi32++ {
            fmt.Println(v30, v30)
            v33 := GenS1{F0: genFn1((fi32 ^ fi32), 4294967295), F1: 173}
            _ = fi32
            _ = v33
        }
        var v34 uint32 = (^(v22))
        for fi35 := 0; fi35 < 12; fi35++ {
            v34 = 1
            fmt.Println(v30, v29, v30)
            for fi36 := 0; fi36 < 13; fi36++ {
                v29 = v29
                var v37 int16 = 32767
                _ = fi36
                _ = v37
            }
            _ = fi35
        }
        _ = v34
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v22
    _ = v23
    _ = v24
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
}
