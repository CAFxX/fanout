package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 uint }
type GenS2 struct { F0 int; F1 string; F2 int32 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int32; F1 bool; F2 bool }
func genFn0() uint64 {
    var v6 string = "foo bar"
    v6 = v6
    _ = v6
    return 861
}
func genFn1() int8 {
    switch ("" + ("" + "\n\t")) {
    case "foo bar":
        var v7 uint32 = 88
        _ = v7
    case "\n\t":
        v8 := []int{32768}
        _ = v8
    default:
        var v9 int = 32767
        _ = v9
    }
    var v10 float64 = (1e10 * 3.14159)
    _ = v10
    return 1
}

func main() {
    var v1 uint16 = 1
    var v2 float32 = 0.1
    var v3 int32 = 1
    v4 := []uint32{}
    v5 := []int{-9223372036854775808, -2147483648}
    fmt.Println(v1)
    v1 -= 978
    var v11 float64 = 1.5
    v12 := GenS1{F0: 222, F1: 0, F2: 255}
    if true {
        fmt.Println(v1)
        switch 32768 {
        case 378:
            v13 := GenS2{F0: -132, F1: ("" + ("a" + "")), F2: (v3 | v3)}
            _ = v13
        case 746:
            var v14 uint8 = 1
            _ = v14
        case 2147483647:
            var v15 float64 = (v11 / 1e10)
            _ = v15
        }
        v16 := GenS2{F0: -2147483648, F1: "hello", F2: v3}
        _ = v16
    }
    if (v2 == (v2 / (v2 * v2))) {
        if false {
            for fi17 := 0; fi17 < 1; fi17++ {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                if (3.14159 > v2) {
                    v11 = 0.0
                }
                var v18 int16 = -1
                _ = fi17
                _ = v18
            }
        }
    }
    v19 := GenS2{F0: -120, F1: "hello", F2: v3}
    v1 = v1
    v3 -= (^(v3))
    if ("a" == "\n\t") {
        v3 -= v3
        if (!false) {
            v20 := GenS3{F0: int16(v1)}
            _ = v20
        } else {
            switch 4294967295 {
            case 2147483647:
                var v21 float32 = (v2 - ((v2 * v2) * -2.25))
                _ = v21
            case 9223372036854775807:
                var v22 uint16 = 130
                _ = v22
            }
            for fi23 := 0; fi23 < 14; fi23++ {
                v1 -= v1
                fmt.Println(v1)
                _ = fi23
            }
        }
        v1 += (v1 - v1)
    } else {
        for fi24 := 0; fi24 < 11; fi24++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            var v25 int = 107
            var v26 uint8 = 255
            _ = fi24
            _ = v25
            _ = v26
        }
        if len(v5) > 0 {
            v5[0] = 42
        }
    }
    if (!(int16(v1) <= int16(v3))) {
        v1 = 42
        _ = len(v4)
    } else {
        if (2147483648 < 2) {
            for fi27 := 0; fi27 < 14; fi27++ {
                v5 = append(v5, (fi27 % ((32768) | 1)))
                switch fi27 {
                case -812:
                    fmt.Println(v2, v3, fi27)
                default:
                    v28 := GenS2{F0: fi27, F1: "foo bar", F2: (v3 << ((422) & 63))}
                    _ = v28
                }
                _ = fi27
            }
            if len(v4) > 0 {
                v4[0] = 0
            }
        }
        var v29 uint16 = 2
        if true {
            v5 = append(v5, int(v3))
            v11 = v11
        } else {
            var v30 uint8 = 161
            var v31 int16 = 1
            v31 = (v31 - (v31 | (v31 - 632)))
            _ = v30
            _ = v31
        }
        _ = v29
    }
    switch "" {
    case "a":
        v2 = 0.0
    case "hello":
        for fi32 := 0; fi32 < 16; fi32++ {
            var v33 float32 = 3.14159
            if false {
                switch "hello" {
                case "hello":
                    for fi34 := 0; fi34 < 13; fi34++ {
                        v1 = (v1 + (v1 % (((v1 % ((v1) | 1))) | 1)))
                        var v35 uint = 324
                        v11 -= float64(1.5)
                        _ = fi34
                        _ = v35
                    }
                case "foo bar":
                    var v36 uint16 = (v1 + 1)
                    _ = v36
                default:
                    v37 := GenS4{F0: 1, F1: false, F2: false}
                    _ = v37
                }
                for fi38 := 0; fi38 < 8; fi38++ {
                    _ = len(v5)
                    v33 += 1.5
                    if (false || ("" >= "a")) {
                        switch ("foo bar" + "\n\t") {
                        case "":
                            v1 += (^(v1))
                        }
                        var v39 float64 = ((0.1 + 3.14159) / float64(v2))
                        _ = v39
                    }
                    var v40 int64 = -828
                    _ = fi38
                    _ = v40
                }
                v11 += 0.0
            } else {
                v33 = v33
                var v41 bool = (67 <= 0)
                for fi42 := 0; fi42 < 3; fi42++ {
                    var v43 uint = 899
                    v1 = (v1 - (v1 | 1))
                    _ = fi42
                    _ = v43
                }
                _ = v41
            }
            _ = fi32
            _ = v33
        }
    case "\n\t":
        fmt.Println(v3)
    }
    for fi44 := 0; fi44 < 4; fi44++ {
        var v45 int32 = v3
        var v46 int32 = v45
        var v47 uint8 = 1
        _ = fi44
        _ = v45
        _ = v46
        _ = v47
    }
    fmt.Println(v3, v3, v2)
    var v48 uint16 = (v1 << ((1) & 63))
    v48 = 1
    var v49 uint32 = 65535
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v49
    _ = v5
}
