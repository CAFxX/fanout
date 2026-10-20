package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 int64 }
type GenS2 struct { F0 string; F1 string }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint; F1 uint32; F2 int }
type GenS5 struct { F0 bool; F1 uint; F2 int64 }

func main() {
    var v1 int16 = 1000
    var v2 bool = (1 != 1)
    var v3 uint8 = 255
    switch "foo bar" {
    case "hello":
        if ((1.5 + 1e10) < (-2.25 - float32(float64(0.1)))) {
            v4 := []int{-785}
            switch 2147483647 {
            case 42:
                fmt.Println(v2, v3)
            case -884:
                fmt.Println(v2)
            case -2147483648:
                switch "hello" {
                case "a":
                    fmt.Println(v3, v3)
                case "hello":
                    v5 := GenS1{F0: 765, F1: 2147483647, F2: -9223372036854775808}
                    _ = v5
                }
            }
            _ = v4
        }
    }
    if (1 > uint64(v3)) {
        if (!(!false)) {
            v3 = v3
        }
    } else {
        if (418 <= -1) {
            var v6 int32 = 1
            _ = v6
        } else {
            if (!((true && (4294967295 >= -2147483648)) && (!(!(!v2))))) {
                fmt.Println(v2, v3, v3)
                var v7 uint = 716
                var v8 float32 = -2.25
                _ = v7
                _ = v8
            }
            switch "" {
            case "\n\t":
                var v9 string = "a"
                _ = v9
            case "hello":
                v1 += v1
            case "a":
                if true {
                    switch 32767 {
                    case -1:
                        v1 = -484
                    case -341:
                        fmt.Println(v2, v1)
                    case 78:
                        if v2 {
                            var v10 uint = 1
                            _ = v10
                        }
                    default:
                        var v11 int64 = 178
                        _ = v11
                    }
                    for fi12 := 0; fi12 < 18; fi12++ {
                        v13 := []int{1, fi12, (fi12 & fi12)}
                        if v2 {
                            var v14 float32 = 3.14159
                            if len(v13) > 0 {
                                v13[0] = (fi12 + fi12)
                            }
                            _ = v14
                        }
                        _ = fi12
                        _ = v13
                    }
                    var v15 uint = 255
                    _ = v15
                }
            default:
                var v16 int16 = (v1 << ((1) & 63))
                _ = v16
            }
        }
        v17 := GenS2{F0: ("\n\t" + "\n\t"), F1: "foo bar"}
        v18 := []int{397, 2147483647, 797}
        _ = v17
        _ = v18
    }
    v19 := GenS3{F0: (v3 - v3)}
    switch "foo bar" {
    case "a":
        if false {
            v20 := []int{9223372036854775807}
            _ = v20
        } else {
            if (1 > 128) {
                v3 += v3
                fmt.Println(v3, v2)
            } else {
                v2 = (0.1 <= ((3.14159 / 1.5) - -2.25))
                fmt.Println(v3, v2)
            }
            var v21 string = ("a" + ("foo bar" + ("\n\t" + "hello")))
            _ = v21
        }
    case "foo bar":
        switch ("hello" + ("" + "foo bar")) {
        case "a":
            if v2 {
                v22 := []int{}
                v3 -= v3
                _ = v22
            } else {
                v23 := []int{565}
                _ = v23
            }
        default:
            var v24 uint = 1
            _ = v24
        }
    default:
        var v25 uint8 = (v3 % ((127) | 1))
        _ = v25
    }
    switch -708 {
    case -975:
        for fi26 := 0; fi26 < 12; fi26++ {
            v27 := []int{(^(fi26))}
            v3 += (v3 * v3)
            _ = fi26
            _ = v27
        }
    case 834:
        v28 := []int{4294967295, 4294967295}
        _ = v28
    }
    if v2 {
        v3 += (v3 | v3)
        var v29 float64 = 3.14159
        var v30 int64 = 1
        _ = v29
        _ = v30
    } else {
        if true {
            var v31 int32 = 2147483647
            _ = v31
        }
        var v32 float64 = 1e10
        _ = v32
    }
    if false {
        v1 += v1
        fmt.Println(v3, v3)
    }
    if (654 < -232) {
        fmt.Println(v1, v1)
        if v2 {
            fmt.Println(v1, v2, v3)
            if v2 {
                switch (("\n\t" + "") + ("a" + (("" + "a") + (("a" + "foo bar") + "hello")))) {
                case "hello":
                    v33 := GenS4{F0: 18446744073709551615, F1: 42, F2: 0}
                    _ = v33
                case "a":
                    for fi34 := 0; fi34 < 12; fi34++ {
                        fmt.Println(v1, fi34)
                        v1 += (v1 | (^(v1)))
                        _ = fi34
                    }
                }
                switch 128 {
                case 42:
                    fmt.Println(v3, v3, v3)
                case -2147483648:
                    v35 := []int{}
                    _ = v35
                case 207:
                    var v36 uint32 = 0
                    _ = v36
                default:
                    v37 := GenS2{F0: "\n\t", F1: ("foo bar" + "hello")}
                    _ = v37
                }
                v38 := GenS3{F0: 246}
                _ = v38
            } else {
                if (42 >= 0) {
                    var v39 uint32 = 415
                    v2 = (2147483648 > v39)
                    fmt.Println(v3, v1)
                    _ = v39
                } else {
                    if (638 != 2) {
                        v3 -= (v3 / ((v3) | 1))
                    }
                    var v40 string = ""
                    if false {
                        var v41 string = v40
                        v42 := []int{-102, 4294967295, -1}
                        _ = v41
                        _ = v42
                    }
                    _ = v40
                }
                var v43 uint32 = 4294967295
                if (int(v43) >= 9223372036854775807) {
                    v44 := []int{597, -2147483648}
                    switch "" {
                    case "":
                        var v45 uint = 319
                        _ = v45
                    case "\n\t":
                        var v46 uint32 = 42
                        _ = v46
                    case "foo bar":
                        v47 := GenS5{F0: ((v2 && v2) && (1e10 != 1.5)), F1: 255, F2: -1}
                        _ = v47
                    }
                    _ = v44
                } else {
                    v48 := []int{1, 128, 128}
                    for fi49 := 0; fi49 < 20; fi49++ {
                        for fi50 := 0; fi50 < 15; fi50++ {
                            switch (fi49 << ((0) & 63)) {
                            case 9223372036854775807:
                                if (1 > 65535) {
                                    v43 += v43
                                    v3 += v3
                                    var v51 int64 = -665
                                    v51 = (-(v51))
                                    v48 = append(v48, (^(fi49)))
                                    _ = v51
                                }
                            }
                            _ = fi50
                        }
                        v43 -= 0
                        _ = fi49
                    }
                    var v52 int8 = 1
                    _ = v48
                    _ = v52
                }
                _ = v43
            }
            v53 := GenS2{F0: "", F1: ""}
            _ = v53
        } else {
            fmt.Println(v2, v2, v1)
        }
        if false {
            if (18446744073709551615 >= 42) {
                for fi54 := 0; fi54 < 3; fi54++ {
                    v3 += (v3 >> ((143) & 63))
                    var v55 int16 = -482
                    _ = fi54
                    _ = v55
                }
            } else {
                v1 += (^(v1))
            }
        }
    }
    fmt.Println(v1)
    fmt.Println(v2)
    v56 := GenS2{F0: "", F1: ("\n\t" + (("\n\t" + "a") + ""))}
    v57 := GenS3{F0: 2}
    v3 -= v3
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v56
    _ = v57
}
