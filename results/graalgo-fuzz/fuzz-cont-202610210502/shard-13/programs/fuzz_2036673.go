package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int64; F1 int }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint16; F1 string }
func genFn0(p0 int8, p1 uint) int32 {
    var v5 bool = false
    for fi6 := 0; fi6 < 10; fi6++ {
        fmt.Println(v5, fi6, p1)
        var v7 uint8 = 0
        _ = fi6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = v5
    return 0
}
func genFn1(p0 uint, p1 string) string {
    switch p1 {
    case "a":
        fmt.Println(p0, p1)
    case "":
        v8 := GenS1{F0: 0, F1: 42}
        _ = v8
    case "foo bar":
        v9 := GenS2{F0: 2}
        _ = v9
    }
    p1 = ((p1 + (p1 + p1)) + p1)
    _ = p0
    _ = p1
    return (p1 + p1)
}

func main() {
    var v1 uint16 = 65535
    var v2 int32 = 483
    var v3 float64 = 0.1
    var v4 uint32 = 42
    v1 = v1
    v2 = 0
    switch -2147483648 {
    case -9223372036854775808:
        if true {
            v1 += (v1 | (v1 ^ 2))
            v10 := GenS3{F0: 934, F1: 4294967295}
            _ = v10
        } else {
            var v11 uint16 = v1
            for fi12 := 0; fi12 < 10; fi12++ {
                var v13 uint = 923
                _ = fi12
                _ = v13
            }
            _ = v11
        }
    default:
        v14 := []int{782}
        _ = v14
    }
    fmt.Println(v3, v4)
    fmt.Println(v4, v3)
    if false {
        var v15 int = 1
        v15 = (v15 >> ((2) & 63))
        if false {
            var v16 int8 = 100
            fmt.Println(v16)
            if (((0.1 * (((1e10 - ((-2.25 - 3.14159) / (1e10 * float32(float64(1.5))))) - (((((0.1 + 1.5) - (-2.25 / 1e10)) + (((3.14159 + -2.25) * 1.5) - ((1.5 / 1.0) + (0.0 + -2.25)))) * -2.25) + -2.25)) * ((1e10 / 1e10) - (0.0 * float32((v3 + v3)))))) >= (0.0 + 0.1)) && ((-132 == 42) || false)) {
                v17 := GenS4{F0: -127}
                _ = v17
            } else {
                for fi18 := 0; fi18 < 13; fi18++ {
                    v16 = int8(v4)
                    var v19 int = v15
                    switch ("a" + "\n\t") {
                    case "a":
                        if true {
                            v4 = v4
                            var v20 bool = (true || true)
                            _ = v20
                        } else {
                            var v21 uint64 = 42
                            var v22 float64 = float64(-2.25)
                            _ = v21
                            _ = v22
                        }
                    case "hello":
                        var v23 float64 = 0.0
                        _ = v23
                    case "\n\t":
                        var v24 int64 = 1
                        _ = v24
                    default:
                        v25 := GenS1{F0: v19, F1: (v2 | (-(v2)))}
                        _ = v25
                    }
                    _ = fi18
                    _ = v19
                }
                var v26 int32 = v2
                var v27 int = (v15 % ((v15) | 1))
                _ = v26
                _ = v27
            }
            _ = v16
        }
        _ = v15
    } else {
        v28 := GenS1{F0: 0, F1: (v2 % ((genFn0(127, 255)) | 1))}
        var v29 uint16 = v1
        _ = v28
        _ = v29
    }
    var v30 uint64 = 2
    if (0.0 == (((1e10 / 1.5) + 1e10) * 3.14159)) {
        for fi31 := 0; fi31 < 3; fi31++ {
            if (!true) {
                var v32 int32 = 42
                _ = v32
            } else {
                fmt.Println(fi31)
                var v33 float64 = (0.1 - v3)
                _ = v33
            }
            fmt.Println(v30)
            for fi34 := 0; fi34 < 15; fi34++ {
                if (true || (255 == 0)) {
                    var v35 uint64 = (v30 >> ((uint(v30)) & 63))
                    v4 += 37
                    _ = v35
                } else {
                    var v36 int64 = -1
                    v3 += v3
                    switch "\n\t" {
                    case "\n\t":
                        var v37 uint8 = 2
                        _ = v37
                    case "":
                        v38 := []int{}
                        _ = v38
                    case "a":
                        v39 := GenS3{F0: -1, F1: (fi31 ^ -156)}
                        _ = v39
                    default:
                        var v40 int16 = 0
                        _ = v40
                    }
                    _ = v36
                }
                if ((94 < 255) && ((v1 << ((65535) & 63)) > v1)) {
                    var v41 int32 = v2
                    v42 := []int{(fi34 * (fi34 % (((fi31 ^ int(v30))) | 1))), fi31, -867}
                    if true {
                        var v43 int = (fi31 * int(v1))
                        v44 := GenS2{F0: 255}
                        v30 += v30
                        if ((0.0 - 0.1) == 0.0) {
                            var v45 uint8 = 42
                            if false {
                                var v46 float32 = ((1.5 + 1e10) + 0.1)
                                _ = v46
                            } else {
                                v41 -= genFn0(127, 104)
                                if false {
                                    _ = len(v42)
                                    v42 = append(v42, fi34)
                                } else {
                                    fmt.Println(v43)
                                    _ = len(v42)
                                    v1 = uint16(v4)
                                }
                                v1 += 2
                            }
                            v3 = (((float64(1e10) - 3.14159) / 3.14159) - float64(float32(v3)))
                            _ = v45
                        }
                        _ = v43
                        _ = v44
                    } else {
                        var v47 int64 = 399
                        v3 -= ((((float64(float32(((0.1 / 0.1) * v3))) + 0.0) - v3) * 1e10) / 0.1)
                        v42 = append(v42, fi34)
                        _ = v47
                    }
                    _ = v41
                    _ = v42
                }
                _ = fi34
            }
            _ = fi31
        }
        var v48 int32 = int32(v1)
        for fi49 := 0; fi49 < 4; fi49++ {
            fmt.Println(fi49)
            v50 := GenS5{F0: 9223372036854775807}
            _ = fi49
            _ = v50
        }
        _ = v48
    }
    var v51 int64 = 9223372036854775807
    v52 := GenS6{F0: 32767, F1: "\n\t"}
    fmt.Println(v3, v1, v1)
    for fi53 := 0; fi53 < 11; fi53++ {
        var v54 uint64 = (v30 ^ 18446744073709551615)
        _ = fi53
        _ = v54
    }
    var v55 float32 = 3.14159
    if true {
        v56 := GenS3{F0: (v51 << ((1) & 63)), F1: -904}
        v57 := []int{-1, 42, -2147483648}
        _ = v56
        _ = v57
    }
    v58 := []int{-708}
    _ = len(v58)
    for fi59 := 0; fi59 < 7; fi59++ {
        _ = len(v58)
        var v60 int = fi59
        for fi61 := 0; fi61 < 11; fi61++ {
            var v62 float32 = v55
            v63 := GenS3{F0: (v51 * v51), F1: 13}
            _ = fi61
            _ = v62
            _ = v63
        }
        _ = fi59
        _ = v60
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v51:", v51)
    fmt.Println("v55:", v55)
    fmt.Println("slicelen v58:", len(v58), v58)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v51
    _ = v52
    _ = v55
    _ = v58
}
