package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int16; F1 string; F2 int8 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int8 }
func genFn0(p0 bool) int16 {
    p0 = (4294967295 > 27)
    if false {
        v5 := []int{32767}
        _ = v5
    }
    fmt.Println(p0)
    for fi6 := 0; fi6 < 7; fi6++ {
        p0 = true
        _ = fi6
    }
    _ = p0
    return 551
}
func genFn1() uint {
    switch "\n\t" {
    case "hello":
        var v7 int32 = -864
        _ = v7
    case "":
        if true {
            v8 := GenS1{F0: 493}
            var v9 uint = 4294967295
            _ = v8
            _ = v9
        } else {
            v10 := GenS2{F0: 1, F1: "a", F2: 100}
            var v11 uint8 = 42
            _ = v10
            _ = v11
        }
    }
    var v12 int64 = 9223372036854775807
    fmt.Println(v12)
    v13 := []int{128}
    for fi14 := 0; fi14 < 19; fi14++ {
        v15 := GenS3{F0: 2}
        _ = fi14
        _ = v15
    }
    _ = v12
    _ = v13
    return 96
}

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 uint16 = 831
    var v3 int8 = -77
    var v4 string = ("\n\t" + "a")
    var v16 float32 = (float32(0.1) / float32(3.14159))
    v17 := GenS4{F0: 9223372036854775807}
    var v18 int64 = 0
    v19 := []int{-9223372036854775808, -707, int(v3)}
    for fi20 := 0; fi20 < 5; fi20++ {
        if len(v19) > 0 {
            v19[0] = (fi20 / ((fi20) | 1))
        }
        var v21 int = 1
        for fi22 := 0; fi22 < 20; fi22++ {
            var v23 int32 = -1
            v16 = (0.0 * (float32(1e10) * 0.0))
            var v24 int64 = (v18 << ((519) & 63))
            _ = fi22
            _ = v23
            _ = v24
        }
        _ = fi20
        _ = v21
    }
    var v25 int = 128
    for fi26 := 0; fi26 < 6; fi26++ {
        if true {
            for fi27 := 0; fi27 < 13; fi27++ {
                fmt.Println(v25, v18)
                var v28 uint = 255
                _ = fi27
                _ = v28
            }
            v29 := GenS5{F0: 42}
            _ = v29
        }
        if len(v19) > 0 {
            v19[0] = v25
        }
        var v30 uint32 = 457
        _ = fi26
        _ = v30
    }
    for fi31 := 0; fi31 < 15; fi31++ {
        v32 := GenS2{F0: 0, F1: v4, F2: v3}
        v33 := GenS2{F0: -1, F1: v4, F2: -19}
        _ = fi31
        _ = v32
        _ = v33
    }
    v4 = v4
    for fi34 := 0; fi34 < 19; fi34++ {
        if false {
            v18 += (v18 - (v18 & 42))
        } else {
            _ = len(v19)
        }
        _ = fi34
    }
    v35 := GenS2{F0: genFn0(false), F1: "\n\t", F2: (v3 + v3)}
    v36 := GenS6{F0: v3}
    v37 := GenS3{F0: 65535}
    v1 -= (v1 << ((1) & 63))
    var v38 int32 = 165
    v4 = v4
    for fi39 := 0; fi39 < 15; fi39++ {
        switch "hello" {
        case "hello":
            fmt.Println(v3)
        case "\n\t":
            v40 := GenS6{F0: v3}
            _ = v40
        case "":
            var v41 int8 = -1
            _ = v41
        default:
            var v42 float64 = (1.5 / 1.5)
            _ = v42
        }
        if true {
            var v43 uint64 = v1
            _ = v43
        }
        _ = fi39
    }
    if len(v19) > 0 {
        v19[0] = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v25
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v4
}
