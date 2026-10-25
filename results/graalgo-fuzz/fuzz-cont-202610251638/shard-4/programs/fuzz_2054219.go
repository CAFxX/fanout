package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8; F2 uint }
type GenS2 struct { F0 uint; F1 string; F2 uint }
func genFn0(p0 int16, p1 int, p2 string) uint8 {
    fmt.Println(p2)
    if (42 > 100) {
        for fi6 := 0; fi6 < 4; fi6++ {
            switch (fi6 * p1) {
            case 42:
                var v7 uint8 = 79
                _ = v7
            case 9223372036854775807:
                for fi8 := 0; fi8 < 3; fi8++ {
                    v9 := []int{fi8, fi6, p1, -552}
                    var v10 bool = (-62 != p0)
                    if true {
                        switch (p1 << ((255) & 63)) {
                        case -9223372036854775808:
                            p0 -= (p0 % ((1000) | 1))
                        default:
                            v10 = true
                        }
                        v11 := GenS1{F0: 884, F1: 127, F2: 785}
                        _ = v11
                    }
                    if len(v9) > 0 {
                        v9[0] = (fi6 % ((fi8) | 1))
                    }
                    _ = fi8
                    _ = v10
                    _ = v9
                }
            default:
                v12 := []int{(fi6 - (fi6 | (fi6 & 1))), (fi6 * (-(fi6))), 42, p1}
                _ = v12
            }
            fmt.Println(p2)
            p2 = p2
            _ = fi6
        }
        switch ("foo bar" + p2) {
        case "hello":
            fmt.Println(p1, p1)
        default:
            v13 := []int{p1}
            _ = v13
        }
        switch "foo bar" {
        case "foo bar":
            var v14 int8 = -128
            _ = v14
        case "hello":
            if (-971 > -2147483648) {
                p1 = (p1 & (p1 * (p1 >> ((42) & 63))))
            } else {
                var v15 uint32 = 2147483647
                _ = v15
            }
        case "":
            v16 := GenS2{F0: 2, F1: ("foo bar" + p2), F2: 29}
            _ = v16
        default:
            for fi17 := 0; fi17 < 18; fi17++ {
                v18 := []int{}
                _ = fi17
                _ = v18
            }
        }
    }
    var v19 float64 = 0.1
    v20 := GenS1{F0: 32767, F1: -25, F2: 255}
    _ = p0
    _ = p1
    _ = p2
    _ = v19
    _ = v20
    return 42
}

func main() {
    v1 := make(map[int]uint32)
    v2 := []uint32{1, 42, 575, 299}
    var v3 uint32 = 4294967295
    v4 := []int{127, 84}
    var v5 uint = 255
    delete(v1, -227)
    var v21 string = ("hello" + ("" + ("" + "a")))
    if true {
        var v22 uint8 = 127
        var v23 int64 = 42
        v24 := GenS2{F0: 296, F1: "hello", F2: 866}
        _ = v22
        _ = v23
        _ = v24
    } else {
        var v25 uint16 = 1
        v3 += v3
        _ = v25
    }
    _ = len(v1)
    for fi26 := 0; fi26 < 1; fi26++ {
        if (260 >= 255) {
            v27 := GenS1{F0: 172, F1: 42, F2: (v5 >> (((v5 ^ 2)) & 63))}
            var v28 string = v21
            if (!("hello" >= (v21 + "hello"))) {
                if false {
                    v3 += (v3 & v3)
                    v3 -= (^(v3))
                    v21 = v28
                } else {
                    if true {
                        var v29 int16 = -208
                        _ = v29
                    }
                    switch fi26 {
                    case -174:
                        for fi30 := 0; fi30 < 4; fi30++ {
                            if false {
                                var v31 int8 = 0
                                _ = v31
                            } else {
                                v21 = v28
                            }
                            v3 += (v3 - (v3 - 64))
                            var v32 int32 = 510
                            v5 += v5
                            fmt.Println(fi26, v21)
                            _ = fi30
                            _ = v32
                        }
                    }
                }
                v5 += (v5 + (v5 ^ 402))
            } else {
                fmt.Println(v5, v5, v5)
            }
            _ = v27
            _ = v28
        }
        switch "a" {
        case "foo bar":
            fmt.Println(v21)
        case "":
            if (1000 < 0) {
                if ((v5 + 18446744073709551615) > v5) {
                    v4 = append(v4, (-(fi26)))
                    v33 := v1[(fi26 ^ 166)]
                    v34 := GenS2{F0: (^(v5)), F1: "", F2: v5}
                    _ = v33
                    _ = v34
                } else {
                    for fi35 := 0; fi35 < 20; fi35++ {
                        v3 += (^(v3))
                        _ = fi35
                    }
                    fmt.Println(v5)
                }
                v5 = (v5 << (((v5 + v5)) & 63))
            } else {
                v21 = "hello"
            }
        }
        for fi36 := 0; fi36 < 13; fi36++ {
            v37 := v1[(fi26 / ((2147483647) | 1))]
            if (-1 != -503) {
                var v38 uint = v5
                v39 := GenS1{F0: 443, F1: 100, F2: (v38 * (v5 ^ v38))}
                _ = v38
                _ = v39
            }
            v5 += (v5 & v5)
            if false {
                var v40 uint64 = 272
                _ = v40
            }
            _ = fi36
            _ = v37
        }
        _ = fi26
    }
    if false {
        fmt.Println(v5, v5, v3)
    } else {
        var v41 float64 = 1.5
        v1[32768] = 65535
        v42 := GenS1{F0: 65535, F1: 42, F2: v5}
        _ = v41
        _ = v42
    }
    for fi43 := 0; fi43 < 16; fi43++ {
        _ = len(v2)
        for fi44 := 0; fi44 < 14; fi44++ {
            v3 -= v3
            _ = fi44
        }
        switch fi43 {
        case -1:
            if true {
                var v45 int64 = -647
                switch 32767 {
                case 32768:
                    switch int(v45) {
                    case -9223372036854775808:
                        var v46 uint64 = 2
                        _ = v46
                    }
                }
                _ = v45
            }
        case 651:
            var v47 string = (((("\n\t" + ((v21 + v21) + v21)) + (v21 + v21)) + "foo bar") + v21)
            _ = v47
        case 42:
            var v48 uint16 = 0
            _ = v48
        }
        _ = fi43
    }
    v21 = "hello"
    if (14 > 100) {
        switch v21 {
        case "a":
            var v49 uint32 = (v3 + v3)
            _ = v49
        case "hello":
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        default:
            var v50 int64 = -9223372036854775808
            _ = v50
        }
        fmt.Println(v5, v21, v5)
        v5 = v5
    } else {
        if false {
            _ = len(v1)
            switch "hello" {
            case "foo bar":
                for fi51 := 0; fi51 < 14; fi51++ {
                    var v52 int16 = 1
                    var v53 float64 = 3.14159
                    v5 -= (v5 + v5)
                    _ = fi51
                    _ = v52
                    _ = v53
                }
            case "hello":
                v3 -= v3
            default:
                fmt.Println(v21)
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
