package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint8; F1 int; F2 uint32 }
type GenS3 struct { F0 string; F1 uint32 }
type GenS4 struct { F0 string; F1 bool }
type GenS5 struct { F0 int32; F1 bool; F2 int64 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 int32; F1 int32 }
func genFn0() bool {
    if (false || (!true)) {
        var v7 bool = true
        var v8 int32 = -2147483648
        v9 := GenS1{F0: 42}
        _ = v7
        _ = v8
        _ = v9
    }
    if false {
        var v10 int = -433
        if false {
            switch "hello" {
            case "a":
                var v11 int64 = 417
                _ = v11
            default:
                fmt.Println(v10)
            }
            var v12 uint8 = 193
            _ = v12
        } else {
            var v13 uint = 2
            v14 := []int{4294967295, (v10 * (v10 >> ((v13) & 63))), v10}
            switch v10 {
            case 1:
                switch (v10 * v10) {
                case 448:
                    var v15 string = ""
                    _ = v15
                case 601:
                    if ("hello" <= ("\n\t" + "hello")) {
                        if (true || false) {
                            if false {
                                v16 := GenS2{F0: 97, F1: v10, F2: 187}
                                v10 += v10
                                fmt.Println(v10, v13)
                                _ = v16
                            }
                            switch v10 {
                            case -2147483648:
                                v17 := GenS1{F0: 2}
                                _ = v17
                            case -173:
                                if (255 == 1) {
                                    v10 = (v10 * (v10 | (^(v10))))
                                }
                            case 166:
                                _ = len(v14)
                            default:
                                switch (v10 << ((18446744073709551615) & 63)) {
                                case 32767:
                                    var v18 int16 = -32768
                                    _ = v18
                                }
                            }
                            v14 = append(v14, (v10 - (v10 >> ((v13) & 63))))
                        }
                        var v19 uint8 = 255
                        v14 = append(v14, 127)
                        _ = v19
                    }
                case 0:
                    if (true && false) {
                        if len(v14) > 0 {
                            v14[0] = v10
                        }
                        v10 = v10
                        var v20 uint = 4294967295
                        _ = v20
                    } else {
                        v10 -= 4294967295
                        var v21 uint64 = 156
                        _ = v21
                    }
                default:
                    for fi22 := 0; fi22 < 17; fi22++ {
                        v13 -= 2
                        v10 = (fi22 * -358)
                        _ = fi22
                    }
                }
            case 128:
                v13 = 1
            case -210:
                if len(v14) >= 2 {
                    v14 = v14[1:len(v14)]
                }
            }
            _ = v13
            _ = v14
        }
        var v23 uint = 42
        _ = v10
        _ = v23
    }
    var v24 int32 = 0
    var v25 uint32 = 0
    _ = v24
    _ = v25
    return ((!true) || (true && false))
}
func genFn1() bool {
    v26 := GenS1{F0: 1}
    if (100 >= 0) {
        var v27 uint16 = 42
        v27 += 65535
        v28 := GenS3{F0: "foo bar", F1: 2}
        _ = v27
        _ = v28
    } else {
        if false {
            for fi29 := 0; fi29 < 17; fi29++ {
                fmt.Println(fi29)
                v30 := GenS2{F0: 128, F1: (fi29 - (fi29 | fi29)), F2: 4294967295}
                _ = fi29
                _ = v30
            }
            var v31 float64 = (((float64(0.0) / -2.25) / 1.0) / (0.1 - 1e10))
            _ = v31
        } else {
            var v32 int32 = -714
            _ = v32
        }
        var v33 float64 = 0.1
        _ = v33
    }
    var v34 int16 = -210
    if (!((-632 == 9223372036854775807) || ((true || (true || (true && (1 <= 0)))) && (true && true)))) {
        fmt.Println(v34)
    }
    v34 -= v34
    _ = v26
    _ = v34
    return ((128 <= 1) && (42 >= 587))
}

func main() {
    var v1 int32 = 1
    v2 := make(map[string]bool)
    var v3 int = 32768
    v4 := make(map[uint32]string)
    var v5 int32 = v1
    var v6 uint32 = 2147483648
    v35 := []int{(v3 - v3), (v3 - (v3 ^ v3)), v3, v3}
    switch "\n\t" {
    case "\n\t":
        if (!(0 < 9223372036854775807)) {
            fmt.Println(v6, v6)
            var v36 int8 = 0
            _ = v36
        }
    case "":
        var v37 uint8 = 248
        _ = v37
    case "foo bar":
        for fi38 := 0; fi38 < 13; fi38++ {
            v39 := GenS4{F0: ("\n\t" + ("a" + "hello")), F1: false}
            var v40 float32 = (1e10 / 0.1)
            _ = fi38
            _ = v39
            _ = v40
        }
    }
    v41 := GenS5{F0: (v5 & v5), F1: ((uint(v6) >= 299) || (32 > 2)), F2: -779}
    var v42 int16 = 203
    for fi43 := 0; fi43 < 7; fi43++ {
        var v44 float32 = 0.1
        _ = fi43
        _ = v44
    }
    v1 += (v5 / (((v5 << ((18) & 63))) | 1))
    v45 := GenS6{F0: uint(v6)}
    v5 -= v5
    v42 += v42
    v46 := GenS7{F0: v5, F1: -469}
    var v47 uint8 = 2
    v3 -= -1
    v48 := GenS7{F0: -1, F1: -2147483648}
    for fi49 := 0; fi49 < 18; fi49++ {
        switch ("foo bar" + "foo bar") {
        case "":
            if (false && (!(!true))) {
                v35 = append(v35, fi49)
                switch ("foo bar" + "\n\t") {
                case "foo bar":
                    var v50 uint = 4294967295
                    _ = v50
                case "":
                    for fi51 := 0; fi51 < 2; fi51++ {
                        v35 = append(v35, (fi49 << ((146) & 63)))
                        _ = fi51
                    }
                }
                var v52 int = (-(v3))
                _ = v52
            }
        default:
            fmt.Println(v1, v5, v42)
        }
        v42 = v42
        _ = fi49
    }
    if true {
        v5 += (v5 & (v1 | v1))
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v41
    _ = v42
    _ = v45
    _ = v46
    _ = v47
    _ = v48
    _ = v5
    _ = v6
}
