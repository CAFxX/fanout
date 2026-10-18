package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint64; F1 uint32; F2 uint32 }
type GenS3 struct { F0 int16; F1 string }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 bool; F1 uint16; F2 int8 }
type GenS6 struct { F0 uint8; F1 uint16 }
type GenS7 struct { F0 int32; F1 int16 }
type GenS8 struct { F0 int64; F1 int32 }

func main() {
    var v1 int = 51
    v2 := make(map[int]string)
    var v3 uint32 = 65535
    var v4 uint = 18446744073709551615
    for fi5 := 0; fi5 < 11; fi5++ {
        fmt.Println(v1)
        if false {
            if true {
                v6 := GenS1{F0: true}
                _ = v6
            }
            switch "foo bar" {
            case "hello":
                v1 += fi5
            case "foo bar":
                for fi7 := 0; fi7 < 15; fi7++ {
                    if (-23 < -1) {
                        if false {
                            v4 -= 18446744073709551615
                            fmt.Println(fi5, fi7)
                            v1 += -1
                        } else {
                            fmt.Println(v4, v3, fi5)
                        }
                        if (true && true) {
                            if true {
                                fmt.Println(v1)
                                v8 := []int{}
                                _ = v8
                            }
                            fmt.Println(fi5)
                            v9 := []int{fi5}
                            _ = v9
                        } else {
                            fmt.Println(v1, fi5, v3)
                            if false {
                                v3 = v3
                                var v10 int8 = 1
                                _ = v10
                            }
                        }
                    } else {
                        v11 := GenS1{F0: true}
                        if (4294967295 < 42) {
                            var v12 uint = v4
                            v13 := []int{fi5, (fi7 >> (((v4 % ((v12) | 1))) & 63))}
                            _ = v12
                            _ = v13
                        }
                        switch fi5 {
                        case 42:
                            if false {
                                if true {
                                    fmt.Println(fi7, v4)
                                    v4 = 42
                                    var v14 uint16 = 32767
                                    v15 := GenS1{F0: false}
                                    if (767 != -117) {
                                        switch (fi7 * 2147483647) {
                                        case 1:
                                            var v16 uint16 = 32767
                                            _ = v16
                                        }
                                        var v17 int64 = 182
                                        var v18 uint64 = 42
                                        _ = v17
                                        _ = v18
                                    } else {
                                        if (1 < 0) {
                                            v19 := GenS1{F0: false}
                                            _ = v19
                                        } else {
                                            v14 -= v14
                                        }
                                    }
                                    _ = v14
                                    _ = v15
                                }
                                switch (fi5 % ((-306) | 1)) {
                                case 4294967295:
                                    if (0 <= 398) {
                                        v3 += (v3 >> (((v4 & (v4 / (((^(v4))) | 1)))) & 63))
                                    }
                                }
                                fmt.Println(v4)
                            }
                        case 0:
                            switch "foo bar" {
                            case "foo bar":
                                switch ("hello" + ("foo bar" + "")) {
                                case "a":
                                    var v20 int64 = int64(v4)
                                    _ = v20
                                case "":
                                    var v21 bool = ("foo bar" != "hello")
                                    _ = v21
                                case "\n\t":
                                    var v22 float64 = ((0.1 * float64((0.0 - float32(-2.25)))) * (((((float64(float32(((0.0 - (1e10 + (((1e10 / 1.0) - 1.5) / (-2.25 + 0.1)))) - 3.14159))) - ((3.14159 + (0.1 * (3.14159 - (-2.25 + (1.5 / (0.0 - (((3.14159 / (0.1 - -2.25)) / 1.0) + (float64(((1.5 + -2.25) + 3.14159)) + 3.14159)))))))) / float64(1e10))) / (0.1 / 1.0)) - ((0.1 * 0.1) * 1e10)) * 0.1) * float64(float32(1e10))))
                                    _ = v22
                                }
                            case "a":
                                var v23 int = (-(fi7))
                                _ = v23
                            default:
                                fmt.Println(v3, v4)
                            }
                        default:
                            var v24 uint16 = 1
                            _ = v24
                        }
                        _ = v11
                    }
                    _ = fi7
                }
            case "\n\t":
                v4 = (v4 % (((v4 / (((v4 / (((^(v4))) | 1))) | 1))) | 1))
            }
            for fi25 := 0; fi25 < 13; fi25++ {
                if (v3 >= 42) {
                    v26 := GenS2{F0: 42, F1: 45, F2: (v3 & (^(v3)))}
                    _ = v26
                }
                var v27 int16 = 509
                var v28 int64 = -9223372036854775808
                var v29 float32 = -2.25
                _ = fi25
                _ = v27
                _ = v28
                _ = v29
            }
        } else {
            if false {
                v1 = 4294967295
            } else {
                v30 := v2[(v1 * 9223372036854775807)]
                var v31 int = fi5
                _ = v30
                _ = v31
            }
            fmt.Println(v4, v3, v4)
            v2[(fi5 - (fi5 ^ 32767))] = "\n\t"
        }
        _ = fi5
    }
    v4 = (v4 >> (((v4 - v4)) & 63))
    var v32 int16 = -32768
    var v33 bool = ((v4 | v4) != (v4 & (v4 >> ((v4) & 63))))
    var v34 uint8 = 255
    fmt.Println(v32, v34, v4)
    if v33 {
        if ((-73 >= 291) && (32767 != 1)) {
            var v35 float32 = 0.0
            for fi36 := 0; fi36 < 5; fi36++ {
                fmt.Println(v3, v34)
                _ = fi36
            }
            switch v1 {
            case 128:
                for fi37 := 0; fi37 < 4; fi37++ {
                    if (!(v32 != (v32 ^ (v32 ^ 32767)))) {
                        switch "foo bar" {
                        case "foo bar":
                            var v38 uint8 = v34
                            _ = v38
                        }
                        for fi39 := 0; fi39 < 15; fi39++ {
                            v3 += 77
                            _ = fi39
                        }
                        fmt.Println(v3, v33, fi37)
                    }
                    v1 = 281
                    _ = fi37
                }
            default:
                for fi40 := 0; fi40 < 16; fi40++ {
                    v41 := GenS3{F0: (v32 / ((v32) | 1)), F1: (("hello" + "a") + "hello")}
                    delete(v2, (fi40 >> ((v4) & 63)))
                    _ = fi40
                    _ = v41
                }
            }
            _ = v35
        }
        switch (-(v1)) {
        case 358:
            v3 -= v3
        case 2147483647:
            fmt.Println(v33, v33, v34)
        default:
            for fi42 := 0; fi42 < 7; fi42++ {
                for fi43 := 0; fi43 < 1; fi43++ {
                    switch "\n\t" {
                    case "a":
                        v44 := GenS4{F0: 609}
                        _ = v44
                    }
                    _ = fi43
                }
                _ = fi42
            }
        }
        fmt.Println(v32, v4)
    }
    fmt.Println(v4, v33)
    v34 += v34
    switch "hello" {
    case "\n\t":
        v45 := []int{v1, -1}
        _ = v45
    case "hello":
        var v46 uint16 = 1
        _ = v46
    default:
        if (v34 >= (v34 | (v34 + v34))) {
            v47 := GenS5{F0: (!(721 != 863)), F1: 652, F2: -1}
            _ = v47
        } else {
            if true {
                if false {
                    if true {
                        if (true || true) {
                            switch "\n\t" {
                            case "hello":
                                if ((-2.25 / -2.25) < 0.0) {
                                    var v48 string = "foo bar"
                                    fmt.Println(v3, v32, v32)
                                    _ = v48
                                } else {
                                    var v49 int8 = 1
                                    v3 += v3
                                    _ = v49
                                }
                            case "foo bar":
                                v1 = (v1 >> ((uint(v34)) & 63))
                            case "\n\t":
                                v50 := GenS6{F0: 165, F1: 260}
                                _ = v50
                            }
                            fmt.Println(v33, v1)
                            v2[(-(v1))] = ("hello" + "foo bar")
                        } else {
                            var v51 int64 = -1
                            if (0 > v4) {
                                switch (v1 ^ v1) {
                                case 946:
                                    fmt.Println(v3, v34, v33)
                                default:
                                    if (100 >= -128) {
                                        v3 += v3
                                    }
                                }
                                var v52 uint8 = uint8(v3)
                                fmt.Println(v51, v32, v4)
                                _ = v52
                            }
                            _ = v51
                        }
                    } else {
                        var v53 uint8 = 42
                        v1 = v1
                        _ = v53
                    }
                }
                var v54 uint8 = (v34 + 1)
                _ = v54
            } else {
                if ((v4 << ((v4) & 63)) != v4) {
                    fmt.Println(v33)
                } else {
                    v33 = true
                    fmt.Println(v32)
                    fmt.Println(v3, v34)
                }
            }
        }
    }
    if ((v32 >> (((v4 / ((0) | 1))) & 63)) == v32) {
        v55 := []int{(^(v1)), 42, (v1 >> ((v4) & 63))}
        var v56 uint32 = 1
        var v57 int = v1
        _ = v55
        _ = v56
        _ = v57
    }
    if (513 > (-(v1))) {
        v58 := GenS7{F0: 1, F1: (^(v32))}
        if (3.14159 < ((((-2.25 * 1e10) * ((0.1 / -2.25) - 1e10)) * (3.14159 - float64(-2.25))) + ((1e10 * float64((((float32(1.5) / float32(3.14159)) / -2.25) + 0.1))) / 1e10))) {
            var v59 uint64 = 0
            var v60 float64 = 3.14159
            _ = v59
            _ = v60
        } else {
            v3 += v3
        }
        _ = v58
    }
    fmt.Println(v32)
    var v61 uint = 1
    switch (v1 + -384) {
    case -9223372036854775808:
        if ((42 >= 2) && v33) {
            if false {
                switch (v1 + (v1 >> ((547) & 63))) {
                case 1:
                    v62 := GenS8{F0: 9223372036854775807, F1: -1}
                    _ = v62
                }
            }
            for fi63 := 0; fi63 < 1; fi63++ {
                v33 = true
                v33 = v33
                _ = fi63
            }
            v64 := GenS5{F0: true, F1: 320, F2: 127}
            _ = v64
        } else {
            v65 := []int{-13, v1, 42, v1}
            _ = v65
        }
    case 728:
        v2[32768] = "foo bar"
    case 42:
        var v66 int16 = (^(v32))
        _ = v66
    default:
        for fi67 := 0; fi67 < 11; fi67++ {
            v1 = 127
            _ = fi67
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v61
}
