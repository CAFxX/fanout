package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint; F2 int32 }
type GenS2 struct { F0 int16; F1 uint; F2 uint8 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int8; F1 int32 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 string; F1 uint; F2 uint16 }
type GenS7 struct { F0 int8; F1 uint16; F2 string }
func genFn0(p0 int, p1 string, p2 int8) uint32 {
    p1 = (p1 + (p1 + p1))
    v4 := GenS1{F0: 562, F1: 65535, F2: 42}
    v5 := []int{}
    var v6 int16 = 1
    v7 := GenS2{F0: (v6 << ((157) & 63)), F1: 255, F2: 1}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return 694
}
func genFn1(p0 int64, p1 int8, p2 int) int64 {
    v8 := []int{p2, p2}
    if false {
        p2 += (p2 >> ((18446744073709551615) & 63))
        var v9 uint8 = 2
        p0 = (p0 - p0)
        _ = v9
    } else {
        fmt.Println(p2, p0)
        switch p2 {
        case 42:
            var v10 float64 = float64((float32((1.5 / ((-2.25 / (1.5 / (float64((1.5 - 1e10)) + (0.1 + 0.0)))) / 0.1))) - (float32((float64(0.0) - 3.14159)) * 3.14159)))
            _ = v10
        }
        p1 -= (p1 ^ (p1 / (((p1 & -128)) | 1)))
    }
    v11 := GenS3{F0: 732}
    var v12 bool = (((3.14159 * 1e10) >= -2.25) && false)
    switch (p2 % ((p2) | 1)) {
    case 42:
        if v12 {
            v12 = v12
            p1 += (p1 * p1)
            if true {
                fmt.Println(v12, p0, p1)
            }
        } else {
            var v13 int8 = -27
            _ = v13
        }
    case 1:
        switch (("a" + "") + (("" + "a") + ("\n\t" + "foo bar"))) {
        case "\n\t":
            var v14 int64 = p0
            _ = v14
        case "hello":
            var v15 float64 = ((-2.25 / 0.1) + 1e10)
            _ = v15
        default:
            var v16 float64 = 1e10
            _ = v16
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    _ = v8
    return 0
}

func main() {
    v1 := []string{"\n\t", ("\n\t" + "hello")}
    var v2 int64 = 0
    var v3 uint32 = 440
    v3 -= (v3 >> ((608) & 63))
    fmt.Println(v2, v3)
    var v17 uint = 4294967295
    v2 += v2
    fmt.Println(v3)
    v18 := GenS4{F0: -54, F1: 0}
    v19 := GenS5{F0: 1}
    v20 := GenS4{F0: -112, F1: 1}
    fmt.Println(v2, v17, v3)
    var v21 uint8 = 128
    for fi22 := 0; fi22 < 16; fi22++ {
        if (!true) {
            fmt.Println(fi22)
            if false {
                v23 := GenS6{F0: ("foo bar" + ""), F1: v17, F2: 997}
                _ = v23
            } else {
                v24 := GenS7{F0: -128, F1: uint16(v3), F2: ""}
                _ = v24
            }
            v21 = 0
        } else {
            for fi25 := 0; fi25 < 20; fi25++ {
                v21 = 127
                var v26 uint = (v17 | v17)
                if ((true || ((3.14159 * 0.1) != 0.1)) && false) {
                    v1 = append(v1, ("foo bar" + ("a" + "\n\t")))
                    var v27 float64 = -2.25
                    v28 := GenS2{F0: 1000, F1: v17, F2: 1}
                    var v29 int = (fi25 % (((^(fi25))) | 1))
                    _ = v27
                    _ = v28
                    _ = v29
                }
                _ = fi25
                _ = v26
            }
            v30 := GenS5{F0: 32767}
            _ = v30
        }
        for fi31 := 0; fi31 < 20; fi31++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            v21 = (v21 - (v21 << (((v17 ^ v17)) & 63)))
            var v32 int64 = genFn1(v2, -1, (fi22 | (fi22 / (((fi22 % ((fi22) | 1))) | 1))))
            v33 := GenS6{F0: "foo bar", F1: v17, F2: 255}
            _ = fi31
            _ = v32
            _ = v33
        }
        _ = fi22
    }
    var v34 uint8 = (v21 | (v21 / (((v21 % (((v21 | 127)) | 1))) | 1)))
    var v35 int32 = 2147483647
    _ = len(v1)
    v34 -= (v21 + 1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v34
    _ = v35
}
