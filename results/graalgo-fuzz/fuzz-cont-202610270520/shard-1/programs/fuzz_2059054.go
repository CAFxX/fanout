package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16 }
type GenS2 struct { F0 int32; F1 uint8; F2 bool }
type GenS3 struct { F0 int16; F1 int64; F2 bool }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 bool; F1 int }
type GenS6 struct { F0 uint32; F1 int32; F2 int16 }
func genFn0() int16 {
    switch "a" {
    case "":
        v5 := []int{0, 9223372036854775807}
        _ = v5
    case "\n\t":
        var v6 float32 = 0.1
        _ = v6
    case "hello":
        v7 := []int{-2147483648}
        _ = v7
    default:
        v8 := GenS1{F0: 127, F1: -32768}
        _ = v8
    }
    if (!true) {
        var v9 uint16 = 792
        var v10 float64 = -2.25
        _ = v10
        _ = v9
    }
    return 42
}

func main() {
    var v1 int8 = 100
    var v2 string = "a"
    var v3 uint32 = 426
    v4 := make(map[string]uint32)
    v11 := v4[((v2 + "\n\t") + v2)]
    v11 = (^(v3))
    var v12 uint64 = 42
    if (-2.25 > (3.14159 - 0.1)) {
        if (!(962 > 1)) {
            v13 := GenS1{F0: 4, F1: 233}
            if (true && (((1e10 - 0.0) - 3.14159) <= 1e10)) {
                for fi14 := 0; fi14 < 8; fi14++ {
                    for fi15 := 0; fi15 < 4; fi15++ {
                        fmt.Println(v12, fi15)
                        var v16 bool = true
                        _ = fi15
                        _ = v16
                    }
                    v17 := []int{9223372036854775807, 32768, fi14, (fi14 * (fi14 ^ fi14))}
                    _ = fi14
                    _ = v17
                }
            }
            _ = v13
        } else {
            var v18 int32 = 2147483647
            var v19 uint64 = (v12 + v12)
            for fi20 := 0; fi20 < 14; fi20++ {
                v19 -= 126
                if false {
                    v19 = (v12 * 781)
                } else {
                    v21 := GenS2{F0: v18, F1: 15, F2: (0.1 > 1e10)}
                    fmt.Println(v19, v2, fi20)
                    fmt.Println(v1, v19)
                    _ = v21
                }
                _ = fi20
            }
            _ = v18
            _ = v19
        }
    } else {
        v2 = (v2 + (((v2 + v2) + ((v2 + "a") + v2)) + v2))
    }
    var v22 uint16 = 297
    switch ((("" + v2) + v2) + v2) {
    case "hello":
        var v23 uint8 = 127
        _ = v23
    case "\n\t":
        _ = len(v4)
    default:
        v22 = (v22 ^ (^(v22)))
    }
    for fi24 := 0; fi24 < 6; fi24++ {
        if false {
            switch v2 {
            case "hello":
                v25 := []int{-2147483648, (fi24 & fi24)}
                _ = v25
            case "a":
                fmt.Println(v1, v1)
            default:
                var v26 uint64 = v12
                _ = v26
            }
            var v27 int32 = -966
            _ = v27
        }
        v28 := []int{fi24, fi24, int(v12)}
        _ = len(v4)
        _ = fi24
        _ = v28
    }
    var v29 uint16 = (v22 >> ((65535) & 63))
    switch -1 {
    case 42:
        v22 += (v29 / (((v22 << ((0) & 63))) | 1))
    case 169:
        var v30 uint = 42
        _ = v30
    case 9223372036854775807:
        if false {
            if ((3.14159 - (float32(float64(((-2.25 + 1.5) - 1e10))) - float32(((1.5 + 0.0) + -2.25)))) > 0.1) {
                for fi31 := 0; fi31 < 1; fi31++ {
                    v3 += 881
                    _ = fi31
                }
                for fi32 := 0; fi32 < 19; fi32++ {
                    v33 := GenS2{F0: 984, F1: 94, F2: (!("hello" == (v2 + "\n\t")))}
                    v34 := GenS3{F0: -111, F1: 417, F2: (0 >= 810)}
                    _ = fi32
                    _ = v33
                    _ = v34
                }
            } else {
                if false {
                    fmt.Println(v2, v22)
                    v35 := GenS2{F0: -835, F1: 255, F2: (-2.25 == (1.5 + 1e10))}
                    _ = v35
                } else {
                    var v36 float32 = float32(((1e10 / 3.14159) * 3.14159))
                    var v37 int = -2147483648
                    _ = v36
                    _ = v37
                }
                v2 = ("hello" + (((("a" + "foo bar") + ((("hello" + "\n\t") + v2) + v2)) + v2) + v2))
                v29 = (v22 + (v29 - 255))
            }
            for fi38 := 0; fi38 < 6; fi38++ {
                v22 = (v22 % ((353) | 1))
                if true {
                    if true {
                        switch fi38 {
                        case -315:
                            fmt.Println(v1, v11, fi38)
                        }
                        switch ("hello" + ("" + v2)) {
                        case "\n\t":
                            for fi39 := 0; fi39 < 13; fi39++ {
                                switch "\n\t" {
                                case "":
                                    v40 := GenS4{F0: 0}
                                    _ = v40
                                }
                                _ = fi39
                            }
                        case "":
                            if (int8(v22) <= 100) {
                                if (((fi38 << ((4294967295) & 63)) <= (fi38 ^ (^(fi38)))) || ((v22 | (^(v22))) > uint16(v3))) {
                                    var v41 uint16 = v29
                                    _ = v41
                                } else {
                                    var v42 float32 = (-2.25 * 0.0)
                                    var v43 float32 = v42
                                    _ = v42
                                    _ = v43
                                }
                            } else {
                                for fi44 := 0; fi44 < 10; fi44++ {
                                    v45 := GenS4{F0: -1}
                                    if false {
                                        switch (^(fi44)) {
                                        case 9223372036854775807:
                                            if false {
                                                var v46 bool = true
                                                fmt.Println(v11, v29, v12)
                                                _ = v46
                                            } else {
                                                delete(v4, (v2 + ((v2 + v2) + v2)))
                                                fmt.Println(v22, v12, v22)
                                            }
                                        default:
                                            v47 := []int{32767}
                                            _ = v47
                                        }
                                    }
                                    var v48 uint = 1
                                    _ = fi44
                                    _ = v45
                                    _ = v48
                                }
                                switch (fi38 >> ((18446744073709551615) & 63)) {
                                case -294:
                                    for fi49 := 0; fi49 < 1; fi49++ {
                                        if (fi49 >= fi49) {
                                            v29 += v29
                                            delete(v4, (v2 + ("" + "foo bar")))
                                            if (true && (true || true)) {
                                                v50 := []int{}
                                                _ = v50
                                            } else {
                                                var v51 string = (v2 + v2)
                                                _ = v51
                                            }
                                        } else {
                                            v2 = ("\n\t" + v2)
                                            if true {
                                                v4[v2] = 2147483647
                                            }
                                            v11 -= v11
                                        }
                                        if false {
                                            if (true || true) {
                                                v4[v2] = (v3 << ((1) & 63))
                                                if false {
                                                    var v52 int = (fi49 / ((fi38) | 1))
                                                    v53 := GenS4{F0: -1}
                                                    if (-2.25 > -2.25) {
                                                        v54 := GenS3{F0: 364, F1: -512, F2: (true || (!false))}
                                                        v22 = uint16(v12)
                                                        var v55 int = (v52 % (((fi49 - (fi38 % (((fi49 / (((fi49 & fi38)) | 1))) | 1)))) | 1))
                                                        _ = v54
                                                        _ = v55
                                                    } else {
                                                        v4[v2] = 4294967295
                                                        if (fi49 > (-(fi38))) {
                                                            fmt.Println(v2, v12, fi49)
                                                        }
                                                        v1 += 100
                                                    }
                                                    switch (fi49 + (-(fi38))) {
                                                    case 2147483647:
                                                        v56 := []int{fi49, (^(fi38)), (v52 ^ fi49)}
                                                        _ = v56
                                                    default:
                                                        v57 := GenS5{F0: true, F1: (v52 ^ 127)}
                                                        _ = v57
                                                    }
                                                    _ = v52
                                                    _ = v53
                                                }
                                            } else {
                                                v58 := GenS6{F0: 479, F1: 1, F2: -536}
                                                fmt.Println(v1, v1, v12)
                                                v2 = ("" + "")
                                                _ = v58
                                            }
                                        }
                                        _ = fi49
                                    }
                                case -159:
                                    delete(v4, ((v2 + v2) + "hello"))
                                }
                                v59 := GenS4{F0: 0}
                                _ = v59
                            }
                        case "hello":
                            var v60 int16 = -1
                            _ = v60
                        default:
                            var v61 uint8 = 178
                            _ = v61
                        }
                        v62 := []int{fi38}
                        _ = v62
                    }
                }
                _ = fi38
            }
        } else {
            v63 := GenS6{F0: 0, F1: 42, F2: 42}
            _ = v63
        }
    default:
        v11 = v3
    }
    delete(v4, v2)
    var v64 bool = true
    v65 := GenS4{F0: 42}
    if (531 != 32767) {
        var v66 uint = 338
        var v67 string = "foo bar"
        _ = len(v4)
        _ = v66
        _ = v67
    }
    for fi68 := 0; fi68 < 7; fi68++ {
        fmt.Println(v2, v22)
        _ = fi68
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v64:", v64)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v22
    _ = v29
    _ = v3
    _ = v4
    _ = v64
    _ = v65
}
