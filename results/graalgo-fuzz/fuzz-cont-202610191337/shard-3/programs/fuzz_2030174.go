package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 int }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int8; F1 int64 }
type GenS4 struct { F0 uint16; F1 string; F2 int }
type GenS5 struct { F0 uint16 }
func genFn0(p0 string, p1 int32, p2 int64) int16 {
    fmt.Println(p0, p2, p2)
    fmt.Println(p1, p2)
    v6 := []int{-8, 128, 42, 9223372036854775807}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 0
}
func genFn1() uint32 {
    var v7 uint32 = 0
    for fi8 := 0; fi8 < 19; fi8++ {
        if (2147483647 > 1) {
            v7 += (v7 ^ (v7 | v7))
        }
        v9 := GenS1{F0: 0, F1: 1, F2: fi8}
        v10 := GenS2{F0: 1000}
        _ = fi8
        _ = v10
        _ = v9
    }
    _ = v7
    return (v7 >> ((0) & 63))
}

func main() {
    var v1 bool = (-1 != -128)
    var v2 int64 = 1
    var v3 int = 633
    var v4 uint16 = 2
    var v5 float32 = (((0.1 * 1.5) * 0.1) / 0.1)
    switch (v3 | v3) {
    case -953:
        var v11 int = -233
        _ = v11
    case -438:
        var v12 uint8 = 87
        _ = v12
    case 128:
        for fi13 := 0; fi13 < 7; fi13++ {
            for fi14 := 0; fi14 < 1; fi14++ {
                var v15 uint16 = 466
                _ = fi14
                _ = v15
            }
            _ = fi13
        }
    }
    for fi16 := 0; fi16 < 19; fi16++ {
        v17 := GenS2{F0: genFn0(("hello" + "foo bar"), 774, v2)}
        _ = fi16
        _ = v17
    }
    var v18 uint32 = uint32(v4)
    var v19 uint = 812
    if ((100 > 127) && (!(100 <= 0))) {
        v20 := []int{}
        for fi21 := 0; fi21 < 6; fi21++ {
            switch (v3 ^ fi21) {
            case 0:
                fmt.Println(v18, v18, v1)
            case -204:
                var v22 uint8 = 0
                _ = v22
            default:
                for fi23 := 0; fi23 < 8; fi23++ {
                    fmt.Println(v2)
                    fmt.Println(v19)
                    _ = fi23
                }
            }
            _ = fi21
        }
        _ = v20
    } else {
        var v24 uint = v19
        _ = v24
    }
    v3 -= (v3 + v3)
    v18 += (v18 >> (((v19 >> ((v19) & 63))) & 63))
    for fi25 := 0; fi25 < 20; fi25++ {
        fmt.Println(v18, v5)
        v26 := GenS1{F0: (v19 - 790), F1: 0, F2: (fi25 & (v3 >> ((18446744073709551615) & 63)))}
        v4 -= uint16(v19)
        _ = fi25
        _ = v26
    }
    fmt.Println(v1, v19)
    if (((v5 / float32(1.5)) / 3.14159) != ((v5 / v5) + float32(float64((v5 / v5))))) {
        switch v3 {
        case 696:
            fmt.Println(v4, v18, v18)
        }
        switch ("hello" + (("\n\t" + ("foo bar" + "")) + ("\n\t" + ("" + "a")))) {
        case "":
            for fi27 := 0; fi27 < 20; fi27++ {
                switch 42 {
                case -9223372036854775808:
                    v5 -= 1e10
                case 1:
                    fmt.Println(fi27)
                case 4294967295:
                    var v28 int64 = -1
                    _ = v28
                default:
                    fmt.Println(v1, v4, v1)
                }
                v29 := GenS3{F0: -1, F1: (v2 / ((v2) | 1))}
                v2 = (v2 ^ (-(v2)))
                _ = fi27
                _ = v29
            }
        case "\n\t":
            v19 -= v19
        }
    } else {
        for fi30 := 0; fi30 < 20; fi30++ {
            var v31 int64 = v2
            v19 += v19
            var v32 uint16 = (v4 / (((v4 + 862)) | 1))
            _ = fi30
            _ = v31
            _ = v32
        }
        var v33 bool = (!v1)
        _ = v33
    }
    for fi34 := 0; fi34 < 15; fi34++ {
        v35 := []int{}
        v35 = append(v35, (v3 << (((^(v19))) & 63)))
        fmt.Println(v5, v19, v19)
        _ = fi34
        _ = v35
    }
    switch 9223372036854775807 {
    case 42:
        if (genFn1() != genFn1()) {
            for fi36 := 0; fi36 < 20; fi36++ {
                v18 = v18
                var v37 float32 = v5
                _ = fi36
                _ = v37
            }
            v3 -= (v3 | -1)
            fmt.Println(v18, v4, v19)
        }
    case 127:
        v38 := GenS4{F0: (v4 - v4), F1: "a", F2: v3}
        _ = v38
    case 391:
        v39 := GenS5{F0: (v4 ^ 255)}
        _ = v39
    default:
        var v40 uint16 = v4
        _ = v40
    }
    v41 := []int{(v3 & (v3 * (v3 % ((v3) | 1))))}
    if len(v41) >= 2 {
        v41 = v41[1:len(v41)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v41:", len(v41), v41)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v5
}
