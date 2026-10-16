package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16 }
type GenS2 struct { F0 uint; F1 int32; F2 int32 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int64; F1 uint; F2 int16 }
type GenS5 struct { F0 int64; F1 uint8; F2 int }
type GenS6 struct { F0 uint64 }
func genFn0(p0 uint8, p1 int16) uint32 {
    var v4 float64 = 3.14159
    var v5 int16 = p1
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 891
}
func genFn1(p0 int) bool {
    for fi6 := 0; fi6 < 5; fi6++ {
        v7 := []int{}
        for fi8 := 0; fi8 < 13; fi8++ {
            var v9 float32 = (0.0 + (((((-2.25 * 0.1) + (1.5 - 0.1)) + 0.1) / -2.25) - 0.1))
            _ = len(v7)
            v10 := GenS1{F0: 788, F1: 1000}
            _ = fi8
            _ = v10
            _ = v9
        }
        var v11 uint64 = 384
        _ = fi6
        _ = v11
        _ = v7
    }
    v12 := []int{p0, (p0 ^ -2147483648), 42}
    _ = p0
    _ = v12
    return (false || (false || false))
}

func main() {
    var v1 int16 = 1
    var v2 int = 127
    var v3 uint32 = 1
    v1 = (v1 | (v1 * 1000))
    if true {
        for fi13 := 0; fi13 < 15; fi13++ {
            v14 := GenS2{F0: 2, F1: 0, F2: 2147483647}
            v15 := []int{(fi13 - (^(v2))), (fi13 - fi13), (fi13 << ((4294967295) & 63))}
            if len(v15) >= 2 {
                v15 = v15[1:len(v15)]
            }
            _ = fi13
            _ = v14
            _ = v15
        }
        var v16 uint32 = v3
        _ = v16
    }
    if true {
        switch v2 {
        case 32767:
            var v17 uint8 = 14
            _ = v17
        case -125:
            var v18 int = v2
            _ = v18
        case -9223372036854775808:
            v19 := GenS3{F0: -66}
            _ = v19
        default:
            var v20 uint64 = 4294967295
            _ = v20
        }
        if ((v3 | 2147483648) <= v3) {
            var v21 uint64 = 18446744073709551615
            var v22 int = v2
            v3 -= v3
            _ = v21
            _ = v22
        }
    } else {
        var v23 uint = 443
        var v24 string = "\n\t"
        _ = v23
        _ = v24
    }
    for fi25 := 0; fi25 < 7; fi25++ {
        v26 := GenS3{F0: -114}
        fmt.Println(v2, fi25, v3)
        switch "" {
        case "":
            switch (v2 & (v2 % ((0) | 1))) {
            case 2147483647:
                v27 := []int{-9223372036854775808, -612, fi25, (v2 | -282)}
                _ = v27
            case -919:
                var v28 string = "a"
                _ = v28
            default:
                var v29 uint32 = v3
                _ = v29
            }
        default:
            v30 := []int{(fi25 - fi25), fi25, fi25, v2}
            _ = v30
        }
        _ = fi25
        _ = v26
    }
    v31 := []int{32768, (v2 * (v2 - 128)), v2}
    for fi32 := 0; fi32 < 13; fi32++ {
        fmt.Println(fi32, fi32)
        v2 = (v2 >> ((65535) & 63))
        _ = fi32
    }
    v2 -= v2
    v1 = 42
    for fi33 := 0; fi33 < 20; fi33++ {
        if len(v31) > 0 {
            v31[0] = fi33
        }
        var v34 uint8 = 2
        var v35 uint64 = 611
        _ = fi33
        _ = v34
        _ = v35
    }
    if (float32(-2.25) < -2.25) {
        if (-2147483648 == -2147483648) {
            if true {
                v3 -= 732
            }
            v36 := GenS4{F0: -9223372036854775808, F1: 264, F2: 42}
            v37 := GenS5{F0: -1, F1: 127, F2: (v2 % ((2147483647) | 1))}
            _ = v36
            _ = v37
        } else {
            _ = len(v31)
            if true {
                for fi38 := 0; fi38 < 5; fi38++ {
                    v2 = (^(v2))
                    v1 -= -457
                    if false {
                        v3 += genFn0(127, v1)
                        var v39 float64 = (0.1 - 0.0)
                        if len(v31) >= 2 {
                            v31 = v31[1:len(v31)]
                        }
                        _ = v39
                    }
                    _ = fi38
                }
            } else {
                if true {
                    v1 = (-(v1))
                    var v40 int64 = -1
                    _ = v40
                } else {
                    v41 := GenS1{F0: 1, F1: v1}
                    if len(v31) > 0 {
                        v31[0] = v2
                    }
                    var v42 int64 = 9223372036854775807
                    _ = v41
                    _ = v42
                }
            }
        }
        for fi43 := 0; fi43 < 7; fi43++ {
            fmt.Println(v2, fi43)
            _ = fi43
        }
        _ = len(v31)
    }
    for fi44 := 0; fi44 < 8; fi44++ {
        for fi45 := 0; fi45 < 12; fi45++ {
            v46 := GenS6{F0: 18446744073709551615}
            var v47 int = fi45
            _ = fi45
            _ = v46
            _ = v47
        }
        var v48 int64 = 1
        v2 -= (^(fi44))
        _ = fi44
        _ = v48
    }
    if len(v31) >= 2 {
        v31 = v31[1:len(v31)]
    }
    fmt.Println(v2, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v31:", len(v31), v31)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
}
