package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }
type GenS2 struct { F0 uint32; F1 uint32; F2 string }
type GenS3 struct { F0 uint; F1 bool }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 string; F1 int8; F2 uint64 }
type GenS6 struct { F0 int16; F1 int }
func genFn0(p0 uint, p1 bool) uint32 {
    p0 = (p0 >> (((p0 ^ p0)) & 63))
    v6 := GenS1{F0: 0, F1: -1}
    v7 := GenS1{F0: 32767, F1: -940}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 388
}
func genFn1(p0 int16, p1 bool) int8 {
    var v8 bool = (true || (!false))
    p1 = true
    v9 := GenS1{F0: 575, F1: 599}
    var v10 int16 = 0
    _ = p0
    _ = p1
    _ = v10
    _ = v8
    _ = v9
    return 1
}

func main() {
    var v1 uint8 = 42
    var v2 int64 = -9223372036854775808
    var v3 uint32 = 2147483648
    var v4 uint = 2
    var v5 uint16 = 32767
    if true {
        v1 = v1
    } else {
        var v11 int32 = 42
        v12 := []int{93, -329, 4294967295}
        _ = v11
        _ = v12
    }
    fmt.Println(v4)
    if ("" == "a") {
        fmt.Println(v4, v1, v2)
        if true {
            var v13 uint8 = (^(v1))
            v14 := []int{-1}
            v3 += v3
            _ = v13
            _ = v14
        }
        for fi15 := 0; fi15 < 5; fi15++ {
            v16 := []int{-254, (fi15 / (((fi15 | (fi15 << (((v4 * (v4 / (((v4 % (((v4 << ((0) & 63))) | 1))) | 1)))) & 63)))) | 1))}
            _ = fi15
            _ = v16
        }
    }
    v17 := []int{616, -933, 32768, -297}
    v18 := GenS1{F0: (v5 ^ (v5 >> ((v4) & 63))), F1: 4294967295}
    if len(v17) > 0 {
        v17[0] = 42
    }
    for fi19 := 0; fi19 < 2; fi19++ {
        for fi20 := 0; fi20 < 5; fi20++ {
            switch ("a" + "\n\t") {
            case "":
                fmt.Println(v4)
            case "a":
                var v21 uint16 = (v5 << (((v4 >> ((49) & 63))) & 63))
                _ = v21
            case "foo bar":
                v22 := GenS2{F0: 0, F1: uint32(v1), F2: ""}
                var v23 uint64 = 343
                v1 = 255
                _ = v22
                _ = v23
            }
            _ = fi20
        }
        _ = fi19
    }
    if true {
        for fi24 := 0; fi24 < 14; fi24++ {
            fmt.Println(v5)
            for fi25 := 0; fi25 < 16; fi25++ {
                v26 := GenS1{F0: v5, F1: (fi24 % (((^(fi25))) | 1))}
                _ = fi25
                _ = v26
            }
            _ = fi24
        }
        if ((!true) && false) {
            v27 := GenS3{F0: 776, F1: false}
            v1 = (v1 >> ((v4) & 63))
            fmt.Println(v4, v5)
            _ = v27
        } else {
            if true {
                var v28 bool = false
                _ = v28
            }
        }
    } else {
        switch 128 {
        case 455:
            v4 += (v4 % ((18446744073709551615) | 1))
        case 1:
            switch -311 {
            case 9223372036854775807:
                var v29 int16 = int16(v4)
                _ = v29
            case 4294967295:
                var v30 int8 = 127
                _ = v30
            case 127:
                if false {
                    switch 804 {
                    case 9223372036854775807:
                        if (!(1e10 < -2.25)) {
                            switch 9223372036854775807 {
                            case -9223372036854775808:
                                var v31 bool = true
                                _ = v31
                            case 42:
                                for fi32 := 0; fi32 < 17; fi32++ {
                                    var v33 uint = 0
                                    var v34 int8 = 100
                                    _ = fi32
                                    _ = v33
                                    _ = v34
                                }
                            }
                            var v35 uint8 = (v1 + (v1 >> ((v4) & 63)))
                            v36 := GenS1{F0: (v5 % ((v5) | 1)), F1: 127}
                            _ = v35
                            _ = v36
                        } else {
                            for fi37 := 0; fi37 < 8; fi37++ {
                                v38 := GenS4{F0: 1000}
                                var v39 bool = (-983 != -1)
                                _ = fi37
                                _ = v38
                                _ = v39
                            }
                        }
                    }
                    if false {
                        v40 := GenS5{F0: "a", F1: 100, F2: 0}
                        v3 += v3
                        for fi41 := 0; fi41 < 2; fi41++ {
                            if (357 >= 2) {
                                if false {
                                    for fi42 := 0; fi42 < 2; fi42++ {
                                        v4 = (^(v4))
                                        var v43 string = "foo bar"
                                        _ = fi42
                                        _ = v43
                                    }
                                    switch "hello" {
                                    case "\n\t":
                                        var v44 int16 = -979
                                        _ = v44
                                    }
                                } else {
                                    if (false || true) {
                                        v1 -= v1
                                    }
                                }
                                fmt.Println(v3, v2, v5)
                                switch fi41 {
                                case 0:
                                    fmt.Println(v4, v1, v5)
                                case 681:
                                    v45 := GenS1{F0: 461, F1: fi41}
                                    _ = v45
                                default:
                                    for fi46 := 0; fi46 < 9; fi46++ {
                                        var v47 uint64 = uint64(v4)
                                        v47 = 0
                                        fmt.Println(v3)
                                        _ = fi46
                                        _ = v47
                                    }
                                }
                            }
                            _ = fi41
                        }
                        _ = v40
                    } else {
                        switch ("hello" + "hello") {
                        case "":
                            var v48 float32 = float32(3.14159)
                            _ = v48
                        }
                        for fi49 := 0; fi49 < 19; fi49++ {
                            for fi50 := 0; fi50 < 14; fi50++ {
                                var v51 uint = (v4 + (v4 & v4))
                                v1 -= v1
                                _ = fi50
                                _ = v51
                            }
                            for fi52 := 0; fi52 < 12; fi52++ {
                                v17 = append(v17, fi52)
                                _ = fi52
                            }
                            v53 := GenS6{F0: 1, F1: (fi49 + 0)}
                            _ = fi49
                            _ = v53
                        }
                    }
                }
            default:
                v5 = v5
            }
        case 128:
            var v54 int = int(v5)
            _ = v54
        }
        switch "hello" {
        case "\n\t":
            var v55 int = 4294967295
            _ = v55
        }
        v3 = (^(v3))
    }
    switch "foo bar" {
    case "\n\t":
        var v56 uint8 = v1
        _ = v56
    case "hello":
        v17 = append(v17, 0)
    case "foo bar":
        v4 -= 605
    }
    var v57 string = ""
    v4 -= v4
    v58 := GenS1{F0: (v5 ^ (v5 | v5)), F1: 671}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v57:", v57)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v57
    _ = v58
}
