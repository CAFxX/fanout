package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int64; F1 string; F2 bool }
type GenS4 struct { F0 uint64; F1 int8 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint; F1 int64; F2 int8 }
type GenS7 struct { F0 uint; F1 int32; F2 int }
type GenS8 struct { F0 string; F1 uint16 }
type GenS9 struct { F0 int; F1 uint32; F2 uint32 }
func genFn0(p0 int16, p1 int64, p2 bool) int16 {
    for fi6 := 0; fi6 < 16; fi6++ {
        if (!(p1 > p1)) {
            if ((p1 * (p1 << ((42) & 63))) <= (p1 - (-(p1)))) {
                var v7 int16 = 576
                if true {
                    if ((!p2) || (!p2)) {
                        v8 := GenS1{F0: 2}
                        p0 = -55
                        _ = v8
                    } else {
                        var v9 int64 = p1
                        p0 += (p0 - p0)
                        switch "\n\t" {
                        case "\n\t":
                            var v10 string = (("\n\t" + "") + "\n\t")
                            _ = v10
                        case "foo bar":
                            v11 := GenS1{F0: 255}
                            _ = v11
                        case "":
                            if (((p1 / ((p1) | 1)) <= (^(v9))) || p2) {
                                p2 = p2
                                for fi12 := 0; fi12 < 19; fi12++ {
                                    var v13 uint32 = 0
                                    switch (("hello" + "") + ("foo bar" + ("\n\t" + "a"))) {
                                    case "a":
                                        var v14 string = ""
                                        var v15 float32 = (((3.14159 - 0.0) / (-2.25 / float32(3.14159))) * (0.0 * (1.5 / -2.25)))
                                        _ = v14
                                        _ = v15
                                    case "\n\t":
                                        var v16 uint64 = 42
                                        _ = v16
                                    case "foo bar":
                                        fmt.Println(p2)
                                    }
                                    v17 := GenS2{F0: (65535 < 300)}
                                    _ = fi12
                                    _ = v13
                                    _ = v17
                                }
                            }
                        default:
                            for fi18 := 0; fi18 < 7; fi18++ {
                                fmt.Println(p0, p2, p0)
                                _ = fi18
                            }
                        }
                        _ = v9
                    }
                    var v19 int64 = p1
                    _ = v19
                } else {
                    v20 := []int{fi6}
                    for fi21 := 0; fi21 < 10; fi21++ {
                        var v22 string = "\n\t"
                        _ = fi21
                        _ = v22
                    }
                    for fi23 := 0; fi23 < 13; fi23++ {
                        if p2 {
                            switch (fi23 ^ (^(fi23))) {
                            case 9223372036854775807:
                                p1 = p1
                            }
                            fmt.Println(fi23, fi23)
                        }
                        _ = fi23
                    }
                    _ = v20
                }
                var v24 uint16 = 255
                _ = v24
                _ = v7
            }
        } else {
            fmt.Println(p2, p0)
            v25 := GenS3{F0: p1, F1: "\n\t", F2: (0 >= 32767)}
            v26 := []int{fi6}
            _ = v25
            _ = v26
        }
        p2 = (!p2)
        var v27 int32 = int32(p1)
        _ = fi6
        _ = v27
    }
    switch "" {
    case "foo bar":
        for fi28 := 0; fi28 < 15; fi28++ {
            p0 = (p0 >> ((42) & 63))
            v29 := GenS2{F0: false}
            var v30 float64 = float64(0.1)
            _ = fi28
            _ = v29
            _ = v30
        }
    case "\n\t":
        fmt.Println(p1)
    default:
        var v31 bool = ((0 == 795) || ((!p2) || (!(2 > 174))))
        _ = v31
    }
    p1 -= p1
    var v32 uint = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v32
    return p0
}
func genFn1(p0 int, p1 int64, p2 int16) uint16 {
    var v33 int32 = -662
    v34 := GenS2{F0: true}
    for fi35 := 0; fi35 < 12; fi35++ {
        for fi36 := 0; fi36 < 6; fi36++ {
            v37 := []int{fi35}
            v33 += -280
            _ = fi36
            _ = v37
        }
        fmt.Println(v33, p1, p1)
        _ = fi35
    }
    for fi38 := 0; fi38 < 15; fi38++ {
        for fi39 := 0; fi39 < 19; fi39++ {
            p2 = (p2 ^ p2)
            _ = fi39
        }
        for fi40 := 0; fi40 < 18; fi40++ {
            if false {
                p1 += -954
                v33 += v33
            } else {
                switch -895 {
                case 778:
                    if false {
                        v41 := []int{}
                        v42 := GenS4{F0: 2, F1: 1}
                        if len(v41) > 0 {
                            v41[0] = fi40
                        }
                        _ = v41
                        _ = v42
                    }
                case -270:
                    v33 += 470
                case 32768:
                    var v43 uint8 = 42
                    _ = v43
                }
            }
            v44 := GenS2{F0: true}
            if true {
                v45 := GenS5{F0: (p1 % (((p1 * -9223372036854775808)) | 1))}
                var v46 int32 = 1
                var v47 bool = ((((!false) || false) && false) || true)
                _ = v45
                _ = v46
                _ = v47
            }
            _ = fi40
            _ = v44
        }
        var v48 int = fi38
        _ = fi38
        _ = v48
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v33
    _ = v34
    return 65535
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int = 9223372036854775807
    var v3 int16 = 32767
    var v4 int64 = 1
    var v5 uint64 = 42
    v4 -= v4
    var v49 string = ""
    fmt.Println(v5)
    var v50 int16 = v3
    for fi51 := 0; fi51 < 20; fi51++ {
        switch (v49 + (v49 + v49)) {
        case "\n\t":
            v52 := GenS3{F0: v4, F1: v49, F2: (genFn1(-9223372036854775808, (v4 % ((42) | 1)), (v3 - v50)) >= 65535)}
            _ = v52
        case "hello":
            if ((!(255 > 42)) && true) {
                v50 -= genFn0(int16(v4), (v4 ^ v4), false)
            }
        case "a":
            v4 += v4
        default:
            v53 := GenS5{F0: 239}
            _ = v53
        }
        var v54 uint64 = 0
        _ = fi51
        _ = v54
    }
    v2 += v2
    switch (-(v2)) {
    case 32768:
        var v55 string = "foo bar"
        _ = v55
    }
    v4 += v4
    for fi56 := 0; fi56 < 20; fi56++ {
        var v57 uint = 1
        switch v49 {
        case "foo bar":
            switch (v2 * 32768) {
            case -9223372036854775808:
                v1[4294967295] = "hello"
            case 9223372036854775807:
                delete(v1, 2)
            }
        case "":
            for fi58 := 0; fi58 < 7; fi58++ {
                if true {
                    v57 = (v57 + v57)
                } else {
                    if ((0 > 32767) || (2147483647 > 616)) {
                        fmt.Println(v2, v4)
                        var v59 uint32 = 65535
                        v3 = genFn0((-(v3)), v4, ((v57 >> (((v57 / ((v57) | 1))) & 63)) != uint(v59)))
                        v60 := GenS4{F0: 1, F1: int8(v3)}
                        var v61 float64 = 1e10
                        _ = v59
                        _ = v60
                        _ = v61
                    } else {
                        fmt.Println(v3, fi56)
                        if false {
                            v4 -= (-(v4))
                        }
                    }
                }
                if true {
                    if true {
                        v62 := v1[2147483648]
                        switch v49 {
                        case "":
                            v63 := GenS6{F0: 65535, F1: v4, F2: 100}
                            _ = v63
                        case "\n\t":
                            v64 := []int{4294967295, -9223372036854775808}
                            _ = v64
                        case "hello":
                            v65 := GenS7{F0: (^(v57)), F1: -1, F2: int(v5)}
                            _ = v65
                        default:
                            v3 -= 604
                        }
                        _ = v62
                    } else {
                        v66 := GenS3{F0: (v4 % ((v4) | 1)), F1: v49, F2: (32768 <= -2147483648)}
                        var v67 int64 = v4
                        _ = v66
                        _ = v67
                    }
                    var v68 int8 = 100
                    _ = v68
                } else {
                    _ = len(v1)
                    fmt.Println(fi58, v2, v3)
                }
                _ = fi58
            }
        default:
            var v69 uint16 = 27
            _ = v69
        }
        for fi70 := 0; fi70 < 1; fi70++ {
            v50 = genFn0((v50 >> ((255) & 63)), (-(v4)), ((true && (1e10 > -2.25)) || ((42 > 908) && (!false))))
            var v71 int32 = 2147483647
            _ = fi70
            _ = v71
        }
        _ = fi56
        _ = v57
    }
    v72 := GenS8{F0: "hello", F1: 32767}
    if (189 != 255) {
        fmt.Println(v3)
    }
    switch (v2 & v2) {
    case 4294967295:
        var v73 float64 = (1.5 / 1e10)
        _ = v73
    case -142:
        fmt.Println(v4, v50)
    case 42:
        v74 := GenS9{F0: v2, F1: 2147483648, F2: 42}
        _ = v74
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v50
    _ = v72
}
