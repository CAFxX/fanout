package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int32; F1 int32; F2 int }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint8; F1 int16 }
type GenS6 struct { F0 int32; F1 int16 }
type GenS7 struct { F0 string }
func genFn0(p0 uint32) bool {
    v5 := GenS1{F0: 128}
    v6 := []int{1}
    for fi7 := 0; fi7 < 7; fi7++ {
        var v8 int = (fi7 * (fi7 - (fi7 % (((fi7 << ((2) & 63))) | 1))))
        for fi9 := 0; fi9 < 11; fi9++ {
            var v10 float32 = -2.25
            var v11 uint64 = 363
            if false {
                if (false || (!true)) {
                    v10 = 0.0
                } else {
                    v11 -= v11
                    v12 := GenS1{F0: 0}
                    _ = v12
                }
                v8 = 128
                v8 -= (^(v8))
                if true {
                    v8 += fi9
                    v13 := GenS1{F0: 193}
                    v14 := GenS2{F0: 850}
                    if true {
                        if true {
                            switch fi7 {
                            case -530:
                                if true {
                                    if (297 >= (p0 + (p0 % (((p0 | (p0 >> ((0) & 63)))) | 1)))) {
                                        v8 = fi9
                                        v10 = ((0.0 - (v10 * v10)) - (v10 / (((0.0 - 1e10) + (v10 / 1.5)) * ((v10 / (v10 / 1.0)) * float32(-2.25)))))
                                        fmt.Println(v8)
                                    }
                                    if (748 != 18446744073709551615) {
                                        v15 := GenS1{F0: 61}
                                        if true {
                                            v16 := GenS2{F0: 174}
                                            switch (v8 * -64) {
                                            case -9223372036854775808:
                                                v17 := GenS3{F0: 301, F1: -1, F2: fi9}
                                                _ = v17
                                            case -425:
                                                var v18 int32 = 42
                                                _ = v18
                                            case 32767:
                                                var v19 int8 = 127
                                                _ = v19
                                            default:
                                                switch "" {
                                                case "a":
                                                    v10 = float32(1e10)
                                                case "\n\t":
                                                    p0 = p0
                                                default:
                                                    p0 = 2
                                                }
                                            }
                                            _ = v16
                                        } else {
                                            v10 -= v10
                                            switch (fi9 << ((18446744073709551615) & 63)) {
                                            case 528:
                                                if false {
                                                    v11 = 42
                                                    v20 := GenS4{F0: 2}
                                                    _ = v20
                                                }
                                            }
                                        }
                                        _ = v15
                                    } else {
                                        v21 := GenS4{F0: (v11 * 42)}
                                        var v22 uint8 = 255
                                        _ = v21
                                        _ = v22
                                    }
                                    var v23 uint = 2
                                    _ = v23
                                }
                            default:
                                switch "hello" {
                                case "":
                                    fmt.Println(v11)
                                case "foo bar":
                                    if len(v6) > 0 {
                                        v6[0] = 62
                                    }
                                }
                            }
                        }
                        var v24 uint = 79
                        var v25 string = "hello"
                        _ = v24
                        _ = v25
                    }
                    var v26 uint8 = 255
                    var v27 int32 = 2147483647
                    _ = len(v6)
                    _ = v13
                    _ = v14
                    _ = v26
                    _ = v27
                } else {
                    var v28 uint64 = (v11 + v11)
                    v29 := GenS3{F0: 0, F1: -1, F2: (v8 ^ (v8 >> ((42) & 63)))}
                    var v30 uint16 = 42
                    _ = v28
                    _ = v29
                    _ = v30
                }
                v31 := GenS4{F0: (v11 | 4294967295)}
                _ = v31
            }
            _ = fi9
            _ = v10
            _ = v11
        }
        _ = fi7
        _ = v8
    }
    switch 636 {
    case -777:
        var v32 string = "hello"
        _ = v32
    }
    _ = p0
    _ = v5
    _ = v6
    return true
}
func genFn1() int16 {
    v33 := []int{2147483647}
    var v34 int64 = 9223372036854775807
    var v35 int = 127
    v35 += (v35 & (v35 | 128))
    v36 := GenS5{F0: 255, F1: 0}
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    return 0
}

func main() {
    var v1 string = "\n\t"
    var v2 float32 = (float32((float64(1.5) - float64(((0.1 * 1.5) * (float32(float64(1e10)) + (float32(3.14159) + 1e10)))))) - (float32(0.0) + 0.1))
    var v3 int64 = 9223372036854775807
    var v4 string = (v1 + ("foo bar" + "foo bar"))
    v37 := GenS6{F0: 683, F1: -32768}
    var v38 int64 = v3
    if (255 == 1) {
        v39 := []int{-872, -2147483648, -772, 9223372036854775807}
        _ = v39
    }
    fmt.Println(v3)
    v38 += -9223372036854775808
    for fi40 := 0; fi40 < 14; fi40++ {
        if (float64((1e10 - ((0.0 / v2) * 3.14159))) >= (0.1 * 1.5)) {
            for fi41 := 0; fi41 < 6; fi41++ {
                var v42 float32 = v2
                v43 := []int{(fi40 - fi41), (fi41 % ((0) | 1)), fi41}
                v2 += 3.14159
                _ = fi41
                _ = v42
                _ = v43
            }
            switch (fi40 % ((fi40) | 1)) {
            case -989:
                var v44 int = (fi40 ^ fi40)
                _ = v44
            case 9223372036854775807:
                v45 := []int{(fi40 ^ fi40), -9223372036854775808, fi40, (fi40 / (((fi40 / ((0) | 1))) | 1))}
                _ = v45
            }
        } else {
            switch fi40 {
            case 1:
                v3 = 9223372036854775807
            case 32767:
                var v46 bool = false
                _ = v46
            case 0:
                if false {
                    var v47 uint = 1
                    v3 = v3
                    _ = v47
                } else {
                    v48 := []int{(fi40 & (fi40 % ((fi40) | 1))), fi40, (fi40 ^ 32767), fi40}
                    if len(v48) > 0 {
                        v48[0] = (fi40 << ((18446744073709551615) & 63))
                    }
                    for fi49 := 0; fi49 < 1; fi49++ {
                        switch v1 {
                        case "a":
                            fmt.Println(fi49)
                        case "hello":
                            v2 = (float32(1.5) + v2)
                        }
                        _ = fi49
                    }
                    _ = v48
                }
            default:
                switch v1 {
                case "hello":
                    var v50 uint32 = 724
                    _ = v50
                case "foo bar":
                    fmt.Println(fi40, v4, v4)
                }
            }
            v51 := GenS1{F0: 128}
            for fi52 := 0; fi52 < 8; fi52++ {
                switch ((v4 + v4) + (v1 + v1)) {
                case "a":
                    var v53 uint32 = 2147483647
                    _ = v53
                case "\n\t":
                    if true {
                        v38 = v38
                        var v54 uint8 = 128
                        _ = v54
                    }
                    v4 = (v4 + "")
                    v55 := []int{(fi40 / ((fi40) | 1))}
                    v38 -= (v3 - -1)
                    var v56 int8 = 42
                    _ = v55
                    _ = v56
                case "":
                    var v57 string = (((("a" + v4) + v1) + "a") + v1)
                    _ = v57
                }
                fmt.Println(v4, v4)
                fmt.Println(v4, v4)
                _ = fi52
            }
            _ = v51
        }
        var v58 uint = 42
        _ = fi40
        _ = v58
    }
    v3 = 0
    if (false || (false || (0 > 92))) {
        var v59 int8 = 56
        v60 := GenS7{F0: v4}
        _ = v59
        _ = v60
    }
    for fi61 := 0; fi61 < 14; fi61++ {
        v3 -= -20
        fmt.Println(v3, v38)
        var v62 uint = 254
        _ = fi61
        _ = v62
    }
    v1 = v4
    var v63 int16 = 1
    fmt.Println(v63, v4)
    v3 = v38
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v63
}
