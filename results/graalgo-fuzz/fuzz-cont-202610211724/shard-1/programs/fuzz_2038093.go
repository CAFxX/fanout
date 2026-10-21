package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 int8 }
type GenS2 struct { F0 uint8; F1 int32 }
type GenS3 struct { F0 int32; F1 uint }
type GenS4 struct { F0 uint; F1 bool; F2 uint64 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 uint32 }
func genFn0(p0 uint, p1 int64) uint64 {
    var v4 uint32 = 571
    v5 := []int{-488, 128}
    v6 := GenS1{F0: 4, F1: 1, F2: -1}
    var v7 bool = false
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return 946
}
func genFn1(p0 int64, p1 int32, p2 int) bool {
    p1 += 2147483647
    v8 := []int{p2, int(p0)}
    for fi9 := 0; fi9 < 19; fi9++ {
        if true {
            p1 = 0
            switch -615 {
            case 1:
                v10 := GenS2{F0: 255, F1: (p1 | p1)}
                _ = v10
            case -773:
                p1 -= (p1 & -349)
            case 381:
                var v11 int8 = int8(p0)
                _ = v11
            default:
                for fi12 := 0; fi12 < 2; fi12++ {
                    if (false && true) {
                        switch p2 {
                        case -1:
                            p0 = p0
                        case 1:
                            _ = len(v8)
                            p1 = (p1 ^ -2147483648)
                            p0 -= (p0 ^ 0)
                        case 42:
                            if len(v8) > 0 {
                                v8[0] = (p2 % ((p2) | 1))
                            }
                        default:
                            var v13 uint64 = 1
                            _ = v13
                        }
                        if true {
                            v14 := GenS3{F0: -946, F1: 65535}
                            p1 -= p1
                            _ = v14
                        }
                        var v15 uint32 = 852
                        _ = v15
                    } else {
                        if true {
                            p2 -= fi12
                            if ((p1 ^ p1) == (-(p1))) {
                                if true {
                                    fmt.Println(fi12, p0)
                                }
                            }
                            _ = len(v8)
                        }
                    }
                    _ = fi12
                }
            }
        }
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
        _ = len(v8)
        _ = fi9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return (p0 != p0)
}

func main() {
    v1 := []int{0}
    var v2 uint = 255
    var v3 bool = (-1 > 530)
    for fi16 := 0; fi16 < 7; fi16++ {
        if true {
            fmt.Println(v3)
        }
        v17 := GenS2{F0: 156, F1: 2147483647}
        if len(v1) > 0 {
            v1[0] = 855
        }
        _ = fi16
        _ = v17
    }
    var v18 uint64 = 304
    if (0 == 1000) {
        switch 42 {
        case 42:
            switch (((("hello" + "a") + (("\n\t" + "a") + "hello")) + "a") + ("foo bar" + "foo bar")) {
            case "hello":
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
        case 128:
            v18 += (v18 << ((v2) & 63))
        }
        if v3 {
            v2 += 1
            var v19 uint8 = 255
            _ = v19
        } else {
            if len(v1) > 0 {
                v1[0] = 2147483647
            }
            if len(v1) > 0 {
                v1[0] = 2147483647
            }
        }
        if v3 {
            var v20 uint = v2
            switch "" {
            case "\n\t":
                v2 = v2
            case "hello":
                if v3 {
                    fmt.Println(v3)
                }
            default:
                _ = len(v1)
            }
            fmt.Println(v18, v2, v18)
            _ = v20
        }
    } else {
        for fi21 := 0; fi21 < 11; fi21++ {
            for fi22 := 0; fi22 < 3; fi22++ {
                v23 := GenS4{F0: (^(v2)), F1: ((fi21 << ((1) & 63)) < int(v2)), F2: (v18 % (((v18 | (v18 & 18446744073709551615))) | 1))}
                var v24 uint64 = v18
                _ = fi22
                _ = v23
                _ = v24
            }
            _ = fi21
        }
    }
    for fi25 := 0; fi25 < 8; fi25++ {
        v26 := GenS5{F0: 1}
        var v27 float32 = (1.5 - (1.5 - (-2.25 - 0.1)))
        _ = fi25
        _ = v26
        _ = v27
    }
    v28 := GenS5{F0: -32768}
    switch -9223372036854775808 {
    case 2147483647:
        if ((128 <= 42) && v3) {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            if len(v1) > 0 {
                v1[0] = 9223372036854775807
            }
            fmt.Println(v18, v18, v3)
        }
    case 32768:
        var v29 int64 = 423
        _ = v29
    }
    v30 := GenS6{F0: v3}
    v31 := GenS7{F0: 65535}
    var v32 uint32 = 2147483647
    switch ("\n\t" + ("a" + ((("hello" + "") + (("foo bar" + "\n\t") + "a")) + "a"))) {
    case "":
        for fi33 := 0; fi33 < 4; fi33++ {
            v1 = append(v1, (fi33 % (((fi33 | fi33)) | 1)))
            _ = fi33
        }
    case "foo bar":
        v32 = 4294967295
    default:
        fmt.Println(v18, v2, v32)
    }
    var v34 int32 = 2147483647
    if (!v3) {
        for fi35 := 0; fi35 < 1; fi35++ {
            for fi36 := 0; fi36 < 20; fi36++ {
                switch fi36 {
                case 4294967295:
                    switch "hello" {
                    case "\n\t":
                        var v37 int = fi35
                        _ = v37
                    }
                case 0:
                    if ((0.0 <= 0.0) && (1e10 == (0.1 * 3.14159))) {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    } else {
                        _ = len(v1)
                        if v3 {
                            if true {
                                var v38 bool = v3
                                v18 = (v18 ^ 675)
                                _ = v38
                            } else {
                                switch fi36 {
                                case -1:
                                    v18 = genFn0((v2 + 0), 579)
                                case 42:
                                    fmt.Println(v18)
                                default:
                                    var v39 bool = (43 <= 127)
                                    fmt.Println(v32)
                                    if v39 {
                                        fmt.Println(v39, v32)
                                        var v40 int = 32768
                                        _ = v40
                                    } else {
                                        v3 = v39
                                        var v41 int16 = 42
                                        if false {
                                            v41 = v41
                                            var v42 uint32 = v32
                                            _ = v42
                                        }
                                        _ = v41
                                    }
                                    _ = v39
                                }
                                var v43 int16 = 0
                                if (fi36 != -1) {
                                    if ((0.1 + -2.25) == 0.0) {
                                        v3 = v3
                                        v32 -= v32
                                        var v44 bool = false
                                        _ = v44
                                    } else {
                                        fmt.Println(v32, fi36)
                                    }
                                    if v3 {
                                        switch "a" {
                                        case "hello":
                                            if len(v1) > 0 {
                                                v1[0] = 0
                                            }
                                        case "a":
                                            if (!true) {
                                                var v45 uint = (v2 / ((v2) | 1))
                                                _ = v45
                                            } else {
                                                var v46 bool = (!(v3 && v3))
                                                _ = v46
                                            }
                                        case "":
                                            fmt.Println(fi36)
                                        }
                                        var v47 uint8 = 186
                                        _ = v47
                                    } else {
                                        _ = len(v1)
                                        v32 -= 4294967295
                                        v2 = v2
                                        v3 = v3
                                        var v48 uint32 = (v32 & (v32 & v32))
                                        v49 := GenS8{F0: genFn0(v2, -1)}
                                        _ = v48
                                        _ = v49
                                    }
                                }
                                _ = v43
                            }
                        } else {
                            if false {
                                v50 := GenS3{F0: v34, F1: uint(v32)}
                                v51 := GenS9{F0: true}
                                fmt.Println(v18, v2, v2)
                                _ = v50
                                _ = v51
                            } else {
                                v34 += v34
                                fmt.Println(v34, v18, v34)
                                v18 -= 18446744073709551615
                            }
                        }
                    }
                }
                if (v32 > (v32 / (((v32 & 2147483647)) | 1))) {
                    if false {
                        if v3 {
                            switch (^(fi35)) {
                            case 128:
                                var v52 float32 = ((-2.25 / (((1.5 + 3.14159) * 1e10) - (float32((0.0 - 1.5)) + 0.0))) - 0.0)
                                _ = v52
                            }
                            if (("foo bar" + "a") > "") {
                                v1 = append(v1, 2147483647)
                                v32 = 65535
                                if (v3 && v3) {
                                    v2 = 488
                                    fmt.Println(v32, v34)
                                    v32 -= v32
                                }
                            }
                        }
                        v34 = (v34 << ((v2) & 63))
                        if (0 <= 731) {
                            var v53 bool = ((v3 && (!v3)) || true)
                            if false {
                                v18 -= genFn0(v2, 42)
                            }
                            if (fi36 != fi36) {
                                v54 := GenS10{F0: (v32 * (v32 | 383))}
                                _ = v54
                            } else {
                                fmt.Println(fi36, v18, v18)
                                fmt.Println(v3, fi35, v32)
                            }
                            _ = v53
                        }
                        v2 = (v2 % (((^(v2))) | 1))
                    }
                    var v55 uint32 = 42
                    _ = v55
                }
                if v3 {
                    v18 = (v18 - 775)
                    var v56 uint32 = v32
                    _ = len(v1)
                    _ = v56
                }
                _ = fi36
            }
            var v57 uint8 = uint8(v18)
            _ = fi35
            _ = v57
        }
        if v3 {
            fmt.Println(v3)
            fmt.Println(v2, v3, v18)
            fmt.Println(v18, v32, v3)
        }
        if (1 > (^(v18))) {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        } else {
            var v58 int64 = -9223372036854775808
            _ = v58
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v18
    _ = v2
    _ = v28
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v34
}
