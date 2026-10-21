package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint32; F1 uint32; F2 uint }
type GenS3 struct { F0 uint; F1 uint }
type GenS4 struct { F0 uint64; F1 uint32; F2 uint32 }
type GenS5 struct { F0 int; F1 int64; F2 bool }
func genFn0() int32 {
    var v5 int16 = 42
    v5 -= (v5 - (v5 << ((458) & 63)))
    for fi6 := 0; fi6 < 11; fi6++ {
        if true {
            if (false && ((76 >= 2) && false)) {
                v5 += (v5 % ((v5) | 1))
                v7 := GenS1{F0: ("" + "\n\t")}
                v8 := GenS1{F0: "hello"}
                _ = v7
                _ = v8
            } else {
                v5 = 1000
                if ((92 == 0) && false) {
                    var v9 uint64 = 0
                    _ = v9
                } else {
                    switch fi6 {
                    case 4294967295:
                        v10 := GenS1{F0: "foo bar"}
                        _ = v10
                    case 307:
                        var v11 uint16 = 864
                        _ = v11
                    default:
                        v5 -= (v5 + 1000)
                    }
                }
            }
            v12 := []int{756, 1, (fi6 >> ((42) & 63))}
            _ = v12
        }
        v13 := []int{}
        if false {
            if false {
                v5 -= 42
                var v14 uint8 = 10
                if false {
                    for fi15 := 0; fi15 < 18; fi15++ {
                        if ("a" == "hello") {
                            var v16 int64 = 42
                            _ = v16
                        } else {
                            switch ("a" + "a") {
                            case "":
                                v17 := GenS1{F0: ""}
                                _ = v17
                            case "\n\t":
                                v18 := GenS2{F0: 42, F1: 893, F2: 1}
                                _ = v18
                            case "hello":
                                v14 = v14
                            default:
                                var v19 uint = 2
                                _ = v19
                            }
                        }
                        _ = fi15
                    }
                } else {
                    for fi20 := 0; fi20 < 12; fi20++ {
                        var v21 uint64 = 234
                        if (-2.25 > 3.14159) {
                            v22 := GenS2{F0: 65535, F1: 2147483647, F2: 4294967295}
                            _ = v22
                        }
                        _ = fi20
                        _ = v21
                    }
                }
                _ = v14
            } else {
                v5 -= (v5 >> ((734) & 63))
            }
            v5 += v5
        }
        _ = fi6
        _ = v13
    }
    var v23 int = -874
    fmt.Println(v5)
    _ = v23
    _ = v5
    return 1
}
func genFn1(p0 uint, p1 uint64) uint {
    for fi24 := 0; fi24 < 4; fi24++ {
        v25 := []int{(fi24 | fi24), (fi24 << (((^(p0))) & 63)), (-(fi24))}
        if len(v25) > 0 {
            v25[0] = (-(fi24))
        }
        _ = fi24
        _ = v25
    }
    switch "" {
    case "hello":
        var v26 uint64 = p1
        _ = v26
    case "":
        fmt.Println(p1, p1)
    case "\n\t":
        for fi27 := 0; fi27 < 19; fi27++ {
            v28 := []int{}
            switch -2147483648 {
            case -3:
                fmt.Println(p0, fi27, p1)
            default:
                for fi29 := 0; fi29 < 13; fi29++ {
                    var v30 uint = p0
                    v31 := GenS2{F0: 2147483647, F1: 412, F2: (p0 ^ (p0 / (((v30 * (^(p0)))) | 1)))}
                    _ = fi29
                    _ = v30
                    _ = v31
                }
            }
            _ = fi27
            _ = v28
        }
    default:
        for fi32 := 0; fi32 < 2; fi32++ {
            switch "\n\t" {
            case "foo bar":
                fmt.Println(p0, p0, p1)
            default:
                var v33 uint = p0
                _ = v33
            }
            p0 = p0
            p1 -= 42
            _ = fi32
        }
    }
    _ = p0
    _ = p1
    return p0
}

func main() {
    var v1 int = -37
    v2 := make(map[string]uint32)
    var v3 float64 = 0.0
    v4 := []int{(^(v1)), (v1 << ((2) & 63)), (^(v1)), 385}
    if ("a" > "foo bar") {
        var v34 uint64 = 0
        v34 = v34
        _ = v34
    } else {
        var v35 bool = false
        _ = v35
    }
    for fi36 := 0; fi36 < 1; fi36++ {
        v3 = 0.1
        _ = fi36
    }
    var v37 int8 = 127
    var v38 uint64 = 0
    if (true && false) {
        var v39 uint16 = 32767
        _ = v39
    } else {
        var v40 int32 = -856
        _ = v40
    }
    fmt.Println(v3, v38)
    v41 := GenS3{F0: 210, F1: 18446744073709551615}
    v42 := GenS4{F0: 4294967295, F1: 2147483648, F2: 4294967295}
    fmt.Println(v37, v1, v1)
    v3 = float64(1.5)
    var v43 string = ""
    v44 := GenS5{F0: 1, F1: int64(v1), F2: true}
    v43 = (((v43 + v43) + "foo bar") + "\n\t")
    if false {
        fmt.Println(v43, v37)
        v45 := GenS4{F0: (^(v38)), F1: 962, F2: 187}
        _ = v45
    } else {
        switch v1 {
        case 253:
            v46 := GenS5{F0: 1, F1: 1, F2: (2147483647 != 557)}
            _ = v46
        case 462:
            _ = len(v2)
        case -9223372036854775808:
            var v47 int64 = 42
            _ = v47
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v44
}
