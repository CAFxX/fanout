package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string }
type GenS2 struct { F0 uint16; F1 int8 }
type GenS3 struct { F0 uint32; F1 uint64; F2 uint }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int; F1 int64 }
type GenS6 struct { F0 uint32; F1 string }
func genFn0(p0 int16, p1 string, p2 string) int16 {
    v4 := GenS1{F0: 94, F1: "hello"}
    p0 = -32768
    p0 -= 1000
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return -1
}
func genFn1(p0 int) int32 {
    if (1e10 != 0.1) {
        var v5 uint16 = 0
        _ = v5
    }
    p0 = p0
    _ = p0
    return 173
}

func main() {
    var v1 bool = true
    var v2 uint64 = 157
    var v3 float64 = (float64(-2.25) * float64(3.14159))
    v6 := GenS2{F0: 0, F1: 110}
    if (!(3.14159 > -2.25)) {
        fmt.Println(v1)
        v7 := []int{9223372036854775807, 32767, 42, 32768}
        v8 := GenS3{F0: 2147483647, F1: v2, F2: 255}
        _ = v7
        _ = v8
    } else {
        v3 -= ((float64(((((((1e10 / -2.25) * (1.5 + 1e10)) * 0.0) - 0.1) - 1e10) + 1e10)) * v3) + -2.25)
    }
    v2 = (^(v2))
    if v1 {
        v2 = v2
        var v9 uint8 = 0
        v10 := []int{-21}
        _ = v10
        _ = v9
    } else {
        if (255 <= 926) {
            for fi11 := 0; fi11 < 14; fi11++ {
                v12 := []int{2147483647, (fi11 >> ((241) & 63)), (fi11 + fi11)}
                _ = fi11
                _ = v12
            }
        } else {
            var v13 int64 = 42
            _ = v13
        }
        switch ("" + ("foo bar" + "a")) {
        case "":
            var v14 int8 = -99
            _ = v14
        case "hello":
            switch -9223372036854775808 {
            case 9223372036854775807:
                var v15 int64 = -1
                _ = v15
            }
        }
        var v16 uint = 518
        _ = v16
    }
    var v17 uint8 = 1
    v18 := []int{9223372036854775807, 128, 9223372036854775807}
    v18 = append(v18, -516)
    v1 = v1
    for fi19 := 0; fi19 < 9; fi19++ {
        switch "\n\t" {
        case "foo bar":
            v20 := GenS2{F0: 237, F1: 38}
            _ = v20
        default:
            for fi21 := 0; fi21 < 8; fi21++ {
                switch "hello" {
                case "hello":
                    v18 = append(v18, fi21)
                case "a":
                    if v1 {
                        if (int32(fi19) <= 2147483647) {
                            v22 := GenS1{F0: 221, F1: "\n\t"}
                            v2 = (^(v2))
                            _ = v22
                        }
                        var v23 uint8 = (v17 >> ((1) & 63))
                        _ = v23
                    } else {
                        var v24 uint64 = 2
                        _ = v24
                    }
                default:
                    if v1 {
                        if true {
                            var v25 int = fi19
                            v26 := GenS2{F0: 255, F1: -18}
                            _ = v25
                            _ = v26
                        }
                    } else {
                        if true {
                            if (v17 == v17) {
                                if len(v18) > 0 {
                                    v18[0] = fi19
                                }
                                if true {
                                    var v27 int16 = -1
                                    if v1 {
                                        var v28 string = "hello"
                                        var v29 int = (fi21 % (((fi21 / ((fi19) | 1))) | 1))
                                        _ = v28
                                        _ = v29
                                    } else {
                                        if (1 <= 100) {
                                            v30 := GenS3{F0: 4294967295, F1: v2, F2: 42}
                                            if v1 {
                                                var v31 float32 = float32(1e10)
                                                v27 = v27
                                                _ = v31
                                            }
                                            v18 = append(v18, -565)
                                            _ = v30
                                        } else {
                                            v32 := GenS3{F0: 2147483648, F1: uint64(v17), F2: 65535}
                                            var v33 uint32 = 2147483647
                                            v34 := GenS1{F0: v17, F1: (("foo bar" + "\n\t") + "a")}
                                            _ = v32
                                            _ = v33
                                            _ = v34
                                        }
                                    }
                                    _ = v27
                                } else {
                                    var v35 int32 = 2147483647
                                    v17 -= v17
                                    _ = v35
                                }
                                if len(v18) >= 2 {
                                    v18 = v18[1:len(v18)]
                                }
                                fmt.Println(fi21, v17)
                                v17 -= v17
                                var v36 int = (^(fi21))
                                _ = len(v18)
                                v2 += v2
                                v17 = 42
                                _ = v36
                            } else {
                                fmt.Println(fi21, v17, v3)
                            }
                            var v37 uint = 18446744073709551615
                            if v1 {
                                v2 -= (v2 >> ((v37) & 63))
                            }
                            _ = v37
                        }
                        var v38 int64 = 0
                        _ = v38
                    }
                }
                if (v3 > ((0.1 * 0.1) - 1.5)) {
                    if ((!v1) || false) {
                        if v1 {
                            var v39 string = "hello"
                            _ = v39
                        } else {
                            if (1 == int32(v2)) {
                                fmt.Println(fi19)
                                var v40 uint16 = 65535
                                v41 := GenS4{F0: 32768}
                                _ = v40
                                _ = v41
                            } else {
                                if len(v18) >= 2 {
                                    v18 = v18[1:len(v18)]
                                }
                                var v42 int = fi21
                                v1 = (1 == 65535)
                                _ = v42
                            }
                            if v1 {
                                fmt.Println(fi21, fi21)
                                if len(v18) > 0 {
                                    v18[0] = -43
                                }
                            }
                            v17 += 128
                        }
                        v43 := GenS5{F0: -9223372036854775808, F1: 916}
                        switch (fi19 & -572) {
                        case 9223372036854775807:
                            var v44 uint = 0
                            _ = v44
                        case -9223372036854775808:
                            v18 = append(v18, fi21)
                        case 42:
                            fmt.Println(fi21)
                        }
                        _ = v43
                    } else {
                        var v45 uint = 479
                        v46 := GenS4{F0: fi19}
                        v17 = v17
                        _ = v45
                        _ = v46
                    }
                    v2 -= 1
                } else {
                    v47 := GenS5{F0: -316, F1: 0}
                    switch "\n\t" {
                    case "hello":
                        v48 := GenS6{F0: 987, F1: ""}
                        _ = v48
                    }
                    _ = v47
                }
                _ = fi21
            }
        }
        _ = fi19
    }
    v2 -= v2
    for fi49 := 0; fi49 < 13; fi49++ {
        if (9223372036854775807 > 0) {
            _ = len(v18)
            var v50 float32 = 0.1
            var v51 uint16 = 42
            _ = v50
            _ = v51
        } else {
            var v52 uint = 0
            if len(v18) > 0 {
                v18[0] = (fi49 << ((906) & 63))
            }
            _ = v52
        }
        var v53 uint = uint(v2)
        var v54 float32 = 1e10
        _ = fi49
        _ = v53
        _ = v54
    }
    v55 := GenS6{F0: 0, F1: "\n\t"}
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v55
    _ = v6
}
