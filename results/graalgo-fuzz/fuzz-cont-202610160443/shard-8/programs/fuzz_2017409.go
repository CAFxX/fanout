package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16; F2 bool }
type GenS2 struct { F0 uint16; F1 uint32; F2 uint64 }
type GenS3 struct { F0 int; F1 uint64 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint64 }

func main() {
    var v1 int16 = 1
    var v2 uint8 = 42
    var v3 uint16 = 255
    var v4 uint8 = v2
    v5 := []int{0, -9223372036854775808, 32768}
    switch 128 {
    case 42:
        fmt.Println(v4, v1)
    case -1:
        if len(v5) > 0 {
            v5[0] = -1
        }
    case -2147483648:
        v3 = 0
    default:
        for fi6 := 0; fi6 < 4; fi6++ {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            v2 = v4
            _ = fi6
        }
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v7 := GenS1{F0: 0, F1: -1, F2: (-235 == 1)}
    v5 = append(v5, -988)
    switch -122 {
    case 1:
        v8 := GenS1{F0: 2147483647, F1: (v1 * (v1 % ((v1) | 1))), F2: (!true)}
        _ = v8
    default:
        var v9 int32 = 0
        _ = v9
    }
    var v10 uint32 = 65535
    v2 -= 2
    switch "hello" {
    case "a":
        fmt.Println(v2, v4, v10)
    case "foo bar":
        var v11 int32 = -2147483648
        _ = v11
    case "":
        fmt.Println(v3, v1, v1)
    default:
        var v12 uint = 266
        _ = v12
    }
    var v13 bool = (((!false) && false) && false)
    var v14 int8 = 1
    var v15 int32 = -1
    v5 = append(v5, -2147483648)
    v16 := GenS2{F0: v3, F1: v10, F2: 18446744073709551615}
    switch "\n\t" {
    case "hello":
        v15 = v15
    default:
        if (515 < 18446744073709551615) {
            v17 := GenS1{F0: (v15 << ((4294967295) & 63)), F1: 1, F2: (!(v15 != (v15 & -451)))}
            switch -9223372036854775808 {
            case 308:
                switch -180 {
                case 4294967295:
                    v18 := GenS3{F0: 42, F1: 882}
                    _ = v18
                }
            default:
                v19 := GenS4{F0: 591}
                _ = v19
            }
            _ = v17
        } else {
            v13 = (!v13)
            for fi20 := 0; fi20 < 18; fi20++ {
                v15 += -921
                switch "\n\t" {
                case "hello":
                    v21 := GenS5{F0: 6}
                    _ = v21
                case "":
                    var v22 int64 = 42
                    _ = v22
                case "foo bar":
                    if v13 {
                        v13 = (912 <= 1)
                        var v23 float32 = float32(float64(-2.25))
                        _ = v23
                    } else {
                        v14 -= int8(v1)
                        _ = len(v5)
                    }
                default:
                    v24 := GenS5{F0: 4294967295}
                    _ = v24
                }
                if len(v5) > 0 {
                    v5[0] = fi20
                }
                _ = fi20
            }
            var v25 float64 = -2.25
            _ = v25
        }
    }
    for fi26 := 0; fi26 < 11; fi26++ {
        var v27 uint16 = 765
        v28 := GenS3{F0: (fi26 / ((fi26) | 1)), F1: 0}
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = fi26
        _ = v27
        _ = v28
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
