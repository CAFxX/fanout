package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int64; F1 int64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 int64 }

func main() {
    var v1 uint32 = 2147483647
    var v2 uint8 = 2
    v3 := make(map[uint32]string)
    var v4 int32 = -631
    var v5 int64 = 9223372036854775807
    fmt.Println(v4, v4, v1)
    v4 = (v4 >> ((65535) & 63))
    var v6 bool = (uint16(v2) == 2)
    v2 = v2
    var v7 int32 = (v4 - v4)
    fmt.Println(v2)
    var v8 uint32 = v1
    v9 := GenS1{F0: v7}
    delete(v3, 851)
    for fi10 := 0; fi10 < 9; fi10++ {
        if v6 {
            switch (("hello" + "\n\t") + ("" + "\n\t")) {
            case "":
                if (uint(v1) == 611) {
                    switch 1 {
                    case 32767:
                        v11 := GenS1{F0: (v4 / ((42) | 1))}
                        _ = v11
                    case -2147483648:
                        var v12 float64 = 0.0
                        _ = v12
                    default:
                        var v13 uint32 = (^(v8))
                        _ = v13
                    }
                    v14 := GenS2{F0: (v5 << ((38) & 63)), F1: (v5 << ((598) & 63))}
                    _ = v14
                } else {
                    if true {
                        v15 := []int{-524, (fi10 ^ -1)}
                        _ = v15
                    }
                    for fi16 := 0; fi16 < 17; fi16++ {
                        var v17 uint8 = 112
                        _ = fi16
                        _ = v17
                    }
                    switch "hello" {
                    case "":
                        if v6 {
                            switch (-(fi10)) {
                            case 0:
                                for fi18 := 0; fi18 < 13; fi18++ {
                                    v19 := GenS1{F0: -2147483648}
                                    v20 := GenS3{F0: (v2 >> ((18446744073709551615) & 63))}
                                    _ = fi18
                                    _ = v19
                                    _ = v20
                                }
                            case 1:
                                v21 := []int{(fi10 + (fi10 & (fi10 ^ fi10))), fi10, (fi10 * 128), (fi10 | fi10)}
                                _ = v21
                            case 128:
                                v22 := GenS1{F0: -1}
                                _ = v22
                            default:
                                for fi23 := 0; fi23 < 1; fi23++ {
                                    if (0 < 2) {
                                        v6 = v6
                                    }
                                    v24 := GenS2{F0: (v5 % ((42) | 1)), F1: (v5 / (((^(v5))) | 1))}
                                    _ = fi23
                                    _ = v24
                                }
                            }
                        } else {
                            v25 := []int{fi10}
                            if v6 {
                                if (1 != 18446744073709551615) {
                                    if (1e10 < 0.0) {
                                        var v26 bool = (828 < 65535)
                                        if len(v25) > 0 {
                                            v25[0] = 6
                                        }
                                        v27 := GenS2{F0: (v5 % ((v5) | 1)), F1: v5}
                                        _ = v26
                                        _ = v27
                                    } else {
                                        fmt.Println(v7, v6, v4)
                                        v2 -= v2
                                        var v28 bool = ((("foo bar" + "a") + "a") != "foo bar")
                                        _ = v28
                                    }
                                } else {
                                    v29 := GenS4{F0: 2}
                                    var v30 float64 = (0.0 + (1.5 + ((1.5 - 1e10) * float64(0.1))))
                                    var v31 int16 = 707
                                    _ = v29
                                    _ = v30
                                    _ = v31
                                }
                                if (0 < 6) {
                                    v3[(^(v1))] = "hello"
                                    v3[(v8 / ((v8) | 1))] = "\n\t"
                                } else {
                                    switch (fi10 | fi10) {
                                    case 1:
                                        v32 := GenS5{F0: 413}
                                        _ = v32
                                    default:
                                        switch ("a" + "foo bar") {
                                        case "\n\t":
                                            fmt.Println(fi10, v6)
                                        case "":
                                            v7 = 1
                                        default:
                                            var v33 float64 = (-2.25 + -2.25)
                                            _ = v33
                                        }
                                    }
                                    v34 := GenS2{F0: (-(v5)), F1: (v5 & (v5 & 0))}
                                    for fi35 := 0; fi35 < 19; fi35++ {
                                        switch (((((("hello" + "a") + "foo bar") + "") + "\n\t") + "hello") + "hello") {
                                        case "":
                                            v6 = ((!(!(90 > 42))) || (!(18446744073709551615 != 417)))
                                        default:
                                            var v36 uint = 838
                                            _ = v36
                                        }
                                        var v37 int16 = 32767
                                        fmt.Println(v6, v2)
                                        _ = fi35
                                        _ = v37
                                    }
                                    _ = v34
                                }
                            } else {
                                var v38 uint32 = (v8 - 0)
                                _ = v38
                            }
                            _ = v25
                        }
                    case "hello":
                        v7 = (v4 << ((uint(v8)) & 63))
                    default:
                        v39 := GenS6{F0: 1}
                        _ = v39
                    }
                }
            default:
                v40 := GenS7{F0: (v5 - v5)}
                _ = v40
            }
            v41 := v3[467]
            for fi42 := 0; fi42 < 6; fi42++ {
                v43 := GenS6{F0: 100}
                if v6 {
                    v3[2] = (v41 + v41)
                    var v44 int8 = 0
                    v3[(^(v8))] = v41
                    v45 := GenS6{F0: v44}
                    _ = v44
                    _ = v45
                }
                v41 = (((v41 + "a") + (v41 + v41)) + ("" + v41))
                _ = fi42
                _ = v43
            }
            _ = v41
        }
        _ = fi10
    }
    v46 := v3[(^(v1))]
    for fi47 := 0; fi47 < 4; fi47++ {
        var v48 int64 = v5
        fmt.Println(v6, v46)
        _ = fi47
        _ = v48
    }
    var v49 bool = (!v6)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v49
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
