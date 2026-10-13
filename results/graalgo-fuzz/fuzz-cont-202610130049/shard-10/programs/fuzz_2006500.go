package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint }
type GenS2 struct { F0 uint; F1 int16; F2 uint }
type GenS3 struct { F0 int32; F1 uint64; F2 uint16 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint64; F1 int16 }
type GenS6 struct { F0 bool; F1 string; F2 int }
type GenS7 struct { F0 int }
type GenS8 struct { F0 uint8; F1 int8 }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 bool; F1 int }
func genFn0(p0 int8) uint32 {
    switch -2147483648 {
    case 1:
        v6 := []int{}
        _ = v6
    case -292:
        p0 += p0
    case -548:
        for fi7 := 0; fi7 < 16; fi7++ {
            for fi8 := 0; fi8 < 4; fi8++ {
                if false {
                    fmt.Println(fi7)
                    v9 := []int{(fi7 >> ((0) & 63))}
                    v10 := GenS1{F0: 1, F1: 1}
                    _ = v10
                    _ = v9
                }
                var v11 string = "foo bar"
                fmt.Println(v11)
                _ = fi8
                _ = v11
            }
            var v12 uint8 = 0
            _ = fi7
            _ = v12
        }
    default:
        var v13 bool = true
        _ = v13
    }
    var v14 int32 = 0
    if true {
        switch -44 {
        case 106:
            var v15 uint = 323
            _ = v15
        case 127:
            p0 += int8(v14)
        case 42:
            if (1000 >= 1000) {
                if (0 > 32767) {
                    var v16 float32 = 1.5
                    _ = v16
                } else {
                    p0 += p0
                }
            }
        default:
            if true {
                switch "hello" {
                case "foo bar":
                    v14 = 0
                default:
                    fmt.Println(p0, p0)
                }
                var v17 uint32 = 2147483647
                fmt.Println(p0, p0)
                _ = v17
            } else {
                fmt.Println(v14, v14)
                p0 += -125
            }
        }
        v14 -= (-(v14))
        if (33 > 98) {
            v18 := GenS2{F0: 1, F1: 32767, F2: 18446744073709551615}
            _ = v18
        }
    }
    for fi19 := 0; fi19 < 3; fi19++ {
        if false {
            v14 += -554
            v14 = 42
            var v20 uint32 = 42
            _ = v20
        }
        _ = fi19
    }
    switch -2147483648 {
    case 1:
        var v21 int32 = 923
        _ = v21
    default:
        switch 158 {
        case 4294967295:
            p0 = (p0 | p0)
        case -2147483648:
            p0 += (-(p0))
        }
    }
    _ = p0
    _ = v14
    return 65535
}

func main() {
    var v1 uint8 = 127
    var v2 int32 = -851
    v3 := make(map[string]bool)
    v4 := make(map[uint32]string)
    var v5 bool = true
    for fi22 := 0; fi22 < 16; fi22++ {
        v5 = (!false)
        fmt.Println(v5, v1)
        _ = fi22
    }
    switch 2147483647 {
    case 32768:
        v23 := GenS3{F0: (v2 % ((2147483647) | 1)), F1: 42, F2: 32767}
        _ = v23
    case 128:
        v1 += v1
    default:
        v24 := []int{86, 0}
        _ = v24
    }
    v25 := GenS4{F0: ("\n\t" + "hello")}
    switch 9223372036854775807 {
    case -554:
        v26 := v3["hello"]
        _ = v26
    case -9223372036854775808:
        switch 9223372036854775807 {
        case -1:
            fmt.Println(v1)
        case 580:
            v27 := []int{292, 428, -9223372036854775808, -9223372036854775808}
            _ = v27
        }
    default:
        if v5 {
            fmt.Println(v1, v5, v2)
        } else {
            v28 := v3["hello"]
            for fi29 := 0; fi29 < 15; fi29++ {
                for fi30 := 0; fi30 < 5; fi30++ {
                    v31 := GenS2{F0: 0, F1: 1, F2: 0}
                    _ = fi30
                    _ = v31
                }
                v32 := GenS5{F0: 0, F1: -32768}
                _ = fi29
                _ = v32
            }
            var v33 float64 = 1.5
            _ = v28
            _ = v33
        }
    }
    var v34 bool = false
    fmt.Println(v1)
    switch ("\n\t" + ("a" + "a")) {
    case "foo bar":
        var v35 int32 = v2
        _ = v35
    case "a":
        v5 = (v2 <= 277)
    default:
        var v36 int8 = 42
        _ = v36
    }
    v2 = v2
    var v37 uint = 0
    var v38 uint8 = v1
    var v39 uint32 = 4294967295
    if false {
        if ((42 > -1) || v5) {
            v40 := v3[""]
            fmt.Println(v40, v38)
            v41 := GenS6{F0: (-2.25 == ((float64((3.14159 - 3.14159)) - (1e10 - (-2.25 - ((0.1 - float64(1e10)) * -2.25)))) + (0.1 / 3.14159))), F1: ((("a" + "a") + (("\n\t" + "hello") + ("foo bar" + "foo bar"))) + ("foo bar" + "\n\t")), F2: -723}
            _ = v40
            _ = v41
        }
    }
    switch 32767 {
    case 32768:
        v39 = (v39 >> ((v37) & 63))
    case 42:
        for fi42 := 0; fi42 < 19; fi42++ {
            v43 := GenS1{F0: (v1 >> ((129) & 63)), F1: 42}
            v44 := GenS2{F0: v37, F1: int16(fi42), F2: v37}
            _ = fi42
            _ = v43
            _ = v44
        }
    }
    if (true && (!true)) {
        v39 = 0
        for fi45 := 0; fi45 < 6; fi45++ {
            for fi46 := 0; fi46 < 9; fi46++ {
                v5 = false
                switch fi45 {
                case -1:
                    delete(v3, "a")
                default:
                    if (v5 || (1e10 < 0.0)) {
                        v39 = (v39 * v39)
                        v37 = (v37 - 4294967295)
                        switch "\n\t" {
                        case "hello":
                            if v5 {
                                var v47 int16 = 1000
                                v1 -= (v1 - v1)
                                var v48 uint = 732
                                _ = v47
                                _ = v48
                            } else {
                                fmt.Println(v2)
                                var v49 float64 = 0.0
                                if (0 != (fi45 >> ((v37) & 63))) {
                                    fmt.Println(fi46, fi46, v49)
                                    v50 := []int{fi45, fi46}
                                    v2 = v2
                                    _ = v50
                                }
                                _ = v49
                            }
                        default:
                            v39 -= (v39 % ((v39) | 1))
                        }
                    } else {
                        v1 += (v1 % ((157) | 1))
                    }
                }
                if v34 {
                    if (0 <= 65535) {
                        var v51 float32 = ((3.14159 - 0.1) / (3.14159 / 1e10))
                        v52 := GenS1{F0: (v1 ^ v38), F1: v37}
                        _ = v51
                        _ = v52
                    } else {
                        v1 -= (v38 + (v38 % ((128) | 1)))
                        switch ("\n\t" + "") {
                        case "a":
                            fmt.Println(v5)
                            v53 := GenS7{F0: (fi45 - fi45)}
                            _ = v53
                        case "foo bar":
                            if ((("" != ("foo bar" + "hello")) || (true || (v38 < 158))) || v34) {
                                if v5 {
                                    var v54 string = ""
                                    switch fi46 {
                                    case 32768:
                                        var v55 float64 = (((0.0 * (float64((1.5 / 1e10)) * float64(-2.25))) * 0.1) + (float64((1e10 + 1e10)) - (float64((((((((((1.5 + 1e10) + 3.14159) * 0.1) + -2.25) - float32(0.1)) - (-2.25 * 1e10)) / (3.14159 * 1.5)) / float32(1e10)) / 3.14159)) * (0.0 / 1.0))))
                                        _ = v55
                                    case 42:
                                        var v56 uint64 = uint64(v1)
                                        _ = v56
                                    default:
                                        v57 := []int{fi46}
                                        _ = v57
                                    }
                                    _ = v54
                                }
                            } else {
                                v37 -= (v37 + v37)
                            }
                        default:
                            v58 := GenS7{F0: 703}
                            _ = v58
                        }
                    }
                } else {
                    switch fi45 {
                    case -438:
                        if false {
                            v59 := GenS8{F0: (v1 - (v38 * v1)), F1: int8(fi45)}
                            v60 := []int{(fi45 >> (((v37 >> ((96) & 63))) & 63)), int(v39), fi46}
                            v2 += (v2 + (v2 + (^(v2))))
                            _ = v59
                            _ = v60
                        }
                        v5 = (!false)
                        v37 -= (v37 % (((v37 * uint(v38))) | 1))
                        v38 = 128
                    case 128:
                        switch (fi45 << (((v37 >> ((v37) & 63))) & 63)) {
                        case 0:
                            switch "a" {
                            case "":
                                v38 = (v38 >> ((785) & 63))
                            case "a":
                                var v61 uint = v37
                                _ = v61
                            case "hello":
                                v1 = v1
                            }
                        default:
                            switch ("a" + "\n\t") {
                            case "a":
                                var v62 float64 = 0.1
                                _ = v62
                            case "hello":
                                v63 := GenS5{F0: 645, F1: 1}
                                v64 := []int{-731, fi45, fi45}
                                _ = v63
                                _ = v64
                            default:
                                var v65 bool = ((1 < 42) && v34)
                                _ = v65
                            }
                        }
                    default:
                        var v66 bool = (!(1e10 == (0.0 / (3.14159 - 0.0))))
                        _ = v66
                    }
                    v1 = v38
                }
                var v67 uint64 = 4294967295
                _ = fi46
                _ = v67
            }
            var v68 uint = 228
            for fi69 := 0; fi69 < 1; fi69++ {
                v38 += 248
                v1 = v38
                if v5 {
                    var v70 int16 = -32768
                    v34 = false
                    _ = v70
                }
                if v5 {
                    fmt.Println(fi69, v1, fi45)
                }
                _ = fi69
            }
            _ = fi45
            _ = v68
        }
        if (float64(0.0) < -2.25) {
            v71 := GenS9{F0: -128}
            delete(v3, "\n\t")
            if v5 {
                if v34 {
                    v39 = genFn0(127)
                } else {
                    var v72 int64 = -18
                    _ = v72
                }
                v38 = 127
            }
            _ = v71
        }
    }
    v34 = v34
    if v34 {
        fmt.Println(v5)
        for fi73 := 0; fi73 < 6; fi73++ {
            v4[v39] = "\n\t"
            v74 := GenS10{F0: true, F1: fi73}
            _ = fi73
            _ = v74
        }
    }
    if true {
        fmt.Println(v34, v1)
    } else {
        v38 -= uint8(v37)
        switch -2147483648 {
        case 821:
            v75 := []int{-295, 9223372036854775807, 42, -679}
            _ = v75
        case -1:
            switch 9223372036854775807 {
            case 9223372036854775807:
                var v76 int8 = 42
                _ = v76
            case -1:
                v77 := v4[201]
                _ = v77
            default:
                v78 := []int{}
                _ = v78
            }
        }
    }
    var v79 float32 = (((0.1 / 1e10) * (-2.25 - 3.14159)) + 0.1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v79:", v79)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v34
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v79
}
