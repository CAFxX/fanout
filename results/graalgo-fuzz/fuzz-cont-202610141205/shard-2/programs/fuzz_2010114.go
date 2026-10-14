package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16 }
type GenS2 struct { F0 int8; F1 uint }
type GenS3 struct { F0 int32; F1 uint }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint }
func genFn0(p0 string) uint64 {
    var v5 uint64 = 18446744073709551615
    v5 = v5
    fmt.Println(p0)
    for fi6 := 0; fi6 < 20; fi6++ {
        v5 -= v5
        _ = fi6
    }
    if false {
        switch p0 {
        case "a":
            v7 := GenS1{F0: (p0 + (p0 + p0)), F1: 1}
            _ = v7
        case "":
            v8 := GenS1{F0: (((p0 + ((p0 + "a") + p0)) + (p0 + "a")) + p0), F1: 749}
            _ = v8
        case "foo bar":
            v9 := GenS1{F0: p0, F1: 342}
            _ = v9
        default:
            var v10 int64 = 812
            _ = v10
        }
        v11 := []int{42}
        if (0 == 2147483647) {
            var v12 uint64 = v5
            if true {
                v5 = 782
            }
            if ((v12 > v5) || true) {
                var v13 int16 = 1000
                _ = v13
            }
            _ = v12
        }
        _ = v11
    } else {
        v14 := GenS2{F0: 100, F1: 636}
        v15 := GenS1{F0: ((p0 + "hello") + p0), F1: 990}
        p0 = (("a" + "\n\t") + "hello")
        _ = v14
        _ = v15
    }
    _ = p0
    _ = v5
    return 857
}

func main() {
    var v1 float64 = (0.0 / (3.14159 / 1.5))
    v2 := make(map[int]uint32)
    v3 := []string{}
    v4 := make(map[uint32]uint32)
    v1 -= ((v1 / v1) - v1)
    var v16 bool = false
    switch 127 {
    case 393:
        for fi17 := 0; fi17 < 15; fi17++ {
            var v18 int8 = 93
            v1 -= -2.25
            _ = fi17
            _ = v18
        }
    case -2147483648:
        for fi19 := 0; fi19 < 10; fi19++ {
            v1 += float64(0.1)
            v20 := v2[784]
            _ = fi19
            _ = v20
        }
    case -139:
        for fi21 := 0; fi21 < 7; fi21++ {
            var v22 uint32 = 2
            switch fi21 {
            case 128:
                fmt.Println(v22, v1, fi21)
            case 32768:
                delete(v4, 1)
            case -381:
                var v23 uint32 = (v22 >> ((4294967295) & 63))
                _ = v23
            default:
                var v24 int16 = 1000
                _ = v24
            }
            fmt.Println(v22, v16, fi21)
            _ = fi21
            _ = v22
        }
    }
    var v25 int32 = 2147483647
    var v26 int16 = -544
    switch ("hello" + "a") {
    case "":
        var v27 bool = v16
        _ = v27
    default:
        switch "a" {
        case "hello":
            if ((1 < -546) || (2 >= 2)) {
                fmt.Println(v25)
                v26 += 42
            } else {
                v28 := GenS3{F0: 1, F1: 2}
                v29 := GenS3{F0: (v25 / ((v25) | 1)), F1: 0}
                _ = v28
                _ = v29
            }
        default:
            var v30 uint64 = 42
            _ = v30
        }
    }
    var v31 int8 = 42
    for fi32 := 0; fi32 < 5; fi32++ {
        v31 = (v31 / ((v31) | 1))
        fmt.Println(v31, v31, v1)
        v26 += v26
        _ = fi32
    }
    if (2147483648 < 2147483647) {
        switch 42 {
        case 472:
            if (422 > 0) {
                v1 = v1
            }
        case 0:
            switch -599 {
            case 576:
                v3 = append(v3, "\n\t")
            case 244:
                for fi33 := 0; fi33 < 11; fi33++ {
                    var v34 float64 = v1
                    var v35 string = "foo bar"
                    if true {
                        if (v16 || v16) {
                            var v36 uint8 = 2
                            var v37 int8 = 11
                            var v38 uint8 = (v36 % (((v36 ^ (v36 * (v36 | v36)))) | 1))
                            _ = v36
                            _ = v37
                            _ = v38
                        } else {
                            v39 := GenS3{F0: v25, F1: 65535}
                            var v40 uint16 = 32767
                            var v41 uint16 = v40
                            _ = v39
                            _ = v40
                            _ = v41
                        }
                        if v16 {
                            v35 = v35
                            var v42 int16 = 42
                            switch (v35 + (v35 + ((v35 + (v35 + (v35 + v35))) + ("hello" + v35)))) {
                            case "foo bar":
                                v43 := GenS1{F0: ("a" + v35), F1: 2}
                                _ = v43
                            case "hello":
                                if len(v3) > 0 {
                                    v3[0] = v35
                                }
                            case "":
                                if v16 {
                                    var v44 uint16 = 32767
                                    var v45 int8 = (v31 - -119)
                                    var v46 int32 = (v25 & v25)
                                    _ = v44
                                    _ = v45
                                    _ = v46
                                } else {
                                    var v47 bool = false
                                    delete(v4, 1)
                                    v48 := GenS4{F0: 0}
                                    _ = v47
                                    _ = v48
                                }
                            }
                            _ = v42
                        }
                        _ = len(v4)
                    }
                    _ = fi33
                    _ = v34
                    _ = v35
                }
            default:
                if v16 {
                    var v49 int16 = (v26 & (v26 ^ 1000))
                    _ = len(v2)
                    _ = v49
                }
            }
        case 42:
            if (32767 > -340) {
                if (v26 <= 1) {
                    var v50 uint8 = 255
                    _ = v50
                } else {
                    if (-562 > -991) {
                        for fi51 := 0; fi51 < 11; fi51++ {
                            var v52 string = ("foo bar" + ("a" + "foo bar"))
                            v25 -= (v25 >> ((1) & 63))
                            v31 -= int8(v26)
                            _ = fi51
                            _ = v52
                        }
                    } else {
                        for fi53 := 0; fi53 < 6; fi53++ {
                            switch (fi53 ^ fi53) {
                            case -160:
                                var v54 int64 = 884
                                _ = v54
                            case -644:
                                var v55 int8 = (-(v31))
                                _ = v55
                            case 9223372036854775807:
                                v25 += v25
                            }
                            _ = fi53
                        }
                    }
                }
            } else {
                for fi56 := 0; fi56 < 8; fi56++ {
                    v57 := GenS5{F0: 4294967295}
                    v26 -= v26
                    if true {
                        v1 += float64(0.0)
                    } else {
                        v58 := GenS4{F0: 1}
                        v16 = (true || v16)
                        _ = v58
                    }
                    _ = fi56
                    _ = v57
                }
                var v59 uint64 = 0
                _ = v59
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v16
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v31
    _ = v4
}
