package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint; F2 bool }
type GenS2 struct { F0 string; F1 uint32 }
type GenS3 struct { F0 uint8; F1 bool }
type GenS4 struct { F0 int32; F1 uint64; F2 int }
type GenS5 struct { F0 int; F1 uint32 }
type GenS6 struct { F0 int; F1 int16; F2 uint }
func genFn0() bool {
    if (1 > 32767) {
        var v6 int32 = -1
        switch "" {
        case "a":
            switch 1 {
            case 1:
                v6 -= v6
            case -9223372036854775808:
                if (122 < -128) {
                    if (("foo bar" + ("a" + "")) > (((("foo bar" + "") + "hello") + "a") + (("" + ((("hello" + "\n\t") + "a") + "hello")) + ("a" + "foo bar")))) {
                        if true {
                            v7 := GenS1{F0: -9223372036854775808, F1: 65535, F2: (-88 < -70)}
                            var v8 int = 32767
                            v8 = v8
                            _ = v7
                            _ = v8
                        }
                    }
                } else {
                    v9 := []int{-1, 603, -637}
                    for fi10 := 0; fi10 < 15; fi10++ {
                        v6 -= v6
                        _ = fi10
                    }
                    _ = v9
                }
            }
        case "hello":
            switch "" {
            case "a":
                var v11 uint8 = 206
                _ = v11
            default:
                var v12 uint = 42
                _ = v12
            }
        default:
            if false {
                v6 = v6
                var v13 float64 = float64((((1e10 + 1.5) + (float32(-2.25) * float32(0.0))) / ((((3.14159 + 1e10) * 1e10) / (((3.14159 + ((3.14159 + 1e10) + 1e10)) / 1.0) * -2.25)) + (float32(1.5) - float32(-2.25)))))
                _ = v13
            }
        }
        for fi14 := 0; fi14 < 6; fi14++ {
            var v15 uint64 = 506
            v16 := GenS2{F0: "a", F1: 2147483647}
            var v17 float64 = ((float64(((((1.5 + (3.14159 - 3.14159)) * 1e10) / (3.14159 / 1.5)) / 3.14159)) * float64((1.5 + -2.25))) / 0.1)
            _ = fi14
            _ = v15
            _ = v16
            _ = v17
        }
        _ = v6
    }
    var v18 bool = (!false)
    fmt.Println(v18)
    _ = v18
    return ((!(978 > 4294967295)) && (!v18))
}
func genFn1(p0 uint) uint16 {
    p0 += p0
    v19 := GenS1{F0: 9223372036854775807, F1: (p0 + p0), F2: true}
    p0 = p0
    _ = p0
    _ = v19
    return 1
}

func main() {
    var v1 uint8 = 2
    var v2 uint = 0
    var v3 uint32 = 2147483647
    var v4 uint16 = 68
    var v5 uint64 = uint64(v2)
    var v20 uint32 = v3
    var v21 uint64 = (v5 >> ((4294967295) & 63))
    v22 := GenS2{F0: "\n\t", F1: (v20 % (((v3 + v20)) | 1))}
    v23 := []int{-9223372036854775808}
    switch -218 {
    case -1:
        v24 := GenS3{F0: 42, F1: false}
        _ = v24
    case 9223372036854775807:
        fmt.Println(v5, v3)
    case -610:
        switch -9223372036854775808 {
        case 128:
            v25 := GenS4{F0: int32(v5), F1: v5, F2: 128}
            _ = v25
        default:
            var v26 int64 = int64(v4)
            _ = v26
        }
    default:
        v27 := GenS5{F0: -953, F1: (v3 ^ (v20 << ((266) & 63)))}
        _ = v27
    }
    var v28 float32 = float32(3.14159)
    var v29 int8 = -113
    if true {
        var v30 bool = false
        var v31 uint8 = (v1 - (v1 ^ v1))
        v32 := GenS2{F0: ("a" + "a"), F1: v20}
        _ = v30
        _ = v31
        _ = v32
    } else {
        v2 = 1
        var v33 float64 = (0.0 + ((0.0 - 0.0) + 3.14159))
        _ = v33
    }
    var v34 string = ("foo bar" + "\n\t")
    _ = len(v23)
    v35 := GenS6{F0: 2147483647, F1: 0, F2: uint(v21)}
    var v36 int = 32768
    var v37 int = (v36 ^ (-(v36)))
    v29 -= v29
    var v38 int64 = -9223372036854775808
    v5 += v21
    for fi39 := 0; fi39 < 5; fi39++ {
        var v40 float32 = v28
        if (v38 != -794) {
            v36 += v37
            v1 = v1
        } else {
            switch "foo bar" {
            case "":
                if len(v23) >= 2 {
                    v23 = v23[1:len(v23)]
                }
            case "\n\t":
                var v41 uint16 = 500
                _ = v41
            default:
                var v42 uint = v2
                _ = v42
            }
            var v43 int64 = (v38 / ((v38) | 1))
            var v44 int8 = (v29 ^ v29)
            _ = v43
            _ = v44
        }
        for fi45 := 0; fi45 < 19; fi45++ {
            v40 -= (v28 - -2.25)
            var v46 int64 = -1
            var v47 int = v37
            _ = fi45
            _ = v46
            _ = v47
        }
        _ = fi39
        _ = v40
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v28
    _ = v29
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v5
}
