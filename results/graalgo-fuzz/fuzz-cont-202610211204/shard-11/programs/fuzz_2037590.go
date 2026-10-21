package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint; F2 uint8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int; F1 int64; F2 uint64 }
type GenS4 struct { F0 uint; F1 int8 }
type GenS5 struct { F0 int16; F1 uint8 }
type GenS6 struct { F0 int16 }
func genFn0(p0 uint) uint16 {
    p0 = (p0 % (((p0 * 4294967295)) | 1))
    p0 -= (p0 % ((p0) | 1))
    var v4 int = 42
    switch 32768 {
    case 4294967295:
        for fi5 := 0; fi5 < 3; fi5++ {
            for fi6 := 0; fi6 < 8; fi6++ {
                v7 := []int{v4}
                var v8 string = (("foo bar" + (("a" + ("\n\t" + "a")) + "")) + "hello")
                if ((3.14159 / 1.0) < -2.25) {
                    var v9 uint32 = uint32(p0)
                    v8 = v8
                    _ = v9
                }
                v10 := GenS1{F0: 78, F1: (p0 & p0), F2: 255}
                _ = fi6
                _ = v10
                _ = v7
                _ = v8
            }
            if (false || (-1 < 923)) {
                var v11 bool = (true || ((-2.25 > (1e10 / 1e10)) && (0.0 != ((float32(0.1) / float32(((0.1 + 1.5) / 1.5))) / 1e10))))
                for fi12 := 0; fi12 < 10; fi12++ {
                    if ((!(65535 == (p0 ^ p0))) || (p0 != p0)) {
                        if v11 {
                            v13 := GenS2{F0: ("\n\t" + ("\n\t" + ("a" + "")))}
                            v14 := []int{fi5}
                            v14 = append(v14, fi12)
                            _ = v13
                            _ = v14
                        } else {
                            switch ("hello" + "\n\t") {
                            case "foo bar":
                                if v11 {
                                    v15 := []int{}
                                    _ = v15
                                }
                            }
                            var v16 int64 = 9223372036854775807
                            _ = v16
                        }
                        var v17 int16 = 32767
                        if v11 {
                            if true {
                                fmt.Println(v11, v4, p0)
                                if (2147483647 >= 89) {
                                    fmt.Println(fi5, v11, p0)
                                    if true {
                                        var v18 uint16 = 344
                                        _ = v18
                                    } else {
                                        v19 := GenS1{F0: 2, F1: (p0 & 2), F2: 42}
                                        v11 = v11
                                        _ = v19
                                    }
                                    if v11 {
                                        v4 = (-(v4))
                                    }
                                }
                                var v20 uint = p0
                                _ = v20
                            } else {
                                var v21 uint32 = 0
                                _ = v21
                            }
                            var v22 int = (v4 * (v4 - (v4 - fi12)))
                            _ = v22
                        } else {
                            v23 := []int{(v4 + (fi5 & (v4 * 9223372036854775807)))}
                            v24 := GenS2{F0: ("" + "\n\t")}
                            _ = v23
                            _ = v24
                        }
                        p0 = 208
                        if (!v11) {
                            switch fi5 {
                            case 747:
                                var v25 int32 = 306
                                _ = v25
                            case -590:
                                if v11 {
                                    var v26 int64 = -1
                                    _ = v26
                                }
                            }
                            if v11 {
                                var v27 int64 = 1
                                var v28 uint8 = 128
                                _ = v27
                                _ = v28
                            }
                        } else {
                            v29 := GenS3{F0: (^(v4)), F1: 9223372036854775807, F2: 42}
                            if false {
                                v30 := []int{256, v4, 0}
                                switch v4 {
                                case 128:
                                    var v31 float32 = (1e10 - -2.25)
                                    _ = v31
                                }
                                var v32 uint32 = 123
                                _ = v30
                                _ = v32
                            }
                            p0 -= p0
                            _ = v29
                        }
                        var v33 int = -63
                        p0 += (p0 & p0)
                        _ = v17
                        _ = v33
                    } else {
                        v34 := []int{v4, (v4 | fi5), -2147483648, fi12}
                        _ = v34
                    }
                    if v11 {
                        v35 := GenS3{F0: (^(v4)), F1: -9223372036854775808, F2: 402}
                        _ = v35
                    } else {
                        if v11 {
                            v36 := []int{}
                            var v37 int64 = -1
                            _ = v36
                            _ = v37
                        } else {
                            v38 := GenS4{F0: (p0 ^ 79), F1: 42}
                            v11 = v11
                            v39 := []int{(^(fi12)), fi5, fi12}
                            _ = len(v39)
                            fmt.Println(fi5, fi5)
                            _ = v38
                            _ = v39
                        }
                    }
                    v40 := GenS1{F0: 88, F1: p0, F2: 42}
                    _ = fi12
                    _ = v40
                }
                _ = v11
            }
            v41 := []int{2147483647, 2147483647, (v4 | fi5), v4}
            _ = fi5
            _ = v41
        }
    case 9223372036854775807:
        for fi42 := 0; fi42 < 7; fi42++ {
            p0 += 436
            v4 = (v4 % (((v4 << ((p0) & 63))) | 1))
            _ = fi42
        }
    default:
        v43 := GenS1{F0: 0, F1: p0, F2: 128}
        _ = v43
    }
    v44 := []int{v4, (v4 << (((p0 + 205)) & 63)), v4, (v4 / (((v4 & (^(v4)))) | 1))}
    _ = p0
    _ = v4
    _ = v44
    return 2
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[uint32]int)
    v3 := make(map[string]uint32)
    v45 := GenS5{F0: 0, F1: 161}
    var v46 bool = false
    for fi47 := 0; fi47 < 3; fi47++ {
        v46 = false
        v48 := []int{(fi47 * (-(fi47))), fi47}
        _ = fi47
        _ = v48
    }
    var v49 string = "foo bar"
    switch (((v49 + (v49 + (v49 + v49))) + "foo bar") + v49) {
    case "foo bar":
        if v46 {
            v49 = "foo bar"
            for fi50 := 0; fi50 < 16; fi50++ {
                if ((124 <= 2147483648) && (42 > 4294967295)) {
                    var v51 int8 = -128
                    if false {
                        if (int16(v51) > 32767) {
                            var v52 string = ("foo bar" + v49)
                            _ = v52
                        }
                        for fi53 := 0; fi53 < 14; fi53++ {
                            v54 := GenS6{F0: 1}
                            var v55 int16 = 42
                            fmt.Println(v55, v46, fi50)
                            _ = fi53
                            _ = v54
                            _ = v55
                        }
                    } else {
                        v51 += v51
                        var v56 bool = (((-1 >= 9223372036854775807) || (((!(!(v46 || v46))) || v46) && false)) && v46)
                        switch 4294967295 {
                        case 4294967295:
                            if v46 {
                                var v57 int16 = -989
                                v58 := []int{(fi50 >> ((1) & 63)), (fi50 << ((606) & 63))}
                                _ = len(v58)
                                _ = v57
                                _ = v58
                            }
                        case -2147483648:
                            if v46 {
                                for fi59 := 0; fi59 < 6; fi59++ {
                                    var v60 float64 = 1e10
                                    _ = fi59
                                    _ = v60
                                }
                                v46 = v56
                                fmt.Println(v49, fi50)
                            }
                        }
                        _ = v56
                    }
                    fmt.Println(v51, v46, v46)
                    _ = v51
                }
                var v61 int16 = 225
                var v62 uint32 = 4294967295
                _ = fi50
                _ = v61
                _ = v62
            }
        }
    case "hello":
        switch 852 {
        case 2147483647:
            for fi63 := 0; fi63 < 16; fi63++ {
                v64 := []int{fi63, fi63}
                var v65 int64 = 0
                switch (fi63 << ((507) & 63)) {
                case 4294967295:
                    var v66 int16 = -32768
                    _ = v66
                case 32767:
                    var v67 int8 = int8(v65)
                    _ = v67
                default:
                    _ = len(v3)
                }
                _ = fi63
                _ = v64
                _ = v65
            }
        case 32767:
            delete(v2, 0)
        case -603:
            v49 = (v49 + "hello")
        }
    }
    var v68 uint32 = 2147483648
    v69 := []int{32768, -879, 4294967295, -345}
    switch -276 {
    case 9223372036854775807:
        switch 4294967295 {
        case 647:
            for fi70 := 0; fi70 < 1; fi70++ {
                v46 = (137 >= 533)
                _ = fi70
            }
        case 583:
            var v71 bool = (891 < 2147483647)
            _ = v71
        default:
            fmt.Println(v46, v68)
        }
    }
    var v72 uint8 = 70
    v68 = v68
    v49 = (("\n\t" + v49) + (((v49 + (v49 + (v49 + ((v49 + v49) + "a")))) + v49) + v49))
    v73 := GenS1{F0: (v72 * v72), F1: 1, F2: (v72 & v72)}
    if len(v69) > 0 {
        v69[0] = -141
    }
    var v74 int64 = -492
    var v75 uint16 = genFn0(65535)
    fmt.Println(v49, v75)
    var v76 int8 = -128
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v46:", v46)
    fmt.Println("v49:", v49)
    fmt.Println("v68:", v68)
    fmt.Println("slicelen v69:", len(v69), v69)
    fmt.Println("v72:", v72)
    fmt.Println("v74:", v74)
    fmt.Println("v75:", v75)
    fmt.Println("v76:", v76)
    _ = v1
    _ = v2
    _ = v3
    _ = v45
    _ = v46
    _ = v49
    _ = v68
    _ = v69
    _ = v72
    _ = v73
    _ = v74
    _ = v75
    _ = v76
}
