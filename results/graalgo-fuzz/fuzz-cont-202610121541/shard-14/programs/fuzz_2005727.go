package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32 }
type GenS2 struct { F0 uint; F1 int16; F2 bool }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int; F1 string; F2 uint8 }
type GenS6 struct { F0 string; F1 string; F2 uint64 }
type GenS7 struct { F0 uint8; F1 int64 }
func genFn0(p0 int16, p1 uint8) string {
    for fi4 := 0; fi4 < 5; fi4++ {
        var v5 int8 = -128
        switch fi4 {
        case 2147483647:
            v6 := GenS1{F0: 4294967295, F1: 42}
            _ = v6
        }
        _ = fi4
        _ = v5
    }
    if ((("a" + "") + "") < (("foo bar" + ("\n\t" + "a")) + ("a" + ""))) {
        var v7 int8 = 127
        var v8 uint32 = 0
        p1 += (p1 << ((1) & 63))
        _ = v7
        _ = v8
    }
    var v9 int8 = 36
    for fi10 := 0; fi10 < 14; fi10++ {
        fmt.Println(p1, fi10, p1)
        p1 -= 42
        _ = fi10
    }
    _ = p0
    _ = p1
    _ = v9
    return "\n\t"
}

func main() {
    var v1 bool = (255 != 2)
    var v2 bool = v1
    var v3 uint32 = 537
    var v11 int = -9223372036854775808
    v12 := GenS1{F0: v3, F1: 4294967295}
    v2 = (!true)
    if (!((v11 ^ v11) <= 1)) {
        for fi13 := 0; fi13 < 17; fi13++ {
            for fi14 := 0; fi14 < 11; fi14++ {
                v15 := GenS1{F0: (v3 * (v3 | (v3 >> ((42) & 63)))), F1: v3}
                switch v11 {
                case -9223372036854775808:
                    if v2 {
                        if false {
                            fmt.Println(fi14, fi13, fi13)
                            v16 := []int{32768, v11, 32767, v11}
                            _ = v16
                        }
                        v17 := GenS2{F0: 11, F1: 1, F2: (((^(v11)) != (fi13 - 9223372036854775807)) && v1)}
                        v18 := []int{}
                        _ = v17
                        _ = v18
                    }
                case -982:
                    if (v2 && v2) {
                        v19 := GenS3{F0: uint8(v3)}
                        _ = v19
                    }
                default:
                    var v20 bool = ((!false) || v2)
                    _ = v20
                }
                v11 = -9223372036854775808
                v3 += v3
                _ = fi14
                _ = v15
            }
            var v21 int = (fi13 - (v11 * 42))
            v22 := GenS4{F0: v3}
            _ = fi13
            _ = v21
            _ = v22
        }
        for fi23 := 0; fi23 < 4; fi23++ {
            var v24 int64 = 659
            var v25 float64 = 1.5
            v26 := []int{0, fi23}
            _ = fi23
            _ = v24
            _ = v25
            _ = v26
        }
    } else {
        var v27 float32 = 1.5
        _ = v27
    }
    fmt.Println(v11)
    v28 := GenS4{F0: (^(v3))}
    for fi29 := 0; fi29 < 3; fi29++ {
        v30 := []int{(^(v11)), fi29, (v11 % (((v11 ^ (v11 * v11))) | 1))}
        var v31 float32 = -2.25
        _ = fi29
        _ = v30
        _ = v31
    }
    for fi32 := 0; fi32 < 15; fi32++ {
        for fi33 := 0; fi33 < 12; fi33++ {
            var v34 uint32 = v3
            v35 := GenS5{F0: (v11 >> ((18446744073709551615) & 63)), F1: "foo bar", F2: 127}
            v36 := GenS6{F0: ((("" + "\n\t") + "") + "a"), F1: "\n\t", F2: 173}
            v37 := []int{}
            v38 := GenS7{F0: 128, F1: -451}
            _ = fi33
            _ = v34
            _ = v35
            _ = v36
            _ = v37
            _ = v38
        }
        v39 := []int{v11, (v11 ^ fi32)}
        _ = fi32
        _ = v39
    }
    v40 := []int{v11}
    v3 -= v3
    var v41 uint16 = 0
    v1 = (!(false || (((!(!((v3 - v3) <= 2147483648))) && (1e10 < ((1e10 + 3.14159) + (1.5 / 1.5)))) || (-32768 == 820))))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v40:", len(v40), v40)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v28
    _ = v3
    _ = v40
    _ = v41
}
