package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 uint16 }
type GenS2 struct { F0 int32; F1 int32 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 bool; F1 int32; F2 uint }

func main() {
    var v1 bool = false
    var v2 uint = 255
    v3 := make(map[string]uint32)
    var v4 int32 = -648
    var v5 uint8 = 2
    for fi6 := 0; fi6 < 1; fi6++ {
        var v7 uint8 = 1
        _ = fi6
        _ = v7
    }
    if v1 {
        var v8 uint8 = (v5 % (((v5 | (v5 >> ((4294967295) & 63)))) | 1))
        _ = v8
    } else {
        v9 := GenS1{F0: -32768, F1: -56, F2: 51}
        _ = v9
    }
    switch ("foo bar" + ("" + "foo bar")) {
    case "a":
        if false {
            if (false && (((!v1) && v1) && true)) {
                v10 := GenS1{F0: 1, F1: 0, F2: 80}
                var v11 int16 = 312
                _ = v10
                _ = v11
            } else {
                v5 += v5
            }
            v1 = ((114 < v5) || (272 <= 2))
            var v12 float64 = 1.5
            _ = v12
        } else {
            v5 -= v5
            var v13 float32 = 0.1
            _ = v13
        }
    case "\n\t":
        v1 = v1
    case "foo bar":
        v4 = v4
    default:
        var v14 uint16 = 633
        _ = v14
    }
    if v1 {
        if v1 {
            v4 -= v4
            v4 = (^(v4))
            if true {
                fmt.Println(v2, v2)
                if v1 {
                    var v15 uint16 = 42
                    v16 := []int{42, 867, 32767, 0}
                    switch "a" {
                    case "\n\t":
                        switch ("\n\t" + ("\n\t" + "foo bar")) {
                        case "\n\t":
                            for fi17 := 0; fi17 < 4; fi17++ {
                                if false {
                                    var v18 uint16 = (^(v15))
                                    var v19 float64 = 0.0
                                    _ = v18
                                    _ = v19
                                }
                                switch 32767 {
                                case 1:
                                    v15 = v15
                                }
                                _ = fi17
                            }
                        case "hello":
                            v20 := GenS1{F0: 32767, F1: 1, F2: v15}
                            _ = v20
                        case "foo bar":
                            for fi21 := 0; fi21 < 4; fi21++ {
                                for fi22 := 0; fi22 < 13; fi22++ {
                                    v23 := GenS1{F0: -220, F1: 72, F2: v15}
                                    if ((!(false && (-1 >= 100))) && v1) {
                                        v3[("\n\t" + ("\n\t" + "a"))] = 65535
                                        var v24 int32 = (v4 & 2147483647)
                                        var v25 uint = v2
                                        _ = v24
                                        _ = v25
                                    } else {
                                        v15 = v15
                                    }
                                    _ = fi22
                                    _ = v23
                                }
                                _ = fi21
                            }
                        default:
                            _ = len(v3)
                        }
                    default:
                        v26 := GenS2{F0: (v4 ^ (v4 + v4)), F1: (-(v4))}
                        _ = v26
                    }
                    _ = v15
                    _ = v16
                } else {
                    for fi27 := 0; fi27 < 15; fi27++ {
                        for fi28 := 0; fi28 < 15; fi28++ {
                            fmt.Println(fi27)
                            fmt.Println(v2, v2, v2)
                            _ = fi28
                        }
                        _ = fi27
                    }
                }
            }
        }
        v1 = (((v4 + (v4 * 285)) > v4) && (-1 < 1))
        v2 = v2
    }
    if v1 {
        v29 := []int{-667, -791, -1, 1}
        fmt.Println(v4, v5, v5)
        _ = v29
    }
    if false {
        for fi30 := 0; fi30 < 13; fi30++ {
            var v31 int8 = -128
            _ = fi30
            _ = v31
        }
    }
    for fi32 := 0; fi32 < 16; fi32++ {
        v33 := v3[""]
        v34 := GenS1{F0: 42, F1: 94, F2: 2}
        v1 = true
        _ = fi32
        _ = v33
        _ = v34
    }
    v2 += v2
    for fi35 := 0; fi35 < 1; fi35++ {
        fmt.Println(fi35)
        v36 := GenS2{F0: -1, F1: -2147483648}
        v1 = (!(9223372036854775807 > 242))
        _ = fi35
        _ = v36
    }
    var v37 int64 = -1
    var v38 uint64 = 4294967295
    for fi39 := 0; fi39 < 14; fi39++ {
        for fi40 := 0; fi40 < 3; fi40++ {
            v2 = v2
            v41 := GenS3{F0: 4294967295}
            switch fi39 {
            case 0:
                v42 := GenS4{F0: (v4 + (v4 * v4))}
                _ = v42
            case 4294967295:
                if false {
                    v43 := GenS3{F0: 11}
                    _ = v43
                } else {
                    v44 := GenS3{F0: 1}
                    if v1 {
                        if (1e10 == 0.1) {
                            var v45 int64 = -1
                            v46 := []int{(fi40 << ((v2) & 63)), 405, (fi39 | (fi40 % ((int(v5)) | 1)))}
                            _ = v45
                            _ = v46
                        }
                        if false {
                            if (v1 || v1) {
                                fmt.Println(fi40, fi39)
                                fmt.Println(v4, v1)
                                v1 = (72 < uint32(v2))
                            }
                        }
                        var v47 float64 = -2.25
                        var v48 int8 = 1
                        var v49 int = (fi39 - 4294967295)
                        _ = v47
                        _ = v48
                        _ = v49
                    } else {
                        fmt.Println(fi39)
                    }
                    var v50 uint16 = 32767
                    v51 := []int{(^(fi39)), 32767}
                    _ = v44
                    _ = v50
                    _ = v51
                }
            default:
                var v52 uint64 = v38
                _ = v52
            }
            _ = fi40
            _ = v41
        }
        switch fi39 {
        case -1:
            v53 := GenS3{F0: 2147483648}
            _ = v53
        case 32767:
            var v54 float32 = (0.1 + 1.5)
            _ = v54
        }
        _ = fi39
    }
    for fi55 := 0; fi55 < 3; fi55++ {
        v38 += (v38 - v38)
        switch (^(fi55)) {
        case 1:
            v56 := []int{9223372036854775807, (fi55 % (((fi55 << ((uint(v38)) & 63))) | 1)), fi55}
            _ = v56
        }
        v57 := GenS5{F0: (!(1 != 807)), F1: (-(v4)), F2: 18446744073709551615}
        _ = fi55
        _ = v57
    }
    v4 -= v4
    var v58 uint32 = 946
    v1 = (!((v5 + (v5 << (((v2 << ((v2) & 63))) & 63))) <= (v5 & (v5 % (((v5 ^ v5)) | 1)))))
    v59 := []int{0, 212, 9223372036854775807, -2147483648}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v58:", v58)
    fmt.Println("slicelen v59:", len(v59), v59)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v58
    _ = v59
}
