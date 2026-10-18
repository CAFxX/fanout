package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 string }
type GenS2 struct { F0 int; F1 int8 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int32; F1 int8 }
type GenS5 struct { F0 uint32; F1 bool; F2 int32 }
func genFn0(p0 uint8, p1 uint32) uint64 {
    var v4 int32 = -1
    for fi5 := 0; fi5 < 13; fi5++ {
        var v6 int16 = 1
        _ = fi5
        _ = v6
    }
    var v7 uint = 4294967295
    v8 := []int{128, -127, -9223372036854775808}
    _ = p0
    _ = p1
    _ = v4
    _ = v7
    _ = v8
    return 4294967295
}
func genFn1(p0 int32, p1 uint64, p2 int32) uint64 {
    switch -839 {
    case 0:
        v9 := []int{901, 2147483647, 9223372036854775807, 7}
        _ = v9
    case -2147483648:
        var v10 float64 = (1.5 * 3.14159)
        _ = v10
    }
    fmt.Println(p0, p2)
    v11 := GenS1{F0: 967, F1: 127, F2: ("a" + ((("foo bar" + "\n\t") + ("\n\t" + "\n\t")) + ""))}
    p1 -= (p1 - p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return p1
}

func main() {
    v1 := []uint32{2147483648, 65535, 963, 4294967295}
    v2 := make(map[string]string)
    var v3 string = "\n\t"
    for fi12 := 0; fi12 < 4; fi12++ {
        if (!(0.0 > 3.14159)) {
            v13 := GenS1{F0: 255, F1: 116, F2: ("foo bar" + ("" + v3))}
            _ = v13
        }
        switch v3 {
        case "":
            v1 = append(v1, 211)
        case "foo bar":
            v3 = v3
        default:
            if false {
                var v14 uint16 = 0
                fmt.Println(fi12, v3)
                _ = v14
            }
        }
        v3 = v3
        _ = fi12
    }
    var v15 float32 = 1.5
    for fi16 := 0; fi16 < 3; fi16++ {
        if true {
            switch ((v3 + "") + v3) {
            case "a":
                var v17 uint16 = 140
                _ = v17
            case "foo bar":
                if false {
                    v18 := GenS2{F0: fi16, F1: -128}
                    _ = v18
                } else {
                    if false {
                        v15 -= v15
                        for fi19 := 0; fi19 < 11; fi19++ {
                            fmt.Println(fi16, v15, v3)
                            v20 := GenS2{F0: fi16, F1: 100}
                            _ = fi19
                            _ = v20
                        }
                        v15 += (v15 - v15)
                    }
                }
            case "":
                var v21 uint64 = 4294967295
                _ = v21
            default:
                v22 := GenS3{F0: ((v3 + "") + ("" + v3))}
                _ = v22
            }
        }
        switch ("" + "\n\t") {
        case "foo bar":
            v15 += float32(((1.5 + 0.1) + 0.0))
        default:
            v15 = (v15 + v15)
        }
        _ = fi16
    }
    if (0 >= 32767) {
        var v23 string = v3
        _ = v23
    }
    v24 := GenS2{F0: -402, F1: 35}
    fmt.Println(v15)
    for fi25 := 0; fi25 < 2; fi25++ {
        v26 := GenS4{F0: 1, F1: 95}
        _ = fi25
        _ = v26
    }
    fmt.Println(v15, v3)
    var v27 uint = 2
    fmt.Println(v3, v15)
    switch "foo bar" {
    case "\n\t":
        _ = len(v2)
    case "hello":
        var v28 int64 = -1
        _ = v28
    default:
        for fi29 := 0; fi29 < 19; fi29++ {
            for fi30 := 0; fi30 < 2; fi30++ {
                var v31 int16 = -1
                _ = fi30
                _ = v31
            }
            _ = fi29
        }
    }
    switch -2147483648 {
    case -11:
        if true {
            v15 -= 0.1
            delete(v2, "")
            fmt.Println(v27)
        } else {
            switch v3 {
            case "\n\t":
                for fi32 := 0; fi32 < 19; fi32++ {
                    v3 = v3
                    if (!true) {
                        switch (fi32 + (fi32 << (((v27 * v27)) & 63))) {
                        case 4294967295:
                            var v33 int32 = int32(fi32)
                            _ = v33
                        case 127:
                            v1 = append(v1, 2)
                        case -791:
                            if true {
                                switch (("hello" + v3) + v3) {
                                case "a":
                                    for fi34 := 0; fi34 < 5; fi34++ {
                                        var v35 uint = v27
                                        fmt.Println(v35)
                                        var v36 int16 = int16(fi32)
                                        _ = fi34
                                        _ = v35
                                        _ = v36
                                    }
                                case "":
                                    fmt.Println(v15)
                                case "foo bar":
                                    switch fi32 {
                                    case -9223372036854775808:
                                        var v37 int64 = 42
                                        _ = v37
                                    case 0:
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                    case 32768:
                                        if (1000 > 42) {
                                            v38 := v2[v3]
                                            for fi39 := 0; fi39 < 3; fi39++ {
                                                v3 = ("a" + v38)
                                                fmt.Println(fi32, v27)
                                                _ = fi39
                                            }
                                            _ = v38
                                        } else {
                                            v15 += v15
                                        }
                                    default:
                                        v27 += 592
                                    }
                                }
                            }
                        default:
                            var v40 int = -2147483648
                            _ = v40
                        }
                    } else {
                        var v41 uint32 = 51
                        fmt.Println(v41, v41)
                        fmt.Println(v3, v3)
                        _ = v41
                    }
                    fmt.Println(fi32, fi32)
                    _ = fi32
                }
            }
        }
    case 32768:
        if (4294967295 < -1) {
            var v42 uint64 = 715
            _ = v42
        }
    case -527:
        v1 = append(v1, 2)
    }
    var v43 string = ("" + "hello")
    if (true && true) {
        for fi44 := 0; fi44 < 17; fi44++ {
            var v45 int = fi44
            v45 -= (fi44 / ((fi44) | 1))
            if (((3.14159 + float32((-2.25 / 1.5))) * v15) < 0.0) {
                for fi46 := 0; fi46 < 3; fi46++ {
                    fmt.Println(v43, v3, v15)
                    var v47 int64 = 1
                    fmt.Println(v47, v27)
                    _ = fi46
                    _ = v47
                }
                for fi48 := 0; fi48 < 9; fi48++ {
                    var v49 uint32 = 2147483648
                    _ = fi48
                    _ = v49
                }
            }
            _ = fi44
            _ = v45
        }
        fmt.Println(v27, v27)
        if false {
            delete(v2, (v43 + "\n\t"))
            var v50 uint32 = 310
            var v51 uint64 = 813
            _ = v50
            _ = v51
        } else {
            v3 = v3
            var v52 int64 = 1
            v53 := GenS2{F0: 0, F1: -1}
            _ = v52
            _ = v53
        }
    }
    v3 = (v3 + v43)
    v43 = ((((v3 + "\n\t") + "\n\t") + "hello") + ((v3 + ((v43 + v43) + v43)) + ""))
    if true {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    } else {
        v54 := GenS5{F0: 328, F1: (!((v43 >= "foo bar") && false)), F2: -102}
        _ = v54
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v15
    _ = v2
    _ = v24
    _ = v27
    _ = v3
    _ = v43
}
