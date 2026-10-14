package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int32; F1 int32; F2 int64 }
type GenS3 struct { F0 bool; F1 int; F2 int16 }
type GenS4 struct { F0 int8 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int8 = 46
    v3 := make(map[uint32]string)
    var v4 string = (("\n\t" + "a") + "hello")
    var v5 int = -2147483648
    v6 := GenS1{F0: v4}
    var v7 uint32 = 0
    var v8 float32 = 0.1
    v9 := v1[v7]
    var v10 int16 = -1
    if (!(127 <= (-(v2)))) {
        v2 = 100
    } else {
        if false {
            if false {
                v5 += v5
                v11 := GenS2{F0: 42, F1: 0, F2: -9223372036854775808}
                for fi12 := 0; fi12 < 11; fi12++ {
                    var v13 int32 = -1
                    _ = fi12
                    _ = v13
                }
                _ = v11
            } else {
                var v14 int32 = 0
                if false {
                    if false {
                        for fi15 := 0; fi15 < 13; fi15++ {
                            fmt.Println(v4, fi15, fi15)
                            if (4294967295 < (v7 | (^(v9)))) {
                                var v16 int8 = v2
                                var v17 int8 = v2
                                _ = v16
                                _ = v17
                            }
                            _ = len(v1)
                            _ = fi15
                        }
                        fmt.Println(v14, v7)
                    }
                    switch v4 {
                    case "foo bar":
                        v18 := GenS2{F0: 0, F1: 42, F2: 0}
                        _ = v18
                    case "a":
                        _ = len(v3)
                    case "hello":
                        var v19 uint32 = (v9 / ((v7) | 1))
                        _ = v19
                    }
                } else {
                    if (76 >= 18446744073709551615) {
                        if (!((-9223372036854775808 != -690) && (270 <= 373))) {
                            for fi20 := 0; fi20 < 2; fi20++ {
                                var v21 int = int(v2)
                                for fi22 := 0; fi22 < 12; fi22++ {
                                    var v23 float32 = v8
                                    var v24 int = (fi20 << ((4294967295) & 63))
                                    _ = fi22
                                    _ = v23
                                    _ = v24
                                }
                                _ = fi20
                                _ = v21
                            }
                        }
                        for fi25 := 0; fi25 < 17; fi25++ {
                            var v26 uint32 = v9
                            if false {
                                switch (v4 + v4) {
                                case "":
                                    var v27 uint16 = 1
                                    _ = v27
                                case "a":
                                    fmt.Println(v5, v14, v5)
                                case "foo bar":
                                    v7 -= 2147483647
                                default:
                                    v7 = v26
                                }
                                v4 = (((v4 + "") + v4) + ((((((("hello" + "foo bar") + ("\n\t" + v4)) + "") + (("foo bar" + v4) + "a")) + v4) + ((("hello" + "\n\t") + (v4 + v4)) + "\n\t")) + v4))
                                v26 -= (v7 ^ (v7 & v26))
                            }
                            v3[v9] = (v4 + "\n\t")
                            _ = fi25
                            _ = v26
                        }
                        var v28 int16 = (-(v10))
                        _ = v28
                    } else {
                        fmt.Println(v4, v7, v9)
                        v29 := GenS3{F0: true, F1: (v5 >> ((391) & 63)), F2: 1000}
                        if true {
                            v5 -= (^(v5))
                            for fi30 := 0; fi30 < 16; fi30++ {
                                v10 -= v10
                                var v31 int = (v5 >> ((18446744073709551615) & 63))
                                _ = fi30
                                _ = v31
                            }
                        } else {
                            var v32 uint8 = 42
                            var v33 int32 = (-(v14))
                            var v34 int8 = -1
                            _ = v32
                            _ = v33
                            _ = v34
                        }
                        _ = v29
                    }
                    v4 = (("foo bar" + v4) + v4)
                }
                switch v4 {
                case "\n\t":
                    v35 := GenS3{F0: (42 >= 732), F1: v5, F2: (v10 * v10)}
                    _ = v35
                case "foo bar":
                    v8 -= ((v8 / v8) / (1.5 / 0.1))
                case "a":
                    if (!((294 > 42) || ((float32(-2.25) > v8) || false))) {
                        v36 := GenS2{F0: v14, F1: v14, F2: 1}
                        for fi37 := 0; fi37 < 15; fi37++ {
                            v2 -= v2
                            v38 := GenS1{F0: (v4 + "a")}
                            if (false && false) {
                                for fi39 := 0; fi39 < 7; fi39++ {
                                    v4 = v4
                                    v40 := GenS1{F0: v4}
                                    _ = fi39
                                    _ = v40
                                }
                                for fi41 := 0; fi41 < 20; fi41++ {
                                    fmt.Println(v9)
                                    _ = fi41
                                }
                            } else {
                                v7 = (v7 << ((2) & 63))
                                var v42 int64 = 9223372036854775807
                                _ = v42
                            }
                            _ = fi37
                            _ = v38
                        }
                        if (707 == 4294967295) {
                            v43 := []int{v5, v5, (v5 * int(v2))}
                            if (float32((((0.0 * 3.14159) + -2.25) + ((0.1 + ((float64(0.1) * 1.5) * 0.1)) - 1.5))) >= v8) {
                                switch v4 {
                                case "\n\t":
                                    var v44 float64 = (3.14159 / 1.0)
                                    _ = v44
                                case "":
                                    switch v4 {
                                    case "":
                                        v45 := v3[v7]
                                        _ = v45
                                    default:
                                        var v46 int32 = v14
                                        _ = v46
                                    }
                                case "a":
                                    var v47 bool = true
                                    _ = v47
                                }
                            } else {
                                v8 = ((v8 * ((((v8 - v8) - v8) / v8) - v8)) + 3.14159)
                            }
                            _ = v43
                        }
                        _ = v36
                    } else {
                        switch (v4 + v4) {
                        case "a":
                            var v48 int64 = 140
                            _ = v48
                        case "foo bar":
                            v4 = (v4 + v4)
                        }
                    }
                }
                _ = v14
            }
            v7 = v7
            v49 := GenS3{F0: (v8 >= 0.1), F1: (-(v5)), F2: 749}
            _ = v49
        } else {
            for fi50 := 0; fi50 < 8; fi50++ {
                v51 := GenS4{F0: 100}
                v52 := []int{}
                _ = fi50
                _ = v51
                _ = v52
            }
        }
        fmt.Println(v9, v9, v8)
        var v53 uint = 795
        _ = v53
    }
    var v54 float32 = ((v8 / -2.25) / 1.0)
    if false {
        v8 = ((v54 + (v54 / v54)) / (v8 * (v54 / 0.1)))
        for fi55 := 0; fi55 < 12; fi55++ {
            v10 -= 1
            v4 = ((v4 + v4) + "hello")
            var v56 bool = (!(!true))
            _ = fi55
            _ = v56
        }
        v57 := []int{-938, 1}
        _ = v57
    } else {
        v58 := GenS2{F0: int32(v5), F1: 2147483647, F2: -9223372036854775808}
        v10 = v10
        _ = v58
    }
    v8 += ((((float32(3.14159) - (-2.25 / v54)) - ((0.1 / ((v8 + (0.1 * (v54 - v8))) + (0.0 / 1e10))) + (0.0 + v8))) * 0.1) - v54)
    for fi59 := 0; fi59 < 11; fi59++ {
        for fi60 := 0; fi60 < 10; fi60++ {
            var v61 int = v5
            _ = fi60
            _ = v61
        }
        v4 = (v4 + "")
        delete(v1, v7)
        _ = fi59
    }
    v10 -= 32767
    for fi62 := 0; fi62 < 3; fi62++ {
        for fi63 := 0; fi63 < 5; fi63++ {
            v64 := GenS3{F0: true, F1: 50, F2: v10}
            var v65 uint8 = 154
            v4 = v4
            _ = fi63
            _ = v64
            _ = v65
        }
        _ = fi62
    }
    v7 = v7
    v66 := v3[v9]
    var v67 float64 = 1e10
    if false {
        v4 = v4
        fmt.Println(v4, v67)
        fmt.Println(v10, v9, v7)
    } else {
        v10 += v10
        switch 2147483647 {
        case 128:
            for fi68 := 0; fi68 < 3; fi68++ {
                if (770 <= 351) {
                    v9 = v9
                }
                _ = fi68
            }
        case -1:
            fmt.Println(v67, v66)
        case 4294967295:
            v69 := v1[1]
            _ = v69
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v54:", v54)
    fmt.Println("v66:", v66)
    fmt.Println("v67:", v67)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v54
    _ = v6
    _ = v66
    _ = v67
    _ = v7
    _ = v8
    _ = v9
}
