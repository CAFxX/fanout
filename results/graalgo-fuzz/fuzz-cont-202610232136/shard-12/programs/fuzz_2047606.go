package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int32; F1 int; F2 uint32 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 string; F1 uint; F2 uint32 }
type GenS5 struct { F0 uint64; F1 int; F2 uint32 }
type GenS6 struct { F0 int32; F1 int16; F2 string }
type GenS7 struct { F0 uint32 }
func genFn0(p0 int, p1 uint32) uint32 {
    var v4 uint64 = 2
    v5 := GenS1{F0: (p1 ^ 755)}
    p0 = (p0 | 127)
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 2
}
func genFn1() int {
    var v6 uint32 = genFn0(-938, 42)
    if (true && (255 != 127)) {
        v7 := []int{0, 2147483647, 2147483647}
        fmt.Println(v6)
        _ = v7
    } else {
        fmt.Println(v6)
    }
    var v8 float64 = 0.1
    v9 := []int{}
    if (false && true) {
        for fi10 := 0; fi10 < 13; fi10++ {
            if (!(!(351 > 701))) {
                for fi11 := 0; fi11 < 9; fi11++ {
                    if len(v9) >= 2 {
                        v9 = v9[1:len(v9)]
                    }
                    var v12 bool = true
                    switch "a" {
                    case "\n\t":
                        var v13 int16 = 0
                        _ = v13
                    case "":
                        if ((!(v12 || v12)) || (fi10 >= 127)) {
                            v6 += (v6 / ((v6) | 1))
                            if len(v9) > 0 {
                                v9[0] = (fi10 % (((fi10 & fi11)) | 1))
                            }
                            var v14 float32 = -2.25
                            _ = v14
                        }
                    default:
                        var v15 uint32 = (v6 % ((v6) | 1))
                        v6 = v15
                        var v16 float32 = 0.1
                        _ = v15
                        _ = v16
                    }
                    if v12 {
                        if true {
                            var v17 int = (fi11 + (fi10 / (((fi11 << ((827) & 63))) | 1)))
                            _ = v17
                        } else {
                            v8 -= (v8 + float64(-2.25))
                        }
                        v18 := GenS1{F0: v6}
                        v19 := GenS1{F0: (v6 | (v6 ^ v6))}
                        _ = v18
                        _ = v19
                    }
                    var v20 int8 = 111
                    _ = fi11
                    _ = v12
                    _ = v20
                }
            } else {
                if true {
                    fmt.Println(fi10, v8, v8)
                    var v21 uint16 = 65535
                    if len(v9) >= 2 {
                        v9 = v9[1:len(v9)]
                    }
                    _ = v21
                } else {
                    v6 += v6
                    for fi22 := 0; fi22 < 14; fi22++ {
                        fmt.Println(v8, v6)
                        v23 := GenS2{F0: 556, F1: 4294967295, F2: (v6 & v6)}
                        var v24 int16 = 42
                        _ = fi22
                        _ = v23
                        _ = v24
                    }
                    v25 := GenS3{F0: (fi10 + 4294967295)}
                    _ = v25
                }
                if len(v9) > 0 {
                    v9[0] = (fi10 >> ((255) & 63))
                }
                v8 += v8
            }
            _ = fi10
        }
        v8 += (v8 + -2.25)
    }
    _ = v6
    _ = v8
    _ = v9
    return -464
}

func main() {
    v1 := make(map[int]uint32)
    var v2 bool = false
    var v3 int64 = 0
    v26 := []int{}
    switch "hello" {
    case "hello":
        if v2 {
            switch -9223372036854775808 {
            case 0:
                var v27 string = ("hello" + (("a" + "\n\t") + "foo bar"))
                _ = v27
            case 32768:
                v28 := GenS1{F0: 65535}
                _ = v28
            case 128:
                for fi29 := 0; fi29 < 14; fi29++ {
                    var v30 uint16 = 32767
                    _ = fi29
                    _ = v30
                }
            default:
                v31 := GenS4{F0: "\n\t", F1: 4294967295, F2: genFn0(128, 2147483647)}
                _ = v31
            }
        } else {
            for fi32 := 0; fi32 < 11; fi32++ {
                fmt.Println(v3)
                _ = fi32
            }
            for fi33 := 0; fi33 < 10; fi33++ {
                switch "a" {
                case "foo bar":
                    fmt.Println(fi33)
                case "":
                    v3 = (v3 & v3)
                default:
                    v2 = v2
                }
                _ = fi33
            }
        }
    default:
        v2 = ((-112 <= 127) && (42 <= 977))
    }
    for fi34 := 0; fi34 < 11; fi34++ {
        v35 := GenS5{F0: 4294967295, F1: (fi34 / ((fi34) | 1)), F2: 344}
        _ = fi34
        _ = v35
    }
    if (!((0.1 * 1e10) != 1.5)) {
        var v36 int8 = -1
        _ = v36
    } else {
        var v37 uint = 18446744073709551615
        for fi38 := 0; fi38 < 14; fi38++ {
            var v39 uint32 = 604
            var v40 uint16 = 1
            var v41 int16 = -1
            _ = fi38
            _ = v39
            _ = v40
            _ = v41
        }
        v42 := GenS6{F0: 0, F1: 1000, F2: "hello"}
        _ = v37
        _ = v42
    }
    var v43 bool = v2
    for fi44 := 0; fi44 < 10; fi44++ {
        var v45 bool = true
        switch ("a" + "\n\t") {
        case "\n\t":
            for fi46 := 0; fi46 < 15; fi46++ {
                switch "\n\t" {
                case "a":
                    switch "a" {
                    case "\n\t":
                        v45 = (32767 > 0)
                        switch fi44 {
                        case 0:
                            delete(v1, 1)
                        case 980:
                            var v47 uint16 = 65535
                            _ = v47
                        default:
                            v45 = ((((-2.25 * float64(1.5)) / (((-2.25 - 0.1) - 0.1) / (3.14159 + 0.0))) / 0.1) == -2.25)
                        }
                    }
                case "":
                    v45 = true
                default:
                    fmt.Println(v43)
                }
                v26 = append(v26, (fi44 >> ((653) & 63)))
                var v48 string = ("hello" + "")
                _ = fi46
                _ = v48
            }
        case "":
            if v45 {
                v43 = v2
                switch "hello" {
                case "":
                    switch (fi44 * fi44) {
                    case 9223372036854775807:
                        switch (("foo bar" + ("hello" + "\n\t")) + "a") {
                        case "foo bar":
                            delete(v1, fi44)
                        case "":
                            fmt.Println(v2)
                        case "\n\t":
                            _ = len(v1)
                        default:
                            v45 = (((42 <= -963) && (1 >= 4294967295)) || false)
                        }
                    case -436:
                        for fi49 := 0; fi49 < 10; fi49++ {
                            fmt.Println(v3, v2)
                            v50 := GenS6{F0: -566, F1: 32767, F2: "\n\t"}
                            _ = fi49
                            _ = v50
                        }
                    }
                case "\n\t":
                    v51 := GenS3{F0: (-(fi44))}
                    _ = v51
                default:
                    switch "hello" {
                    case "":
                        v2 = v45
                    default:
                        if len(v26) >= 2 {
                            v26 = v26[1:len(v26)]
                        }
                    }
                }
            }
        case "a":
            for fi52 := 0; fi52 < 19; fi52++ {
                v1[0] = 725
                var v53 int32 = 2147483647
                _ = fi52
                _ = v53
            }
        }
        v2 = false
        _ = fi44
        _ = v45
    }
    v54 := v1[127]
    for fi55 := 0; fi55 < 10; fi55++ {
        v43 = false
        v3 = 1
        _ = fi55
    }
    if v2 {
        for fi56 := 0; fi56 < 6; fi56++ {
            v57 := v1[(fi56 % ((fi56) | 1))]
            if v2 {
                v57 += (v54 * (v57 % ((v54) | 1)))
                var v58 uint8 = 1
                for fi59 := 0; fi59 < 9; fi59++ {
                    if (float32(1e10) < (0.1 / 3.14159)) {
                        v3 = v3
                        v3 -= (v3 & v3)
                        v3 = v3
                    }
                    if (float64((1.5 + 0.0)) >= -2.25) {
                        var v60 int16 = 1
                        fmt.Println(fi59, v58)
                        _ = v60
                    }
                    v2 = true
                    _ = fi59
                }
                _ = v58
            }
            v2 = (v43 && (1e10 >= 1e10))
            _ = fi56
            _ = v57
        }
    } else {
        v61 := GenS7{F0: (v54 * 2147483647)}
        _ = v61
    }
    v3 -= 42
    v62 := GenS6{F0: 42, F1: -814, F2: "hello"}
    v63 := GenS4{F0: "\n\t", F1: 678, F2: 65535}
    if len(v26) >= 2 {
        v26 = v26[1:len(v26)]
    }
    fmt.Println(v3, v43, v3)
    delete(v1, 128)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v3:", v3)
    fmt.Println("v43:", v43)
    fmt.Println("v54:", v54)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v43
    _ = v54
    _ = v62
    _ = v63
}
