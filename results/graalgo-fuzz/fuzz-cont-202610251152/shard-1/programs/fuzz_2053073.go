package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 int16; F1 bool; F2 string }
type GenS3 struct { F0 string; F1 uint64 }
type GenS4 struct { F0 uint; F1 uint16 }
func genFn0(p0 uint16) int8 {
    if false {
        for fi6 := 0; fi6 < 10; fi6++ {
            p0 += (p0 & p0)
            fmt.Println(fi6, p0)
            var v7 uint32 = 65535
            _ = fi6
            _ = v7
        }
        p0 += (^(p0))
        for fi8 := 0; fi8 < 19; fi8++ {
            fmt.Println(fi8)
            _ = fi8
        }
    } else {
        for fi9 := 0; fi9 < 16; fi9++ {
            v10 := GenS1{F0: 32767, F1: 336}
            p0 = 806
            v11 := []int{fi9, (fi9 << ((0) & 63))}
            _ = fi9
            _ = v10
            _ = v11
        }
        v12 := GenS1{F0: p0, F1: uint32(p0)}
        p0 -= p0
        _ = v12
    }
    for fi13 := 0; fi13 < 6; fi13++ {
        for fi14 := 0; fi14 < 7; fi14++ {
            p0 += (p0 | 255)
            _ = fi14
        }
        _ = fi13
    }
    p0 -= 124
    _ = p0
    return 127
}

func main() {
    var v1 int32 = -412
    var v2 uint = 42
    var v3 uint = 18446744073709551615
    v4 := []uint32{2, 0}
    var v5 uint32 = 2147483647
    switch -1 {
    case 654:
        v5 += v5
    case -564:
        switch ("foo bar" + ("" + "hello")) {
        case "":
            v2 += (^(v2))
        default:
            v1 = (v1 + 0)
        }
    case 2147483647:
        var v15 uint16 = 65535
        _ = v15
    }
    if false {
        var v16 uint64 = 73
        if len(v4) > 0 {
            v4[0] = 1
        }
        v17 := GenS2{F0: -1, F1: ((!false) || (212 > 1)), F2: ("hello" + "")}
        _ = v16
        _ = v17
    } else {
        v18 := GenS3{F0: "a", F1: 1}
        for fi19 := 0; fi19 < 1; fi19++ {
            for fi20 := 0; fi20 < 20; fi20++ {
                var v21 int16 = 42
                if true {
                    fmt.Println(v5, v2)
                    _ = len(v4)
                }
                v21 -= (v21 + (v21 / (((-(v21))) | 1)))
                v1 = v1
                _ = fi20
                _ = v21
            }
            if ((v1 - (v1 >> ((v3) & 63))) < (v1 >> (((v3 / ((12) | 1))) & 63))) {
                for fi22 := 0; fi22 < 9; fi22++ {
                    var v23 float32 = ((-2.25 - ((1e10 - 1e10) + (3.14159 / (0.1 * float32(1e10))))) * 0.0)
                    var v24 int16 = 32767
                    _ = fi22
                    _ = v23
                    _ = v24
                }
                v4 = append(v4, v5)
            }
            v25 := GenS3{F0: ("\n\t" + "\n\t"), F1: 42}
            _ = fi19
            _ = v25
        }
        _ = v18
    }
    var v26 int32 = (v1 >> ((v3) & 63))
    for fi27 := 0; fi27 < 12; fi27++ {
        if true {
            var v28 uint16 = 1
            if (false && (255 > (v28 ^ uint16(v5)))) {
                fmt.Println(v1, v5)
                var v29 int16 = 32767
                if (3.14159 >= 1e10) {
                    if (!false) {
                        var v30 uint = (^(v3))
                        v30 -= v30
                        _ = v30
                    } else {
                        if len(v4) > 0 {
                            v4[0] = (v5 | (v5 << (((v2 & v2)) & 63)))
                        }
                        if len(v4) > 0 {
                            v4[0] = 198
                        }
                    }
                }
                _ = v29
            }
            _ = v28
        } else {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            for fi31 := 0; fi31 < 13; fi31++ {
                var v32 int = (fi27 | (fi31 & fi27))
                v5 = (v5 - (v5 + (v5 << ((18446744073709551615) & 63))))
                _ = fi31
                _ = v32
            }
            var v33 int16 = 42
            _ = v33
        }
        _ = fi27
    }
    if true {
        for fi34 := 0; fi34 < 13; fi34++ {
            var v35 uint = (v3 * v3)
            _ = fi34
            _ = v35
        }
        v36 := GenS1{F0: 2, F1: 2147483647}
        for fi37 := 0; fi37 < 2; fi37++ {
            switch "" {
            case "\n\t":
                fmt.Println(v1, v2, v2)
            case "foo bar":
                for fi38 := 0; fi38 < 1; fi38++ {
                    var v39 uint32 = (v5 % (((v5 ^ (v5 / ((1) | 1)))) | 1))
                    _ = fi38
                    _ = v39
                }
            case "hello":
                v40 := GenS4{F0: v3, F1: 943}
                _ = v40
            }
            for fi41 := 0; fi41 < 14; fi41++ {
                fmt.Println(v2, v1, v2)
                _ = fi41
            }
            fmt.Println(v26, v3, v2)
            _ = fi37
        }
        _ = v36
    }
    for fi42 := 0; fi42 < 9; fi42++ {
        var v43 uint16 = 0
        _ = len(v4)
        v26 = v26
        _ = fi42
        _ = v43
    }
    for fi44 := 0; fi44 < 2; fi44++ {
        var v45 float32 = ((float32(1e10) / 3.14159) * 1.5)
        v1 += (v1 / ((int32(v5)) | 1))
        _ = fi44
        _ = v45
    }
    fmt.Println(v5, v5, v3)
    var v46 bool = (9223372036854775807 <= -588)
    fmt.Println(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v46
    _ = v5
}
