package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool; F2 uint }
type GenS2 struct { F0 int64; F1 int32; F2 bool }
type GenS3 struct { F0 int8; F1 uint; F2 int64 }
func genFn0() uint {
    var v7 int32 = 1
    var v8 int32 = 738
    if false {
        var v9 uint32 = 260
        var v10 float32 = -2.25
        _ = v10
        _ = v9
    }
    _ = v7
    _ = v8
    return 0
}

func main() {
    var v1 int16 = 0
    v2 := []int{}
    v3 := make(map[int]uint32)
    v4 := make(map[uint32]bool)
    var v5 uint = 706
    var v6 int16 = (v1 ^ v1)
    v6 = (v6 / (((v1 >> ((1) & 63))) | 1))
    v6 -= (v1 | v6)
    for fi11 := 0; fi11 < 9; fi11++ {
        var v12 int = int(v1)
        _ = fi11
        _ = v12
    }
    for fi13 := 0; fi13 < 2; fi13++ {
        var v14 float64 = (0.1 + 1.5)
        _ = fi13
        _ = v14
    }
    v5 = (v5 + 18446744073709551615)
    if true {
        switch "" {
        case "\n\t":
            v5 = v5
        case "foo bar":
            _ = len(v2)
        case "hello":
            var v15 int32 = 42
            _ = v15
        default:
            v16 := GenS1{F0: "foo bar", F1: (false || false), F2: 2}
            _ = v16
        }
        var v17 int32 = -2147483648
        _ = v17
    }
    for fi18 := 0; fi18 < 12; fi18++ {
        delete(v3, fi18)
        v6 += 757
        _ = fi18
    }
    switch -994 {
    case 407:
        v19 := GenS1{F0: "\n\t", F1: true, F2: genFn0()}
        _ = v19
    case 4294967295:
        v1 -= v1
    default:
        switch -9223372036854775808 {
        case -305:
            v20 := GenS1{F0: "foo bar", F1: false, F2: 585}
            _ = v20
        case 772:
            for fi21 := 0; fi21 < 8; fi21++ {
                for fi22 := 0; fi22 < 17; fi22++ {
                    var v23 uint8 = 127
                    v23 = (v23 - (v23 >> (((^(v5))) & 63)))
                    var v24 uint32 = 65535
                    v24 -= 4294967295
                    var v25 uint8 = uint8(v24)
                    _ = fi22
                    _ = v23
                    _ = v24
                    _ = v25
                }
                if true {
                    v26 := GenS2{F0: -1, F1: 0, F2: false}
                    _ = v26
                } else {
                    fmt.Println(v1, fi21)
                    v27 := GenS1{F0: ("a" + ""), F1: (-1 >= (v6 | v6)), F2: 4294967295}
                    _ = v27
                }
                v6 -= int16(fi21)
                _ = fi21
            }
        case 42:
            var v28 float32 = 3.14159
            _ = v28
        }
    }
    var v29 int32 = 2147483647
    if true {
        for fi30 := 0; fi30 < 8; fi30++ {
            for fi31 := 0; fi31 < 11; fi31++ {
                v2 = append(v2, 0)
                if true {
                    _ = len(v2)
                    var v32 int64 = -596
                    if len(v2) > 0 {
                        v2[0] = fi30
                    }
                    _ = v32
                } else {
                    if false {
                        v3[(fi30 / (((-(fi31))) | 1))] = 502
                    }
                }
                v5 = v5
                _ = fi31
            }
            _ = fi30
        }
        var v33 int8 = 94
        if (float64(-2.25) != ((3.14159 * float64(1e10)) + (1e10 / (0.1 + (1.5 * -2.25))))) {
            v2 = append(v2, 128)
            v5 = v5
        } else {
            if ((127 != 2) && true) {
                v34 := GenS3{F0: v33, F1: genFn0(), F2: 9223372036854775807}
                v35 := GenS2{F0: -9223372036854775808, F1: (v29 | (v29 | (v29 + v29))), F2: false}
                if len(v2) > 0 {
                    v2[0] = -2147483648
                }
                _ = v34
                _ = v35
            }
            for fi36 := 0; fi36 < 19; fi36++ {
                v5 = (v5 * (v5 + 920))
                _ = fi36
            }
        }
        _ = v33
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
