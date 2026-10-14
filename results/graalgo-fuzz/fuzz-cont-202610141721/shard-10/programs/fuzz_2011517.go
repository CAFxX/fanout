package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64 }
type GenS2 struct { F0 int64; F1 uint32; F2 string }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint8; F1 bool; F2 bool }
func genFn0() string {
    if true {
        var v7 float64 = 0.0
        var v8 int = 127
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 3; fi9++ {
        var v10 uint64 = 175
        _ = fi9
        _ = v10
    }
    for fi11 := 0; fi11 < 7; fi11++ {
        var v12 float32 = 0.1
        v13 := GenS1{F0: 1, F1: 0}
        v14 := []int{fi11}
        _ = fi11
        _ = v12
        _ = v13
        _ = v14
    }
    for fi15 := 0; fi15 < 4; fi15++ {
        if (772 == 32767) {
            v16 := []int{fi15, (^(fi15))}
            for fi17 := 0; fi17 < 15; fi17++ {
                var v18 float64 = (1.5 - ((0.0 * (-2.25 * 0.1)) + 1.5))
                _ = fi17
                _ = v18
            }
            _ = v16
        } else {
            for fi19 := 0; fi19 < 19; fi19++ {
                v20 := GenS1{F0: 32767, F1: 1}
                v21 := []int{}
                v22 := GenS2{F0: -935, F1: 2, F2: ""}
                if (1.5 <= 1.5) {
                    fmt.Println(fi15, fi15)
                    var v23 int16 = 938
                    var v24 int16 = v23
                    _ = v23
                    _ = v24
                } else {
                    var v25 int8 = 1
                    _ = v25
                }
                _ = fi19
                _ = v20
                _ = v21
                _ = v22
            }
            var v26 uint = 65535
            if (0 < 0) {
                fmt.Println(fi15)
            }
            _ = v26
        }
        fmt.Println(fi15)
        v27 := []int{32767, 587}
        _ = fi15
        _ = v27
    }
    var v28 bool = (32767 <= -2147483648)
    _ = v28
    return "foo bar"
}

func main() {
    var v1 uint = 812
    var v2 uint8 = 179
    var v3 int8 = 42
    var v4 int64 = -1
    var v5 uint16 = 32767
    v6 := []int{-775, -9223372036854775808, -901, 128}
    v3 = v3
    var v29 float64 = 0.1
    v3 -= (v3 + (v3 + int8(v5)))
    fmt.Println(v5, v2)
    for fi30 := 0; fi30 < 10; fi30++ {
        v29 -= v29
        if true {
            switch "a" {
            case "\n\t":
                var v31 int = fi30
                _ = v31
            default:
                var v32 int16 = 1000
                _ = v32
            }
            for fi33 := 0; fi33 < 2; fi33++ {
                var v34 bool = (646 < 195)
                v4 -= 42
                _ = fi33
                _ = v34
            }
        }
        _ = fi30
    }
    v35 := GenS2{F0: v4, F1: 2147483648, F2: "\n\t"}
    v6 = append(v6, -1)
    v4 = (v4 * (v4 % (((v4 & -958)) | 1)))
    var v36 int16 = -32768
    v3 -= (v3 % ((1) | 1))
    var v37 uint16 = 383
    if false {
        var v38 int8 = -15
        if false {
            switch "\n\t" {
            case "foo bar":
                v1 += 255
            case "a":
                switch int(v37) {
                case 4294967295:
                    var v39 int8 = (v3 / (((v38 & 42)) | 1))
                    _ = v39
                case 1:
                    switch -2147483648 {
                    case 300:
                        if ((v1 ^ (v1 + v1)) == v1) {
                            if ((!(!true)) && true) {
                                var v40 uint = (v1 - (v1 << (((^(v1))) & 63)))
                                var v41 int16 = (v36 << (((v1 - (^(v1)))) & 63))
                                for fi42 := 0; fi42 < 5; fi42++ {
                                    v43 := GenS1{F0: (^(v37)), F1: (-(v4))}
                                    if true {
                                        if ((false && (false || false)) && (!(((3.14159 * float32((1e10 + 0.0))) * 0.1) == 3.14159))) {
                                            v41 = v41
                                            _ = len(v6)
                                            switch ("a" + "a") {
                                            case "\n\t":
                                                v44 := GenS2{F0: v4, F1: 0, F2: "\n\t"}
                                                _ = v44
                                            case "foo bar":
                                                v3 -= v3
                                            case "":
                                                var v45 uint = (^(v1))
                                                _ = v45
                                            }
                                        }
                                        if ((-(v3)) < v38) {
                                            v46 := GenS2{F0: (v4 << ((2) & 63)), F1: 2147483648, F2: "hello"}
                                            for fi47 := 0; fi47 < 9; fi47++ {
                                                fmt.Println(v1, v40, v1)
                                                _ = fi47
                                            }
                                            var v48 uint16 = (v5 - 2)
                                            _ = v46
                                            _ = v48
                                        } else {
                                            v49 := GenS3{F0: v4}
                                            _ = v49
                                        }
                                        v50 := GenS2{F0: (v4 & 9223372036854775807), F1: 65535, F2: "hello"}
                                        _ = v50
                                    } else {
                                        if len(v6) > 0 {
                                            v6[0] = (fi42 >> (((v40 + v1)) & 63))
                                        }
                                    }
                                    _ = fi42
                                    _ = v43
                                }
                                _ = v40
                                _ = v41
                            }
                        }
                    default:
                        v51 := GenS4{F0: v2, F1: true, F2: true}
                        _ = v51
                    }
                default:
                    var v52 int = 704
                    _ = v52
                }
            default:
                var v53 bool = true
                _ = v53
            }
            switch "\n\t" {
            case "":
                for fi54 := 0; fi54 < 15; fi54++ {
                    if len(v6) > 0 {
                        v6[0] = fi54
                    }
                    switch "hello" {
                    case "":
                        v36 = (v36 + (v36 + (v36 % ((v36) | 1))))
                    case "foo bar":
                        var v55 uint16 = 21
                        _ = v55
                    }
                    var v56 int64 = v4
                    _ = fi54
                    _ = v56
                }
            case "hello":
                if true {
                    v57 := GenS4{F0: v2, F1: (!false), F2: ((15 != 42) && true)}
                    if len(v6) > 0 {
                        v6[0] = 0
                    }
                    var v58 int = 398
                    _ = v57
                    _ = v58
                }
            default:
                v5 -= (v37 & (v5 * v5))
            }
        }
        _ = v38
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
