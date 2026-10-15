package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int16; F1 int }
type GenS3 struct { F0 int8; F1 bool; F2 uint64 }
type GenS4 struct { F0 int; F1 int8 }
type GenS5 struct { F0 uint32 }
func genFn0(p0 int) int {
    v6 := []int{(p0 >> ((2) & 63)), 40, p0, p0}
    p0 -= (p0 << ((4294967295) & 63))
    if len(v6) > 0 {
        v6[0] = -1
    }
    fmt.Println(p0)
    _ = p0
    _ = v6
    return p0
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[uint32]string)
    var v3 int32 = 148
    var v4 uint8 = 111
    v5 := make(map[string]int)
    for fi7 := 0; fi7 < 7; fi7++ {
        switch genFn0((fi7 % (((-(fi7))) | 1))) {
        case -258:
            v8 := v1["foo bar"]
            _ = v8
        default:
            switch 2147483647 {
            case 128:
                v4 -= v4
            case 42:
                v9 := GenS1{F0: 65535}
                _ = v9
            }
        }
        var v10 uint8 = v4
        var v11 int32 = v3
        _ = fi7
        _ = v10
        _ = v11
    }
    switch -518 {
    case 0:
        if (0 >= 42) {
            switch "a" {
            case "foo bar":
                var v12 int32 = v3
                _ = v12
            case "":
                var v13 int32 = (v3 * v3)
                _ = v13
            default:
                fmt.Println(v4, v3)
            }
            v2[1] = "foo bar"
            delete(v2, 0)
        }
    case 1:
        v4 -= (v4 / ((128) | 1))
    }
    for fi14 := 0; fi14 < 12; fi14++ {
        if (191 < 7) {
            var v15 float32 = (0.1 + ((((0.1 * (1.5 + 0.1)) + 3.14159) + 1.5) * float32(0.0)))
            _ = v15
        }
        fmt.Println(fi14, fi14, v4)
        for fi16 := 0; fi16 < 17; fi16++ {
            var v17 uint = 433
            switch ("foo bar" + (("" + "foo bar") + "foo bar")) {
            case "a":
                var v18 uint16 = 255
                _ = v18
            }
            if false {
                v4 -= v4
            } else {
                var v19 float64 = (1.5 / 1.0)
                v20 := GenS2{F0: 633, F1: -9223372036854775808}
                _ = v19
                _ = v20
            }
            _ = fi16
            _ = v17
        }
        _ = fi14
    }
    fmt.Println(v4, v3)
    v3 += (v3 - 1)
    _ = len(v2)
    v4 += (v4 + 42)
    v21 := GenS3{F0: 107, F1: true, F2: 4294967295}
    v22 := []int{946}
    var v23 uint16 = 448
    var v24 int16 = 0
    for fi25 := 0; fi25 < 3; fi25++ {
        if true {
            switch fi25 {
            case 1:
                var v26 uint64 = 0
                _ = v26
            case 127:
                var v27 float32 = 0.0
                _ = v27
            }
            v28 := GenS4{F0: fi25, F1: 0}
            var v29 int8 = 106
            _ = v28
            _ = v29
        } else {
            var v30 uint16 = 2
            for fi31 := 0; fi31 < 17; fi31++ {
                var v32 bool = true
                var v33 int8 = 1
                var v34 int = (fi31 & (fi25 * 0))
                fmt.Println(v24)
                _ = fi31
                _ = v32
                _ = v33
                _ = v34
            }
            _ = v30
        }
        switch genFn0(fi25) {
        case 202:
            v35 := GenS5{F0: uint32(v23)}
            _ = v35
        case -1:
            if len(v22) >= 2 {
                v22 = v22[1:len(v22)]
            }
        }
        var v36 int8 = -128
        _ = fi25
        _ = v36
    }
    fmt.Println(v23)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
