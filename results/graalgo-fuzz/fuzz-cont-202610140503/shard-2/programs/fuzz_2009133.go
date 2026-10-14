package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint; F2 uint }
type GenS2 struct { F0 uint16; F1 uint32 }
type GenS3 struct { F0 uint; F1 string; F2 uint8 }
type GenS4 struct { F0 int64; F1 uint32 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int16; F1 string }
type GenS7 struct { F0 int64; F1 uint64; F2 int16 }

func main() {
    var v1 uint16 = 255
    v2 := make(map[int]bool)
    var v3 uint8 = 127
    var v4 float64 = -2.25
    v5 := GenS1{F0: 42, F1: uint(v3), F2: 255}
    v4 = v4
    var v6 uint32 = 2147483647
    for fi7 := 0; fi7 < 17; fi7++ {
        for fi8 := 0; fi8 < 16; fi8++ {
            v1 += 497
            switch 4294967295 {
            case 32768:
                if true {
                    var v9 float64 = (1e10 / 1.5)
                    v10 := GenS2{F0: (v1 % ((42) | 1)), F1: (v6 - 1)}
                    _ = v10
                    _ = v9
                }
            case 127:
                v6 = (^(v6))
            case 32767:
                v11 := GenS1{F0: 1000, F1: uint(v6), F2: 102}
                _ = v11
            }
            v1 += v1
            v12 := GenS3{F0: 1, F1: ("" + ("foo bar" + ("hello" + (("\n\t" + ("\n\t" + "hello")) + (((("\n\t" + ("" + "")) + "foo bar") + "hello") + ("hello" + ("foo bar" + "\n\t"))))))), F2: (^(v3))}
            _ = fi8
            _ = v12
        }
        v6 -= v6
        for fi13 := 0; fi13 < 18; fi13++ {
            var v14 float32 = 0.0
            _ = fi13
            _ = v14
        }
        _ = fi7
    }
    v2[127] = false
    v15 := []int{128, 4294967295, -9223372036854775808}
    for fi16 := 0; fi16 < 10; fi16++ {
        var v17 int8 = 85
        switch ("" + "") {
        case "a":
            fmt.Println(v4, v3, v1)
        case "":
            v15 = append(v15, fi16)
        case "hello":
            v6 -= (v6 % ((v6) | 1))
        }
        for fi18 := 0; fi18 < 9; fi18++ {
            fmt.Println(fi18, v17, v1)
            var v19 string = "hello"
            if true {
                v20 := GenS3{F0: 255, F1: (v19 + v19), F2: (v3 ^ v3)}
                _ = v20
            } else {
                v2[127] = true
                fmt.Println(v4, v1, v6)
            }
            _ = fi18
            _ = v19
        }
        _ = fi16
        _ = v17
    }
    v21 := GenS1{F0: 477, F1: 2, F2: 357}
    if false {
        var v22 uint8 = v3
        switch "foo bar" {
        case "\n\t":
            switch "" {
            case "":
                if len(v15) >= 2 {
                    v15 = v15[1:len(v15)]
                }
            case "a":
                for fi23 := 0; fi23 < 4; fi23++ {
                    for fi24 := 0; fi24 < 10; fi24++ {
                        if ((42 > 828) || true) {
                            if false {
                                v3 = 2
                            }
                        }
                        var v25 uint8 = v3
                        _ = fi24
                        _ = v25
                    }
                    for fi26 := 0; fi26 < 4; fi26++ {
                        switch (-(fi23)) {
                        case 4294967295:
                            var v27 uint32 = 1
                            _ = v27
                        case -9223372036854775808:
                            v1 = (v1 / ((v1) | 1))
                            var v28 int = (-(fi26))
                            switch "a" {
                            case "":
                                delete(v2, (v28 | fi26))
                            }
                            _ = v28
                        default:
                            if true {
                                if ((float64(1.5) - 1e10) >= v4) {
                                    var v29 bool = (!true)
                                    var v30 uint16 = (^(v1))
                                    _ = v29
                                    _ = v30
                                }
                            }
                        }
                        _ = fi26
                    }
                    _ = fi23
                }
            case "hello":
                v31 := GenS4{F0: 95, F1: 2147483648}
                _ = v31
            }
        default:
            v32 := GenS5{F0: 1}
            _ = v32
        }
        _ = v22
    } else {
        v33 := GenS6{F0: 369, F1: ("foo bar" + "hello")}
        for fi34 := 0; fi34 < 15; fi34++ {
            if true {
                if false {
                    v4 -= (((float64(0.0) / v4) * (v4 / v4)) / v4)
                    switch "\n\t" {
                    case "hello":
                        if (42 < 704) {
                            v35 := GenS7{F0: -1, F1: 18446744073709551615, F2: -694}
                            _ = v35
                        } else {
                            if (!true) {
                                _ = len(v15)
                            } else {
                                var v36 int8 = -41
                                if true {
                                    v1 += (v1 * (v1 >> ((255) & 63)))
                                    var v37 int32 = 2147483647
                                    _ = v37
                                }
                                var v38 string = ("\n\t" + ("\n\t" + "\n\t"))
                                _ = v36
                                _ = v38
                            }
                            var v39 int64 = 42
                            v3 -= uint8(v6)
                            _ = v39
                        }
                    case "":
                        fmt.Println(v3, fi34, v3)
                    }
                    var v40 int64 = 183
                    _ = v40
                }
                v6 -= (v6 | v6)
                var v41 int64 = 1
                _ = v41
            }
            _ = fi34
        }
        v42 := GenS1{F0: 1000, F1: 42, F2: 4294967295}
        _ = v33
        _ = v42
    }
    v3 += v3
    fmt.Println(v6, v6, v6)
    var v43 float32 = 0.0
    var v44 int8 = 0
    var v45 uint32 = 42
    v1 -= (^(v1))
    var v46 string = "\n\t"
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v46
    _ = v5
    _ = v6
}
