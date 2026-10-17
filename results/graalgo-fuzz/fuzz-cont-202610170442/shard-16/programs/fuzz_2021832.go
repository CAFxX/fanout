package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int64; F1 uint32 }
type GenS3 struct { F0 uint64; F1 int8 }
type GenS4 struct { F0 uint16; F1 int; F2 int64 }
func genFn0(p0 uint) uint8 {
    var v6 float64 = (1e10 - -2.25)
    v7 := GenS1{F0: 32767}
    switch ("foo bar" + ("foo bar" + "")) {
    case "":
        if false {
            for fi8 := 0; fi8 < 11; fi8++ {
                v6 += 1.5
                _ = fi8
            }
            p0 -= (p0 << ((p0) & 63))
            fmt.Println(v6, v6)
        }
    case "hello":
        for fi9 := 0; fi9 < 1; fi9++ {
            switch -2147483648 {
            case 127:
                var v10 uint16 = 0
                _ = v10
            case -9223372036854775808:
                p0 = 0
            }
            v6 = (0.1 - (((v6 / v6) + v6) + -2.25))
            p0 += 15
            _ = fi9
        }
    case "foo bar":
        var v11 uint32 = 0
        _ = v11
    default:
        var v12 uint64 = 2
        _ = v12
    }
    _ = p0
    _ = v6
    _ = v7
    return 1
}

func main() {
    var v1 float64 = (((-2.25 - ((float64(-2.25) + ((-2.25 * 0.0) - ((3.14159 / (((0.1 - -2.25) * 1e10) / -2.25)) + 3.14159))) * 0.0)) + 1.5) / 1.5)
    v2 := []uint32{154}
    var v3 int32 = -1
    v4 := make(map[int]string)
    v5 := make(map[uint32]uint32)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    switch "foo bar" {
    case "\n\t":
        v13 := GenS2{F0: -9223372036854775808, F1: 766}
        _ = v13
    case "a":
        var v14 uint64 = 240
        _ = v14
    case "":
        for fi15 := 0; fi15 < 15; fi15++ {
            if (-1 != 573) {
                v3 = v3
            } else {
                v1 = (float64(0.1) - ((1.5 + v1) - 0.1))
            }
            _ = fi15
        }
    default:
        fmt.Println(v1, v1)
    }
    if (!true) {
        for fi16 := 0; fi16 < 1; fi16++ {
            v3 = v3
            _ = fi16
        }
    } else {
        var v17 string = (("" + "hello") + ("a" + ("" + "\n\t")))
        if true {
            var v18 string = v17
            v19 := GenS3{F0: 0, F1: 127}
            var v20 uint8 = 51
            _ = v18
            _ = v19
            _ = v20
        } else {
            fmt.Println(v1, v3, v3)
        }
        if len(v2) > 0 {
            v2[0] = 2147483647
        }
        _ = v17
    }
    v3 = v3
    fmt.Println(v1)
    fmt.Println(v1, v3)
    if (-1 >= -198) {
        var v21 int64 = -230
        v21 += (v21 + v21)
        if (false && false) {
            for fi22 := 0; fi22 < 5; fi22++ {
                var v23 uint16 = 65535
                _ = fi22
                _ = v23
            }
            for fi24 := 0; fi24 < 6; fi24++ {
                var v25 string = ""
                v3 += v3
                var v26 bool = (!(!false))
                _ = fi24
                _ = v25
                _ = v26
            }
        }
        _ = v21
    } else {
        for fi27 := 0; fi27 < 5; fi27++ {
            switch "" {
            case "hello":
                for fi28 := 0; fi28 < 4; fi28++ {
                    fmt.Println(fi27, fi27, fi28)
                    _ = fi28
                }
            case "foo bar":
                v4[(fi27 * (-(fi27)))] = "\n\t"
            case "\n\t":
                var v29 string = "hello"
                _ = v29
            default:
                for fi30 := 0; fi30 < 12; fi30++ {
                    switch fi30 {
                    case -984:
                        switch "foo bar" {
                        case "":
                            if (fi30 < (^(fi27))) {
                                var v31 uint8 = 0
                                v32 := GenS4{F0: 32767, F1: fi30, F2: 0}
                                _ = v31
                                _ = v32
                            } else {
                                _ = len(v2)
                                var v33 float64 = 0.0
                                var v34 int64 = -9223372036854775808
                                _ = v33
                                _ = v34
                            }
                            var v35 bool = (!false)
                            _ = v35
                        case "\n\t":
                            _ = len(v5)
                        case "foo bar":
                            v36 := GenS2{F0: -1, F1: 952}
                            _ = v36
                        }
                    }
                    var v37 uint16 = 42
                    var v38 bool = ("foo bar" <= "foo bar")
                    _ = fi30
                    _ = v37
                    _ = v38
                }
            }
            v3 -= 42
            v3 -= (v3 >> ((255) & 63))
            _ = fi27
        }
        v1 = (v1 / v1)
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    v3 -= (-(v3))
    for fi39 := 0; fi39 < 13; fi39++ {
        for fi40 := 0; fi40 < 14; fi40++ {
            delete(v4, (fi39 >> ((2) & 63)))
            _ = fi40
        }
        v1 -= v1
        _ = fi39
    }
    v41 := GenS1{F0: 158}
    v3 = v3
    fmt.Println(v3)
    v1 = float64(-2.25)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v5
}
