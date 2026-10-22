package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8 }
type GenS2 struct { F0 uint64; F1 int64; F2 int8 }
type GenS3 struct { F0 uint32; F1 uint16; F2 bool }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int32; F1 bool; F2 string }
type GenS6 struct { F0 int16; F1 int32 }
type GenS7 struct { F0 uint }
func genFn0(p0 bool, p1 int) int8 {
    if false {
        if true {
            v5 := GenS1{F0: "a", F1: 1}
            var v6 uint32 = 2147483648
            _ = v5
            _ = v6
        } else {
            v7 := GenS2{F0: 1, F1: -908, F2: 100}
            _ = v7
        }
        var v8 float32 = float32(0.0)
        v9 := []int{(^(p1)), (p1 * (-(p1)))}
        _ = v8
        _ = v9
    } else {
        if true {
            var v10 uint = 1
            _ = v10
        } else {
            var v11 uint8 = 1
            v12 := GenS3{F0: 1, F1: 32767, F2: (p0 && p0)}
            _ = v11
            _ = v12
        }
        p1 = p1
        if (18446744073709551615 <= 289) {
            p1 = -9223372036854775808
            var v13 int = p1
            _ = v13
        }
    }
    if (1 > 409) {
        switch "\n\t" {
        case "\n\t":
            switch "foo bar" {
            case "\n\t":
                switch 984 {
                case 128:
                    var v14 int8 = 127
                    _ = v14
                case 4294967295:
                    switch 128 {
                    case 32768:
                        v15 := []int{-9223372036854775808, p1, (p1 >> ((0) & 63)), p1}
                        _ = v15
                    default:
                        v16 := GenS3{F0: 493, F1: 0, F2: p0}
                        _ = v16
                    }
                }
            case "hello":
                if (1e10 <= 0.1) {
                    if p0 {
                        p0 = true
                        for fi17 := 0; fi17 < 11; fi17++ {
                            v18 := []int{}
                            if (p0 && true) {
                                if true {
                                    for fi19 := 0; fi19 < 8; fi19++ {
                                        var v20 uint16 = 145
                                        _ = len(v18)
                                        if (32767 == 1000) {
                                            v21 := GenS3{F0: 2147483647, F1: v20, F2: (!(v20 != (^(v20))))}
                                            p0 = (42 != 159)
                                            _ = v21
                                        } else {
                                            if p0 {
                                                var v22 float64 = 0.1
                                                var v23 string = "hello"
                                                _ = v22
                                                _ = v23
                                            }
                                            p1 += -347
                                            switch fi17 {
                                            case 4294967295:
                                                v18 = append(v18, -547)
                                            }
                                        }
                                        _ = fi19
                                        _ = v20
                                    }
                                    v24 := GenS4{F0: 42}
                                    p1 = 0
                                    _ = v24
                                }
                                var v25 int = (p1 + p1)
                                _ = v25
                            }
                            for fi26 := 0; fi26 < 8; fi26++ {
                                if len(v18) > 0 {
                                    v18[0] = 32767
                                }
                                _ = fi26
                            }
                            _ = fi17
                            _ = v18
                        }
                    }
                    var v27 int16 = 1
                    _ = v27
                }
            case "a":
                v28 := []int{}
                _ = v28
            default:
                var v29 int16 = 1
                _ = v29
            }
        }
        p0 = p0
        p0 = (65535 >= 42)
    } else {
        v30 := GenS5{F0: 287, F1: p0, F2: "hello"}
        _ = v30
    }
    _ = p0
    _ = p1
    return -128
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[string]string)
    var v3 int32 = -762
    var v4 uint64 = 2
    switch 9223372036854775807 {
    case 127:
        v31 := []int{-938, 487, 874, -9223372036854775808}
        _ = v31
    case 9223372036854775807:
        v32 := []int{-638, -2147483648, -2147483648, 117}
        _ = v32
    case 128:
        for fi33 := 0; fi33 < 5; fi33++ {
            v34 := []int{fi33}
            _ = fi33
            _ = v34
        }
    default:
        for fi35 := 0; fi35 < 1; fi35++ {
            if false {
                var v36 bool = true
                fmt.Println(v3, v36, v4)
                _ = v36
            } else {
                switch fi35 {
                case -586:
                    switch ("foo bar" + ("hello" + "hello")) {
                    case "":
                        v4 += (^(v4))
                    case "hello":
                        _ = len(v1)
                    case "foo bar":
                        v3 -= -1
                    }
                case 4294967295:
                    v37 := []int{(fi35 ^ fi35), fi35, fi35}
                    _ = v37
                case 9223372036854775807:
                    var v38 uint64 = 0
                    _ = v38
                default:
                    var v39 uint = 42
                    _ = v39
                }
                v40 := GenS6{F0: 42, F1: -95}
                _ = v40
            }
            _ = fi35
        }
    }
    v41 := []int{32767, 32768, 695}
    var v42 uint16 = 0
    v42 = 1
    for fi43 := 0; fi43 < 14; fi43++ {
        v44 := GenS1{F0: "", F1: 2}
        v42 = 2
        _ = fi43
        _ = v44
    }
    if len(v41) >= 2 {
        v41 = v41[1:len(v41)]
    }
    if false {
        v41 = append(v41, 0)
        v3 = 0
    }
    var v45 float64 = 0.1
    fmt.Println(v42, v3)
    for fi46 := 0; fi46 < 7; fi46++ {
        switch (-(fi46)) {
        case -674:
            var v47 uint8 = 0
            _ = v47
        case 9223372036854775807:
            switch "foo bar" {
            case "\n\t":
                v48 := GenS7{F0: 18446744073709551615}
                _ = v48
            }
        case -817:
            var v49 int = fi46
            _ = v49
        default:
            v50 := v2[""]
            _ = v50
        }
        var v51 int32 = v3
        _ = fi46
        _ = v51
    }
    if (v42 < v42) {
        switch -841 {
        case 32768:
            if (0 > 42) {
                fmt.Println(v42, v3, v4)
            } else {
                var v52 string = (("foo bar" + ("\n\t" + "a")) + "hello")
                _ = len(v41)
                _ = v52
            }
        case -2147483648:
            var v53 int16 = 5
            _ = v53
        }
        var v54 int16 = -32768
        v45 = (v45 + v45)
        _ = v54
    } else {
        for fi55 := 0; fi55 < 16; fi55++ {
            for fi56 := 0; fi56 < 4; fi56++ {
                switch fi55 {
                case -1:
                    var v57 uint8 = 128
                    _ = v57
                case 128:
                    v4 += v4
                default:
                    v4 -= v4
                }
                _ = fi56
            }
            switch ("hello" + "a") {
            case "\n\t":
                v4 = (^(v4))
            default:
                fmt.Println(v3)
            }
            _ = fi55
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v41:", len(v41), v41)
    fmt.Println("v42:", v42)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v45
}
