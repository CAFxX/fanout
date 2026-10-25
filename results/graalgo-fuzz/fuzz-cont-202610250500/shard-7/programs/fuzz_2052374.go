package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 uint32 }
type GenS2 struct { F0 int8; F1 int8 }
type GenS3 struct { F0 int; F1 int16; F2 uint }
type GenS4 struct { F0 int64; F1 string; F2 int8 }
func genFn0(p0 int16) bool {
    for fi7 := 0; fi7 < 1; fi7++ {
        if ((128 <= 127) || false) {
            p0 += p0
        } else {
            p0 -= p0
            switch "foo bar" {
            case "hello":
                v8 := []int{(fi7 ^ -924), 670}
                _ = v8
            case "\n\t":
                var v9 float32 = (float32(0.0) - (1e10 + (((float32(1e10) + (3.14159 * 3.14159)) * ((0.0 + (3.14159 + (3.14159 / -2.25))) / 1.0)) / 1.0)))
                _ = v9
            }
            var v10 uint = 65535
            _ = v10
        }
        _ = fi7
    }
    switch ("" + "hello") {
    case "hello":
        for fi11 := 0; fi11 < 11; fi11++ {
            if (255 >= 255) {
                if (9223372036854775807 > 0) {
                    switch fi11 {
                    case 32768:
                        v12 := GenS1{F0: 501, F1: 0, F2: 213}
                        _ = v12
                    case -307:
                        var v13 uint = 2
                        _ = v13
                    }
                }
                v14 := GenS1{F0: 756, F1: 2, F2: 2}
                _ = v14
            } else {
                switch 32767 {
                case 9223372036854775807:
                    v15 := GenS2{F0: -128, F1: -128}
                    _ = v15
                case -1:
                    v16 := []int{1, -272, fi11}
                    _ = v16
                case 128:
                    var v17 float64 = 0.1
                    _ = v17
                }
                fmt.Println(fi11, p0)
                switch "a" {
                case "hello":
                    var v18 int64 = 280
                    _ = v18
                case "":
                    p0 -= (p0 ^ p0)
                }
            }
            p0 -= (p0 << ((85) & 63))
            v19 := []int{fi11}
            _ = fi11
            _ = v19
        }
    case "a":
        v20 := GenS3{F0: -9223372036854775808, F1: p0, F2: 255}
        _ = v20
    case "":
        if (-2147483648 <= 42) {
            if (42 != 0) {
                if true {
                    var v21 int64 = 42
                    v21 -= (v21 / (((v21 ^ (v21 % ((v21) | 1)))) | 1))
                    v21 += 0
                    _ = v21
                }
                fmt.Println(p0)
            }
            v22 := GenS4{F0: 42, F1: "\n\t", F2: 0}
            _ = v22
        } else {
            var v23 int = -549
            if (0.1 >= (0.0 * (0.0 / (3.14159 - 0.0)))) {
                v23 = (v23 ^ (v23 >> ((275) & 63)))
                v24 := []int{-615, (v23 * (v23 / (((^(v23))) | 1))), v23}
                var v25 bool = true
                _ = v24
                _ = v25
            }
            _ = v23
        }
    default:
        fmt.Println(p0)
    }
    if (!("foo bar" <= ((("foo bar" + (((("a" + "foo bar") + "hello") + "hello") + ("\n\t" + "hello"))) + "foo bar") + ""))) {
        p0 = (p0 | 1)
    } else {
        v26 := GenS1{F0: 1, F1: 766, F2: 42}
        switch "a" {
        case "a":
            if (42 != 0) {
                var v27 int64 = 1
                _ = v27
            } else {
                v28 := []int{-2147483648, -9223372036854775808}
                var v29 int64 = 0
                var v30 uint16 = 42
                _ = v28
                _ = v29
                _ = v30
            }
        case "":
            p0 = (^(p0))
        case "foo bar":
            var v31 string = ((("foo bar" + "foo bar") + "a") + "\n\t")
            _ = v31
        }
        fmt.Println(p0)
        _ = v26
    }
    v32 := []int{405}
    _ = p0
    _ = v32
    return false
}
func genFn1(p0 int32) string {
    fmt.Println(p0)
    var v33 int64 = int64(p0)
    _ = p0
    _ = v33
    return ""
}

func main() {
    var v1 string = "hello"
    var v2 float64 = 0.0
    v3 := []int{-2147483648, 2147483647, 32767, 4294967295}
    var v4 uint = 466
    v5 := []string{v1, v1}
    v6 := make(map[uint32]bool)
    var v34 int16 = 0
    for fi35 := 0; fi35 < 8; fi35++ {
        var v36 float64 = (v2 / (v2 / (-2.25 * -2.25)))
        _ = fi35
        _ = v36
    }
    v34 += v34
    for fi37 := 0; fi37 < 18; fi37++ {
        v38 := GenS1{F0: 65535, F1: 0, F2: 4294967295}
        var v39 int8 = 0
        for fi40 := 0; fi40 < 20; fi40++ {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = fi40
        }
        _ = fi37
        _ = v38
        _ = v39
    }
    v34 = v34
    v34 = v34
    var v41 uint32 = 306
    switch 1 {
    case 4294967295:
        var v42 uint8 = 255
        _ = v42
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v41
    _ = v5
    _ = v6
}
