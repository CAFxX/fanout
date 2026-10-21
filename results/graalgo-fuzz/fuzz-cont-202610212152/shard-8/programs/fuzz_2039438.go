package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int32; F1 uint8 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint32; F1 int64; F2 int64 }
type GenS6 struct { F0 int }
func genFn0(p0 int8, p1 int, p2 bool) uint {
    fmt.Println(p1)
    var v4 int64 = 42
    v4 -= (v4 | (v4 >> ((18446744073709551615) & 63)))
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 42
}
func genFn1(p0 int, p1 int64, p2 int32) int64 {
    p0 -= 0
    v5 := GenS1{F0: false}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 9223372036854775807
}

func main() {
    var v1 uint64 = 18446744073709551615
    v2 := make(map[uint32]uint32)
    var v3 uint = 806
    for fi6 := 0; fi6 < 6; fi6++ {
        if (!("hello" != "")) {
            for fi7 := 0; fi7 < 9; fi7++ {
                v8 := GenS2{F0: 1, F1: 2}
                var v9 float64 = 1.5
                _ = fi7
                _ = v8
                _ = v9
            }
            v1 -= 18446744073709551615
        }
        _ = fi6
    }
    for fi10 := 0; fi10 < 12; fi10++ {
        var v11 float64 = ((1e10 * 0.0) * float64(float32(0.1)))
        for fi12 := 0; fi12 < 18; fi12++ {
            v13 := []int{}
            _ = fi12
            _ = v13
        }
        switch (fi10 ^ (fi10 << ((v3) & 63))) {
        case 42:
            for fi14 := 0; fi14 < 20; fi14++ {
                switch fi10 {
                case 9223372036854775807:
                    switch "foo bar" {
                    case "hello":
                        var v15 uint = v3
                        _ = v15
                    case "a":
                        var v16 int32 = 582
                        _ = v16
                    case "foo bar":
                        var v17 int64 = -1
                        _ = v17
                    default:
                        if true {
                            v18 := []int{-522, fi14}
                            v2[0] = 413
                            v11 = float64(1.5)
                            fmt.Println(v1, fi14)
                            _ = v18
                        } else {
                            v2[42] = 2147483648
                            var v19 int16 = 359
                            if false {
                                var v20 int16 = v19
                                _ = v20
                            } else {
                                fmt.Println(v3)
                            }
                            _ = v19
                        }
                    }
                default:
                    if (!(false || (!false))) {
                        v11 += v11
                    } else {
                        var v21 uint64 = v1
                        v22 := GenS3{F0: true}
                        _ = v21
                        _ = v22
                    }
                }
                _ = fi14
            }
        default:
            v23 := GenS2{F0: -2147483648, F1: 255}
            _ = v23
        }
        _ = fi10
        _ = v11
    }
    v24 := []int{-793, -1, int(v1)}
    if true {
        if true {
            var v25 int = 32768
            for fi26 := 0; fi26 < 10; fi26++ {
                var v27 uint = (v3 & (v3 + v3))
                v28 := GenS2{F0: 470, F1: 42}
                _ = fi26
                _ = v27
                _ = v28
            }
            _ = v25
        }
        _ = len(v2)
    } else {
        v2[2147483647] = 2147483648
    }
    if false {
        fmt.Println(v1, v3)
        switch "a" {
        case "foo bar":
            for fi29 := 0; fi29 < 1; fi29++ {
                var v30 uint64 = v1
                fmt.Println(v3, v1, v1)
                _ = fi29
                _ = v30
            }
        case "":
            for fi31 := 0; fi31 < 15; fi31++ {
                if false {
                    var v32 int64 = -1
                    v1 -= (v1 * (v1 & v1))
                    _ = v32
                } else {
                    v3 += 65535
                    v33 := GenS1{F0: false}
                    _ = v33
                }
                for fi34 := 0; fi34 < 14; fi34++ {
                    fmt.Println(v3, fi34, fi34)
                    _ = fi34
                }
                _ = fi31
            }
        }
    } else {
        if (((float64(((1.5 / -2.25) + ((1.5 + (1e10 * (float32((0.1 - (0.0 * (1e10 * ((0.1 / -2.25) / 3.14159))))) + 0.1))) + -2.25))) / 1e10) + 1e10) < 1.5) {
            v3 += (^(v3))
            var v35 float64 = (3.14159 / 1.0)
            _ = v35
        }
        for fi36 := 0; fi36 < 6; fi36++ {
            switch 230 {
            case -1:
                _ = len(v24)
            case -484:
                switch -140 {
                case -1:
                    v37 := v2[996]
                    _ = v37
                case 758:
                    v3 = v3
                case 295:
                    for fi38 := 0; fi38 < 4; fi38++ {
                        var v39 bool = false
                        _ = len(v24)
                        _ = fi38
                        _ = v39
                    }
                default:
                    var v40 string = ("a" + "a")
                    _ = v40
                }
            case 41:
                if (!false) {
                    var v41 float64 = 1.5
                    v42 := GenS4{F0: 255}
                    _ = v41
                    _ = v42
                } else {
                    for fi43 := 0; fi43 < 20; fi43++ {
                        var v44 int16 = -549
                        v44 = (v44 + (v44 * 32767))
                        fmt.Println(v3, v1, v44)
                        _ = fi43
                        _ = v44
                    }
                    var v45 float64 = 1e10
                    v3 += (^(v3))
                    _ = v45
                }
            default:
                fmt.Println(v3)
            }
            switch (fi36 / ((fi36) | 1)) {
            case 0:
                if (1.5 >= (1e10 * 1e10)) {
                    var v46 string = (("a" + ("foo bar" + "a")) + "foo bar")
                    var v47 int64 = 0
                    v47 = (v47 / ((genFn1(fi36, 642, 2147483647)) | 1))
                    _ = v46
                    _ = v47
                }
            }
            v48 := v2[42]
            _ = fi36
            _ = v48
        }
    }
    v49 := GenS5{F0: 2147483648, F1: -109, F2: 868}
    var v50 uint8 = 49
    switch -291 {
    case 9223372036854775807:
        v51 := GenS5{F0: 402, F1: 1, F2: 42}
        _ = v51
    default:
        v52 := GenS4{F0: (v50 & 255)}
        _ = v52
    }
    if (!(42 < -9223372036854775808)) {
        v53 := GenS6{F0: 4294967295}
        switch "" {
        case "a":
            if len(v24) >= 2 {
                v24 = v24[1:len(v24)]
            }
        case "":
            var v54 uint16 = 32767
            _ = v54
        default:
            if (false || ((1 <= 0) || true)) {
                _ = len(v24)
            }
        }
        _ = v53
    }
    _ = len(v24)
    v50 -= (v50 - v50)
    v55 := GenS3{F0: ((genFn1(-165, -9223372036854775808, 0) != 1) && (32768 < -512))}
    v56 := GenS4{F0: (v50 >> ((42) & 63))}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v3:", v3)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v49
    _ = v50
    _ = v55
    _ = v56
}
