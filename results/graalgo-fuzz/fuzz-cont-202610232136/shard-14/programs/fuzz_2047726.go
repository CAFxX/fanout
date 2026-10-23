package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 int16 }
type GenS2 struct { F0 uint8; F1 int64; F2 int64 }
type GenS3 struct { F0 string; F1 int32; F2 int64 }
type GenS4 struct { F0 int; F1 string }
func genFn0(p0 int8) int {
    var v7 bool = false
    for fi8 := 0; fi8 < 16; fi8++ {
        v9 := []int{}
        v10 := GenS1{F0: 18446744073709551615, F1: 73, F2: 32767}
        _ = fi8
        _ = v10
        _ = v9
    }
    switch ("" + ("a" + ("foo bar" + ""))) {
    case "a":
        v11 := []int{1, -1}
        _ = v11
    case "hello":
        for fi12 := 0; fi12 < 13; fi12++ {
            if false {
                p0 += p0
            }
            var v13 int8 = p0
            _ = fi12
            _ = v13
        }
    default:
        v7 = false
    }
    _ = p0
    _ = v7
    return 593
}

func main() {
    var v1 float64 = 1e10
    var v2 uint16 = 477
    var v3 int32 = 0
    v4 := make(map[string]bool)
    var v5 int8 = -76
    v6 := make(map[int]int)
    for fi14 := 0; fi14 < 4; fi14++ {
        var v15 uint16 = 57
        v3 -= (v3 - v3)
        fmt.Println(v2)
        _ = fi14
        _ = v15
    }
    for fi16 := 0; fi16 < 15; fi16++ {
        var v17 int64 = -15
        v17 = v17
        v18 := GenS1{F0: 18446744073709551615, F1: 255, F2: -1}
        _ = fi16
        _ = v17
        _ = v18
    }
    if true {
        if ((-1 > 42) || false) {
            if true {
                v19 := GenS2{F0: 42, F1: -9223372036854775808, F2: 42}
                delete(v6, 42)
                for fi20 := 0; fi20 < 17; fi20++ {
                    for fi21 := 0; fi21 < 16; fi21++ {
                        if true {
                            switch "hello" {
                            case "a":
                                var v22 uint32 = 2147483648
                                _ = v22
                            case "hello":
                                var v23 int16 = 0
                                _ = v23
                            case "foo bar":
                                var v24 uint8 = 255
                                var v25 int32 = (v3 ^ v3)
                                if (415 <= 4294967295) {
                                    v3 += v25
                                    fmt.Println(v25, v1, v25)
                                    fmt.Println(v1, v1)
                                } else {
                                    var v26 string = "a"
                                    switch fi21 {
                                    case -1:
                                        v27 := v6[42]
                                        _ = v27
                                    case 32768:
                                        v5 += v5
                                    default:
                                        var v28 uint16 = 1
                                        _ = v28
                                    }
                                    _ = v26
                                }
                                _ = v24
                                _ = v25
                            }
                            if false {
                                var v29 bool = ((1e10 + 0.0) != float32(v1))
                                _ = v29
                            } else {
                                switch fi20 {
                                case -423:
                                    v5 = (v5 + (v5 & 1))
                                case 523:
                                    delete(v4, ("hello" + "hello"))
                                }
                            }
                        } else {
                            var v30 uint64 = 42
                            var v31 int16 = -1
                            var v32 uint64 = 453
                            _ = v30
                            _ = v31
                            _ = v32
                        }
                        if ((v5 * (v5 + v5)) != v5) {
                            v33 := GenS3{F0: ("a" + "foo bar"), F1: v3, F2: -421}
                            fmt.Println(v3, v2, v2)
                            _ = v33
                        } else {
                            v34 := []int{genFn0(v5), genFn0(v5)}
                            _ = v34
                        }
                        _ = fi21
                    }
                    var v35 int32 = (^(v3))
                    _ = fi20
                    _ = v35
                }
                _ = v19
            }
            var v36 float32 = 0.0
            for fi37 := 0; fi37 < 14; fi37++ {
                var v38 int64 = 9223372036854775807
                v6[0] = (-(fi37))
                var v39 int32 = int32(fi37)
                _ = fi37
                _ = v38
                _ = v39
            }
            _ = v36
        }
        switch "foo bar" {
        case "\n\t":
            for fi40 := 0; fi40 < 4; fi40++ {
                switch ("foo bar" + "\n\t") {
                case "\n\t":
                    fmt.Println(v5)
                default:
                    if (1 > 2147483648) {
                        v4["foo bar"] = (!(((-2.25 / 1.5) - -2.25) == 0.0))
                        var v41 bool = true
                        _ = v41
                    } else {
                        for fi42 := 0; fi42 < 14; fi42++ {
                            v3 += v3
                            v43 := []int{4294967295}
                            _ = fi42
                            _ = v43
                        }
                        for fi44 := 0; fi44 < 5; fi44++ {
                            var v45 int32 = v3
                            _ = fi44
                            _ = v45
                        }
                    }
                }
                _ = fi40
            }
        case "":
            v2 += (v2 >> ((1) & 63))
        case "a":
            v5 -= -1
        }
    } else {
        v46 := v4[""]
        for fi47 := 0; fi47 < 3; fi47++ {
            v48 := []int{(fi47 + (fi47 & (fi47 ^ fi47))), fi47, (fi47 << ((65535) & 63)), fi47}
            _ = fi47
            _ = v48
        }
        _ = v46
    }
    var v49 string = "\n\t"
    v6[-570] = genFn0(v5)
    switch 32767 {
    case -1:
        v3 = (^(v3))
    case -2147483648:
        fmt.Println(v49)
    case 0:
        v3 -= v3
    }
    if (false && true) {
        v5 = v5
    } else {
        v49 = "foo bar"
        var v50 int16 = 32767
        _ = v50
    }
    v49 = v49
    var v51 float32 = 1.5
    if false {
        for fi52 := 0; fi52 < 17; fi52++ {
            v53 := []int{fi52, 571, (fi52 / ((fi52) | 1))}
            _ = fi52
            _ = v53
        }
        if true {
            fmt.Println(v51, v5, v3)
            fmt.Println(v2, v5)
            var v54 string = ("hello" + (("" + ("hello" + ("\n\t" + v49))) + (("hello" + "foo bar") + (("hello" + v49) + "\n\t"))))
            _ = v54
        }
        fmt.Println(v1)
    }
    v49 = (v49 + v49)
    v5 -= v5
    if (179 < 42) {
        fmt.Println(v5, v2, v5)
        v3 = 793
        var v55 string = ""
        _ = v55
    }
    v56 := GenS4{F0: 9223372036854775807, F1: ""}
    var v57 bool = false
    var v58 string = "hello"
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("v57:", v57)
    fmt.Println("v58:", v58)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v51
    _ = v56
    _ = v57
    _ = v58
    _ = v6
}
