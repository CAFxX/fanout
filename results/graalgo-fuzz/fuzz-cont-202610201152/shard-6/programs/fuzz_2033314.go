package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16; F2 int16 }
type GenS2 struct { F0 int8; F1 bool }
type GenS3 struct { F0 uint64; F1 uint8 }

func main() {
    v1 := make(map[string]uint32)
    var v2 uint64 = 826
    v3 := make(map[int]string)
    var v4 string = "a"
    v5 := make(map[string]string)
    var v6 int16 = 1
    fmt.Println(v4)
    if true {
        switch 128 {
        case 833:
            var v7 uint32 = 94
            _ = v7
        case 1:
            var v8 int64 = 773
            _ = v8
        case 42:
            var v9 uint16 = 65535
            _ = v9
        default:
            fmt.Println(v4)
        }
        v10 := GenS1{F0: 1, F1: 700, F2: (v6 | 1)}
        _ = v10
    }
    if (!((65535 != 255) || true)) {
        v11 := GenS1{F0: 1, F1: 255, F2: 0}
        var v12 string = "hello"
        if false {
            var v13 float32 = 0.1
            v14 := []int{}
            _ = v13
            _ = v14
        }
        _ = v11
        _ = v12
    } else {
        if (v6 <= -63) {
            if true {
                v15 := GenS2{F0: -128, F1: true}
                v16 := GenS1{F0: 65535, F1: 255, F2: v6}
                _ = v15
                _ = v16
            } else {
                if (false || false) {
                    switch "\n\t" {
                    case "hello":
                        var v17 float32 = ((float32(float64((0.1 - 1.5))) + 0.0) * 0.0)
                        _ = v17
                    case "a":
                        switch 32768 {
                        case 237:
                            if (v6 > v6) {
                                v18 := GenS2{F0: 42, F1: ((1e10 / (0.0 - float64(-2.25))) < -2.25)}
                                _ = v18
                            } else {
                                switch ((v4 + (("hello" + ((v4 + v4) + v4)) + ("hello" + (v4 + v4)))) + v4) {
                                case "hello":
                                    fmt.Println(v6, v4, v2)
                                case "a":
                                    for fi19 := 0; fi19 < 17; fi19++ {
                                        fmt.Println(v4, v4)
                                        if (1 == 2147483647) {
                                            var v20 int32 = 640
                                            if true {
                                                switch ("\n\t" + v4) {
                                                case "":
                                                    switch (fi19 / ((-2147483648) | 1)) {
                                                    case 14:
                                                        for fi21 := 0; fi21 < 10; fi21++ {
                                                            v22 := v1[((v4 + ("foo bar" + v4)) + "")]
                                                            var v23 string = v4
                                                            v24 := v3[fi19]
                                                            _ = fi21
                                                            _ = v22
                                                            _ = v23
                                                            _ = v24
                                                        }
                                                    case -482:
                                                        v25 := GenS3{F0: (v2 ^ v2), F1: 126}
                                                        _ = v25
                                                    }
                                                default:
                                                    var v26 uint8 = 1
                                                    _ = v26
                                                }
                                            }
                                            _ = v20
                                        } else {
                                            if ((fi19 << ((0) & 63)) >= fi19) {
                                                v27 := []int{fi19, fi19}
                                                v28 := GenS1{F0: 129, F1: 65535, F2: (v6 * v6)}
                                                _ = len(v27)
                                                _ = v27
                                                _ = v28
                                            } else {
                                                for fi29 := 0; fi29 < 18; fi29++ {
                                                    var v30 float64 = (float64(0.0) + (0.1 / ((((1e10 * 1.5) * 0.1) / 1.0) - (((1e10 + 3.14159) * 1.5) + (((((0.1 * (1e10 * 0.1)) + -2.25) - (float64(3.14159) / 3.14159)) / 1.5) + 0.1)))))
                                                    v31 := GenS3{F0: (v2 / ((v2) | 1)), F1: 76}
                                                    switch ("\n\t" + "\n\t") {
                                                    case "":
                                                        var v32 int16 = 1000
                                                        _ = v32
                                                    }
                                                    _ = fi29
                                                    _ = v30
                                                    _ = v31
                                                }
                                                switch 127 {
                                                case 42:
                                                    switch (fi19 - (fi19 & 128)) {
                                                    case 128:
                                                        var v33 uint = 949
                                                        _ = v33
                                                    case 1:
                                                        var v34 string = ""
                                                        _ = v34
                                                    default:
                                                        if true {
                                                            var v35 uint8 = 42
                                                            var v36 float32 = 0.1
                                                            switch (fi19 % (((fi19 - -2147483648)) | 1)) {
                                                            case -2147483648:
                                                                v4 = "\n\t"
                                                            case 1:
                                                                delete(v5, v4)
                                                            case 32768:
                                                                v2 = (v2 + (v2 % (((v2 / ((v2) | 1))) | 1)))
                                                            }
                                                            _ = v35
                                                            _ = v36
                                                        }
                                                    }
                                                default:
                                                    v4 = v4
                                                }
                                                var v37 uint = 342
                                                _ = v37
                                            }
                                        }
                                        _ = fi19
                                    }
                                default:
                                    if false {
                                        v2 -= v2
                                        var v38 int = 9223372036854775807
                                        v39 := []int{(-(v38)), 32768, -9223372036854775808, v38}
                                        _ = v38
                                        _ = v39
                                    } else {
                                        for fi40 := 0; fi40 < 15; fi40++ {
                                            v2 += 691
                                            v4 = (v4 + (v4 + (("a" + "\n\t") + v4)))
                                            v6 -= (^(v6))
                                            _ = fi40
                                        }
                                    }
                                }
                                delete(v1, v4)
                                v41 := []int{0, -562, 9223372036854775807}
                                _ = v41
                            }
                        default:
                            if (0 != uint16(v2)) {
                                v2 += (v2 - (v2 >> ((1) & 63)))
                            } else {
                                var v42 uint32 = 1
                                for fi43 := 0; fi43 < 9; fi43++ {
                                    v44 := []int{fi43, fi43, -202}
                                    _ = fi43
                                    _ = v44
                                }
                                _ = v42
                            }
                        }
                    default:
                        v45 := GenS2{F0: 0, F1: true}
                        _ = v45
                    }
                    fmt.Println(v4, v4)
                } else {
                    fmt.Println(v4, v4, v6)
                    v1[v4] = uint32(v2)
                }
                fmt.Println(v2, v6, v4)
                if (3.14159 != 1.5) {
                    var v46 int = 32768
                    for fi47 := 0; fi47 < 15; fi47++ {
                        v46 = int(v6)
                        var v48 int = (v46 + (v46 / ((fi47) | 1)))
                        _ = fi47
                        _ = v48
                    }
                    _ = v46
                }
            }
            v4 = ("a" + (v4 + ("a" + "hello")))
            v49 := []int{127}
            _ = v49
        } else {
            _ = len(v3)
        }
        for fi50 := 0; fi50 < 6; fi50++ {
            var v51 int64 = 0
            _ = fi50
            _ = v51
        }
        if (42 < 128) {
            v52 := []int{4294967295, -256}
            _ = v52
        } else {
            _ = len(v1)
        }
    }
    v53 := GenS2{F0: -59, F1: false}
    for fi54 := 0; fi54 < 4; fi54++ {
        v6 -= v6
        v2 = (v2 ^ 0)
        _ = fi54
    }
    switch "a" {
    case "":
        switch -9223372036854775808 {
        case 2147483647:
            v4 = ("foo bar" + v4)
        case 32767:
            v2 += 23
        }
    default:
        v6 = (v6 + -32768)
    }
    var v55 uint64 = 0
    var v56 uint64 = v55
    v4 = v4
    for fi57 := 0; fi57 < 13; fi57++ {
        v6 = v6
        var v58 bool = (!false)
        _ = fi57
        _ = v58
    }
    for fi59 := 0; fi59 < 6; fi59++ {
        delete(v1, ((v4 + ("\n\t" + v4)) + "hello"))
        _ = fi59
    }
    if (v6 >= 521) {
        var v60 float32 = 1e10
        _ = v60
    }
    delete(v5, "")
    for fi61 := 0; fi61 < 19; fi61++ {
        _ = len(v5)
        var v62 float32 = (1e10 / float32((((1.5 - -2.25) / (0.0 - -2.25)) / (0.1 * (3.14159 / 1.0)))))
        _ = fi61
        _ = v62
    }
    switch v4 {
    case "hello":
        v63 := GenS3{F0: v2, F1: 128}
        _ = v63
    case "":
        var v64 uint16 = 65535
        _ = v64
    default:
        v65 := GenS3{F0: v56, F1: 211}
        _ = v65
    }
    for fi66 := 0; fi66 < 8; fi66++ {
        for fi67 := 0; fi67 < 20; fi67++ {
            v4 = "\n\t"
            v68 := []int{}
            _ = fi67
            _ = v68
        }
        _ = fi66
    }
    if true {
        var v69 uint = 235
        v4 = "\n\t"
        v69 -= 2
        _ = v69
    } else {
        var v70 int16 = (-(v6))
        fmt.Println(v55, v2)
        _ = v70
    }
    delete(v1, v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v55:", v55)
    fmt.Println("v56:", v56)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v53
    _ = v55
    _ = v56
    _ = v6
}
