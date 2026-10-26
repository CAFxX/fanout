package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint64; F1 int }
func genFn0(p0 bool, p1 uint16) uint8 {
    for fi4 := 0; fi4 < 11; fi4++ {
        var v5 int = (^(fi4))
        for fi6 := 0; fi6 < 7; fi6++ {
            var v7 float32 = 3.14159
            _ = fi6
            _ = v7
        }
        _ = fi4
        _ = v5
    }
    fmt.Println(p1)
    _ = p0
    _ = p1
    return 237
}
func genFn1(p0 int32) uint32 {
    if (1 < 2147483647) {
        v8 := []int{int(p0), 0, -2147483648, -9223372036854775808}
        _ = v8
    }
    v9 := []int{-1, 32767, -2147483648}
    v10 := GenS1{F0: 2}
    switch -2147483648 {
    case 42:
        v11 := GenS2{F0: 18446744073709551615, F1: -814}
        _ = v11
    case -9223372036854775808:
        if len(v9) > 0 {
            v9[0] = 42
        }
    }
    _ = p0
    _ = v10
    _ = v9
    return 1
}

func main() {
    v1 := []uint32{}
    var v2 int32 = 0
    var v3 uint16 = 0
    var v12 uint64 = 325
    var v13 uint64 = (^(v12))
    v13 = (^(v13))
    v12 += v13
    v13 += v12
    v13 = 497
    v14 := GenS1{F0: 304}
    v13 -= (^(v13))
    var v15 int16 = -1
    var v16 uint8 = 236
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v17 uint = 42
    var v18 float32 = (3.14159 + 1.5)
    for fi19 := 0; fi19 < 2; fi19++ {
        if (1.5 == ((1e10 - 1.5) - (0.0 - ((0.0 + 1e10) - (float64(float32(0.0)) / float64((3.14159 + 0.0))))))) {
            v1 = append(v1, genFn1(-2147483648))
            fmt.Println(v3, v2)
            v16 -= (v16 / (((v16 >> ((v17) & 63))) | 1))
        } else {
            v20 := GenS1{F0: 4294967295}
            v13 = 2
            _ = v20
        }
        _ = fi19
    }
    if (2 < genFn1((v2 - (v2 % ((0) | 1))))) {
        var v21 int8 = -111
        switch ("a" + "hello") {
        case "foo bar":
            var v22 uint16 = (v3 | 255)
            _ = v22
        case "\n\t":
            v15 -= (v15 * 0)
        case "":
            var v23 int16 = (v15 + (v15 | int16(v16)))
            _ = v23
        default:
            v12 += 1
        }
        _ = v21
    }
    if (true || (v13 > (^(v12)))) {
        if ((v17 + (v17 + v17)) > (v17 & uint(v13))) {
            for fi24 := 0; fi24 < 5; fi24++ {
                v12 += (^(v12))
                v2 -= v2
                _ = fi24
            }
            var v25 int8 = int8(v16)
            _ = v25
        }
        switch -9223372036854775808 {
        case 912:
            fmt.Println(v17, v16)
        case 935:
            switch -735 {
            case 32767:
                if len(v1) > 0 {
                    v1[0] = 65535
                }
            }
        case -1:
            v26 := GenS2{F0: (v12 >> (((v17 + v17)) & 63)), F1: int(v15)}
            _ = v26
        default:
            switch 233 {
            case -833:
                for fi27 := 0; fi27 < 18; fi27++ {
                    var v28 uint64 = 2
                    if len(v1) > 0 {
                        v1[0] = 2147483648
                    }
                    _ = fi27
                    _ = v28
                }
            }
        }
        v17 -= 2
    } else {
        v13 = v12
        v15 -= v15
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
}
