package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint16; F1 uint16; F2 uint }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint8; F1 uint; F2 int64 }
type GenS5 struct { F0 uint; F1 uint }
type GenS6 struct { F0 int64; F1 int64; F2 int16 }
type GenS7 struct { F0 uint32; F1 string; F2 int }
func genFn0(p0 uint64, p1 int, p2 int8) int32 {
    v4 := GenS1{F0: 18446744073709551615}
    for fi5 := 0; fi5 < 4; fi5++ {
        p0 -= p0
        if (176 >= 0) {
            switch ("" + "foo bar") {
            case "foo bar":
                v6 := []int{p1, (p1 ^ fi5), (^(p1))}
                _ = v6
            case "a":
                for fi7 := 0; fi7 < 10; fi7++ {
                    v8 := []int{(fi7 & p1), fi5, 32767}
                    _ = fi7
                    _ = v8
                }
            }
        } else {
            switch "" {
            case "hello":
                v9 := []int{p1}
                _ = v9
            case "":
                p2 += (p2 | (p2 >> ((1) & 63)))
            }
        }
        _ = fi5
    }
    var v10 float32 = 0.0
    var v11 int = (p1 % ((p1) | 1))
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v4
    return -2147483648
}
func genFn1() int32 {
    var v12 uint = 2
    fmt.Println(v12)
    var v13 string = "foo bar"
    for fi14 := 0; fi14 < 19; fi14++ {
        if true {
            v15 := GenS1{F0: 1}
            for fi16 := 0; fi16 < 17; fi16++ {
                switch fi14 {
                case 128:
                    var v17 float32 = (0.1 - 3.14159)
                    v18 := []int{fi14, (fi16 | 0)}
                    _ = v17
                    _ = v18
                case 170:
                    v12 = 0
                    v13 = ""
                    v19 := GenS2{F0: 65535, F1: 255, F2: (v12 + (v12 + 4294967295))}
                    _ = v19
                case 0:
                    v13 = v13
                default:
                    v13 = (v13 + "")
                    v20 := GenS3{F0: (v12 << (((v12 + 1)) & 63))}
                    _ = v20
                }
                _ = fi16
            }
            _ = v15
        }
        fmt.Println(v12, v13)
        v21 := []int{}
        _ = fi14
        _ = v21
    }
    _ = v12
    _ = v13
    return -1
}

func main() {
    var v1 int8 = 1
    var v2 int = -55
    var v3 int64 = -9223372036854775808
    if false {
        for fi22 := 0; fi22 < 2; fi22++ {
            if (true || ((v3 ^ v3) <= (v3 / (((v3 % (((v3 ^ (v3 + (v3 ^ (-(v3)))))) | 1))) | 1)))) {
                v23 := GenS2{F0: 2, F1: 2, F2: 2}
                v2 += v2
                for fi24 := 0; fi24 < 5; fi24++ {
                    if true {
                        var v25 int32 = 0
                        fmt.Println(v3, v2)
                        if false {
                            var v26 uint16 = 540
                            _ = v26
                        } else {
                            fmt.Println(v3)
                            v3 -= (v3 + int64(v25))
                            if true {
                                v1 += 24
                            }
                        }
                        _ = v25
                    }
                    _ = fi24
                }
                _ = v23
            }
            v2 += v2
            for fi27 := 0; fi27 < 7; fi27++ {
                var v28 uint8 = 110
                _ = fi27
                _ = v28
            }
            _ = fi22
        }
        switch "\n\t" {
        case "a":
            if false {
                fmt.Println(v1)
                v3 = (^(v3))
                v29 := GenS4{F0: 1, F1: 661, F2: v3}
                _ = v29
            }
        default:
            v2 += (v2 & 32768)
        }
        var v30 bool = true
        _ = v30
    } else {
        for fi31 := 0; fi31 < 10; fi31++ {
            var v32 int = fi31
            v32 = (v32 % ((-1) | 1))
            for fi33 := 0; fi33 < 15; fi33++ {
                switch "\n\t" {
                case "a":
                    var v34 uint = 143
                    _ = v34
                case "foo bar":
                    switch (-(v32)) {
                    case -2147483648:
                        if true {
                            var v35 uint64 = 0
                            _ = v35
                        }
                    case 127:
                        v1 = (v1 * v1)
                    default:
                        v36 := GenS3{F0: 0}
                        _ = v36
                    }
                case "\n\t":
                    if false {
                        switch "hello" {
                        case "\n\t":
                            if (559 <= 255) {
                                var v37 uint = 963
                                _ = v37
                            } else {
                                if true {
                                    v38 := GenS3{F0: 2}
                                    _ = v38
                                } else {
                                    var v39 uint16 = 1
                                    _ = v39
                                }
                            }
                        default:
                            v40 := []int{}
                            _ = v40
                        }
                        v41 := []int{v32, (-(v32)), fi31}
                        if true {
                            fmt.Println(fi31)
                            v42 := GenS5{F0: 42, F1: 1}
                            _ = v42
                        }
                        _ = v41
                    }
                default:
                    var v43 uint16 = 346
                    _ = v43
                }
                var v44 float32 = float32(3.14159)
                _ = fi33
                _ = v44
            }
            _ = fi31
            _ = v32
        }
        fmt.Println(v2, v3)
    }
    var v45 uint64 = 0
    v46 := []int{(v2 << ((319) & 63)), 0}
    for fi47 := 0; fi47 < 1; fi47++ {
        v3 += (v3 & 918)
        _ = fi47
    }
    if false {
        if false {
            for fi48 := 0; fi48 < 16; fi48++ {
                if (0 <= (v2 - fi48)) {
                    v45 -= v45
                    v46 = append(v46, v2)
                    v2 += fi48
                }
                var v49 bool = (-840 <= 32767)
                _ = fi48
                _ = v49
            }
            v3 -= v3
        } else {
            var v50 uint64 = v45
            _ = v50
        }
        v46 = append(v46, (^(v2)))
        v51 := GenS6{F0: v3, F1: (v3 - (v3 * (-(v3)))), F2: -544}
        _ = v51
    } else {
        var v52 bool = false
        fmt.Println(v3, v45)
        for fi53 := 0; fi53 < 10; fi53++ {
            _ = len(v46)
            var v54 uint16 = 1
            v2 = 9223372036854775807
            _ = fi53
            _ = v54
        }
        _ = v52
    }
    fmt.Println(v1)
    fmt.Println(v3)
    v1 = (v1 + v1)
    fmt.Println(v2, v3)
    var v55 float64 = 0.1
    switch "" {
    case "hello":
        var v56 string = "hello"
        _ = v56
    case "":
        v57 := GenS7{F0: 723, F1: "\n\t", F2: 9223372036854775807}
        _ = v57
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v45:", v45)
    fmt.Println("slicelen v46:", len(v46), v46)
    fmt.Println("v55:", v55)
    _ = v1
    _ = v2
    _ = v3
    _ = v45
    _ = v46
    _ = v55
}
