package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint16; F1 int64 }
type GenS3 struct { F0 int; F1 uint; F2 uint32 }
type GenS4 struct { F0 int16; F1 uint32; F2 bool }

func main() {
    var v1 string = "foo bar"
    v2 := make(map[uint32]bool)
    v3 := []uint32{4294967295}
    var v4 uint32 = 0
    v5 := make(map[uint32]uint32)
    if (-2.25 == 3.14159) {
        fmt.Println(v1, v4)
        fmt.Println(v1, v4)
    }
    if false {
        var v6 int32 = 2147483647
        var v7 float64 = float64(float32((float64(0.1) * (0.0 / (1.5 + (0.1 / float64((1e10 * (((3.14159 / (3.14159 - -2.25)) * float32(3.14159)) / 1.0)))))))))
        _ = v6
        _ = v7
    }
    fmt.Println(v1)
    _ = len(v3)
    v8 := v5[(v4 | v4)]
    switch -697 {
    case 884:
        if len(v3) > 0 {
            v3[0] = (v4 * v8)
        }
    default:
        for fi9 := 0; fi9 < 11; fi9++ {
            v10 := v2[(v4 >> ((42) & 63))]
            v11 := GenS1{F0: -362}
            if false {
                if false {
                    var v12 uint32 = v4
                    v13 := GenS1{F0: 547}
                    _ = v12
                    _ = v13
                }
                var v14 int8 = 42
                for fi15 := 0; fi15 < 7; fi15++ {
                    v4 = (^(v8))
                    var v16 uint64 = 849
                    _ = fi15
                    _ = v16
                }
                _ = v14
            }
            _ = fi9
            _ = v10
            _ = v11
        }
    }
    if (!false) {
        var v17 uint8 = 255
        _ = v17
    }
    var v18 int = -227
    _ = len(v2)
    v1 = v1
    if true {
        var v19 int8 = -128
        for fi20 := 0; fi20 < 5; fi20++ {
            if (!(!false)) {
                v8 = v8
                for fi21 := 0; fi21 < 14; fi21++ {
                    var v22 uint32 = 65535
                    fmt.Println(fi20)
                    v1 = v1
                    _ = fi21
                    _ = v22
                }
                _ = len(v2)
            }
            _ = fi20
        }
        _ = v19
    } else {
        if true {
            for fi23 := 0; fi23 < 5; fi23++ {
                if (uint8(v4) <= 1) {
                    switch (v18 >> ((65535) & 63)) {
                    case 32767:
                        if true {
                            fmt.Println(v8, v4, fi23)
                            if true {
                                var v24 bool = false
                                _ = v24
                            }
                        }
                    default:
                        fmt.Println(v1, v18)
                    }
                    v25 := v2[v4]
                    v1 = v1
                    _ = v25
                } else {
                    v8 = v8
                    switch v18 {
                    case 127:
                        v1 = ((v1 + v1) + v1)
                    }
                    for fi26 := 0; fi26 < 8; fi26++ {
                        var v27 int8 = 1
                        if (1e10 != 3.14159) {
                            var v28 string = "hello"
                            _ = v28
                        } else {
                            v29 := GenS1{F0: 1}
                            _ = len(v3)
                            var v30 uint64 = 4294967295
                            _ = v29
                            _ = v30
                        }
                        _ = fi26
                        _ = v27
                    }
                }
                v31 := GenS1{F0: v18}
                if (0.1 != ((0.1 + (float64(((3.14159 + 0.0) * (0.1 / 3.14159))) * 1e10)) - 0.0)) {
                    v4 = v8
                } else {
                    fmt.Println(fi23, v1, v8)
                    for fi32 := 0; fi32 < 3; fi32++ {
                        v18 = (v18 | fi32)
                        switch 32767 {
                        case -420:
                            if false {
                                fmt.Println(v1)
                                v8 += v4
                                if false {
                                    var v33 int = 238
                                    _ = v33
                                }
                            }
                        case -9223372036854775808:
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        case 122:
                            v1 = "\n\t"
                        default:
                            _ = len(v3)
                        }
                        if ((fi32 & fi32) <= (v18 / (((-(fi23))) | 1))) {
                            var v34 int64 = 582
                            v4 += v4
                            _ = v34
                        }
                        _ = fi32
                    }
                }
                _ = fi23
                _ = v31
            }
            var v35 int32 = 0
            if false {
                var v36 int16 = 584
                var v37 int = int(v35)
                _ = v36
                _ = v37
            } else {
                var v38 uint16 = 0
                if (!true) {
                    if false {
                        if false {
                            if (false || false) {
                                for fi39 := 0; fi39 < 13; fi39++ {
                                    v1 = v1
                                    _ = fi39
                                }
                            }
                            v4 += (v8 + 1)
                        }
                        switch v1 {
                        case "":
                            v18 = (v18 % ((215) | 1))
                        case "foo bar":
                            var v40 int64 = 0
                            _ = v40
                        case "hello":
                            if true {
                                delete(v2, v8)
                            }
                        }
                        var v41 float64 = (float64((float32(((0.1 - 1e10) * (1.5 + float64(((float32(1e10) + 0.0) + -2.25))))) + -2.25)) - float64(0.1))
                        _ = v41
                    }
                }
                v42 := GenS2{F0: 255, F1: 0}
                _ = v38
                _ = v42
            }
            _ = v35
        }
    }
    _ = len(v2)
    var v43 uint16 = 42
    v43 = 510
    for fi44 := 0; fi44 < 20; fi44++ {
        switch ("\n\t" + "a") {
        case "\n\t":
            var v45 float32 = 3.14159
            _ = v45
        case "hello":
            if (!(-1 == 0)) {
                var v46 int8 = -1
                _ = v46
            }
        case "a":
            if ((1000 >= 0) || (0.0 < (3.14159 * float32(1e10)))) {
                v8 -= 4294967295
                switch int(v4) {
                case 127:
                    if (((-2.25 * -2.25) - (1e10 + 1.5)) < 3.14159) {
                        for fi47 := 0; fi47 < 17; fi47++ {
                            switch (v18 | (fi44 / (((fi44 >> ((901) & 63))) | 1))) {
                            case 127:
                                v48 := GenS3{F0: (fi47 << ((213) & 63)), F1: 621, F2: v8}
                                _ = v48
                            case 4294967295:
                                fmt.Println(v4, v4)
                            case 42:
                                var v49 uint64 = 1
                                _ = v49
                            default:
                                fmt.Println(fi47)
                            }
                            _ = fi47
                        }
                    }
                case -9223372036854775808:
                    var v50 uint64 = 42
                    _ = v50
                case 826:
                    switch v1 {
                    case "a":
                        fmt.Println(v8)
                    case "foo bar":
                        for fi51 := 0; fi51 < 10; fi51++ {
                            v5[(v4 + (v8 & (v4 >> ((65535) & 63))))] = (v8 % ((65535) | 1))
                            if (51 > 42) {
                                v1 = v1
                                switch (v18 >> ((165) & 63)) {
                                case 2147483647:
                                    var v52 string = v1
                                    _ = v52
                                case 192:
                                    v53 := GenS4{F0: 1, F1: v4, F2: true}
                                    _ = v53
                                }
                            } else {
                                var v54 int = (v18 | fi51)
                                _ = v54
                            }
                            v55 := GenS1{F0: (fi44 >> ((850) & 63))}
                            _ = fi51
                            _ = v55
                        }
                    case "hello":
                        v8 += v4
                    default:
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    }
                }
            }
        default:
            var v56 int = (v18 ^ 128)
            _ = v56
        }
        if len(v3) > 0 {
            v3[0] = (^(v4))
        }
        _ = fi44
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v5
    _ = v8
}
