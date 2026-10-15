package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int16; F1 bool; F2 int8 }
type GenS3 struct { F0 uint; F1 bool; F2 int8 }
type GenS4 struct { F0 int; F1 uint64; F2 int }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int64 }
func genFn0(p0 bool) uint32 {
    for fi6 := 0; fi6 < 11; fi6++ {
        p0 = true
        var v7 bool = p0
        _ = fi6
        _ = v7
    }
    v8 := []int{-1, 32767, -387, -9223372036854775808}
    p0 = p0
    _ = p0
    _ = v8
    return 65535
}

func main() {
    var v1 string = "\n\t"
    var v2 bool = (true && (!false))
    var v3 int8 = 58
    var v4 int8 = 100
    var v5 int16 = -32768
    fmt.Println(v1)
    for fi9 := 0; fi9 < 5; fi9++ {
        v4 = (^(v4))
        _ = fi9
    }
    v10 := []int{9223372036854775807, 187}
    for fi11 := 0; fi11 < 11; fi11++ {
        fmt.Println(v4, fi11, v1)
        var v12 float32 = float32(0.1)
        _ = fi11
        _ = v12
    }
    v10 = append(v10, 128)
    v10 = append(v10, 128)
    switch "a" {
    case "foo bar":
        v4 = (v3 * -1)
    default:
        v10 = append(v10, 4294967295)
    }
    if len(v10) > 0 {
        v10[0] = 4294967295
    }
    v10 = append(v10, 32768)
    for fi13 := 0; fi13 < 3; fi13++ {
        if v2 {
            fmt.Println(v4)
            if v2 {
                var v14 int32 = 1
                var v15 uint = 672
                _ = v14
                _ = v15
            } else {
                if v2 {
                    switch fi13 {
                    case 688:
                        v1 = v1
                    case -446:
                        var v16 int16 = (v5 << ((2) & 63))
                        _ = v16
                    }
                } else {
                    for fi17 := 0; fi17 < 1; fi17++ {
                        v18 := GenS1{F0: 42}
                        var v19 int64 = 1
                        if (!v2) {
                            var v20 int16 = 825
                            var v21 float32 = (float32((1.5 - (-2.25 + 1.5))) * ((1.5 * (3.14159 * 0.1)) - 1e10))
                            if len(v10) >= 2 {
                                v10 = v10[1:len(v10)]
                            }
                            _ = v20
                            _ = v21
                        }
                        if false {
                            switch (("\n\t" + (("foo bar" + "a") + "foo bar")) + "foo bar") {
                            case "foo bar":
                                var v22 bool = (v2 || ((153 == (fi13 % ((746) | 1))) || (2 < genFn0(true))))
                                _ = v22
                            case "":
                                v2 = true
                            case "\n\t":
                                v19 = (v19 * v19)
                            }
                        }
                        fmt.Println(v3)
                        _ = fi17
                        _ = v18
                        _ = v19
                    }
                    if false {
                        v4 -= v3
                        var v23 float64 = 0.0
                        for fi24 := 0; fi24 < 17; fi24++ {
                            v23 -= 1.5
                            var v25 int8 = v4
                            var v26 uint = 42
                            _ = fi24
                            _ = v25
                            _ = v26
                        }
                        _ = v23
                    } else {
                        fmt.Println(v1, v2, v2)
                        switch "foo bar" {
                        case "\n\t":
                            v4 = v3
                        case "foo bar":
                            if v2 {
                                for fi27 := 0; fi27 < 9; fi27++ {
                                    v5 -= (v5 - 1000)
                                    v28 := GenS1{F0: 133}
                                    if len(v10) > 0 {
                                        v10[0] = 4294967295
                                    }
                                    _ = fi27
                                    _ = v28
                                }
                                var v29 float64 = (1.5 - (0.1 / 1.0))
                                _ = v29
                            }
                        case "":
                            v2 = v2
                        default:
                            v2 = ((0.1 / float64(-2.25)) == (-2.25 - 3.14159))
                        }
                    }
                }
                v30 := GenS1{F0: 42}
                var v31 float32 = (0.1 / -2.25)
                _ = v30
                _ = v31
            }
            if len(v10) > 0 {
                v10[0] = (fi13 << ((65535) & 63))
            }
        } else {
            if v2 {
                switch (fi13 / ((-2147483648) | 1)) {
                case -520:
                    var v32 uint8 = 8
                    _ = v32
                default:
                    if (!v2) {
                        if len(v10) >= 2 {
                            v10 = v10[1:len(v10)]
                        }
                        switch -199 {
                        case 0:
                            v33 := GenS2{F0: (^(v5)), F1: (176 == -367), F2: v3}
                            _ = v33
                        case 264:
                            for fi34 := 0; fi34 < 11; fi34++ {
                                fmt.Println(v1, v4)
                                _ = fi34
                            }
                        case 4294967295:
                            var v35 float32 = (((-2.25 + (3.14159 / -2.25)) + -2.25) + 1.5)
                            _ = v35
                        }
                    }
                }
                switch ((v1 + v1) + ("foo bar" + v1)) {
                case "\n\t":
                    for fi36 := 0; fi36 < 20; fi36++ {
                        if (v1 == "a") {
                            switch (("" + v1) + v1) {
                            case "foo bar":
                                if (v4 >= v3) {
                                    fmt.Println(v2)
                                } else {
                                    v1 = ("foo bar" + (v1 + v1))
                                    switch ("foo bar" + ("\n\t" + v1)) {
                                    case "hello":
                                        if v2 {
                                            fmt.Println(fi13, v3)
                                            v1 = (("foo bar" + (v1 + v1)) + v1)
                                        }
                                    case "":
                                        v3 = v4
                                        v1 = ((v1 + v1) + "hello")
                                        var v37 string = ((v1 + "a") + v1)
                                        _ = v37
                                    }
                                    fmt.Println(v1, v1)
                                }
                            case "\n\t":
                                var v38 string = v1
                                _ = v38
                            default:
                                v39 := GenS2{F0: (v5 - v5), F1: (128 == (^(fi13))), F2: (v3 - (v3 << ((65535) & 63)))}
                                _ = v39
                            }
                            if len(v10) >= 2 {
                                v10 = v10[1:len(v10)]
                            }
                        } else {
                            if len(v10) > 0 {
                                v10[0] = (fi13 ^ (fi13 & (fi13 << ((847) & 63))))
                            }
                            if v2 {
                                fmt.Println(v2, v2)
                                v40 := GenS3{F0: 255, F1: (1 == -1), F2: int8(fi13)}
                                v41 := GenS4{F0: -1, F1: 38, F2: (fi13 ^ (fi13 % ((-9223372036854775808) | 1)))}
                                var v42 bool = v2
                                switch v1 {
                                case "":
                                    v2 = v42
                                case "\n\t":
                                    v1 = ("a" + ("" + v1))
                                case "a":
                                    var v43 uint8 = 164
                                    _ = v43
                                default:
                                    var v44 uint32 = 1
                                    _ = v44
                                }
                                _ = v40
                                _ = v41
                                _ = v42
                            } else {
                                if true {
                                    var v45 int16 = -497
                                    var v46 int = fi13
                                    v4 = v4
                                    _ = v45
                                    _ = v46
                                } else {
                                    var v47 int64 = -9223372036854775808
                                    var v48 int8 = v3
                                    if (628 == 1) {
                                        if len(v10) >= 2 {
                                            v10 = v10[1:len(v10)]
                                        }
                                        var v49 int16 = (v5 >> ((287) & 63))
                                        _ = v49
                                    }
                                    _ = v47
                                    _ = v48
                                }
                                if (369 < -1) {
                                    fmt.Println(fi36)
                                    if len(v10) > 0 {
                                        v10[0] = -706
                                    }
                                    var v50 uint16 = 2
                                    _ = v50
                                } else {
                                    fmt.Println(fi13, v4, v3)
                                }
                                v10 = append(v10, -9223372036854775808)
                            }
                            v51 := GenS4{F0: 2147483647, F1: 17, F2: 127}
                            _ = v51
                        }
                        fmt.Println(fi13, v5)
                        v52 := GenS5{F0: 1}
                        switch "foo bar" {
                        case "a":
                            var v53 int = fi13
                            _ = v53
                        default:
                            if (0.0 >= 1.5) {
                                switch "hello" {
                                case "a":
                                    v10 = append(v10, fi13)
                                case "hello":
                                    v3 = v3
                                }
                            }
                        }
                        if v2 {
                            if v2 {
                                var v54 uint32 = 417
                                var v55 int8 = v4
                                _ = v54
                                _ = v55
                            } else {
                                _ = len(v10)
                                if v2 {
                                    if (490 >= 240) {
                                        _ = len(v10)
                                        fmt.Println(v4, fi36, fi36)
                                    }
                                }
                                v2 = v2
                            }
                            fmt.Println(fi36)
                            var v56 uint = 65535
                            _ = v56
                        }
                        _ = fi36
                        _ = v52
                    }
                case "hello":
                    var v57 float64 = 0.1
                    _ = v57
                case "a":
                    v58 := GenS6{F0: 951}
                    _ = v58
                default:
                    var v59 float64 = (float64(-2.25) + ((1.5 + float64(0.1)) * 0.1))
                    _ = v59
                }
            } else {
                v1 = v1
            }
            switch v1 {
            case "\n\t":
                v1 = "\n\t"
            case "a":
                for fi60 := 0; fi60 < 11; fi60++ {
                    v2 = (fi13 != fi13)
                    _ = fi60
                }
            }
        }
        var v61 uint = 42
        v62 := GenS5{F0: 0}
        _ = fi13
        _ = v61
        _ = v62
    }
    var v63 bool = (42 >= 103)
    fmt.Println(v3)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v63
}
