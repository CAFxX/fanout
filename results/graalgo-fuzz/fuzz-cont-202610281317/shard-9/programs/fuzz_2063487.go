package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 int }
type GenS2 struct { F0 bool; F1 int16; F2 uint }
type GenS3 struct { F0 uint16; F1 int64 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int; F1 uint8; F2 int16 }

func main() {
    var v1 bool = false
    v2 := make(map[int]bool)
    v3 := make(map[int]uint32)
    var v4 float64 = 0.0
    switch 127 {
    case 986:
        for fi5 := 0; fi5 < 17; fi5++ {
            v1 = true
            if (!true) {
                v4 = (float64(-2.25) + 0.0)
                var v6 uint8 = 42
                _ = v6
            }
            for fi7 := 0; fi7 < 3; fi7++ {
                v1 = true
                v8 := GenS1{F0: 484, F1: 698, F2: fi5}
                _ = fi7
                _ = v8
            }
            _ = fi5
        }
    case -937:
        fmt.Println(v1)
    case -9223372036854775808:
        switch ("foo bar" + ("\n\t" + "a")) {
        case "hello":
            var v9 uint64 = 341
            _ = v9
        case "a":
            delete(v3, 4294967295)
        }
    default:
        if false {
            v4 -= (v4 / v4)
            for fi10 := 0; fi10 < 6; fi10++ {
                if (65535 == 255) {
                    v4 = 1.5
                    if ((42 <= 1) || false) {
                        fmt.Println(fi10)
                    }
                } else {
                    var v11 uint16 = 2
                    var v12 int8 = 113
                    fmt.Println(v4, fi10, v11)
                    _ = v11
                    _ = v12
                }
                if (int64(fi10) >= 9223372036854775807) {
                    for fi13 := 0; fi13 < 1; fi13++ {
                        v1 = v1
                        v14 := GenS2{F0: ((fi10 | fi10) >= 2147483647), F1: 32767, F2: 886}
                        v15 := GenS2{F0: v1, F1: 32767, F2: 0}
                        _ = fi13
                        _ = v14
                        _ = v15
                    }
                } else {
                    v16 := GenS1{F0: 1, F1: 32767, F2: 127}
                    if ("foo bar" > (("\n\t" + "\n\t") + "")) {
                        fmt.Println(v4)
                        v4 -= (v4 / v4)
                        var v17 uint64 = 638
                        _ = v17
                    } else {
                        for fi18 := 0; fi18 < 3; fi18++ {
                            var v19 int16 = 882
                            v4 -= v4
                            var v20 float32 = -2.25
                            _ = fi18
                            _ = v19
                            _ = v20
                        }
                    }
                    _ = v16
                }
                _ = fi10
            }
        } else {
            for fi21 := 0; fi21 < 9; fi21++ {
                fmt.Println(v4, v4, v4)
                v2[(fi21 | 32768)] = false
                _ = fi21
            }
            v22 := []int{-477}
            _ = v22
        }
    }
    var v23 uint = 0
    fmt.Println(v1, v1, v23)
    if (42 < 944) {
        fmt.Println(v4, v23)
    } else {
        for fi24 := 0; fi24 < 1; fi24++ {
            for fi25 := 0; fi25 < 6; fi25++ {
                switch "foo bar" {
                case "\n\t":
                    var v26 uint16 = 2
                    _ = v26
                case "hello":
                    v27 := []int{fi24, fi24, fi24, (fi25 % ((fi25) | 1))}
                    _ = v27
                }
                v1 = false
                _ = fi25
            }
            for fi28 := 0; fi28 < 15; fi28++ {
                var v29 int16 = -32768
                v30 := GenS3{F0: 0, F1: 0}
                _ = fi28
                _ = v29
                _ = v30
            }
            _ = fi24
        }
        for fi31 := 0; fi31 < 12; fi31++ {
            if v1 {
                switch (fi31 / ((0) | 1)) {
                case 4294967295:
                    if ((v1 && (v1 || v1)) || ("\n\t" >= (("\n\t" + "") + (("foo bar" + "hello") + "foo bar")))) {
                        for fi32 := 0; fi32 < 4; fi32++ {
                            if true {
                                v1 = (uint32(v23) <= 42)
                                fmt.Println(fi31)
                                _ = len(v3)
                            }
                            v33 := GenS4{F0: (^(fi32))}
                            _ = fi32
                            _ = v33
                        }
                        v34 := v2[(fi31 % ((9223372036854775807) | 1))]
                        v1 = (2147483648 >= 4294967295)
                        _ = v34
                    }
                case 32768:
                    v35 := GenS2{F0: (((((1.5 / 1.5) / (1e10 + ((-2.25 * 3.14159) * 3.14159))) / 1.0) < (-2.25 * -2.25)) && v1), F1: 1000, F2: 42}
                    _ = v35
                case 0:
                    v1 = v1
                default:
                    for fi36 := 0; fi36 < 8; fi36++ {
                        v4 = -2.25
                        v1 = (!(-90 == 70))
                        _ = fi36
                    }
                }
                if (v23 <= v23) {
                    var v37 uint8 = 0
                    _ = v37
                } else {
                    v4 -= ((0.0 - ((v4 * (v4 + v4)) * (float64((-2.25 / (3.14159 * (0.1 * -2.25)))) + (float64(1.5) + v4)))) + (v4 + (v4 - v4)))
                }
            }
            _ = fi31
        }
        for fi38 := 0; fi38 < 10; fi38++ {
            for fi39 := 0; fi39 < 13; fi39++ {
                v40 := GenS2{F0: false, F1: -383, F2: (v23 % ((0) | 1))}
                var v41 uint16 = 32767
                _ = fi39
                _ = v40
                _ = v41
            }
            _ = fi38
        }
    }
    for fi42 := 0; fi42 < 15; fi42++ {
        if true {
            delete(v2, 128)
        }
        v43 := []int{}
        _ = fi42
        _ = v43
    }
    v44 := GenS5{F0: 0, F1: 0, F2: 1}
    if (3.14159 >= float64(1e10)) {
        v45 := GenS1{F0: 0, F1: 0, F2: -725}
        if (0 == -34) {
            for fi46 := 0; fi46 < 3; fi46++ {
                var v47 uint = v23
                _ = fi46
                _ = v47
            }
        } else {
            fmt.Println(v1, v23, v1)
            for fi48 := 0; fi48 < 5; fi48++ {
                v23 = v23
                var v49 int8 = 42
                var v50 float32 = float32((v4 / v4))
                _ = fi48
                _ = v49
                _ = v50
            }
        }
        var v51 string = "hello"
        _ = v45
        _ = v51
    } else {
        var v52 string = "foo bar"
        if (2 <= 0) {
            v53 := []int{258, -762, 86, -9223372036854775808}
            for fi54 := 0; fi54 < 4; fi54++ {
                var v55 float32 = 1e10
                _ = fi54
                _ = v55
            }
            var v56 int = -1
            _ = v53
            _ = v56
        }
        _ = v52
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v44
}
