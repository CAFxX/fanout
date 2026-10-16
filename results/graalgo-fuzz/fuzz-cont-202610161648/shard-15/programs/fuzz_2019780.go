package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint64; F1 int }
type GenS3 struct { F0 uint32; F1 int32 }
type GenS4 struct { F0 uint8; F1 bool; F2 uint }
type GenS5 struct { F0 uint; F1 uint16; F2 uint64 }
type GenS6 struct { F0 uint64; F1 uint32; F2 int8 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 uint64; F1 int16 }
func genFn0(p0 string) uint8 {
    switch "a" {
    case "":
        if (-701 <= -2147483648) {
            var v7 int16 = 419
            _ = v7
        } else {
            var v8 int64 = -116
            switch -9223372036854775808 {
            case -1:
                p0 = p0
            case 0:
                p0 = p0
            case 2147483647:
                v8 -= v8
            }
            _ = v8
        }
    default:
        var v9 int = 4294967295
        _ = v9
    }
    var v10 uint64 = 0
    v11 := GenS1{F0: v10}
    _ = p0
    _ = v10
    _ = v11
    return 29
}
func genFn1(p0 int16, p1 int64, p2 uint) uint16 {
    var v12 uint64 = 18446744073709551615
    v13 := GenS1{F0: (v12 << ((42) & 63))}
    switch "\n\t" {
    case "foo bar":
        var v14 uint64 = v12
        _ = v14
    case "hello":
        var v15 uint16 = 1
        _ = v15
    case "\n\t":
        switch ((("hello" + ("foo bar" + "foo bar")) + ("\n\t" + "hello")) + (("a" + ("" + "hello")) + "\n\t")) {
        case "":
            if false {
                if (-9223372036854775808 < 32767) {
                    v16 := []int{}
                    switch "" {
                    case "":
                        var v17 uint = (p2 & (p2 ^ (p2 % ((p2) | 1))))
                        _ = v17
                    case "a":
                        for fi18 := 0; fi18 < 2; fi18++ {
                            v19 := GenS2{F0: v12, F1: (-(fi18))}
                            p0 = 42
                            _ = fi18
                            _ = v19
                        }
                    default:
                        p0 = int16(p2)
                    }
                    for fi20 := 0; fi20 < 14; fi20++ {
                        _ = len(v16)
                        for fi21 := 0; fi21 < 16; fi21++ {
                            switch 0 {
                            case 1:
                                var v22 int16 = (p0 - p0)
                                _ = v22
                            case 2147483647:
                                v16 = append(v16, int(p0))
                                var v23 int32 = 267
                                _ = v23
                            case 42:
                                switch fi21 {
                                case 127:
                                    var v24 uint = p2
                                    _ = v24
                                case -561:
                                    var v25 uint64 = (^(v12))
                                    _ = v25
                                }
                            }
                            switch "a" {
                            case "\n\t":
                                v12 = (v12 >> (((p2 % ((p2) | 1))) & 63))
                            case "":
                                fmt.Println(p0)
                            }
                            _ = fi21
                        }
                        _ = len(v16)
                        _ = fi20
                    }
                    _ = v16
                } else {
                    var v26 int16 = p0
                    if false {
                        switch "a" {
                        case "a":
                            p2 = p2
                        case "foo bar":
                            if false {
                                if true {
                                    switch "a" {
                                    case "\n\t":
                                        p0 += (p0 | p0)
                                    case "a":
                                        var v27 float64 = float64(float32((-2.25 + (-2.25 / 3.14159))))
                                        _ = v27
                                    default:
                                        var v28 uint32 = 2147483647
                                        _ = v28
                                    }
                                }
                                p2 = (p2 % (((p2 ^ (p2 << ((p2) & 63)))) | 1))
                            }
                        case "hello":
                            var v29 uint64 = (v12 - uint64(p2))
                            _ = v29
                        default:
                            var v30 bool = false
                            _ = v30
                        }
                    }
                    _ = v26
                }
            } else {
                v31 := GenS2{F0: (v12 / ((4294967295) | 1)), F1: -1}
                p1 = (p1 % ((42) | 1))
                _ = v31
            }
        case "\n\t":
            if false {
                if true {
                    fmt.Println(p2)
                    p2 = (p2 ^ 255)
                    switch "hello" {
                    case "":
                        fmt.Println(p0)
                    case "hello":
                        var v32 float64 = 1.5
                        _ = v32
                    default:
                        v33 := GenS3{F0: 42, F1: 42}
                        _ = v33
                    }
                } else {
                    v34 := GenS4{F0: 197, F1: false, F2: 856}
                    _ = v34
                }
                var v35 bool = (true || (true || (!(!((100 <= 127) && ((p2 ^ (p2 / ((255) | 1))) == p2))))))
                v36 := GenS5{F0: p2, F1: 253, F2: (v12 & 219)}
                _ = v35
                _ = v36
            }
        case "hello":
            switch 4294967295 {
            case 635:
                v37 := []int{}
                _ = v37
            case -9223372036854775808:
                v38 := []int{2147483647}
                _ = v38
            default:
                p2 += (p2 * p2)
            }
        default:
            p2 = (^(p2))
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v13
    return 65535
}

func main() {
    var v1 int = -9223372036854775808
    var v2 uint16 = 1
    var v3 int32 = 42
    var v4 int64 = 1
    var v5 uint64 = 587
    var v6 int8 = 127
    v1 -= v1
    fmt.Println(v6, v6)
    var v39 int = v1
    v2 = v2
    v1 = v1
    v40 := GenS6{F0: 0, F1: 1, F2: (v6 | v6)}
    fmt.Println(v5, v2, v1)
    for fi41 := 0; fi41 < 1; fi41++ {
        switch fi41 {
        case 32768:
            switch (fi41 - (-(fi41))) {
            case 1:
                v1 = -266
            case -9223372036854775808:
                v42 := GenS6{F0: (v5 | v5), F1: 558, F2: (v6 | (v6 & v6))}
                _ = v42
            }
        case 9223372036854775807:
            var v43 uint8 = 27
            _ = v43
        default:
            fmt.Println(fi41, v6, v4)
        }
        v44 := GenS7{F0: "foo bar"}
        _ = fi41
        _ = v44
    }
    v4 = v4
    v45 := []int{}
    if false {
        fmt.Println(v2, v2, v3)
    } else {
        if (255 <= 310) {
            if (v5 >= v5) {
                fmt.Println(v6, v4)
                if true {
                    fmt.Println(v1)
                    var v46 uint = 42
                    v1 = (^(v39))
                    _ = v46
                } else {
                    var v47 string = ("foo bar" + "\n\t")
                    var v48 string = ((("" + v47) + (v47 + v47)) + v47)
                    if (328 <= 786) {
                        var v49 uint8 = 134
                        _ = v49
                    } else {
                        var v50 int64 = v4
                        var v51 int32 = (^(v3))
                        _ = v50
                        _ = v51
                    }
                    _ = v47
                    _ = v48
                }
                if (v39 <= 32767) {
                    if true {
                        fmt.Println(v39, v4)
                        if ((-2.25 * (((3.14159 * ((1e10 / 1e10) + -2.25)) / -2.25) + 1e10)) == 0.1) {
                            _ = len(v45)
                        }
                        _ = len(v45)
                    } else {
                        if len(v45) >= 2 {
                            v45 = v45[1:len(v45)]
                        }
                        if true {
                            fmt.Println(v39)
                        }
                        switch "a" {
                        case "foo bar":
                            var v52 float32 = (((0.0 * 0.1) - ((0.0 * -2.25) + ((float32(0.0) * (0.1 - (3.14159 + 1e10))) * 1.5))) + -2.25)
                            _ = v52
                        }
                    }
                    v3 = (v3 ^ v3)
                    for fi53 := 0; fi53 < 6; fi53++ {
                        v45 = append(v45, v39)
                        if true {
                            for fi54 := 0; fi54 < 5; fi54++ {
                                if ((uint8(v2) != 128) || (float64(1e10) == 3.14159)) {
                                    fmt.Println(v3, v5, fi53)
                                    fmt.Println(v2)
                                }
                                _ = fi54
                            }
                            v1 -= 127
                        } else {
                            v5 -= (v5 - 401)
                            for fi55 := 0; fi55 < 20; fi55++ {
                                v6 = (^(v6))
                                var v56 int32 = -3
                                _ = fi55
                                _ = v56
                            }
                            var v57 uint16 = v2
                            _ = v57
                        }
                        _ = fi53
                    }
                } else {
                    if (!("\n\t" <= "")) {
                        v58 := GenS2{F0: v5, F1: (v1 & 4294967295)}
                        var v59 uint8 = 2
                        _ = v58
                        _ = v59
                    }
                    _ = len(v45)
                    v5 = v5
                }
            } else {
                v39 = 1
                if false {
                    fmt.Println(v6, v6, v5)
                    var v60 uint64 = (^(v5))
                    switch "" {
                    case "hello":
                        v45 = append(v45, v1)
                    case "\n\t":
                        var v61 int8 = 42
                        _ = v61
                    default:
                        v3 -= (v3 + 2147483647)
                    }
                    _ = v60
                } else {
                    switch v39 {
                    case 152:
                        v62 := GenS8{F0: v5, F1: 1000}
                        _ = v62
                    }
                    v63 := GenS3{F0: uint32(v5), F1: v3}
                    _ = v63
                }
            }
        } else {
            for fi64 := 0; fi64 < 8; fi64++ {
                var v65 string = "foo bar"
                if true {
                    v66 := GenS7{F0: (v65 + v65)}
                    _ = v66
                } else {
                    for fi67 := 0; fi67 < 14; fi67++ {
                        var v68 int8 = v6
                        v65 = v65
                        _ = fi67
                        _ = v68
                    }
                    fmt.Println(v1, v6, v6)
                }
                _ = fi64
                _ = v65
            }
        }
    }
    v4 -= int64(v3)
    v45 = append(v45, (v39 & v39))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v45:", len(v45), v45)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v45
    _ = v5
    _ = v6
}
