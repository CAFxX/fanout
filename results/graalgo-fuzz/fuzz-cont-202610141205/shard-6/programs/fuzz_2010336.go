package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32 }
type GenS2 struct { F0 string; F1 uint32; F2 int64 }
type GenS3 struct { F0 uint8; F1 int32; F2 int32 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint16; F1 bool; F2 int8 }

func main() {
    var v1 uint8 = 0
    v2 := make(map[string]uint32)
    v3 := make(map[int]int)
    var v4 float32 = 0.0
    v5 := GenS1{F0: 32767, F1: -1}
    switch "hello" {
    case "hello":
        fmt.Println(v1, v1)
    case "foo bar":
        if (493 != 42) {
            var v6 float32 = 1.5
            if ((v6 * v4) != ((v6 * -2.25) - v4)) {
                v7 := []int{-39, -1, 9223372036854775807, 127}
                _ = v7
            } else {
                for fi8 := 0; fi8 < 15; fi8++ {
                    v1 = 81
                    fmt.Println(fi8)
                    _ = fi8
                }
                if ((0 == -1) && true) {
                    for fi9 := 0; fi9 < 1; fi9++ {
                        if (float32(1.5) >= v6) {
                            if false {
                                for fi10 := 0; fi10 < 10; fi10++ {
                                    v6 = v4
                                    var v11 int8 = 95
                                    _ = fi10
                                    _ = v11
                                }
                                var v12 uint = 42
                                var v13 int8 = 127
                                _ = v12
                                _ = v13
                            }
                        }
                        _ = fi9
                    }
                    v14 := []int{}
                    _ = v14
                }
            }
            fmt.Println(v6, v4, v6)
            _ = v6
        } else {
            switch "a" {
            case "":
                var v15 uint = 0
                _ = v15
            case "a":
                for fi16 := 0; fi16 < 2; fi16++ {
                    for fi17 := 0; fi17 < 15; fi17++ {
                        var v18 int8 = -128
                        _ = fi17
                        _ = v18
                    }
                    v3[4294967295] = fi16
                    v1 -= 68
                    _ = fi16
                }
            case "foo bar":
                for fi19 := 0; fi19 < 12; fi19++ {
                    v20 := []int{(fi19 + 698), fi19}
                    var v21 uint = 1
                    _ = fi19
                    _ = v20
                    _ = v21
                }
            }
            var v22 int32 = 728
            _ = v22
        }
    case "a":
        switch 42 {
        case 540:
            v1 += v1
        case 127:
            v23 := GenS1{F0: 506, F1: 2147483647}
            _ = v23
        case -1:
            for fi24 := 0; fi24 < 13; fi24++ {
                fmt.Println(fi24, v1)
                _ = fi24
            }
        default:
            var v25 int8 = 127
            _ = v25
        }
    }
    for fi26 := 0; fi26 < 12; fi26++ {
        var v27 uint32 = 2147483648
        v4 = v4
        v1 = 42
        _ = fi26
        _ = v27
    }
    if ((false && ("foo bar" != "")) && (("foo bar" + "\n\t") > ("hello" + ((("\n\t" + "") + "\n\t") + "a")))) {
        for fi28 := 0; fi28 < 19; fi28++ {
            for fi29 := 0; fi29 < 7; fi29++ {
                v4 -= (0.0 - v4)
                _ = fi29
            }
            v4 = v4
            _ = fi28
        }
    }
    var v30 uint32 = 65535
    for fi31 := 0; fi31 < 9; fi31++ {
        var v32 int32 = 42
        switch "\n\t" {
        case "hello":
            v33 := v3[(fi31 * (^(fi31)))]
            _ = v33
        default:
            switch "a" {
            case "":
                var v34 string = "\n\t"
                _ = v34
            case "foo bar":
                for fi35 := 0; fi35 < 10; fi35++ {
                    if true {
                        v36 := GenS2{F0: "", F1: (v30 & v30), F2: 1}
                        var v37 int = int(v30)
                        _ = v36
                        _ = v37
                    } else {
                        v38 := GenS2{F0: "hello", F1: v30, F2: -9223372036854775808}
                        _ = v38
                    }
                    _ = fi35
                }
            default:
                if (v32 == 1) {
                    v1 = v1
                    if ("a" != ("" + "foo bar")) {
                        for fi39 := 0; fi39 < 9; fi39++ {
                            v40 := GenS2{F0: "foo bar", F1: v30, F2: -211}
                            v41 := GenS3{F0: v1, F1: int32(v1), F2: int32(fi31)}
                            _ = fi39
                            _ = v40
                            _ = v41
                        }
                        var v42 int64 = -1
                        for fi43 := 0; fi43 < 3; fi43++ {
                            var v44 uint32 = v30
                            _ = fi43
                            _ = v44
                        }
                        _ = v42
                    } else {
                        var v45 int8 = 42
                        var v46 uint = 65535
                        fmt.Println(v45, v1, v32)
                        _ = v45
                        _ = v46
                    }
                } else {
                    switch "a" {
                    case "a":
                        for fi47 := 0; fi47 < 4; fi47++ {
                            v30 += v30
                            v2[""] = v30
                            _ = fi47
                        }
                    case "hello":
                        v3[fi31] = -907
                    default:
                        v48 := GenS4{F0: 14}
                        _ = v48
                    }
                }
            }
        }
        v49 := []int{}
        _ = fi31
        _ = v32
        _ = v49
    }
    v50 := GenS2{F0: "\n\t", F1: (v30 / (((v30 * (v30 | (v30 << ((121) & 63))))) | 1)), F2: int64(v1)}
    for fi51 := 0; fi51 < 2; fi51++ {
        v52 := []int{4294967295, (fi51 << ((4294967295) & 63)), (fi51 + -803)}
        _ = fi51
        _ = v52
    }
    v30 -= v30
    v53 := GenS5{F0: 1, F1: true, F2: 0}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v50
    _ = v53
}
