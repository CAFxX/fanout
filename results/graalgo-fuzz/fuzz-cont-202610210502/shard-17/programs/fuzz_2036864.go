package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 int32; F1 int8; F2 int8 }
type GenS9 struct { F0 int16; F1 uint8; F2 string }
type GenS10 struct { F0 int64; F1 int8 }
func genFn0() int32 {
    var v7 float64 = 0.1
    var v8 uint8 = 1
    if false {
        if false {
            v9 := GenS1{F0: 1}
            for fi10 := 0; fi10 < 18; fi10++ {
                var v11 uint = 997
                v12 := []int{9223372036854775807}
                var v13 int = fi10
                _ = fi10
                _ = v11
                _ = v12
                _ = v13
            }
            if (!true) {
                v14 := GenS1{F0: 42}
                _ = v14
            }
            _ = v9
        }
        var v15 uint32 = 2147483647
        var v16 uint64 = 18446744073709551615
        _ = v15
        _ = v16
    } else {
        var v17 bool = (42 <= 42)
        v18 := []int{-2147483648, 9223372036854775807, 32768}
        _ = v17
        _ = v18
    }
    v19 := GenS2{F0: 2}
    fmt.Println(v7, v8)
    _ = v19
    _ = v7
    _ = v8
    return 902
}
func genFn1(p0 string) uint {
    v20 := []int{-659, -1}
    if false {
        var v21 int16 = 748
        for fi22 := 0; fi22 < 16; fi22++ {
            switch (^(fi22)) {
            case -565:
                fmt.Println(v21, v21)
            default:
                if (fi22 <= (-(fi22))) {
                    fmt.Println(v21, p0)
                    for fi23 := 0; fi23 < 20; fi23++ {
                        v24 := GenS3{F0: 42}
                        var v25 uint64 = 162
                        _ = fi23
                        _ = v24
                        _ = v25
                    }
                    var v26 uint32 = 4294967295
                    _ = v26
                }
            }
            p0 = (p0 + (p0 + p0))
            _ = fi22
        }
        _ = v21
    } else {
        switch p0 {
        case "":
            v27 := GenS4{F0: 361}
            _ = v27
        default:
            var v28 float64 = 3.14159
            _ = v28
        }
        fmt.Println(p0)
    }
    _ = p0
    _ = v20
    return 4294967295
}

func main() {
    var v1 uint = 992
    var v2 int8 = 127
    var v3 int16 = -471
    v4 := make(map[int]int)
    var v5 int32 = 182
    var v6 string = "hello"
    switch int(v1) {
    case -9223372036854775808:
        switch ("hello" + "foo bar") {
        case "foo bar":
            var v29 int32 = 1
            _ = v29
        case "hello":
            v3 = (v3 * (-(v3)))
        default:
            var v30 bool = (0 != 1)
            _ = v30
        }
    case 127:
        fmt.Println(v1)
    }
    v2 = v2
    for fi31 := 0; fi31 < 17; fi31++ {
        var v32 int64 = 0
        switch fi31 {
        case 149:
            v33 := GenS5{F0: (v5 << ((735) & 63))}
            _ = v33
        case 127:
            var v34 int32 = (v5 * v5)
            _ = v34
        case 4294967295:
            var v35 uint8 = 255
            _ = v35
        default:
            var v36 int8 = (v2 | (v2 + v2))
            _ = v36
        }
        _ = fi31
        _ = v32
    }
    var v37 uint8 = 128
    switch -389 {
    case 0:
        var v38 uint16 = 1
        _ = v38
    case 813:
        v39 := GenS1{F0: v5}
        _ = v39
    default:
        v40 := GenS4{F0: 9223372036854775807}
        _ = v40
    }
    for fi41 := 0; fi41 < 1; fi41++ {
        var v42 float64 = ((0.1 * 1.5) - ((((1e10 / 1.0) - (1e10 - (float64((0.0 / (3.14159 * ((((float32(0.0) / float32(((0.1 / ((float64(1.5) * (((1.5 + float64((1e10 * (float32(0.0) * 3.14159)))) + -2.25) + 0.0)) / (float64((((float32(float64(1e10)) - -2.25) / 1.0) / -2.25)) + 1e10))) - 0.1))) / 1.0) / 1.0) - 1e10)))) + (0.0 - 0.1)))) * 1e10) / (0.1 + ((((float64(((1.5 - float32(-2.25)) / 1e10)) + 0.1) + (float64(((0.1 + 1e10) / 1.0)) + 1.5)) + -2.25) / 1e10))))
        v6 = (v6 + "hello")
        v6 = v6
        _ = fi41
        _ = v42
    }
    v43 := GenS6{F0: 18446744073709551615}
    for fi44 := 0; fi44 < 14; fi44++ {
        if ((0.1 / (0.1 + (3.14159 * 0.1))) == (float32(0.0) * 3.14159)) {
            v2 = 1
        } else {
            var v45 int16 = (v3 << (((v1 - 2)) & 63))
            switch (fi44 ^ fi44) {
            case 0:
                v46 := v4[fi44]
                _ = v46
            }
            v47 := GenS7{F0: v6}
            _ = v45
            _ = v47
        }
        if (0.0 > 0.1) {
            if false {
                fmt.Println(v2)
                var v48 uint64 = 1
                _ = v48
            } else {
                fmt.Println(v1)
            }
            var v49 uint8 = (v37 * (v37 ^ 255))
            v50 := GenS8{F0: (v5 % ((v5) | 1)), F1: v2, F2: v2}
            _ = v49
            _ = v50
        }
        var v51 uint32 = 65535
        _ = fi44
        _ = v51
    }
    var v52 uint = 596
    v6 = v6
    for fi53 := 0; fi53 < 1; fi53++ {
        var v54 int16 = v3
        _ = fi53
        _ = v54
    }
    var v55 float64 = (3.14159 / float64(-2.25))
    var v56 uint8 = 1
    if true {
        v57 := []int{32768, 42, 32767, 32768}
        _ = v57
    } else {
        v37 = uint8(v52)
        switch (v6 + ("hello" + v6)) {
        case "\n\t":
            fmt.Println(v55, v5)
        case "":
            _ = len(v4)
        default:
            _ = len(v4)
        }
    }
    if false {
        v58 := GenS9{F0: -460, F1: (v56 * 2), F2: (v6 + ((("foo bar" + v6) + v6) + v6))}
        if true {
            v59 := []int{9223372036854775807, 0, 0, 1}
            var v60 int32 = 0
            v59 = append(v59, 127)
            _ = v59
            _ = v60
        } else {
            delete(v4, -362)
            v1 = v1
            for fi61 := 0; fi61 < 20; fi61++ {
                switch v6 {
                case "a":
                    switch (fi61 << ((genFn1("\n\t")) & 63)) {
                    case 32767:
                        var v62 uint32 = 57
                        _ = v62
                    default:
                        v63 := v4[-9223372036854775808]
                        _ = v63
                    }
                case "foo bar":
                    if false {
                        var v64 uint32 = 260
                        delete(v4, fi61)
                        _ = v64
                    } else {
                        v55 -= -2.25
                    }
                case "\n\t":
                    v52 += v52
                default:
                    v37 -= 255
                }
                for fi65 := 0; fi65 < 5; fi65++ {
                    v66 := GenS7{F0: "foo bar"}
                    var v67 float32 = 3.14159
                    switch v6 {
                    case "":
                        fmt.Println(v2, v6, fi61)
                    }
                    fmt.Println(fi65, fi61)
                    v5 += -1
                    switch (v6 + ("" + v6)) {
                    case "hello":
                        var v68 uint64 = 0
                        _ = v68
                    case "a":
                        v52 += v1
                    }
                    _ = fi65
                    _ = v66
                    _ = v67
                }
                if (!true) {
                    v52 = v52
                } else {
                    var v69 bool = true
                    var v70 uint64 = 219
                    v71 := GenS10{F0: 1, F1: 125}
                    _ = v69
                    _ = v70
                    _ = v71
                }
                _ = fi61
            }
        }
        var v72 uint16 = 32767
        _ = v58
        _ = v72
    }
    v73 := GenS10{F0: -16, F1: v2}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    fmt.Println("v55:", v55)
    fmt.Println("v56:", v56)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v43
    _ = v5
    _ = v52
    _ = v55
    _ = v56
    _ = v6
    _ = v73
}
