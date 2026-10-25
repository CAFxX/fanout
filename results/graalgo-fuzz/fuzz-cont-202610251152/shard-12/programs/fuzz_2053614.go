package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32; F2 uint8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint16; F1 bool }
type GenS6 struct { F0 int64; F1 int64; F2 int8 }
type GenS7 struct { F0 int; F1 bool }
type GenS8 struct { F0 int16; F1 int16 }
type GenS9 struct { F0 int; F1 int }
func genFn0(p0 string) uint16 {
    var v5 uint16 = 32767
    if false {
        v6 := GenS1{F0: 447, F1: 2, F2: 137}
        _ = v6
    }
    for fi7 := 0; fi7 < 8; fi7++ {
        v5 += (^(v5))
        if (1 <= 0) {
            for fi8 := 0; fi8 < 11; fi8++ {
                v9 := []int{(fi7 / (((fi8 & 641)) | 1))}
                _ = fi8
                _ = v9
            }
            v10 := GenS1{F0: 697, F1: 0, F2: 217}
            if ((!(((false || (42 != 127)) || (969 == 476)) && (3 > 128))) || false) {
                for fi11 := 0; fi11 < 9; fi11++ {
                    fmt.Println(v5, fi7, fi11)
                    _ = fi11
                }
                v12 := []int{fi7, fi7}
                _ = v12
            }
            _ = v10
        }
        _ = fi7
    }
    if (-128 < 0) {
        for fi13 := 0; fi13 < 2; fi13++ {
            v14 := GenS2{F0: 4294967295}
            for fi15 := 0; fi15 < 19; fi15++ {
                fmt.Println(fi15, fi13)
                fmt.Println(fi15)
                _ = fi15
            }
            _ = fi13
            _ = v14
        }
        v16 := []int{9223372036854775807, 0, -583}
        fmt.Println(v5, p0)
        _ = v16
    } else {
        switch -535 {
        case 42:
            v5 = (^(v5))
        case 4294967295:
            v17 := GenS1{F0: 2147483647, F1: 4294967295, F2: 127}
            _ = v17
        case -2147483648:
            v18 := GenS3{F0: 1}
            _ = v18
        default:
            p0 = "foo bar"
        }
        v19 := []int{-9223372036854775808, 688, 310}
        var v20 float64 = float64(0.1)
        _ = v19
        _ = v20
    }
    fmt.Println(p0, v5)
    _ = p0
    _ = v5
    return 255
}
func genFn1(p0 int16) uint32 {
    if (!false) {
        v21 := GenS4{F0: 0}
        switch ("" + "\n\t") {
        case "\n\t":
            fmt.Println(p0)
        case "a":
            var v22 bool = false
            _ = v22
        default:
            if true {
                p0 -= p0
                for fi23 := 0; fi23 < 9; fi23++ {
                    p0 -= (p0 >> ((65535) & 63))
                    var v24 float32 = (float32(0.1) * 0.0)
                    _ = fi23
                    _ = v24
                }
            }
        }
        p0 = p0
        _ = v21
    }
    v25 := GenS2{F0: 0}
    if true {
        fmt.Println(p0)
    }
    v26 := []int{9223372036854775807, -552}
    _ = p0
    _ = v25
    _ = v26
    return 77
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[string]uint32)
    v3 := make(map[string]uint32)
    var v4 uint8 = 1
    v27 := []int{-9223372036854775808, 32768}
    var v28 uint64 = 108
    v29 := GenS5{F0: 65535, F1: true}
    var v30 int = -1
    switch (("" + ("hello" + ("" + ("foo bar" + ("foo bar" + "\n\t"))))) + "a") {
    case "hello":
        switch "\n\t" {
        case "foo bar":
            if ((v30 + v30) > v30) {
                var v31 uint32 = 790
                switch (v30 >> ((uint(v4)) & 63)) {
                case -2147483648:
                    var v32 int8 = 42
                    _ = v32
                case -229:
                    v33 := GenS6{F0: 42, F1: -1, F2: 15}
                    _ = v33
                case 42:
                    v34 := GenS2{F0: (^(v28))}
                    _ = v34
                }
                _ = v31
            }
        default:
            if false {
                for fi35 := 0; fi35 < 7; fi35++ {
                    var v36 uint32 = 0
                    _ = fi35
                    _ = v36
                }
                v30 = int(v4)
            } else {
                for fi37 := 0; fi37 < 1; fi37++ {
                    switch (fi37 << ((961) & 63)) {
                    case -596:
                        v38 := GenS7{F0: v30, F1: (!true)}
                        _ = v38
                    case 1:
                        if (-612 != 908) {
                            for fi39 := 0; fi39 < 5; fi39++ {
                                v40 := GenS8{F0: 1, F1: -1}
                                _ = fi39
                                _ = v40
                            }
                        }
                    case 926:
                        for fi41 := 0; fi41 < 3; fi41++ {
                            if true {
                                if len(v27) > 0 {
                                    v27[0] = (^(fi37))
                                }
                                v28 += (v28 << ((0) & 63))
                                var v42 int = fi37
                                _ = v42
                            } else {
                                v30 = fi37
                            }
                            v4 += (v4 + 0)
                            _ = fi41
                        }
                    default:
                        var v43 uint = 42
                        _ = v43
                    }
                    _ = fi37
                }
                switch (^(v30)) {
                case -794:
                    if len(v27) > 0 {
                        v27[0] = v30
                    }
                case 32768:
                    switch "a" {
                    case "\n\t":
                        v44 := GenS8{F0: 1000, F1: -32768}
                        _ = v44
                    }
                default:
                    var v45 uint64 = (v28 / ((145) | 1))
                    _ = v45
                }
                fmt.Println(v28)
            }
        }
    case "foo bar":
        v46 := v3[("a" + (("a" + "hello") + "foo bar"))]
        _ = v46
    default:
        var v47 uint32 = 42
        _ = v47
    }
    switch "a" {
    case "a":
        fmt.Println(v4, v4)
    }
    if (282 > 501) {
        v4 += 255
        for fi48 := 0; fi48 < 7; fi48++ {
            var v49 uint32 = 65535
            _ = fi48
            _ = v49
        }
    }
    v50 := GenS9{F0: (^(v30)), F1: (v30 * v30)}
    fmt.Println(v4, v4, v4)
    v4 -= 127
    v28 -= v28
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v50
}
