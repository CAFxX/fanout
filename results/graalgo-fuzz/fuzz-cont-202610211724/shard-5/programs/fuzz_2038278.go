package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint16; F1 bool }
type GenS4 struct { F0 int32; F1 int16 }
type GenS5 struct { F0 uint64; F1 int8 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 uint8; F1 int8; F2 int16 }
type GenS8 struct { F0 int16; F1 int32; F2 uint }
func genFn0(p0 bool) int16 {
    p0 = ((!p0) && (3.14159 >= (-2.25 - -2.25)))
    var v7 int16 = 1
    _ = p0
    _ = v7
    return (^(v7))
}
func genFn1() int64 {
    var v8 int = 4294967295
    v8 += 1
    switch v8 {
    case 216:
        if (592 >= 42) {
            v9 := GenS1{F0: -82, F1: ((32767 != v8) && false)}
            v10 := []int{v8}
            _ = v10
            _ = v9
        }
    case 4294967295:
        if false {
            v11 := GenS2{F0: ("foo bar" + ("foo bar" + "hello"))}
            _ = v11
        }
    default:
        v12 := []int{(v8 & (v8 ^ (v8 & (v8 & v8)))), (v8 ^ (v8 % ((v8) | 1))), (-(v8))}
        _ = v12
    }
    v13 := []int{(v8 >> ((18446744073709551615) & 63)), 127, v8}
    if (false && true) {
        var v14 uint32 = 84
        _ = len(v13)
        v14 = 607
        _ = v14
    } else {
        fmt.Println(v8)
        switch (v8 ^ (v8 | (v8 ^ v8))) {
        case 2147483647:
            v8 -= (v8 & 4294967295)
        case -282:
            for fi15 := 0; fi15 < 1; fi15++ {
                fmt.Println(v8)
                var v16 int = (fi15 / (((v8 | (fi15 ^ v8))) | 1))
                _ = fi15
                _ = v16
            }
        }
    }
    _ = v13
    _ = v8
    return -1
}

func main() {
    var v1 string = "hello"
    v2 := make(map[uint32]int)
    var v3 int16 = 42
    var v4 string = ""
    v5 := make(map[uint32]string)
    var v6 int32 = 0
    v17 := GenS3{F0: 0, F1: (2 != 780)}
    v3 -= v3
    for fi18 := 0; fi18 < 7; fi18++ {
        v19 := GenS4{F0: v6, F1: v3}
        for fi20 := 0; fi20 < 20; fi20++ {
            if (127 == 127) {
                var v21 uint64 = 1
                v4 = "a"
                _ = v21
            } else {
                var v22 uint32 = 2147483647
                if true {
                    switch (^(fi20)) {
                    case 42:
                        var v23 bool = (!(0.0 < (float32(1.5) - (3.14159 - 1.5))))
                        _ = v23
                    }
                    if ((false || true) && (v3 < 761)) {
                        var v24 int = (fi18 % (((fi20 % (((fi18 + fi18)) | 1))) | 1))
                        var v25 uint16 = 1
                        _ = v24
                        _ = v25
                    } else {
                        v26 := GenS2{F0: "\n\t"}
                        var v27 int = 42
                        v6 = -922
                        _ = v26
                        _ = v27
                    }
                    if (0 == 9223372036854775807) {
                        v3 += 42
                    }
                } else {
                    switch fi20 {
                    case 127:
                        v28 := GenS4{F0: (-(v6)), F1: (v3 / ((42) | 1))}
                        _ = v28
                    case -1:
                        var v29 int16 = v3
                        _ = v29
                    }
                    var v30 float64 = (-2.25 / 1.5)
                    _ = v30
                }
                _ = v22
            }
            var v31 int64 = -9223372036854775808
            var v32 uint32 = 2147483648
            _ = fi20
            _ = v31
            _ = v32
        }
        var v33 string = v1
        _ = fi18
        _ = v19
        _ = v33
    }
    var v34 float64 = (float64(-2.25) * float64((((1e10 / 1.5) + (1.5 / -2.25)) + 0.1)))
    v35 := []int{32768, 2147483647, 42, 792}
    var v36 int = 1
    for fi37 := 0; fi37 < 4; fi37++ {
        v36 += fi37
        if (!false) {
            if (v3 != 1000) {
                v3 += (v3 | (v3 + -959))
                var v38 float64 = 1e10
                if false {
                    for fi39 := 0; fi39 < 12; fi39++ {
                        v40 := GenS5{F0: 4294967295, F1: 127}
                        _ = fi39
                        _ = v40
                    }
                }
                _ = v38
            } else {
                v3 = int16(v36)
            }
            var v41 bool = (!(1 != 1))
            _ = v41
        } else {
            switch (fi37 << ((65535) & 63)) {
            case -523:
                var v42 uint = 65535
                _ = v42
            case 4294967295:
                var v43 int32 = v6
                _ = v43
            case -9223372036854775808:
                var v44 int8 = 100
                _ = v44
            default:
                v45 := GenS1{F0: -86, F1: (-31 >= 41)}
                _ = v45
            }
            v6 = 1
            var v46 bool = true
            _ = v46
        }
        _ = fi37
    }
    if (!false) {
        if (0 >= genFn1()) {
            v47 := v5[2147483648]
            for fi48 := 0; fi48 < 9; fi48++ {
                v49 := GenS6{F0: (v6 << ((0) & 63))}
                v50 := GenS7{F0: 127, F1: -128, F2: (v3 + genFn0((-2.25 != 0.0)))}
                for fi51 := 0; fi51 < 4; fi51++ {
                    var v52 bool = false
                    var v53 uint32 = 2147483648
                    var v54 bool = (((v34 == v34) || (4294967295 != 18446744073709551615)) || (v3 == (v3 - genFn0(((v52 || ((!(467 >= v3)) && ((^(v3)) != 627))) && v52)))))
                    _ = fi51
                    _ = v52
                    _ = v53
                    _ = v54
                }
                _ = fi48
                _ = v49
                _ = v50
            }
            _ = v47
        } else {
            v6 -= 0
            v55 := GenS8{F0: (v3 / (((^(v3))) | 1)), F1: (v6 ^ 42), F2: 255}
            if true {
                v3 = genFn0(true)
                v1 = v4
            }
            _ = v55
        }
        switch (^(v36)) {
        case -2147483648:
            fmt.Println(v34)
        case 128:
            if (false || (v1 < v1)) {
                var v56 int16 = 1000
                if false {
                    v34 = 0.1
                }
                _ = v56
            }
        default:
            var v57 int64 = -9223372036854775808
            _ = v57
        }
    }
    var v58 int16 = (v3 % (((v3 + v3)) | 1))
    if (("a" + v4) < "\n\t") {
        for fi59 := 0; fi59 < 19; fi59++ {
            v35 = append(v35, (^(v36)))
            _ = fi59
        }
    } else {
        fmt.Println(v34, v4)
        switch v36 {
        case -1:
            var v60 bool = true
            _ = v60
        case 593:
            v61 := v2[2]
            _ = v61
        default:
            var v62 uint = 18446744073709551615
            _ = v62
        }
        v34 -= v34
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v58:", v58)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v58
    _ = v6
}
