package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 string; F1 uint }
type GenS3 struct { F0 int64; F1 int32 }
type GenS4 struct { F0 bool; F1 uint16; F2 int8 }
type GenS5 struct { F0 bool; F1 uint16; F2 uint64 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint16; F1 uint16; F2 uint }
func genFn0(p0 int32, p1 string, p2 int32) int8 {
    if true {
        p2 += p0
        var v7 int = 9223372036854775807
        _ = v7
    }
    p0 -= -519
    if false {
        if (642 < 42) {
            p0 += (p2 + p0)
            p0 -= p0
        } else {
            p1 = ("foo bar" + p1)
        }
    }
    v8 := GenS1{F0: -705}
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return -1
}

func main() {
    var v1 int16 = 233
    var v2 uint16 = 255
    v3 := make(map[uint32]int)
    var v4 string = "\n\t"
    var v5 uint32 = 988
    v6 := make(map[int]string)
    v9 := GenS2{F0: (("\n\t" + v4) + v4), F1: 42}
    delete(v6, -533)
    v10 := GenS3{F0: -1, F1: 0}
    if false {
        delete(v6, 256)
        switch 439 {
        case 42:
            var v11 int16 = v1
            _ = v11
        case 127:
            var v12 float64 = 1.5
            _ = v12
        }
        switch -9223372036854775808 {
        case 49:
            var v13 bool = (!(-98 > -128))
            _ = v13
        case -9223372036854775808:
            if true {
                v14 := GenS4{F0: false, F1: (^(v2)), F2: 1}
                fmt.Println(v4, v1, v5)
                _ = v14
            }
        case -2147483648:
            v4 = ((v4 + v4) + v4)
        default:
            v1 += v1
        }
    }
    switch -209 {
    case 1:
        if false {
            var v15 int32 = 1
            _ = v15
        } else {
            v3[v5] = -9223372036854775808
        }
    default:
        var v16 int32 = -1
        _ = v16
    }
    v2 = (^(v2))
    if true {
        switch 1 {
        case -536:
            v17 := GenS5{F0: true, F1: (v2 << ((42) & 63)), F2: 2}
            _ = v17
        case 9223372036854775807:
            var v18 uint16 = (v2 | (v2 & (v2 * 255)))
            _ = v18
        case 32767:
            var v19 uint16 = (v2 / ((v2) | 1))
            _ = v19
        }
        switch v4 {
        case "hello":
            switch 2147483647 {
            case -1:
                if true {
                    if false {
                        for fi20 := 0; fi20 < 9; fi20++ {
                            fmt.Println(v5, v1)
                            v3[v5] = (fi20 & -690)
                            _ = fi20
                        }
                        v21 := GenS4{F0: (-559 < 281), F1: v2, F2: -1}
                        _ = v21
                    } else {
                        v22 := GenS3{F0: 1, F1: 0}
                        var v23 int32 = -1
                        _ = v22
                        _ = v23
                    }
                }
            case 732:
                var v24 uint = 4294967295
                _ = v24
            case 9223372036854775807:
                if false {
                    v25 := []int{-2147483648, 128, 127, -2147483648}
                    _ = v25
                } else {
                    for fi26 := 0; fi26 < 18; fi26++ {
                        v4 = ((v4 + (v4 + ("" + v4))) + "\n\t")
                        _ = fi26
                    }
                    fmt.Println(v5)
                }
            default:
                if false {
                    var v27 float64 = -2.25
                    if ((v5 == v5) && (((genFn0(705, v4, -472) < 47) && true) && false)) {
                        switch -9223372036854775808 {
                        case -292:
                            var v28 int64 = 0
                            _ = v28
                        case -2147483648:
                            v5 += (v5 - v5)
                        }
                        v29 := GenS4{F0: false, F1: (v2 - 0), F2: genFn0(-1, (("\n\t" + (((v4 + "\n\t") + ((((v4 + v4) + v4) + v4) + v4)) + v4)) + v4), 42)}
                        if false {
                            if true {
                                var v30 int = -813
                                _ = v30
                            }
                            v27 = v27
                            v4 = v4
                        }
                        _ = v29
                    }
                    v1 += v1
                    _ = v27
                } else {
                    v5 = (v5 | uint32(v2))
                }
            }
        case "foo bar":
            fmt.Println(v4, v4, v5)
        case "":
            v31 := GenS6{F0: v2}
            _ = v31
        }
    }
    if true {
        if (!true) {
            if (true && (39 > 42)) {
                for fi32 := 0; fi32 < 13; fi32++ {
                    if true {
                        v33 := GenS7{F0: v2, F1: v2, F2: 232}
                        _ = v33
                    } else {
                        fmt.Println(v4, v1)
                    }
                    fmt.Println(v5)
                    var v34 int32 = -2147483648
                    _ = fi32
                    _ = v34
                }
                var v35 uint16 = (v2 >> ((1) & 63))
                _ = v35
            }
            var v36 uint8 = 128
            v37 := []int{32767}
            _ = v36
            _ = v37
        }
    }
    for fi38 := 0; fi38 < 19; fi38++ {
        for fi39 := 0; fi39 < 16; fi39++ {
            v1 = -280
            _ = fi39
        }
        _ = fi38
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
