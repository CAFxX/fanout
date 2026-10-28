package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int32; F1 bool }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint; F1 int }
type GenS7 struct { F0 uint16; F1 uint; F2 uint16 }
func genFn0(p0 string, p1 uint, p2 int64) uint {
    var v6 int16 = 1
    fmt.Println(v6, p2, p0)
    p1 = (p1 << (((^(p1))) & 63))
    p0 = (("a" + "\n\t") + p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return p1
}

func main() {
    var v1 float64 = 1.5
    var v2 float64 = (float64((1e10 + 1.5)) - ((v1 * ((v1 - 0.1) / v1)) * 3.14159))
    var v3 string = "foo bar"
    var v4 uint32 = 2147483647
    var v5 float32 = 0.0
    var v7 int8 = 0
    v3 = v3
    v8 := []int{int(v4), 1, -836}
    for fi9 := 0; fi9 < 9; fi9++ {
        v3 = "foo bar"
        _ = fi9
    }
    if false {
        v10 := GenS1{F0: 42}
        _ = v10
    } else {
        v11 := GenS1{F0: 842}
        v12 := GenS2{F0: 688}
        switch "" {
        case "\n\t":
            var v13 uint = genFn0(((v3 + v3) + ("a" + "a")), 255, -9223372036854775808)
            _ = v13
        case "foo bar":
            var v14 int64 = 223
            _ = v14
        case "a":
            var v15 uint64 = 2
            _ = v15
        }
        _ = v11
        _ = v12
    }
    for fi16 := 0; fi16 < 10; fi16++ {
        v17 := GenS2{F0: 799}
        if (1 < fi16) {
            var v18 int64 = -1
            switch v3 {
            case "a":
                var v19 float32 = -2.25
                _ = v19
            case "hello":
                var v20 int8 = v7
                _ = v20
            case "\n\t":
                var v21 uint32 = v4
                _ = v21
            default:
                v22 := GenS3{F0: 255}
                _ = v22
            }
            _ = v18
        }
        for fi23 := 0; fi23 < 10; fi23++ {
            if false {
                switch (fi16 % (((fi23 - (fi23 % (((fi16 - 32768)) | 1)))) | 1)) {
                case -9223372036854775808:
                    fmt.Println(fi16, v3, v2)
                case 636:
                    switch "\n\t" {
                    case "a":
                        switch ("a" + (v3 + v3)) {
                        case "foo bar":
                            if false {
                                if true {
                                    if (false && true) {
                                        v2 = v1
                                    }
                                    v24 := GenS1{F0: 4294967295}
                                    _ = v24
                                }
                                fmt.Println(fi23, v7, fi16)
                            }
                        case "a":
                            var v25 int16 = 1
                            v25 += (v25 >> ((0) & 63))
                            v26 := GenS1{F0: 4294967295}
                            _ = v25
                            _ = v26
                        default:
                            if false {
                                var v27 float32 = v5
                                if (("hello" + "") <= "\n\t") {
                                    var v28 uint16 = 2
                                    if false {
                                        var v29 int16 = -32768
                                        _ = len(v8)
                                        v2 -= v2
                                        _ = v29
                                    }
                                    v30 := GenS4{F0: -335, F1: false}
                                    v31 := GenS4{F0: 1, F1: false}
                                    v2 = v2
                                    _ = v28
                                    _ = v30
                                    _ = v31
                                } else {
                                    v8 = append(v8, fi16)
                                    switch fi23 {
                                    case 32768:
                                        _ = len(v8)
                                    default:
                                        if (2 <= 255) {
                                            fmt.Println(v5)
                                            switch "a" {
                                            case "\n\t":
                                                v32 := GenS3{F0: genFn0(("" + ((v3 + "a") + (v3 + ("hello" + ("hello" + ("hello" + ((v3 + ((v3 + v3) + "")) + v3))))))), 2, -1)}
                                                v33 := GenS4{F0: 2147483647, F1: true}
                                                _ = v32
                                                _ = v33
                                            case "hello":
                                                v1 += float64(float32(-2.25))
                                            case "a":
                                                var v34 uint32 = v4
                                                _ = v34
                                            default:
                                                var v35 uint = 2
                                                _ = v35
                                            }
                                            var v36 bool = (!(255 > 171))
                                            _ = v36
                                        } else {
                                            var v37 uint = 0
                                            v38 := GenS5{F0: 151}
                                            if len(v8) >= 2 {
                                                v8 = v8[1:len(v8)]
                                            }
                                            _ = v37
                                            _ = v38
                                        }
                                    }
                                    var v39 int16 = 1000
                                    _ = v39
                                }
                                _ = v27
                            } else {
                                v5 = float32(((v1 / (1e10 - -2.25)) * (v2 / (3.14159 / 0.1))))
                                v8 = append(v8, -769)
                            }
                        }
                    }
                case 260:
                    if true {
                        fmt.Println(fi16)
                        switch ((v3 + v3) + "a") {
                        case "hello":
                            v4 -= 534
                        case "foo bar":
                            var v40 float32 = (0.1 + float32(-2.25))
                            v2 = v1
                            v5 = ((v40 * (v40 / (1e10 / ((v40 + (v5 - (((v40 / 0.1) / v40) / (float32((0.0 + ((0.0 / (-2.25 / float64(-2.25))) * 1.5))) - v5)))) / float32(v2))))) + v5)
                            _ = v40
                        case "":
                            fmt.Println(v3, v5, fi16)
                        }
                    }
                }
                var v41 uint8 = 10
                _ = v41
            }
            _ = fi23
        }
        _ = fi16
        _ = v17
    }
    var v42 int64 = 9223372036854775807
    fmt.Println(v1, v4)
    if true {
        v7 = (v7 - v7)
        v43 := GenS5{F0: 1}
        _ = v43
    }
    for fi44 := 0; fi44 < 14; fi44++ {
        v45 := GenS6{F0: 65535, F1: (fi44 - -9223372036854775808)}
        var v46 int = int(v4)
        var v47 string = ""
        _ = fi44
        _ = v45
        _ = v46
        _ = v47
    }
    fmt.Println(v1)
    switch ((v3 + ("hello" + v3)) + v3) {
    case "foo bar":
        v2 = v2
    case "hello":
        for fi48 := 0; fi48 < 12; fi48++ {
            v8 = append(v8, fi48)
            var v49 uint32 = (v4 + (^(v4)))
            _ = fi48
            _ = v49
        }
    case "\n\t":
        fmt.Println(v3, v2, v2)
    }
    if len(v8) > 0 {
        v8[0] = 62
    }
    switch -9223372036854775808 {
    case 127:
        if len(v8) > 0 {
            v8[0] = 128
        }
    }
    if false {
        var v50 int16 = -1
        v51 := GenS7{F0: 302, F1: 139, F2: 32767}
        _ = v50
        _ = v51
    } else {
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
    }
    var v52 uint8 = 56
    for fi53 := 0; fi53 < 3; fi53++ {
        var v54 float64 = (0.0 * (v2 + 1.5))
        _ = fi53
        _ = v54
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v5
    _ = v52
    _ = v7
    _ = v8
}
