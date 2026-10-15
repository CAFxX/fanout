package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint8; F1 uint16 }
func genFn0() int32 {
    var v7 int32 = 1
    var v8 uint = 65535
    v9 := GenS1{F0: -1}
    v10 := GenS1{F0: -48}
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return (^(v7))
}
func genFn1(p0 int8, p1 int32) int64 {
    if (int32(p0) == (p1 * genFn0())) {
        fmt.Println(p1, p1)
        var v11 uint32 = 87
        fmt.Println(v11, p1)
        _ = v11
    }
    for fi12 := 0; fi12 < 14; fi12++ {
        p1 = (p1 * genFn0())
        switch 4294967295 {
        case 2147483647:
            p0 += p0
        case 873:
            fmt.Println(fi12, fi12, fi12)
        case 4294967295:
            v13 := GenS1{F0: p0}
            _ = v13
        default:
            if false {
                if (false && ((!true) && false)) {
                    fmt.Println(p0, p1, p1)
                    p1 = (p1 ^ genFn0())
                    var v14 int32 = (p1 * p1)
                    _ = v14
                }
                p1 = genFn0()
                v15 := []int{fi12, fi12, (fi12 - -9223372036854775808)}
                _ = v15
            }
        }
        _ = fi12
    }
    p0 = (p0 / ((42) | 1))
    fmt.Println(p1, p0)
    v16 := []int{128, -876, 4294967295}
    _ = p0
    _ = p1
    _ = v16
    return -918
}

func main() {
    var v1 int = 32767
    var v2 uint = 18446744073709551615
    var v3 float64 = float64((3.14159 / 0.1))
    var v4 int64 = 42
    var v5 uint16 = 42
    v6 := []uint32{749}
    v3 += 1.5
    for fi17 := 0; fi17 < 10; fi17++ {
        _ = len(v6)
        if true {
            if (true || true) {
                var v18 int8 = 81
                _ = v18
            } else {
                var v19 bool = (0 < 616)
                if false {
                    if v19 {
                        fmt.Println(v19, v1)
                        for fi20 := 0; fi20 < 13; fi20++ {
                            var v21 float32 = 1e10
                            if true {
                                var v22 int16 = -1
                                _ = len(v6)
                                var v23 uint = 42
                                _ = v22
                                _ = v23
                            } else {
                                switch "" {
                                case "hello":
                                    var v24 float64 = 3.14159
                                    _ = v24
                                default:
                                    var v25 uint32 = 4294967295
                                    _ = v25
                                }
                            }
                            _ = fi20
                            _ = v21
                        }
                    }
                    var v26 float64 = v3
                    _ = v26
                }
                var v27 uint64 = 18446744073709551615
                _ = v19
                _ = v27
            }
            v3 = v3
            v28 := GenS1{F0: 100}
            _ = v28
        } else {
            v2 = v2
            var v29 int = fi17
            var v30 int = (v1 & (v1 % (((fi17 + (v1 | fi17))) | 1)))
            _ = v29
            _ = v30
        }
        _ = fi17
    }
    var v31 float64 = v3
    v1 -= v1
    var v32 uint32 = 42
    var v33 int8 = int8(v32)
    var v34 int32 = -2147483648
    fmt.Println(v5, v4)
    for fi35 := 0; fi35 < 17; fi35++ {
        fmt.Println(fi35, v3, v5)
        v34 += (v34 & (^(v34)))
        v36 := GenS2{F0: 174, F1: v5}
        _ = fi35
        _ = v36
    }
    v33 -= (v33 | (v33 << (((v2 >> (((v2 ^ 4294967295)) & 63))) & 63)))
    fmt.Println(v31, v5, v3)
    for fi37 := 0; fi37 < 19; fi37++ {
        if (1 > 0) {
            for fi38 := 0; fi38 < 16; fi38++ {
                fmt.Println(v33, v34, v3)
                _ = fi38
            }
            v5 += 32767
        } else {
            v34 = int32(v32)
            var v39 int32 = v34
            _ = v39
        }
        fmt.Println(v31)
        var v40 int32 = (v34 & genFn0())
        _ = fi37
        _ = v40
    }
    var v41 int32 = v34
    for fi42 := 0; fi42 < 17; fi42++ {
        var v43 uint32 = 121
        if false {
            if len(v6) > 0 {
                v6[0] = v32
            }
        } else {
            for fi44 := 0; fi44 < 10; fi44++ {
                switch (v1 | fi42) {
                case 128:
                    if (v3 >= v31) {
                        v32 -= v32
                        var v45 uint = v2
                        switch -443 {
                        case 42:
                            var v46 int = (v1 << ((4294967295) & 63))
                            _ = v46
                        case 0:
                            v34 -= genFn0()
                        case -1:
                            if (3.14159 > v31) {
                                if false {
                                    if len(v6) >= 2 {
                                        v6 = v6[1:len(v6)]
                                    }
                                    switch "hello" {
                                    case "":
                                        switch ("a" + "hello") {
                                        case "foo bar":
                                            var v47 bool = (v4 <= (v4 - int64(v33)))
                                            _ = v47
                                        }
                                    case "foo bar":
                                        switch 0 {
                                        case 128:
                                            v3 += (v3 - (v31 / float64(1.5)))
                                        }
                                    default:
                                        fmt.Println(v5, fi44)
                                    }
                                }
                            } else {
                                if true {
                                    var v48 float64 = v31
                                    _ = v48
                                }
                                fmt.Println(v2, fi42, v2)
                            }
                        default:
                            var v49 uint32 = 405
                            _ = v49
                        }
                        v50 := GenS1{F0: 42}
                        v6 = append(v6, 65535)
                        _ = v45
                        _ = v50
                    }
                case 503:
                    var v51 int = v1
                    _ = v51
                default:
                    if ((1.5 - (-2.25 / 1e10)) < 0.1) {
                        v31 = float64((((-2.25 * -2.25) + 3.14159) * (((0.1 - (((-2.25 - (((3.14159 + ((1.5 / 0.1) + 1e10)) - 1.5) / 3.14159)) + 0.1) / float32(0.1))) - 1e10) - (1e10 - 3.14159))))
                        v2 = (v2 % (((v2 * 506)) | 1))
                        switch "a" {
                        case "a":
                            _ = len(v6)
                            v52 := GenS1{F0: (v33 | (v33 / ((-10) | 1)))}
                            if (int(v2) < fi44) {
                                switch 9223372036854775807 {
                                case -2147483648:
                                    v53 := GenS2{F0: 2, F1: 41}
                                    _ = v53
                                default:
                                    v43 -= (v32 | v43)
                                    v54 := GenS2{F0: 11, F1: v5}
                                    if len(v6) > 0 {
                                        v6[0] = v32
                                    }
                                    v32 = 1
                                    switch (fi42 & (fi42 % ((v1) | 1))) {
                                    case -706:
                                        switch (fi42 - fi44) {
                                        case 595:
                                            switch 9223372036854775807 {
                                            case 32768:
                                                fmt.Println(v33, v32)
                                            case 127:
                                                v55 := GenS2{F0: 127, F1: (^(v5))}
                                                _ = v55
                                            case -159:
                                                if false {
                                                    if ("foo bar" < "a") {
                                                        var v56 int32 = (v41 * 42)
                                                        if true {
                                                            var v57 uint = v2
                                                            _ = v57
                                                        } else {
                                                            var v58 uint64 = 2
                                                            v58 -= v58
                                                            _ = v58
                                                        }
                                                        _ = v56
                                                    } else {
                                                        if len(v6) >= 2 {
                                                            v6 = v6[1:len(v6)]
                                                        }
                                                        v59 := GenS1{F0: (^(v33))}
                                                        fmt.Println(v3)
                                                        _ = v59
                                                    }
                                                } else {
                                                    v3 = v31
                                                }
                                            default:
                                                v6 = append(v6, 2)
                                            }
                                        case 939:
                                            if len(v6) >= 2 {
                                                v6 = v6[1:len(v6)]
                                            }
                                        }
                                    case 42:
                                        if false {
                                            v32 = (^(v43))
                                            v3 = 1e10
                                        }
                                    case 9223372036854775807:
                                        switch "a" {
                                        case "foo bar":
                                            v60 := GenS2{F0: 65, F1: 32767}
                                            _ = v60
                                        case "hello":
                                            v61 := GenS2{F0: 2, F1: 0}
                                            _ = v61
                                        }
                                    }
                                    _ = v54
                                }
                            }
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                            _ = v52
                        case "\n\t":
                            v34 += (v41 * (v34 + (v34 | v34)))
                        case "":
                            fmt.Println(v31, v2)
                        default:
                            var v62 uint32 = v32
                            _ = v62
                        }
                    }
                }
                var v63 uint8 = 255
                var v64 int16 = -32768
                _ = fi44
                _ = v63
                _ = v64
            }
            v6 = append(v6, (v43 % (((v32 ^ uint32(v2))) | 1)))
            var v65 int64 = (v4 << ((v2) & 63))
            _ = v65
        }
        _ = fi42
        _ = v43
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v41
    _ = v5
    _ = v6
}
