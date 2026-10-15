package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32 }
type GenS2 struct { F0 int64; F1 int8 }
type GenS3 struct { F0 uint32; F1 int64 }
type GenS4 struct { F0 uint32; F1 int32 }
type GenS5 struct { F0 int16; F1 uint; F2 int64 }
type GenS6 struct { F0 uint; F1 uint32; F2 uint32 }
func genFn0(p0 uint32) uint16 {
    for fi7 := 0; fi7 < 12; fi7++ {
        for fi8 := 0; fi8 < 16; fi8++ {
            p0 = p0
            switch "\n\t" {
            case "hello":
                fmt.Println(fi7, p0, fi8)
            case "foo bar":
                p0 = 0
            }
            var v9 uint = 957
            _ = fi8
            _ = v9
        }
        var v10 float64 = ((-2.25 / float64(-2.25)) / (0.1 / -2.25))
        v11 := []int{}
        _ = fi7
        _ = v10
        _ = v11
    }
    for fi12 := 0; fi12 < 4; fi12++ {
        if (!true) {
            p0 = 600
        } else {
            fmt.Println(fi12, p0)
        }
        for fi13 := 0; fi13 < 12; fi13++ {
            switch (fi13 & (^(fi13))) {
            case -519:
                fmt.Println(fi12, fi13, p0)
                fmt.Println(p0, fi13, fi13)
            case 4294967295:
                fmt.Println(fi13, p0, fi13)
            case -139:
                p0 = p0
            default:
                var v14 int64 = 9223372036854775807
                _ = v14
            }
            v15 := []int{-66}
            _ = fi13
            _ = v15
        }
        _ = fi12
    }
    _ = p0
    return 255
}
func genFn1(p0 uint16) string {
    v16 := []int{}
    if true {
        if ((false || true) || true) {
            p0 = 979
            v17 := GenS1{F0: true, F1: 2}
            _ = v17
        } else {
            var v18 uint = 200
            _ = v18
        }
        if true {
            v16 = append(v16, -1)
        } else {
            var v19 float32 = 1.5
            v20 := GenS1{F0: true, F1: 2147483647}
            v16 = append(v16, -2147483648)
            _ = v19
            _ = v20
        }
    }
    fmt.Println(p0)
    _ = len(v16)
    v21 := GenS1{F0: ((p0 ^ 255) == (^(p0))), F1: 2}
    _ = p0
    _ = v16
    _ = v21
    return "a"
}

func main() {
    var v1 int = -959
    v2 := make(map[uint32]bool)
    var v3 uint16 = 42
    v4 := make(map[string]string)
    v5 := []uint32{2147483648, 1}
    v6 := []uint32{1, 4294967295}
    v2[973] = true
    fmt.Println(v1, v3)
    v22 := GenS2{F0: 1, F1: 0}
    var v23 bool = (!false)
    switch v1 {
    case -9223372036854775808:
        fmt.Println(v23)
    default:
        for fi24 := 0; fi24 < 4; fi24++ {
            v3 = (v3 & 1)
            v25 := GenS2{F0: 42, F1: -77}
            if ((9223372036854775807 >= 42) && ((0.0 + (float32(-2.25) - 0.0)) >= (3.14159 + 1e10))) {
                for fi26 := 0; fi26 < 1; fi26++ {
                    var v27 bool = (-128 >= 127)
                    _ = fi26
                    _ = v27
                }
                v23 = true
                var v28 uint32 = 412
                _ = v28
            } else {
                var v29 float32 = ((1e10 + (3.14159 + 1.5)) * 0.0)
                if v23 {
                    if len(v6) > 0 {
                        v6[0] = 0
                    }
                    v30 := GenS2{F0: -968, F1: 100}
                    _ = v30
                }
                _ = v29
            }
            _ = fi24
            _ = v25
        }
    }
    switch v1 {
    case 42:
        fmt.Println(v3, v23, v1)
    }
    var v31 int16 = 32767
    var v32 int64 = 9223372036854775807
    for fi33 := 0; fi33 < 12; fi33++ {
        v2[42] = (!false)
        _ = fi33
    }
    switch (v1 << ((18446744073709551615) & 63)) {
    case 344:
        switch v1 {
        case 9223372036854775807:
            v32 -= v32
        case 42:
            for fi34 := 0; fi34 < 20; fi34++ {
                var v35 string = "hello"
                switch fi34 {
                case 127:
                    var v36 int16 = v31
                    _ = v36
                case -598:
                    for fi37 := 0; fi37 < 17; fi37++ {
                        if (v35 != v35) {
                            var v38 int64 = v32
                            _ = v38
                        }
                        _ = fi37
                    }
                default:
                    v39 := GenS3{F0: 930, F1: v32}
                    _ = v39
                }
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                _ = fi34
                _ = v35
            }
        case -2147483648:
            if ((^(v3)) < v3) {
                if false {
                    v23 = v23
                    fmt.Println(v32, v31, v1)
                } else {
                    if len(v5) > 0 {
                        v5[0] = 258
                    }
                    var v40 float64 = (0.1 * (float64(0.1) / 1.0))
                    _ = v40
                }
            }
        }
    case -1:
        v23 = v23
    }
    var v41 float64 = (0.0 - float64(-2.25))
    if (v23 || ((1.5 * 0.1) != 0.1)) {
        v3 = genFn0(0)
        v32 = v32
        v42 := GenS4{F0: 2147483648, F1: -2147483648}
        _ = v42
    }
    switch "" {
    case "foo bar":
        fmt.Println(v23, v1, v3)
    default:
        if v23 {
            v43 := GenS5{F0: (v31 - (v31 << ((0) & 63))), F1: 0, F2: int64(v31)}
            v3 = (^(v3))
            v44 := GenS6{F0: 65535, F1: 715, F2: 1}
            _ = v43
            _ = v44
        } else {
            if v23 {
                v45 := v4[("hello" + ("foo bar" + (((("\n\t" + "hello") + "a") + ("foo bar" + ("foo bar" + ("" + "")))) + ("foo bar" + "a"))))]
                v46 := GenS6{F0: 447, F1: 980, F2: 789}
                _ = v45
                _ = v46
            }
            v1 -= (v1 << ((18446744073709551615) & 63))
        }
    }
    var v47 uint = 42
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v47:", v47)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v41
    _ = v47
    _ = v5
    _ = v6
}
