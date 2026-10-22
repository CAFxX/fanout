package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint64; F1 uint8 }
func genFn0() int32 {
    if true {
        var v5 uint = 255
        var v6 int8 = -1
        _ = v5
        _ = v6
    } else {
        for fi7 := 0; fi7 < 8; fi7++ {
            for fi8 := 0; fi8 < 16; fi8++ {
                fmt.Println(fi8, fi8)
                if (!true) {
                    var v9 int = -9223372036854775808
                    _ = v9
                }
                _ = fi8
            }
            v10 := GenS1{F0: 0, F1: 0}
            var v11 uint8 = 42
            _ = fi7
            _ = v10
            _ = v11
        }
        var v12 int = 128
        if true {
            v13 := []int{128, -121, 32768}
            _ = len(v13)
            _ = v13
        } else {
            if true {
                v14 := GenS1{F0: 127, F1: 18446744073709551615}
                switch ("a" + "hello") {
                case "a":
                    var v15 uint32 = 65535
                    _ = v15
                case "\n\t":
                    var v16 uint = 255
                    _ = v16
                case "hello":
                    var v17 int8 = 1
                    _ = v17
                }
                _ = v14
            } else {
                v12 -= 2147483647
            }
            switch v12 {
            case -662:
                var v18 uint32 = 42
                _ = v18
            default:
                v12 += v12
            }
            v12 = v12
        }
        _ = v12
    }
    for fi19 := 0; fi19 < 9; fi19++ {
        var v20 uint8 = 2
        _ = fi19
        _ = v20
    }
    if (65535 <= 138) {
        var v21 int = -96
        _ = v21
    } else {
        var v22 uint64 = 277
        _ = v22
    }
    return 2147483647
}
func genFn1(p0 uint8, p1 int32) uint16 {
    if (42 != p1) {
        p1 = (p1 | (p1 / ((p1) | 1)))
    } else {
        switch "a" {
        case "foo bar":
            fmt.Println(p0, p0)
        case "a":
            var v23 float64 = (0.0 + float64((float32(1e10) - float32((0.1 - 0.0)))))
            _ = v23
        case "hello":
            var v24 float64 = (((float64((-2.25 + -2.25)) + float64(1e10)) * 1e10) - 1e10)
            _ = v24
        default:
            fmt.Println(p1, p1)
        }
        p1 -= p1
        switch -286 {
        case -2147483648:
            if (true || true) {
                fmt.Println(p0, p0)
                v25 := []int{}
                var v26 int32 = (p1 >> ((791) & 63))
                _ = v25
                _ = v26
            } else {
                p0 -= (^(p0))
                p1 = 2147483647
                var v27 uint64 = 42
                _ = v27
            }
        case 117:
            var v28 uint32 = 1
            _ = v28
        case 128:
            var v29 uint = 391
            _ = v29
        default:
            v30 := GenS2{F0: "\n\t"}
            _ = v30
        }
    }
    for fi31 := 0; fi31 < 18; fi31++ {
        p1 = p1
        var v32 float64 = 1e10
        _ = fi31
        _ = v32
    }
    v33 := []int{1, 9223372036854775807, 9223372036854775807}
    _ = p0
    _ = p1
    _ = v33
    return 42
}

func main() {
    var v1 int = -550
    v2 := []int{v1, (-(v1)), (v1 & 127)}
    v3 := make(map[int]bool)
    v4 := []uint32{2147483647, 4294967295, 0}
    if (!(false || true)) {
        for fi34 := 0; fi34 < 10; fi34++ {
            for fi35 := 0; fi35 < 13; fi35++ {
                v36 := GenS1{F0: 127, F1: 2}
                v37 := GenS2{F0: "\n\t"}
                v1 = fi34
                _ = fi35
                _ = v36
                _ = v37
            }
            v3[v1] = ((float32((0.0 * 1.5)) / (((((3.14159 / (1e10 + float32(((3.14159 + 3.14159) / 1e10)))) + ((-2.25 + -2.25) / 3.14159)) * 3.14159) - float32(((-2.25 - 0.1) - 0.1))) / 1.0)) < 3.14159)
            var v38 int16 = 176
            _ = fi34
            _ = v38
        }
        var v39 float64 = (3.14159 - 0.1)
        fmt.Println(v1, v39)
        _ = v39
    }
    fmt.Println(v1)
    for fi40 := 0; fi40 < 19; fi40++ {
        var v41 int64 = 42
        for fi42 := 0; fi42 < 4; fi42++ {
            switch fi42 {
            case -9223372036854775808:
                fmt.Println(v1, fi40, v1)
            case 127:
                var v43 int16 = 1000
                v44 := GenS3{F0: 0, F1: 2}
                _ = v43
                _ = v44
            default:
                v41 += (v41 ^ v41)
            }
            _ = fi42
        }
        _ = fi40
        _ = v41
    }
    for fi45 := 0; fi45 < 9; fi45++ {
        if true {
            v1 -= 2147483647
        } else {
            var v46 uint64 = 797
            _ = v46
        }
        v1 += 42
        for fi47 := 0; fi47 < 11; fi47++ {
            fmt.Println(v1, fi47, v1)
            v48 := GenS2{F0: "foo bar"}
            var v49 float32 = ((0.1 - float32(1e10)) + 1e10)
            _ = fi47
            _ = v48
            _ = v49
        }
        _ = fi45
    }
    var v50 string = "foo bar"
    for fi51 := 0; fi51 < 16; fi51++ {
        for fi52 := 0; fi52 < 10; fi52++ {
            delete(v3, (fi52 * 639))
            v1 += (fi51 << ((18446744073709551615) & 63))
            var v53 float64 = 0.1
            _ = fi52
            _ = v53
        }
        var v54 float32 = 0.0
        _ = fi51
        _ = v54
    }
    v1 = (v1 ^ (-(v1)))
    if false {
        var v55 int8 = -33
        var v56 float64 = 3.14159
        _ = v55
        _ = v56
    }
    v57 := GenS1{F0: 42, F1: 0}
    switch "hello" {
    case "":
        for fi58 := 0; fi58 < 10; fi58++ {
            var v59 int = v1
            for fi60 := 0; fi60 < 2; fi60++ {
                if (255 > 2) {
                    v61 := GenS1{F0: 100, F1: 4294967295}
                    _ = len(v3)
                    v50 = (v50 + (v50 + ("" + v50)))
                    _ = v61
                }
                fmt.Println(v59, fi58, v59)
                _ = fi60
            }
            _ = fi58
            _ = v59
        }
    case "foo bar":
        v50 = v50
    case "\n\t":
        var v62 float32 = (0.1 + 3.14159)
        _ = v62
    default:
        _ = len(v3)
    }
    v50 = (v50 + "")
    fmt.Println(v1, v50)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v50
    _ = v57
}
