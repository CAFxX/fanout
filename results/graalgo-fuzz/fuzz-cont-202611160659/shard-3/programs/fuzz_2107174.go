package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int; F2 bool }
type GenS2 struct { F0 int16; F1 int64; F2 int8 }
type GenS3 struct { F0 int8; F1 uint; F2 uint64 }
type GenS4 struct { F0 int; F1 bool; F2 int16 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 bool; F1 uint32 }
func genFn0() string {
    v7 := GenS1{F0: "foo bar", F1: 249, F2: true}
    v8 := GenS2{F0: -1, F1: 0, F2: 100}
    v9 := []int{}
    _ = v7
    _ = v8
    _ = v9
    return ""
}

func main() {
    var v1 string = "hello"
    var v2 int16 = 0
    var v3 int8 = 127
    var v4 bool = (v3 == v3)
    v5 := make(map[int]bool)
    v6 := make(map[string]bool)
    v5[9223372036854775807] = true
    v10 := GenS3{F0: (v3 & 127), F1: 1, F2: 42}
    fmt.Println(v4)
    if (!((!false) && (2 == 68))) {
        var v11 uint8 = 0
        _ = v11
    } else {
        if false {
            fmt.Println(v4)
            if v4 {
                v12 := []int{655, -382}
                for fi13 := 0; fi13 < 3; fi13++ {
                    fmt.Println(fi13)
                    var v14 uint16 = 42
                    v3 -= int8(v2)
                    _ = fi13
                    _ = v14
                }
                _ = v12
            } else {
                if false {
                    v4 = (!false)
                }
            }
            switch 128 {
            case 128:
                var v15 bool = (!v4)
                _ = v15
            }
        }
        switch 1 {
        case 125:
            var v16 float32 = 1e10
            _ = v16
        case 2147483647:
            for fi17 := 0; fi17 < 5; fi17++ {
                v3 += v3
                fmt.Println(v1)
                _ = fi17
            }
        case 0:
            for fi18 := 0; fi18 < 14; fi18++ {
                var v19 string = ("a" + v1)
                var v20 string = v19
                _ = fi18
                _ = v19
                _ = v20
            }
        default:
            v1 = "foo bar"
        }
        var v21 uint32 = 919
        _ = v21
    }
    for fi22 := 0; fi22 < 12; fi22++ {
        for fi23 := 0; fi23 < 7; fi23++ {
            v3 -= (v3 / (((v3 << ((1) & 63))) | 1))
            var v24 int16 = v2
            v4 = (-1 < -803)
            _ = fi23
            _ = v24
        }
        v25 := GenS4{F0: 42, F1: false, F2: (v2 - 1)}
        fmt.Println(v4)
        _ = fi22
        _ = v25
    }
    v26 := GenS4{F0: int(v2), F1: true, F2: -725}
    var v27 int32 = 236
    if (169 <= 0) {
        var v28 float32 = (-2.25 / (-2.25 - ((0.1 - (0.1 - 1e10)) * 3.14159)))
        v29 := GenS3{F0: -1, F1: 503, F2: 42}
        _ = v28
        _ = v29
    } else {
        var v30 string = "a"
        _ = v30
    }
    v3 = 42
    v31 := GenS5{F0: 9223372036854775807}
    switch -518 {
    case 2147483647:
        if (1 < 18446744073709551615) {
            v32 := []int{42, 1, -2147483648}
            _ = len(v32)
            _ = v32
        } else {
            var v33 uint64 = 1
            v33 = 76
            fmt.Println(v1)
            _ = v33
        }
    case 32768:
        for fi34 := 0; fi34 < 17; fi34++ {
            v4 = v4
            var v35 uint64 = 1
            var v36 bool = true
            _ = fi34
            _ = v35
            _ = v36
        }
    }
    if true {
        switch -9223372036854775808 {
        case 4294967295:
            for fi37 := 0; fi37 < 18; fi37++ {
                fmt.Println(v27, v1)
                for fi38 := 0; fi38 < 19; fi38++ {
                    v1 = ("a" + v1)
                    var v39 uint8 = 1
                    _ = fi38
                    _ = v39
                }
                v3 -= (-(v3))
                _ = fi37
            }
        case -154:
            delete(v6, "foo bar")
        }
    }
    if v4 {
        for fi40 := 0; fi40 < 12; fi40++ {
            var v41 uint16 = 663
            fmt.Println(v4, v3, v41)
            if v4 {
                fmt.Println(v2)
                if v4 {
                    v42 := GenS2{F0: (-(v2)), F1: -9223372036854775808, F2: v3}
                    v43 := GenS6{F0: v4, F1: 1}
                    _ = v42
                    _ = v43
                } else {
                    for fi44 := 0; fi44 < 20; fi44++ {
                        fmt.Println(v27, v3, v3)
                        _ = fi44
                    }
                }
            }
            _ = fi40
            _ = v41
        }
        if true {
            v2 -= (v2 + v2)
            for fi45 := 0; fi45 < 7; fi45++ {
                v1 = v1
                _ = fi45
            }
        } else {
            var v46 uint8 = 1
            v47 := GenS6{F0: (((42 <= 18446744073709551615) || (924 != 18446744073709551615)) || v4), F1: 1}
            _ = v46
            _ = v47
        }
    }
    if false {
        fmt.Println(v3)
    } else {
        for fi48 := 0; fi48 < 4; fi48++ {
            if (42 != (v27 & v27)) {
                var v49 float32 = float32(float64(0.1))
                if ((v3 % (((v3 >> ((1) & 63))) | 1)) < 0) {
                    v1 = (v1 + (v1 + v1))
                    var v50 float64 = 1e10
                    var v51 bool = (255 == 128)
                    _ = v50
                    _ = v51
                } else {
                    var v52 float64 = float64(v49)
                    _ = v52
                }
                _ = v49
            } else {
                v2 = v2
            }
            _ = fi48
        }
        switch v1 {
        case "":
            v6[""] = (4294967295 < 939)
        case "hello":
            v4 = (65535 >= 277)
        default:
            fmt.Println(v3, v4, v1)
        }
        var v53 int64 = -162
        _ = v53
    }
    v4 = false
    v2 -= 42
    var v54 uint8 = 130
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v54:", v54)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v54
    _ = v6
}
