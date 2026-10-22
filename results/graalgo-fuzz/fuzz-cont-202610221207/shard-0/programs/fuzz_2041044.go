package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint; F1 uint64 }
type GenS3 struct { F0 int8; F1 int8 }
func genFn0(p0 int, p1 int) int {
    var v4 bool = (true && (!true))
    var v5 uint32 = 0
    var v6 int32 = 1
    if (false && (!(0 >= int8(v5)))) {
        switch "" {
        case "":
            v6 = int32(p0)
        case "foo bar":
            switch p0 {
            case 1:
                if v4 {
                    var v7 int32 = 1
                    _ = v7
                }
            case 298:
                var v8 int16 = 1
                _ = v8
            case -688:
                v9 := []int{p0, 803, -941, p1}
                _ = v9
            default:
                if true {
                    var v10 int8 = 127
                    _ = v10
                } else {
                    v11 := GenS1{F0: 0}
                    p1 = -751
                    _ = v11
                }
            }
        case "\n\t":
            var v12 float32 = 0.1
            _ = v12
        default:
            fmt.Println(p1, v6)
        }
        fmt.Println(v6)
    }
    switch p0 {
    case 0:
        v6 -= v6
    case 1:
        switch "" {
        case "hello":
            v4 = (!false)
        }
    default:
        for fi13 := 0; fi13 < 3; fi13++ {
            v14 := GenS1{F0: 1}
            p1 += p0
            _ = fi13
            _ = v14
        }
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return (p1 ^ (p0 & 1))
}
func genFn1(p0 int8, p1 int16, p2 int8) uint64 {
    v15 := []int{1, 0, 42, -2147483648}
    if true {
        p0 = (p0 ^ 35)
        switch "hello" {
        case "foo bar":
            switch 42 {
            case -1:
                p0 += p2
            }
        case "a":
            switch "a" {
            case "\n\t":
                for fi16 := 0; fi16 < 6; fi16++ {
                    switch (fi16 & 0) {
                    case 32767:
                        p2 -= p0
                    case 9223372036854775807:
                        if len(v15) > 0 {
                            v15[0] = (fi16 >> ((455) & 63))
                        }
                    case -60:
                        var v17 float64 = 0.0
                        _ = v17
                    default:
                        p2 = p2
                    }
                    switch "\n\t" {
                    case "":
                        var v18 int = fi16
                        _ = v18
                    case "foo bar":
                        _ = len(v15)
                    case "hello":
                        if false {
                            v19 := GenS1{F0: 218}
                            _ = v19
                        }
                    default:
                        p2 = p0
                    }
                    _ = fi16
                }
            }
        case "hello":
            var v20 int32 = -2147483648
            _ = v20
        }
    }
    v15 = append(v15, 42)
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    return 18446744073709551615
}

func main() {
    v1 := make(map[int]int)
    var v2 uint32 = 2147483648
    v3 := []uint32{0, 2147483647, v2, v2}
    v2 -= v2
    for fi21 := 0; fi21 < 15; fi21++ {
        var v22 int = (-(fi21))
        _ = fi21
        _ = v22
    }
    v23 := GenS1{F0: 85}
    if (((42 <= 65535) || true) || true) {
        v24 := GenS1{F0: 1}
        v2 -= v2
        _ = v24
    } else {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        for fi25 := 0; fi25 < 8; fi25++ {
            switch (fi25 / (((fi25 + (^(fi25)))) | 1)) {
            case -336:
                var v26 int = fi25
                _ = v26
            }
            v2 += v2
            _ = fi25
        }
        var v27 int32 = 0
        _ = v27
    }
    v2 -= (v2 ^ 876)
    var v28 uint8 = 1
    v29 := GenS1{F0: (v28 & 127)}
    var v30 int16 = 42
    for fi31 := 0; fi31 < 15; fi31++ {
        if (358 == 821) {
            for fi32 := 0; fi32 < 12; fi32++ {
                v30 -= -50
                _ = fi32
            }
            delete(v1, fi31)
            for fi33 := 0; fi33 < 2; fi33++ {
                var v34 bool = (!true)
                var v35 float64 = (-2.25 - 0.0)
                _ = fi33
                _ = v34
                _ = v35
            }
        } else {
            if (-1 >= 1) {
                var v36 uint32 = (^(v2))
                if (18446744073709551615 >= 255) {
                    v28 = (v28 + (v28 << ((18446744073709551615) & 63)))
                    v37 := GenS2{F0: 255, F1: 0}
                    _ = v37
                }
                for fi38 := 0; fi38 < 2; fi38++ {
                    v28 = v28
                    fmt.Println(fi38)
                    _ = fi38
                }
                _ = v36
            }
            switch (fi31 - 9223372036854775807) {
            case -9223372036854775808:
                var v39 int8 = -128
                _ = v39
            case 9223372036854775807:
                for fi40 := 0; fi40 < 12; fi40++ {
                    fmt.Println(fi31, v28, v28)
                    _ = fi40
                }
            case 127:
                v2 -= v2
            }
            if (true || (("\n\t" + "\n\t") < ("\n\t" + "foo bar"))) {
                switch fi31 {
                case -2147483648:
                    v3 = append(v3, (v2 * (v2 / (((v2 & 0)) | 1))))
                case 0:
                    v2 += 2
                }
                v30 -= (v30 % (((v30 >> ((1) & 63))) | 1))
            } else {
                switch genFn0(-481, genFn0(fi31, (fi31 * 710))) {
                case 128:
                    for fi41 := 0; fi41 < 12; fi41++ {
                        var v42 float32 = 3.14159
                        _ = fi41
                        _ = v42
                    }
                case 2147483647:
                    switch fi31 {
                    case -1:
                        v43 := v1[fi31]
                        _ = v43
                    case 160:
                        var v44 int64 = 1
                        _ = v44
                    }
                case 896:
                    var v45 uint8 = uint8(v2)
                    _ = v45
                }
            }
        }
        v28 = (^(v28))
        var v46 string = "foo bar"
        _ = fi31
        _ = v46
    }
    if (v2 >= (v2 % (((^(v2))) | 1))) {
        v47 := GenS3{F0: 127, F1: -128}
        for fi48 := 0; fi48 < 7; fi48++ {
            switch ("a" + ((("" + "foo bar") + "hello") + ("" + (("foo bar" + "a") + (("foo bar" + "\n\t") + "foo bar"))))) {
            case "":
                if (!((!((0.0 <= 0.1) || (2147483647 >= (^(v2))))) && true)) {
                    _ = len(v1)
                    if ((((0.0 * 0.0) - ((-2.25 - 3.14159) * 1e10)) * ((float64(((3.14159 + 0.0) + 1e10)) * -2.25) + ((float64(3.14159) - float64(((-2.25 - (0.1 * (0.0 * 0.1))) + 0.1))) / float64((1e10 + (float32(1e10) - float32(3.14159))))))) != -2.25) {
                        v49 := GenS1{F0: v28}
                        v3 = append(v3, v2)
                        _ = v49
                    }
                    v28 = (^(v28))
                }
            default:
                if ((false && (-1 == -106)) && (false && true)) {
                    v2 -= (^(v2))
                    v50 := GenS1{F0: 1}
                    _ = v50
                } else {
                    v28 = (v28 % ((v28) | 1))
                    var v51 uint64 = 494
                    _ = v51
                }
            }
            _ = fi48
        }
        v30 += v30
        _ = v47
    }
    v1[166] = -9223372036854775808
    if ((false || (75 == 32767)) || (((!(!false)) || true) || (0 != 0))) {
        var v52 uint8 = 90
        if (!(!(4294967295 >= 588))) {
            v53 := GenS3{F0: 1, F1: 0}
            v28 -= (v52 * (v52 << ((42) & 63)))
            _ = v53
        }
        _ = v52
    }
    v54 := v1[-962]
    if len(v3) > 0 {
        v3[0] = 791
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v54:", v54)
    _ = v1
    _ = v2
    _ = v23
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v54
}
