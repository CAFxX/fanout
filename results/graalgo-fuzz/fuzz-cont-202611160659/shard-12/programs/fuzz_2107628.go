package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32; F2 uint }
type GenS2 struct { F0 uint64; F1 int64 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int16; F1 int64 }
type GenS5 struct { F0 int32; F1 bool; F2 string }
type GenS6 struct { F0 int8; F1 int64 }
type GenS7 struct { F0 uint8; F1 int8 }
type GenS8 struct { F0 uint32; F1 string; F2 string }
type GenS9 struct { F0 string }
type GenS10 struct { F0 int16; F1 string; F2 int64 }
func genFn0(p0 uint16, p1 uint16) uint64 {
    var v4 int32 = -1
    v4 = v4
    for fi5 := 0; fi5 < 2; fi5++ {
        if (!((true && (1e10 >= 1e10)) && true)) {
            switch (("foo bar" + "foo bar") + "") {
            case "\n\t":
                v6 := GenS1{F0: false, F1: v4, F2: 92}
                _ = v6
            case "foo bar":
                v7 := GenS1{F0: true, F1: (v4 & (v4 | v4)), F2: 0}
                _ = v7
            }
            switch "hello" {
            case "foo bar":
                for fi8 := 0; fi8 < 13; fi8++ {
                    var v9 uint32 = 4294967295
                    v10 := GenS2{F0: 42, F1: 0}
                    _ = fi8
                    _ = v10
                    _ = v9
                }
            case "hello":
                p1 -= (p0 & (p1 * (p0 - 255)))
            default:
                p0 = 42
            }
        }
        fmt.Println(v4)
        _ = fi5
    }
    v4 = (v4 / ((v4) | 1))
    var v11 float32 = ((1.5 - 0.1) * float32(3.14159))
    _ = p0
    _ = p1
    _ = v11
    _ = v4
    return 2
}
func genFn1(p0 uint8) uint {
    v12 := GenS3{F0: -810}
    p0 -= p0
    var v13 uint = 122
    v14 := GenS2{F0: 18446744073709551615, F1: -1}
    _ = p0
    _ = v12
    _ = v13
    _ = v14
    return 1
}

func main() {
    v1 := make(map[int]int)
    var v2 bool = false
    var v3 uint = 503
    delete(v1, -514)
    fmt.Println(v2, v2)
    for fi15 := 0; fi15 < 14; fi15++ {
        var v16 float32 = ((((3.14159 - (-2.25 / 1.0)) - 0.0) * -2.25) - 1e10)
        _ = fi15
        _ = v16
    }
    v3 = (^(v3))
    if (644 <= 0) {
        var v17 string = "a"
        _ = v17
    } else {
        switch "a" {
        case "a":
            v18 := GenS4{F0: 1000, F1: -1}
            _ = v18
        case "hello":
            v3 = (v3 >> ((v3) & 63))
        default:
            delete(v1, 42)
        }
        for fi19 := 0; fi19 < 13; fi19++ {
            v20 := GenS2{F0: 682, F1: 178}
            _ = fi19
            _ = v20
        }
        for fi21 := 0; fi21 < 15; fi21++ {
            for fi22 := 0; fi22 < 17; fi22++ {
                var v23 uint = 0
                if v2 {
                    v3 = genFn1(128)
                } else {
                    v24 := GenS5{F0: 0, F1: (v2 || v2), F2: "a"}
                    switch "hello" {
                    case "a":
                        v3 += v23
                    default:
                        v25 := []int{42, 0, fi21, 0}
                        _ = v25
                    }
                    _ = v24
                }
                if v2 {
                    if (282 != 921) {
                        var v26 string = ""
                        var v27 int64 = 0
                        _ = v26
                        _ = v27
                    } else {
                        fmt.Println(v23, v23)
                        var v28 float64 = (-2.25 / 3.14159)
                        _ = v28
                    }
                    switch fi22 {
                    case 2147483647:
                        fmt.Println(fi21, fi22, v23)
                    case 37:
                        switch "foo bar" {
                        case "hello":
                            var v29 float64 = 0.1
                            if (128 < 0) {
                                var v30 bool = (0.0 != (float32(-2.25) * 1e10))
                                fmt.Println(v29)
                                _ = v30
                            } else {
                                var v31 uint8 = 86
                                var v32 float32 = -2.25
                                _ = v31
                                _ = v32
                            }
                            var v33 int64 = -741
                            _ = v29
                            _ = v33
                        }
                    case 32768:
                        if true {
                            if (-32768 >= 936) {
                                v1[-2147483648] = 0
                            } else {
                                _ = len(v1)
                            }
                            v34 := []int{4294967295, int(v23), fi22, 70}
                            v2 = (42 >= 64)
                            _ = v34
                        }
                    default:
                        var v35 uint64 = 2
                        _ = v35
                    }
                }
                _ = fi22
                _ = v23
            }
            fmt.Println(fi21, fi21)
            v36 := GenS6{F0: 42, F1: -9223372036854775808}
            _ = fi21
            _ = v36
        }
    }
    v1[2147483647] = 9223372036854775807
    if v2 {
        var v37 int8 = 100
        switch 905 {
        case 309:
            if (1000 <= 42) {
                for fi38 := 0; fi38 < 20; fi38++ {
                    for fi39 := 0; fi39 < 11; fi39++ {
                        fmt.Println(v37, fi39)
                        v40 := GenS4{F0: 1, F1: 1}
                        v41 := GenS3{F0: 508}
                        _ = fi39
                        _ = v40
                        _ = v41
                    }
                    fmt.Println(v37, v2, v3)
                    switch fi38 {
                    case 1:
                        _ = len(v1)
                    case 305:
                        if (!v2) {
                            if ((!v2) && true) {
                                fmt.Println(v2, v37, v2)
                                if (v2 || false) {
                                    fmt.Println(v37, v3)
                                    var v42 uint64 = genFn0(1, 1)
                                    switch "" {
                                    case "\n\t":
                                        for fi43 := 0; fi43 < 2; fi43++ {
                                            switch "" {
                                            case "hello":
                                                v2 = true
                                            }
                                            if (42 >= 65535) {
                                                v42 = v42
                                            }
                                            if v2 {
                                                v42 -= (v42 ^ 42)
                                                fmt.Println(v42, v2)
                                                var v44 string = "\n\t"
                                                v45 := []int{(^(fi43)), (fi38 % (((fi38 << (((^(v3))) & 63))) | 1)), (^(fi38)), (-(fi38))}
                                                fmt.Println(v42)
                                                _ = v44
                                                _ = v45
                                            }
                                            _ = fi43
                                        }
                                    case "":
                                        switch (("" + "foo bar") + "foo bar") {
                                        case "a":
                                            v46 := []int{fi38, 700, -2147483648}
                                            _ = v46
                                        case "":
                                            var v47 int8 = v37
                                            _ = v47
                                        case "hello":
                                            v48 := GenS2{F0: genFn0(207, 870), F1: int64(v3)}
                                            _ = v48
                                        }
                                    case "foo bar":
                                        fmt.Println(v3, v3, v42)
                                    }
                                    _ = v42
                                } else {
                                    v49 := []int{668}
                                    v37 = v37
                                    _ = v49
                                }
                            } else {
                                fmt.Println(v37, fi38)
                                fmt.Println(v3, v3)
                            }
                        }
                    case -9223372036854775808:
                        var v50 int = fi38
                        _ = v50
                    default:
                        v2 = v2
                    }
                    _ = fi38
                }
                for fi51 := 0; fi51 < 20; fi51++ {
                    var v52 string = ("" + "\n\t")
                    _ = fi51
                    _ = v52
                }
            } else {
                for fi53 := 0; fi53 < 5; fi53++ {
                    if v2 {
                        _ = len(v1)
                        var v54 uint32 = 63
                        fmt.Println(v2)
                        _ = v54
                    } else {
                        v55 := GenS7{F0: 2, F1: v37}
                        if v2 {
                            var v56 string = "\n\t"
                            fmt.Println(v3)
                            v57 := GenS4{F0: -32768, F1: -1}
                            _ = v56
                            _ = v57
                        } else {
                            switch ("\n\t" + "foo bar") {
                            case "\n\t":
                                if v2 {
                                    for fi58 := 0; fi58 < 16; fi58++ {
                                        fmt.Println(v2, v3)
                                        _ = fi58
                                    }
                                }
                            case "hello":
                                var v59 uint16 = 255
                                _ = v59
                            default:
                                var v60 int32 = -2147483648
                                _ = v60
                            }
                            var v61 uint32 = 64
                            _ = v61
                        }
                        _ = v55
                    }
                    fmt.Println(v3)
                    for fi62 := 0; fi62 < 7; fi62++ {
                        var v63 int = 0
                        _ = fi62
                        _ = v63
                    }
                    _ = fi53
                }
            }
        case 364:
            v64 := GenS8{F0: 2147483647, F1: "\n\t", F2: "a"}
            _ = v64
        default:
            v65 := GenS4{F0: 558, F1: 450}
            _ = v65
        }
        _ = v37
    }
    for fi66 := 0; fi66 < 15; fi66++ {
        var v67 uint16 = 0
        if v2 {
            v68 := GenS3{F0: -1}
            _ = v68
        } else {
            v3 = (v3 / ((v3) | 1))
            for fi69 := 0; fi69 < 4; fi69++ {
                v70 := GenS9{F0: ("" + "")}
                v1[int(v3)] = (fi66 | (fi69 & fi66))
                var v71 string = ("hello" + ("foo bar" + "\n\t"))
                _ = fi69
                _ = v70
                _ = v71
            }
        }
        _ = fi66
        _ = v67
    }
    v72 := GenS10{F0: 0, F1: "", F2: -682}
    _ = len(v1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v3
    _ = v72
}
