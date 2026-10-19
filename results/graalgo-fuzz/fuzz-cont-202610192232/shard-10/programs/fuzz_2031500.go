package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32; F2 uint8 }
type GenS2 struct { F0 int64; F1 int16 }
type GenS3 struct { F0 uint; F1 int8; F2 uint32 }
type GenS4 struct { F0 int32; F1 uint16; F2 int }
type GenS5 struct { F0 int64; F1 int16 }
type GenS6 struct { F0 int16; F1 string }
func genFn0(p0 int, p1 string) uint {
    p0 = p0
    for fi5 := 0; fi5 < 18; fi5++ {
        v6 := GenS1{F0: 202, F1: 65535, F2: 42}
        for fi7 := 0; fi7 < 7; fi7++ {
            if true {
                fmt.Println(fi7, p0, fi7)
            } else {
                p1 = ("hello" + "\n\t")
                p1 = p1
                if true {
                    var v8 int64 = 42
                    fmt.Println(p0, fi7)
                    _ = v8
                }
                fmt.Println(p0, p1)
                v9 := []int{(fi7 * (fi5 | 32767))}
                var v10 float32 = 1.5
                fmt.Println(fi7)
                var v11 int = (fi5 | (fi5 / ((fi5) | 1)))
                switch "\n\t" {
                case "foo bar":
                    if (1 > 32767) {
                        var v12 int = v11
                        var v13 int32 = -1
                        v14 := GenS2{F0: -724, F1: -1}
                        if len(v9) > 0 {
                            v9[0] = -215
                        }
                        var v15 bool = (false && false)
                        _ = v12
                        _ = v13
                        _ = v14
                        _ = v15
                    } else {
                        if (!(true && (255 >= 2))) {
                            var v16 int8 = -128
                            _ = v16
                        }
                        switch 2147483647 {
                        case 0:
                            if true {
                                if len(v9) > 0 {
                                    v9[0] = v11
                                }
                                var v17 int32 = 1
                                _ = v17
                            } else {
                                v10 -= float32(0.0)
                            }
                        case 32768:
                            var v18 float32 = (((v10 - ((3.14159 + 1.5) * 3.14159)) / float32(1.5)) - v10)
                            v11 += fi7
                            _ = v18
                        case 9223372036854775807:
                            v11 -= -9223372036854775808
                        }
                        var v19 float64 = (3.14159 + float64(1e10))
                        _ = v19
                    }
                    if len(v9) > 0 {
                        v9[0] = (p0 & fi5)
                    }
                    _ = len(v9)
                case "\n\t":
                    if len(v9) >= 2 {
                        v9 = v9[1:len(v9)]
                    }
                }
                v10 = 0.0
                _ = v10
                _ = v11
                _ = v9
            }
            v20 := []int{fi5}
            var v21 int64 = 9223372036854775807
            _ = fi7
            _ = v20
            _ = v21
        }
        switch p1 {
        case "foo bar":
            for fi22 := 0; fi22 < 3; fi22++ {
                if (1 < 2147483647) {
                    v23 := []int{0}
                    var v24 uint = 18446744073709551615
                    _ = v23
                    _ = v24
                } else {
                    v25 := []int{-864, (p0 >> ((1) & 63)), p0, 32768}
                    v26 := GenS2{F0: 9223372036854775807, F1: -1}
                    _ = v25
                    _ = v26
                }
                if true {
                    if true {
                        p0 += (p0 + fi22)
                    }
                    v27 := []int{(fi22 ^ (p0 % ((fi5) | 1))), (^(fi22)), 32767, (fi22 + p0)}
                    p0 = (p0 + p0)
                    _ = v27
                }
                if false {
                    v28 := GenS1{F0: 257, F1: 0, F2: 2}
                    _ = v28
                } else {
                    p1 = p1
                }
                var v29 string = ("hello" + "a")
                v30 := []int{}
                _ = fi22
                _ = v29
                _ = v30
            }
        default:
            fmt.Println(p1, fi5)
        }
        _ = fi5
        _ = v6
    }
    v31 := []int{p0, p0}
    p0 = -1
    var v32 float64 = (float64((float32((0.0 * (float64(-2.25) - 1e10))) - float32((0.1 * 0.0)))) + (float64((1.5 - ((((0.1 - (1.5 * 1.5)) - (1.5 + -2.25)) - 0.1) + float32(float64(0.1))))) + 1.5))
    _ = p0
    _ = p1
    _ = v31
    _ = v32
    return 585
}
func genFn1(p0 int, p1 string, p2 int32) uint16 {
    if (((0.1 + 0.0) / (3.14159 + 1.5)) >= ((1e10 / 0.1) / (float32(0.0) + 1e10))) {
        var v33 float32 = -2.25
        if true {
            v34 := GenS2{F0: -1, F1: 676}
            for fi35 := 0; fi35 < 16; fi35++ {
                p0 = fi35
                _ = fi35
            }
            _ = v34
        } else {
            p2 -= p2
            p0 += (p0 >> ((38) & 63))
            v36 := GenS3{F0: 1, F1: 127, F2: 59}
            _ = v36
        }
        _ = v33
    }
    fmt.Println(p1)
    p1 = "hello"
    var v37 uint = 4294967295
    _ = p0
    _ = p1
    _ = p2
    _ = v37
    return 0
}

func main() {
    v1 := make(map[uint32]string)
    var v2 float32 = -2.25
    var v3 uint64 = 222
    var v4 int8 = -1
    var v38 uint32 = 4294967295
    if false {
        v38 -= (v38 << ((18446744073709551615) & 63))
        v39 := GenS4{F0: 25, F1: genFn1(-1, "hello", 515), F2: 2147483647}
        for fi40 := 0; fi40 < 20; fi40++ {
            v38 += (^(v38))
            if (!true) {
                for fi41 := 0; fi41 < 18; fi41++ {
                    if true {
                        fmt.Println(v38, v4)
                        if true {
                            v4 += v4
                        }
                        var v42 int = fi40
                        _ = v42
                    } else {
                        delete(v1, 1)
                        v3 = 4294967295
                        v43 := GenS1{F0: (v3 & 2), F1: (v38 - (v38 / ((v38) | 1))), F2: 167}
                        _ = v43
                    }
                    var v44 int8 = v4
                    _ = fi41
                    _ = v44
                }
            } else {
                for fi45 := 0; fi45 < 13; fi45++ {
                    var v46 uint8 = 128
                    _ = fi45
                    _ = v46
                }
            }
            if true {
                v47 := []int{4294967295, -9223372036854775808, fi40}
                _ = v47
            }
            _ = fi40
        }
        _ = v39
    }
    fmt.Println(v38, v2, v4)
    if ((v4 - (v4 * v4)) > (^(v4))) {
        v4 -= (v4 / ((v4) | 1))
    } else {
        v48 := GenS5{F0: -1, F1: 42}
        _ = v48
    }
    var v49 int = 0
    var v50 int64 = 788
    v2 = v2
    var v51 float64 = ((0.1 / (1e10 / 0.1)) * float64((v2 + ((-2.25 - (v2 - 0.0)) * (v2 + v2)))))
    for fi52 := 0; fi52 < 11; fi52++ {
        if (972 <= 65535) {
            var v53 uint32 = 2
            _ = v53
        } else {
            var v54 uint16 = 32767
            _ = v54
        }
        _ = fi52
    }
    if true {
        v51 = 1e10
        v50 = 0
    }
    var v55 uint16 = uint16(v3)
    switch (v49 / ((int(v55)) | 1)) {
    case 638:
        v56 := []int{v49, 9223372036854775807}
        _ = v56
    }
    for fi57 := 0; fi57 < 8; fi57++ {
        if (-1 <= 647) {
            v2 = v2
            if false {
                if false {
                    v58 := GenS4{F0: 1, F1: genFn1((fi57 % (((fi57 >> ((65535) & 63))) | 1)), ("" + "\n\t"), 2147483647), F2: v49}
                    _ = v58
                }
                for fi59 := 0; fi59 < 13; fi59++ {
                    var v60 string = "\n\t"
                    _ = fi59
                    _ = v60
                }
            } else {
                var v61 uint8 = uint8(v38)
                _ = v61
            }
        } else {
            if (1000 >= 1) {
                v62 := []int{(-(fi57)), 128}
                var v63 int32 = 0
                v64 := GenS6{F0: 299, F1: "a"}
                _ = v62
                _ = v63
                _ = v64
            }
        }
        if true {
            fmt.Println(fi57, v2, v51)
            fmt.Println(v49, v49)
        }
        _ = fi57
    }
    for fi65 := 0; fi65 < 5; fi65++ {
        switch (fi65 % (((fi65 / ((v49) | 1))) | 1)) {
        case 147:
            var v66 uint16 = (v55 - genFn1((v49 & (fi65 * -917)), "a", 735))
            _ = v66
        case 778:
            v50 -= -793
        case -1:
            fmt.Println(v38)
        }
        _ = fi65
    }
    fmt.Println(v55, v55, v50)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("v55:", v55)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v49
    _ = v50
    _ = v51
    _ = v55
}
