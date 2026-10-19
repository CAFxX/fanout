package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64 }
type GenS2 struct { F0 uint16; F1 bool }
type GenS3 struct { F0 string; F1 uint16; F2 int }
type GenS4 struct { F0 uint16; F1 int }
type GenS5 struct { F0 int; F1 int }
type GenS6 struct { F0 uint32; F1 uint64 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 uint8; F1 int8; F2 uint32 }
type GenS9 struct { F0 uint }

func main() {
    var v1 int8 = 0
    var v2 float32 = (1e10 - (-2.25 - 0.1))
    v3 := make(map[int]bool)
    v4 := []int{-750, -1, -2147483648}
    var v5 int8 = v1
    v6 := GenS1{F0: 805, F1: -9223372036854775808}
    var v7 int = 4294967295
    fmt.Println(v2)
    var v8 uint32 = 368
    fmt.Println(v1)
    var v9 float32 = 1e10
    v3[(v7 ^ -430)] = true
    fmt.Println(v8, v9)
    v9 = v9
    for fi10 := 0; fi10 < 8; fi10++ {
        for fi11 := 0; fi11 < 11; fi11++ {
            if true {
                if (1 == 1) {
                    if len(v4) > 0 {
                        v4[0] = (fi10 - (v7 ^ v7))
                    }
                    var v12 int = fi11
                    var v13 int8 = 4
                    var v14 uint16 = 1
                    if false {
                        var v15 int = 127
                        v14 = (v14 ^ (v14 << ((255) & 63)))
                        var v16 int8 = v1
                        _ = v15
                        _ = v16
                    } else {
                        if true {
                            v17 := GenS2{F0: (v14 << ((255) & 63)), F1: false}
                            if true {
                                var v18 string = "a"
                                if false {
                                    if true {
                                        v13 += v5
                                        var v19 int64 = int64(v7)
                                        switch (((v18 + "hello") + ((v18 + ("foo bar" + v18)) + v18)) + ("a" + v18)) {
                                        case "":
                                            v20 := GenS3{F0: v18, F1: 42, F2: fi11}
                                            _ = v20
                                        case "\n\t":
                                            v3[v7] = false
                                        }
                                        if ((int8(v19) >= v5) || false) {
                                            var v21 float64 = 3.14159
                                            _ = v21
                                        }
                                        _ = v19
                                    }
                                }
                                _ = v18
                            } else {
                                var v22 int32 = 790
                                v23 := GenS1{F0: 4294967295, F1: 9223372036854775807}
                                _ = v22
                                _ = v23
                            }
                            var v24 float32 = v2
                            switch -2147483648 {
                            case 42:
                                var v25 uint32 = 2147483648
                                _ = v25
                            case 4294967295:
                                var v26 int = v12
                                _ = v26
                            case 9223372036854775807:
                                v2 -= (v2 + 0.1)
                            default:
                                v14 = (v14 ^ uint16(v8))
                            }
                            _ = v17
                            _ = v24
                        }
                        v27 := GenS4{F0: (v14 | (v14 << ((4294967295) & 63))), F1: -758}
                        if (((((v9 + float32(float64(0.0))) * (v2 / v9)) * ((((((v9 + v2) * 0.1) - 3.14159) + 1.5) - (-2.25 - v2)) * v2)) == -2.25) || (42 == v1)) {
                            if (1 > 9223372036854775807) {
                                var v28 int16 = -1
                                v29 := GenS2{F0: 1, F1: (0.0 >= (((((v9 * ((((v2 * (3.14159 - (1.5 * v2))) - (0.1 - -2.25)) - (v9 + 1.5)) - (v9 + float32(0.0)))) / 1.5) - 0.0) * v2) + ((v9 * 0.1) - (0.0 - 1.5))))}
                                switch fi10 {
                                case 784:
                                    v30 := GenS5{F0: v12, F1: 698}
                                    _ = v30
                                case -53:
                                    switch v7 {
                                    case -700:
                                        if (v14 == (v14 | v14)) {
                                            fmt.Println(v28, v28, fi10)
                                            var v31 int = (v12 / ((-600) | 1))
                                            _ = v31
                                        } else {
                                            delete(v3, v7)
                                            var v32 float64 = (-2.25 + 3.14159)
                                            _ = v32
                                        }
                                    case 127:
                                        var v33 uint32 = v8
                                        _ = v33
                                    case 32767:
                                        v5 -= 100
                                    }
                                case 313:
                                    if false {
                                        if (false || true) {
                                            v34 := GenS5{F0: (fi10 << ((2) & 63)), F1: -512}
                                            v3[(fi11 - (fi11 - 165))] = false
                                            _ = v34
                                        } else {
                                            if true {
                                                v5 += v5
                                                fmt.Println(v13, v28, v13)
                                            } else {
                                                _ = len(v3)
                                                switch "foo bar" {
                                                case "\n\t":
                                                    v8 = 1
                                                case "hello":
                                                    if true {
                                                        fmt.Println(v2, fi10, v1)
                                                        var v35 bool = (int(v28) != (fi11 | -9223372036854775808))
                                                        v36 := GenS2{F0: 1, F1: true}
                                                        var v37 uint32 = (v8 / ((4294967295) | 1))
                                                        _ = v35
                                                        _ = v36
                                                        _ = v37
                                                    } else {
                                                        v14 = uint16(v8)
                                                    }
                                                case "a":
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                default:
                                                    v7 = fi10
                                                }
                                                if (!(18446744073709551615 != 307)) {
                                                    v14 += (v14 | (v14 ^ 988))
                                                    var v38 uint8 = uint8(v14)
                                                    _ = v38
                                                } else {
                                                    v14 -= 2
                                                    v7 -= v7
                                                    v39 := GenS2{F0: (v14 ^ (v14 >> ((941) & 63))), F1: true}
                                                    _ = v39
                                                }
                                            }
                                            if (true || true) {
                                                if (!false) {
                                                    v12 = v7
                                                    var v40 int8 = 1
                                                    var v41 float64 = (0.0 - ((-2.25 - ((1e10 / (0.0 + ((-2.25 * (((3.14159 + (1.5 * 0.0)) - (1.5 / 3.14159)) + (float64(v2) + (0.0 + (((float64(v9) - float64(v9)) - ((3.14159 - 3.14159) + (1e10 / 1.0))) * 1.5))))) / 3.14159))) / -2.25)) * 3.14159))
                                                    v42 := GenS4{F0: (v14 + (v14 + v14)), F1: 32767}
                                                    v43 := GenS6{F0: v8, F1: 4294967295}
                                                    _ = v40
                                                    _ = v41
                                                    _ = v42
                                                    _ = v43
                                                }
                                            }
                                            v12 = fi11
                                        }
                                    } else {
                                        v5 -= v13
                                    }
                                default:
                                    v1 -= v13
                                }
                                _ = v28
                                _ = v29
                            } else {
                                v1 += v1
                            }
                        }
                        var v44 string = "a"
                        _ = v27
                        _ = v44
                    }
                    _ = v12
                    _ = v13
                    _ = v14
                } else {
                    v5 = v5
                }
            }
            if false {
                switch "hello" {
                case "foo bar":
                    if (!(v8 > (v8 * (v8 / ((65535) | 1))))) {
                        var v45 int = (fi10 | (-(fi11)))
                        _ = v45
                    } else {
                        fmt.Println(fi10)
                    }
                case "":
                    if (4294967295 <= 18446744073709551615) {
                        v46 := GenS3{F0: "a", F1: 0, F2: fi10}
                        v1 += (v1 | (v5 - (v5 << ((1) & 63))))
                        _ = v46
                    }
                case "hello":
                    v47 := GenS7{F0: 42}
                    _ = v47
                }
                switch "" {
                case "":
                    if ((v1 % ((-1) | 1)) != int8(v7)) {
                        if (true && (v9 == v9)) {
                            v48 := GenS4{F0: 42, F1: (fi11 & fi10)}
                            v9 -= (3.14159 + -2.25)
                            _ = v48
                        } else {
                            if true {
                                v9 = 1.5
                                if (0 != 2) {
                                    if false {
                                        var v49 int = (v7 | (fi11 * fi10))
                                        v2 += v9
                                        var v50 int = (fi11 / (((-(v49))) | 1))
                                        _ = v49
                                        _ = v50
                                    }
                                    var v51 int16 = 0
                                    switch ("a" + "") {
                                    case "":
                                        var v52 uint64 = 4
                                        _ = v52
                                    case "hello":
                                        v5 -= (v5 % (((v1 - 14)) | 1))
                                    default:
                                        var v53 string = "a"
                                        _ = v53
                                    }
                                    if true {
                                        var v54 float32 = (v9 * 0.0)
                                        v7 = fi10
                                        v2 -= 0.1
                                        _ = v54
                                    }
                                    _ = v51
                                }
                                v55 := v3[(fi11 | (v7 >> ((255) & 63)))]
                                v7 = fi10
                                _ = v55
                            } else {
                                v9 = v2
                                v4 = append(v4, int(v8))
                                switch fi10 {
                                case -9223372036854775808:
                                    v56 := v3[(v7 % (((fi10 | fi10)) | 1))]
                                    _ = v56
                                }
                                var v57 int32 = 1
                                _ = v57
                            }
                            var v58 uint64 = 106
                            fmt.Println(fi10)
                            _ = v58
                        }
                    } else {
                        v59 := GenS8{F0: 128, F1: int8(v8), F2: 2147483648}
                        _ = v59
                    }
                default:
                    if true {
                        v9 = (((1e10 * v2) + (v2 * v9)) - (0.0 * v2))
                    } else {
                        v60 := GenS9{F0: 606}
                        if (220 == 128) {
                            var v61 string = "a"
                            _ = v61
                        }
                        v1 = 90
                        _ = v60
                    }
                }
                if (v9 > float32(0.1)) {
                    if false {
                        var v62 float32 = float32(((((0.1 / float64(-2.25)) - (float64((v2 + (float32(-2.25) - 0.0))) - (((((-2.25 - (0.0 + 3.14159)) / 3.14159) + -2.25) * 1e10) - 1e10))) / 1.5) / 1.5))
                        _ = v62
                    } else {
                        v63 := GenS2{F0: 1, F1: ((100 == v5) || true)}
                        _ = v63
                    }
                    fmt.Println(v9, v8)
                }
                v64 := GenS3{F0: ("foo bar" + ("foo bar" + (("foo bar" + "\n\t") + ("hello" + "a")))), F1: 1, F2: fi10}
                _ = v64
            }
            _ = fi11
        }
        _ = fi10
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
