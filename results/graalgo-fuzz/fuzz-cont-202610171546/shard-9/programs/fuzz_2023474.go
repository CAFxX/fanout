package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint8; F1 int64; F2 uint }
type GenS4 struct { F0 bool; F1 uint8; F2 uint8 }
type GenS5 struct { F0 uint16; F1 uint64 }
type GenS6 struct { F0 uint64; F1 int; F2 int16 }
func genFn0(p0 uint32, p1 int) int16 {
    if (!(32767 > 65535)) {
        var v7 int16 = -32768
        v8 := GenS1{F0: 2147483647, F1: 2}
        _ = v7
        _ = v8
    } else {
        if (32767 < 255) {
            p0 = (p0 / ((p0) | 1))
            var v9 uint = 0
            v10 := []int{42, p1}
            _ = v10
            _ = v9
        }
        v11 := []int{(p1 & p1), (-(p1)), (p1 ^ p1)}
        _ = v11
    }
    var v12 float64 = (float64(1.5) / -2.25)
    _ = p0
    _ = p1
    _ = v12
    return 32767
}
func genFn1(p0 int8, p1 bool, p2 int16) int64 {
    p1 = false
    v13 := []int{-593, -2147483648, -9223372036854775808}
    if (false || true) {
        fmt.Println(p0, p0)
    } else {
        v14 := GenS1{F0: -699, F1: 1}
        _ = v14
    }
    p1 = p1
    var v15 int8 = (p0 / ((0) | 1))
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v15
    return -132
}

func main() {
    v1 := []string{"hello", "\n\t", ""}
    v2 := []int{2147483647}
    var v3 int = 32767
    var v4 int64 = 1
    var v5 bool = true
    var v6 int64 = v4
    v3 = (v3 / ((2147483647) | 1))
    var v16 int8 = 100
    var v17 string = "a"
    if true {
        for fi18 := 0; fi18 < 20; fi18++ {
            v19 := GenS2{F0: 4294967295}
            _ = len(v1)
            switch 128 {
            case -875:
                var v20 uint8 = 128
                _ = v20
            case 2147483647:
                v21 := GenS3{F0: 1, F1: 42, F2: 42}
                _ = v21
            case 388:
                fmt.Println(v5)
            default:
                if v5 {
                    var v22 float64 = 3.14159
                    var v23 uint64 = 1
                    v5 = false
                    _ = v22
                    _ = v23
                }
            }
            _ = fi18
            _ = v19
        }
        v24 := GenS4{F0: v5, F1: 0, F2: 128}
        v25 := GenS5{F0: 498, F1: 2}
        _ = v24
        _ = v25
    } else {
        for fi26 := 0; fi26 < 4; fi26++ {
            fmt.Println(v17)
            v27 := GenS5{F0: 181, F1: 0}
            _ = fi26
            _ = v27
        }
        fmt.Println(v5, v4, v3)
    }
    v28 := GenS5{F0: 42, F1: 101}
    for fi29 := 0; fi29 < 14; fi29++ {
        for fi30 := 0; fi30 < 5; fi30++ {
            if true {
                if ((fi30 >> ((358) & 63)) >= (fi30 / (((v3 >> ((0) & 63))) | 1))) {
                    switch v3 {
                    case 127:
                        if (!(v5 || v5)) {
                            fmt.Println(fi30, v3)
                            v17 = (v17 + ((v17 + v17) + ""))
                            v17 = ("foo bar" + v17)
                        } else {
                            v16 = v16
                            fmt.Println(v17)
                            if v5 {
                                fmt.Println(fi29, v3)
                                v16 = 100
                            }
                            if (1 > -32768) {
                                v17 = ""
                                if (!v5) {
                                    if (255 < 1) {
                                        switch "foo bar" {
                                        case "hello":
                                            v1 = append(v1, ("a" + v17))
                                        case "a":
                                            v6 += (v6 ^ v6)
                                        case "\n\t":
                                            if v5 {
                                                v5 = false
                                            }
                                        }
                                        var v31 float32 = 0.1
                                        v17 = "foo bar"
                                        _ = v31
                                    }
                                }
                                fmt.Println(fi29)
                                v2 = append(v2, (fi29 ^ v3))
                            }
                        }
                    }
                }
            } else {
                fmt.Println(v5, v6)
                fmt.Println(v3, v3)
                v32 := GenS1{F0: 2147483647, F1: 1}
                _ = v32
            }
            _ = fi30
        }
        var v33 float64 = (float64((-2.25 * 0.1)) * 1e10)
        for fi34 := 0; fi34 < 15; fi34++ {
            v1 = append(v1, v17)
            v35 := GenS6{F0: 18446744073709551615, F1: -9223372036854775808, F2: 0}
            _ = fi34
            _ = v35
        }
        _ = fi29
        _ = v33
    }
    fmt.Println(v17)
    var v36 float64 = (float64(0.0) - 1.5)
    switch v3 {
    case 2147483647:
        if len(v1) > 0 {
            v1[0] = "hello"
        }
    case -280:
        v6 -= -970
    default:
        v6 -= (v6 >> ((4294967295) & 63))
    }
    fmt.Println(v36, v36)
    v1 = append(v1, v17)
    var v37 bool = ((0 <= 42) || (65535 < 1))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v28
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
