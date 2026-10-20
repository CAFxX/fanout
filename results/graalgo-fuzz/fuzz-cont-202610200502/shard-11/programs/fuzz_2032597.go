package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint64; F1 int8; F2 int }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int8; F1 uint8 }
type GenS5 struct { F0 string; F1 int }
func genFn0() bool {
    var v7 int = 0
    fmt.Println(v7)
    if true {
        for fi8 := 0; fi8 < 16; fi8++ {
            for fi9 := 0; fi9 < 14; fi9++ {
                var v10 bool = false
                v7 = 402
                v10 = v10
                _ = fi9
                _ = v10
            }
            v11 := []int{(fi8 % ((526) | 1)), 32767}
            _ = fi8
            _ = v11
        }
    }
    _ = v7
    return ((-1 <= 100) || true)
}
func genFn1() int32 {
    v12 := []int{32767}
    if (!false) {
        var v13 string = "a"
        _ = v13
    }
    var v14 int64 = 100
    var v15 string = ("\n\t" + ("hello" + "a"))
    _ = v12
    _ = v14
    _ = v15
    return -2147483648
}

func main() {
    var v1 int32 = 2147483647
    v2 := []uint32{0, 65535, 42, 430}
    var v3 int16 = 32767
    var v4 int16 = -4
    var v5 uint = 1
    var v6 uint64 = 1
    v2 = append(v2, 2)
    if ((true && true) || (255 != 42)) {
        switch 2147483647 {
        case 890:
            fmt.Println(v5)
        case 1:
            v16 := GenS1{F0: 145}
            _ = v16
        default:
            v1 -= v1
        }
        for fi17 := 0; fi17 < 10; fi17++ {
            if true {
                if true {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    v5 += (v5 / ((4294967295) | 1))
                    v5 = (^(v5))
                } else {
                    for fi18 := 0; fi18 < 3; fi18++ {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        v3 += (v4 - v3)
                        _ = fi18
                    }
                    v1 -= v1
                }
                var v19 int16 = (v4 % ((v4) | 1))
                _ = v19
            } else {
                var v20 uint16 = 65535
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = v20
            }
            _ = fi17
        }
        var v21 uint32 = 0
        _ = v21
    }
    if len(v2) > 0 {
        v2[0] = 0
    }
    v6 += v6
    v4 = int16(v5)
    v3 += (v4 & 32767)
    switch ("foo bar" + ("a" + "")) {
    case "hello":
        for fi22 := 0; fi22 < 3; fi22++ {
            v23 := GenS2{F0: 925, F1: int8(v3), F2: fi22}
            if true {
                v5 = 2
                v5 -= (v5 >> (((v5 * (v5 | (v5 & 4294967295)))) & 63))
            }
            v4 = v3
            _ = fi22
            _ = v23
        }
    case "foo bar":
        v24 := GenS2{F0: (v6 | (v6 << ((v5) & 63))), F1: int8(v4), F2: 654}
        _ = v24
    default:
        if true {
            for fi25 := 0; fi25 < 17; fi25++ {
                v26 := GenS3{F0: (((0 == 725) || (v5 <= (^(v5)))) && true)}
                v2 = append(v2, 4294967295)
                fmt.Println(v4, v4)
                _ = fi25
                _ = v26
            }
        }
    }
    for fi27 := 0; fi27 < 18; fi27++ {
        var v28 uint32 = 2147483648
        var v29 uint16 = 669
        switch "\n\t" {
        case "":
            var v30 uint16 = 255
            _ = v30
        case "\n\t":
            v28 += (v28 % ((v28) | 1))
        case "a":
            v31 := GenS2{F0: v6, F1: -128, F2: (fi27 & (-(fi27)))}
            _ = v31
        }
        _ = fi27
        _ = v28
        _ = v29
    }
    v5 -= (v5 >> (((v5 + (v5 * (^(v5))))) & 63))
    if false {
        v6 += (v6 % ((v6) | 1))
    } else {
        switch ("foo bar" + "hello") {
        case "":
            switch int(v1) {
            case 155:
                v32 := GenS4{F0: 127, F1: 2}
                _ = v32
            case -832:
                var v33 int32 = v1
                _ = v33
            }
        case "\n\t":
            _ = len(v2)
        }
    }
    var v34 int64 = 1
    switch "foo bar" {
    case "\n\t":
        if (7 <= 42) {
            v4 = v4
        } else {
            for fi35 := 0; fi35 < 3; fi35++ {
                v36 := GenS5{F0: "\n\t", F1: (fi35 * 128)}
                _ = fi35
                _ = v36
            }
        }
    case "":
        var v37 int32 = 286
        _ = v37
    case "foo bar":
        var v38 float32 = (-2.25 * 3.14159)
        _ = v38
    default:
        for fi39 := 0; fi39 < 7; fi39++ {
            for fi40 := 0; fi40 < 13; fi40++ {
                var v41 uint64 = uint64(v5)
                switch fi39 {
                case -9223372036854775808:
                    var v42 uint = v5
                    _ = v42
                case 127:
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                case -559:
                    var v43 bool = true
                    _ = v43
                }
                v41 = v6
                _ = fi40
                _ = v41
            }
            _ = len(v2)
            _ = fi39
        }
    }
    if true {
        var v44 uint16 = 65535
        v34 += (v34 - (v34 ^ (v34 | (v34 % (((v34 | (^(v34)))) | 1)))))
        _ = v44
    } else {
        v3 -= v3
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
