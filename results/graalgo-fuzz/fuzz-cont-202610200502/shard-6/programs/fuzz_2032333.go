package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16; F2 int16 }
type GenS2 struct { F0 uint8; F1 int8 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 string; F1 int64 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int64; F1 uint16 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 uint16; F1 int16 }
type GenS9 struct { F0 int16; F1 int8; F2 uint32 }
func genFn0(p0 uint) int16 {
    if true {
        for fi4 := 0; fi4 < 12; fi4++ {
            switch ("hello" + "") {
            case "a":
                v5 := []int{fi4, fi4, (fi4 ^ fi4)}
                _ = v5
            case "hello":
                p0 += 255
            case "":
                var v6 int16 = -1
                _ = v6
            }
            _ = fi4
        }
        if false {
            var v7 uint64 = 1
            switch "foo bar" {
            case "":
                p0 = p0
            default:
                fmt.Println(v7, p0)
            }
            _ = v7
        } else {
            fmt.Println(p0)
        }
        var v8 float64 = (0.0 + (-2.25 / ((-2.25 - 0.1) / (1e10 + 1.5))))
        _ = v8
    }
    p0 -= 0
    p0 = p0
    var v9 float32 = float32(float64(0.1))
    _ = p0
    _ = v9
    return int16(p0)
}
func genFn1(p0 int8, p1 int8) int16 {
    if (!true) {
        v10 := GenS1{F0: 2147483647, F1: genFn0(42), F2: -314}
        _ = v10
    } else {
        var v11 uint16 = 255
        _ = v11
    }
    v12 := []int{}
    _ = p0
    _ = p1
    _ = v12
    return 0
}

func main() {
    var v1 uint = 750
    var v2 int = 4294967295
    v3 := make(map[int]uint32)
    for fi13 := 0; fi13 < 10; fi13++ {
        v14 := GenS1{F0: 2147483648, F1: 32767, F2: 671}
        var v15 uint32 = 1
        v16 := GenS2{F0: uint8(v1), F1: -128}
        _ = fi13
        _ = v14
        _ = v15
        _ = v16
    }
    if (!false) {
        v17 := GenS3{F0: -1}
        _ = v17
    } else {
        v3[v2] = 331
        switch v2 {
        case -938:
            v2 -= -972
        default:
            var v18 int32 = -1
            _ = v18
        }
    }
    v2 -= v2
    switch (v2 % ((1) | 1)) {
    case 0:
        if (-9223372036854775808 != 943) {
            var v19 int64 = -826
            if false {
                v20 := []int{}
                switch "hello" {
                case "a":
                    var v21 int = (v2 * (v2 << (((v1 - v1)) & 63)))
                    _ = v21
                case "foo bar":
                    for fi22 := 0; fi22 < 10; fi22++ {
                        v23 := GenS4{F0: "hello", F1: -429}
                        for fi24 := 0; fi24 < 6; fi24++ {
                            fmt.Println(v19)
                            v1 = v1
                            v25 := GenS5{F0: ((!((v1 - v1) != v1)) && false)}
                            _ = fi24
                            _ = v25
                        }
                        if true {
                            switch fi22 {
                            case 32767:
                                fmt.Println(v2, v1, v2)
                            case -274:
                                if true {
                                    var v26 uint8 = 128
                                    _ = v26
                                }
                            case -2147483648:
                                if (221 > 250) {
                                    if (42 < 0) {
                                        fmt.Println(v19)
                                        var v27 uint16 = 32767
                                        _ = v27
                                    } else {
                                        var v28 int64 = v19
                                        var v29 int8 = -92
                                        _ = v28
                                        _ = v29
                                    }
                                    var v30 uint32 = uint32(v1)
                                    _ = v30
                                }
                            default:
                                var v31 int16 = genFn0(18446744073709551615)
                                _ = v31
                            }
                        }
                        _ = fi22
                        _ = v23
                    }
                default:
                    v2 -= 4294967295
                }
                _ = v20
            }
            if false {
                v2 = 777
                var v32 int = (v2 - (v2 + v2))
                _ = v32
            }
            _ = v19
        } else {
            if true {
                var v33 int8 = 1
                var v34 int = -607
                _ = v33
                _ = v34
            }
            v2 = v2
            v35 := GenS1{F0: 2147483647, F1: 1000, F2: -32768}
            _ = v35
        }
    default:
        v2 = v2
    }
    _ = len(v3)
    var v36 uint = v1
    v37 := GenS6{F0: 0, F1: 32767}
    v1 += v36
    switch "hello" {
    case "foo bar":
        for fi38 := 0; fi38 < 2; fi38++ {
            for fi39 := 0; fi39 < 17; fi39++ {
                var v40 uint64 = 558
                var v41 float32 = ((-2.25 + (float32(3.14159) - 0.0)) - 3.14159)
                v42 := GenS1{F0: 0, F1: -96, F2: -544}
                _ = fi39
                _ = v40
                _ = v41
                _ = v42
            }
            if true {
                for fi43 := 0; fi43 < 5; fi43++ {
                    var v44 uint8 = 119
                    v2 = fi38
                    _ = fi43
                    _ = v44
                }
                _ = len(v3)
            } else {
                if true {
                    var v45 uint = v1
                    var v46 string = ""
                    v47 := GenS7{F0: 0}
                    _ = v45
                    _ = v46
                    _ = v47
                } else {
                    fmt.Println(fi38, v1)
                }
            }
            switch "foo bar" {
            case "hello":
                switch fi38 {
                case -88:
                    if (((1 <= 1) || true) || true) {
                        v48 := GenS8{F0: 2, F1: -32768}
                        fmt.Println(v2, fi38)
                        _ = v48
                    }
                default:
                    switch "\n\t" {
                    case "":
                        v49 := GenS3{F0: -116}
                        _ = v49
                    case "a":
                        v36 -= v36
                    case "hello":
                        v2 = v2
                    default:
                        for fi50 := 0; fi50 < 18; fi50++ {
                            switch fi50 {
                            case 128:
                                v36 = (v36 * (^(v1)))
                            case -241:
                                var v51 float64 = 1.5
                                _ = v51
                            default:
                                v52 := GenS9{F0: -572, F1: 100, F2: 4294967295}
                                _ = v52
                            }
                            v53 := GenS5{F0: (2 < 1)}
                            _ = fi50
                            _ = v53
                        }
                    }
                }
            case "a":
                for fi54 := 0; fi54 < 2; fi54++ {
                    var v55 uint16 = 32767
                    _ = len(v3)
                    _ = fi54
                    _ = v55
                }
            }
            _ = fi38
        }
    case "\n\t":
        for fi56 := 0; fi56 < 19; fi56++ {
            fmt.Println(fi56, v36)
            var v57 uint8 = 115
            _ = fi56
            _ = v57
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
}
