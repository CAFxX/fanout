package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64 }
type GenS2 struct { F0 uint8; F1 uint64; F2 uint32 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 bool; F1 int32 }
type GenS5 struct { F0 int16; F1 int }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 uint64; F1 int16; F2 string }

func main() {
    var v1 uint16 = 2
    var v2 int32 = -320
    v3 := make(map[uint32]uint32)
    var v4 float32 = 1e10
    v5 := make(map[uint32]int)
    var v6 int32 = v2
    if (!(!(!false))) {
        for fi7 := 0; fi7 < 10; fi7++ {
            if true {
                if false {
                    var v8 int = fi7
                    _ = v8
                } else {
                    fmt.Println(v2, v2)
                }
                if (("foo bar" > "hello") || true) {
                    if false {
                        var v9 int8 = 127
                        var v10 int8 = (v9 * v9)
                        var v11 string = "\n\t"
                        _ = v10
                        _ = v11
                        _ = v9
                    }
                    switch ("hello" + "hello") {
                    case "hello":
                        if false {
                            switch "foo bar" {
                            case "":
                                fmt.Println(v4, v4, v1)
                            }
                        } else {
                            v6 = (v2 >> ((0) & 63))
                            v12 := GenS1{F0: v2, F1: 42}
                            v13 := GenS1{F0: 1, F1: 9223372036854775807}
                            _ = v12
                            _ = v13
                        }
                    default:
                        if false {
                            switch fi7 {
                            case 622:
                                v14 := GenS2{F0: 18, F1: 2, F2: 4294967295}
                                _ = v14
                            case -9223372036854775808:
                                v15 := GenS3{F0: 42}
                                _ = v15
                            }
                        }
                    }
                } else {
                    var v16 int16 = 42
                    _ = v16
                }
            } else {
                var v17 int32 = v2
                v4 = ((float32(((((0.0 + -2.25) - (float64((-2.25 + v4)) / 1.5)) / -2.25) * float64(v4))) - v4) + v4)
                _ = v17
            }
            switch -1 {
            case -2147483648:
                delete(v5, 2147483647)
            case 4294967295:
                v2 = 50
            default:
                fmt.Println(v2, v2, fi7)
            }
            if true {
                v18 := []int{fi7, (^(fi7)), (-(fi7)), (fi7 % (((^(fi7))) | 1))}
                v2 += (v6 >> ((2) & 63))
                v1 -= v1
                _ = v18
            }
            _ = fi7
        }
        if (!(0 == -32768)) {
            if (0 <= 100) {
                var v19 int = 0
                _ = v19
            }
            var v20 int = 32768
            for fi21 := 0; fi21 < 4; fi21++ {
                fmt.Println(fi21)
                _ = fi21
            }
            _ = v20
        } else {
            v22 := GenS3{F0: 94}
            v2 = v6
            _ = v22
        }
        var v23 int32 = -1
        _ = v23
    } else {
        v24 := GenS4{F0: true, F1: 1}
        _ = v24
    }
    var v25 string = "hello"
    v1 = (v1 | v1)
    v25 = v25
    if ((1e10 + 3.14159) != 1e10) {
        delete(v3, 0)
        var v26 float32 = (v4 * 3.14159)
        var v27 int16 = 1000
        _ = v26
        _ = v27
    }
    switch (("\n\t" + ("\n\t" + "foo bar")) + v25) {
    case "a":
        v28 := GenS5{F0: -730, F1: 9223372036854775807}
        _ = v28
    default:
        if (1 > 4294967295) {
            for fi29 := 0; fi29 < 7; fi29++ {
                v30 := GenS6{F0: 42}
                fmt.Println(fi29, v1, v1)
                switch fi29 {
                case 9223372036854775807:
                    if (219 != 114) {
                        if true {
                            v31 := v3[801]
                            _ = v31
                        } else {
                            switch "\n\t" {
                            case "\n\t":
                                if (false || (!(-30 <= -11))) {
                                    switch 127 {
                                    case 127:
                                        v32 := v5[2147483648]
                                        _ = v32
                                    default:
                                        for fi33 := 0; fi33 < 12; fi33++ {
                                            v34 := GenS3{F0: 1}
                                            var v35 uint = 42
                                            if (false || false) {
                                                if false {
                                                    v35 += v35
                                                    fmt.Println(v2, fi33)
                                                    switch (v25 + v25) {
                                                    case "hello":
                                                        if (1 <= v35) {
                                                            var v36 uint32 = 4294967295
                                                            var v37 int32 = v6
                                                            v6 = v37
                                                            _ = v36
                                                            _ = v37
                                                        }
                                                    default:
                                                        v6 = (v2 + (-(v6)))
                                                    }
                                                } else {
                                                    v35 += (v35 + 1)
                                                    v25 = v25
                                                    var v38 uint16 = (v1 * 2)
                                                    _ = v38
                                                }
                                                v39 := []int{fi33}
                                                _ = v39
                                            } else {
                                                switch "" {
                                                case "\n\t":
                                                    fmt.Println(v4, v2)
                                                }
                                                v1 = 32767
                                            }
                                            _ = fi33
                                            _ = v34
                                            _ = v35
                                        }
                                    }
                                    switch (-(fi29)) {
                                    case 786:
                                        var v40 int16 = 32767
                                        _ = v40
                                    case 1:
                                        for fi41 := 0; fi41 < 9; fi41++ {
                                            v42 := GenS7{F0: (v25 + (v25 + (v25 + v25)))}
                                            if (false || ((444 == 32767) || true)) {
                                                switch "foo bar" {
                                                case "hello":
                                                    v2 -= (^(v2))
                                                case "foo bar":
                                                    fmt.Println(v1, fi41)
                                                case "":
                                                    v1 += (v1 << ((18446744073709551615) & 63))
                                                }
                                            } else {
                                                switch "a" {
                                                case "\n\t":
                                                    v4 = (((1e10 - (1.5 * v4)) + v4) - v4)
                                                case "hello":
                                                    v43 := GenS3{F0: 1}
                                                    _ = v43
                                                }
                                                v2 = (v6 + 1)
                                                v44 := []int{(fi29 << ((65535) & 63)), (fi29 / ((-595) | 1))}
                                                _ = v44
                                            }
                                            _ = fi41
                                            _ = v42
                                        }
                                    }
                                    var v45 string = "\n\t"
                                    _ = v45
                                }
                            case "hello":
                                v46 := []int{fi29, int(v1), (fi29 % (((fi29 * (fi29 + 31))) | 1)), (fi29 * fi29)}
                                _ = v46
                            }
                            if true {
                                if (v6 > v2) {
                                    v6 = 1
                                    v4 = (v4 + float32(3.14159))
                                }
                                v47 := GenS8{F0: 0, F1: 42, F2: "a"}
                                _ = v47
                            } else {
                                for fi48 := 0; fi48 < 10; fi48++ {
                                    v4 = 1e10
                                    _ = fi48
                                }
                                v49 := GenS5{F0: 911, F1: fi29}
                                _ = v49
                            }
                        }
                        v4 = v4
                    }
                case 32768:
                    var v50 uint8 = 8
                    _ = v50
                case 42:
                    delete(v3, 1)
                }
                _ = fi29
                _ = v30
            }
            for fi51 := 0; fi51 < 8; fi51++ {
                v52 := []int{fi51, fi51, fi51, fi51}
                _ = fi51
                _ = v52
            }
            switch 4294967295 {
            case -1:
                fmt.Println(v6)
            case -948:
                v53 := GenS7{F0: v25}
                _ = v53
            }
        }
    }
    v6 += (v2 >> ((65535) & 63))
    var v54 bool = (-1 <= 32767)
    v55 := GenS1{F0: (v2 & 777), F1: -684}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v54:", v54)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v54
    _ = v55
    _ = v6
}
