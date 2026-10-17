package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64; F2 int64 }
type GenS2 struct { F0 uint16; F1 int; F2 int }
type GenS3 struct { F0 uint64; F1 uint; F2 int64 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int16; F1 int8 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 uint; F1 string; F2 int }
type GenS9 struct { F0 uint16; F1 int }
type GenS10 struct { F0 int32; F1 uint8; F2 uint }

func main() {
    var v1 int16 = 42
    v2 := []string{("foo bar" + "\n\t"), "hello", "\n\t", "foo bar"}
    v3 := make(map[string]string)
    v4 := []int{32768, -817, 2147483647, 4294967295}
    v1 += 42
    if true {
        _ = len(v4)
        var v5 int8 = 42
        _ = v5
    } else {
        if true {
            v1 -= 936
            fmt.Println(v1)
            fmt.Println(v1)
        } else {
            v6 := GenS1{F0: 127, F1: -427, F2: -39}
            v1 = (v1 & (v1 << ((4294967295) & 63)))
            var v7 uint = 1
            _ = v6
            _ = v7
        }
        if (!false) {
            fmt.Println(v1)
            var v8 int16 = (v1 | 32767)
            _ = v8
        } else {
            if (9223372036854775807 > 0) {
                switch ("" + "\n\t") {
                case "\n\t":
                    var v9 float64 = 1e10
                    _ = v9
                case "":
                    v10 := GenS2{F0: 679, F1: -381, F2: 9223372036854775807}
                    _ = v10
                case "hello":
                    if true {
                        v1 += 32767
                        v1 -= (v1 + (v1 & v1))
                    }
                default:
                    fmt.Println(v1)
                }
                fmt.Println(v1)
                v1 += (^(v1))
            }
        }
    }
    for fi11 := 0; fi11 < 3; fi11++ {
        v1 -= v1
        for fi12 := 0; fi12 < 12; fi12++ {
            fmt.Println(fi12, fi12, fi12)
            switch 443 {
            case 934:
                var v13 float32 = (3.14159 / -2.25)
                _ = v13
            case 127:
                if true {
                    fmt.Println(fi12, v1, fi11)
                    v14 := GenS3{F0: 18446744073709551615, F1: 65535, F2: -514}
                    v1 -= v1
                    var v15 int64 = -1
                    _ = v14
                    _ = v15
                } else {
                    if false {
                        var v16 uint16 = 989
                        _ = v16
                    }
                }
            default:
                if true {
                    if (255 == 500) {
                        if (!(1e10 >= 0.0)) {
                            fmt.Println(fi11, fi11, fi11)
                        } else {
                            switch ("hello" + "a") {
                            case "a":
                                fmt.Println(fi12, v1)
                            case "foo bar":
                                v1 = v1
                                var v17 uint64 = 4294967295
                                if (1 >= -128) {
                                    v1 = v1
                                    v3["hello"] = (("hello" + ("a" + ("a" + ("foo bar" + ("" + "\n\t"))))) + "")
                                    if true {
                                        v17 = (v17 ^ 1)
                                        switch "foo bar" {
                                        case "":
                                            var v18 int8 = 42
                                            _ = v18
                                        case "a":
                                            v19 := GenS1{F0: -128, F1: -1, F2: 0}
                                            _ = v19
                                        default:
                                            v1 += v1
                                        }
                                    } else {
                                        var v20 string = (("\n\t" + ("hello" + ("hello" + "hello"))) + "\n\t")
                                        v20 = (v20 + v20)
                                        _ = v20
                                    }
                                } else {
                                    var v21 float64 = 1e10
                                    if true {
                                        var v22 int64 = int64(v17)
                                        switch "" {
                                        case "a":
                                            v1 -= v1
                                            var v23 uint16 = 255
                                            _ = v23
                                        default:
                                            if false {
                                                if true {
                                                    v24 := v3[((("hello" + "a") + ("foo bar" + "\n\t")) + ("hello" + ""))]
                                                    v25 := GenS3{F0: (v17 + (^(v17))), F1: 2, F2: v22}
                                                    v17 -= v17
                                                    _ = v24
                                                    _ = v25
                                                } else {
                                                    fmt.Println(v1, v1, v22)
                                                    fmt.Println(fi11)
                                                    v22 += -1
                                                }
                                                v26 := GenS4{F0: 4294967295}
                                                switch fi11 {
                                                case 0:
                                                    if false {
                                                        var v27 int32 = -1
                                                        _ = v27
                                                    }
                                                case -9223372036854775808:
                                                    var v28 int8 = 100
                                                    _ = v28
                                                }
                                                switch "" {
                                                case "\n\t":
                                                    v29 := GenS5{F0: -32768, F1: 42}
                                                    _ = v29
                                                case "hello":
                                                    fmt.Println(v22, fi11, v21)
                                                }
                                                var v30 float32 = 1.5
                                                var v31 int = (fi11 >> ((65535) & 63))
                                                _ = v26
                                                _ = v30
                                                _ = v31
                                            } else {
                                                switch (-(fi11)) {
                                                case 0:
                                                    switch "\n\t" {
                                                    case "a":
                                                        v22 += v22
                                                    case "hello":
                                                        var v32 bool = true
                                                        _ = v32
                                                    case "\n\t":
                                                        var v33 int64 = (^(v22))
                                                        v34 := GenS6{F0: 0}
                                                        _ = v33
                                                        _ = v34
                                                    }
                                                default:
                                                    v22 = v22
                                                    fmt.Println(v21, v1, fi11)
                                                    if true {
                                                        v35 := GenS4{F0: 869}
                                                        var v36 uint32 = 2
                                                        v1 -= (v1 ^ (v1 - -32768))
                                                        _ = v35
                                                        _ = v36
                                                    } else {
                                                        v37 := GenS5{F0: v1, F1: 127}
                                                        _ = v37
                                                    }
                                                }
                                                var v38 float32 = 3.14159
                                                var v39 int32 = -869
                                                _ = v38
                                                _ = v39
                                            }
                                        }
                                        v40 := GenS1{F0: 42, F1: (v22 - -819), F2: v22}
                                        _ = v22
                                        _ = v40
                                    }
                                    _ = v21
                                }
                                v41 := GenS1{F0: 93, F1: -9223372036854775808, F2: 42}
                                var v42 float32 = 0.1
                                fmt.Println(v42, fi11)
                                _ = v17
                                _ = v41
                                _ = v42
                            default:
                                v1 = v1
                            }
                            var v43 string = "\n\t"
                            _ = v43
                        }
                        var v44 int16 = 0
                        _ = v44
                    }
                }
            }
            var v45 int = fi11
            v45 -= fi12
            _ = fi12
            _ = v45
        }
        for fi46 := 0; fi46 < 17; fi46++ {
            fmt.Println(fi11, v1, fi46)
            v47 := GenS6{F0: 2}
            if true {
                var v48 float64 = 1e10
                if false {
                    fmt.Println(v48)
                    v49 := GenS5{F0: 32767, F1: 114}
                    v50 := GenS7{F0: (false || (true || false))}
                    _ = v49
                    _ = v50
                } else {
                    switch fi46 {
                    case 9223372036854775807:
                        if len(v2) > 0 {
                            v2[0] = "hello"
                        }
                    case 32768:
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    case -1:
                        switch "foo bar" {
                        case "":
                            v1 = v1
                        }
                    }
                    v3["hello"] = ("foo bar" + ("foo bar" + ""))
                }
                _ = v48
            } else {
                delete(v3, "foo bar")
            }
            fmt.Println(fi46, fi11)
            _ = fi46
            _ = v47
        }
        _ = fi11
    }
    if len(v4) > 0 {
        v4[0] = -43
    }
    if (255 >= 2) {
        var v51 uint32 = 631
        _ = v51
    }
    _ = len(v3)
    var v52 int8 = -1
    for fi53 := 0; fi53 < 14; fi53++ {
        v1 -= 0
        var v54 uint = 2
        _ = fi53
        _ = v54
    }
    switch -2147483648 {
    case 2147483647:
        for fi55 := 0; fi55 < 8; fi55++ {
            var v56 int8 = v52
            v57 := GenS1{F0: v56, F1: -1, F2: -1}
            v58 := GenS8{F0: 430, F1: ("hello" + "foo bar"), F2: fi55}
            _ = fi55
            _ = v56
            _ = v57
            _ = v58
        }
    case -64:
        var v59 float32 = float32(0.0)
        _ = v59
    default:
        switch "foo bar" {
        case "foo bar":
            v60 := GenS9{F0: 1, F1: 9223372036854775807}
            _ = v60
        default:
            switch 0 {
            case 1:
                for fi61 := 0; fi61 < 7; fi61++ {
                    if false {
                        fmt.Println(v1)
                        v62 := v3[("foo bar" + "\n\t")]
                        _ = v62
                    } else {
                        fmt.Println(fi61, v52, fi61)
                    }
                    v63 := GenS10{F0: 42, F1: 127, F2: 65535}
                    v52 += (v52 | v52)
                    _ = fi61
                    _ = v63
                }
            case 42:
                v52 -= v52
            case -498:
                for fi64 := 0; fi64 < 12; fi64++ {
                    for fi65 := 0; fi65 < 8; fi65++ {
                        var v66 int32 = 1
                        _ = fi65
                        _ = v66
                    }
                    v67 := GenS8{F0: 255, F1: "foo bar", F2: -1}
                    for fi68 := 0; fi68 < 4; fi68++ {
                        v2 = append(v2, ("foo bar" + (("foo bar" + "hello") + "\n\t")))
                        v1 = (v1 | v1)
                        _ = fi68
                    }
                    _ = fi64
                    _ = v67
                }
            }
        }
    }
    var v69 float32 = 1.5
    fmt.Println(v69, v1)
    var v70 uint8 = 255
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v52:", v52)
    fmt.Println("v69:", v69)
    fmt.Println("v70:", v70)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v52
    _ = v69
    _ = v70
}
