package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int16; F1 int64 }
type GenS4 struct { F0 uint32; F1 int32 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 int32; F1 int32 }
func genFn0(p0 int16) uint32 {
    fmt.Println(p0)
    var v6 int32 = -2147483648
    _ = p0
    _ = v6
    return 1
}
func genFn1(p0 int32) string {
    v7 := GenS1{F0: 18446744073709551615}
    p0 = (-(p0))
    p0 = p0
    switch "a" {
    case "foo bar":
        p0 = (p0 % (((p0 | (p0 | (p0 + (p0 * (p0 | (p0 - (p0 | (p0 / ((-972) | 1)))))))))) | 1))
    case "a":
        v8 := []int{-9223372036854775808, 324}
        _ = v8
    case "\n\t":
        fmt.Println(p0)
    }
    _ = p0
    _ = v7
    return "foo bar"
}

func main() {
    var v1 int8 = -128
    v2 := []int{1}
    var v3 uint = 4294967295
    var v4 float64 = (1.5 + 0.0)
    var v5 int32 = 2147483647
    switch 32768 {
    case 2147483647:
        var v9 int8 = (v1 / ((v1) | 1))
        _ = v9
    case 32767:
        v10 := GenS2{F0: (v3 / (((v3 >> ((v3) & 63))) | 1))}
        _ = v10
    case 395:
        for fi11 := 0; fi11 < 8; fi11++ {
            v12 := GenS1{F0: 0}
            _ = fi11
            _ = v12
        }
    default:
        v13 := GenS2{F0: 255}
        _ = v13
    }
    fmt.Println(v4, v4)
    var v14 uint = (v3 | (v3 * v3))
    var v15 string = ("a" + "foo bar")
    v15 = v15
    v16 := GenS1{F0: 0}
    if len(v2) > 0 {
        v2[0] = int(v1)
    }
    var v17 int8 = v1
    var v18 int64 = 1
    if false {
        v17 = (v1 % ((-8) | 1))
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    switch 1 {
    case 128:
        switch -540 {
        case 4294967295:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        case 2147483647:
            var v19 string = ((v15 + v15) + v15)
            _ = v19
        case -830:
            for fi20 := 0; fi20 < 12; fi20++ {
                v21 := GenS3{F0: -32768, F1: -1}
                if true {
                    fmt.Println(v15)
                }
                if true {
                    var v22 int8 = v1
                    var v23 uint32 = 0
                    _ = v22
                    _ = v23
                }
                _ = fi20
                _ = v21
            }
        }
    default:
        var v24 uint16 = 832
        _ = v24
    }
    if (true || (129 < 0)) {
        v25 := GenS1{F0: 0}
        _ = v25
    }
    _ = len(v2)
    if true {
        switch 4294967295 {
        case 4294967295:
            v3 += v3
        case 2147483647:
            v26 := GenS1{F0: 69}
            _ = v26
        case 32767:
            v14 = 1
        }
        if (568 > (v14 / ((v3) | 1))) {
            if (-1 > v5) {
                switch "" {
                case "a":
                    var v27 float32 = float32(((1e10 / (((v4 / 0.1) / 3.14159) + 1.5)) / 1e10))
                    _ = v27
                default:
                    v28 := GenS4{F0: 25, F1: v5}
                    _ = v28
                }
                v29 := GenS1{F0: 891}
                _ = v29
            } else {
                for fi30 := 0; fi30 < 16; fi30++ {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    for fi31 := 0; fi31 < 6; fi31++ {
                        if false {
                            _ = len(v2)
                            v18 = -500
                        } else {
                            v32 := GenS1{F0: 42}
                            v5 = v5
                            v33 := GenS5{F0: 65535}
                            _ = v32
                            _ = v33
                        }
                        v18 = (v18 - v18)
                        _ = fi31
                    }
                    if true {
                        v15 = ((("foo bar" + "foo bar") + (v15 + v15)) + v15)
                    } else {
                        for fi34 := 0; fi34 < 19; fi34++ {
                            if (v4 > (float64(0.0) * v4)) {
                                v35 := GenS6{F0: v18}
                                v36 := GenS2{F0: v3}
                                var v37 int64 = int64(v14)
                                _ = v35
                                _ = v36
                                _ = v37
                            }
                            _ = fi34
                        }
                        var v38 int64 = v18
                        _ = v38
                    }
                    _ = fi30
                }
            }
            v18 -= 1
            v39 := GenS7{F0: (-(v5)), F1: (v5 >> ((v3) & 63))}
            _ = v39
        } else {
            switch ("a" + (v15 + (v15 + v15))) {
            case "foo bar":
                var v40 int = 128
                _ = v40
            default:
                fmt.Println(v18)
            }
            v41 := GenS4{F0: genFn0(32767), F1: 42}
            _ = len(v2)
            _ = v41
        }
        var v42 uint = (v14 | (^(v14)))
        _ = v42
    } else {
        for fi43 := 0; fi43 < 16; fi43++ {
            v44 := GenS5{F0: 65535}
            var v45 uint32 = 2
            var v46 uint16 = 147
            _ = fi43
            _ = v44
            _ = v45
            _ = v46
        }
        if ((4294967295 >= 330) && true) {
            v2 = append(v2, -1)
            var v47 bool = false
            var v48 uint32 = 2147483648
            _ = v47
            _ = v48
        } else {
            v17 = (v17 - (v17 - v17))
            if (!(!((1 < uint32(v3)) && (false && (false || true))))) {
                v49 := GenS1{F0: 1}
                _ = v49
            } else {
                v50 := GenS2{F0: 42}
                _ = v50
            }
        }
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi51 := 0; fi51 < 10; fi51++ {
        switch ((((v15 + v15) + v15) + v15) + (v15 + v15)) {
        case "\n\t":
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        case "":
            v3 = v3
        default:
            var v52 uint32 = 0
            _ = v52
        }
        _ = fi51
    }
    switch -559 {
    case 222:
        for fi53 := 0; fi53 < 9; fi53++ {
            fmt.Println(v15, v14, fi53)
            _ = fi53
        }
    case -1:
        var v54 float64 = v4
        _ = v54
    default:
        if false {
            var v55 string = (v15 + "")
            var v56 int = 128
            _ = v55
            _ = v56
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
