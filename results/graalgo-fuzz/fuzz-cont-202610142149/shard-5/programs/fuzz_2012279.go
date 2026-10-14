package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint64; F1 int32; F2 uint64 }
type GenS5 struct { F0 string; F1 bool }
type GenS6 struct { F0 int16; F1 int16; F2 int64 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 bool; F1 int16; F2 uint }
func genFn0(p0 int32, p1 int32, p2 bool) uint16 {
    v5 := []int{-9223372036854775808, 923, 4294967295}
    if p2 {
        var v6 int64 = 9223372036854775807
        fmt.Println(p2)
        p1 = -1
        _ = v6
    } else {
        for fi7 := 0; fi7 < 15; fi7++ {
            if p2 {
                v8 := GenS1{F0: -128, F1: (fi7 * (^(fi7)))}
                var v9 uint = 65535
                v10 := GenS2{F0: 127}
                _ = v10
                _ = v8
                _ = v9
            } else {
                switch fi7 {
                case -2147483648:
                    for fi11 := 0; fi11 < 3; fi11++ {
                        if len(v5) > 0 {
                            v5[0] = 0
                        }
                        if p2 {
                            p1 = 0
                            if p2 {
                                if p2 {
                                    v5 = append(v5, (fi11 | fi7))
                                    if (0 <= 0) {
                                        var v12 int16 = 32767
                                        _ = v12
                                    }
                                } else {
                                    if len(v5) > 0 {
                                        v5[0] = (fi11 | (fi7 | (fi7 % (((fi11 & fi11)) | 1))))
                                    }
                                    p2 = true
                                }
                                var v13 uint32 = 2
                                _ = v13
                            } else {
                                v14 := GenS2{F0: -128}
                                _ = v14
                            }
                        }
                        switch ("foo bar" + "") {
                        case "a":
                            if (((-2.25 * (1.5 / 3.14159)) + 0.1) >= 1.5) {
                                v15 := GenS3{F0: 42}
                                fmt.Println(fi7, fi7)
                                var v16 uint64 = 1
                                _ = v15
                                _ = v16
                            } else {
                                var v17 uint32 = 731
                                _ = len(v5)
                                switch (fi7 - (fi7 >> ((2) & 63))) {
                                case 1:
                                    v17 += (v17 << ((242) & 63))
                                }
                                var v18 string = "hello"
                                _ = v17
                                _ = v18
                            }
                        }
                        var v19 uint32 = 923
                        _ = fi11
                        _ = v19
                    }
                case 429:
                    var v20 string = "foo bar"
                    _ = v20
                }
            }
            _ = fi7
        }
        switch "a" {
        case "":
            var v21 uint32 = 1
            _ = v21
        case "hello":
            for fi22 := 0; fi22 < 3; fi22++ {
                switch fi22 {
                case 0:
                    v5 = append(v5, fi22)
                }
                if (9223372036854775807 > 0) {
                    p2 = (0 > 775)
                    for fi23 := 0; fi23 < 5; fi23++ {
                        _ = len(v5)
                        var v24 int32 = p0
                        _ = fi23
                        _ = v24
                    }
                }
                switch fi22 {
                case 0:
                    switch "\n\t" {
                    case "foo bar":
                        if ((!true) && p2) {
                            var v25 int32 = (p1 / ((p0) | 1))
                            p0 -= 0
                            _ = v25
                        }
                    case "\n\t":
                        v26 := GenS4{F0: 501, F1: (p1 & p1), F2: 42}
                        _ = v26
                    }
                default:
                    v27 := GenS2{F0: -1}
                    _ = v27
                }
                _ = fi22
            }
        }
        v28 := GenS5{F0: "a", F1: (128 < -1)}
        _ = v28
    }
    for fi29 := 0; fi29 < 10; fi29++ {
        p2 = (fi29 >= (fi29 << ((65535) & 63)))
        _ = fi29
    }
    switch "\n\t" {
    case "\n\t":
        v30 := GenS1{F0: 1, F1: -2147483648}
        _ = v30
    case "foo bar":
        p0 -= (p0 * p1)
    case "a":
        switch "" {
        case "a":
            p0 = (p0 & (p1 - (-(p1))))
        case "\n\t":
            fmt.Println(p0, p2, p0)
        }
    }
    v31 := GenS3{F0: 9223372036854775807}
    _ = p0
    _ = p1
    _ = p2
    _ = v31
    _ = v5
    return 530
}

func main() {
    v1 := make(map[uint32]int)
    var v2 float64 = 0.0
    var v3 int8 = 1
    var v4 float64 = (float64(-2.25) + 1e10)
    for fi32 := 0; fi32 < 20; fi32++ {
        v33 := GenS6{F0: 1000, F1: -136, F2: -9223372036854775808}
        _ = len(v1)
        _ = fi32
        _ = v33
    }
    v34 := v1[1]
    v2 = v2
    switch "foo bar" {
    case "":
        for fi35 := 0; fi35 < 4; fi35++ {
            v3 += 100
            _ = fi35
        }
    case "a":
        var v36 int16 = 923
        _ = v36
    default:
        v37 := v1[2]
        _ = v37
    }
    if true {
        var v38 uint64 = 4294967295
        v1[42] = v34
        _ = v38
    }
    if (1 <= -958) {
        for fi39 := 0; fi39 < 6; fi39++ {
            v40 := GenS7{F0: 42}
            v34 -= (fi39 + 98)
            _ = fi39
            _ = v40
        }
        var v41 bool = false
        v3 -= (-(v3))
        _ = v41
    }
    fmt.Println(v3, v4, v34)
    v2 = ((v2 - v4) / v4)
    v42 := v1[40]
    v43 := GenS8{F0: (true || (!false)), F1: 1, F2: 0}
    var v44 uint64 = 0
    for fi45 := 0; fi45 < 6; fi45++ {
        delete(v1, 4294967295)
        v3 = (v3 | v3)
        for fi46 := 0; fi46 < 16; fi46++ {
            v4 += v2
            _ = fi46
        }
        _ = fi45
    }
    if true {
        v34 += 793
    } else {
        v4 = (-2.25 / float64(0.1))
        switch ("" + "hello") {
        case "foo bar":
            delete(v1, 2147483648)
        case "a":
            v42 += 128
        case "\n\t":
            v3 += (-(v3))
        default:
            for fi47 := 0; fi47 < 18; fi47++ {
                var v48 bool = true
                _ = fi47
                _ = v48
            }
        }
    }
    switch (v34 * v34) {
    case 4294967295:
        fmt.Println(v4)
    case -234:
        v34 += (-(v42))
    }
    var v49 float64 = float64(3.14159)
    if false {
        var v50 float32 = 1.5
        _ = v50
    } else {
        for fi51 := 0; fi51 < 20; fi51++ {
            fmt.Println(v2)
            _ = fi51
        }
        var v52 int32 = 633
        _ = v52
    }
    var v53 uint32 = 546
    v49 = (v49 - (v49 - v4))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v44:", v44)
    fmt.Println("v49:", v49)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v49
    _ = v53
}
