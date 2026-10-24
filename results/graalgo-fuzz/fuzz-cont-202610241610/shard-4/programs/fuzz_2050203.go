package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int; F1 uint64; F2 uint }
type GenS4 struct { F0 string; F1 uint; F2 uint8 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 uint16; F1 int16 }
func genFn0(p0 uint32) uint64 {
    p0 = (p0 << ((18446744073709551615) & 63))
    fmt.Println(p0)
    _ = p0
    return 0
}

func main() {
    v1 := []int{-9223372036854775808, -1, 935, -2147483648}
    v2 := []uint32{96, 603}
    var v3 int64 = 9223372036854775807
    v4 := make(map[int]string)
    var v5 uint16 = 1
    for fi6 := 0; fi6 < 17; fi6++ {
        v7 := GenS1{F0: 1, F1: fi6}
        v5 -= v5
        _ = len(v4)
        _ = fi6
        _ = v7
    }
    v3 = (v3 * 42)
    if ((true && ("hello" > "a")) && true) {
        for fi8 := 0; fi8 < 15; fi8++ {
            if false {
                var v9 int8 = -1
                v5 = v5
                _ = v9
            } else {
                var v10 uint32 = 42
                if ((!(true && true)) && (true || false)) {
                    switch fi8 {
                    case 127:
                        if false {
                            var v11 uint8 = 1
                            var v12 float32 = 0.1
                            _ = v11
                            _ = v12
                        }
                    case 128:
                        var v13 uint16 = 255
                        _ = v13
                    }
                }
                _ = v10
            }
            var v14 int16 = -381
            for fi15 := 0; fi15 < 11; fi15++ {
                v14 = v14
                _ = len(v4)
                _ = fi15
            }
            _ = fi8
            _ = v14
        }
        v3 = (v3 % ((v3) | 1))
    } else {
        v5 -= (v5 >> ((900) & 63))
        if (!(65535 < 0)) {
            switch ("\n\t" + "a") {
            case "foo bar":
                if true {
                    for fi16 := 0; fi16 < 14; fi16++ {
                        v3 = (v3 << ((0) & 63))
                        _ = fi16
                    }
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    var v17 uint16 = (v5 - 65535)
                    _ = v17
                } else {
                    switch ("hello" + "hello") {
                    case "a":
                        var v18 float64 = 1e10
                        _ = v18
                    case "hello":
                        var v19 uint = 255
                        _ = v19
                    case "":
                        v5 -= v5
                    }
                    switch -439 {
                    case 42:
                        var v20 string = "foo bar"
                        _ = v20
                    case 730:
                        v5 = v5
                    case 2147483647:
                        var v21 uint32 = 42
                        _ = v21
                    default:
                        v5 = (^(v5))
                    }
                    for fi22 := 0; fi22 < 16; fi22++ {
                        v1 = append(v1, fi22)
                        delete(v4, fi22)
                        _ = fi22
                    }
                }
            }
            var v23 uint8 = 65
            v3 = (v3 >> ((42) & 63))
            _ = v23
        } else {
            if (!false) {
                v3 = 0
            } else {
                v24 := GenS2{F0: 2}
                var v25 int32 = -2147483648
                switch ("hello" + "foo bar") {
                case "foo bar":
                    var v26 int32 = (v25 & v25)
                    _ = v26
                }
                _ = v24
                _ = v25
            }
        }
        var v27 int = 32767
        _ = v27
    }
    switch "\n\t" {
    case "a":
        v28 := GenS3{F0: 128, F1: 102, F2: 4294967295}
        _ = v28
    case "hello":
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    default:
        for fi29 := 0; fi29 < 15; fi29++ {
            _ = len(v4)
            var v30 float64 = -2.25
            _ = fi29
            _ = v30
        }
    }
    if false {
        switch 4294967295 {
        case 433:
            var v31 int16 = 42
            _ = v31
        default:
            fmt.Println(v5, v3)
        }
    }
    var v32 string = "\n\t"
    if (false || false) {
        var v33 bool = (v32 >= "hello")
        var v34 int64 = (v3 & v3)
        _ = v33
        _ = v34
    } else {
        var v35 uint16 = v5
        switch (v32 + (v32 + v32)) {
        case "\n\t":
            v5 = 65535
        case "":
            var v36 bool = true
            _ = v36
        case "hello":
            if (314 <= 18446744073709551615) {
                if false {
                    var v37 int16 = 1
                    v38 := GenS4{F0: (v32 + v32), F1: 0, F2: 127}
                    v3 += v3
                    _ = v37
                    _ = v38
                } else {
                    _ = len(v4)
                    v32 = v32
                    if ((true && true) && (42 < -32768)) {
                        var v39 float32 = 0.0
                        _ = v39
                    }
                }
            } else {
                if true {
                    v40 := GenS1{F0: 2, F1: 201}
                    _ = v40
                } else {
                    v41 := GenS3{F0: -2147483648, F1: 18446744073709551615, F2: 18446744073709551615}
                    for fi42 := 0; fi42 < 7; fi42++ {
                        for fi43 := 0; fi43 < 18; fi43++ {
                            v32 = "foo bar"
                            switch fi43 {
                            case 127:
                                fmt.Println(v35)
                                v1 = append(v1, fi43)
                            case 232:
                                v44 := GenS1{F0: 926, F1: (fi42 % ((4294967295) | 1))}
                                _ = v44
                            case 2147483647:
                                fmt.Println(v32)
                            default:
                                v35 -= v35
                            }
                            v4[4294967295] = "a"
                            var v45 uint16 = v5
                            _ = fi43
                            _ = v45
                        }
                        for fi46 := 0; fi46 < 19; fi46++ {
                            var v47 int = (fi42 << ((18446744073709551615) & 63))
                            v48 := GenS5{F0: 58}
                            _ = fi46
                            _ = v47
                            _ = v48
                        }
                        if true {
                            if ("foo bar" < "hello") {
                                v5 = v5
                                var v49 uint8 = 42
                                _ = v49
                            }
                            var v50 uint8 = 2
                            if len(v1) > 0 {
                                v1[0] = (-(fi42))
                            }
                            _ = v50
                        }
                        _ = fi42
                    }
                    _ = v41
                }
            }
        default:
            switch 746 {
            case 32768:
                if (42 < 2) {
                    if true {
                        var v51 int8 = -1
                        for fi52 := 0; fi52 < 3; fi52++ {
                            v51 += v51
                            switch (fi52 / (((-(fi52))) | 1)) {
                            case 32768:
                                v53 := GenS6{F0: 4294967295}
                                _ = v53
                            case 42:
                                v54 := v4[fi52]
                                _ = v54
                            case 9223372036854775807:
                                var v55 float64 = 3.14159
                                _ = v55
                            }
                            _ = fi52
                        }
                        _ = v51
                    } else {
                        v3 = v3
                    }
                }
            case -2147483648:
                if false {
                    _ = len(v2)
                    for fi56 := 0; fi56 < 14; fi56++ {
                        if false {
                            v35 -= v35
                            v35 += v35
                        }
                        for fi57 := 0; fi57 < 5; fi57++ {
                            v32 = v32
                            fmt.Println(v5)
                            delete(v4, fi57)
                            _ = fi57
                        }
                        v58 := GenS6{F0: 255}
                        _ = fi56
                        _ = v58
                    }
                } else {
                    var v59 uint8 = 2
                    _ = v59
                }
            default:
                _ = len(v1)
            }
        }
        v5 += (v5 - (^(v5)))
        _ = v35
    }
    switch v32 {
    case "":
        v5 = 42
    case "hello":
        if true {
            if ((v5 <= v5) && true) {
                v3 = (v3 * v3)
            } else {
                v60 := GenS7{F0: v5, F1: -1}
                for fi61 := 0; fi61 < 18; fi61++ {
                    v5 += v5
                    if false {
                        v2 = append(v2, 157)
                    } else {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                    var v62 uint16 = v5
                    _ = fi61
                    _ = v62
                }
                _ = len(v4)
                _ = v60
            }
        }
    }
    v32 = v32
    var v63 uint64 = 0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v63
}
