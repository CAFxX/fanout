package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int64; F1 int32 }
type GenS4 struct { F0 bool; F1 int16 }
type GenS5 struct { F0 uint8; F1 uint32 }
type GenS6 struct { F0 uint; F1 int; F2 uint16 }
type GenS7 struct { F0 string }

func main() {
    v1 := []int{128, -484}
    v2 := []int{-746, 9223372036854775807, 1}
    var v3 bool = (!true)
    var v4 int64 = -9223372036854775808
    var v5 uint = 429
    v6 := make(map[string]bool)
    if len(v2) > 0 {
        v2[0] = 4294967295
    }
    var v7 int64 = v4
    switch -9223372036854775808 {
    case 246:
        v8 := GenS1{F0: 128, F1: (v7 | 42)}
        _ = v8
    case 961:
        switch "a" {
        case "hello":
            v1 = append(v1, -198)
        case "\n\t":
            var v9 float64 = 3.14159
            _ = v9
        default:
            delete(v6, ("a" + "foo bar"))
        }
    default:
        if (42 <= 32767) {
            if v3 {
                var v10 uint = v5
                v11 := GenS1{F0: 128, F1: (v4 | (v7 % ((v7) | 1)))}
                for fi12 := 0; fi12 < 18; fi12++ {
                    v3 = (-1 > 2147483647)
                    _ = fi12
                }
                _ = v10
                _ = v11
            } else {
                if (v3 && (120 > 18446744073709551615)) {
                    if v3 {
                        fmt.Println(v7)
                        fmt.Println(v7)
                        v7 = (v4 << ((v5) & 63))
                    }
                    _ = len(v2)
                } else {
                    var v13 int16 = -261
                    v14 := GenS1{F0: 0, F1: (v4 / ((v7) | 1))}
                    _ = v13
                    _ = v14
                }
                v15 := v6["\n\t"]
                _ = v15
            }
            switch "hello" {
            case "a":
                var v16 uint16 = 252
                _ = v16
            }
            for fi17 := 0; fi17 < 7; fi17++ {
                v5 = (^(v5))
                v18 := GenS1{F0: 42, F1: v4}
                _ = fi17
                _ = v18
            }
        } else {
            if (41 == 255) {
                switch "foo bar" {
                case "\n\t":
                    v4 += v4
                case "hello":
                    fmt.Println(v7, v4, v7)
                case "a":
                    v3 = v3
                }
                if (!v3) {
                    v19 := GenS1{F0: 73, F1: v4}
                    _ = len(v1)
                    _ = v19
                }
            }
            v3 = ((18446744073709551615 >= 902) || v3)
            var v20 string = ("\n\t" + "\n\t")
            _ = v20
        }
    }
    var v21 int = 962
    var v22 int16 = 1
    if ((v7 >= v4) || (!(v3 && v3))) {
        v23 := GenS1{F0: 2, F1: -9223372036854775808}
        fmt.Println(v21, v7, v21)
        v24 := GenS2{F0: 668}
        _ = v23
        _ = v24
    } else {
        if false {
            var v25 uint32 = 1
            var v26 uint32 = v25
            for fi27 := 0; fi27 < 8; fi27++ {
                v3 = v3
                _ = fi27
            }
            _ = v25
            _ = v26
        }
        v28 := GenS1{F0: 1, F1: v4}
        v22 = v22
        _ = v28
    }
    for fi29 := 0; fi29 < 12; fi29++ {
        v30 := GenS3{F0: v7, F1: -1}
        _ = fi29
        _ = v30
    }
    v21 -= 186
    _ = len(v1)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v7 += 42
    switch (v21 ^ v21) {
    case -131:
        for fi31 := 0; fi31 < 6; fi31++ {
            if (true && false) {
                if (!true) {
                    for fi32 := 0; fi32 < 7; fi32++ {
                        v33 := GenS4{F0: ((-123 >= 117) && v3), F1: (-(v22))}
                        _ = fi32
                        _ = v33
                    }
                    for fi34 := 0; fi34 < 1; fi34++ {
                        if v3 {
                            fmt.Println(v21)
                            v21 = fi31
                            v3 = (-9223372036854775808 >= (fi31 & v21))
                            var v35 uint16 = 2
                            var v36 uint16 = (v35 << ((v5) & 63))
                            _ = v35
                            _ = v36
                        } else {
                            v5 += v5
                            var v37 string = "a"
                            if ((!true) && (42 > -128)) {
                                var v38 bool = ("foo bar" > (v37 + (v37 + "foo bar")))
                                var v39 int8 = 100
                                var v40 uint32 = 2147483648
                                if v38 {
                                    v38 = (v3 && ((!v3) || v3))
                                    v41 := v6[("foo bar" + v37)]
                                    var v42 uint16 = 0
                                    _ = v41
                                    _ = v42
                                } else {
                                    switch "hello" {
                                    case "foo bar":
                                        v40 -= 2
                                    case "hello":
                                        var v43 uint64 = 42
                                        _ = v43
                                    default:
                                        switch "a" {
                                        case "foo bar":
                                            v44 := v6[""]
                                            _ = v44
                                        case "\n\t":
                                            if ((42 >= (^(v40))) || (((v37 + v37) + v37) != v37)) {
                                                v37 = v37
                                                var v45 uint8 = 226
                                                v21 = (v21 - (fi34 ^ (v21 >> (((v5 >> ((v5) & 63))) & 63))))
                                                _ = v45
                                            }
                                        case "":
                                            var v46 float32 = 1e10
                                            _ = v46
                                        default:
                                            v47 := GenS5{F0: 1, F1: v40}
                                            _ = v47
                                        }
                                    }
                                    v4 += v7
                                }
                                v39 = 100
                                _ = v38
                                _ = v39
                                _ = v40
                            } else {
                                fmt.Println(v37, v5, fi31)
                                v37 = v37
                            }
                            _ = v37
                        }
                        v7 -= v4
                        var v48 uint32 = 413
                        _ = fi34
                        _ = v48
                    }
                    switch fi31 {
                    case 9223372036854775807:
                        v7 = v7
                    case -2147483648:
                        for fi49 := 0; fi49 < 12; fi49++ {
                            switch v21 {
                            case 127:
                                v6["a"] = ((3.14159 >= 3.14159) || (int16(v7) != (^(v22))))
                            default:
                                v50 := v6[""]
                                _ = v50
                            }
                            if v3 {
                                if v3 {
                                    var v51 uint8 = 128
                                    fmt.Println(v3, v21)
                                    _ = v51
                                }
                                v52 := GenS6{F0: v5, F1: fi49, F2: 1}
                                var v53 int16 = (v22 | 32767)
                                _ = v52
                                _ = v53
                            } else {
                                v3 = (!v3)
                                if true {
                                    v5 = (v5 ^ (v5 >> ((v5) & 63)))
                                }
                                var v54 int = (fi31 + fi49)
                                _ = v54
                            }
                            _ = fi49
                        }
                    case 32767:
                        v55 := GenS4{F0: (!(v22 <= v22)), F1: 709}
                        _ = v55
                    }
                } else {
                    _ = len(v6)
                    fmt.Println(v3, v21)
                    var v56 uint32 = 0
                    _ = v56
                }
                v57 := GenS7{F0: "hello"}
                fmt.Println(v4, v3)
                _ = v57
            }
            v4 = 1
            _ = fi31
        }
    default:
        v21 -= v21
    }
    var v58 int32 = 42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v58:", v58)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v58
    _ = v6
    _ = v7
}
