package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint; F2 uint16 }
type GenS2 struct { F0 int32; F1 int16 }
type GenS3 struct { F0 uint16; F1 uint32 }
type GenS4 struct { F0 bool }
func genFn0() string {
    v4 := GenS1{F0: "", F1: 4294967295, F2: 46}
    v5 := GenS2{F0: -1, F1: 251}
    _ = v4
    _ = v5
    return "a"
}
func genFn1() int8 {
    for fi6 := 0; fi6 < 2; fi6++ {
        var v7 uint32 = 1
        if ((-596 != 42) && true) {
            for fi8 := 0; fi8 < 4; fi8++ {
                var v9 uint16 = 42
                _ = fi8
                _ = v9
            }
            v7 = v7
            v7 -= v7
        }
        _ = fi6
        _ = v7
    }
    switch 1 {
    case 128:
        v10 := GenS2{F0: 44, F1: 42}
        _ = v10
    case 223:
        v11 := []int{}
        _ = v11
    }
    var v12 float32 = ((3.14159 - 1.5) / float32(float64(-2.25)))
    fmt.Println(v12)
    var v13 uint64 = 59
    _ = v12
    _ = v13
    return 42
}

func main() {
    var v1 int16 = 0
    var v2 int32 = -133
    var v3 uint64 = 1
    var v14 uint16 = 2
    for fi15 := 0; fi15 < 15; fi15++ {
        v16 := []int{-9223372036854775808, -261, 32768, fi15}
        fmt.Println(v2)
        _ = len(v16)
        _ = fi15
        _ = v16
    }
    switch "a" {
    case "":
        v2 -= v2
    case "\n\t":
        v17 := GenS2{F0: (v2 >> ((65535) & 63)), F1: -1}
        _ = v17
    case "a":
        switch "\n\t" {
        case "":
            if (42 <= 42) {
                var v18 uint32 = 0
                v14 -= (v14 - v14)
                v2 = (^(v2))
                _ = v18
            } else {
                v2 += v2
                for fi19 := 0; fi19 < 18; fi19++ {
                    var v20 uint = 0
                    if (v2 < 2147483647) {
                        var v21 int = (fi19 << (((v20 & v20)) & 63))
                        if false {
                            v3 = 4294967295
                            for fi22 := 0; fi22 < 2; fi22++ {
                                v14 = v14
                                fmt.Println(v1, v2)
                                _ = fi22
                            }
                            if false {
                                if ("foo bar" >= ("hello" + "\n\t")) {
                                    if (-9223372036854775808 <= 1) {
                                        var v23 uint64 = 18446744073709551615
                                        for fi24 := 0; fi24 < 9; fi24++ {
                                            v23 = v3
                                            v23 = 42
                                            v25 := []int{int(v14), 42, int(v20)}
                                            _ = fi24
                                            _ = v25
                                        }
                                        v26 := []int{fi19, -339, 647, (v21 / ((v21) | 1))}
                                        _ = v23
                                        _ = v26
                                    } else {
                                        if ((uint8(v14) < 0) || ((true && true) && false)) {
                                            for fi27 := 0; fi27 < 15; fi27++ {
                                                var v28 uint32 = uint32(v20)
                                                v29 := GenS1{F0: ("foo bar" + "a"), F1: 18446744073709551615, F2: (v14 >> ((v20) & 63))}
                                                var v30 uint8 = 120
                                                switch "hello" {
                                                case "\n\t":
                                                    v21 = -9223372036854775808
                                                    var v31 bool = ((v28 & v28) > 1)
                                                    _ = v31
                                                }
                                                if (-201 != (v21 + (fi19 ^ -9223372036854775808))) {
                                                    v2 = v2
                                                } else {
                                                    var v32 int8 = -111
                                                    var v33 uint16 = 0
                                                    switch "hello" {
                                                    case "foo bar":
                                                        fmt.Println(v30, v1, v2)
                                                    }
                                                    var v34 uint8 = (v30 - (v30 >> (((v20 / (((v20 | (v20 * (v20 & (v20 << ((v20) & 63)))))) | 1))) & 63)))
                                                    var v35 uint8 = (v30 * (^(v30)))
                                                    v36 := []int{(-(v21))}
                                                    v1 = v1
                                                    _ = v32
                                                    _ = v33
                                                    _ = v34
                                                    _ = v35
                                                    _ = v36
                                                }
                                                _ = fi27
                                                _ = v28
                                                _ = v29
                                                _ = v30
                                            }
                                        }
                                    }
                                } else {
                                    for fi37 := 0; fi37 < 9; fi37++ {
                                        v38 := []int{(^(fi19)), 127, fi19}
                                        v38 = append(v38, 245)
                                        var v39 int32 = (^(v2))
                                        _ = fi37
                                        _ = v38
                                        _ = v39
                                    }
                                    v2 += v2
                                }
                                var v40 bool = ((((v3 % ((734) | 1)) > v3) && true) && ((v1 >> ((65535) & 63)) <= int16(v21)))
                                _ = v40
                            }
                        } else {
                            v41 := []int{fi19}
                            _ = v41
                        }
                        switch "hello" {
                        case "foo bar":
                            v14 += 2
                        case "\n\t":
                            for fi42 := 0; fi42 < 10; fi42++ {
                                if true {
                                    v43 := GenS2{F0: v2, F1: (v1 - v1)}
                                    v44 := []int{}
                                    var v45 uint8 = 42
                                    switch ("foo bar" + "hello") {
                                    case "foo bar":
                                        fmt.Println(fi19)
                                        _ = len(v44)
                                    default:
                                        fmt.Println(v14)
                                    }
                                    _ = v43
                                    _ = v44
                                    _ = v45
                                } else {
                                    var v46 uint16 = v14
                                    var v47 float64 = 1.5
                                    _ = v46
                                    _ = v47
                                }
                                v21 = fi19
                                fmt.Println(fi19)
                                _ = fi42
                            }
                        case "hello":
                            var v48 uint16 = v14
                            _ = v48
                        default:
                            for fi49 := 0; fi49 < 7; fi49++ {
                                if true {
                                    var v50 string = (("a" + "hello") + "")
                                    v51 := GenS3{F0: 512, F1: 824}
                                    _ = v50
                                    _ = v51
                                }
                                fmt.Println(v21, v2, v14)
                                v52 := GenS4{F0: false}
                                _ = fi49
                                _ = v52
                            }
                        }
                        _ = v21
                    }
                    _ = fi19
                    _ = v20
                }
            }
        case "\n\t":
            for fi53 := 0; fi53 < 2; fi53++ {
                switch ("hello" + "a") {
                case "hello":
                    for fi54 := 0; fi54 < 16; fi54++ {
                        fmt.Println(v14, v3)
                        _ = fi54
                    }
                case "a":
                    for fi55 := 0; fi55 < 15; fi55++ {
                        var v56 float32 = (0.0 * ((-2.25 / 1.5) + 0.0))
                        _ = fi55
                        _ = v56
                    }
                case "\n\t":
                    var v57 uint64 = v3
                    _ = v57
                default:
                    v58 := []int{(-(fi53)), (fi53 - 42), fi53}
                    _ = v58
                }
                for fi59 := 0; fi59 < 17; fi59++ {
                    var v60 uint16 = v14
                    fmt.Println(fi53, v1)
                    _ = fi59
                    _ = v60
                }
                _ = fi53
            }
        }
    }
    var v61 float32 = 0.1
    if false {
        for fi62 := 0; fi62 < 13; fi62++ {
            var v63 int16 = v1
            v64 := []int{int(v63), fi62}
            v2 = (v2 ^ 42)
            _ = fi62
            _ = v63
            _ = v64
        }
    } else {
        fmt.Println(v61)
    }
    var v65 uint32 = 2147483647
    var v66 bool = (1.5 <= (float32(-2.25) - v61))
    v1 = v1
    v66 = true
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v61:", v61)
    fmt.Println("v65:", v65)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v61
    _ = v65
    _ = v66
}
