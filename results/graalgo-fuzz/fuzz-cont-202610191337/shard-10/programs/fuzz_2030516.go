package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }
type GenS2 struct { F0 int16; F1 uint8 }
type GenS3 struct { F0 int64; F1 int64 }
type GenS4 struct { F0 uint8; F1 int64 }
type GenS5 struct { F0 uint16 }
func genFn0(p0 uint64) int32 {
    p0 = (p0 & p0)
    p0 -= (^(p0))
    var v4 uint = 65535
    var v5 float64 = -2.25
    _ = p0
    _ = v4
    _ = v5
    return 1
}
func genFn1(p0 string, p1 uint16, p2 bool) uint16 {
    v6 := []int{int(p1), 127}
    if true {
        for fi7 := 0; fi7 < 10; fi7++ {
            var v8 int = fi7
            var v9 int8 = 100
            _ = fi7
            _ = v8
            _ = v9
        }
        v10 := GenS1{F0: 9223372036854775807, F1: -973}
        var v11 uint64 = 155
        _ = v10
        _ = v11
    }
    fmt.Println(p2, p1)
    var v12 uint32 = 65535
    fmt.Println(v12)
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v6
    return 1
}

func main() {
    var v1 int8 = 127
    v2 := make(map[int]uint32)
    v3 := make(map[string]uint32)
    v13 := GenS1{F0: -436, F1: 1}
    v14 := GenS2{F0: 42, F1: 192}
    fmt.Println(v1)
    fmt.Println(v1)
    v15 := []int{}
    switch "foo bar" {
    case "":
        var v16 string = "\n\t"
        _ = v16
    case "\n\t":
        var v17 int64 = 720
        _ = v17
    case "hello":
        delete(v2, 604)
    default:
        _ = len(v15)
    }
    if (127 > 14) {
        var v18 int = -1
        var v19 bool = (v1 >= (v1 | 127))
        _ = v18
        _ = v19
    } else {
        var v20 uint = 2
        switch 0 {
        case -9223372036854775808:
            if false {
                v21 := GenS3{F0: 0, F1: 0}
                if (2147483647 == 4294967295) {
                    for fi22 := 0; fi22 < 19; fi22++ {
                        switch (-(fi22)) {
                        case 32767:
                            var v23 uint8 = 128
                            _ = v23
                        case -2147483648:
                            if (false && (2147483647 < 2147483648)) {
                                var v24 int16 = 983
                                for fi25 := 0; fi25 < 7; fi25++ {
                                    v1 -= (v1 - 57)
                                    fmt.Println(v20, v24, fi22)
                                    v26 := GenS2{F0: v24, F1: 127}
                                    _ = fi25
                                    _ = v26
                                }
                                _ = v24
                            }
                        case 4294967295:
                            v20 = v20
                        }
                        var v27 int64 = -720
                        for fi28 := 0; fi28 < 18; fi28++ {
                            v29 := GenS2{F0: 1, F1: 1}
                            fmt.Println(fi22)
                            v15 = append(v15, (fi28 % (((fi28 << ((4294967295) & 63))) | 1)))
                            _ = fi28
                            _ = v29
                        }
                        _ = fi22
                        _ = v27
                    }
                    v30 := GenS2{F0: 1000, F1: 127}
                    _ = v30
                } else {
                    if false {
                        if (!false) {
                            v31 := GenS4{F0: 154, F1: 1}
                            for fi32 := 0; fi32 < 5; fi32++ {
                                if false {
                                    fmt.Println(v20, v1, v1)
                                } else {
                                    fmt.Println(v1, v20, fi32)
                                }
                                delete(v3, "hello")
                                _ = fi32
                            }
                            if (!((!(false && (1 == uint64(v20)))) && false)) {
                                var v33 float32 = 1.5
                                _ = v33
                            }
                            _ = v31
                        } else {
                            var v34 uint = 0
                            _ = v34
                        }
                    }
                    v35 := GenS3{F0: -9223372036854775808, F1: -9223372036854775808}
                    v15 = append(v15, 465)
                    _ = v35
                }
                _ = v21
            } else {
                for fi36 := 0; fi36 < 5; fi36++ {
                    var v37 int16 = 0
                    for fi38 := 0; fi38 < 14; fi38++ {
                        var v39 uint64 = 689
                        var v40 int32 = 1
                        v20 += 1
                        _ = fi38
                        _ = v39
                        _ = v40
                    }
                    if (2 <= 65535) {
                        for fi41 := 0; fi41 < 16; fi41++ {
                            if ((42 <= 18446744073709551615) && false) {
                                switch fi41 {
                                case -1:
                                    var v42 uint16 = 760
                                    _ = v42
                                default:
                                    v37 += v37
                                }
                                if ((v1 * (v1 << ((42) & 63))) == int8(fi41)) {
                                    if (-2.25 == -2.25) {
                                        var v43 uint16 = 255
                                        switch (fi41 / (((fi41 & 32768)) | 1)) {
                                        case -9223372036854775808:
                                            v44 := v2[(fi36 >> ((v20) & 63))]
                                            _ = v44
                                        case 128:
                                            fmt.Println(fi41, v43, v37)
                                        }
                                        var v45 uint64 = 63
                                        var v46 uint8 = 32
                                        _ = v43
                                        _ = v45
                                        _ = v46
                                    } else {
                                        var v47 uint32 = 1
                                        var v48 int64 = -1
                                        v49 := v3[""]
                                        _ = v47
                                        _ = v48
                                        _ = v49
                                    }
                                }
                                v37 = v37
                            }
                            var v50 int8 = -128
                            v50 = (v1 / ((v1) | 1))
                            _ = fi41
                            _ = v50
                        }
                        v1 -= (v1 ^ v1)
                        _ = len(v15)
                    } else {
                        fmt.Println(v20)
                        v1 += (v1 + (v1 % (((^(v1))) | 1)))
                        _ = len(v15)
                    }
                    _ = fi36
                    _ = v37
                }
                var v51 uint8 = 0
                for fi52 := 0; fi52 < 1; fi52++ {
                    for fi53 := 0; fi53 < 12; fi53++ {
                        if true {
                            var v54 string = "hello"
                            _ = v54
                        }
                        var v55 int8 = v1
                        _ = fi53
                        _ = v55
                    }
                    _ = fi52
                }
                _ = v51
            }
        }
        _ = v20
    }
    if false {
        for fi56 := 0; fi56 < 7; fi56++ {
            var v57 float64 = float64((3.14159 * 3.14159))
            _ = fi56
            _ = v57
        }
        var v58 float64 = -2.25
        _ = v58
    }
    var v59 bool = (!false)
    for fi60 := 0; fi60 < 6; fi60++ {
        v1 = (v1 - -86)
        _ = fi60
    }
    var v61 int64 = -9223372036854775808
    for fi62 := 0; fi62 < 10; fi62++ {
        for fi63 := 0; fi63 < 18; fi63++ {
            v64 := GenS1{F0: (fi63 / ((fi63) | 1)), F1: 397}
            v65 := GenS5{F0: genFn1("foo bar", 131, ("foo bar" == ("a" + "\n\t")))}
            v59 = (1.5 != 1e10)
            _ = fi63
            _ = v64
            _ = v65
        }
        _ = fi62
    }
    var v66 int64 = 42
    if (32767 >= 2147483647) {
        v59 = v59
        if len(v15) > 0 {
            v15[0] = -1
        }
    } else {
        v67 := v3[("foo bar" + "\n\t")]
        _ = v67
    }
    if v59 {
        fmt.Println(v66)
        var v68 bool = v59
        _ = v68
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v59:", v59)
    fmt.Println("v61:", v61)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v59
    _ = v61
    _ = v66
}
