package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 int32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 bool; F1 int8; F2 int8 }
func genFn0(p0 int, p1 bool, p2 uint) uint {
    v4 := []int{(p0 + 32768)}
    if true {
        var v5 int8 = 42
        v6 := GenS1{F0: uint16(p2), F1: 255, F2: -2147483648}
        v4 = append(v4, (p0 / ((832) | 1)))
        _ = v5
        _ = v6
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return (^(p2))
}

func main() {
    var v1 uint = 42
    v2 := []int{-2147483648, 1}
    var v3 float64 = (-2.25 - -2.25)
    v1 = (^(v1))
    switch -657 {
    case 336:
        if false {
            fmt.Println(v3)
        }
    case 127:
        var v7 uint16 = 29
        _ = v7
    default:
        for fi8 := 0; fi8 < 19; fi8++ {
            var v9 uint32 = 2147483648
            var v10 int = fi8
            _ = fi8
            _ = v10
            _ = v9
        }
    }
    var v11 int32 = -203
    switch -2147483648 {
    case 407:
        var v12 int8 = 42
        _ = v12
    case 1:
        switch -843 {
        case -2147483648:
            var v13 uint32 = 2
            _ = v13
        case 128:
            v3 += v3
        case 0:
            v11 = v11
        }
    }
    v2 = append(v2, -484)
    v1 = 795
    for fi14 := 0; fi14 < 15; fi14++ {
        v3 += v3
        if true {
            var v15 float64 = v3
            _ = v15
        } else {
            v16 := GenS2{F0: 9223372036854775807}
            if (!false) {
                v2 = append(v2, 2147483647)
                if ((42 != v11) && (524 >= 42)) {
                    v3 = v3
                    var v17 uint = 155
                    _ = v17
                } else {
                    _ = len(v2)
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
                if true {
                    fmt.Println(v1)
                    if true {
                        fmt.Println(v3, v3)
                        if true {
                            fmt.Println(v1, v1, fi14)
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        } else {
                            for fi18 := 0; fi18 < 9; fi18++ {
                                v19 := GenS2{F0: -1}
                                _ = fi18
                                _ = v19
                            }
                            v11 -= v11
                        }
                    }
                    var v20 float32 = -2.25
                    _ = v20
                } else {
                    v21 := GenS3{F0: false}
                    v1 = v1
                    _ = v21
                }
            } else {
                for fi22 := 0; fi22 < 13; fi22++ {
                    if (false && true) {
                        _ = len(v2)
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        if false {
                            if (fi14 > (fi14 / ((fi22) | 1))) {
                                if false {
                                    if len(v2) > 0 {
                                        v2[0] = -1
                                    }
                                } else {
                                    var v23 float64 = (v3 + 0.0)
                                    switch ((("\n\t" + (("" + ("a" + ("a" + "a"))) + (((("" + "foo bar") + (("a" + ("hello" + "\n\t")) + "a")) + "") + "hello"))) + "\n\t") + "\n\t") {
                                    case "a":
                                        if ((0 == 1) || true) {
                                            var v24 uint8 = 255
                                            _ = v24
                                        }
                                    case "":
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                    case "foo bar":
                                        var v25 bool = ("foo bar" != (("" + "foo bar") + ""))
                                        var v26 int8 = 0
                                        v1 += v1
                                        v11 -= v11
                                        _ = v25
                                        _ = v26
                                    }
                                    _ = v23
                                }
                                switch "\n\t" {
                                case "foo bar":
                                    v27 := GenS1{F0: 1, F1: 32767, F2: v11}
                                    _ = v27
                                default:
                                    switch "\n\t" {
                                    case "\n\t":
                                        v28 := GenS4{F0: 34}
                                        _ = v28
                                    case "hello":
                                        v11 -= (v11 - (^(v11)))
                                    }
                                }
                                v1 += (v1 / ((v1) | 1))
                            } else {
                                v11 -= (^(v11))
                            }
                            v29 := GenS5{F0: 2}
                            v3 += ((v3 + (3.14159 / v3)) / v3)
                            _ = v29
                        } else {
                            var v30 int32 = -798
                            v3 = v3
                            _ = v30
                        }
                        var v31 string = "hello"
                        v1 += v1
                        _ = v31
                    }
                    var v32 uint = (v1 / ((v1) | 1))
                    v33 := GenS2{F0: -1}
                    _ = fi22
                    _ = v32
                    _ = v33
                }
            }
            fmt.Println(v11)
            _ = v16
        }
        _ = fi14
    }
    fmt.Println(v11, v11, v3)
    for fi34 := 0; fi34 < 20; fi34++ {
        v3 = 1e10
        _ = fi34
    }
    if ((v1 * 18446744073709551615) != genFn0(32767, (!true), v1)) {
        switch (("hello" + ("\n\t" + "")) + ("a" + "\n\t")) {
        case "\n\t":
            v1 = v1
        case "a":
            var v35 int32 = 0
            _ = v35
        }
        var v36 uint64 = 128
        _ = v36
    } else {
        v37 := GenS6{F0: (0.1 == 1.5), F1: 127, F2: -1}
        _ = v37
    }
    fmt.Println(v11, v11)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
}
