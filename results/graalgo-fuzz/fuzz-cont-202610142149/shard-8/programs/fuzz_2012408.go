package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
type GenS2 struct { F0 string; F1 int8; F2 uint16 }
type GenS3 struct { F0 int64; F1 int16 }
type GenS4 struct { F0 int8; F1 uint16 }
type GenS5 struct { F0 int8; F1 bool }
type GenS6 struct { F0 bool; F1 int64 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 uint16; F1 uint16; F2 bool }
type GenS9 struct { F0 uint64; F1 string }
func genFn0(p0 int8) uint16 {
    switch "" {
    case "hello":
        p0 = -128
    case "\n\t":
        fmt.Println(p0)
    case "a":
        p0 = (p0 / ((-22) | 1))
    default:
        var v6 float32 = 3.14159
        _ = v6
    }
    if false {
        p0 = (p0 | 127)
        p0 = 1
    } else {
        p0 = p0
        var v7 uint32 = 865
        _ = v7
    }
    p0 += p0
    _ = p0
    return 413
}

func main() {
    var v1 bool = (982 == 42)
    v2 := make(map[string]bool)
    var v3 bool = (325 == 32768)
    var v4 uint8 = 199
    var v5 int32 = -2147483648
    v8 := GenS1{F0: 32767, F1: 1}
    if v3 {
        var v9 uint16 = 42
        switch "foo bar" {
        case "a":
            v10 := GenS2{F0: "\n\t", F1: 127, F2: genFn0(100)}
            _ = v10
        default:
            for fi11 := 0; fi11 < 7; fi11++ {
                var v12 int32 = v5
                switch "hello" {
                case "":
                    var v13 uint = 1
                    _ = v13
                case "\n\t":
                    v3 = (!(!v1))
                default:
                    if (356 < 9223372036854775807) {
                        v2[("hello" + "")] = ("hello" < "a")
                    } else {
                        v4 = v4
                        v3 = (v1 || (!(9223372036854775807 == 42)))
                        v12 += (v5 - (v12 - v5))
                    }
                }
                v14 := GenS3{F0: int64(v4), F1: 537}
                _ = fi11
                _ = v12
                _ = v14
            }
        }
        if ((((0.0 + 0.0) / -2.25) - 1.5) < 3.14159) {
            fmt.Println(v1)
        }
        _ = v9
    }
    var v15 uint8 = v4
    if v3 {
        if ((float64(((float32(float64(3.14159)) / -2.25) + float32(float64(1.5)))) != (0.0 * 3.14159)) && (1 <= -789)) {
            v16 := v2["a"]
            v16 = (true || (32767 != 511))
            var v17 bool = v3
            _ = v16
            _ = v17
        } else {
            v5 -= (v5 + v5)
            v18 := []int{int(v5)}
            var v19 uint64 = 4294967295
            _ = v18
            _ = v19
        }
        fmt.Println(v3)
        var v20 int = 127
        _ = v20
    }
    if ((127 < 1) && (v5 != v5)) {
        fmt.Println(v3)
    } else {
        var v21 uint64 = 4294967295
        v1 = false
        switch "foo bar" {
        case "\n\t":
            for fi22 := 0; fi22 < 11; fi22++ {
                for fi23 := 0; fi23 < 2; fi23++ {
                    switch ("" + ("\n\t" + "")) {
                    case "":
                        v5 -= v5
                    default:
                        v2[""] = (42 >= -468)
                    }
                    v3 = ((763 != 364) && (v3 || v3))
                    _ = fi23
                }
                for fi24 := 0; fi24 < 17; fi24++ {
                    if false {
                        switch int(v21) {
                        case -9223372036854775808:
                            var v25 int16 = 1
                            _ = v25
                        case 1:
                            var v26 int8 = 94
                            _ = v26
                        default:
                            v3 = v1
                        }
                    }
                    v5 += v5
                    v27 := []int{}
                    _ = fi24
                    _ = v27
                }
                if v1 {
                    switch "\n\t" {
                    case "a":
                        var v28 int32 = 207
                        _ = v28
                    case "\n\t":
                        for fi29 := 0; fi29 < 1; fi29++ {
                            var v30 int64 = 0
                            _ = fi29
                            _ = v30
                        }
                    default:
                        switch "" {
                        case "":
                            v1 = true
                        case "hello":
                            v5 = v5
                        case "a":
                            if true {
                                v1 = (v3 || (((1e10 - -2.25) * float64(float32(0.1))) > 0.0))
                                if v3 {
                                    v1 = ((int64(v21) >= 42) && true)
                                    for fi31 := 0; fi31 < 16; fi31++ {
                                        var v32 uint64 = 262
                                        fmt.Println(fi31, v4)
                                        v33 := GenS2{F0: (("foo bar" + "foo bar") + "hello"), F1: 42, F2: genFn0(-1)}
                                        _ = fi31
                                        _ = v32
                                        _ = v33
                                    }
                                }
                            }
                        default:
                            v15 = (v4 >> ((488) & 63))
                        }
                    }
                    v34 := GenS2{F0: "", F1: 113, F2: 537}
                    if v3 {
                        if (!v3) {
                            var v35 int8 = 1
                            switch "foo bar" {
                            case "":
                                v36 := GenS4{F0: 118, F1: 187}
                                _ = v36
                            default:
                                var v37 uint = 184
                                _ = v37
                            }
                            _ = v35
                        }
                        v38 := GenS5{F0: -128, F1: v1}
                        if v1 {
                            v21 = (^(v21))
                        }
                        _ = v38
                    }
                    _ = v34
                }
                _ = fi22
            }
        }
        _ = v21
    }
    v39 := GenS6{F0: v3, F1: 416}
    v15 -= (v4 << ((65535) & 63))
    v5 += -2147483648
    v15 += (v15 + v15)
    v40 := GenS7{F0: 4294967295}
    switch "hello" {
    case "\n\t":
        var v41 int8 = 1
        _ = v41
    case "a":
        if false {
            v4 = (v15 + (^(v15)))
            var v42 int = 128
            _ = v42
        }
    case "hello":
        fmt.Println(v3)
    default:
        v4 -= (v4 & 229)
    }
    for fi43 := 0; fi43 < 10; fi43++ {
        v2["foo bar"] = (true || true)
        var v44 string = "hello"
        var v45 int = (fi43 >> ((255) & 63))
        _ = fi43
        _ = v44
        _ = v45
    }
    v46 := GenS8{F0: 0, F1: 247, F2: ((100 != 15) || (18446744073709551615 != 42))}
    v2["hello"] = (445 > 0)
    for fi47 := 0; fi47 < 5; fi47++ {
        var v48 int64 = 0
        var v49 uint16 = genFn0(-1)
        if (!(v5 == -1)) {
            for fi50 := 0; fi50 < 15; fi50++ {
                v51 := []int{fi50, -718, fi50}
                var v52 int32 = v5
                _ = fi50
                _ = v51
                _ = v52
            }
            v53 := GenS8{F0: (v49 - (v49 / ((genFn0(1)) | 1))), F1: (v49 + (v49 | (v49 & genFn0(127)))), F2: ((0 <= int8(v5)) && v1)}
            v54 := []int{0, fi47, int(v48)}
            _ = v53
            _ = v54
        }
        _ = fi47
        _ = v48
        _ = v49
    }
    for fi55 := 0; fi55 < 9; fi55++ {
        v56 := GenS8{F0: 765, F1: 255, F2: true}
        _ = fi55
        _ = v56
    }
    for fi57 := 0; fi57 < 17; fi57++ {
        for fi58 := 0; fi58 < 11; fi58++ {
            v59 := GenS9{F0: 4294967295, F1: "foo bar"}
            _ = fi58
            _ = v59
        }
        _ = fi57
    }
    fmt.Println(v1, v3, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v46
    _ = v5
    _ = v8
}
