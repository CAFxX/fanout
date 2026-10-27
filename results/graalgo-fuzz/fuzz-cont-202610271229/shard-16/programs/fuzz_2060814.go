package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint; F1 int32 }
type GenS4 struct { F0 uint16; F1 uint16; F2 uint8 }
type GenS5 struct { F0 int; F1 bool; F2 bool }
type GenS6 struct { F0 uint16; F1 int8 }
type GenS7 struct { F0 uint8 }
func genFn0(p0 string) int {
    for fi7 := 0; fi7 < 14; fi7++ {
        switch p0 {
        case "":
            var v8 uint8 = 127
            _ = v8
        default:
            p0 = p0
        }
        if (2147483647 <= 468) {
            p0 = p0
        }
        _ = fi7
    }
    v9 := []int{}
    p0 = "hello"
    if (0.1 <= 1.5) {
        var v10 uint64 = 0
        for fi11 := 0; fi11 < 7; fi11++ {
            for fi12 := 0; fi12 < 11; fi12++ {
                var v13 int16 = -433
                var v14 int = fi12
                if false {
                    if true {
                        v14 += (fi12 & (v14 | fi11))
                    }
                    if (-642 != 42) {
                        var v15 uint64 = v10
                        v15 = v15
                        _ = v15
                    }
                } else {
                    v16 := GenS1{F0: "a", F1: 42}
                    _ = v16
                }
                if true {
                    var v17 int32 = 2147483647
                    var v18 uint32 = 2147483647
                    var v19 uint32 = (v18 / ((v18) | 1))
                    if true {
                        v17 = -905
                        if len(v9) > 0 {
                            v9[0] = v14
                        }
                    }
                    v20 := GenS2{F0: 128}
                    _ = v17
                    _ = v18
                    _ = v19
                    _ = v20
                } else {
                    if len(v9) >= 2 {
                        v9 = v9[1:len(v9)]
                    }
                }
                var v21 int = fi11
                _ = fi12
                _ = v13
                _ = v14
                _ = v21
            }
            for fi22 := 0; fi22 < 8; fi22++ {
                p0 = ("foo bar" + p0)
                var v23 uint8 = 127
                if false {
                    var v24 uint32 = 803
                    var v25 uint8 = 2
                    _ = v24
                    _ = v25
                } else {
                    if true {
                        v9 = append(v9, fi22)
                        fmt.Println(v23)
                    } else {
                        if (!true) {
                            if true {
                                v10 -= v10
                            }
                            v23 -= 2
                            fmt.Println(fi22)
                            fmt.Println(v10, p0)
                        } else {
                            var v26 bool = ((0 <= -40) && (!(-430 != -960)))
                            var v27 string = "foo bar"
                            _ = v26
                            _ = v27
                        }
                        v10 += 0
                        fmt.Println(v10, v10)
                    }
                    var v28 uint32 = 999
                    v28 = v28
                    _ = v28
                }
                _ = fi22
                _ = v23
            }
            p0 = ""
            _ = fi11
        }
        _ = v10
    }
    _ = p0
    _ = v9
    return 9223372036854775807
}
func genFn1(p0 uint, p1 string, p2 int16) int32 {
    var v29 uint32 = 557
    for fi30 := 0; fi30 < 1; fi30++ {
        for fi31 := 0; fi31 < 1; fi31++ {
            p0 -= (^(p0))
            _ = fi31
        }
        if true {
            if (0 >= 128) {
                var v32 int = int(p0)
                switch (v32 + fi30) {
                case -2147483648:
                    v33 := GenS3{F0: (p0 | (p0 & 0)), F1: 0}
                    _ = v33
                case 581:
                    p2 += (-(p2))
                }
                _ = v32
            } else {
                for fi34 := 0; fi34 < 8; fi34++ {
                    p2 = (p2 << (((p0 & p0)) & 63))
                    v35 := []int{-2147483648, (fi34 / ((fi34) | 1))}
                    if false {
                        p0 = p0
                        p1 = (p1 + (("a" + p1) + (p1 + p1)))
                        v29 -= (v29 | v29)
                    } else {
                        v36 := GenS4{F0: 255, F1: 997, F2: 0}
                        v37 := GenS5{F0: 128, F1: false, F2: false}
                        _ = v36
                        _ = v37
                    }
                    _ = fi34
                    _ = v35
                }
            }
            v38 := GenS3{F0: (^(p0)), F1: 163}
            _ = v38
        }
        switch ("" + ((p1 + p1) + p1)) {
        case "foo bar":
            v39 := []int{fi30}
            _ = v39
        case "a":
            v40 := []int{(^(fi30)), genFn0("\n\t"), fi30, fi30}
            _ = v40
        default:
            p2 -= (^(p2))
        }
        _ = fi30
    }
    var v41 string = ""
    for fi42 := 0; fi42 < 13; fi42++ {
        v43 := GenS2{F0: 154}
        _ = fi42
        _ = v43
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v29
    _ = v41
    return 42
}

func main() {
    var v1 float64 = 0.0
    v2 := make(map[uint32]int)
    v3 := make(map[string]bool)
    var v4 float64 = (v1 - v1)
    v5 := []int{-9223372036854775808, -9223372036854775808}
    var v6 float32 = (0.1 / float32(1e10))
    if false {
        fmt.Println(v1)
        if true {
            for fi44 := 0; fi44 < 14; fi44++ {
                if ((1 >= 1) && (false || (true || true))) {
                    var v45 float32 = -2.25
                    fmt.Println(v1, v4, v6)
                    v46 := GenS1{F0: "\n\t", F1: 42}
                    _ = v45
                    _ = v46
                } else {
                    v1 = v4
                }
                switch (fi44 >> ((98) & 63)) {
                case -1:
                    fmt.Println(v1)
                }
                if (!false) {
                    v47 := GenS6{F0: 65535, F1: -128}
                    for fi48 := 0; fi48 < 4; fi48++ {
                        if (100 == -128) {
                            _ = len(v5)
                        }
                        _ = fi48
                    }
                    _ = v47
                }
                _ = fi44
            }
        }
        var v49 bool = (v6 != v6)
        _ = v49
    } else {
        v50 := GenS2{F0: 42}
        if (0.0 < v4) {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            var v51 uint = 841
            _ = v51
        }
        _ = v50
    }
    var v52 int = genFn0(("\n\t" + ("" + "a")))
    v53 := GenS5{F0: 32767, F1: true, F2: false}
    if (!true) {
        var v54 int = (v52 / (((v52 + (v52 * -754))) | 1))
        if true {
            switch "" {
            case "a":
                if true {
                    fmt.Println(v52)
                }
            case "foo bar":
                var v55 int8 = -106
                _ = v55
            case "\n\t":
                v1 += ((v1 / (v1 / v1)) / (v4 * v1))
            }
        }
        fmt.Println(v52, v52, v4)
        _ = v54
    }
    v56 := GenS4{F0: 0, F1: 65535, F2: 127}
    if true {
        v1 = (3.14159 + v4)
        v1 += ((v4 - (v4 - v4)) * v1)
    } else {
        if len(v5) > 0 {
            v5[0] = v52
        }
    }
    v57 := GenS7{F0: 45}
    var v58 string = ""
    v1 = float64(3.14159)
    var v59 uint8 = 1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v52:", v52)
    fmt.Println("v58:", v58)
    fmt.Println("v59:", v59)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v53
    _ = v56
    _ = v57
    _ = v58
    _ = v59
    _ = v6
}
