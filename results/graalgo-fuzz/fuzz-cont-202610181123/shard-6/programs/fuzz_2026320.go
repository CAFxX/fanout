package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 uint16 }
type GenS2 struct { F0 int32; F1 int64 }
type GenS3 struct { F0 uint; F1 uint16 }
type GenS4 struct { F0 uint16; F1 int; F2 bool }

func main() {
    var v1 uint8 = 42
    var v2 uint32 = 4294967295
    var v3 uint16 = uint16(v1)
    v4 := make(map[string]uint32)
    v5 := make(map[string]string)
    var v6 uint = 111
    var v7 float64 = 1e10
    fmt.Println(v3)
    switch "\n\t" {
    case "hello":
        var v8 uint32 = v2
        _ = v8
    case "\n\t":
        switch "\n\t" {
        case "hello":
            v9 := GenS1{F0: 2, F1: 0, F2: v3}
            _ = v9
        case "\n\t":
            v10 := GenS1{F0: (v6 + v6), F1: 686, F2: (v3 & v3)}
            _ = v10
        case "foo bar":
            for fi11 := 0; fi11 < 15; fi11++ {
                if (uint8(v6) == (v1 / (((v1 & (^(v1)))) | 1))) {
                    var v12 int8 = 1
                    switch fi11 {
                    case 32768:
                        v13 := []int{-1, fi11, fi11}
                        _ = v13
                    case 896:
                        v14 := v5["\n\t"]
                        _ = v14
                    case -702:
                        if (!true) {
                            var v15 float32 = 1e10
                            _ = v15
                        } else {
                            v16 := []int{(fi11 | int(v12)), -126}
                            for fi17 := 0; fi17 < 17; fi17++ {
                                v18 := GenS1{F0: (v6 * v6), F1: 1000, F2: (^(v3))}
                                v2 = (v2 + v2)
                                var v19 int8 = (v12 + int8(v6))
                                _ = fi17
                                _ = v18
                                _ = v19
                            }
                            var v20 uint8 = uint8(v2)
                            _ = v16
                            _ = v20
                        }
                    }
                    if false {
                        var v21 int16 = 227
                        var v22 int16 = v21
                        var v23 uint8 = v1
                        _ = v21
                        _ = v22
                        _ = v23
                    }
                    _ = v12
                } else {
                    if true {
                        v4["hello"] = (v2 % ((v2) | 1))
                    }
                    var v24 string = "a"
                    _ = v24
                }
                var v25 int16 = -32768
                v26 := GenS1{F0: v6, F1: (^(v25)), F2: v3}
                _ = fi11
                _ = v25
                _ = v26
            }
        }
    }
    for fi27 := 0; fi27 < 14; fi27++ {
        if true {
            var v28 int32 = -2147483648
            _ = v28
        } else {
            fmt.Println(v1, fi27, fi27)
            var v29 uint = v6
            for fi30 := 0; fi30 < 13; fi30++ {
                fmt.Println(v6)
                v2 = (v2 + (v2 % (((v2 + (v2 & 0))) | 1)))
                fmt.Println(v1)
                _ = fi30
            }
            _ = v29
        }
        var v31 uint8 = v1
        v32 := v4[(("hello" + "\n\t") + "\n\t")]
        _ = fi27
        _ = v31
        _ = v32
    }
    var v33 int16 = -401
    v33 += 250
    v34 := GenS1{F0: (^(v6)), F1: v33, F2: v3}
    v35 := []int{-119, 464, -9223372036854775808}
    switch (("a" + "foo bar") + "") {
    case "a":
        v36 := GenS1{F0: v6, F1: v33, F2: v3}
        _ = v36
    case "hello":
        if false {
            if (65535 < v6) {
                var v37 uint32 = v2
                var v38 int = 2147483647
                _ = v37
                _ = v38
            } else {
                v6 += (v6 % ((v6) | 1))
            }
            if false {
                var v39 uint32 = (v2 - 42)
                v40 := GenS2{F0: int32(v3), F1: -1}
                v1 = v1
                _ = v39
                _ = v40
            }
            var v41 int8 = 34
            _ = v41
        } else {
            if (2 != v3) {
                _ = len(v35)
                v7 = (v7 * v7)
            } else {
                if false {
                    switch 0 {
                    case 42:
                        v2 -= 65535
                    default:
                        if true {
                            fmt.Println(v6, v3, v3)
                            v1 = v1
                        } else {
                            v42 := GenS2{F0: -1, F1: 42}
                            _ = v42
                        }
                    }
                    if false {
                        fmt.Println(v33, v3)
                        v3 -= (v3 | 42)
                    } else {
                        var v43 uint = v6
                        _ = v43
                    }
                } else {
                    _ = len(v5)
                    v44 := v4["a"]
                    _ = v44
                }
                var v45 float64 = (v7 / v7)
                switch 9223372036854775807 {
                case -9223372036854775808:
                    if (-2147483648 > int32(v6)) {
                        v46 := GenS3{F0: (^(v6)), F1: 987}
                        var v47 int16 = (-(v33))
                        var v48 uint8 = (v1 & (v1 >> ((957) & 63)))
                        _ = v46
                        _ = v47
                        _ = v48
                    }
                case -1:
                    v7 += v45
                case 32767:
                    v49 := GenS2{F0: 42, F1: 326}
                    _ = v49
                default:
                    _ = len(v35)
                }
                _ = v45
            }
        }
    case "foo bar":
        var v50 uint8 = v1
        _ = v50
    default:
        if (388 != 538) {
            v6 = v6
            var v51 uint32 = uint32(v6)
            _ = len(v4)
            _ = v51
        }
    }
    var v52 int32 = -423
    switch -599 {
    case -832:
        if true {
            fmt.Println(v2, v2, v33)
            fmt.Println(v6, v2)
            var v53 uint16 = (v3 >> ((204) & 63))
            _ = v53
        } else {
            if (false || ((!(!(!((!(42 > 971)) || (v1 <= v1))))) && (v3 <= v3))) {
                for fi54 := 0; fi54 < 2; fi54++ {
                    v7 = v7
                    if true {
                        v55 := GenS1{F0: v6, F1: (v33 / (((^(v33))) | 1)), F2: v3}
                        v3 = v3
                        fmt.Println(fi54, fi54)
                        _ = v55
                    }
                    _ = fi54
                }
                delete(v4, ("\n\t" + ("\n\t" + "a")))
                var v56 string = "a"
                _ = v56
            }
            var v57 uint16 = 380
            _ = v57
        }
    default:
        v6 -= 224
    }
    if (32767 <= (v33 | (v33 & (v33 * -32768)))) {
        for fi58 := 0; fi58 < 2; fi58++ {
            if (!(!(true && false))) {
                fmt.Println(v2, fi58, v33)
            } else {
                v59 := GenS2{F0: v52, F1: -1}
                _ = v59
            }
            _ = fi58
        }
    }
    v33 = 552
    for fi60 := 0; fi60 < 15; fi60++ {
        var v61 float64 = v7
        var v62 float64 = v7
        _ = fi60
        _ = v61
        _ = v62
    }
    v63 := GenS4{F0: v3, F1: 166, F2: ((!false) && false)}
    var v64 uint16 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v52:", v52)
    fmt.Println("v6:", v6)
    fmt.Println("v64:", v64)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v52
    _ = v6
    _ = v63
    _ = v64
    _ = v7
}
