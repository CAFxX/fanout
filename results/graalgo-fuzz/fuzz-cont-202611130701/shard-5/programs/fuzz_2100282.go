package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8; F2 uint16 }
type GenS2 struct { F0 int64; F1 int16; F2 int }
type GenS3 struct { F0 uint32; F1 int64; F2 int32 }
func genFn0(p0 uint64, p1 int32, p2 uint8) uint32 {
    v7 := GenS1{F0: 42, F1: p2, F2: 2}
    v8 := GenS1{F0: 0, F1: p2, F2: 2}
    p0 = 1
    v9 := []int{4294967295, 269, 2147483647, 9223372036854775807}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return 2
}
func genFn1(p0 bool, p1 uint8) uint {
    fmt.Println(p1)
    var v10 uint16 = 380
    v10 = 0
    v11 := []int{2147483647, 32768, 128}
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return 4294967295
}

func main() {
    v1 := make(map[string]string)
    var v2 float64 = (((3.14159 - 0.0) / float64(3.14159)) * (((3.14159 / 1.0) + 3.14159) + float64(1.5)))
    var v3 int64 = 1
    v4 := make(map[string]string)
    var v5 uint = 2
    var v6 uint64 = 0
    var v12 float64 = v2
    var v13 uint = v5
    if (524 != 2147483647) {
        for fi14 := 0; fi14 < 6; fi14++ {
            for fi15 := 0; fi15 < 7; fi15++ {
                var v16 int32 = 1
                switch (fi14 * (fi15 & fi15)) {
                case 2147483647:
                    if ((0.0 != v2) || (v13 < (v5 - (v5 % (((v13 + (v13 | 42))) | 1))))) {
                        v13 = uint(v6)
                        var v17 int16 = 1000
                        delete(v4, "hello")
                        v5 = (v13 - genFn1(false, 127))
                        if (16 != 100) {
                            var v18 uint16 = 0
                            var v19 float64 = v12
                            v6 = (^(v6))
                            _ = v18
                            _ = v19
                        } else {
                            v3 -= (-(v3))
                        }
                        _ = v17
                    }
                case 42:
                    v13 = (v5 << (((v5 - 0)) & 63))
                default:
                    v6 = 4294967295
                }
                _ = fi15
                _ = v16
            }
            _ = fi14
        }
        if (((1.5 - 0.1) <= (0.0 + v12)) || false) {
            if false {
                v5 -= genFn1(true, 42)
                var v20 uint64 = (v6 / ((v6) | 1))
                _ = v20
            } else {
                v21 := GenS2{F0: 0, F1: -806, F2: 128}
                for fi22 := 0; fi22 < 7; fi22++ {
                    if false {
                        var v23 int16 = 42
                        var v24 float64 = v2
                        _ = v23
                        _ = v24
                    }
                    _ = fi22
                }
                _ = v21
            }
            switch "\n\t" {
            case "\n\t":
                switch -595 {
                case -9223372036854775808:
                    switch int(v6) {
                    case 4294967295:
                        var v25 int16 = -32768
                        _ = v25
                    case -2147483648:
                        var v26 float64 = (v12 - ((v2 - (float64((-2.25 / 1e10)) / 1.0)) - 1.5))
                        _ = v26
                    case 948:
                        delete(v4, "hello")
                    }
                case 2147483647:
                    if true {
                        var v27 uint = (^(v13))
                        var v28 uint8 = 215
                        switch -9223372036854775808 {
                        case 4294967295:
                            v29 := []int{9223372036854775807, 9223372036854775807, 274}
                            _ = v29
                        case -753:
                            fmt.Println(v6, v13, v27)
                        case 42:
                            var v30 float64 = 3.14159
                            _ = v30
                        }
                        _ = v27
                        _ = v28
                    } else {
                        switch "\n\t" {
                        case "\n\t":
                            var v31 uint32 = genFn0(1, -2147483648, 0)
                            _ = v31
                        default:
                            for fi32 := 0; fi32 < 18; fi32++ {
                                v33 := GenS1{F0: 743, F1: 128, F2: 380}
                                _ = fi32
                                _ = v33
                            }
                        }
                    }
                case -270:
                    v3 = v3
                default:
                    v1["\n\t"] = (("hello" + (("a" + "\n\t") + "foo bar")) + "")
                }
            default:
                v12 += 1e10
            }
            fmt.Println(v5, v5, v6)
        } else {
            v2 = 0.0
        }
    }
    if true {
        v2 -= 1.5
    } else {
        var v34 int32 = -1
        var v35 uint16 = 1
        _ = v34
        _ = v35
    }
    for fi36 := 0; fi36 < 9; fi36++ {
        if (0.0 < (v12 / v2)) {
            var v37 float32 = 0.1
            _ = v37
        } else {
            v38 := GenS3{F0: 8, F1: v3, F2: -1}
            for fi39 := 0; fi39 < 3; fi39++ {
                var v40 int16 = 1
                _ = fi39
                _ = v40
            }
            _ = v38
        }
        _ = fi36
    }
    var v41 int64 = (v3 | 9223372036854775807)
    for fi42 := 0; fi42 < 6; fi42++ {
        if true {
            for fi43 := 0; fi43 < 9; fi43++ {
                var v44 bool = false
                v2 = 1.5
                var v45 int = fi43
                _ = fi43
                _ = v44
                _ = v45
            }
            for fi46 := 0; fi46 < 3; fi46++ {
                if (100 <= 127) {
                    var v47 string = "\n\t"
                    _ = v47
                } else {
                    v12 = float64(float32(v12))
                    var v48 uint16 = 42
                    if (!(v48 > v48)) {
                        v3 = int64(v6)
                    }
                    v1["a"] = ""
                    var v49 uint16 = v48
                    _ = v48
                    _ = v49
                }
                var v50 string = "\n\t"
                v51 := v4[v50]
                var v52 uint8 = 94
                _ = fi46
                _ = v50
                _ = v51
                _ = v52
            }
            var v53 uint32 = 4294967295
            _ = v53
        }
        _ = fi42
    }
    fmt.Println(v3, v2)
    var v54 string = ((("hello" + ("\n\t" + "hello")) + (("hello" + ("foo bar" + "a")) + "\n\t")) + "foo bar")
    v6 = (v6 >> ((2) & 63))
    var v55 int = 32767
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v5
    _ = v54
    _ = v55
    _ = v6
}
