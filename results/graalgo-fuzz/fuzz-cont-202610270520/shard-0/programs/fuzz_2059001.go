package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string; F2 int32 }
type GenS2 struct { F0 uint64; F1 uint64; F2 uint16 }
type GenS3 struct { F0 uint; F1 int16 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 bool; F1 uint16 }
type GenS7 struct { F0 int64; F1 int8 }
type GenS8 struct { F0 uint16; F1 string; F2 uint }
func genFn0(p0 int64, p1 uint64, p2 uint64) int32 {
    var v6 bool = (-128 >= 0)
    v6 = (!v6)
    p2 = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return -1
}
func genFn1() string {
    if true {
        v7 := GenS1{F0: 4294967295, F1: "foo bar", F2: 97}
        if true {
            v8 := []int{-221, 32768, -345}
            _ = v8
        } else {
            var v9 int32 = -1
            if true {
                var v10 int8 = 100
                _ = v10
            }
            _ = v9
        }
        var v11 int16 = -266
        _ = v11
        _ = v7
    }
    var v12 int = -58
    var v13 uint16 = 975
    _ = v12
    _ = v13
    return "\n\t"
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[uint32]string)
    var v3 uint32 = 0
    var v4 int32 = -1
    v5 := make(map[string]uint32)
    v3 = v3
    var v14 int64 = 42
    if false {
        v15 := GenS2{F0: 4294967295, F1: 229, F2: 0}
        _ = v15
    } else {
        var v16 int8 = 1
        _ = v16
    }
    for fi17 := 0; fi17 < 7; fi17++ {
        if (-2.25 != 1.5) {
            v3 += 2147483647
            for fi18 := 0; fi18 < 15; fi18++ {
                v4 -= v4
                v4 = -2147483648
                v19 := GenS1{F0: (v3 & (v3 % (((^(v3))) | 1))), F1: "hello", F2: v4}
                _ = fi18
                _ = v19
            }
        } else {
            v3 -= (v3 / ((971) | 1))
        }
        _ = fi17
    }
    if false {
        for fi20 := 0; fi20 < 1; fi20++ {
            if (((811 != 543) && true) || false) {
                var v21 uint64 = 0
                _ = v21
            }
            _ = fi20
        }
        fmt.Println(v3)
        for fi22 := 0; fi22 < 16; fi22++ {
            switch "hello" {
            case "foo bar":
                v23 := GenS3{F0: 65535, F1: 90}
                _ = v23
            case "":
                var v24 uint16 = 32767
                _ = v24
            }
            fmt.Println(fi22, v3)
            if true {
                if false {
                    var v25 string = "a"
                    for fi26 := 0; fi26 < 12; fi26++ {
                        var v27 uint8 = 160
                        v27 += v27
                        _ = fi26
                        _ = v27
                    }
                    _ = v25
                } else {
                    for fi28 := 0; fi28 < 14; fi28++ {
                        var v29 uint8 = 2
                        v14 = (v14 << ((4294967295) & 63))
                        switch fi28 {
                        case -415:
                            v14 = v14
                        }
                        _ = fi28
                        _ = v29
                    }
                    if false {
                        if (362 <= 840) {
                            if (!false) {
                                v14 -= (v14 + v14)
                                v3 = (^(v3))
                                fmt.Println(fi22, v3, v3)
                            }
                            v3 = v3
                            var v30 uint8 = 42
                            _ = v30
                        } else {
                            var v31 uint = 255
                            if (!true) {
                                v32 := GenS4{F0: 1000}
                                for fi33 := 0; fi33 < 3; fi33++ {
                                    v34 := GenS4{F0: 1}
                                    _ = fi33
                                    _ = v34
                                }
                                v35 := GenS5{F0: fi22}
                                _ = v32
                                _ = v35
                            } else {
                                var v36 int8 = -1
                                _ = v36
                            }
                            for fi37 := 0; fi37 < 6; fi37++ {
                                fmt.Println(v31, v3)
                                if (!(42 < 65535)) {
                                    var v38 uint = (^(v31))
                                    v39 := GenS4{F0: 1}
                                    var v40 int16 = 1000
                                    _ = v38
                                    _ = v39
                                    _ = v40
                                } else {
                                    v3 += (v3 & v3)
                                    var v41 uint64 = 0
                                    fmt.Println(v14)
                                    _ = v41
                                }
                                _ = fi37
                            }
                            _ = v31
                        }
                        var v42 uint = 65535
                        v4 += (v4 ^ v4)
                        _ = v42
                    }
                }
            } else {
                var v43 int16 = -1
                fmt.Println(fi22, fi22)
                _ = v43
            }
            _ = fi22
        }
    } else {
        if true {
            var v44 uint8 = 128
            _ = v44
        } else {
            if (!false) {
                v45 := GenS1{F0: (v3 & 2), F1: "hello", F2: (v4 * genFn0((v14 & (v14 / (((-(v14))) | 1))), 1, 18446744073709551615))}
                v46 := GenS6{F0: (-2.25 != 1.5), F1: 1}
                _ = v45
                _ = v46
            }
        }
    }
    switch -59 {
    case -195:
        switch 9223372036854775807 {
        case -41:
            v5["foo bar"] = v3
        case 128:
            switch "\n\t" {
            case "\n\t":
                v3 = 735
            case "a":
                fmt.Println(v3)
            default:
                for fi47 := 0; fi47 < 16; fi47++ {
                    var v48 bool = ((((float64(0.1) == 1.5) && (65535 > 0)) || false) || (0 == 0))
                    for fi49 := 0; fi49 < 15; fi49++ {
                        v48 = ((fi47 ^ fi49) >= (fi47 ^ int(v4)))
                        _ = fi49
                    }
                    _ = fi47
                    _ = v48
                }
            }
        default:
            v14 -= (v14 / (((-(v14))) | 1))
        }
    case 264:
        for fi50 := 0; fi50 < 4; fi50++ {
            v51 := []int{fi50}
            if true {
                var v52 string = (("" + "\n\t") + ("\n\t" + "foo bar"))
                _ = v52
            } else {
                if true {
                    v3 = 42
                    v53 := GenS5{F0: fi50}
                    _ = v53
                } else {
                    var v54 string = ("hello" + "")
                    _ = v54
                }
            }
            _ = fi50
            _ = v51
        }
    }
    v4 = (^(v4))
    switch "foo bar" {
    case "\n\t":
        v4 = (-(v4))
    case "":
        v14 = 1
    default:
        v55 := GenS7{F0: (v14 - v14), F1: -112}
        _ = v55
    }
    v56 := GenS8{F0: 1, F1: (("\n\t" + "\n\t") + (("\n\t" + "foo bar") + "\n\t")), F2: 65535}
    v14 = 42
    v57 := GenS7{F0: v14, F1: -1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v56
    _ = v57
}
