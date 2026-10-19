package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int8; F1 string; F2 uint32 }
type GenS4 struct { F0 int32; F1 int8; F2 int }
type GenS5 struct { F0 int32 }
func genFn0(p0 uint32, p1 uint16) uint {
    var v7 float64 = 1e10
    var v8 string = ("a" + "a")
    var v9 float64 = float64(1.5)
    v10 := GenS1{F0: 42, F1: v8}
    _ = p0
    _ = p1
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 2
}
func genFn1(p0 int64, p1 uint16, p2 int16) int32 {
    var v11 float32 = (0.1 * 3.14159)
    v12 := GenS2{F0: false}
    var v13 string = "a"
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    _ = v13
    return -774
}

func main() {
    var v1 int16 = -365
    var v2 int = 9223372036854775807
    var v3 bool = true
    v4 := []int{(v2 * -860), (v2 - 1), v2, v2}
    var v5 float32 = (0.0 + 0.1)
    v6 := make(map[int]int)
    switch "foo bar" {
    case "\n\t":
        v14 := GenS1{F0: v1, F1: ("hello" + ("\n\t" + "hello"))}
        _ = v14
    default:
        v5 -= (((v5 * -2.25) / -2.25) / ((v5 / 3.14159) * v5))
    }
    switch (("foo bar" + "hello") + "hello") {
    case "a":
        v15 := GenS1{F0: v1, F1: "foo bar"}
        _ = v15
    case "foo bar":
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    v1 = 1000
    var v16 bool = v3
    switch v2 {
    case 286:
        var v17 int16 = v1
        _ = v17
    case 32768:
        fmt.Println(v3)
    }
    fmt.Println(v3)
    for fi18 := 0; fi18 < 12; fi18++ {
        if (544 != 926) {
            v19 := GenS1{F0: -280, F1: ""}
            v4 = append(v4, 2147483647)
            _ = v19
        }
        fmt.Println(v5, v2, v16)
        _ = fi18
    }
    v16 = (float32((((((-2.25 / 1.0) - -2.25) / 0.1) / 1.0) - 0.1)) > v5)
    var v20 uint8 = 155
    if v3 {
        fmt.Println(v1, v16)
        for fi21 := 0; fi21 < 10; fi21++ {
            for fi22 := 0; fi22 < 1; fi22++ {
                var v23 int32 = 2147483647
                v23 += (v23 >> ((1) & 63))
                _ = fi22
                _ = v23
            }
            if v3 {
                for fi24 := 0; fi24 < 10; fi24++ {
                    v5 = (v5 - (float32(((3.14159 / 1.0) + 0.0)) * v5))
                    fmt.Println(v2, v2, v1)
                    var v25 uint32 = 861
                    _ = fi24
                    _ = v25
                }
            }
            var v26 string = "hello"
            _ = fi21
            _ = v26
        }
    } else {
        for fi27 := 0; fi27 < 4; fi27++ {
            for fi28 := 0; fi28 < 7; fi28++ {
                v5 += v5
                _ = fi28
            }
            for fi29 := 0; fi29 < 12; fi29++ {
                v30 := GenS3{F0: 60, F1: "", F2: 0}
                if (!(65535 > 328)) {
                    switch "foo bar" {
                    case "hello":
                        switch 128 {
                        case -949:
                            v31 := GenS4{F0: 1, F1: 4, F2: v2}
                            _ = v31
                        default:
                            switch ("" + ("foo bar" + "a")) {
                            case "\n\t":
                                _ = len(v6)
                            case "hello":
                                v16 = v16
                            default:
                                var v32 uint = 42
                                _ = v32
                            }
                        }
                    case "\n\t":
                        if (v16 && v16) {
                            var v33 int = fi27
                            var v34 int16 = 1000
                            var v35 float64 = 0.1
                            _ = v33
                            _ = v34
                            _ = v35
                        } else {
                            var v36 uint8 = 197
                            _ = v36
                        }
                    case "":
                        switch -337 {
                        case 9223372036854775807:
                            var v37 int = -720
                            _ = v37
                        case 128:
                            var v38 bool = (v16 && true)
                            _ = v38
                        default:
                            switch "" {
                            case "foo bar":
                                v39 := GenS2{F0: v16}
                                _ = v39
                            case "a":
                                v3 = v16
                            case "hello":
                                var v40 int8 = -128
                                _ = v40
                            default:
                                _ = len(v4)
                            }
                        }
                    }
                } else {
                    if (false || (200 == 728)) {
                        switch fi27 {
                        case 386:
                            switch "hello" {
                            case "foo bar":
                                v16 = (!v3)
                            case "a":
                                var v41 uint32 = 775
                                _ = v41
                            case "hello":
                                v5 = v5
                            }
                        }
                        var v42 int64 = 9223372036854775807
                        fmt.Println(v5, v1)
                        _ = v42
                    } else {
                        v16 = v3
                        var v43 bool = ((!v3) && true)
                        if v3 {
                            var v44 uint = 442
                            v20 -= 127
                            v45 := GenS5{F0: 2147483647}
                            _ = v44
                            _ = v45
                        }
                        _ = v43
                    }
                }
                _ = fi29
                _ = v30
            }
            for fi46 := 0; fi46 < 2; fi46++ {
                if v16 {
                    v3 = (false || (-1 <= 42))
                    v2 -= fi46
                } else {
                    if v16 {
                        v47 := GenS1{F0: (^(v1)), F1: ("foo bar" + "")}
                        _ = len(v6)
                        if (18446744073709551615 >= 18446744073709551615) {
                            fmt.Println(fi27)
                        }
                        _ = v47
                    }
                    v2 -= (^(fi27))
                    var v48 uint8 = 42
                    v16 = (!v16)
                    v6[fi27] = fi46
                    _ = v48
                }
                if false {
                    v49 := GenS4{F0: -1, F1: 0, F2: -421}
                    var v50 int32 = genFn1(-170, 1, (-(v1)))
                    _ = v49
                    _ = v50
                } else {
                    v2 -= (fi46 << ((4294967295) & 63))
                }
                v20 = v20
                _ = fi46
            }
            _ = fi27
        }
    }
    for fi51 := 0; fi51 < 5; fi51++ {
        v3 = false
        for fi52 := 0; fi52 < 2; fi52++ {
            delete(v6, -2147483648)
            v4 = append(v4, (^(v2)))
            _ = fi52
        }
        _ = fi51
    }
    var v53 float64 = 0.0
    if (4294967295 > (v2 - 9223372036854775807)) {
        var v54 int64 = 120
        switch 9223372036854775807 {
        case 42:
            v3 = false
        case -9223372036854775808:
            delete(v6, 32767)
        case -203:
            v55 := GenS4{F0: -96, F1: 42, F2: 9223372036854775807}
            _ = v55
        default:
            for fi56 := 0; fi56 < 8; fi56++ {
                var v57 uint16 = 32767
                if (42 == 1) {
                    v20 += v20
                    _ = len(v4)
                    fmt.Println(v54)
                } else {
                    fmt.Println(v3, v54)
                    switch (fi56 << ((577) & 63)) {
                    case 2147483647:
                        var v58 int32 = -634
                        _ = v58
                    case -293:
                        if false {
                            v59 := v6[-615]
                            fmt.Println(v16, v1)
                            v20 -= 128
                            _ = v59
                        }
                    default:
                        v20 = (^(v20))
                    }
                }
                var v60 int16 = -32768
                _ = fi56
                _ = v57
                _ = v60
            }
        }
        if (!(42 >= genFn1(v54, 42, v1))) {
            fmt.Println(v16)
        }
        _ = v54
    } else {
        v61 := v6[(v2 / (((v2 >> ((genFn0(42, 376)) & 63))) | 1))]
        _ = v61
    }
    for fi62 := 0; fi62 < 12; fi62++ {
        if v3 {
            var v63 uint16 = 2
            _ = v63
        }
        _ = fi62
    }
    _ = len(v6)
    var v64 uint16 = 65535
    v65 := GenS4{F0: 719, F1: 127, F2: -9223372036854775808}
    var v66 int8 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v53:", v53)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v64:", v64)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v16
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v53
    _ = v6
    _ = v64
    _ = v65
    _ = v66
}
