package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32 }
type GenS2 struct { F0 uint64; F1 uint64; F2 int8 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int }
func genFn0(p0 uint8, p1 int32, p2 int8) uint16 {
    v6 := []int{0, 4294967295, 9223372036854775807, 573}
    p0 -= (p0 - p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 32767
}
func genFn1(p0 int32) uint {
    switch "foo bar" {
    case "foo bar":
        var v7 uint = 254
        _ = v7
    case "hello":
        if (false || true) {
            v8 := []int{}
            switch ("hello" + "foo bar") {
            case "hello":
                var v9 bool = true
                _ = v9
            default:
                v10 := GenS1{F0: p0, F1: 2147483648}
                _ = v10
            }
            _ = v8
        }
    }
    if ((false || true) || (-129 != -434)) {
        var v11 int = 9223372036854775807
        var v12 int64 = 370
        if (-420 >= 0) {
            fmt.Println(v12, p0, v11)
            v13 := []int{}
            fmt.Println(v12, p0, v12)
            _ = v13
        }
        _ = v11
        _ = v12
    } else {
        if false {
            var v14 int64 = -9223372036854775808
            if (1 >= 18) {
                fmt.Println(p0, p0)
            } else {
                var v15 int16 = 0
                _ = v15
            }
            _ = v14
        }
        var v16 int32 = (p0 + p0)
        if (false || true) {
            var v17 int = -1
            var v18 int8 = 75
            for fi19 := 0; fi19 < 3; fi19++ {
                v20 := GenS2{F0: 4, F1: 18446744073709551615, F2: (v18 % ((100) | 1))}
                _ = fi19
                _ = v20
            }
            _ = v17
            _ = v18
        } else {
            fmt.Println(v16, p0)
            if (18446744073709551615 <= 392) {
                v21 := GenS3{F0: (v16 * p0)}
                var v22 uint = 4294967295
                _ = v21
                _ = v22
            }
            v23 := []int{42}
            _ = v23
        }
        _ = v16
    }
    v24 := GenS4{F0: -9223372036854775808}
    p0 = (p0 - 1)
    if false {
        p0 = (^(p0))
        v25 := []int{-2147483648, -688, int(p0)}
        _ = v25
    }
    _ = p0
    _ = v24
    return 123
}

func main() {
    var v1 uint32 = 0
    v2 := make(map[uint32]int)
    var v3 float64 = (((-2.25 / 3.14159) + 1.5) - 1e10)
    var v4 bool = (-401 != 4294967295)
    v5 := []int{-9223372036854775808}
    for fi26 := 0; fi26 < 10; fi26++ {
        var v27 int = fi26
        if v4 {
            v4 = v4
        }
        _ = fi26
        _ = v27
    }
    var v28 uint8 = 128
    if len(v5) > 0 {
        v5[0] = 301
    }
    var v29 float32 = 0.0
    fmt.Println(v4)
    v2[v1] = -406
    if v4 {
        if len(v5) > 0 {
            v5[0] = -597
        }
    } else {
        if (v29 >= v29) {
            var v30 int32 = 2147483647
            _ = v30
        } else {
            v1 += (v1 + (v1 / (((v1 >> ((genFn1(1)) & 63))) | 1)))
            var v31 int64 = 42
            v29 += (v29 - (((1.5 / 1.5) - float32(3.14159)) * (v29 - v29)))
            _ = v31
        }
    }
    switch 2147483647 {
    case 1:
        switch "" {
        case "foo bar":
            v1 = (v1 ^ 0)
        case "hello":
            if v4 {
                if false {
                    switch "a" {
                    case "hello":
                        fmt.Println(v1, v3)
                    case "a":
                        if false {
                            v28 = (v28 >> ((4294967295) & 63))
                            v32 := GenS5{F0: 2147483647}
                            var v33 int16 = 1
                            _ = v32
                            _ = v33
                        } else {
                            var v34 uint8 = (v28 * v28)
                            var v35 int32 = 627
                            _ = v34
                            _ = v35
                        }
                    case "foo bar":
                        v2[(v1 & 576)] = 32768
                    default:
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    }
                    for fi36 := 0; fi36 < 1; fi36++ {
                        delete(v2, v1)
                        _ = fi36
                    }
                }
            } else {
                var v37 int = -2147483648
                _ = len(v5)
                _ = v37
            }
        default:
            fmt.Println(v1)
        }
    case 32768:
        v38 := v2[v1]
        _ = v38
    default:
        var v39 uint8 = (^(v28))
        _ = v39
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
