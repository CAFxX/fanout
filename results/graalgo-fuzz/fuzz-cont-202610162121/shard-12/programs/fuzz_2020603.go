package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string; F2 uint8 }
type GenS2 struct { F0 uint16; F1 uint32 }
type GenS3 struct { F0 int16; F1 uint32; F2 uint16 }
type GenS4 struct { F0 int64; F1 int16; F2 int }
type GenS5 struct { F0 int8; F1 uint64 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 uint }

func main() {
    var v1 float32 = (3.14159 + float32(0.1))
    var v2 uint32 = 4294967295
    var v3 bool = false
    switch "a" {
    case "a":
        switch "\n\t" {
        case "a":
            v4 := GenS1{F0: 255, F1: "a", F2: 127}
            _ = v4
        case "":
            var v5 int = 32768
            _ = v5
        case "hello":
            var v6 string = "a"
            _ = v6
        default:
            v1 = v1
        }
    case "\n\t":
        for fi7 := 0; fi7 < 5; fi7++ {
            v2 = (^(v2))
            for fi8 := 0; fi8 < 13; fi8++ {
                if (true || (24 > 0)) {
                    v9 := []int{}
                    v3 = v3
                    v10 := GenS1{F0: uint(v2), F1: ("\n\t" + "\n\t"), F2: 51}
                    _ = v10
                    _ = v9
                } else {
                    v11 := []int{(fi7 | fi8), -944, 128}
                    v12 := GenS2{F0: 255, F1: 21}
                    v3 = ("foo bar" <= "a")
                    _ = v11
                    _ = v12
                }
                fmt.Println(v3)
                v1 = 1.5
                _ = fi8
            }
            _ = fi7
        }
    case "hello":
        v2 = 2
    }
    switch "\n\t" {
    case "":
        if false {
            v13 := GenS2{F0: 488, F1: (^(v2))}
            v14 := GenS3{F0: 479, F1: v2, F2: 418}
            var v15 bool = (2 <= 255)
            _ = v13
            _ = v14
            _ = v15
        }
    case "a":
        switch "" {
        case "\n\t":
            v16 := GenS3{F0: 32767, F1: v2, F2: 2}
            _ = v16
        case "a":
            var v17 uint32 = v2
            _ = v17
        case "hello":
            fmt.Println(v1, v1, v1)
        }
    }
    v18 := []int{-106, -9223372036854775808, -639, 666}
    switch 0 {
    case 32767:
        switch "a" {
        case "hello":
            var v19 uint64 = 531
            _ = v19
        case "a":
            v20 := GenS4{F0: 9223372036854775807, F1: -32768, F2: -1}
            _ = v20
        case "foo bar":
            if len(v18) > 0 {
                v18[0] = 42
            }
        default:
            var v21 int64 = -9223372036854775808
            _ = v21
        }
    }
    var v22 int32 = -2147483648
    v1 = v1
    v18 = append(v18, -9223372036854775808)
    for fi23 := 0; fi23 < 1; fi23++ {
        v2 += v2
        v18 = append(v18, (fi23 - (fi23 / ((fi23) | 1))))
        _ = fi23
    }
    if (!v3) {
        v24 := GenS3{F0: 402, F1: v2, F2: 255}
        v1 = v1
        _ = v24
    }
    for fi25 := 0; fi25 < 15; fi25++ {
        fmt.Println(v22, v1)
        _ = fi25
    }
    var v26 uint = 684
    if (v1 <= v1) {
        switch ("hello" + "a") {
        case "foo bar":
            var v27 uint = (^(v26))
            _ = v27
        default:
            if (!(!true)) {
                v1 += (v1 * v1)
            } else {
                _ = len(v18)
            }
        }
        if (!(42 > -22)) {
            _ = len(v18)
            _ = len(v18)
            switch "\n\t" {
            case "":
                v28 := GenS5{F0: 79, F1: 1}
                _ = v28
            }
        } else {
            v1 = v1
            if true {
                v29 := GenS3{F0: int16(v22), F1: (v2 & 42), F2: 42}
                var v30 int = 9223372036854775807
                _ = v29
                _ = v30
            } else {
                switch -2147483648 {
                case -355:
                    if (("hello" + "\n\t") <= ("\n\t" + "hello")) {
                        switch -759 {
                        case 1:
                            var v31 float32 = v1
                            _ = v31
                        case 32767:
                            var v32 float64 = 1.5
                            _ = v32
                        default:
                            if len(v18) > 0 {
                                v18[0] = -775
                            }
                        }
                        switch "hello" {
                        case "\n\t":
                            var v33 int16 = 1
                            _ = v33
                        case "foo bar":
                            v34 := GenS6{F0: v3}
                            _ = v34
                        case "a":
                            fmt.Println(v3)
                        default:
                            switch "foo bar" {
                            case "foo bar":
                                if ((v2 >> ((v26) & 63)) >= (v2 + 4294967295)) {
                                    v26 = 796
                                } else {
                                    v22 = v22
                                    _ = len(v18)
                                    fmt.Println(v26, v22)
                                }
                            case "":
                                if (("foo bar" + "a") < "hello") {
                                    if (1 < 0) {
                                        if v3 {
                                            var v35 string = "\n\t"
                                            for fi36 := 0; fi36 < 20; fi36++ {
                                                if len(v18) >= 2 {
                                                    v18 = v18[1:len(v18)]
                                                }
                                                _ = fi36
                                            }
                                            _ = v35
                                        } else {
                                            v26 = uint(v2)
                                            if len(v18) >= 2 {
                                                v18 = v18[1:len(v18)]
                                            }
                                            var v37 uint64 = 42
                                            _ = v37
                                        }
                                    }
                                }
                            default:
                                var v38 int64 = -150
                                _ = v38
                            }
                        }
                        var v39 int8 = -107
                        _ = v39
                    }
                case 32768:
                    v40 := GenS7{F0: v26}
                    _ = v40
                default:
                    var v41 uint8 = 54
                    _ = v41
                }
                if v3 {
                    v26 += (v26 ^ 363)
                    v42 := GenS3{F0: -608, F1: (v2 & 65535), F2: 527}
                    _ = v42
                }
            }
        }
    } else {
        for fi43 := 0; fi43 < 5; fi43++ {
            switch "\n\t" {
            case "foo bar":
                fmt.Println(v26, v1, fi43)
            case "\n\t":
                switch -78 {
                case 32767:
                    var v44 uint8 = 128
                    _ = v44
                case 42:
                    for fi45 := 0; fi45 < 6; fi45++ {
                        var v46 bool = (-59 == 0)
                        v46 = (!v46)
                        fmt.Println(v22, v22, v46)
                        _ = fi45
                        _ = v46
                    }
                default:
                    v2 += (v2 + (^(v2)))
                }
            }
            v18 = append(v18, fi43)
            _ = fi43
        }
    }
    for fi47 := 0; fi47 < 11; fi47++ {
        var v48 uint64 = 42
        v1 += v1
        switch fi47 {
        case 580:
            v18 = append(v18, fi47)
        case 4294967295:
            v49 := GenS6{F0: v3}
            _ = v49
        case 328:
            var v50 int16 = 42
            _ = v50
        default:
            v22 += v22
        }
        _ = fi47
        _ = v48
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v18
    _ = v2
    _ = v22
    _ = v26
    _ = v3
}
