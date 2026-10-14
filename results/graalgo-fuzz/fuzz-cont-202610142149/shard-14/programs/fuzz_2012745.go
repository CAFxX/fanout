package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 int8 }
type GenS2 struct { F0 uint32; F1 uint8; F2 uint16 }
type GenS3 struct { F0 uint16; F1 int16; F2 uint8 }
type GenS4 struct { F0 uint16; F1 uint16; F2 uint16 }
type GenS5 struct { F0 uint64; F1 bool }
type GenS6 struct { F0 int32; F1 int }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 int16; F1 uint64; F2 int32 }
func genFn0(p0 int64, p1 uint16, p2 uint8) int64 {
    v7 := GenS1{F0: "foo bar", F1: int32(p1), F2: 100}
    switch 2147483647 {
    case 9223372036854775807:
        p1 -= p1
    case 1:
        fmt.Println(p0, p0)
    case -1:
        v8 := GenS2{F0: 606, F1: (p2 >> ((1) & 63)), F2: p1}
        _ = v8
    default:
        p2 += 127
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return p0
}

func main() {
    var v1 int32 = 1
    var v2 bool = false
    var v3 float64 = (3.14159 + 0.0)
    v4 := []int{}
    v5 := []string{}
    v6 := []uint32{2147483647, 2, 2147483648}
    var v9 bool = v2
    if (605 != 610) {
        var v10 int = 2147483647
        if v2 {
            var v11 bool = v9
            if false {
                var v12 uint64 = 0
                v13 := GenS1{F0: "", F1: v1, F2: int8(v12)}
                fmt.Println(v11)
                _ = v12
                _ = v13
            } else {
                var v14 uint32 = 42
                v15 := GenS3{F0: 1, F1: 1, F2: 191}
                for fi16 := 0; fi16 < 11; fi16++ {
                    var v17 uint = 255
                    v10 += fi16
                    fmt.Println(v2, v3)
                    _ = fi16
                    _ = v17
                }
                _ = v14
                _ = v15
            }
            _ = v11
        }
        _ = v10
    }
    v18 := GenS4{F0: 32767, F1: 0, F2: 2}
    v1 -= -2147483648
    if ((!(!(v9 && (-35 >= -1)))) || v9) {
        v19 := GenS4{F0: 681, F1: 65535, F2: 255}
        v9 = false
        _ = v19
    } else {
        v1 = v1
    }
    v20 := GenS4{F0: 2, F1: 65535, F2: 65535}
    v21 := GenS5{F0: 488, F1: v9}
    switch 216 {
    case -1:
        if true {
            for fi22 := 0; fi22 < 11; fi22++ {
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                if (42 > 42) {
                    if v9 {
                        switch fi22 {
                        case 128:
                            v9 = v9
                        case 2147483647:
                            var v23 uint16 = 2
                            _ = v23
                        default:
                            if ((!v9) && false) {
                                var v24 string = (("a" + "foo bar") + "a")
                                var v25 int32 = v1
                                fmt.Println(v24, v3, v2)
                                _ = v24
                                _ = v25
                            }
                        }
                        var v26 int32 = 2147483647
                        v27 := GenS6{F0: v26, F1: fi22}
                        _ = v26
                        _ = v27
                    }
                }
                _ = fi22
            }
            v1 = -1
            for fi28 := 0; fi28 < 19; fi28++ {
                v29 := GenS4{F0: 142, F1: 660, F2: 42}
                var v30 float64 = 3.14159
                _ = fi28
                _ = v29
                _ = v30
            }
        } else {
            _ = len(v4)
            var v31 uint32 = 172
            _ = v31
        }
    case 321:
        if len(v6) > 0 {
            v6[0] = 4294967295
        }
    default:
        v2 = (!(!v9))
    }
    for fi32 := 0; fi32 < 8; fi32++ {
        fmt.Println(v1)
        var v33 uint = 386
        if (v33 <= (^(v33))) {
            for fi34 := 0; fi34 < 16; fi34++ {
                var v35 uint16 = 364
                if (1 < 720) {
                    v36 := GenS4{F0: 32767, F1: 255, F2: (v35 + v35)}
                    if false {
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                        fmt.Println(v33, v9)
                        v3 = v3
                        var v37 string = "foo bar"
                        _ = v37
                    }
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    fmt.Println(v9, v2)
                    _ = v36
                } else {
                    if len(v6) > 0 {
                        v6[0] = 0
                    }
                    fmt.Println(v2)
                    switch 476 {
                    case 42:
                        v1 = -63
                    default:
                        var v38 uint64 = 374
                        _ = v38
                    }
                }
                v39 := GenS7{F0: -2147483648}
                v4 = append(v4, fi34)
                v40 := GenS8{F0: 4294967295}
                if (v3 < (1.5 + v3)) {
                    if ((v1 - (v1 ^ (v1 >> ((v33) & 63)))) >= v1) {
                        v1 -= 42
                        var v41 bool = (int(v33) == 117)
                        if (533 <= 18446744073709551615) {
                            var v42 int64 = -9223372036854775808
                            _ = v42
                        }
                        _ = v41
                    } else {
                        v43 := GenS9{F0: -163, F1: 18446744073709551615, F2: (v1 & v1)}
                        v9 = ((v2 || v9) && true)
                        v4 = append(v4, (-(fi34)))
                        _ = v43
                    }
                }
                if (42 > 127) {
                    var v44 uint16 = 47
                    fmt.Println(fi32, v3, fi32)
                    _ = v44
                } else {
                    var v45 uint32 = 0
                    _ = v45
                }
                var v46 bool = v9
                v47 := GenS8{F0: 1}
                _ = fi34
                _ = v35
                _ = v39
                _ = v40
                _ = v46
                _ = v47
            }
            var v48 float64 = (((v3 - 0.0) * v3) + 0.1)
            _ = v48
        } else {
            fmt.Println(v9, v2)
        }
        _ = fi32
        _ = v33
    }
    v2 = ("a" > ((("hello" + "a") + "") + (("hello" + (("foo bar" + "") + "a")) + "foo bar")))
    v49 := GenS8{F0: 65535}
    for fi50 := 0; fi50 < 18; fi50++ {
        switch fi50 {
        case 1:
            var v51 float64 = v3
            _ = v51
        case -815:
            if len(v6) > 0 {
                v6[0] = 1
            }
        }
        _ = fi50
    }
    switch 2147483647 {
    case 42:
        var v52 int8 = 42
        _ = v52
    case 240:
        var v53 uint8 = 127
        _ = v53
    }
    if (127 == 128) {
        fmt.Println(v3, v1, v3)
    }
    if len(v4) > 0 {
        v4[0] = 0
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v6
    _ = v9
}
