package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int; F1 uint64 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int8; F1 uint64; F2 uint32 }
type GenS5 struct { F0 uint64; F1 int }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 int; F1 uint; F2 string }
func genFn0(p0 int16) int {
    v6 := GenS1{F0: 2}
    v7 := GenS2{F0: 14, F1: 2}
    for fi8 := 0; fi8 < 7; fi8++ {
        var v9 uint16 = 0
        v10 := GenS1{F0: 127}
        for fi11 := 0; fi11 < 6; fi11++ {
            var v12 int = fi8
            var v13 float64 = 0.0
            _ = fi11
            _ = v12
            _ = v13
        }
        _ = fi8
        _ = v10
        _ = v9
    }
    var v14 uint64 = 77
    _ = p0
    _ = v14
    _ = v6
    _ = v7
    return -2147483648
}

func main() {
    var v1 uint = 65535
    v2 := []uint32{}
    v3 := make(map[string]uint32)
    var v4 uint64 = 121
    var v5 bool = (-25 >= 68)
    fmt.Println(v4, v5)
    if v5 {
        fmt.Println(v5)
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v15 string = "\n\t"
    var v16 float64 = (1.5 / 1.0)
    var v17 uint = (^(v1))
    v4 += v4
    v4 = v4
    var v18 int16 = 1
    v19 := GenS3{F0: 786}
    v15 = v15
    for fi20 := 0; fi20 < 14; fi20++ {
        fmt.Println(v17, v4)
        for fi21 := 0; fi21 < 8; fi21++ {
            if (-815 != 9223372036854775807) {
                v22 := GenS2{F0: (fi20 >> ((v17) & 63)), F1: 564}
                v16 -= ((v16 + v16) - (3.14159 + 0.0))
                _ = v22
            } else {
                v23 := GenS4{F0: -128, F1: v4, F2: 719}
                if v5 {
                    var v24 uint8 = 2
                    _ = v24
                }
                var v25 string = "foo bar"
                var v26 uint32 = 306
                _ = v23
                _ = v25
                _ = v26
            }
            _ = fi21
        }
        _ = fi20
    }
    switch 42 {
    case 32768:
        var v27 uint32 = 0
        _ = v27
    case 127:
        var v28 uint64 = (v4 - v4)
        _ = v28
    case -864:
        v1 = (v17 >> ((v1) & 63))
    }
    switch 151 {
    case -1:
        var v29 int = 32767
        _ = v29
    case 1:
        if v5 {
            _ = len(v3)
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            v17 = (v17 + v1)
        }
    case -9223372036854775808:
        v30 := GenS5{F0: 0, F1: -1}
        _ = v30
    default:
        if (2147483647 != -257) {
            var v31 uint16 = 0
            _ = v31
        } else {
            for fi32 := 0; fi32 < 9; fi32++ {
                var v33 uint = v1
                switch (^(fi32)) {
                case 9223372036854775807:
                    var v34 uint16 = 325
                    _ = v34
                default:
                    var v35 float32 = 0.1
                    _ = v35
                }
                _ = fi32
                _ = v33
            }
            var v36 uint64 = (^(v4))
            v37 := GenS6{F0: v4}
            _ = v36
            _ = v37
        }
    }
    if v5 {
        if (0 > 4294967295) {
            var v38 int = 9223372036854775807
            for fi39 := 0; fi39 < 3; fi39++ {
                fmt.Println(v15, v5, fi39)
                v3[v15] = 65535
                var v40 int = fi39
                _ = fi39
                _ = v40
            }
            if v5 {
                var v41 uint32 = 418
                for fi42 := 0; fi42 < 8; fi42++ {
                    var v43 float32 = 3.14159
                    _ = len(v3)
                    if (3.14159 < v16) {
                        for fi44 := 0; fi44 < 3; fi44++ {
                            v45 := GenS5{F0: (^(v4)), F1: 42}
                            _ = fi44
                            _ = v45
                        }
                        var v46 int = (fi42 % (((fi42 + -2)) | 1))
                        _ = v46
                    } else {
                        var v47 int8 = 1
                        switch v38 {
                        case 9223372036854775807:
                            v17 -= v17
                        case 475:
                            fmt.Println(v1)
                        case 4294967295:
                            if v5 {
                                v17 -= 42
                            }
                        default:
                            for fi48 := 0; fi48 < 1; fi48++ {
                                if len(v2) > 0 {
                                    v2[0] = (v41 & (^(v41)))
                                }
                                v4 += 2
                                v49 := v3["a"]
                                _ = fi48
                                _ = v49
                            }
                        }
                        _ = v47
                    }
                    _ = fi42
                    _ = v43
                }
                _ = v41
            } else {
                for fi50 := 0; fi50 < 3; fi50++ {
                    v51 := GenS7{F0: (v38 >> ((uint(v4)) & 63)), F1: uint(v4), F2: "hello"}
                    _ = fi50
                    _ = v51
                }
                var v52 int8 = 100
                _ = v52
            }
            _ = v38
        }
        v53 := GenS3{F0: 977}
        for fi54 := 0; fi54 < 4; fi54++ {
            var v55 int32 = 2147483647
            _ = fi54
            _ = v55
        }
        _ = v53
    } else {
        v4 = v4
    }
    v1 -= (v17 | v1)
    v5 = v5
    fmt.Println(v17, v17, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
