package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8; F2 int32 }
type GenS2 struct { F0 uint; F1 string }
type GenS3 struct { F0 int8; F1 int32 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 string; F1 string; F2 int16 }
func genFn0(p0 uint8) int {
    var v6 int = -587
    p0 = p0
    v6 += (v6 / ((v6) | 1))
    _ = p0
    _ = v6
    return 127
}
func genFn1(p0 bool, p1 string) uint16 {
    if p0 {
        v7 := []int{4294967295}
        if p0 {
            var v8 bool = p0
            var v9 string = "hello"
            _ = v8
            _ = v9
        } else {
            if (!false) {
                if p0 {
                    if p0 {
                        var v10 int16 = 1
                        _ = len(v7)
                        _ = v10
                    }
                } else {
                    if (42 != 660) {
                        if len(v7) > 0 {
                            v7[0] = 32768
                        }
                        var v11 float32 = float32(float64(0.1))
                        v12 := GenS1{F0: (!p0), F1: 1, F2: -180}
                        _ = v11
                        _ = v12
                    }
                }
                for fi13 := 0; fi13 < 9; fi13++ {
                    v7 = append(v7, 32767)
                    if len(v7) > 0 {
                        v7[0] = genFn0(1)
                    }
                    if (127 != 220) {
                        for fi14 := 0; fi14 < 4; fi14++ {
                            if len(v7) > 0 {
                                v7[0] = fi13
                            }
                            v15 := GenS2{F0: 1, F1: "foo bar"}
                            _ = fi14
                            _ = v15
                        }
                        v7 = append(v7, (^(fi13)))
                    }
                    _ = fi13
                }
                var v16 bool = (0.1 <= 0.0)
                _ = v16
            } else {
                switch -952 {
                case 298:
                    var v17 int = 209
                    _ = v17
                case 4294967295:
                    fmt.Println(p0, p0)
                case -1:
                    fmt.Println(p0, p1)
                }
                fmt.Println(p0)
                var v18 uint = 514
                _ = v18
            }
        }
        if false {
            var v19 int64 = 42
            _ = v19
        } else {
            if false {
                if (!(p0 && p0)) {
                    p1 = (((p1 + "a") + "\n\t") + (((p1 + "a") + p1) + "a"))
                    var v20 float64 = (3.14159 + (1.5 / 3.14159))
                    if true {
                        if false {
                            v21 := GenS1{F0: false, F1: -115, F2: 2147483647}
                            v22 := GenS2{F0: 42, F1: "\n\t"}
                            _ = v21
                            _ = v22
                        } else {
                            v23 := GenS2{F0: 6, F1: "foo bar"}
                            var v24 int16 = -1
                            switch 709 {
                            case 9223372036854775807:
                                p0 = (p1 < p1)
                            case 1:
                                var v25 int8 = 127
                                _ = v25
                            case -9223372036854775808:
                                fmt.Println(p1, p0)
                            default:
                                v20 += ((v20 * v20) / v20)
                            }
                            _ = v23
                            _ = v24
                        }
                        var v26 int = 9223372036854775807
                        _ = v26
                    }
                    _ = v20
                } else {
                    var v27 int8 = -12
                    if (226 > 0) {
                        fmt.Println(v27, v27, v27)
                    } else {
                        var v28 int8 = v27
                        _ = v28
                    }
                    if len(v7) > 0 {
                        v7[0] = 42
                    }
                    _ = v27
                }
            }
            for fi29 := 0; fi29 < 12; fi29++ {
                switch (("a" + p1) + (("a" + ((p1 + ("a" + p1)) + p1)) + "foo bar")) {
                case "\n\t":
                    var v30 uint8 = 54
                    _ = v30
                case "hello":
                    v31 := GenS2{F0: 490, F1: "\n\t"}
                    _ = v31
                }
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
                p1 = ("a" + p1)
                _ = fi29
            }
        }
        _ = v7
    }
    p1 = (("a" + p1) + ("" + p1))
    _ = p0
    _ = p1
    return 2
}

func main() {
    var v1 uint = 4294967295
    var v2 bool = ((2147483648 != 4294967295) || true)
    v3 := make(map[int]uint32)
    v4 := []uint32{413}
    var v5 string = "a"
    for fi32 := 0; fi32 < 14; fi32++ {
        v2 = v2
        switch genFn0(165) {
        case -9223372036854775808:
            if (0.0 == 1.5) {
                for fi33 := 0; fi33 < 1; fi33++ {
                    v34 := GenS1{F0: v2, F1: 43, F2: 390}
                    v5 = "hello"
                    _ = fi33
                    _ = v34
                }
            } else {
                v1 = v1
                if false {
                    v35 := GenS2{F0: (v1 % (((v1 % ((42) | 1))) | 1)), F1: v5}
                    var v36 uint8 = 1
                    v36 = (v36 >> ((v1) & 63))
                    _ = v35
                    _ = v36
                }
                switch v5 {
                case "":
                    var v37 float64 = (1.5 / 1.5)
                    _ = v37
                case "\n\t":
                    v3[genFn0(128)] = 798
                default:
                    v5 = v5
                }
            }
        }
        _ = fi32
    }
    fmt.Println(v2)
    var v38 int = -832
    var v39 uint = (v1 + 65535)
    fmt.Println(v2, v39)
    var v40 uint32 = 1
    for fi41 := 0; fi41 < 10; fi41++ {
        var v42 int64 = 0
        _ = fi41
        _ = v42
    }
    switch (v38 << (((v1 >> (((v1 | 60)) & 63))) & 63)) {
    case 32768:
        if false {
            if (2147483647 == -1) {
                for fi43 := 0; fi43 < 9; fi43++ {
                    v44 := GenS3{F0: 42, F1: -2147483648}
                    switch genFn0(128) {
                    case 127:
                        var v45 string = (v5 + v5)
                        _ = v45
                    default:
                        fmt.Println(v40, v2, v2)
                    }
                    v46 := GenS4{F0: -32768}
                    _ = fi43
                    _ = v44
                    _ = v46
                }
            }
            delete(v3, (v38 % ((4294967295) | 1)))
            fmt.Println(v39, v1, v2)
        }
    case -9223372036854775808:
        v39 += (v39 - (^(v1)))
    case 1:
        var v47 uint8 = 1
        _ = v47
    }
    v2 = (4294967295 == 2)
    switch (v38 ^ genFn0(42)) {
    case 358:
        _ = len(v3)
    case 32767:
        for fi48 := 0; fi48 < 4; fi48++ {
            delete(v3, 42)
            _ = fi48
        }
    }
    switch (v5 + (v5 + v5)) {
    case "\n\t":
        if false {
            v39 = v39
            v49 := GenS5{F0: v5, F1: v5, F2: 0}
            _ = v49
        }
    case "":
        switch (v5 + ((v5 + v5) + "\n\t")) {
        case "foo bar":
            if ((82 > 1) || false) {
                if v2 {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    v1 -= (^(v39))
                } else {
                    var v50 uint16 = 2
                    _ = v50
                }
                v40 = 2147483648
            }
        }
    default:
        var v51 uint8 = 127
        _ = v51
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v5
}
