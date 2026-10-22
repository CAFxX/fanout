package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint16; F1 string }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 bool; F1 uint64 }
func genFn0(p0 uint8, p1 uint16, p2 string) int32 {
    for fi4 := 0; fi4 < 8; fi4++ {
        if (2 != 65535) {
            p2 = (("hello" + p2) + "foo bar")
        }
        if false {
            v5 := GenS1{F0: 42}
            var v6 uint32 = 653
            _ = v5
            _ = v6
        }
        if (32767 <= 42) {
            if (-1 >= -762) {
                var v7 bool = false
                var v8 uint64 = 994
                _ = v7
                _ = v8
            }
        }
        _ = fi4
    }
    p1 = 328
    p1 += 437
    v9 := []int{128}
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return -2147483648
}
func genFn1(p0 int16, p1 uint32, p2 int16) uint8 {
    var v10 int32 = -2147483648
    p1 = (^(p1))
    switch 32767 {
    case 61:
        v11 := []int{-1}
        _ = v11
    case -541:
        p0 += -534
    case -359:
        switch 2147483647 {
        case -391:
            var v12 uint = 0
            _ = v12
        case 32768:
            p0 += 609
        default:
            v13 := []int{9223372036854775807, -2147483648, int(p0)}
            _ = v13
        }
    }
    fmt.Println(p1, v10)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return 84
}

func main() {
    v1 := make(map[string]bool)
    var v2 float64 = (-2.25 / 1.5)
    v3 := make(map[uint32]string)
    v2 -= (v2 - 0.0)
    if true {
        v14 := GenS1{F0: 2147483647}
        var v15 uint32 = 4294967295
        v2 = (0.1 + -2.25)
        _ = v14
        _ = v15
    }
    for fi16 := 0; fi16 < 8; fi16++ {
        var v17 uint64 = 416
        for fi18 := 0; fi18 < 4; fi18++ {
            v19 := []int{(-(fi18))}
            if len(v19) >= 2 {
                v19 = v19[1:len(v19)]
            }
            if len(v19) >= 2 {
                v19 = v19[1:len(v19)]
            }
            _ = fi18
            _ = v19
        }
        for fi20 := 0; fi20 < 3; fi20++ {
            var v21 bool = false
            v22 := GenS1{F0: 1}
            _ = fi20
            _ = v21
            _ = v22
        }
        _ = fi16
        _ = v17
    }
    v23 := []int{32768, 157, 232, 32767}
    fmt.Println(v2)
    v2 += ((1e10 + -2.25) - ((1e10 - float64(1.5)) + v2))
    for fi24 := 0; fi24 < 9; fi24++ {
        switch ("\n\t" + "\n\t") {
        case "hello":
            var v25 uint64 = 1
            _ = v25
        case "":
            fmt.Println(v2, fi24)
        case "\n\t":
            v26 := GenS2{F0: 32767, F1: ""}
            _ = v26
        default:
            v27 := GenS3{F0: 127}
            _ = v27
        }
        _ = fi24
    }
    v28 := v3[4294967295]
    fmt.Println(v2)
    v28 = "foo bar"
    v29 := GenS2{F0: 2, F1: v28}
    var v30 uint8 = 128
    _ = len(v23)
    v2 += (v2 * (v2 - 0.1))
    if false {
        var v31 int8 = -128
        _ = v31
    } else {
        v23 = append(v23, 4294967295)
        var v32 float64 = (v2 / v2)
        v33 := GenS4{F0: (962 <= -9223372036854775808), F1: 42}
        _ = v32
        _ = v33
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    _ = v1
    _ = v2
    _ = v23
    _ = v28
    _ = v29
    _ = v3
    _ = v30
}
