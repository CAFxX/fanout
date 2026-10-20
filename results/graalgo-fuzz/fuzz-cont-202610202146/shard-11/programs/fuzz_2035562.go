package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int32; F1 int; F2 uint }
type GenS5 struct { F0 int8; F1 uint8; F2 int8 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 string }
type GenS8 struct { F0 int }
type GenS9 struct { F0 uint64; F1 int16 }
type GenS10 struct { F0 int64; F1 int64 }
type GenS11 struct { F0 uint8; F1 int32 }
func genFn0() bool {
    for fi4 := 0; fi4 < 18; fi4++ {
        if (false || ((true && true) || false)) {
            if false {
                v5 := []int{(fi4 ^ 561), -586, 1}
                var v6 float64 = 1.5
                _ = v5
                _ = v6
            } else {
                fmt.Println(fi4)
                v7 := []int{}
                var v8 float64 = 0.0
                _ = v7
                _ = v8
            }
            v9 := GenS1{F0: 66, F1: 18446744073709551615}
            var v10 uint64 = 2
            _ = v10
            _ = v9
        } else {
            var v11 int8 = 1
            switch "\n\t" {
            case "foo bar":
                var v12 uint64 = 188
                _ = v12
            default:
                v13 := GenS2{F0: 4294967295}
                _ = v13
            }
            _ = v11
        }
        var v14 uint8 = 1
        _ = fi4
        _ = v14
    }
    var v15 bool = ((12 > -364) || false)
    fmt.Println(v15)
    v15 = false
    _ = v15
    return v15
}

func main() {
    var v1 int = -1
    var v2 uint16 = 32767
    v3 := make(map[string]string)
    var v16 uint8 = 127
    v17 := GenS3{F0: 42}
    v16 -= (v16 - (v16 + v16))
    var v18 uint16 = 2
    fmt.Println(v18, v16, v18)
    if (true && (914 <= 0)) {
        if false {
            var v19 string = ""
            for fi20 := 0; fi20 < 3; fi20++ {
                for fi21 := 0; fi21 < 19; fi21++ {
                    v22 := GenS4{F0: 625, F1: v1, F2: 18446744073709551615}
                    _ = fi21
                    _ = v22
                }
                var v23 string = (v19 + (v19 + v19))
                switch "hello" {
                case "hello":
                    v23 = "a"
                }
                _ = fi20
                _ = v23
            }
            v24 := GenS5{F0: -1, F1: v16, F2: -1}
            _ = v19
            _ = v24
        }
        var v25 uint8 = (v16 - (v16 + 255))
        for fi26 := 0; fi26 < 12; fi26++ {
            for fi27 := 0; fi27 < 5; fi27++ {
                var v28 uint32 = 103
                v18 = v18
                v28 -= (v28 & (v28 % ((65535) | 1)))
                _ = fi27
                _ = v28
            }
            for fi29 := 0; fi29 < 16; fi29++ {
                var v30 int8 = 127
                v31 := GenS4{F0: -44, F1: -9223372036854775808, F2: 18446744073709551615}
                v3["a"] = "\n\t"
                v25 += 1
                _ = fi29
                _ = v30
                _ = v31
            }
            _ = fi26
        }
        _ = v25
    } else {
        if ((-32768 != 0) && true) {
            for fi32 := 0; fi32 < 16; fi32++ {
                for fi33 := 0; fi33 < 8; fi33++ {
                    v1 = int(v16)
                    _ = fi33
                }
                if (42 <= 4294967295) {
                    v34 := GenS1{F0: 255, F1: 0}
                    _ = v34
                } else {
                    v35 := GenS1{F0: 42, F1: 1}
                    _ = v35
                }
                _ = fi32
            }
            var v36 uint32 = 176
            _ = v36
        } else {
            switch -414 {
            case 32767:
                if false {
                    var v37 uint64 = 18446744073709551615
                    _ = v37
                } else {
                    v18 = 968
                }
            case 538:
                v3["foo bar"] = ""
            case 127:
                fmt.Println(v16, v2, v16)
            default:
                v18 -= v2
            }
            v38 := v3["hello"]
            _ = v38
        }
        v39 := GenS3{F0: 1}
        switch "hello" {
        case "hello":
            var v40 string = "hello"
            _ = v40
        case "a":
            var v41 int16 = 597
            _ = v41
        case "\n\t":
            var v42 uint32 = 65535
            _ = v42
        default:
            var v43 uint16 = uint16(v16)
            _ = v43
        }
        _ = v39
    }
    fmt.Println(v16, v18, v18)
    var v44 uint16 = (v2 | v2)
    var v45 uint64 = 2
    v46 := v3["a"]
    var v47 float64 = 1e10
    for fi48 := 0; fi48 < 9; fi48++ {
        switch (v46 + v46) {
        case "a":
            for fi49 := 0; fi49 < 9; fi49++ {
                v50 := GenS6{F0: v46}
                fmt.Println(v45, fi49)
                switch (^(fi49)) {
                case 127:
                    if (255 != 0) {
                        v45 += 246
                    } else {
                        v51 := GenS5{F0: -44, F1: v16, F2: 100}
                        v52 := v3[v46]
                        switch (fi49 << ((18) & 63)) {
                        case 1:
                            v44 = v44
                        case -551:
                            v1 += fi48
                        case 633:
                            var v53 int8 = 100
                            _ = v53
                        }
                        _ = v51
                        _ = v52
                    }
                case -1:
                    if false {
                        v44 = (v44 * (v2 * (v44 - v44)))
                        v47 -= (v47 / 3.14159)
                    }
                default:
                    var v54 int64 = 0
                    _ = v54
                }
                _ = fi49
                _ = v50
            }
        case "hello":
            v3["hello"] = "a"
        default:
            v55 := []int{32767, fi48, v1, -244}
            _ = v55
        }
        _ = fi48
    }
    v56 := GenS7{F0: (v46 + v46)}
    fmt.Println(v47, v45)
    v57 := GenS8{F0: (^(v1))}
    if false {
        v16 -= 132
        v58 := []int{v1, v1, (-(v1)), (v1 | v1)}
        for fi59 := 0; fi59 < 8; fi59++ {
            fmt.Println(v44, v46)
            if (true && false) {
                v3[""] = "a"
            } else {
                for fi60 := 0; fi60 < 7; fi60++ {
                    var v61 uint = 18446744073709551615
                    if false {
                        v62 := GenS9{F0: (v45 * (v45 | (^(v45)))), F1: 32767}
                        if false {
                            v44 = v44
                            v46 = ((v46 + v46) + (v46 + v46))
                        }
                        switch "" {
                        case "foo bar":
                            fmt.Println(v16, v61)
                        case "hello":
                            _ = len(v58)
                        case "\n\t":
                            var v63 float32 = ((3.14159 * (0.1 * 0.0)) + (((1.5 - 1e10) * 1.5) + 3.14159))
                            _ = v63
                        }
                        _ = len(v58)
                        _ = v62
                    } else {
                        _ = len(v3)
                        v64 := GenS10{F0: 9223372036854775807, F1: 9223372036854775807}
                        var v65 uint32 = 2
                        if true {
                            v47 -= 0.0
                            v16 = (v16 % (((v16 << ((v61) & 63))) | 1))
                            fmt.Println(v2, v1, v65)
                        }
                        _ = v64
                        _ = v65
                    }
                    _ = fi60
                    _ = v61
                }
            }
            v2 -= (v44 - (v2 / ((v18) | 1)))
            _ = fi59
        }
        _ = v58
    } else {
        v66 := GenS2{F0: 4294967295}
        v67 := GenS11{F0: (v16 % ((v16) | 1)), F1: 79}
        v1 += v1
        _ = v66
        _ = v67
    }
    var v68 string = v46
    v69 := GenS1{F0: 255, F1: 18446744073709551615}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v68:", v68)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v44
    _ = v45
    _ = v46
    _ = v47
    _ = v56
    _ = v57
    _ = v68
    _ = v69
}
