package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int; F1 uint32; F2 int64 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int8; F1 uint; F2 string }
type GenS5 struct { F0 int8; F1 int16 }
type GenS6 struct { F0 uint; F1 int; F2 int8 }
func genFn0(p0 int16, p1 int, p2 int64) string {
    switch p1 {
    case 32768:
        var v4 string = ("foo bar" + "foo bar")
        _ = v4
    default:
        for fi5 := 0; fi5 < 10; fi5++ {
            var v6 float64 = -2.25
            v7 := GenS1{F0: (4294967295 < 0)}
            _ = fi5
            _ = v6
            _ = v7
        }
    }
    p0 = p0
    if (1e10 <= 3.14159) {
        if ("" != "hello") {
            p0 -= 42
            p1 += (p1 >> ((42) & 63))
            var v8 uint32 = 2
            _ = v8
        } else {
            p0 = 309
            v9 := GenS2{F0: (p1 * p1), F1: 835, F2: (p2 ^ 0)}
            var v10 int16 = p0
            _ = v10
            _ = v9
        }
    } else {
        p1 = p1
        v11 := []int{445, 42, (p1 / (((p1 & p1)) | 1)), (p1 >> ((18446744073709551615) & 63))}
        _ = v11
    }
    v12 := []int{(p1 + (^(p1)))}
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return ("a" + "a")
}
func genFn1(p0 uint64, p1 int16) int32 {
    v13 := GenS1{F0: true}
    for fi14 := 0; fi14 < 18; fi14++ {
        var v15 uint16 = 65535
        _ = fi14
        _ = v15
    }
    p1 = p1
    _ = p0
    _ = p1
    _ = v13
    return 0
}

func main() {
    var v1 uint64 = 42
    v2 := make(map[string]bool)
    var v3 float32 = 1e10
    var v16 string = ("a" + "")
    if true {
        v3 = ((((float32(0.1) * v3) - (v3 + (0.0 / v3))) * v3) * (float32((float64(v3) + float64(3.14159))) - v3))
        v1 += v1
        fmt.Println(v1, v16, v3)
    } else {
        if false {
            if (162 > 171) {
                for fi17 := 0; fi17 < 4; fi17++ {
                    v18 := GenS3{F0: 200}
                    _ = fi17
                    _ = v18
                }
            }
        } else {
            v19 := GenS2{F0: 32767, F1: 1, F2: 9223372036854775807}
            _ = v19
        }
    }
    var v20 int16 = -32768
    if true {
        var v21 uint16 = 255
        v16 = "hello"
        if false {
            v22 := []int{-31, 0, -9223372036854775808}
            v23 := GenS3{F0: v21}
            _ = v22
            _ = v23
        } else {
            v21 = (v21 << ((4294967295) & 63))
            if true {
                v21 = 0
                if (false || true) {
                    v20 = v20
                    var v24 int8 = -103
                    _ = v24
                } else {
                    var v25 uint = 18446744073709551615
                    v26 := []int{1, int(v25), 4294967295}
                    var v27 uint16 = uint16(v1)
                    _ = v25
                    _ = v26
                    _ = v27
                }
                fmt.Println(v1, v16, v3)
            }
            v20 -= (v20 ^ (v20 >> ((65535) & 63)))
        }
        _ = v21
    }
    switch ("foo bar" + ("\n\t" + (v16 + "a"))) {
    case "foo bar":
        switch "hello" {
        case "foo bar":
            v28 := GenS4{F0: 42, F1: 42, F2: v16}
            _ = v28
        case "hello":
            if (!(127 <= 0)) {
                var v29 int8 = 42
                v3 = v3
                _ = v29
            }
        case "\n\t":
            var v30 uint16 = uint16(v1)
            _ = v30
        default:
            if (0 == (v20 >> ((411) & 63))) {
                var v31 uint = 2
                switch -2147483648 {
                case 4294967295:
                    if true {
                        v32 := GenS3{F0: 605}
                        v33 := GenS4{F0: -73, F1: v31, F2: "hello"}
                        _ = v32
                        _ = v33
                    } else {
                        v34 := GenS3{F0: 77}
                        if true {
                            for fi35 := 0; fi35 < 9; fi35++ {
                                v36 := []int{fi35, (fi35 % ((2147483647) | 1)), (fi35 / ((fi35) | 1)), (fi35 + fi35)}
                                v37 := GenS2{F0: 9223372036854775807, F1: 962, F2: 1}
                                for fi38 := 0; fi38 < 8; fi38++ {
                                    var v39 float64 = 1e10
                                    _ = fi38
                                    _ = v39
                                }
                                _ = fi35
                                _ = v36
                                _ = v37
                            }
                            var v40 uint = v31
                            _ = v40
                        } else {
                            var v41 uint16 = uint16(v31)
                            v42 := GenS1{F0: false}
                            _ = v41
                            _ = v42
                        }
                        var v43 float64 = -2.25
                        _ = v34
                        _ = v43
                    }
                default:
                    v44 := []int{-366, 32768}
                    _ = v44
                }
                _ = v31
            } else {
                v3 = 3.14159
            }
        }
    case "a":
        var v45 string = (v16 + (v16 + "\n\t"))
        _ = v45
    case "":
        fmt.Println(v20)
    default:
        for fi46 := 0; fi46 < 15; fi46++ {
            for fi47 := 0; fi47 < 15; fi47++ {
                var v48 uint = 950
                _ = fi47
                _ = v48
            }
            v49 := []int{42, (-(fi46)), (fi46 + 300), fi46}
            _ = fi46
            _ = v49
        }
    }
    v1 += 534
    if false {
        if (!(v3 > (((v3 + v3) * ((1.5 - (v3 + 1e10)) / v3)) + float32((((((0.0 * (0.1 / -2.25)) * 0.1) * (1.5 * (3.14159 - (float64(((float32(1.5) / v3) - (v3 / v3))) + -2.25)))) - 1e10) + 1e10))))) {
            v16 = (v16 + "hello")
            fmt.Println(v20)
            switch 217 {
            case 514:
                if false {
                    var v50 float32 = v3
                    for fi51 := 0; fi51 < 16; fi51++ {
                        v2[("a" + v16)] = false
                        _ = fi51
                    }
                    v16 = v16
                    _ = v50
                }
            case 32767:
                var v52 uint16 = 0
                _ = v52
            default:
                for fi53 := 0; fi53 < 1; fi53++ {
                    var v54 float32 = float32(1.5)
                    v55 := GenS5{F0: 127, F1: v20}
                    fmt.Println(v3)
                    _ = fi53
                    _ = v54
                    _ = v55
                }
            }
        } else {
            if true {
                v56 := GenS1{F0: false}
                var v57 string = v16
                _ = v56
                _ = v57
            } else {
                v3 -= 0.0
                for fi58 := 0; fi58 < 12; fi58++ {
                    var v59 int64 = -501
                    _ = fi58
                    _ = v59
                }
            }
            v60 := GenS6{F0: 1, F1: 128, F2: 1}
            v16 = v16
            _ = v60
        }
        var v61 int8 = 42
        _ = len(v2)
        _ = v61
    }
    v62 := []int{-9223372036854775808}
    if false {
        fmt.Println(v20, v16, v16)
        var v63 int32 = genFn1(v1, v20)
        _ = v63
    } else {
        if len(v62) >= 2 {
            v62 = v62[1:len(v62)]
        }
        var v64 uint = 0
        _ = v64
    }
    v1 = 11
    if (false || (9223372036854775807 == 0)) {
        fmt.Println(v16, v20, v3)
        if (true && (1e10 >= 0.0)) {
            v65 := GenS5{F0: 1, F1: v20}
            v3 = 3.14159
            fmt.Println(v1)
            _ = v65
        } else {
            v1 += 517
            v20 -= (v20 >> ((313) & 63))
        }
    } else {
        var v66 int32 = 760
        v66 += (v66 << ((65535) & 63))
        fmt.Println(v66, v66, v3)
        _ = v66
    }
    var v67 int64 = 1
    var v68 uint64 = (v1 * 548)
    var v69 float32 = v3
    v20 = (v20 % ((v20) | 1))
    var v70 int32 = 2147483647
    switch (v16 + "\n\t") {
    case "hello":
        for fi71 := 0; fi71 < 15; fi71++ {
            v69 -= (float32(0.1) / v3)
            if true {
                fmt.Println(v20)
            }
            var v72 uint = 626
            _ = fi71
            _ = v72
        }
    case "a":
        v1 = 359
    case "":
        v69 -= 3.14159
    default:
        var v73 int = 117
        _ = v73
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v62:", len(v62), v62)
    fmt.Println("v67:", v67)
    fmt.Println("v68:", v68)
    fmt.Println("v69:", v69)
    fmt.Println("v70:", v70)
    _ = v1
    _ = v16
    _ = v2
    _ = v20
    _ = v3
    _ = v62
    _ = v67
    _ = v68
    _ = v69
    _ = v70
}
