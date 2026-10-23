package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string }
type GenS2 struct { F0 int; F1 uint64; F2 uint32 }
type GenS3 struct { F0 uint64; F1 uint8; F2 int }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint16; F1 int64 }
type GenS6 struct { F0 uint; F1 int8; F2 uint8 }
type GenS7 struct { F0 int8; F1 uint32; F2 uint16 }
type GenS8 struct { F0 bool; F1 uint64; F2 int64 }
func genFn0(p0 bool, p1 uint, p2 uint32) uint16 {
    for fi7 := 0; fi7 < 6; fi7++ {
        if (-464 <= 146) {
            switch fi7 {
            case 2147483647:
                v8 := GenS1{F0: fi7, F1: "hello"}
                _ = v8
            case 128:
                p2 = (p2 >> ((p1) & 63))
            case 32767:
                v9 := GenS2{F0: fi7, F1: 1, F2: (p2 << ((p1) & 63))}
                _ = v9
            }
        } else {
            fmt.Println(fi7)
        }
        p2 = p2
        _ = fi7
    }
    v10 := GenS3{F0: 3, F1: 255, F2: 128}
    v11 := GenS1{F0: -153, F1: "hello"}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    return 42
}
func genFn1(p0 int32, p1 int32) uint8 {
    var v12 uint8 = 134
    for fi13 := 0; fi13 < 3; fi13++ {
        switch "a" {
        case "\n\t":
            var v14 float64 = -2.25
            _ = v14
        case "hello":
            for fi15 := 0; fi15 < 6; fi15++ {
                var v16 uint16 = 769
                if false {
                    p1 += p0
                } else {
                    var v17 uint = 1
                    v17 += 65535
                    if false {
                        var v18 uint16 = v16
                        v19 := GenS1{F0: (fi15 % (((fi15 + 128)) | 1)), F1: "hello"}
                        if false {
                            v20 := []int{fi13, 0, 42, -2147483648}
                            var v21 uint8 = 1
                            _ = v20
                            _ = v21
                        } else {
                            fmt.Println(v17, v12)
                        }
                        _ = v18
                        _ = v19
                    } else {
                        switch "foo bar" {
                        case "a":
                            var v22 int32 = p0
                            _ = v22
                        case "":
                            var v23 uint = (v17 >> ((v17) & 63))
                            _ = v23
                        case "foo bar":
                            if (42 != 4294967295) {
                                var v24 float64 = -2.25
                                switch (("hello" + "") + "foo bar") {
                                case "a":
                                    v25 := []int{358, (fi13 + (fi13 / ((-1) | 1))), fi15, int(v16)}
                                    _ = v25
                                case "\n\t":
                                    v26 := GenS3{F0: 2, F1: v12, F2: fi13}
                                    _ = v26
                                }
                                _ = v24
                            }
                        default:
                            fmt.Println(v17, v16, p1)
                        }
                        fmt.Println(p0, fi13, p1)
                    }
                    if true {
                        var v27 int64 = 1
                        switch int(p0) {
                        case -2147483648:
                            v28 := []int{fi13}
                            _ = v28
                        case -9223372036854775808:
                            v29 := []int{(fi15 << ((4294967295) & 63)), (-(fi13))}
                            _ = v29
                        case 42:
                            v30 := []int{}
                            _ = v30
                        }
                        _ = v27
                    } else {
                        var v31 string = (("hello" + "") + ("" + "hello"))
                        var v32 float32 = (1.5 / float32(1e10))
                        p1 -= (p0 | (p0 * (p0 | (p0 ^ (p0 / (((p0 >> ((v17) & 63))) | 1))))))
                        _ = v31
                        _ = v32
                    }
                    _ = v17
                }
                p0 += p1
                _ = fi15
                _ = v16
            }
        default:
            for fi33 := 0; fi33 < 19; fi33++ {
                fmt.Println(v12, v12)
                _ = fi33
            }
        }
        _ = fi13
    }
    v34 := []int{-2147483648, -1}
    for fi35 := 0; fi35 < 12; fi35++ {
        if (!(!false)) {
            v12 = (v12 << ((250) & 63))
        } else {
            v36 := GenS2{F0: (fi35 % (((fi35 * fi35)) | 1)), F1: 2, F2: 2}
            v34 = append(v34, (fi35 % (((fi35 - 1)) | 1)))
            for fi37 := 0; fi37 < 2; fi37++ {
                _ = len(v34)
                if true {
                    p1 += p1
                } else {
                    var v38 int8 = 100
                    var v39 int64 = 100
                    _ = v38
                    _ = v39
                }
                switch "foo bar" {
                case "a":
                    if len(v34) > 0 {
                        v34[0] = (fi37 * fi35)
                    }
                }
                _ = fi37
            }
            _ = v36
        }
        p1 += 2147483647
        _ = fi35
    }
    _ = p0
    _ = p1
    _ = v12
    _ = v34
    return 2
}

func main() {
    v1 := []string{(("a" + "a") + "")}
    var v2 uint = 288
    var v3 int = 32767
    v4 := []string{}
    v5 := []int{(v3 | (v3 + (^(v3)))), v3, (v3 - v3)}
    v6 := []int{v3, v3}
    var v40 int32 = 1
    if (!((v40 / ((v40) | 1)) == 0)) {
        switch (v3 * 32768) {
        case -9223372036854775808:
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
        }
        v41 := GenS4{F0: v2}
        v40 = v40
        _ = v41
    }
    var v42 uint32 = 2147483648
    if (((float32(-2.25) / (1e10 + float32(-2.25))) - float32(0.1)) != 1.5) {
        v2 -= (^(v2))
        for fi43 := 0; fi43 < 10; fi43++ {
            fmt.Println(fi43)
            _ = fi43
        }
        v44 := GenS5{F0: genFn0(false, v2, (v42 * (v42 / ((42) | 1)))), F1: 1}
        _ = v44
    } else {
        if true {
            var v45 int16 = -32768
            if false {
                var v46 int8 = 100
                v40 = (v40 ^ -577)
                _ = len(v1)
                _ = v46
            }
            _ = v45
        }
        var v47 int64 = 322
        v3 -= v3
        _ = v47
    }
    var v48 int64 = -9223372036854775808
    v42 = v42
    var v49 float64 = 1.5
    v50 := GenS2{F0: v3, F1: 144, F2: (v42 * (v42 - (v42 + (v42 & (v42 % ((42) | 1))))))}
    v48 -= (v48 % ((v48) | 1))
    fmt.Println(v3, v2, v40)
    v2 -= (v2 & v2)
    for fi51 := 0; fi51 < 19; fi51++ {
        switch "\n\t" {
        case "\n\t":
            if (!(false || true)) {
                var v52 float32 = float32(((0.1 - float64(0.1)) + (v49 * (v49 + (v49 * float64(0.1))))))
                v52 = v52
                switch "\n\t" {
                case "a":
                    var v53 uint16 = 65535
                    _ = v53
                case "\n\t":
                    if (65535 == 255) {
                        v54 := GenS2{F0: (fi51 % ((v3) | 1)), F1: 4294967295, F2: 65535}
                        if false {
                            fmt.Println(v42)
                            if (!false) {
                                var v55 bool = ("\n\t" >= "")
                                _ = v55
                            } else {
                                for fi56 := 0; fi56 < 16; fi56++ {
                                    var v57 uint32 = 114
                                    switch ("\n\t" + "foo bar") {
                                    case "":
                                        if false {
                                            v58 := GenS3{F0: 4294967295, F1: 219, F2: (fi51 - 85)}
                                            v59 := GenS6{F0: (^(v2)), F1: 85, F2: uint8(v2)}
                                            _ = v58
                                            _ = v59
                                        } else {
                                            var v60 uint = 336
                                            v57 += (v57 + v57)
                                            _ = v60
                                        }
                                    case "a":
                                        switch ("" + "foo bar") {
                                        case "":
                                            var v61 int = (fi51 - v3)
                                            _ = v61
                                        case "hello":
                                            var v62 string = (("" + "\n\t") + ("foo bar" + "\n\t"))
                                            _ = v62
                                        case "a":
                                            var v63 bool = true
                                            _ = v63
                                        }
                                    case "foo bar":
                                        v64 := GenS7{F0: 100, F1: v42, F2: 2}
                                        var v65 int64 = 1
                                        if (true || false) {
                                            if false {
                                                var v66 uint8 = 119
                                                fmt.Println(v2, v57)
                                                fmt.Println(v48)
                                                _ = v66
                                            } else {
                                                fmt.Println(fi51)
                                                if true {
                                                    v67 := GenS3{F0: 2, F1: 127, F2: (v3 + (fi56 * v3))}
                                                    v42 += v57
                                                    var v68 float64 = v49
                                                    _ = v67
                                                    _ = v68
                                                } else {
                                                    fmt.Println(fi51, v65, v3)
                                                    var v69 int16 = 349
                                                    var v70 float32 = v52
                                                    _ = v69
                                                    _ = v70
                                                }
                                            }
                                        }
                                        _ = v64
                                        _ = v65
                                    default:
                                        fmt.Println(v48, v3, v2)
                                    }
                                    v71 := GenS2{F0: fi56, F1: 2, F2: v42}
                                    _ = fi56
                                    _ = v57
                                    _ = v71
                                }
                                var v72 uint64 = 985
                                var v73 int64 = v48
                                _ = v72
                                _ = v73
                            }
                            v74 := GenS8{F0: (!((0 <= 54) && true)), F1: 18446744073709551615, F2: 0}
                            _ = v74
                        }
                        _ = v54
                    }
                }
                _ = v52
            } else {
                var v75 uint16 = 42
                v2 = v2
                _ = v75
            }
        case "hello":
            for fi76 := 0; fi76 < 9; fi76++ {
                if (-32768 != 623) {
                    v77 := GenS6{F0: (v2 ^ v2), F1: 0, F2: 128}
                    if (0 >= 250) {
                        if false {
                            v78 := GenS2{F0: int(v2), F1: 42, F2: v42}
                            v3 = fi76
                            _ = v78
                        } else {
                            var v79 int16 = 0
                            var v80 string = ("foo bar" + "\n\t")
                            switch (fi51 >> ((v2) & 63)) {
                            case 318:
                                if (201 < 255) {
                                    v79 = (v79 / (((v79 & (v79 >> ((v2) & 63)))) | 1))
                                }
                            default:
                                v4 = append(v4, v80)
                            }
                            _ = v79
                            _ = v80
                        }
                    }
                    _ = v77
                }
                _ = fi76
            }
        }
        _ = fi51
    }
    v40 = int32(v48)
    v42 = (v42 % (((v42 / ((42) | 1))) | 1))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("v42:", v42)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v42
    _ = v48
    _ = v49
    _ = v5
    _ = v50
    _ = v6
}
