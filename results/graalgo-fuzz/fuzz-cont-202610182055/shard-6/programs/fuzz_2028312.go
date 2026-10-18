package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 int32 }
type GenS2 struct { F0 uint64; F1 uint; F2 string }
type GenS3 struct { F0 int32; F1 uint32 }
type GenS4 struct { F0 uint32; F1 int16 }
type GenS5 struct { F0 uint32; F1 int32; F2 string }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 uint; F1 uint8; F2 uint64 }
type GenS8 struct { F0 string }
func genFn0(p0 int16, p1 uint8) string {
    var v4 int = -104
    p0 += (p0 & (p0 - p0))
    var v5 uint16 = 32767
    p0 = (^(p0))
    for fi6 := 0; fi6 < 13; fi6++ {
        switch (-(v4)) {
        case 32767:
            if ((p0 + (p0 << ((0) & 63))) != p0) {
                switch (v4 / (((fi6 & fi6)) | 1)) {
                case -9223372036854775808:
                    switch "a" {
                    case "foo bar":
                        var v7 int8 = 100
                        _ = v7
                    case "":
                        fmt.Println(v4, p1, v5)
                    case "\n\t":
                        if false {
                            var v8 uint32 = 1
                            v9 := GenS1{F0: p1, F1: 4294967295, F2: 1}
                            _ = v8
                            _ = v9
                        }
                    }
                case 32767:
                    var v10 uint8 = 221
                    _ = v10
                case 4294967295:
                    switch 527 {
                    case 42:
                        v11 := GenS2{F0: 4294967295, F1: 771, F2: "\n\t"}
                        _ = v11
                    default:
                        v5 = v5
                    }
                default:
                    v4 = (fi6 << ((1) & 63))
                }
                v12 := []int{}
                _ = v12
            } else {
                switch -1 {
                case 0:
                    var v13 uint = 18446744073709551615
                    _ = v13
                case 127:
                    var v14 uint16 = v5
                    _ = v14
                default:
                    fmt.Println(p1, v5)
                }
                v5 = 65535
            }
        case 0:
            fmt.Println(v5, v5, p1)
        case -339:
            for fi15 := 0; fi15 < 13; fi15++ {
                if false {
                    if false {
                        v16 := GenS3{F0: 2147483647, F1: 808}
                        if (!(p1 > 42)) {
                            var v17 int8 = 100
                            _ = v17
                        }
                        _ = v16
                    }
                    fmt.Println(v5, v4)
                    v18 := []int{fi6, (v4 % ((v4) | 1))}
                    _ = v18
                } else {
                    switch 2147483647 {
                    case 128:
                        var v19 uint64 = 798
                        _ = v19
                    case -1:
                        if (-381 > -1) {
                            v20 := GenS3{F0: -384, F1: 1}
                            _ = v20
                        } else {
                            v21 := []int{fi6, int(p0), (-(v4))}
                            v22 := GenS1{F0: p1, F1: 551, F2: 566}
                            var v23 int = v4
                            _ = v21
                            _ = v22
                            _ = v23
                        }
                    case 2147483647:
                        v24 := GenS4{F0: 954, F1: (p0 >> ((4294967295) & 63))}
                        _ = v24
                    }
                    var v25 bool = (true && false)
                    switch (fi6 % ((4294967295) | 1)) {
                    case 128:
                        v26 := GenS5{F0: 1, F1: 1, F2: ("a" + "a")}
                        _ = v26
                    default:
                        var v27 uint16 = (v5 + (^(v5)))
                        _ = v27
                    }
                    _ = v25
                }
                if (42 >= 255) {
                    v28 := []int{fi6, (fi15 >> ((786) & 63)), fi15, (fi15 % (((v4 * (-(v4)))) | 1))}
                    fmt.Println(v5, p1, v4)
                    switch "\n\t" {
                    case "\n\t":
                        switch v4 {
                        case -2147483648:
                            fmt.Println(p1, p0)
                        default:
                            var v29 int64 = 1
                            switch (v4 * -282) {
                            case 128:
                                var v30 uint8 = (p1 / (((p1 / ((127) | 1))) | 1))
                                _ = v30
                            case -2147483648:
                                p1 -= 128
                            }
                            var v31 uint8 = p1
                            _ = v29
                            _ = v31
                        }
                    case "a":
                        var v32 uint16 = (v5 * 1)
                        _ = v32
                    case "hello":
                        v28 = append(v28, fi15)
                    }
                    _ = v28
                } else {
                    v33 := []int{(fi15 / (((v4 + -1)) | 1)), v4, (v4 - fi6), (fi6 ^ (v4 % ((fi6) | 1)))}
                    var v34 int32 = -1
                    var v35 float64 = (1.5 - 3.14159)
                    var v36 uint16 = v5
                    if len(v33) >= 2 {
                        v33 = v33[1:len(v33)]
                    }
                    var v37 int64 = int64(p1)
                    _ = v33
                    _ = v34
                    _ = v35
                    _ = v36
                    _ = v37
                }
                var v38 int64 = 1
                _ = fi15
                _ = v38
            }
        }
        var v39 float32 = ((((0.1 - 3.14159) - (1e10 - (3.14159 + (1e10 / float32(1.5))))) / ((0.1 * -2.25) + float32(1e10))) * float32(-2.25))
        v40 := GenS1{F0: 127, F1: 4294967295, F2: -491}
        _ = fi6
        _ = v39
        _ = v40
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return ("a" + "a")
}

func main() {
    v1 := []string{("" + ""), "hello", "foo bar"}
    var v2 int8 = 100
    var v3 uint = 42
    fmt.Println(v2)
    fmt.Println(v3, v2)
    for fi41 := 0; fi41 < 20; fi41++ {
        for fi42 := 0; fi42 < 13; fi42++ {
            switch ("foo bar" + "") {
            case "\n\t":
                var v43 uint32 = 156
                _ = v43
            case "":
                var v44 uint64 = 42
                _ = v44
            default:
                switch fi41 {
                case 4294967295:
                    if len(v1) > 0 {
                        v1[0] = ""
                    }
                    var v45 int64 = int64(fi41)
                    _ = v45
                case 42:
                    if (!((4294967295 > 2) && true)) {
                        var v46 uint8 = 128
                        _ = v46
                    } else {
                        fmt.Println(fi42, fi41, v2)
                    }
                case -2147483648:
                    v2 -= int8(v3)
                }
            }
            _ = fi42
        }
        _ = fi41
    }
    var v47 uint8 = 2
    v48 := GenS6{F0: v3}
    fmt.Println(v3, v3)
    v49 := GenS5{F0: 2, F1: 0, F2: "a"}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v50 int16 = 42
    if false {
        for fi51 := 0; fi51 < 2; fi51++ {
            v52 := GenS2{F0: uint64(v3), F1: (v3 + v3), F2: ("a" + "hello")}
            v50 += (v50 - (^(v50)))
            if (((((0.1 * float32(0.0)) + 0.1) / ((float32(0.1) * 1.5) / 1e10)) * float32((3.14159 + (0.1 / (0.1 / (0.1 - -2.25)))))) > (1.5 - (1e10 - (float32(3.14159) - 3.14159)))) {
                for fi53 := 0; fi53 < 9; fi53++ {
                    var v54 int64 = -813
                    v55 := GenS7{F0: (v3 << ((452) & 63)), F1: 184, F2: 0}
                    _ = fi53
                    _ = v54
                    _ = v55
                }
            } else {
                var v56 uint32 = 2147483648
                v56 = v56
                for fi57 := 0; fi57 < 13; fi57++ {
                    if true {
                        v58 := GenS7{F0: 0, F1: v47, F2: 892}
                        if (0 < 42) {
                            var v59 uint16 = 42
                            if false {
                                if (!(v56 >= 2)) {
                                    fmt.Println(v56, v2, v2)
                                }
                            } else {
                                switch (-(fi51)) {
                                case -434:
                                    v59 += v59
                                case 4294967295:
                                    var v60 float32 = 0.0
                                    _ = v60
                                case 0:
                                    var v61 float32 = 0.1
                                    _ = v61
                                }
                                var v62 uint16 = 177
                                _ = len(v1)
                                _ = v62
                            }
                            _ = v59
                        }
                        _ = v58
                    } else {
                        if true {
                            fmt.Println(v47)
                            fmt.Println(fi57, fi57, fi51)
                            v2 += (v2 ^ -1)
                        } else {
                            v63 := GenS8{F0: "a"}
                            _ = v63
                        }
                        if len(v1) > 0 {
                            v1[0] = "\n\t"
                        }
                        v3 += v3
                    }
                    switch (("" + ("hello" + "hello")) + ("\n\t" + (("" + "") + ""))) {
                    case "a":
                        var v64 int8 = v2
                        _ = v64
                    }
                    _ = fi57
                }
                _ = v56
            }
            _ = fi51
            _ = v52
        }
    }
    v2 += 127
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v47:", v47)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v3
    _ = v47
    _ = v48
    _ = v49
    _ = v50
}
