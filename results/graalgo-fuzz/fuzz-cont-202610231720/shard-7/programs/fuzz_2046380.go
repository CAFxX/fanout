package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint16; F1 uint64 }
type GenS6 struct { F0 uint8; F1 string }
type GenS7 struct { F0 bool; F1 int32 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 int8; F1 uint32; F2 int64 }
func genFn0(p0 bool) int8 {
    if p0 {
        for fi6 := 0; fi6 < 4; fi6++ {
            if false {
                for fi7 := 0; fi7 < 14; fi7++ {
                    var v8 uint32 = 1
                    _ = fi7
                    _ = v8
                }
                var v9 string = "foo bar"
                _ = v9
            } else {
                var v10 int8 = 1
                for fi11 := 0; fi11 < 20; fi11++ {
                    if false {
                        v12 := GenS1{F0: 225, F1: 1}
                        v13 := GenS1{F0: 222, F1: 65535}
                        _ = v12
                        _ = v13
                    }
                    _ = fi11
                }
                if (int32(fi6) > 1) {
                    var v14 uint32 = 279
                    _ = v14
                }
                _ = v10
            }
            p0 = (("\n\t" != (("foo bar" + "foo bar") + "\n\t")) || (0.1 != ((float32(((-2.25 - -2.25) + 3.14159)) + 0.0) - 1e10)))
            var v15 uint32 = 282
            _ = fi6
            _ = v15
        }
        p0 = (3.14159 != (0.0 / (0.1 + 0.0)))
    }
    switch -181 {
    case -404:
        var v16 uint16 = 32767
        _ = v16
    case -9223372036854775808:
        if p0 {
            fmt.Println(p0)
        } else {
            switch 1 {
            case 558:
                var v17 uint = 42
                _ = v17
            }
        }
    case -2147483648:
        var v18 bool = p0
        _ = v18
    }
    p0 = (127 != -123)
    _ = p0
    return 1
}
func genFn1() bool {
    switch 2147483647 {
    case 32768:
        var v19 int = 127
        _ = v19
    default:
        var v20 bool = (!true)
        _ = v20
    }
    if false {
        switch "\n\t" {
        case "hello":
            var v21 uint8 = 123
            _ = v21
        default:
            var v22 uint16 = 32767
            _ = v22
        }
        for fi23 := 0; fi23 < 19; fi23++ {
            if (189 <= 0) {
                for fi24 := 0; fi24 < 8; fi24++ {
                    var v25 float64 = -2.25
                    var v26 uint64 = 2
                    var v27 string = "a"
                    _ = fi24
                    _ = v25
                    _ = v26
                    _ = v27
                }
            }
            var v28 bool = false
            for fi29 := 0; fi29 < 19; fi29++ {
                v30 := []int{fi23, (fi23 | fi29)}
                _ = fi29
                _ = v30
            }
            _ = fi23
            _ = v28
        }
        v31 := []int{-15, 128}
        _ = v31
    }
    for fi32 := 0; fi32 < 16; fi32++ {
        v33 := GenS2{F0: 200}
        var v34 float64 = ((-2.25 + (float64((((((3.14159 - ((1.5 / 1e10) / (3.14159 - (0.0 - -2.25)))) + (((((1e10 + 0.1) - 0.1) / 3.14159) + float32(0.0)) + 1.5)) * (3.14159 - (-2.25 / -2.25))) * 3.14159) / (0.1 * 0.1))) / 1.0)) - (float64(0.0) / -2.25))
        switch ("\n\t" + "a") {
        case "":
            fmt.Println(v34, fi32)
        case "\n\t":
            switch (^(fi32)) {
            case -678:
                switch "a" {
                case "a":
                    v35 := GenS2{F0: 4294967295}
                    _ = v35
                }
            default:
                for fi36 := 0; fi36 < 19; fi36++ {
                    v37 := GenS2{F0: 337}
                    var v38 int64 = 0
                    _ = fi36
                    _ = v37
                    _ = v38
                }
            }
        default:
            if true {
                var v39 float64 = (float64((3.14159 - (0.1 - -2.25))) / ((v34 - v34) + v34))
                v39 += (v34 / (v39 - (float64(((1.5 * ((0.0 * 0.1) / (0.0 + (1e10 / 1.5)))) - (0.1 + (-2.25 + ((float32((float64(-2.25) / v39)) / 1.0) / 1.5))))) - (-2.25 - (1e10 * (0.0 * v39))))))
                _ = v39
            }
        }
        _ = fi32
        _ = v33
        _ = v34
    }
    for fi40 := 0; fi40 < 14; fi40++ {
        var v41 int8 = genFn0(false)
        var v42 int16 = 42
        v43 := GenS1{F0: 22, F1: 645}
        _ = fi40
        _ = v41
        _ = v42
        _ = v43
    }
    return false
}

func main() {
    var v1 int8 = 127
    var v2 float64 = 1e10
    v3 := []int{}
    var v4 float32 = 0.0
    v5 := make(map[int]bool)
    for fi44 := 0; fi44 < 13; fi44++ {
        if len(v3) > 0 {
            v3[0] = (^(fi44))
        }
        var v45 int16 = 1
        v46 := GenS3{F0: v1}
        _ = fi44
        _ = v45
        _ = v46
    }
    v4 += ((-2.25 + 0.1) * float32(1.5))
    v47 := GenS4{F0: v1}
    var v48 int = 32767
    switch 127 {
    case 278:
        for fi49 := 0; fi49 < 16; fi49++ {
            var v50 uint16 = 579
            switch -396 {
            case 42:
                v51 := GenS5{F0: v50, F1: 0}
                _ = v51
            case 1:
                if (((v2 * (v2 / v2)) * v2) == (((float64(float32(1.5)) - float64((v4 / ((v4 - ((1e10 / 1.0) / v4)) / float32((v2 - 0.0)))))) - v2) * (v2 * (1e10 / (((v2 + (3.14159 / ((((v2 / 1.5) - 0.0) * float64(v4)) + (3.14159 * float64(v4))))) + (v2 / (v2 / (v2 + 3.14159)))) + 1.5))))) {
                    v50 = (v50 << ((255) & 63))
                    switch "foo bar" {
                    case "a":
                        v2 = float64(float32(v2))
                    }
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                } else {
                    v1 -= v1
                    v3 = append(v3, fi49)
                }
            default:
                switch (v48 >> ((1) & 63)) {
                case 9223372036854775807:
                    v48 += (fi49 << ((255) & 63))
                case 128:
                    v52 := GenS4{F0: v1}
                    _ = v52
                default:
                    if false {
                        fmt.Println(v1, v1)
                    }
                }
            }
            _ = fi49
            _ = v50
        }
    default:
        v53 := GenS6{F0: 0, F1: "hello"}
        _ = v53
    }
    v54 := GenS5{F0: 221, F1: 390}
    if (!false) {
        var v55 int = v48
        var v56 string = "foo bar"
        _ = v55
        _ = v56
    } else {
        v1 += v1
    }
    v1 -= genFn0((1.5 > (v2 / v2)))
    var v57 bool = false
    switch "a" {
    case "hello":
        if ((-1 != 0) || (("\n\t" + "\n\t") == ("foo bar" + ""))) {
            if (1 == 42) {
                switch ("hello" + "foo bar") {
                case "\n\t":
                    if v57 {
                        for fi58 := 0; fi58 < 8; fi58++ {
                            v1 -= v1
                            if (579 < 255) {
                                var v59 string = "hello"
                                var v60 int8 = (v1 % (((v1 + -91)) | 1))
                                v59 = (v59 + v59)
                                _ = v59
                                _ = v60
                            } else {
                                var v61 uint64 = 4294967295
                                _ = v61
                            }
                            _ = fi58
                        }
                        v62 := GenS7{F0: v57, F1: -1}
                        _ = len(v5)
                        _ = v62
                    } else {
                        v48 += v48
                    }
                case "hello":
                    if ((v4 <= v4) && true) {
                        _ = len(v3)
                    }
                case "":
                    fmt.Println(v57, v57)
                }
                v48 = v48
                v63 := GenS6{F0: 255, F1: ("a" + "hello")}
                _ = v63
            }
            fmt.Println(v2)
        } else {
            if v57 {
                fmt.Println(v57, v2)
                if len(v3) > 0 {
                    v3[0] = 2147483647
                }
            } else {
                if v57 {
                    var v64 uint = 713
                    v65 := GenS8{F0: "\n\t"}
                    v66 := GenS5{F0: 65535, F1: 2}
                    _ = v64
                    _ = v65
                    _ = v66
                }
            }
        }
    case "foo bar":
        v57 = false
    default:
        v67 := GenS9{F0: v1, F1: 65535, F2: 1}
        _ = v67
    }
    var v68 uint16 = 42
    for fi69 := 0; fi69 < 2; fi69++ {
        v1 -= genFn0((42 != -23))
        fmt.Println(v57, v57)
        _ = fi69
    }
    v2 += (3.14159 * -2.25)
    var v70 uint32 = 2
    if v57 {
        var v71 uint64 = 1
        _ = v71
    }
    v2 -= (v2 - 0.0)
    var v72 string = "foo bar"
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v57:", v57)
    fmt.Println("v68:", v68)
    fmt.Println("v70:", v70)
    fmt.Println("v72:", v72)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v48
    _ = v5
    _ = v54
    _ = v57
    _ = v68
    _ = v70
    _ = v72
}
