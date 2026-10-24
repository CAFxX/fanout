package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int32; F1 int32; F2 uint8 }
type GenS4 struct { F0 uint; F1 int16; F2 int32 }
type GenS5 struct { F0 bool; F1 bool; F2 uint64 }
type GenS6 struct { F0 string }
func genFn0(p0 bool, p1 bool) uint8 {
    for fi5 := 0; fi5 < 1; fi5++ {
        var v6 float64 = ((0.0 + 1e10) + 1.5)
        _ = fi5
        _ = v6
    }
    var v7 int8 = 2
    p1 = p0
    for fi8 := 0; fi8 < 9; fi8++ {
        v9 := GenS1{F0: 1, F1: v7}
        _ = fi8
        _ = v9
    }
    p0 = ("a" <= "a")
    _ = p0
    _ = p1
    _ = v7
    return 128
}

func main() {
    v1 := make(map[int]bool)
    var v2 float64 = 1e10
    var v3 float64 = (v2 * 1e10)
    v4 := make(map[string]int)
    var v10 uint64 = 1
    v3 += v3
    fmt.Println(v10, v2)
    v10 = 2
    if (-1 < -2147483648) {
        var v11 int8 = -47
        _ = v11
    } else {
        for fi12 := 0; fi12 < 20; fi12++ {
            v13 := GenS1{F0: -32768, F1: 1}
            var v14 uint8 = 42
            v14 -= genFn0((1 != 818), (-828 < 283))
            _ = fi12
            _ = v13
            _ = v14
        }
    }
    if (3.14159 >= 3.14159) {
        if true {
            var v15 int = -805
            v10 -= (v10 >> ((343) & 63))
            _ = v15
        } else {
            fmt.Println(v3, v10)
            for fi16 := 0; fi16 < 3; fi16++ {
                v17 := []int{32767, (fi16 * fi16)}
                _ = fi16
                _ = v17
            }
        }
    } else {
        for fi18 := 0; fi18 < 12; fi18++ {
            v10 = v10
            v19 := []int{fi18, 32767, (fi18 * (fi18 | (fi18 | (fi18 | 127)))), fi18}
            _ = len(v4)
            _ = fi18
            _ = v19
        }
        if false {
            if false {
                if (2147483648 != 2147483648) {
                    if true {
                        for fi20 := 0; fi20 < 5; fi20++ {
                            switch (fi20 / ((fi20) | 1)) {
                            case 32767:
                                var v21 int64 = 1
                                _ = v21
                            case 32768:
                                switch fi20 {
                                case -770:
                                    for fi22 := 0; fi22 < 2; fi22++ {
                                        v4["foo bar"] = 2147483647
                                        var v23 int = (-(fi22))
                                        var v24 float64 = v3
                                        _ = fi22
                                        _ = v23
                                        _ = v24
                                    }
                                case 9223372036854775807:
                                    v10 += (v10 >> ((232) & 63))
                                case -2147483648:
                                    v10 = 18446744073709551615
                                }
                            case -286:
                                for fi25 := 0; fi25 < 15; fi25++ {
                                    var v26 uint32 = 65535
                                    v27 := []int{fi20, (fi20 >> ((65535) & 63))}
                                    _ = fi25
                                    _ = v26
                                    _ = v27
                                }
                            default:
                                v28 := []int{(fi20 - (^(fi20))), fi20, (fi20 << ((255) & 63))}
                                _ = v28
                            }
                            var v29 string = "hello"
                            _ = fi20
                            _ = v29
                        }
                        for fi30 := 0; fi30 < 16; fi30++ {
                            var v31 uint16 = 0
                            _ = fi30
                            _ = v31
                        }
                    }
                }
                fmt.Println(v10)
            }
        } else {
            var v32 int16 = 646
            v33 := []int{2147483647, 2147483647}
            _ = v32
            _ = v33
        }
        v2 = (v3 + (0.0 - v2))
    }
    var v34 float32 = 3.14159
    switch 166 {
    case 32768:
        v3 -= v3
    }
    switch "a" {
    case "a":
        v35 := GenS1{F0: -1, F1: 0}
        _ = v35
    }
    for fi36 := 0; fi36 < 11; fi36++ {
        v37 := v1[(fi36 + (fi36 + (fi36 << ((42) & 63))))]
        _ = fi36
        _ = v37
    }
    fmt.Println(v2)
    var v38 int32 = 872
    var v39 float64 = v3
    var v40 string = "\n\t"
    switch (v40 + (("foo bar" + v40) + "hello")) {
    case "hello":
        fmt.Println(v2)
    case "\n\t":
        delete(v1, -9223372036854775808)
    }
    var v41 uint16 = 32767
    for fi42 := 0; fi42 < 16; fi42++ {
        switch ("foo bar" + "\n\t") {
        case "foo bar":
            if false {
                v43 := GenS2{F0: genFn0((!false), (!true))}
                _ = v43
            }
        default:
            for fi44 := 0; fi44 < 12; fi44++ {
                _ = len(v4)
                if true {
                    if true {
                        v45 := GenS3{F0: (v38 >> ((0) & 63)), F1: -784, F2: 255}
                        switch (-(fi42)) {
                        case -259:
                            if (((1.5 + v34) >= (v34 / v34)) || false) {
                                fmt.Println(fi42)
                                switch (v40 + v40) {
                                case "foo bar":
                                    v46 := GenS4{F0: 979, F1: -1, F2: (v38 % (((v38 % (((-(v38))) | 1))) | 1))}
                                    _ = v46
                                default:
                                    if false {
                                        fmt.Println(v38, v2)
                                    }
                                }
                                v47 := v4[v40]
                                _ = v47
                            }
                        case 42:
                            if false {
                                v48 := GenS2{F0: 127}
                                v3 += v2
                                fmt.Println(fi44)
                                _ = v48
                            } else {
                                if true {
                                    v49 := GenS2{F0: 127}
                                    _ = v49
                                } else {
                                    v50 := GenS5{F0: true, F1: (!((!false) || (-1 >= 415))), F2: v10}
                                    _ = v50
                                }
                                v51 := GenS6{F0: "hello"}
                                _ = v51
                            }
                        }
                        if true {
                            var v52 uint16 = v41
                            fmt.Println(v39, v40)
                            _ = v52
                        }
                        var v53 float64 = v2
                        _ = v45
                        _ = v53
                    }
                } else {
                    var v54 int32 = v38
                    var v55 float32 = (float32(1.5) / (v34 * ((v34 - v34) + 0.1)))
                    _ = v54
                    _ = v55
                }
                _ = fi44
            }
        }
        var v56 float64 = float64((v34 + v34))
        var v57 int = fi42
        _ = fi42
        _ = v56
        _ = v57
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v34
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
}
