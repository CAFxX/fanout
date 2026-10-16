package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 string }
func genFn0(p0 int32, p1 string, p2 uint32) uint32 {
    p1 = "\n\t"
    v7 := []int{}
    var v8 uint32 = p2
    if false {
        p0 -= p0
        if (-481 > 1) {
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
            v7 = append(v7, 1)
            fmt.Println(p1)
        }
        var v9 int32 = -744
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return (p2 - 2)
}
func genFn1(p0 int8) uint16 {
    v10 := []int{-289, 9223372036854775807, 165}
    var v11 int32 = 0
    switch "" {
    case "hello":
        switch -510 {
        case 32767:
            if true {
                var v12 uint32 = 0
                _ = v12
            } else {
                var v13 float64 = 0.1
                _ = v13
            }
        case 127:
            for fi14 := 0; fi14 < 4; fi14++ {
                for fi15 := 0; fi15 < 13; fi15++ {
                    _ = len(v10)
                    _ = fi15
                }
                var v16 float32 = (-2.25 + float32(1.5))
                fmt.Println(v11, fi14)
                _ = fi14
                _ = v16
            }
        case 32768:
            if ("a" > "a") {
                v17 := GenS1{F0: 512, F1: 9223372036854775807}
                _ = v17
            } else {
                v18 := GenS1{F0: (v11 * (^(v11))), F1: 82}
                _ = v18
            }
        default:
            var v19 int8 = p0
            _ = v19
        }
    case "foo bar":
        v20 := GenS1{F0: v11, F1: 42}
        _ = v20
    case "a":
        p0 -= p0
    default:
        var v21 bool = (!false)
        _ = v21
    }
    if (((1.5 / float64(0.1)) >= 0.1) && true) {
        if (((((1e10 * 3.14159) * -2.25) / 0.1) / 1.5) >= 0.1) {
            if len(v10) >= 2 {
                v10 = v10[1:len(v10)]
            }
            fmt.Println(p0, v11)
            v11 = v11
        }
        v11 = 954
    } else {
        switch 224 {
        case 42:
            if true {
                var v22 int8 = (p0 - p0)
                v11 = v11
                _ = v22
            }
        case -2147483648:
            var v23 int32 = -2147483648
            _ = v23
        default:
            switch ("a" + ("hello" + "hello")) {
            case "\n\t":
                fmt.Println(v11, p0)
            case "foo bar":
                v11 += v11
            case "a":
                v24 := GenS2{F0: (v11 < (v11 + v11))}
                _ = v24
            default:
                if len(v10) >= 2 {
                    v10 = v10[1:len(v10)]
                }
            }
        }
    }
    var v25 int = 1
    _ = p0
    _ = v10
    _ = v11
    _ = v25
    return 255
}

func main() {
    v1 := make(map[uint32]int)
    v2 := []int{4294967295, 32767}
    v3 := make(map[string]bool)
    v4 := []uint32{167, 65535, 2, 2}
    var v5 uint = 1
    v6 := make(map[int]string)
    for fi26 := 0; fi26 < 14; fi26++ {
        var v27 int8 = -51
        _ = fi26
        _ = v27
    }
    var v28 int16 = -32768
    var v29 uint32 = 840
    if len(v4) > 0 {
        v4[0] = v29
    }
    switch ("" + "foo bar") {
    case "\n\t":
        v30 := GenS3{F0: v5}
        _ = v30
    case "foo bar":
        v28 = v28
    case "a":
        if (!((160 >= 127) && (!true))) {
            v5 = v5
            for fi31 := 0; fi31 < 3; fi31++ {
                v29 += v29
                for fi32 := 0; fi32 < 12; fi32++ {
                    fmt.Println(v28, v5, fi31)
                    _ = fi32
                }
                _ = fi31
            }
            fmt.Println(v5, v5)
        }
    }
    if (-1 >= 1) {
        switch "foo bar" {
        case "a":
            var v33 uint = 117
            _ = v33
        case "\n\t":
            v34 := GenS4{F0: "foo bar"}
            _ = v34
        case "foo bar":
            for fi35 := 0; fi35 < 19; fi35++ {
                v36 := GenS3{F0: v5}
                _ = fi35
                _ = v36
            }
        default:
            _ = len(v4)
        }
        var v37 float32 = -2.25
        v38 := GenS3{F0: (v5 * (v5 * (v5 | (v5 - (v5 << (((^(v5))) & 63))))))}
        _ = v37
        _ = v38
    }
    for fi39 := 0; fi39 < 4; fi39++ {
        if (((v29 <= 2147483648) && (2 < (v29 ^ genFn0(-317, "foo bar", v29)))) && false) {
            v29 += (v29 / ((v29) | 1))
        }
        if (1 == -1) {
            var v40 bool = (0.1 > 3.14159)
            v41 := GenS3{F0: v5}
            _ = v40
            _ = v41
        }
        _ = fi39
    }
    v5 = 473
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
