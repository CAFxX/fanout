package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int32; F1 uint64 }
type GenS4 struct { F0 int64; F1 int; F2 string }
type GenS5 struct { F0 uint; F1 int64; F2 int }
type GenS6 struct { F0 string; F1 bool; F2 int16 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 uint8; F1 int32 }

func main() {
    var v1 float32 = float32((((3.14159 - 1e10) / (float64((-2.25 / 0.1)) * 1e10)) * 0.0))
    var v2 float32 = v1
    var v3 float32 = (v2 + ((((((3.14159 / (float32((((1e10 + -2.25) * ((0.1 + 3.14159) * 1.5)) - -2.25)) * 3.14159)) - 1e10) - ((v1 - 0.1) - v1)) * ((3.14159 * v1) * (v2 + 3.14159))) - ((v1 / ((float32(float64(v2)) - (((-2.25 * v2) / v2) + v2)) / 3.14159)) - float32(3.14159))) / -2.25))
    v4 := []int{}
    var v5 int8 = -128
    v6 := make(map[string]uint32)
    v2 -= v1
    for fi7 := 0; fi7 < 5; fi7++ {
        if true {
            var v8 int8 = (v5 + 1)
            _ = v8
        } else {
            switch (fi7 * fi7) {
            case 9223372036854775807:
                for fi9 := 0; fi9 < 6; fi9++ {
                    v10 := GenS1{F0: 1}
                    var v11 bool = false
                    var v12 uint8 = 193
                    _ = fi9
                    _ = v10
                    _ = v11
                    _ = v12
                }
            case 128:
                var v13 int = (fi7 ^ (fi7 - fi7))
                _ = v13
            default:
                var v14 float64 = -2.25
                _ = v14
            }
        }
        fmt.Println(v5)
        _ = len(v4)
        _ = fi7
    }
    var v15 float32 = 3.14159
    v16 := GenS2{F0: 42}
    if false {
        _ = len(v6)
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    var v17 int64 = 779
    for fi18 := 0; fi18 < 19; fi18++ {
        v19 := GenS3{F0: -286, F1: 2}
        for fi20 := 0; fi20 < 10; fi20++ {
            var v21 int64 = v17
            _ = fi20
            _ = v21
        }
        fmt.Println(v3)
        _ = fi18
        _ = v19
    }
    v22 := GenS2{F0: 42}
    v3 += v1
    for fi23 := 0; fi23 < 5; fi23++ {
        for fi24 := 0; fi24 < 8; fi24++ {
            var v25 int64 = 1
            _ = fi24
            _ = v25
        }
        var v26 uint64 = 165
        _ = fi23
        _ = v26
    }
    if false {
        if false {
            v27 := GenS4{F0: 1, F1: 32768, F2: (("\n\t" + "hello") + (("hello" + "\n\t") + ("foo bar" + "\n\t")))}
            v17 = (v17 * v17)
            switch "a" {
            case "":
                var v28 string = "foo bar"
                _ = v28
            case "a":
                var v29 int16 = 42
                _ = v29
            case "hello":
                v30 := GenS5{F0: 1, F1: v17, F2: 32767}
                _ = v30
            default:
                v31 := GenS6{F0: "\n\t", F1: (!((801 >= 2147483648) && false)), F2: -406}
                _ = v31
            }
            _ = v27
        }
    }
    var v32 int = 32768
    switch "\n\t" {
    case "a":
        var v33 uint16 = 65535
        _ = v33
    case "":
        for fi34 := 0; fi34 < 16; fi34++ {
            v32 = fi34
            _ = len(v6)
            _ = fi34
        }
    case "\n\t":
        var v35 int8 = v5
        _ = v35
    }
    var v36 int32 = -1
    for fi37 := 0; fi37 < 11; fi37++ {
        var v38 float32 = float32((1e10 + -2.25))
        for fi39 := 0; fi39 < 1; fi39++ {
            v40 := GenS6{F0: ("hello" + "a"), F1: true, F2: 32767}
            v1 = float32((1e10 / (1.5 + float64(v15))))
            _ = fi39
            _ = v40
        }
        _ = fi37
        _ = v38
    }
    if (!((v17 ^ (v17 + (v17 / (((v17 * (v17 >> ((93) & 63)))) | 1)))) >= v17)) {
        for fi41 := 0; fi41 < 9; fi41++ {
            v5 = (v5 >> ((65535) & 63))
            for fi42 := 0; fi42 < 7; fi42++ {
                switch "a" {
                case "":
                    switch ((("foo bar" + (("foo bar" + "") + "\n\t")) + "\n\t") + ("\n\t" + ((("a" + "\n\t") + ("foo bar" + ("foo bar" + ("foo bar" + "foo bar")))) + "\n\t"))) {
                    case "foo bar":
                        v1 += ((0.1 - v2) * 0.0)
                    case "":
                        v43 := GenS7{F0: (v36 | int32(v17))}
                        fmt.Println(v36, v17, v1)
                        v44 := GenS7{F0: (v36 * v36)}
                        _ = v43
                        _ = v44
                    case "\n\t":
                        fmt.Println(v5, v15, v3)
                    default:
                        var v45 int32 = 42
                        _ = v45
                    }
                default:
                    fmt.Println(fi41)
                }
                v46 := GenS8{F0: 128, F1: (v36 + (v36 % (((v36 / (((^(v36))) | 1))) | 1)))}
                v4 = append(v4, v32)
                _ = fi42
                _ = v46
            }
            var v47 int16 = 0
            _ = fi41
            _ = v47
        }
        v4 = append(v4, (v32 + (v32 + 42)))
    }
    if len(v4) > 0 {
        v4[0] = v32
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v22
    _ = v3
    _ = v32
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
