package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint; F2 uint }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint8; F1 uint }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int32; F1 uint }
func genFn0(p0 uint8) string {
    var v6 float64 = 0.0
    v6 += (float64(1.5) + v6)
    _ = p0
    _ = v6
    return "foo bar"
}
func genFn1() uint64 {
    var v7 int16 = 546
    for fi8 := 0; fi8 < 19; fi8++ {
        if (-658 != v7) {
            v9 := GenS1{F0: (v7 / ((-34) | 1)), F1: 1, F2: 65535}
            _ = v9
        } else {
            v7 = (v7 ^ v7)
        }
        switch 42 {
        case 127:
            v10 := GenS2{F0: 55}
            _ = v10
        case 264:
            var v11 uint64 = 1
            _ = v11
        case 32767:
            switch ("" + "a") {
            case "":
                v7 += 1
            }
        }
        _ = fi8
    }
    var v12 bool = (!(!(65535 > 181)))
    if false {
        var v13 int8 = -109
        v14 := []int{0, -888, -2147483648}
        v12 = (!true)
        _ = v13
        _ = v14
    } else {
        v15 := []int{427, 757}
        var v16 bool = v12
        _ = v15
        _ = v16
    }
    if (!v12) {
        var v17 float64 = (1e10 * (0.0 + float64(1.5)))
        switch -904 {
        case 66:
            v18 := []int{-887, 128}
            _ = v18
        default:
            v17 = v17
        }
        v19 := GenS3{F0: 127, F1: 973}
        _ = v17
        _ = v19
    }
    _ = v12
    _ = v7
    return 18446744073709551615
}

func main() {
    var v1 float32 = 3.14159
    var v2 uint16 = 1
    var v3 uint16 = (v2 % ((v2) | 1))
    v4 := []string{"hello", ("hello" + "foo bar"), ("" + "foo bar")}
    var v5 int32 = -608
    fmt.Println(v1, v5, v2)
    _ = len(v4)
    _ = len(v4)
    var v20 uint8 = 202
    v5 = (v5 >> ((709) & 63))
    if false {
        switch "hello" {
        case "":
            var v21 uint8 = (v20 & v20)
            _ = v21
        case "a":
            _ = len(v4)
        case "foo bar":
            v5 = 0
        }
        fmt.Println(v20)
        var v22 float32 = (v1 * v1)
        _ = v22
    } else {
        for fi23 := 0; fi23 < 1; fi23++ {
            v3 = (v2 >> ((255) & 63))
            v24 := GenS3{F0: v20, F1: 1}
            _ = fi23
            _ = v24
        }
        v1 -= (1.5 * v1)
    }
    switch "" {
    case "":
        _ = len(v4)
    case "a":
        if true {
            v5 += v5
            var v25 int8 = 127
            if (false && false) {
                var v26 uint8 = (v20 / ((v20) | 1))
                for fi27 := 0; fi27 < 19; fi27++ {
                    for fi28 := 0; fi28 < 20; fi28++ {
                        if false {
                            v1 -= v1
                            var v29 int8 = (v25 >> ((18446744073709551615) & 63))
                            _ = v29
                        }
                        if false {
                            if len(v4) > 0 {
                                v4[0] = "\n\t"
                            }
                            fmt.Println(fi28, v25)
                        } else {
                            var v30 int32 = v5
                            if false {
                                v31 := GenS4{F0: ((v26 << ((508) & 63)) < uint8(v2))}
                                _ = v31
                            }
                            var v32 uint = 1
                            if len(v4) > 0 {
                                v4[0] = ("" + "a")
                            }
                            _ = v30
                            _ = v32
                        }
                        if false {
                            var v33 float32 = v1
                            v34 := GenS3{F0: (v20 * (v26 + (^(v20)))), F1: 65535}
                            _ = v33
                            _ = v34
                        }
                        _ = fi28
                    }
                    var v35 float32 = 0.0
                    for fi36 := 0; fi36 < 6; fi36++ {
                        v37 := GenS5{F0: -1, F1: 18446744073709551615}
                        v5 -= v5
                        _ = fi36
                        _ = v37
                    }
                    _ = fi27
                    _ = v35
                }
                v2 = v2
                _ = v26
            }
            _ = v25
        }
    case "\n\t":
        v38 := GenS4{F0: (true && false)}
        _ = v38
    }
    v5 = v5
    for fi39 := 0; fi39 < 20; fi39++ {
        var v40 int32 = -1
        if (42 == 556) {
            var v41 int64 = -360
            switch fi39 {
            case -2147483648:
                v42 := GenS2{F0: 2147483647}
                _ = v42
            case 128:
                v2 -= v3
            case -9223372036854775808:
                var v43 int32 = (v5 >> ((uint(v3)) & 63))
                _ = v43
            }
            for fi44 := 0; fi44 < 18; fi44++ {
                v4 = append(v4, "foo bar")
                _ = fi44
            }
            _ = v41
        }
        if true {
            _ = len(v4)
        }
        _ = fi39
        _ = v40
    }
    for fi45 := 0; fi45 < 14; fi45++ {
        var v46 uint64 = 1
        if len(v4) > 0 {
            v4[0] = ""
        }
        _ = fi45
        _ = v46
    }
    if (int8(v3) == 100) {
        var v47 float64 = 1e10
        _ = v47
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
