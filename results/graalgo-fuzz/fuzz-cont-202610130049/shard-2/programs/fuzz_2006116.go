package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 string; F1 uint }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint8; F1 int8; F2 bool }
type GenS5 struct { F0 uint64; F1 uint64; F2 int32 }
type GenS6 struct { F0 uint16; F1 int8 }
type GenS7 struct { F0 int16; F1 uint8; F2 uint16 }
type GenS8 struct { F0 uint16; F1 uint16; F2 bool }
func genFn0(p0 int8, p1 uint8, p2 int16) uint {
    p1 += (p1 / (((p1 / ((p1) | 1))) | 1))
    for fi5 := 0; fi5 < 13; fi5++ {
        p1 -= (p1 | (p1 - p1))
        if (p1 != 9) {
            var v6 float32 = 3.14159
            _ = v6
        } else {
            p2 = p2
            v7 := []int{89, (^(fi5)), -269, 1}
            fmt.Println(p1)
            _ = v7
        }
        var v8 float64 = -2.25
        _ = fi5
        _ = v8
    }
    p2 += p2
    _ = p0
    _ = p1
    _ = p2
    return 0
}
func genFn1(p0 uint32) int32 {
    var v9 int16 = 32767
    var v10 string = ""
    for fi11 := 0; fi11 < 15; fi11++ {
        var v12 uint16 = 1
        v13 := []int{}
        _ = fi11
        _ = v12
        _ = v13
    }
    _ = p0
    _ = v10
    _ = v9
    return 182
}

func main() {
    v1 := []string{((("" + ("\n\t" + "")) + ("\n\t" + "\n\t")) + ("\n\t" + "foo bar"))}
    v2 := make(map[int]bool)
    var v3 int64 = -9223372036854775808
    var v4 int = 4294967295
    v3 -= 158
    for fi14 := 0; fi14 < 4; fi14++ {
        var v15 uint32 = 4294967295
        if false {
            for fi16 := 0; fi16 < 13; fi16++ {
                if (int16(v15) >= -484) {
                    if false {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        v17 := v2[(v4 ^ fi16)]
                        fmt.Println(v15, v17, v15)
                        var v18 int16 = -994
                        _ = v17
                        _ = v18
                    }
                } else {
                    switch ("a" + "") {
                    case "\n\t":
                        if (-2.25 > 0.0) {
                            v19 := v2[fi16]
                            _ = v19
                        } else {
                            switch "hello" {
                            case "\n\t":
                                v3 = (v3 << ((0) & 63))
                            }
                            var v20 uint16 = 32
                            _ = v20
                        }
                    case "foo bar":
                        fmt.Println(v15, fi14, v3)
                    case "hello":
                        fmt.Println(fi16)
                    }
                    v15 -= 2147483647
                    v4 -= fi14
                }
                _ = fi16
            }
        }
        _ = fi14
        _ = v15
    }
    var v21 float32 = float32(0.1)
    for fi22 := 0; fi22 < 14; fi22++ {
        if ((!true) || ("foo bar" <= "a")) {
            for fi23 := 0; fi23 < 5; fi23++ {
                v24 := GenS1{F0: ((false || false) && true)}
                v25 := GenS2{F0: ("a" + ("foo bar" + "\n\t")), F1: genFn0(53, 161, -1)}
                v26 := GenS3{F0: 255}
                _ = fi23
                _ = v24
                _ = v25
                _ = v26
            }
            _ = len(v1)
            v21 = -2.25
        }
        for fi27 := 0; fi27 < 9; fi27++ {
            if (32767 <= 32767) {
                if true {
                    v21 = v21
                }
                var v28 int16 = -945
                var v29 int16 = v28
                _ = v28
                _ = v29
            } else {
                if false {
                    var v30 int16 = 1000
                    fmt.Println(v3)
                    _ = v30
                }
                v3 += v3
            }
            var v31 uint16 = 42
            v32 := v2[(fi22 * v4)]
            _ = fi27
            _ = v31
            _ = v32
        }
        v33 := GenS2{F0: "\n\t", F1: 362}
        _ = fi22
        _ = v33
    }
    switch v4 {
    case -9223372036854775808:
        for fi34 := 0; fi34 < 2; fi34++ {
            v35 := v2[(^(fi34))]
            _ = fi34
            _ = v35
        }
    case 9223372036854775807:
        var v36 int32 = 192
        _ = v36
    case 392:
        var v37 float32 = v21
        _ = v37
    default:
        if (!(1 <= 1)) {
            var v38 uint16 = 0
            v39 := GenS2{F0: ("\n\t" + ""), F1: genFn0(-1, 2, 279)}
            var v40 uint32 = 2
            _ = v38
            _ = v39
            _ = v40
        }
    }
    v41 := GenS4{F0: 42, F1: 1, F2: (0 == 364)}
    if false {
        if true {
            fmt.Println(v3)
            var v42 int8 = 127
            _ = v42
        } else {
            for fi43 := 0; fi43 < 7; fi43++ {
                v44 := GenS5{F0: 18446744073709551615, F1: 42, F2: -1}
                v3 = (v3 << ((18446744073709551615) & 63))
                _ = fi43
                _ = v44
            }
        }
        for fi45 := 0; fi45 < 17; fi45++ {
            var v46 int = -942
            _ = fi45
            _ = v46
        }
    } else {
        var v47 uint = 0
        _ = v47
    }
    if (!false) {
        v48 := GenS5{F0: 1, F1: 18446744073709551615, F2: -1}
        v2[v4] = false
        if (!(false || false)) {
            v49 := GenS6{F0: 2, F1: -1}
            var v50 uint16 = 32767
            v21 = v21
            _ = v49
            _ = v50
        }
        _ = v48
    }
    var v51 uint64 = 4294967295
    switch (v4 - (v4 << ((0) & 63))) {
    case 128:
        if true {
            v51 += (v51 - 18446744073709551615)
        } else {
            _ = len(v1)
            for fi52 := 0; fi52 < 20; fi52++ {
                var v53 string = "hello"
                delete(v2, (v4 + (^(v4))))
                _ = fi52
                _ = v53
            }
        }
    case -9223372036854775808:
        for fi54 := 0; fi54 < 10; fi54++ {
            v21 += 1e10
            for fi55 := 0; fi55 < 11; fi55++ {
                var v56 uint8 = 127
                v4 = fi54
                if false {
                    var v57 string = "\n\t"
                    if (!true) {
                        if false {
                            v58 := GenS7{F0: 42, F1: (v56 % ((232) | 1)), F2: 2}
                            _ = v58
                        }
                    } else {
                        _ = len(v1)
                        v59 := v2[(fi54 & v4)]
                        var v60 string = (v57 + ((v57 + (((((("hello" + v57) + ("\n\t" + v57)) + "hello") + v57) + (v57 + "hello")) + v57)) + v57))
                        _ = v59
                        _ = v60
                    }
                    fmt.Println(v4)
                    var v61 bool = (-304 <= -1)
                    var v62 float64 = float64(v21)
                    _ = v57
                    _ = v61
                    _ = v62
                }
                fmt.Println(v3)
                _ = fi55
                _ = v56
            }
            _ = fi54
        }
    }
    if (genFn0(127, 255, 0) > 255) {
        v63 := GenS3{F0: genFn0(-120, 127, 1000)}
        _ = v63
    } else {
        fmt.Println(v21, v3, v3)
        for fi64 := 0; fi64 < 18; fi64++ {
            v21 -= 3.14159
            v3 = v3
            if false {
                var v65 uint8 = 227
                v66 := GenS6{F0: 292, F1: 1}
                var v67 bool = ((fi64 ^ int(v51)) != v4)
                _ = v65
                _ = v66
                _ = v67
            }
            _ = fi64
        }
        if (2 == 11) {
            fmt.Println(v21)
            var v68 float64 = (1.5 / -2.25)
            _ = v68
        } else {
            v69 := GenS4{F0: 127, F1: 1, F2: false}
            v70 := GenS5{F0: (v51 % ((v51) | 1)), F1: (^(v51)), F2: 1}
            for fi71 := 0; fi71 < 13; fi71++ {
                for fi72 := 0; fi72 < 3; fi72++ {
                    v4 = -420
                    var v73 uint16 = 255
                    var v74 float32 = v21
                    _ = fi72
                    _ = v73
                    _ = v74
                }
                _ = fi71
            }
            _ = v69
            _ = v70
        }
    }
    v75 := GenS4{F0: 68, F1: 110, F2: (!(!(((v51 & v51) < (^(v51))) || false)))}
    if (!((-539 < 1) || true)) {
        v76 := GenS8{F0: 65535, F1: 32767, F2: (!true)}
        v1 = append(v1, "\n\t")
        _ = v76
    }
    v51 = 873
    v4 = v4
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v41
    _ = v51
    _ = v75
}
