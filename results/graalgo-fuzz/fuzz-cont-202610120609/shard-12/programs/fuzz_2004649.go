package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16; F2 uint }
type GenS2 struct { F0 bool; F1 uint; F2 uint }
type GenS3 struct { F0 string; F1 int64; F2 uint64 }
type GenS4 struct { F0 bool; F1 uint8; F2 int }
func genFn0(p0 int64, p1 string, p2 int32) uint8 {
    var v5 string = ("a" + p1)
    v6 := GenS1{F0: 128, F1: 32767, F2: 803}
    v7 := GenS1{F0: 412, F1: -525, F2: 4294967295}
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return 127
}
func genFn1(p0 uint8) int {
    var v8 string = (("" + "a") + (("foo bar" + "\n\t") + "\n\t"))
    v8 = v8
    switch 42 {
    case 42:
        v9 := []int{-271, -9223372036854775808, -210}
        _ = v9
    case 684:
        var v10 uint8 = (p0 << ((18446744073709551615) & 63))
        _ = v10
    case -717:
        if ((p0 + (p0 - 1)) <= (^(p0))) {
            if true {
                var v11 uint16 = 2
                var v12 float64 = (3.14159 + 0.1)
                _ = v11
                _ = v12
            } else {
                switch v8 {
                case "foo bar":
                    switch (v8 + v8) {
                    case "\n\t":
                        v13 := GenS2{F0: (!(v8 <= v8)), F1: 255, F2: 149}
                        _ = v13
                    default:
                        p0 = p0
                    }
                case "":
                    switch v8 {
                    case "":
                        p0 = p0
                    case "hello":
                        p0 -= (^(p0))
                    }
                case "a":
                    v14 := GenS1{F0: -2147483648, F1: -32768, F2: 65535}
                    _ = v14
                default:
                    for fi15 := 0; fi15 < 13; fi15++ {
                        for fi16 := 0; fi16 < 16; fi16++ {
                            p0 += p0
                            var v17 uint64 = 720
                            v18 := GenS1{F0: fi16, F1: -1, F2: 42}
                            _ = fi16
                            _ = v17
                            _ = v18
                        }
                        _ = fi15
                    }
                }
            }
        } else {
            v19 := []int{}
            for fi20 := 0; fi20 < 19; fi20++ {
                for fi21 := 0; fi21 < 20; fi21++ {
                    var v22 float32 = -2.25
                    _ = fi21
                    _ = v22
                }
                for fi23 := 0; fi23 < 17; fi23++ {
                    v24 := GenS1{F0: 444, F1: -32768, F2: 18446744073709551615}
                    var v25 uint = 42
                    var v26 int64 = -557
                    v27 := GenS3{F0: "foo bar", F1: (^(v26)), F2: 502}
                    v8 = ((v8 + v8) + "a")
                    _ = fi23
                    _ = v24
                    _ = v25
                    _ = v26
                    _ = v27
                }
                _ = fi20
            }
            var v28 int16 = 1000
            _ = v19
            _ = v28
        }
    }
    var v29 int16 = 1
    switch -740 {
    case 0:
        v30 := []int{312, 2147483647, 32768}
        _ = v30
    default:
        var v31 int64 = -9223372036854775808
        _ = v31
    }
    _ = p0
    _ = v29
    _ = v8
    return 2147483647
}

func main() {
    var v1 float32 = (0.0 * (1e10 + 3.14159))
    var v2 float64 = 0.0
    v3 := []uint32{65535, 4294967295, 0, 120}
    v4 := make(map[uint32]int)
    v32 := GenS4{F0: (!true), F1: 2, F2: 426}
    if ((1e10 * 1.5) != ((1e10 * (v2 + 1e10)) + 0.1)) {
        var v33 uint8 = 29
        for fi34 := 0; fi34 < 9; fi34++ {
            var v35 uint32 = 2147483647
            v4[65535] = (fi34 & fi34)
            _ = fi34
            _ = v35
        }
        switch 4294967295 {
        case 9223372036854775807:
            for fi36 := 0; fi36 < 19; fi36++ {
                switch fi36 {
                case 42:
                    v33 -= (v33 * (v33 << ((42) & 63)))
                case 887:
                    var v37 uint16 = 255
                    _ = v37
                }
                _ = fi36
            }
        case 32768:
            if (!((42 > 0) || true)) {
                v38 := GenS3{F0: "\n\t", F1: 1, F2: 42}
                v33 += genFn0(-1, "", -393)
                _ = v38
            } else {
                v2 = v2
                if true {
                    var v39 int16 = 1000
                    var v40 uint8 = 0
                    _ = v39
                    _ = v40
                }
                v2 += 3.14159
            }
        }
        _ = v33
    }
    delete(v4, 42)
    v1 -= v1
    v2 += v2
    for fi41 := 0; fi41 < 5; fi41++ {
        var v42 string = "hello"
        if false {
            var v43 uint32 = 639
            _ = v43
        } else {
            v4[2] = (fi41 / ((-948) | 1))
        }
        _ = fi41
        _ = v42
    }
    v4[509] = -507
    if true {
        if true {
            v2 += v2
            var v44 uint8 = 42
            v2 = v2
            _ = v44
        }
        v1 = (1e10 - v1)
        for fi45 := 0; fi45 < 15; fi45++ {
            var v46 int64 = 9223372036854775807
            var v47 float64 = v2
            _ = fi45
            _ = v46
            _ = v47
        }
    } else {
        for fi48 := 0; fi48 < 16; fi48++ {
            _ = len(v3)
            v49 := GenS4{F0: (!((42 >= 671) && true)), F1: 127, F2: (-(fi48))}
            switch (("foo bar" + ("\n\t" + "\n\t")) + "\n\t") {
            case "\n\t":
                v50 := GenS2{F0: true, F1: 1, F2: 18446744073709551615}
                _ = v50
            case "foo bar":
                if false {
                    var v51 uint32 = 535
                    _ = v51
                } else {
                    _ = len(v4)
                    fmt.Println(v1, v2)
                    v2 += float64(((((v1 / (v1 / (v1 - v1))) / v1) * v1) + v1))
                }
            default:
                switch "hello" {
                case "hello":
                    fmt.Println(fi48, fi48, v1)
                case "":
                    if (fi48 >= fi48) {
                        var v52 int = fi48
                        v2 -= (v2 + v2)
                        v2 = ((v2 * -2.25) + 0.1)
                        _ = v52
                    } else {
                        v2 -= v2
                        v3 = append(v3, 4294967295)
                        for fi53 := 0; fi53 < 1; fi53++ {
                            v1 += (v1 + float32(-2.25))
                            var v54 uint8 = 227
                            _ = fi53
                            _ = v54
                        }
                    }
                }
            }
            _ = fi48
            _ = v49
        }
        _ = len(v3)
    }
    if ((1.5 - v1) > ((((v1 - (0.1 / v1)) - float32((((v2 * (float64(-2.25) / float64(v1))) + ((1.5 - (v2 / ((v2 - 1e10) + v2))) + (((v2 / (v2 / 1.0)) * 1e10) + v2))) - 3.14159))) / 3.14159) / -2.25)) {
        v55 := GenS3{F0: "", F1: -1, F2: 4294967295}
        switch "" {
        case "a":
            v1 = v1
        case "":
            v1 = 0.0
        case "foo bar":
            fmt.Println(v2)
        default:
            v2 = (float64(1e10) / 1.5)
        }
        _ = v55
    }
    for fi56 := 0; fi56 < 20; fi56++ {
        v2 = v2
        v2 += v2
        _ = fi56
    }
    _ = len(v4)
    v57 := GenS3{F0: "hello", F1: -1, F2: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v57
}
