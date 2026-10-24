package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint64; F1 uint32; F2 bool }
type GenS3 struct { F0 bool; F1 int32; F2 uint }
func genFn0() int {
    var v7 uint8 = 127
    if (false && false) {
        v7 = v7
        v8 := []int{32767}
        switch "foo bar" {
        case "hello":
            var v9 uint64 = 569
            _ = v9
        case "a":
            var v10 int64 = 9223372036854775807
            _ = v10
        default:
            if true {
                var v11 uint32 = 608
                for fi12 := 0; fi12 < 10; fi12++ {
                    v13 := GenS1{F0: -32768}
                    _ = fi12
                    _ = v13
                }
                if len(v8) >= 2 {
                    v8 = v8[1:len(v8)]
                }
                _ = v11
            } else {
                for fi14 := 0; fi14 < 4; fi14++ {
                    var v15 int16 = 42
                    var v16 uint = 42
                    _ = fi14
                    _ = v15
                    _ = v16
                }
            }
        }
        _ = v8
    }
    for fi17 := 0; fi17 < 20; fi17++ {
        fmt.Println(fi17, fi17)
        _ = fi17
    }
    v7 += (v7 >> ((212) & 63))
    if (!true) {
        if (9223372036854775807 > -1) {
            var v18 uint8 = v7
            v7 -= v18
            _ = v18
        } else {
            switch 9223372036854775807 {
            case -495:
                for fi19 := 0; fi19 < 13; fi19++ {
                    v20 := []int{-184, fi19, (fi19 | -9223372036854775808), -2147483648}
                    _ = fi19
                    _ = v20
                }
            }
            v21 := []int{1}
            _ = v21
        }
        v7 = (^(v7))
        switch 32768 {
        case 42:
            v7 += v7
        case -9223372036854775808:
            v22 := []int{}
            _ = v22
        case 2147483647:
            for fi23 := 0; fi23 < 6; fi23++ {
                v7 += (^(v7))
                _ = fi23
            }
        default:
            for fi24 := 0; fi24 < 16; fi24++ {
                v25 := []int{fi24, (fi24 << ((767) & 63)), fi24, (fi24 - fi24)}
                _ = fi24
                _ = v25
            }
        }
    }
    _ = v7
    return 2147483647
}
func genFn1() int16 {
    var v26 uint32 = 875
    fmt.Println(v26)
    v27 := []int{4294967295, 128}
    if (127 >= -128) {
        var v28 int16 = 32767
        fmt.Println(v26, v26)
        v29 := GenS2{F0: 18446744073709551615, F1: 1, F2: false}
        _ = v28
        _ = v29
    } else {
        for fi30 := 0; fi30 < 3; fi30++ {
            for fi31 := 0; fi31 < 9; fi31++ {
                v26 = (v26 / ((859) | 1))
                var v32 uint32 = (v26 / (((^(v26))) | 1))
                fmt.Println(fi31, v26, fi31)
                _ = fi31
                _ = v32
            }
            _ = fi30
        }
        switch 127 {
        case 32768:
            v26 -= (^(v26))
        case -306:
            _ = len(v27)
        case 9223372036854775807:
            var v33 float64 = (1.5 / (1e10 * 0.1))
            _ = v33
        }
    }
    fmt.Println(v26)
    _ = v26
    _ = v27
    return -594
}

func main() {
    var v1 int64 = -9223372036854775808
    var v2 int16 = -1
    v3 := make(map[int]uint32)
    v4 := make(map[string]int)
    var v5 int32 = 0
    var v6 int64 = (v1 % ((v1) | 1))
    fmt.Println(v5, v5, v1)
    for fi34 := 0; fi34 < 5; fi34++ {
        v35 := v4[""]
        _ = fi34
        _ = v35
    }
    v36 := GenS1{F0: v2}
    for fi37 := 0; fi37 < 10; fi37++ {
        v6 = (v1 << ((42) & 63))
        var v38 uint = 2
        var v39 int64 = (v1 << (((v38 + v38)) & 63))
        _ = fi37
        _ = v38
        _ = v39
    }
    v40 := GenS2{F0: 0, F1: 42, F2: false}
    fmt.Println(v5, v5, v6)
    if false {
        v41 := []int{genFn0()}
        _ = v41
    } else {
        v5 += v5
    }
    v42 := []int{587, 32768, 32767, 32768}
    for fi43 := 0; fi43 < 17; fi43++ {
        var v44 float64 = (float64(float32((((1e10 / float64(((1.5 / -2.25) - ((((0.1 - (-2.25 / 1.5)) + 3.14159) * 1e10) / (1.5 + (((float32((0.0 / 3.14159)) + 0.0) / (-2.25 / float32((3.14159 / -2.25)))) - 3.14159)))))) * 0.1) - -2.25))) * (-2.25 * 1e10))
        switch fi43 {
        case 32768:
            v1 += v6
        }
        fmt.Println(v1, v44)
        _ = fi43
        _ = v44
    }
    v2 -= -32768
    if len(v42) >= 2 {
        v42 = v42[1:len(v42)]
    }
    if false {
        switch 42 {
        case 127:
            if ((2 != 0) && true) {
                v2 -= (^(v2))
                if len(v42) >= 2 {
                    v42 = v42[1:len(v42)]
                }
            }
        case -157:
            v1 = v1
        case 826:
            if true {
                if true {
                    fmt.Println(v1, v1, v5)
                } else {
                    v6 += v6
                    v45 := GenS1{F0: v2}
                    for fi46 := 0; fi46 < 4; fi46++ {
                        v5 = -2147483648
                        _ = fi46
                    }
                    _ = v45
                }
            }
        default:
            if (("\n\t" >= "") && true) {
                if (65535 == 65535) {
                    var v47 float32 = 3.14159
                    if true {
                        for fi48 := 0; fi48 < 19; fi48++ {
                            var v49 int32 = v5
                            _ = fi48
                            _ = v49
                        }
                        v1 = v1
                    }
                    _ = v47
                }
                v6 += v6
                switch "foo bar" {
                case "":
                    _ = len(v42)
                case "\n\t":
                    fmt.Println(v5, v2, v6)
                }
            } else {
                var v50 int8 = -84
                for fi51 := 0; fi51 < 1; fi51++ {
                    v52 := v3[-9223372036854775808]
                    v53 := GenS3{F0: (-710 <= int64(v2)), F1: (v5 >> ((42) & 63)), F2: 827}
                    _ = fi51
                    _ = v52
                    _ = v53
                }
                var v54 uint = 4294967295
                _ = v50
                _ = v54
            }
        }
        var v55 int16 = (^(v2))
        _ = len(v4)
        _ = v55
    } else {
        for fi56 := 0; fi56 < 12; fi56++ {
            var v57 string = ""
            _ = fi56
            _ = v57
        }
    }
    v5 -= (v5 << ((665) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v42:", len(v42), v42)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v40
    _ = v42
    _ = v5
    _ = v6
}
