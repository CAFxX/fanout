package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
type GenS2 struct { F0 uint32; F1 uint16 }
type GenS3 struct { F0 uint; F1 int32; F2 string }
type GenS4 struct { F0 int; F1 int32 }
type GenS5 struct { F0 int; F1 uint64; F2 int32 }
type GenS6 struct { F0 string; F1 int32; F2 uint32 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 uint64; F1 uint64; F2 int32 }
func genFn0(p0 uint64) uint32 {
    if (!(0 == -419)) {
        fmt.Println(p0)
        for fi4 := 0; fi4 < 1; fi4++ {
            v5 := GenS1{F0: 59, F1: 0}
            p0 = 1
            if (0.0 < float32(((float64(0.0) - (3.14159 * (0.0 - 1.5))) + 1e10))) {
                switch 47 {
                case 454:
                    v6 := GenS2{F0: 0, F1: 255}
                    _ = v6
                }
                for fi7 := 0; fi7 < 12; fi7++ {
                    fmt.Println(p0, fi4)
                    var v8 string = ("hello" + ("\n\t" + "foo bar"))
                    _ = fi7
                    _ = v8
                }
                var v9 float64 = 0.0
                _ = v9
            } else {
                var v10 float32 = 0.0
                v11 := GenS1{F0: 560, F1: 65535}
                for fi12 := 0; fi12 < 15; fi12++ {
                    fmt.Println(p0)
                    var v13 float32 = v10
                    var v14 int16 = 0
                    switch fi4 {
                    case 61:
                        if true {
                            if (true && false) {
                                v15 := GenS2{F0: 4294967295, F1: 186}
                                _ = v15
                            }
                        } else {
                            switch fi12 {
                            case 32768:
                                if true {
                                    v10 += (v13 / v10)
                                } else {
                                    v10 += v10
                                    var v16 uint = 1
                                    _ = v16
                                }
                            case 963:
                                var v17 string = "foo bar"
                                _ = v17
                            case -1:
                                switch "a" {
                                case "\n\t":
                                    var v18 float32 = v13
                                    _ = v18
                                case "foo bar":
                                    v13 -= (0.1 + (v13 * v10))
                                }
                            default:
                                v19 := []int{32767, (fi4 >> ((255) & 63))}
                                _ = v19
                            }
                            v20 := GenS3{F0: uint(p0), F1: 1, F2: "hello"}
                            _ = v20
                        }
                    case -9223372036854775808:
                        switch 0 {
                        case 0:
                            fmt.Println(p0)
                        case -403:
                            v14 -= 1000
                        default:
                            var v21 int = (fi12 | (fi4 * -2147483648))
                            _ = v21
                        }
                    case 9223372036854775807:
                        if false {
                            v10 -= 1e10
                            v22 := GenS1{F0: 32767, F1: 1}
                            _ = v22
                        } else {
                            v23 := GenS4{F0: (fi12 & fi4), F1: 117}
                            _ = v23
                        }
                    default:
                        v24 := GenS5{F0: fi4, F1: p0, F2: 1}
                        _ = v24
                    }
                    _ = fi12
                    _ = v13
                    _ = v14
                }
                _ = v10
                _ = v11
            }
            _ = fi4
            _ = v5
        }
        v25 := GenS6{F0: ("hello" + "\n\t"), F1: 149, F2: 326}
        _ = v25
    }
    v26 := GenS7{F0: 100}
    _ = p0
    _ = v26
    return 2147483648
}
func genFn1() uint64 {
    v27 := []int{4294967295, 32767, 457, -55}
    if (-1 != 9223372036854775807) {
        var v28 int16 = 42
        var v29 bool = true
        if (!(false || v29)) {
            var v30 int16 = v28
            _ = len(v27)
            if (((0.0 / 3.14159) >= 0.1) || false) {
                v31 := GenS8{F0: 42, F1: 2, F2: -1}
                _ = v31
            } else {
                if (v28 < (v30 >> ((4294967295) & 63))) {
                    _ = len(v27)
                    for fi32 := 0; fi32 < 16; fi32++ {
                        var v33 int = 1
                        v27 = append(v27, (v33 % (((v33 + v33)) | 1)))
                        _ = fi32
                        _ = v33
                    }
                    v28 = (v28 / ((v28) | 1))
                }
                for fi34 := 0; fi34 < 4; fi34++ {
                    switch ((("a" + ("hello" + "foo bar")) + "\n\t") + ("" + (((("hello" + "hello") + "") + "") + ""))) {
                    case "\n\t":
                        if (400 >= 1) {
                            switch 189 {
                            case 971:
                                fmt.Println(v29, v28)
                            }
                            if len(v27) > 0 {
                                v27[0] = (fi34 >> ((255) & 63))
                            }
                        }
                    default:
                        var v35 int = (fi34 % (((fi34 << ((65535) & 63))) | 1))
                        _ = v35
                    }
                    _ = fi34
                }
            }
            _ = v30
        } else {
            if (!(v29 || (2147483647 == 2147483647))) {
                v28 = v28
                var v36 int16 = v28
                for fi37 := 0; fi37 < 10; fi37++ {
                    _ = len(v27)
                    _ = fi37
                }
                _ = v36
            } else {
                if (948 < -322) {
                    v28 = (-(v28))
                    if len(v27) > 0 {
                        v27[0] = -267
                    }
                    var v38 uint = 42
                    _ = v38
                }
                if len(v27) >= 2 {
                    v27 = v27[1:len(v27)]
                }
            }
        }
        _ = v28
        _ = v29
    }
    var v39 int16 = -309
    var v40 uint32 = 2147483648
    _ = v27
    _ = v39
    _ = v40
    return 0
}

func main() {
    v1 := []int{0, -807, 128}
    var v2 int8 = 0
    var v3 int8 = (^(v2))
    if false {
        v41 := GenS8{F0: 425, F1: 42, F2: -536}
        _ = v41
    }
    v3 = v3
    switch "a" {
    case "a":
        var v42 int64 = 9223372036854775807
        _ = v42
    case "":
        if len(v1) > 0 {
            v1[0] = 608
        }
    case "\n\t":
        switch 32767 {
        case -786:
            if (42 < -32768) {
                if false {
                    var v43 int64 = -608
                    _ = v43
                } else {
                    v2 -= v2
                    v44 := GenS4{F0: 0, F1: 1}
                    v3 -= v2
                    _ = v44
                }
            }
        default:
            for fi45 := 0; fi45 < 14; fi45++ {
                v46 := GenS8{F0: 0, F1: 18446744073709551615, F2: 0}
                var v47 int16 = -655
                _ = fi45
                _ = v46
                _ = v47
            }
        }
    default:
        var v48 float32 = 0.0
        _ = v48
    }
    fmt.Println(v3, v3)
    v1 = append(v1, 4294967295)
    for fi49 := 0; fi49 < 15; fi49++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        var v50 uint16 = 255
        switch ("" + ((("hello" + "") + "") + ("a" + "foo bar"))) {
        case "\n\t":
            v51 := GenS2{F0: 0, F1: (v50 * (v50 & (v50 - v50)))}
            _ = v51
        default:
            for fi52 := 0; fi52 < 13; fi52++ {
                if false {
                    switch "a" {
                    case "foo bar":
                        var v53 uint8 = 42
                        _ = v53
                    case "hello":
                        var v54 int16 = 32767
                        _ = v54
                    case "a":
                        var v55 bool = false
                        _ = v55
                    default:
                        switch fi49 {
                        case 1:
                            switch ("\n\t" + ((("foo bar" + "a") + "hello") + "a")) {
                            case "a":
                                fmt.Println(v50, v50)
                            case "foo bar":
                                v1 = append(v1, fi52)
                            }
                        }
                    }
                } else {
                    var v56 uint64 = uint64(v50)
                    var v57 uint8 = 1
                    v58 := GenS3{F0: 4294967295, F1: 648, F2: "foo bar"}
                    _ = v56
                    _ = v57
                    _ = v58
                }
                _ = fi52
            }
        }
        _ = fi49
        _ = v50
    }
    for fi59 := 0; fi59 < 18; fi59++ {
        var v60 uint64 = 4294967295
        _ = fi59
        _ = v60
    }
    var v61 int16 = -241
    v61 = (v61 + v61)
    var v62 float32 = 3.14159
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v61:", v61)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v2
    _ = v3
    _ = v61
    _ = v62
}
