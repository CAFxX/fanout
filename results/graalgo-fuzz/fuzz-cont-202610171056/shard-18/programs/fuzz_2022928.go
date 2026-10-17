package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint; F1 bool; F2 uint32 }
type GenS5 struct { F0 uint64; F1 int32; F2 uint }

func main() {
    var v1 uint = 4294967295
    v2 := []string{}
    v3 := make(map[int]int)
    var v4 int64 = 42
    var v5 int8 = -128
    for fi6 := 0; fi6 < 3; fi6++ {
        v1 += (v1 - v1)
        v7 := GenS1{F0: 255}
        _ = fi6
        _ = v7
    }
    for fi8 := 0; fi8 < 9; fi8++ {
        v5 += (v5 + v5)
        var v9 int16 = -1
        _ = fi8
        _ = v9
    }
    v5 = v5
    switch (("a" + "hello") + "a") {
    case "":
        if true {
            v5 = v5
        } else {
            v10 := GenS2{F0: (-(v4))}
            _ = v10
        }
    default:
        _ = len(v2)
    }
    for fi11 := 0; fi11 < 12; fi11++ {
        for fi12 := 0; fi12 < 10; fi12++ {
            v1 = v1
            if (0 >= 718) {
                var v13 uint16 = 255
                _ = v13
            } else {
                switch "hello" {
                case "\n\t":
                    var v14 int64 = 0
                    v14 = (v4 % (((v4 & (^(v4)))) | 1))
                    fmt.Println(fi11, v5)
                    _ = v14
                case "a":
                    if true {
                        var v15 uint8 = 1
                        switch "foo bar" {
                        case "hello":
                            v16 := v3[fi12]
                            _ = v16
                        case "foo bar":
                            if (!(!((v1 >> ((2) & 63)) >= (v1 * (v1 << ((v1) & 63)))))) {
                                _ = len(v2)
                            }
                        }
                        if true {
                            v17 := GenS3{F0: v1}
                            switch ("hello" + ("hello" + "a")) {
                            case "":
                                var v18 int16 = 865
                                var v19 int16 = (v18 ^ v18)
                                if (0 <= 4294967295) {
                                    if true {
                                        var v20 uint64 = 1
                                        delete(v3, (fi11 ^ int(v4)))
                                        _ = v20
                                    } else {
                                        fmt.Println(fi11)
                                    }
                                    v5 -= (^(v5))
                                }
                                _ = v18
                                _ = v19
                            default:
                                if (-1 < -1) {
                                    if false {
                                        var v21 uint8 = (v15 & v15)
                                        _ = v21
                                    }
                                    var v22 int = (fi11 * fi12)
                                    _ = v22
                                }
                            }
                            _ = v17
                        } else {
                            if false {
                                var v23 uint64 = 18446744073709551615
                                var v24 int64 = (v4 % (((v4 * (v4 << ((v1) & 63)))) | 1))
                                _ = v23
                                _ = v24
                            } else {
                                fmt.Println(v1, fi12, v15)
                                v15 = (^(v15))
                            }
                        }
                        _ = v15
                    }
                case "":
                    var v25 uint = v1
                    _ = v25
                }
            }
            v26 := GenS4{F0: 509, F1: (true || (!(0 != 1))), F2: 0}
            v2 = append(v2, ("foo bar" + "a"))
            _ = fi12
            _ = v26
        }
        var v27 float32 = 1e10
        _ = fi11
        _ = v27
    }
    switch 441 {
    case 832:
        v2 = append(v2, "\n\t")
    case -271:
        v1 = v1
    case -605:
        v28 := v3[-2147483648]
        _ = v28
    default:
        var v29 bool = false
        _ = v29
    }
    var v30 uint16 = 32767
    for fi31 := 0; fi31 < 6; fi31++ {
        if (-383 <= -32768) {
            if (!(v5 != v5)) {
                v1 = v1
                v30 = (v30 / ((135) | 1))
            }
        } else {
            if (false || false) {
                switch (fi31 - 32768) {
                case 721:
                    if true {
                        v5 = v5
                        if ((fi31 >> (((v1 / (((v1 % ((42) | 1))) | 1))) & 63)) < 9223372036854775807) {
                            for fi32 := 0; fi32 < 20; fi32++ {
                                var v33 uint16 = v30
                                v34 := GenS1{F0: 127}
                                _ = fi32
                                _ = v33
                                _ = v34
                            }
                            v35 := GenS1{F0: uint8(v30)}
                            _ = v35
                        }
                    } else {
                        if true {
                            v36 := GenS4{F0: 2, F1: false, F2: 2}
                            _ = v36
                        } else {
                            v37 := GenS2{F0: 1}
                            _ = v37
                        }
                        v38 := GenS4{F0: v1, F1: (1 <= 906), F2: 980}
                        _ = v38
                    }
                case -179:
                    v39 := GenS5{F0: 24, F1: 2147483647, F2: uint(v30)}
                    _ = v39
                default:
                    var v40 string = "hello"
                    _ = v40
                }
            }
        }
        delete(v3, -1)
        _ = fi31
    }
    var v41 bool = (!(-9223372036854775808 == 374))
    var v42 float32 = -2.25
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v41
    _ = v42
    _ = v5
}
