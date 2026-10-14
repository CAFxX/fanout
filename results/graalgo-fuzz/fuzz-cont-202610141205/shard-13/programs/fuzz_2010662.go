package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64 }
type GenS2 struct { F0 uint32; F1 uint32 }
type GenS3 struct { F0 uint16; F1 uint64 }
type GenS4 struct { F0 int32; F1 int; F2 int8 }
func genFn0() uint64 {
    if false {
        var v4 float32 = 1.5
        v5 := []int{-1, -2147483648, 0, 128}
        var v6 uint16 = 2
        _ = v4
        _ = v5
        _ = v6
    }
    v7 := []int{}
    _ = v7
    return 1
}
func genFn1(p0 bool, p1 int64) uint64 {
    switch 2147483647 {
    case -1:
        fmt.Println(p0, p1)
    }
    switch "hello" {
    case "hello":
        switch 131 {
        case -1:
            v8 := GenS1{F0: false, F1: 1}
            _ = v8
        }
    case "":
        v9 := GenS1{F0: p0, F1: -144}
        _ = v9
    default:
        p0 = p0
    }
    fmt.Println(p1, p0)
    v10 := []int{0, -2147483648, 32768}
    switch "" {
    case "hello":
        for fi11 := 0; fi11 < 17; fi11++ {
            p0 = (-2.25 > 0.0)
            var v12 float32 = (1.5 / (1e10 / -2.25))
            var v13 int32 = 0
            _ = fi11
            _ = v12
            _ = v13
        }
    case "\n\t":
        var v14 uint8 = 128
        _ = v14
    default:
        p0 = (!true)
    }
    _ = p0
    _ = p1
    _ = v10
    return 1
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint = 790
    var v3 int16 = -889
    v2 = 729
    v3 = v3
    if (766 != 9223372036854775807) {
        delete(v1, ("a" + "a"))
        for fi15 := 0; fi15 < 10; fi15++ {
            v16 := []int{(fi15 + (fi15 / ((fi15) | 1))), 4294967295}
            _ = fi15
            _ = v16
        }
        switch "a" {
        case "foo bar":
            for fi17 := 0; fi17 < 1; fi17++ {
                var v18 bool = ((fi17 != 1) && (-12 < 42))
                var v19 uint = (v2 + v2)
                _ = fi17
                _ = v18
                _ = v19
            }
        case "a":
            switch "\n\t" {
            case "foo bar":
                var v20 uint16 = 65535
                _ = v20
            case "hello":
                var v21 uint16 = 32767
                _ = v21
            case "a":
                for fi22 := 0; fi22 < 19; fi22++ {
                    if true {
                        v23 := GenS2{F0: 2, F1: 929}
                        _ = v23
                    }
                    if false {
                        for fi24 := 0; fi24 < 10; fi24++ {
                            v25 := []int{}
                            var v26 float32 = 1.5
                            switch ("" + ("\n\t" + ("a" + "\n\t"))) {
                            case "\n\t":
                                var v27 int64 = 696
                                _ = v27
                            case "":
                                delete(v1, "")
                            case "hello":
                                var v28 bool = ((("hello" + "hello") != ("a" + ("a" + ""))) && ((255 >= 32767) || ((v26 * 1e10) <= (v26 / (v26 + v26)))))
                                _ = v28
                            }
                            _ = fi24
                            _ = v25
                            _ = v26
                        }
                        switch (fi22 / ((1) | 1)) {
                        case -892:
                            switch "a" {
                            case "\n\t":
                                v3 -= v3
                            case "":
                                if ((-2.25 != ((0.1 + 1.5) + float64(float32((float64(-2.25) + float64((0.1 * 1e10))))))) && false) {
                                    v2 += v2
                                    v29 := GenS3{F0: 246, F1: 349}
                                    var v30 int8 = -128
                                    _ = v29
                                    _ = v30
                                }
                            }
                        case -958:
                            for fi31 := 0; fi31 < 6; fi31++ {
                                if true {
                                    v32 := []int{fi22, fi31, int(v3), (^(fi31))}
                                    _ = v32
                                } else {
                                    var v33 uint = (v2 % (((v2 ^ 809)) | 1))
                                    v34 := []int{(fi31 * (-(fi22))), (fi22 & (fi31 << (((^(v33))) & 63))), fi22, -2147483648}
                                    fmt.Println(v3, v3)
                                    _ = v33
                                    _ = v34
                                }
                                v3 += 1
                                if (-110 < -69) {
                                    fmt.Println(fi31, v2, fi31)
                                    var v35 uint16 = 32767
                                    v3 = v3
                                    _ = v35
                                } else {
                                    fmt.Println(fi22, v3)
                                    v36 := []int{fi31, 978, fi31}
                                    _ = v36
                                }
                                v3 = v3
                                _ = fi31
                            }
                        case 32768:
                            var v37 uint = (v2 / ((v2) | 1))
                            _ = v37
                        default:
                            if true {
                                var v38 int8 = 100
                                if (1 == 42) {
                                    for fi39 := 0; fi39 < 10; fi39++ {
                                        v40 := GenS3{F0: 508, F1: genFn1((!true), -9223372036854775808)}
                                        _ = fi39
                                        _ = v40
                                    }
                                    v1["hello"] = false
                                }
                                _ = v38
                            } else {
                                if (v3 >= (v3 ^ -845)) {
                                    for fi41 := 0; fi41 < 15; fi41++ {
                                        v42 := GenS1{F0: (307 < 2), F1: -9223372036854775808}
                                        _ = fi41
                                        _ = v42
                                    }
                                    v2 -= v2
                                    _ = len(v1)
                                }
                            }
                        }
                    } else {
                        v43 := GenS1{F0: false, F1: 126}
                        v44 := GenS4{F0: 1, F1: 32767, F2: 127}
                        if ((v3 + (^(v3))) >= (v3 * -168)) {
                            var v45 int16 = v3
                            v2 += v2
                            _ = v45
                        } else {
                            v2 = 790
                            v2 = v2
                        }
                        _ = v43
                        _ = v44
                    }
                    _ = fi22
                }
            }
        case "":
            v2 = (v2 * (v2 / ((59) | 1)))
        default:
            fmt.Println(v2, v3)
        }
    } else {
        var v46 int8 = 109
        fmt.Println(v3)
        for fi47 := 0; fi47 < 17; fi47++ {
            for fi48 := 0; fi48 < 6; fi48++ {
                v3 -= 1
                v49 := []int{0, (-(fi47)), (fi48 - (-(fi47)))}
                v50 := GenS2{F0: 65535, F1: 2}
                switch 127 {
                case 32768:
                    fmt.Println(fi47)
                default:
                    var v51 float32 = 1.5
                    _ = v51
                }
                v52 := GenS2{F0: uint32(v2), F1: 2147483648}
                var v53 uint = (v2 + 672)
                _ = fi48
                _ = v49
                _ = v50
                _ = v52
                _ = v53
            }
            v3 -= (v3 & 1)
            _ = fi47
        }
        _ = v46
    }
    for fi54 := 0; fi54 < 16; fi54++ {
        v55 := []int{fi54, (^(fi54))}
        _ = fi54
        _ = v55
    }
    v56 := []int{719, -1, 2147483647}
    for fi57 := 0; fi57 < 2; fi57++ {
        var v58 float64 = -2.25
        _ = fi57
        _ = v58
    }
    var v59 uint64 = 909
    var v60 int64 = 411
    v61 := v1[""]
    switch int(v2) {
    case 9223372036854775807:
        if v61 {
            if (v3 > (-(v3))) {
                for fi62 := 0; fi62 < 20; fi62++ {
                    delete(v1, ("" + ("" + ("\n\t" + ""))))
                    if (32767 > uint16(v2)) {
                        if len(v56) >= 2 {
                            v56 = v56[1:len(v56)]
                        }
                    }
                    _ = fi62
                }
            }
        } else {
            _ = len(v56)
            v56 = append(v56, -9223372036854775808)
        }
    default:
        v59 += v59
    }
    for fi63 := 0; fi63 < 1; fi63++ {
        v56 = append(v56, (fi63 / (((^(fi63))) | 1)))
        v2 = v2
        _ = fi63
    }
    v3 = v3
    var v64 float64 = (float64(0.0) * ((3.14159 - ((1e10 / float64(-2.25)) * float64(0.1))) / float64(1.5)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v56:", len(v56), v56)
    fmt.Println("v59:", v59)
    fmt.Println("v60:", v60)
    fmt.Println("v61:", v61)
    fmt.Println("v64:", v64)
    _ = v1
    _ = v2
    _ = v3
    _ = v56
    _ = v59
    _ = v60
    _ = v61
    _ = v64
}
