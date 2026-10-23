package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int8; F1 string; F2 int64 }
type GenS4 struct { F0 int32; F1 int8; F2 uint32 }
type GenS5 struct { F0 uint32; F1 int64 }
type GenS6 struct { F0 int8; F1 uint64 }
type GenS7 struct { F0 uint64 }
func genFn0(p0 string, p1 uint16, p2 uint64) uint32 {
    fmt.Println(p0, p1)
    var v4 bool = (!(false && (0 <= 2)))
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 90
}
func genFn1(p0 int8, p1 uint64) int16 {
    p1 = p1
    for fi5 := 0; fi5 < 19; fi5++ {
        v6 := GenS1{F0: -198}
        for fi7 := 0; fi7 < 17; fi7++ {
            v8 := []int{fi5, (^(fi7)), (fi7 * (fi7 / (((fi5 - (fi5 / (((fi5 | (fi5 % (((fi5 | fi7)) | 1)))) | 1)))) | 1))), 32767}
            if false {
                if len(v8) > 0 {
                    v8[0] = fi7
                }
                var v9 uint8 = 128
                switch int(p0) {
                case -2147483648:
                    if (false || true) {
                        switch fi5 {
                        case 338:
                            if true {
                                fmt.Println(p0, p0, p1)
                            } else {
                                if len(v8) > 0 {
                                    v8[0] = fi7
                                }
                                var v10 float64 = -2.25
                                var v11 float64 = v10
                                _ = v10
                                _ = v11
                            }
                        }
                        fmt.Println(p1)
                        v12 := GenS2{F0: true}
                        _ = v12
                    } else {
                        v13 := GenS1{F0: 757}
                        fmt.Println(fi5, fi7, fi7)
                        _ = v13
                    }
                default:
                    if len(v8) > 0 {
                        v8[0] = (fi7 >> ((20) & 63))
                    }
                }
                var v14 int64 = -9223372036854775808
                p0 = 0
                v14 -= v14
                v15 := GenS1{F0: v14}
                if true {
                    var v16 float32 = (float32(0.0) * (-2.25 * (3.14159 * 0.0)))
                    _ = v16
                } else {
                    var v17 uint64 = 212
                    var v18 int64 = v14
                    p1 -= v17
                    _ = v17
                    _ = v18
                }
                _ = v14
                _ = v15
                _ = v9
            } else {
                if (false && (42 > 0)) {
                    v19 := GenS2{F0: (127 <= 221)}
                    _ = v19
                }
                var v20 int = 32767
                v21 := GenS2{F0: (!true)}
                var v22 uint16 = 516
                switch "foo bar" {
                case "a":
                    var v23 uint16 = v22
                    _ = v23
                }
                var v24 uint32 = 2
                _ = v20
                _ = v21
                _ = v22
                _ = v24
            }
            var v25 bool = ((("hello" > "\n\t") && (65535 > 0)) && false)
            _ = fi7
            _ = v25
            _ = v8
        }
        _ = fi5
        _ = v6
    }
    p0 = (p0 % ((-1) | 1))
    var v26 int = 0
    v26 = 127
    _ = p0
    _ = p1
    _ = v26
    return -32768
}

func main() {
    var v1 float64 = (1.5 + 0.0)
    v2 := make(map[string]int)
    var v3 uint = 42
    var v27 float64 = v1
    v2[("a" + "hello")] = -2147483648
    for fi28 := 0; fi28 < 16; fi28++ {
        v29 := []int{-1, -1, 42}
        _ = fi28
        _ = v29
    }
    var v30 int64 = -9223372036854775808
    var v31 uint16 = 42
    if (-98 < 83) {
        v27 = (1.5 - v27)
        v32 := GenS2{F0: true}
        var v33 uint8 = 127
        _ = v32
        _ = v33
    }
    for fi34 := 0; fi34 < 1; fi34++ {
        var v35 uint8 = 0
        v36 := []int{}
        _ = fi34
        _ = v35
        _ = v36
    }
    v31 -= (v31 & (v31 + 0))
    switch "hello" {
    case "hello":
        for fi37 := 0; fi37 < 19; fi37++ {
            switch (fi37 << ((v3) & 63)) {
            case 42:
                if (-1 != -2147483648) {
                    var v38 uint64 = 0
                    for fi39 := 0; fi39 < 8; fi39++ {
                        v40 := GenS3{F0: 42, F1: "\n\t", F2: v30}
                        _ = fi39
                        _ = v40
                    }
                    _ = v38
                }
            case 128:
                fmt.Println(v30, v3)
            case 4294967295:
                fmt.Println(v30, v27)
            default:
                if true {
                    if false {
                        v27 -= v1
                    }
                } else {
                    v41 := GenS3{F0: 0, F1: (("\n\t" + ("foo bar" + "hello")) + "hello"), F2: -9223372036854775808}
                    var v42 int64 = v30
                    v43 := GenS4{F0: 72, F1: 41, F2: 4294967295}
                    _ = v41
                    _ = v42
                    _ = v43
                }
            }
            if false {
                v44 := []int{}
                _ = v44
            } else {
                v45 := []int{9223372036854775807, 311}
                v46 := GenS5{F0: 42, F1: 89}
                _ = len(v45)
                _ = v45
                _ = v46
            }
            delete(v2, "hello")
            _ = fi37
        }
    case "\n\t":
        for fi47 := 0; fi47 < 10; fi47++ {
            fmt.Println(v31)
            fmt.Println(v3, v3)
            if (("\n\t" + "") != ("a" + ("hello" + ""))) {
                v48 := GenS3{F0: int8(v30), F1: "foo bar", F2: v30}
                _ = v48
            } else {
                if true {
                    switch (fi47 + 2147483647) {
                    case 42:
                        if true {
                            var v49 uint64 = uint64(v31)
                            v50 := GenS2{F0: (false || false)}
                            if true {
                                v51 := []int{fi47, -1, 337, (fi47 + (fi47 ^ 350))}
                                switch "" {
                                case "":
                                    fmt.Println(v1, v30)
                                case "hello":
                                    if false {
                                        var v52 int16 = 1
                                        _ = v52
                                    }
                                case "foo bar":
                                    if (!("" > "hello")) {
                                        switch "" {
                                        case "hello":
                                            v30 += v30
                                        }
                                        var v53 int64 = 42
                                        switch ("a" + "a") {
                                        case "a":
                                            var v54 int8 = -128
                                            _ = v54
                                        case "hello":
                                            var v55 int16 = 647
                                            _ = v55
                                        default:
                                            if false {
                                                for fi56 := 0; fi56 < 18; fi56++ {
                                                    switch "" {
                                                    case "foo bar":
                                                        v49 += v49
                                                    case "a":
                                                        var v57 int8 = 1
                                                        v49 += 979
                                                        _ = v57
                                                    case "\n\t":
                                                        var v58 string = "a"
                                                        _ = v58
                                                    default:
                                                        v3 += v3
                                                    }
                                                    v30 += v53
                                                    _ = fi56
                                                }
                                                if (true || (1e10 > 1e10)) {
                                                    for fi59 := 0; fi59 < 16; fi59++ {
                                                        v60 := GenS3{F0: -8, F1: "foo bar", F2: v30}
                                                        if len(v51) >= 2 {
                                                            v51 = v51[1:len(v51)]
                                                        }
                                                        _ = len(v2)
                                                        _ = fi59
                                                        _ = v60
                                                    }
                                                    v49 = 206
                                                    v3 = v3
                                                }
                                            } else {
                                                var v61 uint32 = genFn0("\n\t", v31, v49)
                                                _ = v61
                                            }
                                        }
                                        _ = v53
                                    } else {
                                        v62 := v2[("foo bar" + ("\n\t" + "a"))]
                                        var v63 int = 32767
                                        _ = v62
                                        _ = v63
                                    }
                                }
                                _ = v51
                            }
                            _ = v49
                            _ = v50
                        } else {
                            fmt.Println(v3)
                        }
                    case 32767:
                        v64 := []int{}
                        _ = v64
                    case 127:
                        for fi65 := 0; fi65 < 6; fi65++ {
                            v66 := []int{128, (fi65 << (((v3 % (((v3 ^ 768)) | 1))) & 63))}
                            v66 = append(v66, (-(fi65)))
                            var v67 float64 = 3.14159
                            _ = fi65
                            _ = v66
                            _ = v67
                        }
                    default:
                        v68 := v2["a"]
                        _ = v68
                    }
                    v69 := GenS6{F0: 0, F1: 4294967295}
                    _ = v69
                }
            }
            _ = fi47
        }
    }
    if (-546 != 128) {
        fmt.Println(v3, v31, v27)
        var v70 int8 = 26
        v71 := GenS7{F0: 850}
        _ = v70
        _ = v71
    } else {
        _ = len(v2)
        v30 += (-(v30))
    }
    v31 = (v31 << ((v3) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v30
    _ = v31
}
