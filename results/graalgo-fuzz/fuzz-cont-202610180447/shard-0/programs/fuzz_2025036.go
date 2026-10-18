package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string }
type GenS2 struct { F0 int; F1 bool; F2 int64 }
type GenS3 struct { F0 uint64; F1 int64; F2 uint8 }
type GenS4 struct { F0 uint64; F1 uint16 }
type GenS5 struct { F0 uint8; F1 uint32 }
type GenS6 struct { F0 int64; F1 uint32; F2 int32 }
type GenS7 struct { F0 int8; F1 uint64 }
func genFn0(p0 int8, p1 string, p2 bool) string {
    p1 = p1
    if false {
        var v7 uint16 = 2
        fmt.Println(p1, p2, v7)
        switch p1 {
        case "foo bar":
            v7 = 32767
        default:
            v8 := GenS1{F0: 71, F1: "hello"}
            _ = v8
        }
        _ = v7
    }
    for fi9 := 0; fi9 < 13; fi9++ {
        p2 = (((fi9 >= (fi9 >> ((1) & 63))) || p2) && true)
        for fi10 := 0; fi10 < 12; fi10++ {
            fmt.Println(p0)
            var v11 bool = (p2 || false)
            p2 = false
            _ = fi10
            _ = v11
        }
        fmt.Println(p1, p2, p0)
        _ = fi9
    }
    _ = p0
    _ = p1
    _ = p2
    return p1
}
func genFn1(p0 int8, p1 uint16) int8 {
    if false {
        if true {
            p1 = p1
            p1 = p1
            v12 := []int{120}
            _ = v12
        } else {
            if true {
                switch "" {
                case "":
                    v13 := GenS2{F0: -2147483648, F1: (("hello" + ("\n\t" + "hello")) < "\n\t"), F2: -1}
                    _ = v13
                case "foo bar":
                    v14 := GenS3{F0: 648, F1: 42, F2: 1}
                    _ = v14
                }
            }
        }
    } else {
        v15 := []int{}
        v16 := GenS2{F0: 42, F1: (!(42 >= 2147483647)), F2: -9223372036854775808}
        _ = v15
        _ = v16
    }
    p1 = (p1 / ((p1) | 1))
    _ = p0
    _ = p1
    return (-(p0))
}

func main() {
    var v1 bool = false
    var v2 int64 = -9223372036854775808
    v3 := []int{-280}
    var v4 int8 = 42
    var v5 int64 = v2
    var v6 int32 = -749
    v17 := GenS4{F0: 42, F1: 475}
    var v18 uint32 = 42
    var v19 int = 32767
    var v20 int16 = 1000
    if v1 {
        _ = len(v3)
    } else {
        v18 = (v18 & v18)
        var v21 uint = 968
        _ = v21
    }
    v2 = (-(v5))
    if true {
        v22 := GenS5{F0: 13, F1: (v18 + (v18 * 42))}
        v20 += (v20 << ((1) & 63))
        _ = v22
    } else {
        v3 = append(v3, (v19 >> ((4294967295) & 63)))
    }
    if v1 {
        for fi23 := 0; fi23 < 4; fi23++ {
            if (!false) {
                v24 := GenS2{F0: 0, F1: (v1 && v1), F2: (v2 + v2)}
                var v25 uint32 = v18
                _ = v24
                _ = v25
            } else {
                switch "hello" {
                case "a":
                    var v26 uint64 = 4294967295
                    _ = v26
                case "hello":
                    v18 -= v18
                case "\n\t":
                    for fi27 := 0; fi27 < 1; fi27++ {
                        switch "" {
                        case "\n\t":
                            switch "" {
                            case "hello":
                                fmt.Println(fi27)
                            }
                            switch (^(v19)) {
                            case -54:
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            case 204:
                                v28 := GenS6{F0: (v5 / (((v5 % ((int64(v18)) | 1))) | 1)), F1: (^(v18)), F2: (v6 % ((2147483647) | 1))}
                                _ = v28
                            case -9223372036854775808:
                                if false {
                                    switch ("a" + "foo bar") {
                                    case "hello":
                                        if true {
                                            var v29 uint32 = (v18 >> ((751) & 63))
                                            v3 = append(v3, 0)
                                            var v30 float32 = (1e10 * 0.0)
                                            var v31 uint64 = 223
                                            var v32 uint32 = v18
                                            _ = v29
                                            _ = v30
                                            _ = v31
                                            _ = v32
                                        } else {
                                            v19 -= 1
                                            var v33 int = fi27
                                            _ = v33
                                        }
                                    case "foo bar":
                                        var v34 uint8 = 1
                                        _ = v34
                                    case "":
                                        v3 = append(v3, (v19 ^ (fi23 >> ((18446744073709551615) & 63))))
                                    default:
                                        v4 = (v4 % ((v4) | 1))
                                    }
                                    v18 += (v18 & (v18 ^ (v18 + v18)))
                                    var v35 string = ""
                                    _ = v35
                                } else {
                                    switch 103 {
                                    case 1:
                                        v20 += v20
                                    case 32767:
                                        if (2147483648 >= (v18 * 260)) {
                                            var v36 uint64 = 0
                                            v36 += 4294967295
                                            if (!(genFn1(-56, 855) > v4)) {
                                                var v37 int8 = genFn1(v4, 1)
                                                switch "foo bar" {
                                                case "\n\t":
                                                    switch (-(fi23)) {
                                                    case 48:
                                                        fmt.Println(fi27, v4, v36)
                                                    case 128:
                                                        var v38 string = "hello"
                                                        _ = v38
                                                    default:
                                                        var v39 int = v19
                                                        _ = v39
                                                    }
                                                }
                                                v2 = (^(v2))
                                                if (!v1) {
                                                    v5 = (v2 + (v2 / ((v5) | 1)))
                                                }
                                                _ = v37
                                            }
                                            _ = v36
                                        }
                                    case -196:
                                        var v40 uint = 2
                                        _ = v40
                                    default:
                                        var v41 int16 = (v20 / (((v20 ^ (-(v20)))) | 1))
                                        _ = v41
                                    }
                                    fmt.Println(v6, v2)
                                }
                            default:
                                v3 = append(v3, (^(v19)))
                            }
                            v18 = v18
                        case "foo bar":
                            v42 := GenS2{F0: fi27, F1: (128 < (^(v19))), F2: -1}
                            _ = v42
                        case "hello":
                            v20 = (-(v20))
                        }
                        if (((v2 / (((v2 >> ((65535) & 63))) | 1)) != 0) || true) {
                            switch 947 {
                            case 2147483647:
                                v43 := GenS2{F0: (v19 % ((fi23) | 1)), F1: (1.5 != (0.0 * float32(0.0))), F2: (^(v5))}
                                _ = v43
                            case 732:
                                v4 = (v4 * 1)
                            case 4294967295:
                                v18 -= 0
                            }
                            v4 = (v4 + -128)
                            v18 = v18
                        } else {
                            v6 += v6
                        }
                        _ = fi27
                    }
                }
                v2 = 87
                if (!v1) {
                    v19 += (v19 / ((-9223372036854775808) | 1))
                }
            }
            if ((v5 << ((0) & 63)) <= v5) {
                v4 += v4
                var v44 int32 = -1
                if (!v1) {
                    v20 += v20
                    var v45 int16 = (v20 | -393)
                    _ = v45
                }
                _ = v44
            } else {
                if v1 {
                    var v46 int8 = -77
                    var v47 int8 = (v4 << ((2) & 63))
                    _ = v46
                    _ = v47
                } else {
                    v48 := GenS6{F0: (v5 - (v2 + (^(v2)))), F1: v18, F2: v6}
                    var v49 float64 = -2.25
                    _ = v48
                    _ = v49
                }
            }
            v1 = v1
            _ = fi23
        }
    }
    if (!((v4 + (v4 - v4)) < (^(v4)))) {
        fmt.Println(v20, v19)
    } else {
        if (!(154 < 581)) {
            var v50 uint8 = 0
            _ = v50
        } else {
            fmt.Println(v6, v4)
        }
        for fi51 := 0; fi51 < 5; fi51++ {
            v19 -= 1
            var v52 uint16 = 1
            _ = fi51
            _ = v52
        }
    }
    for fi53 := 0; fi53 < 12; fi53++ {
        fmt.Println(fi53, v19)
        _ = fi53
    }
    for fi54 := 0; fi54 < 6; fi54++ {
        var v55 int64 = v5
        v56 := GenS7{F0: v4, F1: 18446744073709551615}
        if (v1 || (v6 > v6)) {
            var v57 bool = false
            v19 -= (v19 % ((459) | 1))
            var v58 uint64 = 2
            _ = v57
            _ = v58
        } else {
            v20 -= v20
            var v59 int64 = v2
            var v60 uint16 = 255
            _ = v59
            _ = v60
        }
        _ = fi54
        _ = v55
        _ = v56
    }
    var v61 float64 = -2.25
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v61
}
