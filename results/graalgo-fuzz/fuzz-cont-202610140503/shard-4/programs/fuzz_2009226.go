package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
type GenS2 struct { F0 int16; F1 int8; F2 uint16 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint64; F1 int; F2 uint8 }
type GenS5 struct { F0 string; F1 int; F2 uint32 }
type GenS6 struct { F0 int32; F1 bool }

func main() {
    var v1 int8 = 127
    v2 := make(map[uint32]uint32)
    v3 := make(map[string]bool)
    var v4 bool = false
    for fi5 := 0; fi5 < 11; fi5++ {
        if v4 {
            v6 := GenS1{F0: fi5, F1: 344}
            _ = v6
        } else {
            v1 -= (^(v1))
            for fi7 := 0; fi7 < 11; fi7++ {
                if (v4 || true) {
                    if v4 {
                        var v8 int = fi5
                        _ = v8
                    }
                    var v9 string = "a"
                    _ = v9
                } else {
                    v10 := GenS1{F0: 1, F1: 2147483648}
                    _ = v10
                }
                if false {
                    switch "\n\t" {
                    case "hello":
                        var v11 uint16 = 460
                        _ = v11
                    case "\n\t":
                        v12 := GenS1{F0: 319, F1: 4294967295}
                        _ = v12
                    case "foo bar":
                        switch "a" {
                        case "hello":
                            v13 := v2[356]
                            _ = v13
                        default:
                            if (!(true || ((v4 && (359 <= 1)) || (1 <= 0)))) {
                                var v14 uint64 = 1
                                v1 += v1
                                v4 = false
                                _ = v14
                            } else {
                                v15 := []int{-1, (fi5 >> ((454) & 63)), (-(fi7)), fi5}
                                v16 := GenS2{F0: -1, F1: v1, F2: 0}
                                var v17 bool = (!false)
                                _ = v15
                                _ = v16
                                _ = v17
                            }
                        }
                    default:
                        var v18 uint64 = 2
                        _ = v18
                    }
                } else {
                    fmt.Println(v1, v4, fi7)
                    switch (fi5 / ((fi7) | 1)) {
                    case 32767:
                        fmt.Println(v4)
                    case 32768:
                        v1 += (v1 / (((v1 | -118)) | 1))
                    case -626:
                        v1 = v1
                    default:
                        var v19 uint64 = 1
                        _ = v19
                    }
                    switch (^(fi7)) {
                    case 0:
                        v20 := []int{0, fi7, (fi7 ^ (fi7 >> ((65535) & 63)))}
                        _ = v20
                    case 127:
                        v21 := []int{(^(fi7)), (fi7 % ((9223372036854775807) | 1)), (fi7 % ((fi5) | 1))}
                        _ = v21
                    case 1:
                        v1 = int8(fi5)
                    }
                }
                _ = fi7
            }
            v22 := GenS3{F0: 173}
            _ = v22
        }
        v1 -= v1
        switch "a" {
        case "\n\t":
            var v23 int32 = -1
            _ = v23
        }
        _ = fi5
    }
    v24 := GenS2{F0: 32767, F1: v1, F2: 503}
    switch -2147483648 {
    case 844:
        fmt.Println(v4, v1)
    case -2147483648:
        switch "\n\t" {
        case "\n\t":
            for fi25 := 0; fi25 < 5; fi25++ {
                v26 := GenS4{F0: 1, F1: 128, F2: 119}
                _ = fi25
                _ = v26
            }
        case "a":
            if v4 {
                var v27 string = "\n\t"
                v3["\n\t"] = ((229 != 936) && ((false || v4) && v4))
                switch "\n\t" {
                case "foo bar":
                    var v28 uint64 = 0
                    _ = v28
                case "\n\t":
                    var v29 uint32 = 2147483647
                    _ = v29
                }
                _ = v27
            }
        }
    }
    if v4 {
        var v30 float64 = (-2.25 * (((3.14159 - ((1e10 / 1.0) * -2.25)) * 3.14159) - 1e10))
        var v31 float64 = 0.0
        _ = v30
        _ = v31
    } else {
        if v4 {
            if v4 {
                switch "" {
                case "hello":
                    v1 -= v1
                case "foo bar":
                    var v32 bool = v4
                    _ = v32
                case "a":
                    if (0.0 < (1e10 / 3.14159)) {
                        fmt.Println(v4, v4)
                        var v33 uint64 = 166
                        if v4 {
                            var v34 uint16 = 70
                            for fi35 := 0; fi35 < 1; fi35++ {
                                var v36 int = fi35
                                for fi37 := 0; fi37 < 17; fi37++ {
                                    var v38 int = -404
                                    _ = fi37
                                    _ = v38
                                }
                                if false {
                                    var v39 float32 = 0.1
                                    v40 := []int{664, v36, (^(fi35))}
                                    if (0 <= 42) {
                                        if len(v40) > 0 {
                                            v40[0] = -767
                                        }
                                        v41 := GenS5{F0: "foo bar", F1: fi35, F2: 0}
                                        _ = v41
                                    }
                                    _ = v39
                                    _ = v40
                                } else {
                                    v34 = (v34 << ((18446744073709551615) & 63))
                                }
                                _ = fi35
                                _ = v36
                            }
                            _ = v34
                        } else {
                            var v42 int64 = -1
                            _ = v42
                        }
                        _ = v33
                    }
                }
            }
            v1 = v1
        } else {
            v43 := GenS3{F0: 2}
            _ = v43
        }
        v44 := []int{32768}
        v1 = v1
        _ = v44
    }
    for fi45 := 0; fi45 < 17; fi45++ {
        var v46 int8 = (-(v1))
        for fi47 := 0; fi47 < 20; fi47++ {
            var v48 string = ("foo bar" + (("\n\t" + ("a" + ("" + ""))) + "hello"))
            _ = fi47
            _ = v48
        }
        _ = fi45
        _ = v46
    }
    fmt.Println(v1, v1)
    v49 := GenS6{F0: -98, F1: v4}
    v1 += v1
    var v50 uint32 = 724
    var v51 float64 = (float64(1e10) / (1e10 - (3.14159 * (0.0 / 1.5))))
    var v52 int8 = v1
    var v53 uint8 = 2
    v4 = (!(((65535 < 1) && (0 >= 34)) || (!(42 < 1))))
    switch 32768 {
    case 32768:
        var v54 int64 = -9223372036854775808
        _ = v54
    case 9223372036854775807:
        for fi55 := 0; fi55 < 19; fi55++ {
            switch ("" + "foo bar") {
            case "hello":
                var v56 float64 = 0.0
                _ = v56
            default:
                v51 += 0.1
            }
            for fi57 := 0; fi57 < 5; fi57++ {
                v53 -= (v53 - v53)
                v58 := GenS4{F0: 6, F1: (-(fi57)), F2: v53}
                _ = fi57
                _ = v58
            }
            var v59 uint32 = (v50 & 2147483648)
            _ = fi55
            _ = v59
        }
    default:
        fmt.Println(v52, v51, v52)
    }
    v52 = v52
    var v60 uint64 = uint64(v50)
    v61 := []int{2147483647}
    var v62 int32 = -2147483648
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    fmt.Println("v60:", v60)
    fmt.Println("slicelen v61:", len(v61), v61)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v49
    _ = v50
    _ = v51
    _ = v52
    _ = v53
    _ = v60
    _ = v61
    _ = v62
}
