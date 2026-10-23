package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32; F2 string }
type GenS2 struct { F0 bool; F1 int8; F2 uint8 }
type GenS3 struct { F0 int; F1 uint64; F2 uint }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int8; F1 int32; F2 int }
type GenS6 struct { F0 uint; F1 uint; F2 uint32 }
type GenS7 struct { F0 bool; F1 uint32 }
func genFn0(p0 int) string {
    v6 := GenS1{F0: (true || false), F1: 709, F2: ""}
    p0 = p0
    for fi7 := 0; fi7 < 14; fi7++ {
        fmt.Println(p0, p0)
        _ = fi7
    }
    _ = p0
    _ = v6
    return ""
}
func genFn1(p0 uint64) int16 {
    var v8 int32 = 0
    v9 := GenS1{F0: (((!true) && false) || true), F1: 570, F2: "hello"}
    for fi10 := 0; fi10 < 14; fi10++ {
        if (!((p0 ^ (p0 + (p0 - 2))) >= (^(p0)))) {
            for fi11 := 0; fi11 < 11; fi11++ {
                switch 4294967295 {
                case 2147483647:
                    fmt.Println(v8, p0)
                case 127:
                    v8 -= (^(v8))
                case 42:
                    switch fi10 {
                    case -945:
                        var v12 uint8 = 255
                        _ = v12
                    case 576:
                        v13 := []int{9223372036854775807, (fi11 % (((fi11 / ((-1) | 1))) | 1)), (fi11 * fi10), fi11}
                        _ = v13
                    }
                }
                v14 := []int{int(p0), (fi11 - fi10), (-(fi11))}
                var v15 int32 = v8
                _ = fi11
                _ = v14
                _ = v15
            }
            for fi16 := 0; fi16 < 6; fi16++ {
                v17 := GenS1{F0: (2147483647 > -379), F1: 4294967295, F2: "a"}
                v18 := GenS2{F0: false, F1: -128, F2: 123}
                v19 := GenS2{F0: (uint16(p0) < 1), F1: -128, F2: 0}
                _ = fi16
                _ = v17
                _ = v18
                _ = v19
            }
        } else {
            p0 = 42
            for fi20 := 0; fi20 < 3; fi20++ {
                v21 := []int{fi20, (fi10 >> ((0) & 63)), 127}
                v8 -= 300
                var v22 uint64 = p0
                _ = fi20
                _ = v21
                _ = v22
            }
            if ((((0.0 / 1.0) + ((float32(float64(3.14159)) + ((1e10 + 1.5) / 1e10)) - 3.14159)) > (0.0 - 0.0)) && false) {
                var v23 bool = ((1.5 / 1.5) >= 0.1)
                v24 := GenS3{F0: fi10, F1: (p0 % (((^(p0))) | 1)), F2: 255}
                _ = v23
                _ = v24
            } else {
                switch (^(fi10)) {
                case 2147483647:
                    v25 := []int{fi10}
                    _ = v25
                case 42:
                    v26 := GenS4{F0: 255}
                    _ = v26
                default:
                    v27 := []int{fi10}
                    _ = v27
                }
                fmt.Println(p0, fi10)
                p0 += 1
            }
        }
        var v28 uint8 = 2
        v29 := GenS5{F0: 63, F1: (v8 - v8), F2: fi10}
        _ = fi10
        _ = v28
        _ = v29
    }
    switch ((("hello" + "foo bar") + "") + "hello") {
    case "foo bar":
        switch "a" {
        case "hello":
            for fi30 := 0; fi30 < 4; fi30++ {
                for fi31 := 0; fi31 < 20; fi31++ {
                    var v32 uint = 1
                    _ = fi31
                    _ = v32
                }
                _ = fi30
            }
        default:
            v33 := []int{0, 32768}
            _ = v33
        }
    }
    var v34 string = "hello"
    _ = p0
    _ = v34
    _ = v8
    _ = v9
    return 1
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[uint32]string)
    v3 := make(map[string]int)
    v4 := make(map[int]string)
    var v5 string = ("\n\t" + "hello")
    v5 = v5
    fmt.Println(v5)
    var v35 int16 = 883
    var v36 int32 = -1
    fmt.Println(v35, v35, v5)
    if (255 == 2) {
        if false {
            for fi37 := 0; fi37 < 14; fi37++ {
                if (!(true || true)) {
                    v35 = 571
                } else {
                    v5 = "\n\t"
                }
                v5 = (v5 + v5)
                _ = fi37
            }
            delete(v3, v5)
        } else {
            v5 = v5
        }
        if (3.14159 >= float64(1.5)) {
            for fi38 := 0; fi38 < 9; fi38++ {
                v39 := GenS6{F0: 255, F1: 4294967295, F2: 65535}
                v35 = (v35 / ((genFn1(126)) | 1))
                _ = fi38
                _ = v39
            }
            v5 = (("" + "") + v5)
        }
        for fi40 := 0; fi40 < 9; fi40++ {
            switch 1 {
            case -2147483648:
                fmt.Println(fi40)
            case 1:
                switch v5 {
                case "foo bar":
                    var v41 float32 = -2.25
                    _ = v41
                case "\n\t":
                    v42 := v4[fi40]
                    _ = v42
                default:
                    v43 := []int{}
                    _ = v43
                }
            case 126:
                if true {
                    switch ("hello" + v5) {
                    case "hello":
                        v44 := GenS7{F0: (681 > 820), F1: 2}
                        _ = v44
                    case "a":
                        v36 += (^(v36))
                    }
                }
            }
            v5 = ""
            _ = fi40
        }
    } else {
        v35 -= (v35 >> ((184) & 63))
        var v45 uint8 = 248
        for fi46 := 0; fi46 < 1; fi46++ {
            v47 := v1[2147483648]
            _ = fi46
            _ = v47
        }
        _ = v45
    }
    v4[-2147483648] = v5
    for fi48 := 0; fi48 < 13; fi48++ {
        var v49 int32 = 2147483647
        var v50 bool = (0 > 610)
        for fi51 := 0; fi51 < 6; fi51++ {
            v5 = v5
            switch (^(fi51)) {
            case 42:
                fmt.Println(fi48, fi48)
            case 0:
                var v52 uint16 = 65535
                _ = v52
            case 552:
                if false {
                    _ = len(v3)
                    v53 := []int{fi51}
                    _ = v53
                }
            default:
                v35 = (-(v35))
                if (v50 || v50) {
                    if (true && v50) {
                        v5 = v5
                        if v50 {
                            v35 = v35
                        } else {
                            var v54 int8 = 0
                            _ = v54
                        }
                        v55 := GenS7{F0: v50, F1: 410}
                        _ = v55
                    }
                    _ = len(v4)
                }
                var v56 int64 = 135
                var v57 uint8 = 0
                delete(v4, fi51)
                _ = v56
                _ = v57
            }
            _ = fi51
        }
        _ = fi48
        _ = v49
        _ = v50
    }
    v58 := GenS6{F0: 255, F1: 42, F2: 42}
    v5 = "\n\t"
    if true {
        if (0 != 730) {
            for fi59 := 0; fi59 < 12; fi59++ {
                var v60 uint64 = 18446744073709551615
                _ = fi59
                _ = v60
            }
            v61 := GenS4{F0: 32767}
            var v62 int = -9223372036854775808
            _ = v61
            _ = v62
        }
    } else {
        var v63 uint64 = 42
        var v64 int32 = -377
        _ = v63
        _ = v64
    }
    v5 = v5
    v35 -= v35
    for fi65 := 0; fi65 < 19; fi65++ {
        var v66 uint32 = 883
        if false {
            var v67 float32 = (3.14159 * (0.1 - 1e10))
            v68 := []int{fi65, (fi65 | (-(fi65))), (fi65 * fi65)}
            switch v5 {
            case "":
                var v69 uint = 0
                _ = v69
            case "\n\t":
                if false {
                    v36 = (v36 - 42)
                    v67 -= (v67 + v67)
                } else {
                    for fi70 := 0; fi70 < 20; fi70++ {
                        var v71 int8 = -1
                        if len(v68) >= 2 {
                            v68 = v68[1:len(v68)]
                        }
                        fmt.Println(v67)
                        _ = fi70
                        _ = v71
                    }
                }
            case "hello":
                var v72 int64 = 42
                _ = v72
            }
            _ = v67
            _ = v68
        }
        _ = fi65
        _ = v66
    }
    var v73 uint = 255
    var v74 int16 = (v35 << (((^(v73))) & 63))
    for fi75 := 0; fi75 < 5; fi75++ {
        v76 := []int{fi75, fi75, 4294967295}
        _ = fi75
        _ = v76
    }
    if (!((((((2 >= 0) || (true || (2 >= 406))) && (false && (v36 == 0))) || ((v73 << ((0) & 63)) <= (^(v73)))) && true) || (2147483648 > 42))) {
        fmt.Println(v74, v35, v74)
        switch 32768 {
        case 316:
            v35 += (v35 ^ v74)
        case 2147483647:
            v77 := []int{int(v36)}
            _ = v77
        case 0:
            var v78 uint8 = 67
            _ = v78
        }
    } else {
        fmt.Println(v35)
        fmt.Println(v35, v5)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v73:", v73)
    fmt.Println("v74:", v74)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v58
    _ = v73
    _ = v74
}
