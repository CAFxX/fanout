package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint; F2 bool }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint; F1 int8; F2 int8 }
func genFn0(p0 string, p1 int16, p2 int8) int32 {
    var v7 int64 = 0
    for fi8 := 0; fi8 < 19; fi8++ {
        fmt.Println(p2, p0, p2)
        switch p0 {
        case "foo bar":
            v9 := GenS1{F0: 2, F1: 823, F2: false}
            _ = v9
        case "":
            for fi10 := 0; fi10 < 16; fi10++ {
                switch fi10 {
                case 127:
                    v11 := GenS1{F0: 2, F1: 4294967295, F2: false}
                    v7 -= int64(p2)
                    p1 -= (p1 / (((p1 | (-(p1)))) | 1))
                    v12 := []int{(fi8 >> ((95) & 63)), fi8, (fi8 / ((0) | 1))}
                    _ = v11
                    _ = v12
                }
                _ = fi10
            }
        case "a":
            var v13 int = (^(fi8))
            _ = v13
        default:
            v14 := []int{}
            _ = v14
        }
        for fi15 := 0; fi15 < 17; fi15++ {
            var v16 uint64 = 42
            var v17 float64 = (0.0 + (0.0 + float64(0.0)))
            var v18 int64 = 684
            _ = fi15
            _ = v16
            _ = v17
            _ = v18
        }
        _ = fi8
    }
    v7 += (v7 | v7)
    v19 := []int{0, 473, -9223372036854775808, 865}
    var v20 int = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v19
    _ = v20
    _ = v7
    return 0
}

func main() {
    var v1 float64 = 0.0
    var v2 int64 = 842
    v3 := []uint32{0}
    v4 := make(map[int]int)
    var v5 int32 = 980
    var v6 float64 = 3.14159
    var v21 int16 = 42
    var v22 uint8 = 193
    for fi23 := 0; fi23 < 1; fi23++ {
        v24 := GenS1{F0: 0, F1: 834, F2: false}
        v25 := GenS1{F0: 123, F1: 0, F2: (872 != 413)}
        _ = fi23
        _ = v24
        _ = v25
    }
    v5 += genFn0("foo bar", v21, 99)
    fmt.Println(v5)
    if (42 == 425) {
        v3 = append(v3, 12)
        var v26 int32 = (v5 * 670)
        _ = v26
    }
    for fi27 := 0; fi27 < 2; fi27++ {
        for fi28 := 0; fi28 < 18; fi28++ {
            if (42 >= 0) {
                fmt.Println(v6, v22)
                v2 += (v2 + (v2 % ((v2) | 1)))
                var v29 uint8 = v22
                _ = v29
            }
            v2 += v2
            var v30 bool = true
            var v31 int64 = -221
            var v32 int8 = -128
            _ = fi28
            _ = v30
            _ = v31
            _ = v32
        }
        fmt.Println(v6, fi27, fi27)
        _ = fi27
    }
    for fi33 := 0; fi33 < 5; fi33++ {
        v34 := GenS2{F0: 415}
        for fi35 := 0; fi35 < 11; fi35++ {
            var v36 uint32 = 65535
            var v37 float32 = (1e10 - 0.0)
            v4[(fi35 ^ fi33)] = fi35
            _ = fi35
            _ = v36
            _ = v37
        }
        if (("hello" + "\n\t") >= ("" + "\n\t")) {
            fmt.Println(v6, fi33)
            v38 := GenS3{F0: 393, F1: 42, F2: -128}
            var v39 int16 = 32767
            _ = v38
            _ = v39
        }
        _ = fi33
        _ = v34
    }
    v1 = v1
    switch 32768 {
    case 4294967295:
        if false {
            v21 -= v21
            var v40 float32 = -2.25
            if (v5 <= v5) {
                if ((!(!(42 != 1))) && (false || true)) {
                    v41 := GenS1{F0: 216, F1: 18446744073709551615, F2: ((true || (int(v21) == 42)) || false)}
                    switch "\n\t" {
                    case "\n\t":
                        for fi42 := 0; fi42 < 5; fi42++ {
                            v5 = v5
                            var v43 uint8 = v22
                            _ = fi42
                            _ = v43
                        }
                    case "hello":
                        switch -943 {
                        case -2147483648:
                            v44 := GenS2{F0: 42}
                            _ = v44
                        case -9223372036854775808:
                            v2 += (v2 >> ((65535) & 63))
                        default:
                            switch ("hello" + "foo bar") {
                            case "hello":
                                for fi45 := 0; fi45 < 19; fi45++ {
                                    v21 -= v21
                                    v5 -= (v5 % ((19) | 1))
                                    _ = fi45
                                }
                            case "\n\t":
                                fmt.Println(v5)
                            case "foo bar":
                                _ = len(v3)
                            }
                        }
                    default:
                        var v46 string = "a"
                        _ = v46
                    }
                    var v47 uint = 0
                    _ = v41
                    _ = v47
                } else {
                    fmt.Println(v22, v1)
                    var v48 uint16 = 32767
                    _ = v48
                }
                v3 = append(v3, 4294967295)
                delete(v4, 128)
            }
            _ = v40
        }
    case 1:
        var v49 int64 = v2
        _ = v49
    case 32767:
        var v50 float32 = (((float32(v1) + 1e10) * ((3.14159 * 1e10) * (0.0 * (float32(v6) * 0.1)))) - 3.14159)
        _ = v50
    }
    v5 -= v5
    if false {
        var v51 int64 = (v2 | v2)
        v4[int(v2)] = 708
        _ = v51
    }
    v22 -= (^(v22))
    v21 -= 0
    var v52 uint8 = 255
    if (int64(v21) >= 42) {
        for fi53 := 0; fi53 < 6; fi53++ {
            if true {
                if (42 > 463) {
                    v1 = ((v1 - v6) * 0.0)
                } else {
                    fmt.Println(v6)
                    v21 -= (v21 | (v21 / ((766) | 1)))
                    fmt.Println(v2)
                }
            }
            var v54 uint8 = v52
            _ = fi53
            _ = v54
        }
        v6 = (v1 + v6)
        for fi55 := 0; fi55 < 9; fi55++ {
            v22 += v22
            _ = fi55
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v6
}
