package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16 }
type GenS2 struct { F0 uint16; F1 int64; F2 int16 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 bool }
func genFn0(p0 uint8) uint {
    var v6 uint64 = 18446744073709551615
    v7 := []int{9223372036854775807, 9223372036854775807, -984, 808}
    switch "foo bar" {
    case "a":
        for fi8 := 0; fi8 < 6; fi8++ {
            v9 := GenS1{F0: 0, F1: 32767}
            fmt.Println(p0, v6, v6)
            if false {
                v7 = append(v7, (-(fi8)))
            }
            _ = fi8
            _ = v9
        }
    case "hello":
        if ("hello" != ("\n\t" + "foo bar")) {
            for fi10 := 0; fi10 < 8; fi10++ {
                v6 = v6
                _ = fi10
            }
            if (true || false) {
                p0 = uint8(v6)
            }
        } else {
            for fi11 := 0; fi11 < 9; fi11++ {
                v6 += v6
                if false {
                    v6 -= v6
                }
                var v12 uint = 18446744073709551615
                _ = fi11
                _ = v12
            }
            p0 = 127
        }
    case "":
        if (false || false) {
            v7 = append(v7, 838)
            for fi13 := 0; fi13 < 18; fi13++ {
                p0 -= p0
                _ = fi13
            }
            var v14 uint64 = 2
            _ = v14
        } else {
            if false {
                v15 := GenS2{F0: uint16(v6), F1: 1, F2: 42}
                if (127 > 36) {
                    switch 42 {
                    case -1:
                        for fi16 := 0; fi16 < 2; fi16++ {
                            switch ("a" + "hello") {
                            case "\n\t":
                                v17 := GenS3{F0: "hello"}
                                _ = v17
                            default:
                                v18 := GenS4{F0: (uint(v6) > 886)}
                                _ = v18
                            }
                            _ = fi16
                        }
                    case -25:
                        for fi19 := 0; fi19 < 4; fi19++ {
                            v20 := GenS3{F0: ""}
                            fmt.Println(v6, fi19)
                            _ = fi19
                            _ = v20
                        }
                    case 127:
                        v6 -= 895
                    default:
                        if (42 <= (^(v6))) {
                            for fi21 := 0; fi21 < 8; fi21++ {
                                for fi22 := 0; fi22 < 1; fi22++ {
                                    v23 := GenS2{F0: 662, F1: 637, F2: 1}
                                    _ = fi22
                                    _ = v23
                                }
                                p0 -= p0
                                v24 := GenS2{F0: 85, F1: 784, F2: 32767}
                                _ = fi21
                                _ = v24
                            }
                        } else {
                            v7 = append(v7, -2147483648)
                        }
                    }
                    var v25 uint16 = 42
                    _ = v25
                } else {
                    var v26 int16 = 42
                    _ = v26
                }
                _ = v15
            } else {
                for fi27 := 0; fi27 < 11; fi27++ {
                    var v28 uint16 = 0
                    v6 -= v6
                    _ = fi27
                    _ = v28
                }
            }
            fmt.Println(p0)
            if true {
                var v29 int64 = 261
                switch "foo bar" {
                case "\n\t":
                    switch 32768 {
                    case 2147483647:
                        switch "foo bar" {
                        case "foo bar":
                            p0 -= p0
                        }
                    case 418:
                        fmt.Println(v29, v29)
                    }
                case "a":
                    if true {
                        var v30 uint = 498
                        v30 -= v30
                        _ = v30
                    }
                default:
                    v6 -= (v6 & (v6 + v6))
                }
                _ = v29
            }
        }
    }
    _ = p0
    _ = v6
    _ = v7
    return 42
}
func genFn1(p0 int32, p1 uint) uint16 {
    var v31 float64 = (float64(1e10) * 1.5)
    v32 := []int{-2147483648, 128}
    v31 = 0.0
    v33 := GenS2{F0: 0, F1: 1, F2: 1000}
    _ = p0
    _ = p1
    _ = v31
    _ = v32
    _ = v33
    return 42
}

func main() {
    v1 := make(map[string]int)
    var v2 uint16 = 42
    var v3 int64 = 0
    var v4 float64 = 1.5
    var v5 uint8 = 0
    v5 -= (^(v5))
    v3 += (-(v3))
    v34 := GenS1{F0: 18446744073709551615, F1: (v2 ^ v2)}
    for fi35 := 0; fi35 < 9; fi35++ {
        for fi36 := 0; fi36 < 3; fi36++ {
            v5 += 255
            var v37 int = (-(fi36))
            switch fi35 {
            case 1:
                var v38 bool = (42 != 42)
                fmt.Println(v2)
                _ = v38
            default:
                v39 := v1["foo bar"]
                _ = v39
            }
            _ = fi36
            _ = v37
        }
        for fi40 := 0; fi40 < 18; fi40++ {
            if true {
                var v41 int64 = 42
                v4 = v4
                v3 = -382
                _ = v41
            } else {
                v2 -= 794
                v5 += v5
                fmt.Println(v4, fi35, fi40)
            }
            var v42 int32 = -133
            delete(v1, "\n\t")
            _ = fi40
            _ = v42
        }
        _ = fi35
    }
    v5 -= 255
    var v43 int64 = 0
    v3 += v43
    var v44 uint64 = 42
    fmt.Println(v2, v3, v3)
    v2 -= v2
    for fi45 := 0; fi45 < 3; fi45++ {
        var v46 uint32 = 2147483648
        var v47 uint = 745
        _ = fi45
        _ = v46
        _ = v47
    }
    if true {
        v1["foo bar"] = 42
        for fi48 := 0; fi48 < 7; fi48++ {
            v49 := []int{-1, (fi48 + (fi48 << ((2) & 63)))}
            _ = fi48
            _ = v49
        }
    }
    if true {
        v50 := GenS4{F0: (false && ((false || (-1 >= 100)) || (!(!false))))}
        _ = v50
    } else {
        var v51 int16 = -653
        _ = v51
    }
    var v52 int32 = -2147483648
    switch 32767 {
    case 1:
        var v53 string = ("" + "hello")
        _ = v53
    }
    v3 -= (v43 & v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v43
    _ = v44
    _ = v5
    _ = v52
}
