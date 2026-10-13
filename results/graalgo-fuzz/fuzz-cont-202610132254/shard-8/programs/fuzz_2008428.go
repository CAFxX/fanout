package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
type GenS2 struct { F0 uint; F1 int32 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int; F1 int32; F2 uint8 }
type GenS5 struct { F0 int64; F1 uint16 }
func genFn0(p0 bool) uint32 {
    switch 42 {
    case -2147483648:
        for fi5 := 0; fi5 < 4; fi5++ {
            fmt.Println(p0)
            if p0 {
                var v6 float64 = 0.0
                _ = v6
            }
            _ = fi5
        }
    case -780:
        switch "" {
        case "hello":
            p0 = p0
        case "\n\t":
            p0 = true
        }
    }
    v7 := GenS1{F0: 42, F1: 330}
    _ = p0
    _ = v7
    return 132
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int16 = 32767
    var v3 int16 = (-(v2))
    v4 := []int{420, 1, 4294967295, 229}
    fmt.Println(v2)
    for fi8 := 0; fi8 < 5; fi8++ {
        fmt.Println(v2, v3, v2)
        var v9 string = ("" + "a")
        switch (v9 + "hello") {
        case "foo bar":
            for fi10 := 0; fi10 < 12; fi10++ {
                v3 = (v3 / ((v2) | 1))
                v3 -= int16(fi8)
                _ = fi10
            }
        case "a":
            var v11 uint16 = 32767
            _ = v11
        default:
            v12 := GenS1{F0: 0, F1: 42}
            _ = v12
        }
        _ = fi8
        _ = v9
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if false {
        if true {
            var v13 uint = 18446744073709551615
            _ = v13
        } else {
            v2 += (v3 ^ 32767)
        }
        v14 := GenS2{F0: 338, F1: 0}
        v3 = 0
        _ = v14
    } else {
        v15 := GenS2{F0: 0, F1: 0}
        if true {
            switch 2147483647 {
            case 42:
                v2 -= (v2 * v2)
            case 287:
                var v16 float64 = -2.25
                _ = v16
            default:
                var v17 string = "hello"
                _ = v17
            }
            var v18 int8 = -98
            _ = v18
        } else {
            for fi19 := 0; fi19 < 17; fi19++ {
                v20 := GenS3{F0: 100}
                var v21 int64 = 1
                _ = fi19
                _ = v20
                _ = v21
            }
        }
        var v22 int16 = (v2 % ((v3) | 1))
        _ = v15
        _ = v22
    }
    if (42 >= 255) {
        switch -59 {
        case 541:
            v23 := GenS1{F0: 745, F1: 2}
            _ = v23
        case -89:
            var v24 uint8 = 89
            _ = v24
        case 2147483647:
            fmt.Println(v3, v3)
        default:
            var v25 uint64 = 0
            _ = v25
        }
        var v26 uint16 = 2
        v26 = v26
        _ = v26
    } else {
        var v27 int8 = 1
        _ = v27
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v28 := GenS4{F0: 127, F1: 745, F2: 11}
    var v29 uint8 = 255
    for fi30 := 0; fi30 < 4; fi30++ {
        for fi31 := 0; fi31 < 2; fi31++ {
            v32 := GenS1{F0: 42, F1: 192}
            var v33 int16 = (v2 - v3)
            v3 -= v33
            _ = fi31
            _ = v32
            _ = v33
        }
        v4 = append(v4, 2147483647)
        fmt.Println(fi30)
        _ = fi30
    }
    v1[65535] = 42
    v2 -= (v3 >> ((2) & 63))
    switch -550 {
    case 42:
        if true {
            v3 = 0
            var v34 uint64 = 0
            _ = v34
        } else {
            fmt.Println(v3)
        }
    }
    for fi35 := 0; fi35 < 18; fi35++ {
        v2 = (v2 >> ((0) & 63))
        _ = fi35
    }
    v36 := GenS5{F0: 42, F1: 32767}
    v37 := GenS2{F0: 18446744073709551615, F1: -286}
    var v38 float32 = 0.0
    v4 = append(v4, -952)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v4
}
