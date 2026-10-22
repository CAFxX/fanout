package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint }
type GenS2 struct { F0 int16; F1 int64 }
type GenS3 struct { F0 bool; F1 int16 }
type GenS4 struct { F0 int8; F1 uint8 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int64; F1 int16 }
func genFn0(p0 int64, p1 int64) uint64 {
    p1 = 300
    if (42 != 1000) {
        v7 := []int{2147483647, 128, 9223372036854775807}
        for fi8 := 0; fi8 < 4; fi8++ {
            v9 := GenS1{F0: 227, F1: 405}
            if true {
                var v10 uint64 = 0
                v7 = append(v7, 42)
                for fi11 := 0; fi11 < 19; fi11++ {
                    v12 := GenS2{F0: 728, F1: (^(p1))}
                    if (-929 > fi8) {
                        var v13 uint64 = v10
                        _ = v13
                    } else {
                        var v14 uint8 = 128
                        _ = v14
                    }
                    _ = fi11
                    _ = v12
                }
                _ = v10
            }
            _ = fi8
            _ = v9
        }
        _ = v7
    } else {
        if false {
            var v15 float32 = 3.14159
            _ = v15
        }
        fmt.Println(p1, p0)
    }
    v16 := []int{0, -273, -9223372036854775808}
    if len(v16) >= 2 {
        v16 = v16[1:len(v16)]
    }
    for fi17 := 0; fi17 < 8; fi17++ {
        p0 = (-(p1))
        switch "" {
        case "hello":
            var v18 float64 = 0.1
            _ = v18
        case "foo bar":
            for fi19 := 0; fi19 < 11; fi19++ {
                fmt.Println(fi19, fi17, fi17)
                _ = len(v16)
                _ = fi19
            }
        case "":
            switch fi17 {
            case -2147483648:
                p0 = (p1 / ((p0) | 1))
            case 42:
                v20 := GenS2{F0: -233, F1: 9223372036854775807}
                _ = v20
            case 309:
                fmt.Println(p1)
            }
        }
        v21 := GenS3{F0: (2147483647 < 590), F1: 1}
        _ = fi17
        _ = v21
    }
    _ = p0
    _ = p1
    _ = v16
    return 641
}
func genFn1(p0 bool, p1 int32, p2 uint) uint32 {
    var v22 int = 127
    if (2 <= 239) {
        var v23 float32 = (float32(3.14159) / 3.14159)
        for fi24 := 0; fi24 < 14; fi24++ {
            var v25 string = (("hello" + ((("hello" + "hello") + "foo bar") + ("foo bar" + "a"))) + "")
            fmt.Println(p2, p1, p2)
            var v26 string = v25
            _ = fi24
            _ = v25
            _ = v26
        }
        _ = v23
    } else {
        if p0 {
            var v27 float64 = (1.5 - 0.1)
            if p0 {
                v28 := []int{0, -120, (v22 & v22)}
                v29 := GenS3{F0: p0, F1: -1}
                for fi30 := 0; fi30 < 14; fi30++ {
                    p1 = (p1 >> ((p2) & 63))
                    for fi31 := 0; fi31 < 4; fi31++ {
                        v32 := GenS1{F0: 0, F1: (p2 >> (((^(p2))) & 63))}
                        switch ("" + "") {
                        case "\n\t":
                            var v33 uint32 = 2147483647
                            _ = v33
                        case "hello":
                            var v34 uint8 = 160
                            _ = v34
                        case "foo bar":
                            switch "foo bar" {
                            case "\n\t":
                                if len(v28) >= 2 {
                                    v28 = v28[1:len(v28)]
                                }
                            case "hello":
                                switch "" {
                                case "foo bar":
                                    p0 = p0
                                }
                            case "a":
                                var v35 string = ((("" + ("" + "a")) + "hello") + "hello")
                                _ = v35
                            }
                        default:
                            var v36 int8 = 1
                            _ = v36
                        }
                        v37 := GenS4{F0: -58, F1: 219}
                        _ = fi31
                        _ = v32
                        _ = v37
                    }
                    if len(v28) >= 2 {
                        v28 = v28[1:len(v28)]
                    }
                    _ = fi30
                }
                _ = v28
                _ = v29
            }
            _ = v27
        } else {
            for fi38 := 0; fi38 < 15; fi38++ {
                fmt.Println(p1, fi38, fi38)
                v39 := GenS5{F0: p0}
                var v40 uint8 = 1
                _ = fi38
                _ = v39
                _ = v40
            }
            var v41 uint8 = 42
            _ = v41
        }
        if p0 {
            switch (v22 | v22) {
            case -9223372036854775808:
                v42 := []int{(-(v22)), v22, (v22 | (v22 % ((v22) | 1)))}
                _ = v42
            default:
                v43 := []int{v22, (v22 | v22), (v22 + (^(v22))), (v22 - v22)}
                _ = v43
            }
            var v44 int16 = 131
            _ = v44
        }
    }
    for fi45 := 0; fi45 < 16; fi45++ {
        for fi46 := 0; fi46 < 10; fi46++ {
            var v47 int64 = 0
            _ = fi46
            _ = v47
        }
        _ = fi45
    }
    v48 := GenS6{F0: 42, F1: 760}
    _ = p0
    _ = p1
    _ = p2
    _ = v22
    _ = v48
    return 2147483648
}

func main() {
    var v1 bool = (2 > 2)
    v2 := make(map[string]int)
    var v3 float32 = 0.1
    var v4 uint32 = 166
    var v5 uint16 = 458
    var v6 uint8 = 128
    fmt.Println(v4)
    if (!true) {
        for fi49 := 0; fi49 < 20; fi49++ {
            v50 := []int{127, (^(fi49))}
            var v51 int64 = 9223372036854775807
            fmt.Println(v3)
            _ = fi49
            _ = v50
            _ = v51
        }
        for fi52 := 0; fi52 < 4; fi52++ {
            v4 = 2147483648
            v4 -= v4
            _ = fi52
        }
        v53 := []int{32768, -2147483648, -9223372036854775808, 4294967295}
        _ = v53
    }
    v5 -= v5
    v54 := v2[(("\n\t" + "a") + ("" + ("hello" + "foo bar")))]
    v55 := []int{}
    switch ("" + "") {
    case "hello":
        v3 -= (v3 + 1.5)
    case "foo bar":
        if v1 {
            v5 = (^(v5))
            var v56 int16 = 1000
            _ = v56
        }
    case "\n\t":
        if true {
            var v57 uint8 = (^(v6))
            _ = v57
        }
    }
    v54 -= (v54 | (v54 & v54))
    var v58 uint16 = v5
    if ((0 <= -883) && (!(false || (65535 != 4294967295)))) {
        v3 = ((0.0 / (v3 * float32(((((1.5 / (((float64(float32((0.0 + (0.0 + -2.25)))) - 0.1) + 3.14159) / 1e10)) + ((0.1 - (0.1 / (1.5 - (0.0 + (-2.25 + (0.1 / 1.5)))))) - 1e10)) / 1.0) - (3.14159 + 0.1))))) + 1e10)
        var v59 uint64 = 42
        _ = v59
    } else {
        if false {
            if false {
                v58 += 621
            } else {
                v3 = ((1e10 - 1e10) * 3.14159)
                var v60 uint8 = 255
                _ = len(v2)
                _ = v60
            }
        }
        v55 = append(v55, v54)
        var v61 uint8 = v6
        _ = v61
    }
    switch "\n\t" {
    case "hello":
        v5 = (v58 % ((0) | 1))
    }
    var v62 int32 = 0
    v63 := GenS3{F0: (!(!((v5 & (v58 * 42)) >= v58))), F1: 602}
    var v64 int64 = 1
    for fi65 := 0; fi65 < 11; fi65++ {
        var v66 string = ("\n\t" + "\n\t")
        var v67 int16 = 0
        _ = fi65
        _ = v66
        _ = v67
    }
    var v68 uint64 = 18446744073709551615
    v3 += v3
    v54 += (-(v54))
    var v69 string = ("" + "\n\t")
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v54:", v54)
    fmt.Println("slicelen v55:", len(v55), v55)
    fmt.Println("v58:", v58)
    fmt.Println("v6:", v6)
    fmt.Println("v62:", v62)
    fmt.Println("v64:", v64)
    fmt.Println("v68:", v68)
    fmt.Println("v69:", v69)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v54
    _ = v55
    _ = v58
    _ = v6
    _ = v62
    _ = v63
    _ = v64
    _ = v68
    _ = v69
}
