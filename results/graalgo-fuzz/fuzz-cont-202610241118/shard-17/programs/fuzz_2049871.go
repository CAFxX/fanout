package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint8; F1 uint64 }
type GenS3 struct { F0 int64; F1 uint64; F2 int64 }
type GenS4 struct { F0 uint16; F1 uint8; F2 uint8 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 string }
type GenS7 struct { F0 uint32; F1 bool }
func genFn0() int32 {
    v4 := GenS1{F0: 661}
    var v5 int64 = 0
    switch ((("" + "\n\t") + ("a" + "\n\t")) + "a") {
    case "":
        v6 := []int{32768, -466, -786, -9223372036854775808}
        _ = v6
    case "\n\t":
        v7 := []int{-862, 42, 9223372036854775807}
        _ = v7
    }
    switch "" {
    case "\n\t":
        v8 := GenS1{F0: 1}
        _ = v8
    }
    switch -633 {
    case -381:
        v9 := []int{}
        _ = v9
    case 42:
        v5 = v5
    }
    _ = v4
    _ = v5
    return 42
}
func genFn1(p0 uint16, p1 int64, p2 uint32) int {
    p0 += (p0 + 590)
    if ((-2.25 + 1e10) <= 1.5) {
        var v10 string = "hello"
        _ = v10
    }
    var v11 int16 = 1
    v12 := []int{1, -326, 2147483647, 128}
    switch "hello" {
    case "":
        fmt.Println(p1)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    return 9223372036854775807
}

func main() {
    var v1 uint16 = 0
    v2 := []string{"a", ("foo bar" + "\n\t"), "\n\t", "foo bar"}
    var v3 int64 = 771
    v3 = v3
    switch "a" {
    case "\n\t":
        for fi13 := 0; fi13 < 10; fi13++ {
            v14 := GenS1{F0: 65535}
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            for fi15 := 0; fi15 < 2; fi15++ {
                v1 = v1
                v16 := GenS1{F0: 43}
                if false {
                    v17 := GenS2{F0: 128, F1: 477}
                    _ = v17
                }
                _ = fi15
                _ = v16
            }
            _ = fi13
            _ = v14
        }
    case "hello":
        v1 = v1
    case "a":
        var v18 uint64 = 4294967295
        _ = v18
    default:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    if true {
        if false {
            var v19 uint = 42
            v2 = append(v2, ("foo bar" + ""))
            _ = v19
        }
        switch int(v3) {
        case -2147483648:
            var v20 uint32 = 240
            _ = v20
        case 32767:
            v21 := GenS2{F0: 127, F1: 526}
            _ = v21
        }
        for fi22 := 0; fi22 < 4; fi22++ {
            for fi23 := 0; fi23 < 10; fi23++ {
                v24 := GenS3{F0: (v3 >> ((1) & 63)), F1: 1, F2: 42}
                fmt.Println(v3, fi22)
                _ = fi23
                _ = v24
            }
            if (((-2.25 + 0.1) != (0.0 / 1.5)) || true) {
                for fi25 := 0; fi25 < 4; fi25++ {
                    v26 := GenS3{F0: -425, F1: 18446744073709551615, F2: (v3 ^ v3)}
                    _ = fi25
                    _ = v26
                }
                v3 = v3
            }
            _ = fi22
        }
    }
    var v27 float32 = 0.1
    for fi28 := 0; fi28 < 11; fi28++ {
        if (true || (2 != 566)) {
            switch "hello" {
            case "foo bar":
                switch 757 {
                case 42:
                    var v29 int16 = 42
                    _ = v29
                case 714:
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                default:
                    switch "foo bar" {
                    case "foo bar":
                        if false {
                            _ = len(v2)
                        } else {
                            var v30 uint16 = v1
                            switch "a" {
                            case "hello":
                                var v31 uint64 = 0
                                _ = v31
                            case "\n\t":
                                v27 = v27
                            case "":
                                var v32 float64 = 3.14159
                                _ = v32
                            }
                            var v33 uint16 = v30
                            _ = v30
                            _ = v33
                        }
                    case "\n\t":
                        if true {
                            for fi34 := 0; fi34 < 15; fi34++ {
                                var v35 uint = 2
                                switch ("" + "a") {
                                case "\n\t":
                                    v1 = 42
                                case "hello":
                                    v36 := GenS2{F0: 255, F1: 1}
                                    fmt.Println(fi28, fi34, fi28)
                                    var v37 int8 = 1
                                    _ = v36
                                    _ = v37
                                }
                                var v38 float64 = 1e10
                                _ = fi34
                                _ = v35
                                _ = v38
                            }
                            v39 := GenS2{F0: 37, F1: 20}
                            _ = v39
                        }
                    }
                }
            case "hello":
                fmt.Println(v3)
            }
            var v40 uint32 = 65535
            if true {
                for fi41 := 0; fi41 < 7; fi41++ {
                    if false {
                        var v42 int32 = 0
                        var v43 float64 = 1e10
                        _ = v42
                        _ = v43
                    } else {
                        v3 = (v3 / (((v3 / ((int64(v1)) | 1))) | 1))
                        fmt.Println(v40, fi28, fi28)
                        var v44 float32 = -2.25
                        _ = v44
                    }
                    _ = fi41
                }
            }
            _ = v40
        }
        v45 := GenS4{F0: v1, F1: 255, F2: 42}
        _ = fi28
        _ = v45
    }
    var v46 int = 879
    v47 := GenS4{F0: v1, F1: 128, F2: 44}
    switch "a" {
    case "foo bar":
        switch v46 {
        case 2147483647:
            v27 -= 1.5
        case 42:
            fmt.Println(v1)
        default:
            for fi48 := 0; fi48 < 15; fi48++ {
                var v49 float64 = ((1.5 + (3.14159 - 0.1)) / -2.25)
                _ = fi48
                _ = v49
            }
        }
    case "hello":
        fmt.Println(v3, v46, v46)
    case "":
        if true {
            if true {
                for fi50 := 0; fi50 < 16; fi50++ {
                    if true {
                        var v51 uint8 = 249
                        switch "\n\t" {
                        case "hello":
                            var v52 int32 = -2147483648
                            _ = v52
                        case "foo bar":
                            v51 += 0
                        case "a":
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        }
                        _ = v51
                    }
                    var v53 int8 = -128
                    _ = fi50
                    _ = v53
                }
                v2 = append(v2, "hello")
                if (629 <= 42) {
                    fmt.Println(v3, v3, v27)
                    if len(v2) > 0 {
                        v2[0] = "hello"
                    }
                }
            } else {
                for fi54 := 0; fi54 < 4; fi54++ {
                    for fi55 := 0; fi55 < 1; fi55++ {
                        var v56 uint32 = 4294967295
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        _ = fi55
                        _ = v56
                    }
                    v57 := GenS5{F0: 947}
                    if (2147483648 != 0) {
                        if len(v2) > 0 {
                            v2[0] = (("a" + "foo bar") + ("hello" + "a"))
                        }
                    } else {
                        v3 -= v3
                        if false {
                            v1 += v1
                            v1 -= 42
                        } else {
                            if (!false) {
                                if ((v27 < (v27 * 3.14159)) || false) {
                                    v46 -= genFn1((v1 << ((1) & 63)), (v3 << ((1) & 63)), 1)
                                    v46 = v46
                                } else {
                                    for fi58 := 0; fi58 < 9; fi58++ {
                                        fmt.Println(v27, v3, v1)
                                        _ = fi58
                                    }
                                }
                            } else {
                                v59 := GenS1{F0: 99}
                                v60 := GenS6{F0: ("foo bar" + ("a" + "foo bar"))}
                                _ = len(v2)
                                _ = v59
                                _ = v60
                            }
                            v1 -= 65535
                        }
                        v46 -= fi54
                    }
                    _ = fi54
                    _ = v57
                }
                if (false || false) {
                    v2 = append(v2, "foo bar")
                } else {
                    fmt.Println(v3)
                    if len(v2) > 0 {
                        v2[0] = "\n\t"
                    }
                    for fi61 := 0; fi61 < 16; fi61++ {
                        fmt.Println(v3)
                        _ = fi61
                    }
                }
            }
            v2 = append(v2, ("foo bar" + "a"))
        } else {
            if ("" == "a") {
                fmt.Println(v1, v46, v1)
            } else {
                v1 = v1
            }
            for fi62 := 0; fi62 < 19; fi62++ {
                v3 -= v3
                if true {
                    var v63 int8 = -23
                    for fi64 := 0; fi64 < 10; fi64++ {
                        v46 += (^(fi62))
                        _ = fi64
                    }
                    _ = v63
                }
                _ = fi62
            }
        }
    default:
        v2 = append(v2, "hello")
    }
    v65 := GenS7{F0: 65535, F1: true}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v46
    _ = v47
    _ = v65
}
