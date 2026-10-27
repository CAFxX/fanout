package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16; F2 int64 }
type GenS2 struct { F0 uint8; F1 uint16 }
type GenS3 struct { F0 uint; F1 uint16; F2 uint64 }
type GenS4 struct { F0 int; F1 int32; F2 int16 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int64; F1 int; F2 string }
type GenS7 struct { F0 int8 }

func main() {
    var v1 bool = (!(!true))
    var v2 int64 = 42
    var v3 uint = 255
    if (v1 || true) {
        if true {
            v1 = v1
            v3 = (v3 % ((v3) | 1))
            v4 := []int{9223372036854775807}
            _ = v4
        } else {
            switch "\n\t" {
            case "a":
                var v5 int16 = 1000
                _ = v5
            }
        }
        switch ("foo bar" + (("hello" + "\n\t") + "foo bar")) {
        case "a":
            var v6 bool = v1
            _ = v6
        default:
            fmt.Println(v2)
        }
        fmt.Println(v1, v2, v1)
    } else {
        v7 := []int{0, 2147483647, 1}
        _ = v7
    }
    for fi8 := 0; fi8 < 17; fi8++ {
        if (!(0 == -491)) {
            v9 := []int{(fi8 ^ fi8), (fi8 >> ((65535) & 63))}
            _ = v9
        }
        var v10 int = fi8
        v1 = v1
        _ = fi8
        _ = v10
    }
    v11 := GenS1{F0: 0, F1: 197, F2: v2}
    for fi12 := 0; fi12 < 11; fi12++ {
        var v13 float64 = 1e10
        switch fi12 {
        case 292:
            switch fi12 {
            case 128:
                for fi14 := 0; fi14 < 20; fi14++ {
                    var v15 float64 = float64((1.5 * 1e10))
                    if (v3 >= v3) {
                        v2 -= (^(v2))
                        v16 := []int{(fi14 * 42)}
                        v2 += v2
                        _ = v16
                    } else {
                        v15 = v15
                        v17 := GenS1{F0: 0, F1: 65535, F2: (-(v2))}
                        var v18 int64 = -1
                        _ = v17
                        _ = v18
                    }
                    _ = fi14
                    _ = v15
                }
            }
        case 128:
            var v19 string = ("foo bar" + (((("\n\t" + "\n\t") + "foo bar") + "hello") + ("" + "a")))
            _ = v19
        case 32767:
            for fi20 := 0; fi20 < 7; fi20++ {
                if v1 {
                    fmt.Println(v1, v13)
                    v21 := []int{fi12, fi12, (fi12 / ((fi12) | 1))}
                    _ = v21
                }
                if v1 {
                    switch ("" + "foo bar") {
                    case "\n\t":
                        v3 += (^(v3))
                    case "a":
                        v22 := GenS1{F0: -128, F1: 863, F2: v2}
                        switch "hello" {
                        case "foo bar":
                            switch -1 {
                            case 32767:
                                switch "\n\t" {
                                case "":
                                    v13 += float64(0.1)
                                default:
                                    v3 += v3
                                }
                                fmt.Println(fi20)
                                v1 = (-636 != -1)
                                var v23 float64 = 0.0
                                _ = v23
                            default:
                                v2 -= v2
                            }
                        case "":
                            v24 := []int{fi20}
                            _ = v24
                        case "a":
                            v3 += (v3 >> ((v3) & 63))
                        default:
                            var v25 int64 = (v2 >> ((v3) & 63))
                            _ = v25
                        }
                        fmt.Println(fi12)
                        _ = v22
                    }
                    v26 := GenS1{F0: -87, F1: 2, F2: 0}
                    switch (fi12 << ((1) & 63)) {
                    case 32768:
                        v13 -= 3.14159
                    case -2147483648:
                        v27 := GenS2{F0: 0, F1: 645}
                        _ = v27
                    default:
                        var v28 int8 = 97
                        _ = v28
                    }
                    _ = v26
                }
                _ = fi20
            }
        default:
            if v1 {
                for fi29 := 0; fi29 < 10; fi29++ {
                    v3 = 724
                    if (0.1 >= 0.1) {
                        v2 = v2
                        fmt.Println(fi12)
                    } else {
                        v1 = (!true)
                        v30 := GenS3{F0: (v3 | (v3 >> (((v3 + (^(v3)))) & 63))), F1: 533, F2: 2}
                        var v31 uint8 = 238
                        _ = v30
                        _ = v31
                    }
                    var v32 uint8 = 53
                    _ = fi29
                    _ = v32
                }
                v33 := []int{}
                if (false && v1) {
                    v34 := GenS4{F0: fi12, F1: 77, F2: -1}
                    _ = v34
                }
                _ = v33
            }
        }
        _ = fi12
        _ = v13
    }
    fmt.Println(v3)
    if true {
        v35 := GenS5{F0: 1}
        _ = v35
    }
    var v36 uint16 = 32767
    for fi37 := 0; fi37 < 13; fi37++ {
        v38 := GenS6{F0: -974, F1: fi37, F2: "a"}
        _ = fi37
        _ = v38
    }
    var v39 bool = v1
    v40 := GenS3{F0: v3, F1: 2, F2: 903}
    switch 42 {
    case -504:
        v1 = ((v3 | (v3 << ((65535) & 63))) != 255)
    case 42:
        fmt.Println(v2, v1, v3)
    case 2147483647:
        if (!v39) {
            v41 := []int{}
            _ = v41
        }
    default:
        fmt.Println(v36, v39, v1)
    }
    var v42 int16 = 121
    var v43 int16 = (v42 | v42)
    for fi44 := 0; fi44 < 12; fi44++ {
        v45 := []int{0, fi44, -644, (-(fi44))}
        switch 32767 {
        case 32768:
            v46 := GenS7{F0: 18}
            _ = v46
        default:
            var v47 uint = (^(v3))
            _ = v47
        }
        for fi48 := 0; fi48 < 7; fi48++ {
            var v49 uint16 = 398
            _ = fi48
            _ = v49
        }
        _ = fi44
        _ = v45
    }
    v42 -= v43
    v50 := []int{335, -2147483648, 4294967295}
    v39 = ((-2.25 + (1e10 * float64(-2.25))) != ((1e10 + 0.1) / ((3.14159 * (float64(-2.25) * (-2.25 + 0.1))) - -2.25)))
    for fi51 := 0; fi51 < 2; fi51++ {
        if v39 {
            var v52 uint = v3
            for fi53 := 0; fi53 < 15; fi53++ {
                _ = len(v50)
                v54 := GenS5{F0: uint8(v36)}
                var v55 int32 = 1
                switch "foo bar" {
                case "":
                    v3 = v3
                default:
                    var v56 uint32 = 829
                    _ = v56
                }
                var v57 uint64 = 60
                v42 += (v42 - (v43 | -867))
                _ = fi53
                _ = v54
                _ = v55
                _ = v57
            }
            _ = v52
        }
        _ = fi51
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v39:", v39)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("slicelen v50:", len(v50), v50)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v36
    _ = v39
    _ = v40
    _ = v42
    _ = v43
    _ = v50
}
