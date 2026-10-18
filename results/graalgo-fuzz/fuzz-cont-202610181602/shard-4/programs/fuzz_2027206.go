package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32; F2 string }
type GenS2 struct { F0 int64; F1 int; F2 uint }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int; F1 bool }
type GenS5 struct { F0 int16; F1 uint64 }
type GenS6 struct { F0 int32; F1 int64 }

func main() {
    var v1 float64 = (((0.1 + 1e10) + 1.5) - 3.14159)
    var v2 uint32 = 0
    var v3 uint8 = 100
    v4 := make(map[string]bool)
    v5 := make(map[string]string)
    v6 := make(map[int]uint32)
    v7 := GenS1{F0: 4294967295, F1: 2147483647, F2: "hello"}
    for fi8 := 0; fi8 < 12; fi8++ {
        v9 := v4[("hello" + (("a" + ("foo bar" + ("hello" + "\n\t"))) + "hello"))]
        v9 = v9
        v10 := GenS1{F0: 2, F1: -2147483648, F2: "\n\t"}
        _ = fi8
        _ = v10
        _ = v9
    }
    if (122 < 2) {
        for fi11 := 0; fi11 < 5; fi11++ {
            var v12 int8 = -1
            fmt.Println(v3, v2)
            _ = fi11
            _ = v12
        }
        for fi13 := 0; fi13 < 6; fi13++ {
            var v14 int64 = 0
            var v15 float64 = (1e10 - (v1 * (1.5 - 3.14159)))
            fmt.Println(fi13, v1)
            _ = fi13
            _ = v14
            _ = v15
        }
        v2 += 2
    }
    v1 = v1
    v16 := GenS2{F0: 1, F1: 987, F2: 2}
    v1 += float64(float32(((v1 + v1) / ((float64((1e10 - 0.1)) / v1) - 1e10))))
    if true {
        v2 += 2147483648
        if false {
            switch 774 {
            case -9223372036854775808:
                v17 := []int{4294967295, 32768, 4294967295}
                _ = v17
            case 405:
                v18 := []int{32767, -2147483648}
                _ = v18
            default:
                for fi19 := 0; fi19 < 12; fi19++ {
                    var v20 float64 = v1
                    v3 -= v3
                    _ = fi19
                    _ = v20
                }
            }
            if true {
                for fi21 := 0; fi21 < 18; fi21++ {
                    v1 = v1
                    _ = fi21
                }
                v22 := v5["hello"]
                var v23 uint8 = (v3 << ((673) & 63))
                _ = v22
                _ = v23
            } else {
                var v24 int = 32768
                var v25 uint64 = 4294967295
                _ = v24
                _ = v25
            }
        } else {
            var v26 uint64 = 4294967295
            _ = v26
        }
    }
    if (!false) {
        if (true || true) {
            if true {
                v1 += (v1 * float64(1e10))
                v27 := GenS1{F0: 1, F1: 2147483647, F2: ("" + ("" + ""))}
                _ = v27
            }
            v3 = (v3 + 2)
            v28 := GenS2{F0: -1, F1: 127, F2: 2}
            _ = v28
        }
    } else {
        v2 -= v2
        var v29 float64 = 1e10
        _ = v29
    }
    if (340 != -1) {
        v30 := GenS2{F0: -9223372036854775808, F1: -9223372036854775808, F2: 0}
        switch -775 {
        case 42:
            v31 := []int{9223372036854775807, 0, 128}
            _ = v31
        case 32768:
            if true {
                v32 := GenS2{F0: 411, F1: -1, F2: 121}
                v33 := GenS3{F0: 18446744073709551615}
                _ = v32
                _ = v33
            } else {
                if true {
                    v1 = v1
                    var v34 int8 = 14
                    _ = v34
                }
            }
        case 170:
            v3 += 255
        }
        _ = v30
    }
    fmt.Println(v1)
    v3 += 42
    var v35 float32 = 1.5
    for fi36 := 0; fi36 < 20; fi36++ {
        for fi37 := 0; fi37 < 15; fi37++ {
            switch "\n\t" {
            case "foo bar":
                var v38 uint32 = v2
                var v39 uint = 4294967295
                v40 := GenS2{F0: -1, F1: (fi37 & fi36), F2: v39}
                _ = v38
                _ = v39
                _ = v40
            case "":
                if ((v2 | v2) > (v2 * (v2 % ((4294967295) | 1)))) {
                    switch ("foo bar" + "foo bar") {
                    case "hello":
                        if true {
                            v2 -= v2
                            var v41 float32 = ((v35 + v35) / v35)
                            v42 := GenS3{F0: 0}
                            _ = v41
                            _ = v42
                        }
                        fmt.Println(fi36)
                        v43 := GenS4{F0: (fi36 >> ((981) & 63)), F1: true}
                        switch "\n\t" {
                        case "foo bar":
                            if false {
                                v1 = ((v1 - (v1 / (0.1 * (0.1 / 1.5)))) + v1)
                            } else {
                                var v44 int64 = 42
                                _ = v44
                            }
                        default:
                            fmt.Println(v2, fi36)
                        }
                        _ = v43
                    case "a":
                        switch "" {
                        case "hello":
                            var v45 uint = 616
                            _ = v45
                        case "a":
                            v46 := GenS4{F0: fi36, F1: true}
                            _ = v46
                        case "\n\t":
                            v2 = (v2 / ((2) | 1))
                        default:
                            v47 := GenS4{F0: (fi36 | fi36), F1: (((fi37 * -769) != fi36) || true)}
                            _ = v47
                        }
                    }
                }
            case "hello":
                v48 := GenS2{F0: 374, F1: (fi36 / (((fi37 + 2147483647)) | 1)), F2: 169}
                _ = v48
            default:
                v49 := GenS4{F0: fi36, F1: true}
                _ = v49
            }
            if false {
                v50 := GenS5{F0: 0, F1: 521}
                _ = v50
            } else {
                v51 := GenS6{F0: 42, F1: -9223372036854775808}
                switch fi36 {
                case 42:
                    switch "a" {
                    case "a":
                        v52 := GenS5{F0: -261, F1: 4294967295}
                        _ = v52
                    case "hello":
                        var v53 int64 = -630
                        _ = v53
                    }
                }
                v54 := []int{fi36, -701, fi36, (fi36 % ((-2147483648) | 1))}
                switch ("hello" + ("hello" + "\n\t")) {
                case "foo bar":
                    if ((v2 / ((2147483648) | 1)) != 0) {
                        delete(v5, "\n\t")
                    }
                default:
                    v55 := GenS3{F0: 4294967295}
                    _ = v55
                }
                var v56 float32 = v35
                _ = v51
                _ = v54
                _ = v56
            }
            _ = fi37
        }
        v5["foo bar"] = "\n\t"
        _ = fi36
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
