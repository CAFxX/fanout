package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint }
type GenS2 struct { F0 uint32; F1 int32; F2 int64 }
type GenS3 struct { F0 int8; F1 uint32; F2 uint32 }
type GenS4 struct { F0 uint8; F1 int }
type GenS5 struct { F0 string; F1 uint64 }
type GenS6 struct { F0 int; F1 uint8 }
type GenS7 struct { F0 uint64; F1 int64; F2 uint }
func genFn0() uint8 {
    var v7 bool = (!(9223372036854775807 != 128))
    var v8 bool = (18446744073709551615 == 1)
    v9 := []int{2147483647, 42, 15, -9223372036854775808}
    v7 = true
    v10 := GenS1{F0: 1, F1: 437}
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 26
}
func genFn1(p0 bool, p1 uint32, p2 int32) uint {
    var v11 uint64 = 362
    if (-2.25 < ((0.0 / 1.0) / (1e10 * -2.25))) {
        for fi12 := 0; fi12 < 3; fi12++ {
            for fi13 := 0; fi13 < 8; fi13++ {
                var v14 int64 = -9223372036854775808
                _ = fi13
                _ = v14
            }
            if (-463 != 32768) {
                v15 := GenS2{F0: (p1 | p1), F1: p2, F2: 296}
                v16 := GenS3{F0: 0, F1: 0, F2: (p1 - p1)}
                for fi17 := 0; fi17 < 4; fi17++ {
                    p1 -= 127
                    switch (("" + ("\n\t" + "hello")) + ((("" + "\n\t") + "a") + ("foo bar" + ("a" + "\n\t")))) {
                    case "hello":
                        p0 = (fi12 == (fi17 << ((42) & 63)))
                    }
                    p1 = (^(p1))
                    _ = fi17
                }
                _ = v15
                _ = v16
            }
            switch "hello" {
            case "foo bar":
                fmt.Println(p2)
            case "hello":
                p2 -= -180
            default:
                v18 := GenS3{F0: 46, F1: 65535, F2: 42}
                _ = v18
            }
            _ = fi12
        }
    } else {
        v11 += v11
        p0 = p0
        var v19 int8 = -35
        _ = v19
    }
    v20 := []int{4294967295, 419, 975, 2147483647}
    var v21 int32 = 77
    fmt.Println(p2, p0, v11)
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v20
    _ = v21
    return 18446744073709551615
}

func main() {
    var v1 uint16 = 1
    v2 := []string{"", "foo bar", (("hello" + "hello") + ("\n\t" + "a"))}
    v3 := make(map[string]int)
    v4 := make(map[int]uint32)
    v5 := []uint32{820, 1, 2147483648, 2147483648}
    var v6 uint32 = 1
    v1 -= (^(v1))
    fmt.Println(v1, v1)
    var v22 uint32 = v6
    v1 += v1
    var v23 string = "\n\t"
    for fi24 := 0; fi24 < 4; fi24++ {
        var v25 uint16 = (v1 | uint16(v22))
        for fi26 := 0; fi26 < 7; fi26++ {
            if ((319 != 2) || ((false || (!false)) && true)) {
                v25 = 42
                v27 := GenS1{F0: 1, F1: 1}
                var v28 uint = 1
                if true {
                    v29 := GenS3{F0: -124, F1: 2147483647, F2: 2}
                    var v30 int64 = int64(fi24)
                    v31 := GenS2{F0: 2147483647, F1: 754, F2: (-(v30))}
                    v32 := GenS4{F0: 110, F1: fi24}
                    var v33 int64 = -1
                    fmt.Println(v30, v22)
                    _ = v29
                    _ = v30
                    _ = v31
                    _ = v32
                    _ = v33
                }
                if false {
                    var v34 uint16 = uint16(v28)
                    v3["\n\t"] = 2147483647
                    _ = v34
                }
                _ = v27
                _ = v28
            } else {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                if (int(v22) > (fi24 * 2147483647)) {
                    if true {
                        if (v23 >= ("hello" + v23)) {
                            var v35 uint64 = 18446744073709551615
                            _ = v35
                        }
                        fmt.Println(v22, v1, v1)
                    }
                    fmt.Println(v6)
                    v36 := GenS4{F0: 42, F1: fi26}
                    _ = v36
                }
                v37 := GenS1{F0: 0, F1: 18446744073709551615}
                _ = v37
            }
            v38 := v4[(^(fi26))]
            _ = fi26
            _ = v38
        }
        _ = fi24
        _ = v25
    }
    var v39 float64 = -2.25
    if (32767 < -1) {
        switch "foo bar" {
        case "a":
            if true {
                _ = len(v4)
                v2 = append(v2, (v23 + v23))
                switch -605 {
                case 127:
                    var v40 uint8 = uint8(v22)
                    _ = v40
                case 42:
                    delete(v4, 0)
                }
            }
        }
        if ((-583 <= 1) || false) {
            v41 := GenS5{F0: ("foo bar" + ("" + v23)), F1: 502}
            _ = v41
        } else {
            v6 = 2147483647
            for fi42 := 0; fi42 < 11; fi42++ {
                var v43 string = v23
                _ = fi42
                _ = v43
            }
            var v44 uint = 749
            _ = v44
        }
    }
    fmt.Println(v1, v22, v23)
    for fi45 := 0; fi45 < 7; fi45++ {
        for fi46 := 0; fi46 < 8; fi46++ {
            var v47 int16 = -218
            v48 := GenS6{F0: fi46, F1: 128}
            delete(v4, (fi45 << ((4294967295) & 63)))
            _ = len(v3)
            _ = fi46
            _ = v47
            _ = v48
        }
        for fi49 := 0; fi49 < 6; fi49++ {
            var v50 int = 42
            _ = len(v4)
            _ = fi49
            _ = v50
        }
        _ = fi45
    }
    v51 := GenS7{F0: 1, F1: 9223372036854775807, F2: 42}
    var v52 bool = (-1 >= -1)
    for fi53 := 0; fi53 < 19; fi53++ {
        v52 = (uint64(v1) == 532)
        v1 += 32767
        switch (fi53 & (fi53 & (fi53 + 2147483647))) {
        case 0:
            var v54 float32 = ((0.1 + (((1.5 + 0.1) * (1e10 * 1e10)) - 3.14159)) - 1e10)
            _ = v54
        case -505:
            var v55 float64 = (float64(((0.1 - ((float32(v39) - 1.5) - 1e10)) - 3.14159)) - 0.0)
            _ = v55
        case 32768:
            v5 = append(v5, 80)
        }
        _ = fi53
    }
    v5 = append(v5, v6)
    v56 := GenS7{F0: 2, F1: 9223372036854775807, F2: 18446744073709551615}
    v6 -= v6
    if len(v2) > 0 {
        v2[0] = v23
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v52:", v52)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v39
    _ = v4
    _ = v5
    _ = v51
    _ = v52
    _ = v56
    _ = v6
}
