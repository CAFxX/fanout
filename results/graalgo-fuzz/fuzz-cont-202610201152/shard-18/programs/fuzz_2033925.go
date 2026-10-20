package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8 }
type GenS2 struct { F0 int16; F1 bool; F2 uint16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 string }

func main() {
    var v1 float64 = 3.14159
    v2 := make(map[int]uint32)
    var v3 int64 = -9223372036854775808
    var v4 int = 42
    fmt.Println(v3)
    var v5 int16 = 1
    switch "foo bar" {
    case "a":
        v6 := GenS1{F0: 2, F1: 42}
        _ = v6
    case "hello":
        if false {
            var v7 int16 = v5
            _ = v7
        } else {
            var v8 uint = 18446744073709551615
            v3 += -9223372036854775808
            if (!false) {
                if false {
                    var v9 uint64 = 42
                    v5 += -771
                    _ = v9
                }
                v2[v4] = 1
            } else {
                if (!false) {
                    delete(v2, v4)
                    var v10 int8 = 127
                    if true {
                        for fi11 := 0; fi11 < 12; fi11++ {
                            delete(v2, (v4 - (fi11 >> ((v8) & 63))))
                            _ = fi11
                        }
                    } else {
                        for fi12 := 0; fi12 < 5; fi12++ {
                            switch (v4 % ((v4) | 1)) {
                            case 32767:
                                var v13 uint8 = 1
                                _ = v13
                            case -72:
                                v1 = 0.1
                            default:
                                var v14 int32 = 966
                                _ = v14
                            }
                            v15 := GenS2{F0: (^(v5)), F1: (!true), F2: 255}
                            var v16 int64 = (v3 ^ 1)
                            _ = fi12
                            _ = v15
                            _ = v16
                        }
                    }
                    _ = v10
                } else {
                    var v17 bool = (!(v8 != (^(v8))))
                    for fi18 := 0; fi18 < 17; fi18++ {
                        switch (fi18 << ((v8) & 63)) {
                        case -2147483648:
                            for fi19 := 0; fi19 < 4; fi19++ {
                                fmt.Println(v1, v4, v4)
                                var v20 int32 = 0
                                v21 := GenS3{F0: (^(v5))}
                                _ = fi19
                                _ = v20
                                _ = v21
                            }
                        case 869:
                            v17 = v17
                        default:
                            switch (fi18 - fi18) {
                            case 42:
                                var v22 float32 = ((-2.25 * 1e10) / -2.25)
                                _ = v22
                            case 365:
                                v23 := []int{fi18, -2147483648, (v4 / (((fi18 + v4)) | 1))}
                                _ = v23
                            default:
                                if false {
                                    fmt.Println(v17, v3, v5)
                                    for fi24 := 0; fi24 < 18; fi24++ {
                                        v17 = (v17 || (false || false))
                                        _ = fi24
                                    }
                                    if false {
                                        v1 = (v1 + ((v1 * 0.1) * v1))
                                        if v17 {
                                            if ((-(v3)) > (v3 + 9223372036854775807)) {
                                                v4 = fi18
                                            } else {
                                                for fi25 := 0; fi25 < 17; fi25++ {
                                                    v26 := GenS3{F0: (v5 % (((v5 + v5)) | 1))}
                                                    var v27 int8 = -128
                                                    fmt.Println(fi18)
                                                    _ = fi25
                                                    _ = v26
                                                    _ = v27
                                                }
                                                if v17 {
                                                    var v28 int16 = v5
                                                    _ = v28
                                                }
                                                var v29 int16 = (^(v5))
                                                _ = v29
                                            }
                                            v8 -= (^(v8))
                                            if v17 {
                                                v8 = (v8 >> (((v8 * 18446744073709551615)) & 63))
                                                v4 += v4
                                                fmt.Println(v17)
                                            } else {
                                                var v30 int16 = v5
                                                v2[(v4 % (((fi18 + -9223372036854775808)) | 1))] = 2
                                                v17 = (false || (!v17))
                                                _ = v30
                                            }
                                        }
                                        v3 += (v3 & (v3 + 42))
                                    } else {
                                        fmt.Println(v5)
                                    }
                                }
                            }
                        }
                        for fi31 := 0; fi31 < 10; fi31++ {
                            var v32 int32 = -1
                            delete(v2, fi18)
                            var v33 float32 = ((1.5 / (float32(1e10) + ((((1.5 / 1.5) + 3.14159) + (0.1 * -2.25)) + (((-2.25 * 0.1) * 3.14159) / (-2.25 + (3.14159 * -2.25)))))) / 1e10)
                            _ = fi31
                            _ = v32
                            _ = v33
                        }
                        if false {
                            v4 += -986
                        }
                        _ = fi18
                    }
                    if v17 {
                        var v34 int16 = v5
                        v3 += 9223372036854775807
                        _ = v34
                    } else {
                        if v17 {
                            fmt.Println(v1, v5)
                        } else {
                            for fi35 := 0; fi35 < 10; fi35++ {
                                fmt.Println(v5)
                                var v36 bool = (true && (406 != 65535))
                                fmt.Println(v4, v4)
                                _ = fi35
                                _ = v36
                            }
                            v1 -= v1
                        }
                    }
                    _ = v17
                }
                v4 -= (v4 ^ (-(v4)))
            }
            _ = v8
        }
    default:
        if false {
            v37 := GenS2{F0: (v5 << ((65535) & 63)), F1: ((v5 >> ((15) & 63)) >= -905), F2: 1}
            _ = v37
        }
    }
    var v38 int = v4
    var v39 string = ("foo bar" + "a")
    if true {
        if (255 < 1) {
            switch -1 {
            case 32768:
                var v40 int32 = 0
                _ = v40
            case -252:
                if false {
                    v41 := []int{}
                    _ = v41
                }
            case 42:
                v38 = (v38 | (v38 + v4))
            }
            switch v4 {
            case 1:
                v42 := []int{v4, 1, (v38 ^ (v38 / (((v4 * (v4 * (v38 * (v38 | (v4 + -2147483648)))))) | 1))), 4294967295}
                _ = v42
            case 0:
                v2[(v38 & -219)] = 2147483647
            }
        }
        v43 := GenS3{F0: (v5 ^ 42)}
        _ = v43
    } else {
        for fi44 := 0; fi44 < 20; fi44++ {
            v45 := GenS2{F0: v5, F1: false, F2: 967}
            if (!true) {
                v1 -= -2.25
                if ((v3 / ((v3) | 1)) < v3) {
                    if (((^(v5)) <= 32767) || true) {
                        if true {
                            v46 := []int{-2147483648}
                            v1 += v1
                            _ = v46
                        } else {
                            fmt.Println(v1)
                            if false {
                                for fi47 := 0; fi47 < 16; fi47++ {
                                    var v48 float32 = -2.25
                                    v49 := GenS2{F0: 42, F1: (516 == 0), F2: 0}
                                    if true {
                                        if (!(1 == 42)) {
                                            var v50 int8 = -128
                                            _ = v50
                                        } else {
                                            v51 := GenS4{F0: ""}
                                            var v52 int64 = (v3 | 758)
                                            var v53 int = fi47
                                            fmt.Println(v5, v3)
                                            _ = v51
                                            _ = v52
                                            _ = v53
                                        }
                                    } else {
                                        v39 = ("a" + v39)
                                    }
                                    _ = fi47
                                    _ = v48
                                    _ = v49
                                }
                            } else {
                                fmt.Println(v39)
                                v39 = "a"
                                v3 += v3
                            }
                        }
                        v54 := GenS3{F0: v5}
                        _ = v54
                    } else {
                        for fi55 := 0; fi55 < 12; fi55++ {
                            fmt.Println(fi44, v5)
                            _ = fi55
                        }
                    }
                    switch v39 {
                    case "a":
                        var v56 float32 = -2.25
                        _ = v56
                    }
                    var v57 float32 = float32(v1)
                    _ = v57
                }
            } else {
                switch v39 {
                case "hello":
                    var v58 int16 = -111
                    _ = v58
                }
                v4 = fi44
                v5 = v5
            }
            v59 := GenS3{F0: v5}
            _ = fi44
            _ = v45
            _ = v59
        }
    }
    v38 -= v4
    var v60 uint32 = 0
    for fi61 := 0; fi61 < 7; fi61++ {
        var v62 float32 = float32(float64(0.1))
        var v63 uint8 = 128
        _ = fi61
        _ = v62
        _ = v63
    }
    fmt.Println(v1, v5)
    v60 = (v60 / ((v60) | 1))
    switch v39 {
    case "foo bar":
        v64 := []int{1, v4}
        _ = v64
    case "a":
        var v65 uint32 = v60
        _ = v65
    }
    v4 = v4
    v4 = v38
    v66 := []int{2147483647, 32767, (^(v38))}
    var v67 int32 = -592
    var v68 bool = false
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v60:", v60)
    fmt.Println("slicelen v66:", len(v66), v66)
    fmt.Println("v67:", v67)
    fmt.Println("v68:", v68)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v60
    _ = v66
    _ = v67
    _ = v68
}
