package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int; F2 uint64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int; F1 int16; F2 int16 }
type GenS5 struct { F0 uint32; F1 int32; F2 uint8 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 uint16; F1 int16 }
type GenS8 struct { F0 string; F1 uint16 }
func genFn0(p0 int32) int64 {
    var v5 int32 = (p0 - (p0 - p0))
    var v6 int16 = 0
    v5 = -2147483648
    var v7 uint64 = 0
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return -9223372036854775808
}
func genFn1(p0 uint64, p1 string) uint32 {
    p1 = (p1 + "hello")
    v8 := GenS1{F0: 203, F1: -451, F2: p0}
    _ = p0
    _ = p1
    _ = v8
    return 761
}

func main() {
    var v1 int32 = 42
    var v2 float64 = (0.1 - 0.0)
    var v3 float32 = 1.5
    var v4 float32 = 1.5
    if true {
        v4 -= v3
        var v9 uint8 = 125
        _ = v9
    } else {
        switch 42 {
        case 457:
            var v10 int64 = 0
            _ = v10
        case 128:
            if ((209 <= 250) && (380 > 0)) {
                for fi11 := 0; fi11 < 8; fi11++ {
                    if false {
                        var v12 uint8 = 128
                        fmt.Println(v4)
                        v13 := []int{}
                        _ = v12
                        _ = v13
                    } else {
                        if true {
                            fmt.Println(fi11, v1)
                        }
                        v3 = ((v4 + v3) / (-2.25 + v3))
                        if true {
                            var v14 int = -9223372036854775808
                            var v15 int16 = 0
                            if false {
                                v16 := []int{int(v15), 4294967295}
                                v3 = (v4 + v3)
                                _ = v16
                            }
                            _ = v14
                            _ = v15
                        } else {
                            v4 = (((v3 - v3) + (((((0.0 - 1.5) / v3) * v3) * v4) / (v3 * (((((3.14159 + 0.0) / -2.25) / v3) * v4) / float32(v2))))) - v4)
                            var v17 int8 = 127
                            _ = v17
                        }
                    }
                    _ = fi11
                }
                for fi18 := 0; fi18 < 10; fi18++ {
                    v2 += -2.25
                    v19 := []int{}
                    _ = fi18
                    _ = v19
                }
                v20 := []int{9223372036854775807, -9223372036854775808, -2147483648}
                _ = v20
            }
        case 9223372036854775807:
            v21 := GenS1{F0: 4294967295, F1: 4294967295, F2: 1}
            _ = v21
        }
    }
    v3 = v4
    fmt.Println(v4, v2, v3)
    for fi22 := 0; fi22 < 11; fi22++ {
        var v23 uint16 = 32767
        for fi24 := 0; fi24 < 10; fi24++ {
            var v25 uint = 181
            if true {
                v26 := []int{fi24, (fi22 + fi22), -9223372036854775808, -642}
                v23 -= v23
                _ = v26
            } else {
                switch -1 {
                case 934:
                    switch "hello" {
                    case "\n\t":
                        var v27 int8 = -12
                        _ = v27
                    case "a":
                        if (genFn1(2, "a") == genFn1(4294967295, "hello")) {
                            switch "\n\t" {
                            case "hello":
                                v28 := GenS1{F0: 908, F1: 674, F2: 4294967295}
                                _ = v28
                            case "foo bar":
                                v4 = 0.1
                            default:
                                var v29 uint8 = 2
                                _ = v29
                            }
                            v30 := []int{152, (-(fi22))}
                            _ = v30
                        } else {
                            v31 := GenS2{F0: (!(0 > -851))}
                            if false {
                                fmt.Println(v4)
                                var v32 int8 = 121
                                _ = v32
                            } else {
                                v33 := GenS3{F0: 42}
                                var v34 uint64 = 4294967295
                                var v35 uint32 = 42
                                _ = v33
                                _ = v34
                                _ = v35
                            }
                            _ = v31
                        }
                    case "foo bar":
                        if ((v1 >= v1) && false) {
                            v2 += v2
                            var v36 float64 = (0.0 * v2)
                            _ = v36
                        }
                    }
                case 128:
                    v1 = v1
                default:
                    v25 = (v25 / ((v25) | 1))
                    if (695 == (v23 ^ v23)) {
                        v37 := []int{fi22}
                        _ = v37
                    }
                }
                if (((0.1 / 3.14159) * (1e10 + v2)) < (float64(v3) / 0.1)) {
                    v38 := GenS4{F0: (-(fi22)), F1: -279, F2: 42}
                    _ = v38
                }
                var v39 uint16 = (^(v23))
                _ = v39
            }
            _ = fi24
            _ = v25
        }
        _ = fi22
        _ = v23
    }
    v40 := []int{662, 669}
    switch 9223372036854775807 {
    case 0:
        if len(v40) > 0 {
            v40[0] = 2147483647
        }
    case -491:
        var v41 uint8 = 1
        _ = v41
    }
    v42 := GenS5{F0: 752, F1: 0, F2: 128}
    v43 := GenS5{F0: 65535, F1: -199, F2: 42}
    v44 := GenS6{F0: 797}
    _ = len(v40)
    if (!false) {
        if len(v40) >= 2 {
            v40 = v40[1:len(v40)]
        }
    } else {
        fmt.Println(v2, v1, v1)
    }
    for fi45 := 0; fi45 < 19; fi45++ {
        v46 := GenS7{F0: 876, F1: -986}
        var v47 int32 = (v1 / ((v1) | 1))
        _ = fi45
        _ = v46
        _ = v47
    }
    var v48 int = 32768
    fmt.Println(v48, v1)
    switch ("foo bar" + "a") {
    case "hello":
        var v49 int8 = 127
        _ = v49
    default:
        var v50 uint8 = 0
        _ = v50
    }
    if false {
        if len(v40) >= 2 {
            v40 = v40[1:len(v40)]
        }
        fmt.Println(v2)
    } else {
        var v51 uint32 = 42
        if (false && ((0.1 / -2.25) >= (float32(((0.0 * v2) / v2)) * v4))) {
            var v52 uint16 = 32767
            v53 := GenS8{F0: ("\n\t" + "hello"), F1: 686}
            _ = v52
            _ = v53
        } else {
            for fi54 := 0; fi54 < 18; fi54++ {
                if (false || false) {
                    fmt.Println(v48)
                    var v55 uint16 = 524
                    _ = v55
                } else {
                    var v56 string = "\n\t"
                    switch v48 {
                    case 23:
                        var v57 uint8 = 128
                        _ = v57
                    case 42:
                        var v58 bool = (false && (1 > 32767))
                        _ = v58
                    }
                    _ = v56
                }
                _ = fi54
            }
        }
        _ = v51
    }
    var v59 uint32 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v40:", len(v40), v40)
    fmt.Println("v48:", v48)
    fmt.Println("v59:", v59)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v42
    _ = v43
    _ = v44
    _ = v48
    _ = v59
}
