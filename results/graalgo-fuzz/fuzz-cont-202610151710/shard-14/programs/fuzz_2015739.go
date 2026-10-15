package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int; F1 string }
type GenS3 struct { F0 uint32; F1 bool; F2 uint16 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint32; F1 uint64; F2 uint32 }
type GenS6 struct { F0 int8; F1 uint; F2 string }
type GenS7 struct { F0 int64; F1 uint16 }
type GenS8 struct { F0 uint; F1 uint }
func genFn0(p0 bool, p1 uint16) uint16 {
    v5 := GenS1{F0: 1}
    var v6 uint16 = p1
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 42
}
func genFn1() int {
    for fi7 := 0; fi7 < 20; fi7++ {
        if false {
            for fi8 := 0; fi8 < 18; fi8++ {
                switch "a" {
                case "":
                    var v9 uint = 0
                    var v10 uint16 = 33
                    v11 := []int{-2147483648, fi7, fi8}
                    _ = v10
                    _ = v11
                    _ = v9
                case "hello":
                    var v12 float32 = 0.1
                    _ = v12
                case "foo bar":
                    v13 := GenS2{F0: (fi7 << ((905) & 63)), F1: ("\n\t" + ("\n\t" + "foo bar"))}
                    _ = v13
                default:
                    fmt.Println(fi7)
                }
                _ = fi8
            }
            var v14 int16 = -1
            switch ("hello" + ("a" + "hello")) {
            case "foo bar":
                for fi15 := 0; fi15 < 14; fi15++ {
                    switch ("a" + ("hello" + ("hello" + "a"))) {
                    case "a":
                        var v16 uint16 = 168
                        _ = v16
                    case "hello":
                        switch fi15 {
                        case 42:
                            switch (fi7 * fi7) {
                            case 32767:
                                switch 4294967295 {
                                case 302:
                                    switch ("hello" + "foo bar") {
                                    case "hello":
                                        var v17 uint = 762
                                        _ = v17
                                    default:
                                        fmt.Println(fi7, fi15, fi15)
                                    }
                                    v18 := GenS3{F0: 281, F1: false, F2: 2}
                                    v19 := GenS2{F0: fi15, F1: ("foo bar" + (("\n\t" + "hello") + ("\n\t" + "")))}
                                    _ = v18
                                    _ = v19
                                default:
                                    fmt.Println(fi7, v14)
                                }
                            case -80:
                                v20 := GenS3{F0: 2147483648, F1: (true || true), F2: 32767}
                                _ = v20
                            default:
                                v21 := []int{(fi15 ^ 108), fi15, 0}
                                _ = v21
                            }
                        case -667:
                            var v22 string = "foo bar"
                            _ = v22
                        case 0:
                            fmt.Println(fi7, v14)
                        }
                    default:
                        v14 = -1
                        v23 := []int{(fi7 / (((fi15 & fi15)) | 1)), -313, (fi7 & fi7)}
                        _ = v23
                    }
                    v14 = (v14 / (((^(v14))) | 1))
                    v24 := GenS4{F0: (v14 ^ v14)}
                    _ = fi15
                    _ = v24
                }
            case "hello":
                for fi25 := 0; fi25 < 11; fi25++ {
                    v26 := GenS4{F0: (v14 / ((v14) | 1))}
                    var v27 uint64 = 408
                    _ = fi25
                    _ = v26
                    _ = v27
                }
            case "a":
                v14 += (v14 / (((v14 & -705)) | 1))
            default:
                v14 -= 1
            }
            _ = v14
        }
        _ = fi7
    }
    if (0 < 1) {
        v28 := []int{-94, 834, 42, 32767}
        switch 795 {
        case -2147483648:
            var v29 int32 = 306
            _ = v29
        default:
            v30 := GenS1{F0: 9223372036854775807}
            _ = v30
        }
        for fi31 := 0; fi31 < 4; fi31++ {
            var v32 uint32 = 42
            if true {
                for fi33 := 0; fi33 < 15; fi33++ {
                    switch ("" + "hello") {
                    case "hello":
                        if (!(-1 == 42)) {
                            fmt.Println(fi31, fi33)
                        }
                    case "\n\t":
                        var v34 int32 = -2147483648
                        _ = v34
                    }
                    _ = fi33
                }
                fmt.Println(fi31)
            }
            _ = fi31
            _ = v32
        }
        _ = v28
    } else {
        var v35 int64 = 42
        for fi36 := 0; fi36 < 18; fi36++ {
            for fi37 := 0; fi37 < 14; fi37++ {
                v35 = (v35 | (v35 ^ 0))
                _ = fi37
            }
            var v38 int8 = -128
            var v39 uint = 18446744073709551615
            _ = fi36
            _ = v38
            _ = v39
        }
        if true {
            fmt.Println(v35)
            var v40 bool = false
            _ = v40
        } else {
            var v41 int16 = 1
            _ = v41
        }
        _ = v35
    }
    var v42 float32 = float32((float64(((float32(-2.25) + 3.14159) * 3.14159)) - 1e10))
    for fi43 := 0; fi43 < 15; fi43++ {
        v44 := GenS1{F0: 42}
        v45 := GenS5{F0: 2147483647, F1: 18446744073709551615, F2: 0}
        _ = fi43
        _ = v44
        _ = v45
    }
    _ = v42
    return 2147483647
}

func main() {
    v1 := make(map[string]int)
    v2 := []int{-78, 32768, 845}
    v3 := []int{128, 1}
    var v4 float32 = float32(1.5)
    var v46 uint64 = 117
    v46 += 649
    v47 := v1["a"]
    if true {
        var v48 float64 = 0.0
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v48 = ((0.0 + (1.5 + v48)) / v48)
        _ = v48
    }
    var v49 int32 = -892
    var v50 float32 = -2.25
    v46 -= (v46 / ((v46) | 1))
    v46 = v46
    delete(v1, "\n\t")
    if true {
        v51 := GenS6{F0: 0, F1: 255, F2: "hello"}
        v49 += v49
        var v52 int64 = 0
        _ = v51
        _ = v52
    }
    v49 = v49
    v53 := GenS7{F0: -9223372036854775808, F1: 666}
    v54 := GenS8{F0: 709, F1: 18446744073709551615}
    switch "a" {
    case "foo bar":
        for fi55 := 0; fi55 < 14; fi55++ {
            var v56 bool = (255 == 95)
            v57 := GenS8{F0: 0, F1: 65535}
            _ = fi55
            _ = v56
            _ = v57
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v49:", v49)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v47
    _ = v49
    _ = v50
    _ = v53
    _ = v54
}
