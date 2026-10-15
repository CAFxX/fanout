package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint; F2 uint16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint8; F1 uint16; F2 uint16 }
type GenS4 struct { F0 bool; F1 int64; F2 uint64 }
type GenS5 struct { F0 int64; F1 uint16 }
type GenS6 struct { F0 int64; F1 int8; F2 int8 }
func genFn0(p0 int16, p1 uint8, p2 bool) int {
    p0 += p0
    fmt.Println(p1, p1)
    p2 = true
    _ = p0
    _ = p1
    _ = p2
    return 10
}
func genFn1(p0 int8, p1 bool) int64 {
    p1 = (2147483648 >= 0)
    v7 := []int{-2147483648, 32768, 2147483647}
    _ = len(v7)
    _ = p0
    _ = p1
    _ = v7
    return 42
}

func main() {
    var v1 uint32 = 42
    var v2 int8 = -24
    var v3 int64 = 42
    var v4 int = 32768
    v5 := []int{v4, (^(v4))}
    var v6 int16 = -392
    if (!true) {
        var v8 int16 = v6
        _ = v8
    } else {
        fmt.Println(v4, v1)
        if (!false) {
            for fi9 := 0; fi9 < 16; fi9++ {
                v2 += 1
                switch v4 {
                case 32767:
                    switch (fi9 << ((18446744073709551615) & 63)) {
                    case -142:
                        fmt.Println(v3)
                    default:
                        fmt.Println(v2, v3, v3)
                    }
                default:
                    fmt.Println(v3, v4, v6)
                }
                fmt.Println(v3)
                _ = fi9
            }
        }
        var v10 int = 135
        _ = v10
    }
    var v11 bool = ((false || (!true)) && true)
    switch 0 {
    case -653:
        v12 := GenS1{F0: v3, F1: 18446744073709551615, F2: 32767}
        _ = v12
    case 127:
        for fi13 := 0; fi13 < 11; fi13++ {
            for fi14 := 0; fi14 < 5; fi14++ {
                v15 := GenS2{F0: -9223372036854775808}
                _ = fi14
                _ = v15
            }
            v1 = 2
            switch ("\n\t" + (("a" + "\n\t") + (("a" + ("" + "foo bar")) + (("hello" + "foo bar") + ("\n\t" + ""))))) {
            case "":
                var v16 bool = true
                _ = v16
            case "hello":
                fmt.Println(v6, v4, v1)
            case "a":
                v17 := GenS1{F0: v3, F1: 4294967295, F2: 2}
                _ = v17
            default:
                var v18 uint16 = 393
                _ = v18
            }
            _ = fi13
        }
    default:
        v11 = v11
    }
    v19 := GenS3{F0: 228, F1: 711, F2: 0}
    switch v4 {
    case 9223372036854775807:
        var v20 int = 414
        _ = v20
    case -903:
        var v21 float64 = (1e10 / 1.0)
        _ = v21
    case 32768:
        v3 = v3
    }
    switch ("foo bar" + "hello") {
    case "\n\t":
        for fi22 := 0; fi22 < 16; fi22++ {
            for fi23 := 0; fi23 < 14; fi23++ {
                if (1.5 <= 0.0) {
                    if len(v5) > 0 {
                        v5[0] = (fi22 - fi22)
                    }
                    if (220 > 184) {
                        v24 := GenS4{F0: ("foo bar" != "\n\t"), F1: int64(v6), F2: 784}
                        var v25 bool = v11
                        fmt.Println(v2, v2, v1)
                        _ = v24
                        _ = v25
                    } else {
                        v4 += v4
                        if len(v5) > 0 {
                            v5[0] = 42
                        }
                    }
                    switch "a" {
                    case "hello":
                        v11 = v11
                        var v26 int16 = (v6 / (((^(v6))) | 1))
                        fmt.Println(v2, v3, v11)
                        _ = v26
                    }
                }
                v3 -= 9223372036854775807
                _ = fi23
            }
            fmt.Println(v6, v1)
            var v27 float32 = 1.5
            _ = fi22
            _ = v27
        }
    case "a":
        v1 = (v1 % ((0) | 1))
    case "hello":
        v28 := GenS5{F0: (v3 & v3), F1: 0}
        _ = v28
    default:
        v29 := GenS6{F0: (-(v3)), F1: (v2 << ((585) & 63)), F2: v2}
        _ = v29
    }
    v3 = int64(v6)
    if v11 {
        if len(v5) > 0 {
            v5[0] = (-(v4))
        }
        var v30 int8 = v2
        var v31 uint64 = 17
        _ = v30
        _ = v31
    }
    var v32 int16 = (-(v6))
    v33 := GenS6{F0: 9223372036854775807, F1: (v2 % ((v2) | 1)), F2: 0}
    fmt.Println(v11, v3)
    var v34 float32 = (1.5 + 0.0)
    v35 := GenS2{F0: (-(v4))}
    var v36 int32 = -309
    fmt.Println(v11, v4)
    v11 = true
    v1 += (v1 * (^(v1)))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v19
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
