package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int; F2 uint16 }
type GenS2 struct { F0 string; F1 bool; F2 int32 }
type GenS3 struct { F0 uint; F1 uint; F2 uint64 }
type GenS4 struct { F0 uint64; F1 int16; F2 int }
type GenS5 struct { F0 uint64; F1 int64; F2 int8 }
type GenS6 struct { F0 uint16; F1 int8 }
func genFn0(p0 int32, p1 string, p2 int16) int32 {
    var v7 uint8 = 19
    p2 += (p2 * (p2 * p2))
    var v8 int8 = -56
    v9 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return p0
}
func genFn1() int64 {
    if false {
        var v10 string = "foo bar"
        _ = v10
    } else {
        var v11 float32 = 3.14159
        var v12 uint8 = 166
        var v13 int = 361
        _ = v11
        _ = v12
        _ = v13
    }
    var v14 uint = 42
    v15 := []int{-9223372036854775808, 127, 1, -2147483648}
    fmt.Println(v14)
    _ = v14
    _ = v15
    return -9223372036854775808
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int32 = 909
    v3 := []uint32{42, 276, 4294967295, 259}
    v4 := make(map[string]string)
    var v5 float32 = 0.1
    v6 := []uint32{}
    if (!true) {
        v16 := GenS1{F0: 32767, F1: -784, F2: 134}
        var v17 int8 = -3
        for fi18 := 0; fi18 < 6; fi18++ {
            var v19 int16 = 42
            var v20 uint64 = 140
            v19 = -197
            _ = fi18
            _ = v19
            _ = v20
        }
        _ = v16
        _ = v17
    }
    var v21 uint32 = 535
    var v22 uint8 = 127
    var v23 int = 0
    switch (v23 ^ (v23 ^ (v23 + -1))) {
    case -9223372036854775808:
        var v24 uint64 = 586
        _ = v24
    case 810:
        switch (("foo bar" + "a") + "\n\t") {
        case "a":
            if true {
                v2 += v2
                var v25 uint8 = (^(v22))
                if false {
                    v26 := GenS2{F0: "hello", F1: (true && (!((("hello" + "hello") + ("a" + (("\n\t" + "a") + "foo bar"))) > ""))), F2: (v2 >> ((255) & 63))}
                    v27 := GenS3{F0: 18446744073709551615, F1: 2, F2: 884}
                    var v28 string = ("a" + ("" + "\n\t"))
                    _ = v26
                    _ = v27
                    _ = v28
                } else {
                    var v29 string = ("a" + ((("hello" + "") + "") + ""))
                    var v30 float64 = -2.25
                    if (v21 != uint32(v25)) {
                        v22 = 2
                    }
                    _ = v29
                    _ = v30
                }
                _ = v25
            } else {
                var v31 uint16 = 42
                _ = v31
            }
        case "":
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        default:
            switch (^(v23)) {
            case 1:
                var v32 uint32 = 2147483648
                _ = v32
            case 0:
                var v33 bool = true
                _ = v33
            default:
                v22 -= 2
            }
        }
    case -476:
        var v34 uint64 = 212
        _ = v34
    default:
        if (!(!(!false))) {
            var v35 int64 = 208
            for fi36 := 0; fi36 < 12; fi36++ {
                _ = len(v1)
                switch v23 {
                case -108:
                    if (((^(v22)) != v22) || false) {
                        var v37 uint = 567
                        v23 += (fi36 * (-(fi36)))
                        _ = v37
                    }
                case -9223372036854775808:
                    var v38 bool = true
                    _ = v38
                case 0:
                    var v39 float64 = -2.25
                    _ = v39
                }
                var v40 string = ("" + "")
                _ = fi36
                _ = v40
            }
            var v41 int = v23
            _ = v35
            _ = v41
        } else {
            v4["\n\t"] = "\n\t"
            v23 = v23
        }
    }
    var v42 float32 = (v5 / v5)
    v2 = 99
    v43 := GenS4{F0: 2, F1: -1, F2: 128}
    v23 -= (v23 / ((-2147483648) | 1))
    fmt.Println(v42, v21)
    if (!true) {
        for fi44 := 0; fi44 < 15; fi44++ {
            var v45 uint16 = 0
            if false {
                fmt.Println(fi44, fi44)
                switch "\n\t" {
                case "hello":
                    if true {
                        switch "" {
                        case "foo bar":
                            for fi46 := 0; fi46 < 8; fi46++ {
                                if (!true) {
                                    switch (("hello" + ("a" + "\n\t")) + "hello") {
                                    case "":
                                        v2 = (v2 / ((v2) | 1))
                                        var v47 int32 = genFn0((v2 / ((428) | 1)), ("\n\t" + "a"), 1000)
                                        v2 = (v2 ^ 721)
                                        v48 := GenS3{F0: 4294967295, F1: 598, F2: 0}
                                        v6 = append(v6, v21)
                                        _ = v47
                                        _ = v48
                                    default:
                                        if true {
                                            var v49 int64 = -1
                                            delete(v4, ("a" + ""))
                                            _ = v49
                                        }
                                    }
                                    v50 := GenS1{F0: -1, F1: v23, F2: v45}
                                    _ = v50
                                }
                                _ = fi46
                            }
                        case "":
                            v21 = v21
                        case "\n\t":
                            if false {
                                fmt.Println(v21)
                                v51 := v1[(v21 * v21)]
                                _ = v51
                            }
                        }
                        var v52 uint8 = (^(v22))
                        v5 -= (1.5 - float32(-2.25))
                        _ = v52
                    }
                case "":
                    for fi53 := 0; fi53 < 14; fi53++ {
                        fmt.Println(v5, v22)
                        v54 := GenS5{F0: 42, F1: -1, F2: -128}
                        _ = fi53
                        _ = v54
                    }
                }
                var v55 bool = true
                _ = v55
            } else {
                var v56 string = "\n\t"
                _ = v56
            }
            v4[("" + "hello")] = ("" + "")
            _ = fi44
            _ = v45
        }
        if false {
            var v57 int64 = 9223372036854775807
            v58 := GenS6{F0: 255, F1: 43}
            switch "\n\t" {
            case "a":
                var v59 uint64 = 18446744073709551615
                _ = v59
            case "":
                if (0 < 18) {
                    var v60 uint16 = 2
                    _ = v60
                } else {
                    var v61 int32 = (v2 % (((v2 >> ((992) & 63))) | 1))
                    delete(v4, (("hello" + ("hello" + ("foo bar" + ""))) + ("\n\t" + "")))
                    fmt.Println(v22)
                    _ = v61
                }
            case "\n\t":
                if (!false) {
                    fmt.Println(v21, v42)
                    for fi62 := 0; fi62 < 1; fi62++ {
                        v63 := GenS2{F0: "foo bar", F1: true, F2: (v2 ^ v2)}
                        v64 := v1[v21]
                        v2 -= (v2 | v2)
                        _ = fi62
                        _ = v63
                        _ = v64
                    }
                } else {
                    v4["foo bar"] = "hello"
                }
            }
            _ = v57
            _ = v58
        }
        v3 = append(v3, v21)
    }
    v22 -= (v22 % ((v22) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v5
    _ = v6
}
