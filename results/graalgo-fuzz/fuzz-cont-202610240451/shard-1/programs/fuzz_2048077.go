package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32; F2 uint16 }
type GenS2 struct { F0 uint64; F1 string; F2 uint8 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int32; F1 uint64 }
type GenS5 struct { F0 uint; F1 uint64; F2 uint8 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 uint64; F1 int }

func main() {
    var v1 int16 = 42
    var v2 uint32 = 2
    var v3 uint16 = 0
    var v4 int32 = -281
    if true {
        v4 = v4
        fmt.Println(v4, v3)
    } else {
        fmt.Println(v3, v4, v4)
        v5 := GenS1{F0: v4, F1: (v2 % ((0) | 1)), F2: v3}
        for fi6 := 0; fi6 < 20; fi6++ {
            var v7 string = ("a" + ("\n\t" + "foo bar"))
            if ((v4 * (v4 / (((v4 >> ((661) & 63))) | 1))) > (-(v4))) {
                v3 = (^(v3))
                var v8 uint = 2
                _ = v8
            }
            _ = fi6
            _ = v7
        }
        _ = v5
    }
    var v9 int32 = v4
    v10 := GenS1{F0: v9, F1: (v2 << ((760) & 63)), F2: uint16(v2)}
    v3 += (v3 % ((v3) | 1))
    var v11 string = "foo bar"
    v12 := GenS1{F0: (v4 + v4), F1: v2, F2: v3}
    if (false && (false || false)) {
        v1 = v1
        if true {
            var v13 uint64 = 18446744073709551615
            var v14 string = v11
            for fi15 := 0; fi15 < 17; fi15++ {
                switch v14 {
                case "foo bar":
                    v16 := []int{}
                    _ = v16
                case "hello":
                    for fi17 := 0; fi17 < 17; fi17++ {
                        if (fi17 != 128) {
                            v18 := []int{(fi17 >> ((255) & 63)), fi15, 4294967295, (fi17 | (fi15 << ((0) & 63)))}
                            if len(v18) > 0 {
                                v18[0] = (fi15 >> ((42) & 63))
                            }
                            fmt.Println(v13, fi15)
                            if true {
                                if (false || false) {
                                    if false {
                                        var v19 int16 = (v1 - v1)
                                        var v20 uint16 = (v3 & (v3 >> ((65535) & 63)))
                                        _ = v19
                                        _ = v20
                                    }
                                    var v21 int64 = -9223372036854775808
                                    _ = v21
                                }
                                fmt.Println(v3, fi17, fi17)
                                v3 += (v3 % ((v3) | 1))
                                if len(v18) >= 2 {
                                    v18 = v18[1:len(v18)]
                                }
                                var v22 int = -242
                                _ = v22
                            }
                            v1 = (-(v1))
                            v3 = v3
                            _ = v18
                        } else {
                            if (float64(3.14159) < (0.1 + float64(0.0))) {
                                var v23 int16 = -806
                                _ = v23
                            } else {
                                if true {
                                    v24 := GenS2{F0: v13, F1: v11, F2: 255}
                                    var v25 int16 = 47
                                    var v26 int = (fi17 >> ((42) & 63))
                                    _ = v24
                                    _ = v25
                                    _ = v26
                                } else {
                                    if false {
                                        v27 := GenS2{F0: (v13 % (((v13 * v13)) | 1)), F1: (v11 + v14), F2: 155}
                                        v11 = (v11 + v14)
                                        _ = v27
                                    } else {
                                        v9 += 42
                                        switch ("hello" + v11) {
                                        case "":
                                            v28 := GenS3{F0: v9}
                                            _ = v28
                                        default:
                                            v29 := []int{(fi17 + (fi15 / ((fi17) | 1)))}
                                            _ = v29
                                        }
                                    }
                                    if (false || true) {
                                        switch v11 {
                                        case "":
                                            if false {
                                                fmt.Println(v4, fi17, v1)
                                                var v30 int64 = -1
                                                v31 := []int{fi17}
                                                _ = v30
                                                _ = v31
                                            } else {
                                                var v32 float32 = 3.14159
                                                _ = v32
                                            }
                                        case "foo bar":
                                            switch (v14 + v11) {
                                            case "\n\t":
                                                v33 := []int{(fi17 + fi15), fi17, fi17}
                                                _ = v33
                                            case "":
                                                var v34 int32 = (v4 | 0)
                                                _ = v34
                                            case "hello":
                                                if true {
                                                    fmt.Println(v3)
                                                    switch (fi15 ^ 128) {
                                                    case 32767:
                                                        if (!("foo bar" <= v14)) {
                                                            var v35 int32 = -2147483648
                                                            if (0.1 <= (0.0 * 3.14159)) {
                                                                v36 := []int{fi17, fi15, (fi15 - fi17), (fi17 << ((514) & 63))}
                                                                var v37 float64 = (0.0 / (3.14159 / 0.1))
                                                                v4 -= v35
                                                                _ = v36
                                                                _ = v37
                                                            }
                                                            _ = v35
                                                        } else {
                                                            if (!true) {
                                                                v38 := []int{fi17, fi15, 0, -954}
                                                                switch fi15 {
                                                                case -9223372036854775808:
                                                                    v38 = append(v38, 9223372036854775807)
                                                                default:
                                                                    v14 = v14
                                                                }
                                                                _ = v38
                                                            }
                                                            v39 := GenS4{F0: v9, F1: (v13 | (v13 & (v13 + 42)))}
                                                            fmt.Println(v1)
                                                            _ = v39
                                                        }
                                                    }
                                                    v40 := GenS2{F0: v13, F1: v11, F2: 206}
                                                    _ = v40
                                                }
                                                v9 += 1
                                                v9 -= v4
                                            }
                                        }
                                        v13 -= (v13 ^ (v13 | v13))
                                    } else {
                                        if (int64(v2) > 1) {
                                            fmt.Println(v9, v1, v9)
                                            if (!(fi15 <= (fi15 % ((fi17) | 1)))) {
                                                if true {
                                                    v4 = (v4 << ((528) & 63))
                                                    var v41 bool = (42 != -361)
                                                    v1 = (v1 - v1)
                                                    _ = v41
                                                }
                                                v42 := []int{}
                                                var v43 uint16 = v3
                                                if false {
                                                    v44 := GenS5{F0: 233, F1: (^(v13)), F2: 0}
                                                    switch v14 {
                                                    case "a":
                                                        fmt.Println(v3)
                                                    }
                                                    var v45 int8 = -1
                                                    _ = v44
                                                    _ = v45
                                                }
                                                var v46 float64 = (((3.14159 - ((0.0 + (((1.5 + (1e10 * 1.5)) - -2.25) + ((3.14159 / float64(3.14159)) - 0.1))) * 0.1)) + (0.0 - 3.14159)) * 1.5)
                                                _ = v42
                                                _ = v43
                                                _ = v46
                                            } else {
                                                v47 := []int{}
                                                fmt.Println(v3, v13)
                                                if len(v47) >= 2 {
                                                    v47 = v47[1:len(v47)]
                                                }
                                                _ = v47
                                            }
                                            if ((v1 + v1) >= v1) {
                                                var v48 uint16 = v3
                                                _ = v48
                                            }
                                        }
                                        switch "foo bar" {
                                        case "\n\t":
                                            switch (v14 + v14) {
                                            case "\n\t":
                                                var v49 uint16 = v3
                                                v50 := GenS3{F0: v9}
                                                v1 = (-(v1))
                                                _ = v49
                                                _ = v50
                                            case "":
                                                var v51 uint64 = 945
                                                _ = v51
                                            case "foo bar":
                                                if (true || (1e10 < ((0.1 + 1.5) / float32(((3.14159 / 1.0) / 1.5))))) {
                                                    v52 := []int{}
                                                    var v53 uint32 = (v2 ^ 0)
                                                    _ = len(v52)
                                                    _ = v52
                                                    _ = v53
                                                }
                                            default:
                                                switch (((v14 + v11) + "foo bar") + (v11 + v14)) {
                                                case "\n\t":
                                                    fmt.Println(v14, v1)
                                                    if false {
                                                        var v54 float32 = 1.5
                                                        v55 := []int{fi17, (^(fi15))}
                                                        v9 += (v4 & v4)
                                                        _ = v54
                                                        _ = v55
                                                    } else {
                                                        v1 = (v1 + (^(v1)))
                                                    }
                                                default:
                                                    switch "a" {
                                                    case "foo bar":
                                                        if true {
                                                            v56 := GenS5{F0: 255, F1: (v13 << ((4294967295) & 63)), F2: uint8(v13)}
                                                            _ = v56
                                                        }
                                                    case "":
                                                        v57 := GenS1{F0: -1, F1: (v2 & 543), F2: v3}
                                                        _ = v57
                                                    default:
                                                        fmt.Println(fi15, v1)
                                                    }
                                                }
                                            }
                                        }
                                    }
                                }
                            }
                            var v58 float32 = 1.5
                            _ = v58
                        }
                        v59 := GenS4{F0: v9, F1: 18446744073709551615}
                        v14 = (v11 + v11)
                        _ = fi17
                        _ = v59
                    }
                }
                _ = fi15
            }
            _ = v13
            _ = v14
        } else {
            for fi60 := 0; fi60 < 5; fi60++ {
                v61 := GenS6{F0: v9}
                v2 += (v2 & (v2 << ((27) & 63)))
                _ = fi60
                _ = v61
            }
            for fi62 := 0; fi62 < 19; fi62++ {
                if (-2147483648 == v9) {
                    v4 += (v9 - (v9 % ((v9) | 1)))
                    var v63 int8 = 72
                    _ = v63
                } else {
                    for fi64 := 0; fi64 < 11; fi64++ {
                        v3 += v3
                        fmt.Println(v2, v2, v2)
                        _ = fi64
                    }
                }
                if true {
                    fmt.Println(v11)
                    if (!false) {
                        var v65 int32 = 42
                        v11 = v11
                        var v66 string = v11
                        _ = v65
                        _ = v66
                    } else {
                        for fi67 := 0; fi67 < 10; fi67++ {
                            var v68 int32 = (-(v4))
                            _ = fi67
                            _ = v68
                        }
                        fmt.Println(v1, fi62, v1)
                    }
                } else {
                    v69 := GenS3{F0: v4}
                    v70 := []int{}
                    _ = v69
                    _ = v70
                }
                fmt.Println(v9, v9, fi62)
                _ = fi62
            }
            for fi71 := 0; fi71 < 16; fi71++ {
                v72 := []int{-9223372036854775808, fi71}
                _ = fi71
                _ = v72
            }
        }
    }
    var v73 uint8 = 212
    if (false && ("foo bar" != v11)) {
        v74 := GenS7{F0: 177, F1: 0}
        _ = v74
    }
    if (4294967295 <= 193) {
        if ((v1 / ((32767) | 1)) == v1) {
            var v75 int8 = 56
            _ = v75
        }
        var v76 int32 = (v9 + (v9 % ((0) | 1)))
        v76 += v9
        _ = v76
    } else {
        fmt.Println(v2, v1, v4)
    }
    v77 := []int{}
    var v78 uint32 = 1
    switch -1 {
    case 680:
        v1 -= v1
    }
    v9 -= v4
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v73:", v73)
    fmt.Println("slicelen v77:", len(v77), v77)
    fmt.Println("v78:", v78)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v73
    _ = v77
    _ = v78
    _ = v9
}
