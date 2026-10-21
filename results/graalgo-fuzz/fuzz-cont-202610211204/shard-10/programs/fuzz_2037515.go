package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8; F2 uint32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 bool }
func genFn0(p0 int64) int64 {
    fmt.Println(p0)
    var v7 int64 = (^(p0))
    _ = p0
    _ = v7
    return 833
}
func genFn1(p0 bool, p1 bool) int {
    var v8 int16 = 32767
    var v9 string = ("hello" + "foo bar")
    var v10 uint8 = 1
    v8 = (v8 / (((v8 ^ v8)) | 1))
    _ = p0
    _ = p1
    _ = v10
    _ = v8
    _ = v9
    return 128
}

func main() {
    var v1 float32 = (float32(((0.0 + (1.5 + ((-2.25 + -2.25) - 3.14159))) * float64((0.1 - 0.0)))) * 1.5)
    v2 := make(map[string]bool)
    var v3 int64 = -480
    var v4 int = 1
    var v5 bool = (1 > -2147483648)
    var v6 uint = 395
    if false {
        if v5 {
            if v5 {
                var v11 int32 = 42
                var v12 uint64 = 18446744073709551615
                v13 := GenS1{F0: v12, F1: 1, F2: 1}
                _ = v11
                _ = v12
                _ = v13
            } else {
                v5 = false
                for fi14 := 0; fi14 < 9; fi14++ {
                    var v15 uint = 0
                    _ = fi14
                    _ = v15
                }
            }
            switch ("a" + "\n\t") {
            case "hello":
                v6 = (v6 | (^(v6)))
            }
        } else {
            fmt.Println(v4)
            if (v4 < (v4 << ((v6) & 63))) {
                v4 -= v4
            }
        }
    } else {
        fmt.Println(v4, v5)
        switch "" {
        case "\n\t":
            v6 -= (^(v6))
        default:
            var v16 int8 = int8(v4)
            _ = v16
        }
        if (v5 || ((1 > -2147483648) || false)) {
            v6 = 617
            v5 = false
            v17 := []int{genFn1((561 > 566), (v5 || (!v5)))}
            _ = v17
        }
    }
    v4 += genFn1(v5, v5)
    v18 := GenS1{F0: 42, F1: 187, F2: 42}
    switch "\n\t" {
    case "hello":
        v2["hello"] = (18446744073709551615 == 1)
    case "\n\t":
        v19 := []int{genFn1(false, v5), (v4 + -376), 128}
        _ = v19
    case "foo bar":
        v20 := GenS1{F0: 2, F1: 112, F2: 2147483647}
        _ = v20
    default:
        if true {
            switch "hello" {
            case "\n\t":
                _ = len(v2)
            case "a":
                fmt.Println(v3)
            }
        }
    }
    var v21 uint8 = 180
    if v5 {
        var v22 uint64 = 760
        v23 := GenS1{F0: v22, F1: 168, F2: uint32(v6)}
        _ = v22
        _ = v23
    } else {
        if v5 {
            if (((-2.25 - 1e10) * (float64((((0.1 * (1e10 + ((-2.25 * (v1 + (float32(-2.25) - ((v1 * ((v1 / float32(3.14159)) + v1)) * v1)))) * float32(3.14159)))) + ((v1 - v1) - (float32(float64((v1 * v1))) * (((((0.1 + (float32(0.0) + (v1 / (v1 - ((1e10 + ((((float32((1.5 * (float64(1e10) * (float64((v1 + -2.25)) * 1e10)))) + (v1 * v1)) * -2.25) / ((v1 * 0.0) * float32(float64(1e10)))) * v1)) / (v1 - (0.1 * float32(0.1)))))))) / float32(1e10)) - (v1 / float32((-2.25 + 1e10)))) + (float32((1.5 * ((1e10 * float64(0.1)) - 0.0))) - v1)) - 0.0)))) * 1e10)) * 0.0)) != (1e10 / 1.0)) {
                if (!(1000 == -950)) {
                    v5 = (!(true && v5))
                    var v24 float32 = ((((v1 + float32(3.14159)) + (v1 * (v1 - float32(1.5)))) * (v1 + (v1 + (v1 - (float32((0.1 + (3.14159 / 1e10))) - ((v1 / (v1 - 1.5)) * 1.5)))))) + 3.14159)
                    _ = v24
                } else {
                    v25 := GenS2{F0: 0}
                    v26 := []int{(-(v4)), v4, 869}
                    _ = v25
                    _ = v26
                }
            }
            v27 := []int{}
            _ = v27
        } else {
            if v5 {
                fmt.Println(v4)
                v28 := v2["foo bar"]
                _ = v28
            }
        }
    }
    v29 := GenS3{F0: v5}
    var v30 bool = (-1 == int16(v21))
    for fi31 := 0; fi31 < 9; fi31++ {
        for fi32 := 0; fi32 < 16; fi32++ {
            switch "a" {
            case "\n\t":
                v33 := GenS2{F0: 302}
                _ = v33
            case "":
                v34 := GenS4{F0: v6}
                _ = v34
            case "a":
                v3 = (-(v3))
            }
            v35 := GenS5{F0: (int(v21) > (fi31 & (^(v4))))}
            _ = fi32
            _ = v35
        }
        v3 = v3
        var v36 uint = v6
        _ = fi31
        _ = v36
    }
    v4 = genFn1((((("a" + "") + "") + "") != "a"), (-402 <= 0))
    switch (("hello" + "\n\t") + "") {
    case "a":
        v37 := []int{(v4 % (((v4 & (-(v4)))) | 1)), (v4 * genFn1(v30, false)), (^(v4)), int(v6)}
        _ = v37
    }
    if (v5 || (!v5)) {
        if ((255 == 65535) && (!v5)) {
            v38 := GenS3{F0: ((3 > 65535) && (3.14159 <= (1e10 + float64(v1))))}
            switch -9223372036854775808 {
            case -9223372036854775808:
                for fi39 := 0; fi39 < 15; fi39++ {
                    for fi40 := 0; fi40 < 1; fi40++ {
                        var v41 int16 = 729
                        switch "foo bar" {
                        case "hello":
                            v3 -= v3
                        default:
                            v42 := GenS1{F0: 893, F1: (v21 ^ uint8(v6)), F2: 2147483648}
                            _ = v42
                        }
                        _ = fi40
                        _ = v41
                    }
                    var v43 int64 = 0
                    _ = fi39
                    _ = v43
                }
            case -2147483648:
                for fi44 := 0; fi44 < 20; fi44++ {
                    var v45 int64 = genFn0(v3)
                    var v46 float64 = 1.5
                    v5 = false
                    _ = fi44
                    _ = v45
                    _ = v46
                }
            }
            fmt.Println(v21, v4, v5)
            _ = v38
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v21
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
