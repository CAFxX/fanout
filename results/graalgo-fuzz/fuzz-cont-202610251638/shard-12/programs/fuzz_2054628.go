package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32 }
type GenS2 struct { F0 uint8; F1 uint32; F2 int64 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint16; F1 int32; F2 int32 }
type GenS5 struct { F0 int32 }
func genFn0() uint32 {
    var v4 float64 = 3.14159
    var v5 float32 = (((((0.1 / (-2.25 * (0.1 + 3.14159))) + (((3.14159 * (0.1 * 3.14159)) * 1e10) + 1e10)) - (0.0 * 3.14159)) / 1.5) + (float32(v4) + (float32((((v4 * v4) - float64((float32(v4) - 3.14159))) * v4)) * ((3.14159 - 0.1) / (3.14159 - (3.14159 + (((1.5 * (0.0 + float32(v4))) + (0.0 * float32((v4 * v4)))) / 1.5)))))))
    v6 := []int{2147483647}
    var v7 uint16 = 255
    var v8 int16 = 32767
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 2147483647
}
func genFn1(p0 uint32, p1 uint) int {
    switch ("foo bar" + "hello") {
    case "":
        for fi9 := 0; fi9 < 8; fi9++ {
            var v10 float64 = -2.25
            if true {
                if true {
                    v11 := []int{-249, (fi9 - fi9), -750}
                    _ = v11
                }
            }
            _ = fi9
            _ = v10
        }
    case "\n\t":
        if (0 > 1) {
            p1 -= (p1 + p1)
        } else {
            var v12 uint32 = genFn0()
            v13 := []int{-222, 338, -365}
            _ = v12
            _ = v13
        }
    }
    var v14 uint16 = 32767
    for fi15 := 0; fi15 < 12; fi15++ {
        for fi16 := 0; fi16 < 9; fi16++ {
            var v17 uint16 = 42
            v17 -= (^(v14))
            v18 := []int{32767, (fi16 << (((p1 >> (((p1 / ((p1) | 1))) & 63))) & 63)), fi16, fi16}
            _ = fi16
            _ = v17
            _ = v18
        }
        fmt.Println(p0, v14)
        switch 0 {
        case 1:
            var v19 int = (fi15 - (fi15 << ((65535) & 63)))
            _ = v19
        case 91:
            v20 := GenS1{F0: "hello", F1: 42}
            _ = v20
        }
        _ = fi15
    }
    for fi21 := 0; fi21 < 19; fi21++ {
        v22 := []int{}
        _ = fi21
        _ = v22
    }
    if (1 < 42) {
        for fi23 := 0; fi23 < 13; fi23++ {
            var v24 bool = ((("foo bar" + "hello") + "") >= ("hello" + ("hello" + "hello")))
            fmt.Println(p0)
            _ = fi23
            _ = v24
        }
    } else {
        v25 := []int{32768, -9223372036854775808, 9223372036854775807, 782}
        fmt.Println(v14, p0)
        _ = v25
    }
    _ = p0
    _ = p1
    _ = v14
    return 9223372036854775807
}

func main() {
    var v1 string = "hello"
    var v2 uint16 = 255
    var v3 uint16 = 825
    var v26 string = ("" + "\n\t")
    var v27 uint32 = 430
    v2 -= 2
    if ((!(-128 == 127)) && (1 > 299)) {
        for fi28 := 0; fi28 < 5; fi28++ {
            fmt.Println(fi28, v2, v26)
            _ = fi28
        }
        fmt.Println(v27, v26, v1)
    }
    for fi29 := 0; fi29 < 8; fi29++ {
        fmt.Println(fi29, v26)
        _ = fi29
    }
    if true {
        v30 := GenS1{F0: ((v26 + (("\n\t" + "a") + "foo bar")) + ("" + v26)), F1: -1}
        switch 127 {
        case 128:
            v31 := GenS1{F0: ("a" + (((v1 + v1) + v26) + v1)), F1: -1}
            _ = v31
        case 1:
            if (255 == 92) {
                v3 += (v2 | v2)
            } else {
                for fi32 := 0; fi32 < 4; fi32++ {
                    v33 := []int{genFn1(genFn0(), 255), genFn1(v27, 449), genFn1((v27 & v27), 255)}
                    _ = fi32
                    _ = v33
                }
                v34 := GenS2{F0: 26, F1: v27, F2: 795}
                if ((346 != (^(v3))) && (!(((^(v3)) < v2) && true))) {
                    v35 := GenS3{F0: 255}
                    _ = v35
                }
                _ = v34
            }
        case 529:
            for fi36 := 0; fi36 < 3; fi36++ {
                v3 = v2
                _ = fi36
            }
        }
        fmt.Println(v3)
        _ = v30
    }
    v37 := GenS1{F0: "hello", F1: -2147483648}
    v38 := GenS3{F0: 950}
    v39 := []int{}
    if (true || false) {
        var v40 uint = 65535
        if ((-87 >= genFn1(v27, v40)) || true) {
            v27 -= (v27 ^ (^(v27)))
            var v41 int32 = -2147483648
            if len(v39) >= 2 {
                v39 = v39[1:len(v39)]
            }
            _ = v41
        }
        var v42 string = v1
        _ = v40
        _ = v42
    }
    if ((^(v3)) < v3) {
        switch 71 {
        case 9223372036854775807:
            fmt.Println(v27, v27)
        case -891:
            v43 := GenS3{F0: 65535}
            _ = v43
        default:
            v39 = append(v39, -2147483648)
        }
    }
    switch v1 {
    case "":
        v44 := GenS4{F0: v3, F1: 42, F2: 1}
        _ = v44
    case "a":
        fmt.Println(v26, v3, v3)
    case "hello":
        if len(v39) >= 2 {
            v39 = v39[1:len(v39)]
        }
    default:
        switch v1 {
        case "\n\t":
            if (((3.14159 == (0.0 / float32(-2.25))) || true) && ((v2 ^ 0) == v3)) {
                switch 127 {
                case 2147483647:
                    switch v26 {
                    case "hello":
                        var v45 string = v26
                        _ = v45
                    case "":
                        var v46 int64 = -669
                        _ = v46
                    }
                case 471:
                    if false {
                        v47 := GenS1{F0: "\n\t", F1: int32(v2)}
                        if len(v39) >= 2 {
                            v39 = v39[1:len(v39)]
                        }
                        if false {
                            for fi48 := 0; fi48 < 5; fi48++ {
                                _ = len(v39)
                                switch ("\n\t" + "\n\t") {
                                case "\n\t":
                                    if ((fi48 | 648) < fi48) {
                                        v26 = v1
                                        v2 = v2
                                    } else {
                                        v27 = v27
                                        fmt.Println(v3)
                                    }
                                default:
                                    fmt.Println(v1)
                                }
                                if (1000 != int16(v27)) {
                                    var v49 int32 = 42
                                    _ = v49
                                }
                                _ = fi48
                            }
                            v50 := GenS2{F0: 4, F1: genFn0(), F2: 1}
                            _ = v50
                        }
                        _ = v47
                    } else {
                        var v51 uint16 = 42
                        var v52 bool = true
                        if v52 {
                            var v53 float64 = 0.0
                            _ = v53
                        } else {
                            v27 += (^(v27))
                            v54 := GenS4{F0: v2, F1: int32(v27), F2: 1}
                            _ = v54
                        }
                        _ = v51
                        _ = v52
                    }
                case -203:
                    for fi55 := 0; fi55 < 4; fi55++ {
                        fmt.Println(fi55, v27)
                        var v56 int8 = 127
                        _ = fi55
                        _ = v56
                    }
                }
                var v57 uint64 = 2
                v58 := GenS4{F0: (v2 * 306), F1: 0, F2: 862}
                _ = v57
                _ = v58
            }
        default:
            var v59 uint64 = 4294967295
            _ = v59
        }
    }
    for fi60 := 0; fi60 < 1; fi60++ {
        var v61 float64 = (0.0 - (((0.0 / 1.0) - 0.0) - (float64((1.5 - 0.0)) - (3.14159 + 0.0))))
        for fi62 := 0; fi62 < 11; fi62++ {
            v63 := GenS2{F0: 2, F1: 1, F2: 556}
            switch (fi60 - 983) {
            case -472:
                v39 = append(v39, (fi60 * fi60))
            case 32768:
                switch genFn1(2147483647, 4294967295) {
                case 42:
                    switch "foo bar" {
                    case "hello":
                        fmt.Println(v2, v27, v1)
                    case "":
                        v1 = (v1 + ("foo bar" + "foo bar"))
                    case "\n\t":
                        var v64 uint = 233
                        _ = v64
                    default:
                        switch "foo bar" {
                        case "foo bar":
                            var v65 uint32 = v27
                            _ = v65
                        }
                    }
                }
            }
            v1 = "\n\t"
            _ = fi62
            _ = v63
        }
        v66 := GenS5{F0: 0}
        _ = fi60
        _ = v61
        _ = v66
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v39:", len(v39), v39)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v37
    _ = v38
    _ = v39
}
