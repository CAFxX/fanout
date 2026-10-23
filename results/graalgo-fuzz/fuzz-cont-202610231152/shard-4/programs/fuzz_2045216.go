package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8; F2 bool }
type GenS2 struct { F0 int32; F1 uint16 }
type GenS3 struct { F0 int16; F1 uint; F2 int }
type GenS4 struct { F0 uint64; F1 int; F2 int16 }
type GenS5 struct { F0 int32; F1 uint }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 string; F1 int8 }
type GenS8 struct { F0 int8; F1 int }
type GenS9 struct { F0 uint }

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[int]bool)
    var v3 string = "a"
    v4 := make(map[string]bool)
    if (1.5 >= 0.1) {
        if true {
            var v5 int = 32767
            _ = v5
        }
        var v6 int16 = 32767
        _ = v6
    }
    v7 := GenS1{F0: 128, F1: 0, F2: false}
    switch "a" {
    case "\n\t":
        if (1e10 >= 1.5) {
            if false {
                for fi8 := 0; fi8 < 18; fi8++ {
                    v3 = ("a" + v3)
                    _ = fi8
                }
                for fi9 := 0; fi9 < 2; fi9++ {
                    fmt.Println(v3, v3)
                    _ = fi9
                }
                var v10 float64 = 3.14159
                _ = v10
            }
            fmt.Println(v3)
        }
    case "hello":
        v3 = v3
    }
    if (2 > 605) {
        fmt.Println(v3)
        switch v3 {
        case "foo bar":
            var v11 string = ("hello" + "foo bar")
            _ = v11
        case "a":
            v3 = v3
        case "":
            for fi12 := 0; fi12 < 9; fi12++ {
                v13 := GenS1{F0: 4294967295, F1: 1, F2: (!((^(fi12)) < -820))}
                v14 := GenS2{F0: -1, F1: 65535}
                _ = fi12
                _ = v13
                _ = v14
            }
        default:
            var v15 int64 = -182
            _ = v15
        }
        v16 := GenS1{F0: 127, F1: 1, F2: (!false)}
        _ = v16
    } else {
        if true {
            v3 = v3
            switch v3 {
            case "hello":
                for fi17 := 0; fi17 < 6; fi17++ {
                    switch fi17 {
                    case 3:
                        v18 := GenS3{F0: -817, F1: 255, F2: (fi17 & fi17)}
                        _ = v18
                    case 0:
                        v19 := []int{-661, fi17, fi17, (^(fi17))}
                        _ = v19
                    }
                    v3 = v3
                    _ = fi17
                }
            }
            for fi20 := 0; fi20 < 2; fi20++ {
                switch v3 {
                case "a":
                    v3 = ("" + v3)
                case "\n\t":
                    if false {
                        var v21 uint = 2
                        var v22 int8 = 0
                        _ = v21
                        _ = v22
                    }
                case "foo bar":
                    v23 := GenS4{F0: 2, F1: (fi20 % (((fi20 + fi20)) | 1)), F2: 804}
                    _ = v23
                default:
                    var v24 int32 = -2147483648
                    _ = v24
                }
                fmt.Println(v3, fi20)
                switch fi20 {
                case 171:
                    switch v3 {
                    case "hello":
                        var v25 uint64 = 42
                        _ = v25
                    }
                case 32767:
                    for fi26 := 0; fi26 < 4; fi26++ {
                        fmt.Println(v3, fi26, fi20)
                        fmt.Println(v3)
                        _ = fi26
                    }
                }
                _ = fi20
            }
        }
        v27 := GenS2{F0: -1, F1: 351}
        _ = v27
    }
    var v28 uint64 = 4294967295
    v3 = v3
    v29 := GenS2{F0: -2147483648, F1: 255}
    var v30 bool = true
    fmt.Println(v28, v28, v30)
    if (42 == 1) {
        for fi31 := 0; fi31 < 10; fi31++ {
            v32 := GenS5{F0: -2147483648, F1: 65535}
            _ = fi31
            _ = v32
        }
    }
    if v30 {
        if v30 {
            v28 -= 1
            var v33 int32 = -2147483648
            _ = v33
        } else {
            var v34 int32 = -679
            if (2 < 2147483647) {
                var v35 uint = 0
                v30 = (993 >= -1)
                _ = v35
            }
            var v36 uint16 = 65535
            _ = v34
            _ = v36
        }
        v37 := []int{42, 4294967295, 2147483647, 1}
        _ = v37
    }
    v38 := GenS6{F0: 42}
    for fi39 := 0; fi39 < 8; fi39++ {
        if v30 {
            for fi40 := 0; fi40 < 13; fi40++ {
                if ((!((v28 >> ((uint(v28)) & 63)) != v28)) && ((fi40 * fi39) != 203)) {
                    switch v3 {
                    case "hello":
                        if v30 {
                            v41 := GenS2{F0: 1, F1: 1}
                            v3 = "foo bar"
                            var v42 int32 = -530
                            _ = v41
                            _ = v42
                        }
                    case "foo bar":
                        var v43 int32 = -1
                        _ = v43
                    case "\n\t":
                        fmt.Println(v28, fi40)
                    }
                }
                v30 = v30
                v44 := GenS7{F0: (v3 + v3), F1: -1}
                _ = fi40
                _ = v44
            }
            for fi45 := 0; fi45 < 1; fi45++ {
                v28 += v28
                v46 := GenS8{F0: 104, F1: fi39}
                _ = fi45
                _ = v46
            }
            v47 := GenS7{F0: v3, F1: 42}
            _ = v47
        } else {
            for fi48 := 0; fi48 < 5; fi48++ {
                fmt.Println(v28, fi39, fi39)
                _ = fi48
            }
            for fi49 := 0; fi49 < 12; fi49++ {
                var v50 int16 = 1000
                if (730 <= 9223372036854775807) {
                    var v51 bool = (2 < 65535)
                    var v52 float32 = 1.5
                    if v51 {
                        var v53 int16 = v50
                        if v30 {
                            v54 := GenS9{F0: 2}
                            v53 -= int16(v28)
                            _ = v54
                        } else {
                            v51 = (32767 >= 891)
                        }
                        v55 := GenS3{F0: v50, F1: 1, F2: (fi49 >> ((719) & 63))}
                        var v56 float64 = 0.1
                        _ = v53
                        _ = v55
                        _ = v56
                    } else {
                        switch "" {
                        case "":
                            var v57 float64 = (0.1 + -2.25)
                            _ = v57
                        default:
                            fmt.Println(v30, v52, fi49)
                        }
                    }
                    _ = v51
                    _ = v52
                }
                _ = fi49
                _ = v50
            }
        }
        var v58 int16 = 42
        var v59 float64 = float64(1.5)
        _ = fi39
        _ = v58
        _ = v59
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v38
    _ = v4
    _ = v7
}
