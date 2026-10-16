package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64; F2 uint16 }
type GenS2 struct { F0 int16; F1 uint32 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint32; F1 uint64; F2 uint32 }
type GenS5 struct { F0 int8; F1 bool }
func genFn0(p0 uint64, p1 int32, p2 int32) uint32 {
    var v6 int8 = -8
    v6 = v6
    var v7 uint = 0
    if false {
        switch 0 {
        case 32767:
            v8 := GenS1{F0: 1, F1: 9223372036854775807, F2: 65535}
            _ = v8
        }
        fmt.Println(v7)
        var v9 int64 = 42
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 614
}
func genFn1(p0 uint64, p1 bool) bool {
    p0 -= (p0 >> ((0) & 63))
    p0 = (p0 % ((2) | 1))
    var v10 int8 = 42
    p0 = 4294967295
    v11 := GenS1{F0: 128, F1: 1, F2: 377}
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return false
}

func main() {
    v1 := []int{-1, 32768}
    var v2 string = "hello"
    var v3 int64 = 883
    v4 := make(map[int]int)
    v5 := make(map[uint32]uint32)
    var v12 float32 = -2.25
    fmt.Println(v12, v12)
    if false {
        switch -2147483648 {
        case -340:
            v13 := GenS1{F0: 42, F1: v3, F2: 2}
            _ = v13
        case -777:
            var v14 uint8 = 0
            _ = v14
        case 2147483647:
            if (!false) {
                if false {
                    v15 := GenS2{F0: 0, F1: 42}
                    v3 -= (v3 * (v3 | v3))
                    delete(v4, -1)
                    _ = v15
                }
                for fi16 := 0; fi16 < 17; fi16++ {
                    v17 := GenS1{F0: 68, F1: 795, F2: 1}
                    var v18 float64 = 0.0
                    var v19 uint8 = 1
                    _ = fi16
                    _ = v17
                    _ = v18
                    _ = v19
                }
            }
        default:
            v3 = (v3 & 0)
        }
        for fi20 := 0; fi20 < 5; fi20++ {
            v2 = v2
            fmt.Println(v2, v12)
            var v21 float32 = v12
            _ = fi20
            _ = v21
        }
    }
    v22 := GenS3{F0: 2}
    for fi23 := 0; fi23 < 3; fi23++ {
        v24 := GenS3{F0: 2}
        delete(v5, 4294967295)
        for fi25 := 0; fi25 < 18; fi25++ {
            switch (v2 + v2) {
            case "hello":
                fmt.Println(v12, fi25, v2)
            default:
                switch (fi25 ^ (fi25 / ((fi23) | 1))) {
                case 32767:
                    var v26 uint = 0
                    _ = v26
                case -2147483648:
                    v3 += v3
                }
            }
            _ = fi25
        }
        _ = fi23
        _ = v24
    }
    v12 -= (v12 / v12)
    v12 -= float32(3.14159)
    for fi27 := 0; fi27 < 8; fi27++ {
        if (!(-2147483648 <= 276)) {
            switch 42 {
            case 4294967295:
                switch "\n\t" {
                case "":
                    var v28 int8 = 60
                    _ = v28
                case "a":
                    _ = len(v1)
                case "foo bar":
                    v1 = append(v1, fi27)
                }
            case 1:
                for fi29 := 0; fi29 < 12; fi29++ {
                    v30 := GenS1{F0: 0, F1: (v3 << ((59) & 63)), F2: 42}
                    if (((float32(1.5) * v12) - v12) >= ((v12 * v12) / (1.5 + (float32(1.5) / (((0.0 / v12) - (float32(3.14159) - 0.1)) - 3.14159))))) {
                        var v31 bool = (1 != 586)
                        v2 = v2
                        _ = v31
                    } else {
                        var v32 bool = ((194 > 255) || true)
                        _ = v32
                    }
                    fmt.Println(fi29, fi29)
                    _ = fi29
                    _ = v30
                }
            case -919:
                fmt.Println(v12, fi27)
            default:
                switch v2 {
                case "":
                    var v33 uint = 4294967295
                    _ = v33
                case "\n\t":
                    if false {
                        v34 := v5[2147483647]
                        _ = v34
                    }
                case "hello":
                    switch fi27 {
                    case 9223372036854775807:
                        var v35 uint64 = 234
                        _ = v35
                    case 962:
                        v3 -= (v3 ^ (^(v3)))
                    }
                default:
                    var v36 uint32 = 2
                    _ = v36
                }
            }
            var v37 bool = true
            var v38 uint32 = 705
            _ = v37
            _ = v38
        }
        if true {
            for fi39 := 0; fi39 < 5; fi39++ {
                var v40 bool = false
                var v41 uint32 = genFn0(2, -1, -2147483648)
                switch fi39 {
                case -9223372036854775808:
                    v3 = (v3 << ((904) & 63))
                case -2147483648:
                    switch fi39 {
                    case 1:
                        v42 := GenS4{F0: v41, F1: 18446744073709551615, F2: 4294967295}
                        _ = v42
                    case 128:
                        delete(v5, (v41 & 65535))
                    case 32768:
                        fmt.Println(v12, v41, fi39)
                        switch ("foo bar" + "a") {
                        case "a":
                            if (565 == 18446744073709551615) {
                                var v43 int16 = -253
                                _ = v43
                            }
                        }
                        var v44 uint32 = v41
                        v45 := v4[(fi27 & fi39)]
                        _ = v44
                        _ = v45
                    }
                case 717:
                    v2 = v2
                default:
                    v46 := GenS5{F0: 127, F1: v40}
                    _ = v46
                }
                _ = fi39
                _ = v40
                _ = v41
            }
            _ = len(v1)
        } else {
            v12 -= v12
        }
        _ = fi27
    }
    for fi47 := 0; fi47 < 16; fi47++ {
        v48 := GenS4{F0: 2147483648, F1: 0, F2: 954}
        _ = fi47
        _ = v48
    }
    v49 := GenS3{F0: 834}
    switch v2 {
    case "hello":
        var v50 string = "\n\t"
        _ = v50
    case "\n\t":
        if (true || false) {
            fmt.Println(v2)
        }
    case "a":
        v51 := GenS4{F0: 495, F1: 403, F2: 693}
        _ = v51
    }
    delete(v5, 639)
    v3 = (v3 - 42)
    var v52 uint = 235
    var v53 uint8 = 0
    if len(v1) > 0 {
        v1[0] = 32767
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v12
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v52
    _ = v53
}
