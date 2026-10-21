package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8; F2 string }
type GenS2 struct { F0 uint8; F1 uint }
type GenS3 struct { F0 string; F1 uint8; F2 int8 }
type GenS4 struct { F0 int8; F1 int; F2 int32 }
type GenS5 struct { F0 uint64 }
func genFn0(p0 uint32, p1 string, p2 int8) int32 {
    p0 = p0
    var v5 string = "hello"
    if (!(false && true)) {
        for fi6 := 0; fi6 < 19; fi6++ {
            var v7 bool = false
            var v8 float32 = (float32(1.5) / 1.0)
            var v9 int16 = -1
            _ = fi6
            _ = v7
            _ = v8
            _ = v9
        }
    }
    if false {
        var v10 int32 = 42
        var v11 uint = 65535
        _ = v10
        _ = v11
    } else {
        v12 := []int{}
        v13 := GenS1{F0: 2147483647, F1: 42, F2: ""}
        v12 = append(v12, 114)
        _ = v12
        _ = v13
    }
    fmt.Println(p2, v5)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 716
}
func genFn1(p0 int, p1 int16) uint {
    p0 = (-(p0))
    var v14 int16 = (p1 & (p1 + -984))
    for fi15 := 0; fi15 < 15; fi15++ {
        var v16 uint32 = 4294967295
        if (-9223372036854775808 < -1) {
            switch "" {
            case "hello":
                for fi17 := 0; fi17 < 16; fi17++ {
                    var v18 float64 = 1e10
                    var v19 float64 = (v18 - 0.1)
                    _ = fi17
                    _ = v18
                    _ = v19
                }
            case "a":
                if false {
                    for fi20 := 0; fi20 < 17; fi20++ {
                        fmt.Println(v16, p1)
                        v21 := GenS1{F0: genFn0((v16 % ((42) | 1)), ("hello" + "hello"), -1), F1: 2, F2: "foo bar"}
                        fmt.Println(v16)
                        _ = fi20
                        _ = v21
                    }
                    for fi22 := 0; fi22 < 10; fi22++ {
                        var v23 bool = false
                        p0 = (p0 << ((42) & 63))
                        v16 += (^(v16))
                        _ = fi22
                        _ = v23
                    }
                } else {
                    if true {
                        for fi24 := 0; fi24 < 10; fi24++ {
                            p0 -= (fi24 | -286)
                            v25 := []int{1, 0, p0}
                            fmt.Println(v14, p0)
                            var v26 uint32 = v16
                            _ = fi24
                            _ = v25
                            _ = v26
                        }
                    }
                }
            case "foo bar":
                var v27 bool = (true || true)
                _ = v27
            default:
                if (p1 <= v14) {
                    for fi28 := 0; fi28 < 20; fi28++ {
                        if false {
                            var v29 uint = 629
                            v14 -= (v14 - 1)
                            switch "a" {
                            case "":
                                v29 = 1
                            case "\n\t":
                                v30 := GenS2{F0: 2, F1: 18446744073709551615}
                                v31 := []int{int(v29)}
                                v32 := GenS2{F0: 128, F1: (v29 >> (((^(v29))) & 63))}
                                _ = v30
                                _ = v31
                                _ = v32
                            default:
                                v29 = uint(v16)
                            }
                            v33 := GenS3{F0: "\n\t", F1: 127, F2: 119}
                            fmt.Println(v14)
                            _ = v29
                            _ = v33
                        } else {
                            v34 := []int{(fi15 & 40), p0}
                            v35 := GenS2{F0: 128, F1: 214}
                            _ = v34
                            _ = v35
                        }
                        v16 = 42
                        _ = fi28
                    }
                }
            }
            v36 := GenS4{F0: 100, F1: (p0 + 991), F2: -1}
            var v37 bool = ((!false) || (!false))
            _ = v36
            _ = v37
        }
        _ = fi15
        _ = v16
    }
    _ = p0
    _ = p1
    _ = v14
    return 922
}

func main() {
    var v1 uint16 = 2
    var v2 uint32 = 0
    v3 := []int{-1}
    var v4 int32 = 2147483647
    if len(v3) > 0 {
        v3[0] = 32768
    }
    var v38 int32 = genFn0((v2 | 0), "a", 100)
    v39 := GenS1{F0: v38, F1: 35, F2: ""}
    for fi40 := 0; fi40 < 2; fi40++ {
        switch "foo bar" {
        case "foo bar":
            if false {
                var v41 string = "foo bar"
                switch "" {
                case "foo bar":
                    v41 = "\n\t"
                default:
                    v42 := GenS5{F0: uint64(v2)}
                    _ = v42
                }
                _ = v41
            }
        default:
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
        var v43 float32 = -2.25
        _ = fi40
        _ = v43
    }
    if (((0.0 / -2.25) + float32(1.5)) == (1.5 - (3.14159 / float32(1.5)))) {
        if false {
            var v44 uint = 2
            var v45 bool = (true && (1e10 == ((1.5 / 3.14159) * 0.0)))
            _ = v44
            _ = v45
        } else {
            v4 = 1
        }
        for fi46 := 0; fi46 < 7; fi46++ {
            var v47 int32 = (v4 | v4)
            _ = fi46
            _ = v47
        }
    }
    var v48 int64 = 9223372036854775807
    v49 := GenS5{F0: 18446744073709551615}
    var v50 uint8 = 216
    switch -1 {
    case 4294967295:
        var v51 int32 = (v38 * 2147483647)
        _ = v51
    }
    v1 += v1
    var v52 uint64 = 18446744073709551615
    fmt.Println(v2, v1, v4)
    v53 := GenS5{F0: v52}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("v50:", v50)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v48
    _ = v49
    _ = v50
    _ = v52
    _ = v53
}
