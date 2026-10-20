package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint; F1 int }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint64; F1 uint16; F2 uint64 }
type GenS6 struct { F0 int64; F1 int64; F2 int64 }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 string; F1 uint; F2 int }
func genFn0(p0 uint32) int64 {
    switch "foo bar" {
    case "\n\t":
        switch ("\n\t" + "hello") {
        case "":
            v4 := GenS1{F0: 199}
            _ = v4
        case "a":
            for fi5 := 0; fi5 < 11; fi5++ {
                p0 -= 65535
                var v6 int16 = -1
                p0 -= (p0 << ((65535) & 63))
                _ = fi5
                _ = v6
            }
        case "foo bar":
            p0 = p0
        }
    default:
        for fi7 := 0; fi7 < 10; fi7++ {
            for fi8 := 0; fi8 < 13; fi8++ {
                v9 := GenS1{F0: 128}
                var v10 int16 = -931
                v10 = (-(v10))
                v11 := GenS1{F0: 255}
                v12 := GenS2{F0: 42, F1: fi8}
                v13 := []int{fi7, (fi7 - (fi8 << ((42) & 63))), fi7}
                _ = fi8
                _ = v10
                _ = v11
                _ = v12
                _ = v13
                _ = v9
            }
            switch (("\n\t" + "a") + (("a" + "a") + "foo bar")) {
            case "\n\t":
                p0 = p0
            }
            v14 := []int{fi7, (fi7 % ((fi7) | 1)), 675}
            _ = fi7
            _ = v14
        }
    }
    fmt.Println(p0)
    var v15 float32 = (1e10 - (((-2.25 - (-2.25 * ((1e10 + 1.5) * float32(-2.25)))) / 1.0) - float32(-2.25)))
    _ = p0
    _ = v15
    return -30
}
func genFn1(p0 int16, p1 int64, p2 int) int8 {
    fmt.Println(p2, p0)
    for fi16 := 0; fi16 < 7; fi16++ {
        for fi17 := 0; fi17 < 2; fi17++ {
            v18 := []int{(fi16 << ((0) & 63)), fi16, (p2 % ((fi16) | 1))}
            _ = fi17
            _ = v18
        }
        switch p2 {
        case -9223372036854775808:
            p2 = 9223372036854775807
        case 9223372036854775807:
            v19 := GenS3{F0: 200}
            _ = v19
        case -1:
            p1 += (p1 >> ((65535) & 63))
        }
        _ = fi16
    }
    var v20 int32 = 2147483647
    var v21 uint32 = 65535
    _ = p0
    _ = p1
    _ = p2
    _ = v20
    _ = v21
    return 53
}

func main() {
    var v1 int32 = -1
    var v2 int64 = 42
    v3 := make(map[string]bool)
    fmt.Println(v1, v1)
    var v22 string = ("" + "")
    var v23 int32 = v1
    v23 += 42
    if false {
        var v24 int32 = -353
        _ = v24
    } else {
        for fi25 := 0; fi25 < 20; fi25++ {
            v23 -= (-(v23))
            var v26 uint8 = 127
            _ = fi25
            _ = v26
        }
    }
    var v27 string = (((v22 + ("foo bar" + v22)) + v22) + "a")
    if false {
        v22 = ((((v22 + ((v22 + "a") + "a")) + v22) + v27) + v22)
        for fi28 := 0; fi28 < 9; fi28++ {
            v3["foo bar"] = true
            _ = fi28
        }
        v29 := GenS1{F0: 0}
        _ = v29
    } else {
        if (true && true) {
            switch 695 {
            case 32767:
                v1 += (v23 / ((v1) | 1))
            case -2147483648:
                var v30 uint64 = 1
                _ = v30
            case 1:
                v3[(v22 + v22)] = ((255 > 2) && false)
            }
            v22 = v27
            delete(v3, v27)
        }
    }
    if false {
        fmt.Println(v22, v27)
    } else {
        v1 -= (^(v23))
        fmt.Println(v27, v27)
        v31 := GenS2{F0: 834, F1: 807}
        _ = v31
    }
    v1 += (v23 & v23)
    if true {
        for fi32 := 0; fi32 < 16; fi32++ {
            switch (fi32 / ((127) | 1)) {
            case 32768:
                v33 := []int{}
                _ = v33
            default:
                if false {
                    v23 = (v1 << ((255) & 63))
                }
            }
            _ = fi32
        }
        switch (("foo bar" + "") + v22) {
        case "hello":
            v34 := []int{4294967295, 32767}
            _ = v34
        case "\n\t":
            fmt.Println(v2)
        default:
            if (0.0 == (1e10 / 1.0)) {
                for fi35 := 0; fi35 < 18; fi35++ {
                    v22 = ((("foo bar" + "hello") + "\n\t") + ((v22 + v27) + ("a" + (("hello" + (v27 + v27)) + v22))))
                    fmt.Println(fi35, v2)
                    _ = fi35
                }
                if true {
                    v23 = v23
                    var v36 bool = (!true)
                    var v37 float64 = ((-2.25 + (1e10 - (1e10 - 1e10))) * 3.14159)
                    _ = v36
                    _ = v37
                } else {
                    if true {
                        if true {
                            if false {
                                if (42 == 255) {
                                    v2 += 42
                                    for fi38 := 0; fi38 < 12; fi38++ {
                                        v22 = v27
                                        _ = fi38
                                    }
                                } else {
                                    v1 = v23
                                    var v39 uint32 = 2147483648
                                    _ = v39
                                }
                                var v40 uint = 18446744073709551615
                                v27 = v22
                                _ = v40
                            }
                        }
                        if (301 > 4294967295) {
                            if true {
                                v41 := GenS2{F0: 18446744073709551615, F1: 1}
                                fmt.Println(v2, v27, v1)
                                v42 := GenS4{F0: 0}
                                _ = v41
                                _ = v42
                            } else {
                                if (!(!true)) {
                                    for fi43 := 0; fi43 < 15; fi43++ {
                                        var v44 uint = 42
                                        v22 = (v27 + "hello")
                                        v27 = v27
                                        _ = fi43
                                        _ = v44
                                    }
                                    var v45 uint64 = 4294967295
                                    _ = v45
                                }
                                if true {
                                    v22 = "\n\t"
                                    switch 1 {
                                    case 127:
                                        var v46 int32 = v23
                                        _ = v46
                                    case 32768:
                                        var v47 uint8 = 0
                                        _ = v47
                                    default:
                                        v1 = v23
                                    }
                                }
                            }
                            switch 345 {
                            case -9223372036854775808:
                                v23 = v23
                            default:
                                v48 := GenS5{F0: 4294967295, F1: 53, F2: 4294967295}
                                _ = v48
                            }
                        } else {
                            var v49 uint32 = 2147483648
                            v50 := GenS1{F0: 127}
                            switch v27 {
                            case "hello":
                                v27 = "a"
                            }
                            _ = v49
                            _ = v50
                        }
                        switch -2147483648 {
                        case 32767:
                            switch 32768 {
                            case -147:
                                if true {
                                    switch (v27 + "foo bar") {
                                    case "a":
                                        v51 := GenS6{F0: v2, F1: 0, F2: v2}
                                        _ = v51
                                    case "":
                                        var v52 uint16 = 32767
                                        _ = v52
                                    case "foo bar":
                                        if true {
                                            v2 += (v2 << ((1) & 63))
                                            for fi53 := 0; fi53 < 7; fi53++ {
                                                if (!(-32768 > 1000)) {
                                                    if false {
                                                        v23 -= (v1 * (^(v23)))
                                                        v1 -= (^(v23))
                                                        fmt.Println(v1, v23, fi53)
                                                    }
                                                    v54 := []int{(fi53 << ((65535) & 63)), (fi53 ^ (fi53 | fi53)), (fi53 % (((fi53 - fi53)) | 1)), (fi53 * int(v1))}
                                                    var v55 uint16 = 376
                                                    _ = v54
                                                    _ = v55
                                                }
                                                var v56 float64 = float64(float32(1e10))
                                                v2 += v2
                                                _ = fi53
                                                _ = v56
                                            }
                                            v57 := GenS2{F0: 255, F1: 128}
                                            _ = v57
                                        } else {
                                            for fi58 := 0; fi58 < 19; fi58++ {
                                                v59 := []int{127, fi58, int(v2), -2147483648}
                                                _ = fi58
                                                _ = v59
                                            }
                                        }
                                    }
                                } else {
                                    for fi60 := 0; fi60 < 16; fi60++ {
                                        var v61 uint8 = 128
                                        _ = fi60
                                        _ = v61
                                    }
                                    fmt.Println(v1, v1)
                                }
                            case -1:
                                var v62 int32 = 380
                                _ = v62
                            }
                        default:
                            v63 := GenS7{F0: 2}
                            _ = v63
                        }
                    }
                }
                var v64 uint = 326
                _ = v64
            }
        }
    }
    for fi65 := 0; fi65 < 7; fi65++ {
        v66 := []int{-130, (-(fi65))}
        var v67 int16 = -32768
        _ = fi65
        _ = v66
        _ = v67
    }
    if (!false) {
        var v68 uint = 476
        if (58 >= 105) {
            var v69 int = 128
            _ = v69
        }
        switch 127 {
        case -568:
            var v70 uint8 = 105
            _ = v70
        }
        _ = v68
    } else {
        for fi71 := 0; fi71 < 4; fi71++ {
            v72 := GenS8{F0: "hello", F1: 2, F2: 127}
            _ = fi71
            _ = v72
        }
        fmt.Println(v23, v2, v23)
    }
    v73 := GenS8{F0: v27, F1: 705, F2: 0}
    v74 := GenS7{F0: 2}
    v1 -= (v1 | 90)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v27
    _ = v3
    _ = v73
    _ = v74
}
