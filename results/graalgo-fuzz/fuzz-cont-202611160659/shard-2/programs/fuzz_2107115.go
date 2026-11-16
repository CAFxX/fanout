package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint32; F1 int16; F2 int16 }
type GenS3 struct { F0 bool; F1 int64 }
type GenS4 struct { F0 int64; F1 uint16 }
type GenS5 struct { F0 uint8; F1 int32; F2 string }
type GenS6 struct { F0 int8; F1 int8; F2 uint8 }
type GenS7 struct { F0 uint16 }
func genFn0(p0 uint) uint64 {
    for fi6 := 0; fi6 < 12; fi6++ {
        v7 := GenS1{F0: 789}
        switch "\n\t" {
        case "foo bar":
            p0 = 52
        }
        switch (^(fi6)) {
        case -333:
            fmt.Println(fi6)
        case 9223372036854775807:
            v8 := []int{(fi6 >> ((p0) & 63))}
            _ = v8
        }
        _ = fi6
        _ = v7
    }
    for fi9 := 0; fi9 < 13; fi9++ {
        v10 := GenS1{F0: 42}
        v11 := []int{(fi9 << ((p0) & 63)), fi9, fi9}
        v12 := GenS2{F0: 389, F1: -1, F2: 742}
        _ = fi9
        _ = v10
        _ = v11
        _ = v12
    }
    _ = p0
    return 42
}
func genFn1(p0 int8, p1 uint8) uint16 {
    if false {
        v13 := []int{1, -487, -1}
        v14 := GenS1{F0: 884}
        _ = v13
        _ = v14
    } else {
        switch 0 {
        case 1:
            for fi15 := 0; fi15 < 9; fi15++ {
                for fi16 := 0; fi16 < 11; fi16++ {
                    v17 := []int{fi16, (fi15 << ((4294967295) & 63)), (fi15 >> ((255) & 63)), fi15}
                    p1 = (p1 & (^(p1)))
                    _ = fi16
                    _ = v17
                }
                v18 := []int{fi15, fi15, 1}
                if ((1 != 42) && true) {
                    if false {
                        p0 = (p0 & (p0 % (((p0 + p0)) | 1)))
                        for fi19 := 0; fi19 < 13; fi19++ {
                            p0 = 42
                            if (0 <= 2) {
                                p0 += 127
                                switch "a" {
                                case "foo bar":
                                    var v20 uint = 4294967295
                                    _ = v20
                                case "\n\t":
                                    p0 += int8(fi15)
                                }
                            } else {
                                var v21 int32 = 803
                                if ((^(p1)) > 164) {
                                    p1 = p1
                                }
                                switch "hello" {
                                case "hello":
                                    if false {
                                        v22 := GenS3{F0: true, F1: 0}
                                        var v23 uint64 = 1
                                        _ = v22
                                        _ = v23
                                    }
                                case "\n\t":
                                    v24 := GenS3{F0: false, F1: 0}
                                    _ = v24
                                case "a":
                                    var v25 string = "foo bar"
                                    _ = v25
                                }
                                _ = v21
                            }
                            if true {
                                p1 = p1
                                _ = len(v18)
                            }
                            _ = fi19
                        }
                        p0 -= 47
                    }
                    p1 += 1
                } else {
                    if (32767 >= -121) {
                        var v26 int = -2147483648
                        _ = v26
                    }
                    if (!(!false)) {
                        v27 := GenS1{F0: 0}
                        _ = v27
                    }
                    switch "\n\t" {
                    case "a":
                        if (2147483647 < 2) {
                            v28 := GenS3{F0: (!true), F1: 0}
                            _ = v28
                        }
                    case "foo bar":
                        for fi29 := 0; fi29 < 18; fi29++ {
                            if (1 != 194) {
                                v30 := GenS2{F0: 1, F1: 1000, F2: 435}
                                var v31 uint64 = 4294967295
                                var v32 uint64 = (v31 * v31)
                                _ = v30
                                _ = v31
                                _ = v32
                            } else {
                                p0 += (p0 % (((p0 + p0)) | 1))
                                if (false && true) {
                                    v18 = append(v18, (fi29 ^ 504))
                                } else {
                                    fmt.Println(fi15, fi15, fi29)
                                    p1 -= 94
                                    p0 = p0
                                }
                            }
                            _ = fi29
                        }
                    default:
                        var v33 int = 546
                        _ = v33
                    }
                }
                _ = fi15
                _ = v18
            }
        case -450:
            p1 = (p1 | (p1 | 1))
        default:
            var v34 uint32 = uint32(p1)
            _ = v34
        }
        v35 := GenS1{F0: 1}
        fmt.Println(p0, p1)
        _ = v35
    }
    var v36 uint16 = 336
    fmt.Println(p0, p1)
    p0 = p0
    _ = p0
    _ = p1
    _ = v36
    return 255
}

func main() {
    var v1 float32 = 0.1
    var v2 string = ((("a" + "\n\t") + "hello") + "a")
    var v3 int32 = 1
    v4 := make(map[uint32]uint32)
    v5 := make(map[string]string)
    var v37 int = 32768
    if (1 < 558) {
        switch (v37 & v37) {
        case 4294967295:
            fmt.Println(v1)
        case -1:
            for fi38 := 0; fi38 < 19; fi38++ {
                for fi39 := 0; fi39 < 11; fi39++ {
                    fmt.Println(v3, v1, fi38)
                    _ = fi39
                }
                fmt.Println(v1)
                _ = fi38
            }
        }
        switch v37 {
        case 127:
            var v40 uint16 = 255
            _ = v40
        case 42:
            var v41 uint8 = 127
            _ = v41
        default:
            v3 += (^(v3))
        }
    } else {
        v1 = v1
    }
    var v42 int = v37
    if false {
        for fi43 := 0; fi43 < 18; fi43++ {
            for fi44 := 0; fi44 < 13; fi44++ {
                v45 := []int{v42}
                var v46 float32 = (1e10 * 1.5)
                v47 := v5[("foo bar" + "foo bar")]
                var v48 int64 = 1
                v49 := GenS4{F0: 161, F1: 42}
                _ = fi44
                _ = v45
                _ = v46
                _ = v47
                _ = v48
                _ = v49
            }
            v50 := GenS5{F0: 2, F1: -272, F2: "hello"}
            _ = fi43
            _ = v50
        }
    }
    switch 473 {
    case 900:
        fmt.Println(v42, v42, v2)
    }
    var v51 uint64 = 2
    v52 := GenS6{F0: 0, F1: 1, F2: 128}
    switch v42 {
    case 2147483647:
        switch (v37 % ((v37) | 1)) {
        case -22:
            if false {
                v53 := GenS6{F0: -88, F1: 22, F2: 0}
                fmt.Println(v3, v1)
                v4[42] = 4294967295
                _ = v53
            } else {
                if false {
                    fmt.Println(v1, v42, v2)
                } else {
                    delete(v5, "hello")
                }
                v37 += v37
                var v54 bool = ((!(135 > 127)) && true)
                _ = v54
            }
        case 32768:
            if true {
                var v55 bool = false
                for fi56 := 0; fi56 < 13; fi56++ {
                    for fi57 := 0; fi57 < 4; fi57++ {
                        fmt.Println(v1, v1, v37)
                        var v58 float32 = v1
                        _ = fi57
                        _ = v58
                    }
                    var v59 uint16 = 35
                    _ = fi56
                    _ = v59
                }
                var v60 uint8 = 229
                _ = v55
                _ = v60
            }
        }
    }
    delete(v4, 4294967295)
    v61 := []int{(v42 - (^(v42)))}
    if (float32(1.5) != float32((-2.25 + 1e10))) {
        v62 := GenS7{F0: 65535}
        switch v37 {
        case 2147483647:
            var v63 int32 = v3
            _ = v63
        case 32768:
            v3 -= (-(v3))
        }
        _ = v62
    }
    v3 -= v3
    fmt.Println(v3)
    v42 += 800
    v3 += (-(v3))
    if true {
        v37 += (-(v42))
        v61 = append(v61, 2147483647)
    }
    v37 -= -888
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v51:", v51)
    fmt.Println("slicelen v61:", len(v61), v61)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v42
    _ = v5
    _ = v51
    _ = v52
    _ = v61
}
