package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 bool; F1 uint16 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int; F1 uint32; F2 string }
type GenS6 struct { F0 uint16; F1 int64 }
type GenS7 struct { F0 uint16; F1 uint64 }
func genFn0(p0 bool, p1 int64, p2 int) uint16 {
    var v5 float64 = (-2.25 / ((1e10 / 1.0) - (((1.5 - ((-2.25 * ((0.1 - 0.1) * 3.14159)) / 1e10)) / -2.25) / (3.14159 / 1e10))))
    v6 := GenS1{F0: 2}
    if true {
        v7 := GenS1{F0: 0}
        var v8 uint32 = 361
        var v9 uint = 65535
        _ = v7
        _ = v8
        _ = v9
    } else {
        v10 := []int{}
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 218
}
func genFn1(p0 int64) string {
    if (0 == 42) {
        v11 := GenS1{F0: 0}
        _ = v11
    } else {
        p0 = (-(p0))
        for fi12 := 0; fi12 < 9; fi12++ {
            var v13 int = (fi12 >> ((18446744073709551615) & 63))
            var v14 bool = false
            _ = fi12
            _ = v13
            _ = v14
        }
    }
    if ((0 != 1) && (42 >= 78)) {
        if false {
            var v15 uint8 = 167
            _ = v15
        } else {
            var v16 uint32 = 597
            p0 += -9223372036854775808
            v16 += (^(v16))
            _ = v16
        }
    } else {
        p0 = (p0 + p0)
        v17 := GenS1{F0: 0}
        p0 += p0
        _ = v17
    }
    _ = p0
    return "a"
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint = 4294967295
    v3 := make(map[uint32]uint32)
    v4 := make(map[string]uint32)
    v2 = v2
    _ = len(v1)
    v2 += (v2 + v2)
    var v18 uint = v2
    v2 = 42
    fmt.Println(v2, v18)
    switch "a" {
    case "\n\t":
        fmt.Println(v2, v2)
    case "":
        if true {
            if (42 <= -1) {
                if ((v18 ^ v18) == (v18 & 195)) {
                    _ = len(v3)
                    if false {
                        v1[uint32(v2)] = 32767
                    }
                } else {
                    v2 = 0
                }
                v2 += v18
                if true {
                    v19 := GenS2{F0: 65535}
                    switch 32768 {
                    case 42:
                        if ((!((-9223372036854775808 != 32767) && true)) || ((((317 < 2147483647) && (42 <= 177)) || (!(!true))) || false)) {
                            v18 = 2
                        } else {
                            v20 := []int{-1}
                            var v21 float32 = 0.1
                            var v22 float64 = (((1e10 + 1.5) + -2.25) / -2.25)
                            _ = v20
                            _ = v21
                            _ = v22
                        }
                    case -2147483648:
                        var v23 float32 = ((1.5 / 1.5) / 3.14159)
                        _ = v23
                    case -1:
                        v2 += (v18 >> ((v2) & 63))
                    }
                    _ = v19
                }
            } else {
                v18 -= (v2 ^ 2)
                switch 128 {
                case -1:
                    switch 39 {
                    case 4294967295:
                        var v24 int = 29
                        _ = v24
                    case -1:
                        v25 := GenS1{F0: v18}
                        _ = v25
                    default:
                        switch "" {
                        case "foo bar":
                            if false {
                                for fi26 := 0; fi26 < 18; fi26++ {
                                    v27 := GenS2{F0: 496}
                                    _ = fi26
                                    _ = v27
                                }
                                fmt.Println(v18, v2)
                                fmt.Println(v2, v2)
                            } else {
                                fmt.Println(v18)
                                switch "\n\t" {
                                case "foo bar":
                                    switch "a" {
                                    case "foo bar":
                                        var v28 int8 = -47
                                        _ = v28
                                    case "hello":
                                        if false {
                                            if false {
                                                v29 := GenS3{F0: true, F1: 721}
                                                _ = v29
                                            } else {
                                                v30 := GenS3{F0: false, F1: 2}
                                                _ = v30
                                            }
                                            if (2147483647 != 1) {
                                                v31 := GenS4{F0: (!false)}
                                                for fi32 := 0; fi32 < 1; fi32++ {
                                                    var v33 string = "\n\t"
                                                    v33 = ""
                                                    v2 = 0
                                                    _ = fi32
                                                    _ = v33
                                                }
                                                _ = v31
                                            }
                                        } else {
                                            v34 := GenS5{F0: int(v2), F1: 65535, F2: "hello"}
                                            _ = v34
                                        }
                                    }
                                default:
                                    var v35 float64 = float64((1.5 + ((1e10 * 3.14159) - (1.5 - 0.1))))
                                    _ = v35
                                }
                                fmt.Println(v2, v18)
                            }
                        case "a":
                            v4["hello"] = 42
                        default:
                            for fi36 := 0; fi36 < 8; fi36++ {
                                v18 += (v2 + (^(v2)))
                                v37 := GenS3{F0: (2147483647 <= fi36), F1: genFn0(true, -1, -162)}
                                v38 := GenS5{F0: fi36, F1: 443, F2: "a"}
                                _ = fi36
                                _ = v37
                                _ = v38
                            }
                        }
                    }
                default:
                    switch 9223372036854775807 {
                    case 32768:
                        for fi39 := 0; fi39 < 11; fi39++ {
                            var v40 int8 = 25
                            _ = len(v4)
                            var v41 uint = (v2 * v18)
                            _ = fi39
                            _ = v40
                            _ = v41
                        }
                    default:
                        v42 := v4[(("" + "") + ("foo bar" + "hello"))]
                        _ = v42
                    }
                }
                for fi43 := 0; fi43 < 9; fi43++ {
                    var v44 int = fi43
                    _ = fi43
                    _ = v44
                }
            }
        } else {
            v18 = (v2 % ((v18) | 1))
            if (!true) {
                if (!false) {
                    v18 = v18
                    v18 += (v2 >> (((v18 % (((v2 & (v2 * v2))) | 1))) & 63))
                } else {
                    for fi45 := 0; fi45 < 8; fi45++ {
                        v2 = (v18 / ((255) | 1))
                        var v46 float64 = -2.25
                        v47 := []int{32768, 42, 2147483647}
                        _ = fi45
                        _ = v46
                        _ = v47
                    }
                    v48 := GenS5{F0: 127, F1: 42, F2: ""}
                    _ = v48
                }
                v18 += v18
            } else {
                for fi49 := 0; fi49 < 15; fi49++ {
                    for fi50 := 0; fi50 < 6; fi50++ {
                        v51 := GenS4{F0: false}
                        v52 := GenS6{F0: 1, F1: -714}
                        v53 := []int{32768, (fi49 & (-(fi49))), fi49}
                        _ = fi50
                        _ = v51
                        _ = v52
                        _ = v53
                    }
                    _ = fi49
                }
                var v54 bool = true
                _ = v54
            }
        }
    default:
        v2 = (v2 / ((v2) | 1))
    }
    fmt.Println(v18, v2)
    for fi55 := 0; fi55 < 17; fi55++ {
        for fi56 := 0; fi56 < 17; fi56++ {
            fmt.Println(v2, v2, fi56)
            v57 := GenS3{F0: ("\n\t" <= "hello"), F1: 2}
            _ = fi56
            _ = v57
        }
        _ = fi55
    }
    var v58 uint = (^(v18))
    switch 993 {
    case -817:
        v59 := []int{1, 32768}
        _ = v59
    case 128:
        if (-979 <= 1000) {
            fmt.Println(v18, v18)
            var v60 uint = (^(v58))
            _ = v60
        } else {
            var v61 bool = true
            if v61 {
                v2 += (^(v18))
                var v62 string = "hello"
                for fi63 := 0; fi63 < 8; fi63++ {
                    v64 := []int{128}
                    v61 = (false && (42 <= 0))
                    _ = fi63
                    _ = v64
                }
                _ = v62
            }
            v65 := []int{-652}
            _ = v61
            _ = v65
        }
    case 4294967295:
        v66 := GenS7{F0: 2, F1: 2}
        _ = v66
    default:
        var v67 float64 = ((1.5 / ((float64(0.0) / 1.0) - -2.25)) * 0.1)
        _ = v67
    }
    v68 := v4["hello"]
    var v69 uint64 = 0
    var v70 int = -702
    v69 -= v69
    var v71 uint = (v2 + 1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v58:", v58)
    fmt.Println("v68:", v68)
    fmt.Println("v69:", v69)
    fmt.Println("v70:", v70)
    fmt.Println("v71:", v71)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v58
    _ = v68
    _ = v69
    _ = v70
    _ = v71
}
