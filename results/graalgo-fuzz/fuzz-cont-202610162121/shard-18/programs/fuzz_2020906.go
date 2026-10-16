package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int8; F1 int32; F2 int8 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 bool; F1 uint }
type GenS6 struct { F0 int }
type GenS7 struct { F0 int16; F1 uint16; F2 uint }
func genFn0(p0 uint, p1 uint32) int64 {
    switch ("\n\t" + "foo bar") {
    case "":
        fmt.Println(p0, p0)
    case "hello":
        v6 := GenS1{F0: ((true && (float32((-2.25 - 3.14159)) != (((0.1 + (-2.25 / (float32(float64((1.5 + 0.1))) / float32(float64(1e10))))) * -2.25) / (0.1 - ((1e10 - (3.14159 * (3.14159 - -2.25))) / float32(1.5)))))) && true), F1: 42}
        _ = v6
    default:
        fmt.Println(p0, p0)
    }
    var v7 uint8 = 2
    p1 = p1
    var v8 int8 = 127
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 18
}
func genFn1(p0 uint8, p1 uint, p2 uint) uint16 {
    v9 := []int{}
    fmt.Println(p2, p0)
    for fi10 := 0; fi10 < 14; fi10++ {
        v11 := GenS1{F0: (true && (false && (2 <= 0))), F1: 32767}
        var v12 float32 = -2.25
        _ = fi10
        _ = v11
        _ = v12
    }
    v9 = append(v9, 32768)
    p2 -= (^(p2))
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 0
}

func main() {
    v1 := make(map[int]bool)
    v2 := []string{"foo bar"}
    var v3 uint64 = 617
    var v4 uint32 = 65535
    var v5 uint16 = 255
    if (!true) {
        var v13 uint64 = 42
        switch ("" + "\n\t") {
        case "":
            if (0 >= -1) {
                delete(v1, 4294967295)
                var v14 bool = false
                for fi15 := 0; fi15 < 16; fi15++ {
                    for fi16 := 0; fi16 < 10; fi16++ {
                        switch "a" {
                        case "a":
                            if (((3.14159 - -2.25) / 0.1) == -2.25) {
                                if true {
                                    fmt.Println(v4, v13, v4)
                                    switch (fi15 * (fi15 * (fi16 | fi16))) {
                                    case -988:
                                        v3 += v3
                                        v17 := GenS1{F0: v14, F1: (v5 % (((v5 * (v5 + (^(v5))))) | 1))}
                                        var v18 uint16 = (v5 & uint16(v4))
                                        v19 := GenS2{F0: fi16}
                                        fmt.Println(v4)
                                        v3 += v13
                                        _ = v17
                                        _ = v18
                                        _ = v19
                                    case 32768:
                                        v20 := GenS2{F0: (fi16 - (fi15 + (fi15 / ((32768) | 1))))}
                                        _ = v20
                                    case 4294967295:
                                        var v21 float32 = ((((1e10 - 1e10) + 0.0) / 1.5) / 1.0)
                                        _ = v21
                                    default:
                                        v13 = (v13 % ((4294967295) | 1))
                                    }
                                    v14 = v14
                                } else {
                                    var v22 float32 = 0.0
                                    if (1.5 != 3.14159) {
                                        v3 = (v13 & (v13 - 1))
                                    } else {
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        v23 := v1[(fi15 << ((799) & 63))]
                                        v4 = 2147483648
                                        var v24 int16 = -342
                                        _ = v23
                                        _ = v24
                                    }
                                    _ = v22
                                }
                            } else {
                                v14 = (v4 < v4)
                                if ((v3 * (v13 % ((v13) | 1))) > v13) {
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    if len(v2) > 0 {
                                        v2[0] = "a"
                                    }
                                    v4 -= (v4 << ((18446744073709551615) & 63))
                                }
                                v25 := GenS1{F0: v14, F1: genFn1(2, 0, 303)}
                                _ = v25
                            }
                        case "":
                            if len(v2) > 0 {
                                v2[0] = "\n\t"
                            }
                        }
                        _ = fi16
                    }
                    var v26 float32 = 3.14159
                    fmt.Println(v26)
                    _ = fi15
                    _ = v26
                }
                _ = v14
            }
        }
        v27 := GenS3{F0: int8(v4), F1: -2147483648, F2: 127}
        _ = v13
        _ = v27
    }
    var v28 uint8 = 0
    switch -2147483648 {
    case 4294967295:
        if (128 > v28) {
            v4 += v4
        } else {
            for fi29 := 0; fi29 < 17; fi29++ {
                if false {
                    var v30 string = ("" + ("foo bar" + ("foo bar" + "foo bar")))
                    if false {
                        if (v4 >= 457) {
                            fmt.Println(fi29, v30)
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        }
                    }
                    _ = v30
                }
                _ = fi29
            }
            v3 = 42
        }
    default:
        fmt.Println(v3)
    }
    var v31 int = 127
    delete(v1, (v31 / (((v31 << ((2) & 63))) | 1)))
    v32 := GenS4{F0: uint(v4)}
    v33 := GenS5{F0: true, F1: 456}
    if false {
        v34 := GenS1{F0: (!false), F1: v5}
        _ = v34
    }
    if true {
        for fi35 := 0; fi35 < 2; fi35++ {
            switch "" {
            case "":
                fmt.Println(fi35, v31, v3)
            case "foo bar":
                var v36 uint32 = (v4 + (v4 * v4))
                _ = v36
            case "a":
                var v37 bool = ("hello" < ("foo bar" + ""))
                _ = v37
            default:
                v5 += 42
            }
            v38 := GenS6{F0: 0}
            fmt.Println(v5, v5)
            _ = fi35
            _ = v38
        }
        if len(v2) > 0 {
            v2[0] = "a"
        }
    } else {
        if true {
            v4 = (v4 % (((v4 | 1)) | 1))
        }
        if ((v3 >= 2) && true) {
            fmt.Println(v4)
            for fi39 := 0; fi39 < 3; fi39++ {
                if true {
                    v40 := GenS7{F0: 32767, F1: genFn1((v28 - v28), uint(v28), 18446744073709551615), F2: 4294967295}
                    switch 32767 {
                    case 300:
                        var v41 float32 = 1e10
                        _ = v41
                    case 9223372036854775807:
                        var v42 uint16 = (v5 + 65535)
                        _ = v42
                    default:
                        fmt.Println(v31)
                    }
                    var v43 string = "hello"
                    _ = v40
                    _ = v43
                }
                for fi44 := 0; fi44 < 17; fi44++ {
                    v3 -= (^(v3))
                    _ = fi44
                }
                _ = fi39
            }
            v5 = (^(v5))
        } else {
            if len(v2) > 0 {
                v2[0] = "foo bar"
            }
        }
        if true {
            var v45 int16 = int16(v28)
            _ = v45
        }
    }
    if true {
        fmt.Println(v4, v4)
    }
    var v46 int8 = -34
    if false {
        var v47 int64 = 695
        var v48 uint8 = (v28 | v28)
        _ = v47
        _ = v48
    } else {
        fmt.Println(v31)
    }
    v28 = 127
    v1[v31] = true
    for fi49 := 0; fi49 < 1; fi49++ {
        if ("hello" >= (("" + "a") + ("a" + ("" + "hello")))) {
            v2 = append(v2, "")
        }
        _ = fi49
    }
    fmt.Println(v31, v28, v28)
    v50 := GenS1{F0: (!false), F1: uint16(v4)}
    v3 = v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v46
    _ = v5
    _ = v50
}
