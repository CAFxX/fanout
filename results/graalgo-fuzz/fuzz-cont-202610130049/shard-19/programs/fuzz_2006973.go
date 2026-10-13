package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16 }
type GenS2 struct { F0 int16; F1 int32 }
type GenS3 struct { F0 uint; F1 int }
type GenS4 struct { F0 int16; F1 int8 }
type GenS5 struct { F0 int; F1 int }
func genFn0(p0 int64, p1 int8, p2 int8) uint64 {
    switch -9223372036854775808 {
    case 127:
        var v4 uint8 = 128
        _ = v4
    case 128:
        var v5 string = ("a" + "\n\t")
        _ = v5
    case -19:
        if true {
            if false {
                fmt.Println(p2)
            } else {
                p2 += p2
                p1 = (p1 + 1)
            }
            var v6 int = 42
            switch -980 {
            case 538:
                var v7 string = "hello"
                _ = v7
            }
            _ = v6
        }
    }
    p0 -= p0
    _ = p0
    _ = p1
    _ = p2
    return 329
}
func genFn1(p0 int, p1 int32, p2 int16) uint16 {
    for fi8 := 0; fi8 < 2; fi8++ {
        v9 := []int{}
        var v10 int64 = 42
        var v11 int32 = p1
        _ = fi8
        _ = v10
        _ = v11
        _ = v9
    }
    var v12 uint16 = 32767
    p2 += 1
    v13 := []int{}
    v14 := GenS1{F0: 282, F1: (v12 << ((4294967295) & 63))}
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v13
    _ = v14
    return (v12 * v12)
}

func main() {
    var v1 int8 = -1
    v2 := make(map[int]bool)
    var v3 float64 = 3.14159
    fmt.Println(v1)
    if false {
        fmt.Println(v1)
        v2[127] = true
        if (2 > 2) {
            var v15 uint16 = 2
            v2[88] = false
            _ = v15
        } else {
            switch "foo bar" {
            case "a":
                fmt.Println(v3, v3)
            case "\n\t":
                var v16 string = "\n\t"
                _ = v16
            case "foo bar":
                if true {
                    v17 := []int{}
                    fmt.Println(v1, v3)
                    v18 := GenS1{F0: 255, F1: 487}
                    _ = v17
                    _ = v18
                } else {
                    switch 0 {
                    case 4294967295:
                        switch "" {
                        case "hello":
                            fmt.Println(v1)
                        case "a":
                            v19 := []int{0}
                            _ = v19
                        }
                    }
                    v20 := GenS1{F0: 42, F1: 42}
                    var v21 int = 128
                    _ = v20
                    _ = v21
                }
            }
            var v22 uint = 255
            var v23 uint64 = 771
            _ = v22
            _ = v23
        }
    }
    var v24 float64 = 0.1
    v24 += v24
    v25 := v2[2147483647]
    v26 := GenS2{F0: 1, F1: 42}
    var v27 bool = false
    fmt.Println(v1, v1, v24)
    switch "\n\t" {
    case "":
        fmt.Println(v24)
    case "hello":
        if v27 {
            switch "" {
            case "foo bar":
                if (220 > 2) {
                    fmt.Println(v1, v3, v27)
                    var v28 int64 = -1
                    _ = v28
                } else {
                    v24 = (1.5 / 0.1)
                    v27 = (1 < 2147483647)
                }
            case "a":
                if false {
                    if (2 >= 345) {
                        v24 += 1.5
                        v29 := []int{128}
                        _ = v29
                    } else {
                        v27 = (32767 == 0)
                        v30 := GenS1{F0: 4294967295, F1: 2}
                        _ = len(v2)
                        _ = v30
                    }
                } else {
                    var v31 int64 = 9223372036854775807
                    _ = v31
                }
            case "":
                var v32 int8 = (-(v1))
                _ = v32
            default:
                switch "a" {
                case "foo bar":
                    var v33 int32 = 146
                    _ = v33
                }
            }
            for fi34 := 0; fi34 < 12; fi34++ {
                switch "a" {
                case "a":
                    v35 := []int{fi34}
                    _ = v35
                case "hello":
                    var v36 uint = 788
                    _ = v36
                }
                var v37 uint64 = 0
                var v38 uint = 18446744073709551615
                _ = fi34
                _ = v37
                _ = v38
            }
        } else {
            v3 = (v3 * 0.1)
            if v27 {
                var v39 int8 = 42
                for fi40 := 0; fi40 < 15; fi40++ {
                    v41 := GenS3{F0: 18446744073709551615, F1: (fi40 ^ fi40)}
                    _ = fi40
                    _ = v41
                }
                var v42 int8 = (v1 * (v39 | v1))
                _ = v39
                _ = v42
            } else {
                v43 := GenS4{F0: 409, F1: (v1 % ((-46) | 1))}
                _ = v43
            }
            if v25 {
                if true {
                    for fi44 := 0; fi44 < 7; fi44++ {
                        for fi45 := 0; fi45 < 8; fi45++ {
                            if v25 {
                                if true {
                                    v46 := GenS5{F0: (^(fi44)), F1: fi45}
                                    v47 := GenS4{F0: int16(fi44), F1: (^(v1))}
                                    _ = v46
                                    _ = v47
                                } else {
                                    v48 := GenS3{F0: 4294967295, F1: (fi44 << ((4294967295) & 63))}
                                    _ = v48
                                }
                                if (!(-1 != 32767)) {
                                    v27 = v25
                                    fmt.Println(v1, fi45, v1)
                                    var v49 int8 = (v1 % (((v1 / ((v1) | 1))) | 1))
                                    _ = v49
                                } else {
                                    fmt.Println(v3)
                                }
                            }
                            _ = len(v2)
                            v24 -= (v3 - v3)
                            _ = fi45
                        }
                        var v50 bool = (true && v25)
                        _ = fi44
                        _ = v50
                    }
                    for fi51 := 0; fi51 < 5; fi51++ {
                        if (!(0 == int64(v1))) {
                            var v52 int64 = 9223372036854775807
                            _ = v52
                        } else {
                            fmt.Println(v24, v3, v1)
                        }
                        switch "a" {
                        case "a":
                            v24 += 0.0
                        case "hello":
                            v27 = ((4294967295 != 42) && v25)
                        case "foo bar":
                            for fi53 := 0; fi53 < 14; fi53++ {
                                var v54 float32 = (0.0 - (0.1 - (1e10 + 1.5)))
                                _ = fi53
                                _ = v54
                            }
                        }
                        _ = fi51
                    }
                }
            } else {
                for fi55 := 0; fi55 < 2; fi55++ {
                    v56 := []int{fi55, (fi55 | (^(fi55))), fi55}
                    _ = fi55
                    _ = v56
                }
            }
        }
    case "foo bar":
        v24 -= v3
    default:
        v24 = v24
    }
    if v25 {
        if (-238 == 0) {
            switch "\n\t" {
            case "\n\t":
                v24 = v24
            case "":
                var v57 int32 = 2147483647
                _ = v57
            }
            var v58 uint64 = 2
            var v59 float32 = (0.1 + 0.0)
            _ = v58
            _ = v59
        }
        var v60 uint32 = 777
        var v61 int32 = -1
        _ = v60
        _ = v61
    }
    for fi62 := 0; fi62 < 5; fi62++ {
        var v63 int16 = 708
        _ = fi62
        _ = v63
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
}
