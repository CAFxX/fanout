package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint8; F1 int64; F2 int64 }
type GenS4 struct { F0 bool; F1 int8; F2 int32 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 uint8; F1 uint64; F2 string }
type GenS8 struct { F0 uint; F1 uint64; F2 int }
type GenS9 struct { F0 int16; F1 string }
func genFn0() uint {
    if (!true) {
        var v7 int16 = 32767
        for fi8 := 0; fi8 < 5; fi8++ {
            v7 = (^(v7))
            v9 := GenS1{F0: 0}
            _ = fi8
            _ = v9
        }
        fmt.Println(v7)
        _ = v7
    }
    var v10 int8 = 42
    if false {
        if true {
            for fi11 := 0; fi11 < 10; fi11++ {
                if true {
                    if false {
                        v10 = -128
                        v12 := GenS2{F0: 0}
                        _ = v12
                    }
                    v10 += (v10 & (v10 / ((v10) | 1)))
                    var v13 uint64 = 42
                    _ = v13
                }
                _ = fi11
            }
        } else {
            for fi14 := 0; fi14 < 7; fi14++ {
                v15 := GenS3{F0: 1, F1: 42, F2: 1}
                _ = fi14
                _ = v15
            }
        }
    } else {
        v16 := GenS1{F0: 0}
        var v17 uint = 65535
        switch -1 {
        case 741:
            switch ("\n\t" + ("\n\t" + "a")) {
            case "a":
                switch 127 {
                case 4294967295:
                    switch -582 {
                    case 32767:
                        v10 = v10
                    default:
                        if (42 <= (^(v10))) {
                            var v18 uint32 = 864
                            _ = v18
                        } else {
                            v19 := []int{32768, -372}
                            v17 = v17
                            var v20 int32 = -345
                            _ = v19
                            _ = v20
                        }
                    }
                case -973:
                    v21 := GenS2{F0: (v17 ^ v17)}
                    _ = v21
                }
            case "":
                var v22 uint8 = 127
                _ = v22
            default:
                switch 32767 {
                case 42:
                    fmt.Println(v17, v17)
                default:
                    if true {
                        fmt.Println(v10, v17)
                        v23 := GenS1{F0: 9223372036854775807}
                        _ = v23
                    }
                }
            }
        case 1:
            for fi24 := 0; fi24 < 19; fi24++ {
                var v25 uint = v17
                v26 := GenS4{F0: false, F1: (v10 & (^(v10))), F2: -430}
                _ = fi24
                _ = v25
                _ = v26
            }
        case 9223372036854775807:
            v27 := []int{537, 128}
            _ = v27
        }
        _ = v16
        _ = v17
    }
    v28 := GenS5{F0: 1}
    for fi29 := 0; fi29 < 11; fi29++ {
        fmt.Println(v10, v10)
        var v30 int16 = -91
        _ = fi29
        _ = v30
    }
    _ = v10
    _ = v28
    return 65535
}
func genFn1(p0 int8) uint8 {
    if (!false) {
        fmt.Println(p0)
        if (9223372036854775807 >= 42) {
            p0 -= p0
        }
        p0 = (p0 << ((0) & 63))
    } else {
        switch ("hello" + "foo bar") {
        case "":
            for fi31 := 0; fi31 < 2; fi31++ {
                for fi32 := 0; fi32 < 13; fi32++ {
                    var v33 uint32 = 616
                    if true {
                        var v34 uint = 18446744073709551615
                        var v35 int32 = 42
                        var v36 bool = false
                        _ = v34
                        _ = v35
                        _ = v36
                    } else {
                        var v37 int = (fi32 << ((65535) & 63))
                        if (true && ((v33 ^ 2147483647) < (v33 | v33))) {
                            v33 -= 42
                            v38 := GenS6{F0: (fi32 / ((fi31) | 1))}
                            v37 -= 1
                            _ = v38
                        }
                        fmt.Println(p0, v33)
                        _ = v37
                    }
                    if (false && true) {
                        p0 -= (p0 * 42)
                        switch ("" + "") {
                        case "":
                            p0 = -128
                        case "a":
                            var v39 uint = 468
                            _ = v39
                        case "\n\t":
                            var v40 uint = 18446744073709551615
                            _ = v40
                        }
                        v41 := GenS2{F0: 18446744073709551615}
                        var v42 bool = false
                        var v43 int = (fi31 ^ fi31)
                        var v44 float32 = 1e10
                        fmt.Println(v42, v43, p0)
                        _ = v41
                        _ = v42
                        _ = v43
                        _ = v44
                    } else {
                        v45 := []int{fi32, (fi31 & (fi32 << ((451) & 63)))}
                        fmt.Println(v33)
                        v46 := GenS1{F0: 0}
                        _ = v45
                        _ = v46
                    }
                    v47 := []int{-1, (fi32 * 42), -912, 128}
                    var v48 uint = 18446744073709551615
                    _ = fi32
                    _ = v33
                    _ = v47
                    _ = v48
                }
                _ = fi31
            }
        }
    }
    var v49 uint64 = 0
    for fi50 := 0; fi50 < 5; fi50++ {
        var v51 uint32 = 1
        if (true && true) {
            v52 := GenS7{F0: 48, F1: 42, F2: "\n\t"}
            _ = v52
        }
        _ = fi50
        _ = v51
    }
    v53 := GenS7{F0: 84, F1: (v49 >> ((genFn0()) & 63)), F2: (("foo bar" + "") + "hello")}
    _ = p0
    _ = v49
    _ = v53
    return 16
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[uint32]uint32)
    var v3 uint8 = 127
    var v4 int = 0
    v5 := make(map[int]uint32)
    v6 := make(map[string]int)
    v4 = (-(v4))
    switch v4 {
    case 293:
        if ((2 > 2147483648) && true) {
            fmt.Println(v4, v4)
        }
    case 0:
        v54 := []int{v4, (v4 % ((4294967295) | 1)), (v4 % ((v4) | 1))}
        _ = v54
    }
    v55 := GenS8{F0: 4294967295, F1: 1, F2: -1}
    if true {
        var v56 int = -9223372036854775808
        _ = v56
    }
    var v57 string = (("\n\t" + (("foo bar" + "foo bar") + "\n\t")) + "a")
    if ((float64(1e10) - 0.1) >= 3.14159) {
        switch v57 {
        case "":
            if (144 == 65535) {
                var v58 uint16 = 32767
                v59 := []int{(^(v4)), 0, v4}
                var v60 float64 = 1e10
                _ = v58
                _ = v59
                _ = v60
            }
        case "\n\t":
            for fi61 := 0; fi61 < 17; fi61++ {
                var v62 float64 = (((1e10 + 1e10) * 3.14159) + 3.14159)
                if false {
                    v4 = (v4 / ((1) | 1))
                } else {
                    var v63 int16 = 1
                    for fi64 := 0; fi64 < 15; fi64++ {
                        v65 := GenS2{F0: 0}
                        _ = fi64
                        _ = v65
                    }
                    _ = v63
                }
                v66 := v5[212]
                _ = fi61
                _ = v62
                _ = v66
            }
        }
        v67 := []int{}
        _ = v67
    }
    for fi68 := 0; fi68 < 16; fi68++ {
        for fi69 := 0; fi69 < 9; fi69++ {
            var v70 uint8 = (v3 ^ (^(v3)))
            switch ("foo bar" + v57) {
            case "\n\t":
                v71 := GenS9{F0: 1, F1: v57}
                _ = v71
            default:
                v72 := []int{(v4 | -1)}
                _ = v72
            }
            _ = fi69
            _ = v70
        }
        v4 -= 969
        v73 := []int{9223372036854775807}
        _ = fi68
        _ = v73
    }
    var v74 int32 = 1
    fmt.Println(v3)
    var v75 float64 = 1e10
    fmt.Println(v74)
    fmt.Println(v75, v74)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v57:", v57)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v74:", v74)
    fmt.Println("v75:", v75)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v55
    _ = v57
    _ = v6
    _ = v74
    _ = v75
}
