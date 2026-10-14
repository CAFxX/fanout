package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int; F1 int64 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int8; F1 int64; F2 int8 }

func main() {
    var v1 float64 = -2.25
    var v2 uint32 = 203
    var v3 string = ""
    var v4 string = ("" + (v3 + (v3 + v3)))
    if true {
        switch v4 {
        case "hello":
            v5 := []int{553}
            _ = v5
        case "\n\t":
            fmt.Println(v2, v2)
        default:
            switch v3 {
            case "":
                var v6 float64 = 0.0
                _ = v6
            case "foo bar":
                switch v4 {
                case "a":
                    var v7 int8 = -86
                    _ = v7
                case "foo bar":
                    switch v3 {
                    case "foo bar":
                        for fi8 := 0; fi8 < 8; fi8++ {
                            for fi9 := 0; fi9 < 9; fi9++ {
                                var v10 int8 = 100
                                _ = fi9
                                _ = v10
                            }
                            _ = fi8
                        }
                    case "":
                        v11 := GenS1{F0: (v4 == v4), F1: 209}
                        _ = v11
                    }
                }
            default:
                if true {
                    v1 += 1e10
                    for fi12 := 0; fi12 < 7; fi12++ {
                        v3 = "a"
                        fmt.Println(v2, v3)
                        v1 -= (1e10 + -2.25)
                        _ = fi12
                    }
                } else {
                    v13 := GenS2{F0: 763}
                    switch 32767 {
                    case 1:
                        fmt.Println(v3)
                    case -1:
                        if true {
                            if (40 < 42) {
                                fmt.Println(v2)
                                var v14 uint32 = (^(v2))
                                var v15 bool = (!(false && (2 > 4294967295)))
                                _ = v14
                                _ = v15
                            }
                            if true {
                                fmt.Println(v2, v3)
                            } else {
                                if (-2.25 <= ((float64((float32(((float64(1.5) / ((0.0 + (float64(0.1) / v1)) - v1)) * 1e10)) / 1.5)) + v1) - (float64((0.0 * 0.0)) * (v1 - float64((0.0 / 1e10)))))) {
                                    switch -1 {
                                    case -9223372036854775808:
                                        v4 = ("" + ("a" + v4))
                                    case 1:
                                        for fi16 := 0; fi16 < 16; fi16++ {
                                            v17 := GenS3{F0: v2}
                                            v2 -= (v2 | v2)
                                            v18 := GenS2{F0: 4294967295}
                                            _ = fi16
                                            _ = v17
                                            _ = v18
                                        }
                                    case 128:
                                        v2 = (v2 >> ((42) & 63))
                                    default:
                                        v4 = ""
                                    }
                                    var v19 uint32 = v2
                                    v2 -= 0
                                    _ = v19
                                }
                            }
                        }
                    case 128:
                        v20 := []int{32767, 32768, -954, 9223372036854775807}
                        _ = v20
                    default:
                        for fi21 := 0; fi21 < 5; fi21++ {
                            v2 = 76
                            _ = fi21
                        }
                    }
                    _ = v13
                }
            }
        }
        v1 = ((v1 * 0.0) - ((3.14159 / v1) + v1))
    }
    v22 := GenS4{F0: 618, F1: 1}
    v4 = v4
    v23 := GenS5{F0: (!(-1 <= 1))}
    fmt.Println(v1, v1, v3)
    if false {
        v3 = (v4 + v3)
        var v24 uint32 = v2
        _ = v24
    } else {
        v25 := []int{-501}
        if true {
            _ = len(v25)
            for fi26 := 0; fi26 < 16; fi26++ {
                if false {
                    var v27 uint8 = 255
                    v3 = "foo bar"
                    var v28 uint32 = v2
                    _ = v27
                    _ = v28
                } else {
                    for fi29 := 0; fi29 < 8; fi29++ {
                        var v30 int = fi29
                        switch "" {
                        case "":
                            v30 = (fi29 >> ((2) & 63))
                        default:
                            var v31 float32 = -2.25
                            _ = v31
                        }
                        _ = fi29
                        _ = v30
                    }
                }
                v1 -= v1
                var v32 bool = false
                _ = fi26
                _ = v32
            }
            for fi33 := 0; fi33 < 4; fi33++ {
                if len(v25) >= 2 {
                    v25 = v25[1:len(v25)]
                }
                _ = fi33
            }
        }
        v2 -= (v2 ^ 662)
        _ = v25
    }
    var v34 uint16 = 255
    v34 += (v34 >> ((0) & 63))
    if false {
        v34 -= v34
    }
    v3 = v3
    for fi35 := 0; fi35 < 4; fi35++ {
        if ((v1 != -2.25) || (!false)) {
            for fi36 := 0; fi36 < 3; fi36++ {
                switch (^(fi36)) {
                case 32768:
                    v4 = v4
                case 618:
                    v37 := []int{(fi35 << ((42) & 63)), -799}
                    _ = v37
                }
                v38 := GenS6{F0: -128, F1: -652, F2: 127}
                _ = fi36
                _ = v38
            }
        } else {
            switch v3 {
            case "hello":
                for fi39 := 0; fi39 < 13; fi39++ {
                    v40 := []int{(^(fi39)), fi35, (fi39 * fi35), -626}
                    _ = fi39
                    _ = v40
                }
            default:
                for fi41 := 0; fi41 < 10; fi41++ {
                    switch ((("hello" + (("" + v3) + v4)) + v4) + v4) {
                    case "":
                        var v42 int8 = 42
                        _ = v42
                    case "\n\t":
                        if (42 <= 127) {
                            switch v4 {
                            case "foo bar":
                                v43 := GenS1{F0: (4 < 911), F1: 255}
                                _ = v43
                            case "\n\t":
                                switch fi35 {
                                case -751:
                                    var v44 int8 = -1
                                    _ = v44
                                case -2147483648:
                                    var v45 float32 = 3.14159
                                    _ = v45
                                default:
                                    v46 := []int{fi35, fi41}
                                    _ = v46
                                }
                            }
                            var v47 bool = ("hello" > (v3 + ("" + ("" + (((v4 + ("hello" + "a")) + v4) + ("foo bar" + ((v3 + ("hello" + v4)) + v4)))))))
                            v4 = v3
                            _ = v47
                        } else {
                            var v48 uint16 = (v34 * v34)
                            var v49 int8 = -128
                            _ = v48
                            _ = v49
                        }
                    }
                    var v50 uint32 = (^(v2))
                    fmt.Println(v3, fi35)
                    fmt.Println(fi35)
                    switch 718 {
                    case 9223372036854775807:
                        fmt.Println(v34, v34)
                    }
                    fmt.Println(fi41)
                    v34 += (v34 & (v34 * v34))
                    _ = fi41
                    _ = v50
                }
            }
            v51 := []int{}
            _ = v51
        }
        switch "\n\t" {
        case "hello":
            v34 = (v34 / (((^(v34))) | 1))
        }
        _ = fi35
    }
    switch "foo bar" {
    case "hello":
        switch 42 {
        case 32768:
            var v52 float32 = 1.5
            _ = v52
        case -421:
            switch "hello" {
            case "foo bar":
                switch "hello" {
                case "hello":
                    v53 := GenS4{F0: 63, F1: -9223372036854775808}
                    _ = v53
                default:
                    v3 = "a"
                }
            case "\n\t":
                fmt.Println(v1, v4, v1)
            case "":
                v54 := []int{-711}
                _ = v54
            }
        default:
            fmt.Println(v4, v2)
        }
    default:
        var v55 uint16 = (v34 >> ((0) & 63))
        _ = v55
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v34
    _ = v4
}
