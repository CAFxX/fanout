package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32; F2 uint8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int8; F1 uint16; F2 uint }
type GenS4 struct { F0 bool; F1 int16 }
type GenS5 struct { F0 int32; F1 int16; F2 uint64 }

func main() {
    var v1 uint16 = 65535
    var v2 bool = false
    var v3 int32 = 1
    var v4 float64 = 0.1
    v5 := []int{-304, 347}
    switch -9223372036854775808 {
    case 714:
        v3 = (v3 & v3)
    case 32767:
        var v6 uint32 = 0
        _ = v6
    case 391:
        v3 += (-(v3))
    default:
        for fi7 := 0; fi7 < 5; fi7++ {
            var v8 string = (((("" + (("\n\t" + (("\n\t" + "") + "foo bar")) + "hello")) + "") + ("a" + "")) + "a")
            var v9 uint8 = 128
            if true {
                v8 = ((v8 + "hello") + v8)
                fmt.Println(v9, v2, v3)
            } else {
                var v10 uint16 = (^(v1))
                _ = v10
            }
            _ = fi7
            _ = v8
            _ = v9
        }
    }
    var v11 int32 = (-(v3))
    var v12 string = (("a" + "a") + "")
    v3 = -62
    v11 -= v3
    if (42 < 361) {
        var v13 int16 = -32768
        if (-1 == 32767) {
            var v14 uint16 = 279
            v1 = (v14 | (v14 + (v1 | v14)))
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = v14
        }
        if (!true) {
            if v2 {
                var v15 int8 = -125
                _ = v15
            } else {
                var v16 uint = 255
                var v17 float64 = ((1.5 + v4) - float64(1e10))
                _ = v16
                _ = v17
            }
            switch v12 {
            case "foo bar":
                _ = len(v5)
            case "a":
                v18 := GenS1{F0: false, F1: 2147483648, F2: 2}
                _ = v18
            case "\n\t":
                for fi19 := 0; fi19 < 3; fi19++ {
                    v4 = (v4 + float64(1e10))
                    var v20 uint64 = 42
                    _ = fi19
                    _ = v20
                }
            default:
                for fi21 := 0; fi21 < 14; fi21++ {
                    for fi22 := 0; fi22 < 20; fi22++ {
                        switch "hello" {
                        case "hello":
                            v1 = 2
                        case "foo bar":
                            if ((v11 >> ((1) & 63)) < (-(v3))) {
                                if (0 >= 42) {
                                    if v2 {
                                        v23 := GenS1{F0: (((!(v2 && false)) && (v2 || true)) || ((v12 <= (v12 + (v12 + v12))) || v2)), F1: 981, F2: 127}
                                        _ = v23
                                    } else {
                                        v24 := GenS2{F0: v11}
                                        _ = v24
                                    }
                                    v4 += 3.14159
                                    v5 = append(v5, fi21)
                                } else {
                                    _ = len(v5)
                                    var v25 int8 = 0
                                    v1 -= v1
                                    var v26 uint = 843
                                    v5 = append(v5, (fi22 << ((v26) & 63)))
                                    v12 = "foo bar"
                                    fmt.Println(fi22)
                                    _ = v25
                                    _ = v26
                                }
                                v3 = v11
                                var v27 bool = (false && (!(!(-2147483648 != v3))))
                                _ = v27
                            }
                        case "\n\t":
                            var v28 string = "foo bar"
                            _ = v28
                        }
                        _ = fi22
                    }
                    v29 := GenS2{F0: -1}
                    v11 = -2147483648
                    _ = fi21
                    _ = v29
                }
            }
            fmt.Println(v1)
        }
        _ = v13
    } else {
        if false {
            v5 = append(v5, -2147483648)
            switch ((v12 + ((v12 + "") + (("" + v12) + v12))) + (v12 + (("foo bar" + "") + "a"))) {
            case "foo bar":
                var v30 string = "\n\t"
                _ = v30
            case "":
                v31 := GenS1{F0: ((!v2) || ((128 >= 60) || true)), F1: 832, F2: 255}
                _ = v31
            case "\n\t":
                if v2 {
                    fmt.Println(v4, v2, v11)
                    v32 := GenS1{F0: (v2 || (v2 && (3.14159 <= -2.25))), F1: 0, F2: 128}
                    var v33 uint64 = 0
                    _ = v32
                    _ = v33
                }
            }
        } else {
            v4 -= v4
            for fi34 := 0; fi34 < 1; fi34++ {
                if true {
                    v3 = (v3 ^ (^(v3)))
                    switch fi34 {
                    case 128:
                        var v35 uint8 = 42
                        _ = v35
                    case -9223372036854775808:
                        v5 = append(v5, (fi34 / (((^(fi34))) | 1)))
                    }
                    for fi36 := 0; fi36 < 5; fi36++ {
                        v4 = (((0.0 + 1e10) / v4) - v4)
                        if v2 {
                            v12 = v12
                        } else {
                            if v2 {
                                v3 += int32(fi34)
                                v37 := GenS3{F0: -38, F1: v1, F2: 327}
                                var v38 int = fi36
                                _ = v37
                                _ = v38
                            }
                            v5 = append(v5, fi36)
                            if false {
                                var v39 float32 = (1e10 - -2.25)
                                _ = v39
                            } else {
                                fmt.Println(fi36)
                            }
                        }
                        _ = fi36
                    }
                }
                var v40 int64 = -776
                _ = fi34
                _ = v40
            }
            for fi41 := 0; fi41 < 14; fi41++ {
                switch fi41 {
                case 863:
                    if len(v5) > 0 {
                        v5[0] = fi41
                    }
                case -9223372036854775808:
                    v42 := GenS4{F0: v2, F1: -731}
                    _ = v42
                default:
                    var v43 int32 = v11
                    _ = v43
                }
                var v44 uint16 = v1
                for fi45 := 0; fi45 < 16; fi45++ {
                    switch -2147483648 {
                    case 4294967295:
                        v2 = ((v44 >> ((18446744073709551615) & 63)) < (v44 / ((v44) | 1)))
                    default:
                        if (!((9223372036854775807 > -1) || false)) {
                            if (2147483648 < 0) {
                                var v46 uint = 255
                                v47 := GenS2{F0: 0}
                                v48 := GenS5{F0: (v3 | (^(v11))), F1: int16(v46), F2: 809}
                                _ = v46
                                _ = v47
                                _ = v48
                            }
                        }
                    }
                    _ = fi45
                }
                _ = fi41
                _ = v44
            }
        }
        var v49 float32 = -2.25
        var v50 float32 = v49
        _ = v49
        _ = v50
    }
    var v51 int32 = (v3 * (v11 % (((v3 ^ v3)) | 1)))
    for fi52 := 0; fi52 < 14; fi52++ {
        var v53 int8 = -1
        _ = fi52
        _ = v53
    }
    v4 = v4
    if v2 {
        if false {
            fmt.Println(v4, v11, v4)
        }
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v51
}
