package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string }
type GenS2 struct { F0 string }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint32 }

func main() {
    var v1 uint = 255
    var v2 uint8 = 211
    var v3 int16 = 32767
    fmt.Println(v1, v2)
    v4 := GenS1{F0: 0, F1: "\n\t"}
    var v5 uint16 = 293
    v6 := GenS1{F0: 2147483647, F1: "foo bar"}
    for fi7 := 0; fi7 < 9; fi7++ {
        var v8 int64 = -9223372036854775808
        v9 := []int{(fi7 >> (((v1 >> ((v1) & 63))) & 63)), fi7}
        _ = fi7
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 20; fi10++ {
        for fi11 := 0; fi11 < 3; fi11++ {
            switch fi11 {
            case 42:
                var v12 uint = v1
                _ = v12
            case 924:
                v2 = 127
            case 32767:
                v13 := GenS1{F0: 0, F1: "foo bar"}
                v2 = v2
                _ = v13
            default:
                switch ("foo bar" + "") {
                case "foo bar":
                    var v14 uint32 = 599
                    _ = v14
                case "":
                    v3 = 42
                case "a":
                    var v15 uint16 = 2
                    _ = v15
                }
            }
            if ((-2.25 / 0.1) != (3.14159 / 1e10)) {
                switch int(v1) {
                case 127:
                    v16 := []int{}
                    _ = v16
                case 145:
                    if true {
                        var v17 int32 = 543
                        if (true || (-1 > 12)) {
                            switch fi10 {
                            case 128:
                                if (!true) {
                                    var v18 uint16 = (v5 * (v5 ^ (v5 & v5)))
                                    _ = v18
                                } else {
                                    fmt.Println(v5, v5, v3)
                                    v19 := GenS1{F0: v17, F1: ""}
                                    _ = v19
                                }
                            case -51:
                                v1 = (v1 ^ 65535)
                            case 4294967295:
                                switch "\n\t" {
                                case "foo bar":
                                    if true {
                                        var v20 uint32 = 65535
                                        _ = v20
                                    } else {
                                        v5 -= v5
                                        v5 = (v5 + 687)
                                        if true {
                                            if (false && (-23 > 127)) {
                                                v21 := []int{fi10}
                                                _ = v21
                                            } else {
                                                v1 = (v1 - (^(v1)))
                                                var v22 int = fi11
                                                if (-2.25 > (float64(0.0) - 0.1)) {
                                                    if false {
                                                        v23 := []int{(v22 % (((fi10 % ((623) | 1))) | 1)), 917, (fi11 | (fi11 | (v22 & (v22 & v22)))), (fi11 - v22)}
                                                        _ = v23
                                                    } else {
                                                        if (42 > -922) {
                                                            v17 -= (^(v17))
                                                            switch (fi11 ^ fi11) {
                                                            case -289:
                                                                switch "hello" {
                                                                case "\n\t":
                                                                    var v24 int32 = int32(v5)
                                                                    _ = v24
                                                                case "foo bar":
                                                                    v17 += (v17 | (v17 ^ int32(v3)))
                                                                default:
                                                                    v25 := []int{-593, fi10, (fi10 << ((v1) & 63)), fi10}
                                                                    _ = v25
                                                                }
                                                            default:
                                                                v26 := GenS1{F0: -1, F1: ("a" + ("hello" + ("hello" + "\n\t")))}
                                                                _ = v26
                                                            }
                                                        }
                                                        v27 := []int{v22, fi10, fi11, v22}
                                                        _ = v27
                                                    }
                                                    v28 := GenS1{F0: 42, F1: "\n\t"}
                                                    v29 := GenS2{F0: ("\n\t" + "hello")}
                                                    v3 = (v3 % ((32767) | 1))
                                                    _ = v28
                                                    _ = v29
                                                }
                                                _ = v22
                                            }
                                        } else {
                                            var v30 bool = true
                                            v31 := []int{}
                                            fmt.Println(fi10)
                                            _ = v30
                                            _ = v31
                                        }
                                    }
                                default:
                                    if true {
                                        switch 32767 {
                                        case 9223372036854775807:
                                            v32 := []int{fi10, fi10, fi11, -877}
                                            _ = v32
                                        case 498:
                                            if ((v5 / ((v5) | 1)) <= v5) {
                                                var v33 int = (fi10 ^ 127)
                                                var v34 float32 = ((-2.25 + (1.5 - -2.25)) * (1e10 / ((-2.25 * 1.5) * 0.1)))
                                                v35 := []int{(fi10 >> ((v1) & 63)), (v33 & v33)}
                                                var v36 uint = (v1 >> ((v1) & 63))
                                                _ = len(v35)
                                                _ = v33
                                                _ = v34
                                                _ = v35
                                                _ = v36
                                            }
                                        case 0:
                                            v37 := []int{}
                                            _ = v37
                                        default:
                                            fmt.Println(v17, v3)
                                        }
                                        fmt.Println(fi10, v5)
                                        v38 := []int{}
                                        _ = v38
                                    } else {
                                        v39 := []int{fi11, fi10, (fi11 % (((fi10 % (((-(fi11))) | 1))) | 1)), fi11}
                                        switch ("a" + "foo bar") {
                                        case "hello":
                                            v1 = v1
                                        case "a":
                                            if (((65535 != (v1 + 838)) && (!((!((v2 & v2) >= (v2 / (((v2 & v2)) | 1)))) || (false || ((1.5 <= 0.0) || (!true)))))) && true) {
                                                var v40 int16 = v3
                                                v41 := GenS2{F0: (("" + "a") + "hello")}
                                                _ = v40
                                                _ = v41
                                            }
                                            if true {
                                                v42 := GenS3{F0: (int64(v2) != 42)}
                                                v1 -= v1
                                                _ = v42
                                            }
                                            v1 += 396
                                            var v43 bool = false
                                            v44 := GenS4{F0: v1}
                                            _ = v43
                                            _ = v44
                                        case "foo bar":
                                            v2 = v2
                                        default:
                                            v45 := GenS1{F0: v17, F1: "foo bar"}
                                            _ = v45
                                        }
                                        _ = v39
                                    }
                                }
                            }
                        } else {
                            v5 -= (v5 - v5)
                            v17 = v17
                            var v46 uint64 = 4294967295
                            _ = v46
                        }
                        var v47 int = (-(fi11))
                        _ = v17
                        _ = v47
                    } else {
                        if true {
                            fmt.Println(v1)
                        } else {
                            v2 = v2
                        }
                        v48 := GenS4{F0: v1}
                        _ = v48
                    }
                case 605:
                    switch "a" {
                    case "foo bar":
                        fmt.Println(v1, fi11, fi11)
                        var v49 uint64 = 18446744073709551615
                        v3 = v3
                        _ = v49
                    case "a":
                        var v50 int8 = 127
                        _ = v50
                    }
                default:
                    switch -221 {
                    case 42:
                        v3 += (v3 & v3)
                    case 9223372036854775807:
                        v51 := GenS5{F0: 534}
                        _ = v51
                    default:
                        fmt.Println(fi11, v1)
                    }
                }
            }
            _ = fi11
        }
        _ = fi10
    }
    v3 = v3
    if (!false) {
        v52 := GenS3{F0: (false || true)}
        _ = v52
    } else {
        v53 := GenS4{F0: (^(v1))}
        v54 := []int{42, 42, 0, -2147483648}
        for fi55 := 0; fi55 < 18; fi55++ {
            for fi56 := 0; fi56 < 13; fi56++ {
                if ((((!false) || true) && (!false)) && (1.5 < 3.14159)) {
                    var v57 int = fi55
                    v3 = -383
                    _ = v57
                } else {
                    if len(v54) >= 2 {
                        v54 = v54[1:len(v54)]
                    }
                    if true {
                        if (!false) {
                            v3 = (v3 & v3)
                            var v58 uint8 = v2
                            _ = v58
                        } else {
                            if (v3 < (v3 | (-(v3)))) {
                                var v59 uint32 = uint32(v1)
                                _ = v59
                            }
                            if true {
                                fmt.Println(fi56, v2)
                                v1 = 2
                                if len(v54) >= 2 {
                                    v54 = v54[1:len(v54)]
                                }
                            }
                            v2 += 128
                        }
                        var v60 int = fi56
                        _ = v60
                    } else {
                        var v61 float64 = (float64((1.5 * 1e10)) * (float64(-2.25) - 1.5))
                        if true {
                            var v62 int = fi55
                            _ = v62
                        }
                        _ = v61
                    }
                }
                _ = fi56
            }
            _ = fi55
        }
        _ = v53
        _ = v54
    }
    for fi63 := 0; fi63 < 9; fi63++ {
        if false {
            var v64 string = (("\n\t" + "\n\t") + ("hello" + "\n\t"))
            v65 := GenS2{F0: (((("a" + ("a" + v64)) + (("foo bar" + "") + "\n\t")) + (v64 + v64)) + "foo bar")}
            _ = v64
            _ = v65
        }
        _ = fi63
    }
    var v66 int16 = (v3 | v3)
    v67 := GenS1{F0: 0, F1: "a"}
    v5 -= v5
    v3 += (v3 ^ (v3 & v3))
    for fi68 := 0; fi68 < 12; fi68++ {
        for fi69 := 0; fi69 < 15; fi69++ {
            v3 = v3
            _ = fi69
        }
        fmt.Println(v5)
        v2 += v2
        _ = fi68
    }
    v70 := GenS6{F0: 1}
    var v71 float64 = float64(1e10)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v66:", v66)
    fmt.Println("v71:", v71)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v66
    _ = v67
    _ = v70
    _ = v71
}
