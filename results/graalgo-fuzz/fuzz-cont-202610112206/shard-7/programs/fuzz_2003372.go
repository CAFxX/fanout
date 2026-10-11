package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 string; F1 int64 }
type GenS3 struct { F0 uint8; F1 int32; F2 uint16 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint8; F1 bool }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 int; F1 int32 }
type GenS8 struct { F0 int8; F1 uint8 }

func main() {
    v1 := make(map[uint32]string)
    v2 := []int{128, -50}
    var v3 int64 = 1
    v4 := []int{127, 128}
    var v5 int64 = (v3 - (v3 + v3))
    v6 := make(map[uint32]uint32)
    if false {
        if (3.14159 == float64(-2.25)) {
            var v7 float64 = (3.14159 - -2.25)
            _ = v7
        }
    }
    v5 = v3
    v8 := v6[4294967295]
    v3 -= (v3 + v3)
    v9 := v6[2147483648]
    for fi10 := 0; fi10 < 19; fi10++ {
        v11 := GenS1{F0: 1}
        v9 -= v9
        if (42 > (v5 >> ((42) & 63))) {
            if true {
                var v12 int16 = 1
                _ = v12
            } else {
                for fi13 := 0; fi13 < 14; fi13++ {
                    fmt.Println(v5, fi13)
                    fmt.Println(v3, v8)
                    _ = fi13
                }
                v4 = append(v4, 42)
            }
            v5 -= v3
            var v14 int8 = -1
            _ = v14
        } else {
            var v15 int = fi10
            v9 -= 2
            var v16 int = -749
            _ = v15
            _ = v16
        }
        _ = fi10
        _ = v11
    }
    _ = len(v6)
    if false {
        v9 = v9
        if false {
            v8 += 950
            switch -883 {
            case 0:
                for fi17 := 0; fi17 < 9; fi17++ {
                    if false {
                        fmt.Println(v5, v5)
                        v3 -= -1
                    } else {
                        var v18 float32 = float32(3.14159)
                        for fi19 := 0; fi19 < 15; fi19++ {
                            if false {
                                v3 += v3
                            } else {
                                if (!(0.1 != ((0.1 / 1e10) - 1e10))) {
                                    fmt.Println(fi19, fi19, fi17)
                                    var v20 int = fi17
                                    var v21 string = "a"
                                    _ = v20
                                    _ = v21
                                } else {
                                    v3 = 283
                                }
                                v22 := GenS1{F0: 1000}
                                _ = v22
                            }
                            v5 += -9223372036854775808
                            var v23 int8 = 37
                            _ = fi19
                            _ = v23
                        }
                        var v24 uint16 = 32767
                        _ = v18
                        _ = v24
                    }
                    _ = fi17
                }
            default:
                v5 += (v3 + v5)
            }
            var v25 uint16 = 65535
            _ = v25
        } else {
            v3 = -859
            v3 = v5
            v26 := GenS2{F0: "foo bar", F1: v5}
            _ = v26
        }
        fmt.Println(v9, v8, v3)
    } else {
        v9 = (^(v8))
        v27 := GenS1{F0: 1000}
        var v28 bool = (-1 != 128)
        _ = v27
        _ = v28
    }
    if false {
        for fi29 := 0; fi29 < 5; fi29++ {
            if len(v2) > 0 {
                v2[0] = (fi29 % (((-(fi29))) | 1))
            }
            _ = fi29
        }
        if (1e10 < 0.1) {
            var v30 int64 = v3
            fmt.Println(v3, v8)
            _ = v30
        } else {
            fmt.Println(v8, v3)
        }
        for fi31 := 0; fi31 < 10; fi31++ {
            if (0.1 == (float64(((0.0 - (float32(float64(1e10)) * 1e10)) / ((0.0 + 0.0) - 1.5))) + float64(1.5))) {
                var v32 uint16 = 883
                _ = v32
            } else {
                v3 = (v5 % ((1) | 1))
                var v33 uint64 = 18446744073709551615
                _ = v33
            }
            v34 := GenS3{F0: 127, F1: -1, F2: 65535}
            _ = fi31
            _ = v34
        }
    }
    var v35 int = 4294967295
    v36 := GenS3{F0: 255, F1: -925, F2: 32767}
    switch "a" {
    case "hello":
        if false {
            if true {
                v37 := GenS4{F0: v35}
                _ = v37
            } else {
                if false {
                    var v38 int = v35
                    _ = v38
                } else {
                    for fi39 := 0; fi39 < 3; fi39++ {
                        v35 = fi39
                        _ = len(v4)
                        _ = fi39
                    }
                    var v40 uint8 = 42
                    _ = v40
                }
                for fi41 := 0; fi41 < 12; fi41++ {
                    for fi42 := 0; fi42 < 20; fi42++ {
                        v35 -= 9223372036854775807
                        _ = fi42
                    }
                    _ = fi41
                }
            }
        } else {
            v35 = v35
        }
    case "a":
        fmt.Println(v3, v3)
    }
    v2 = append(v2, (-(v35)))
    fmt.Println(v8)
    v9 = v8
    if true {
        if ((false && false) && true) {
            if (false && true) {
                if false {
                    v43 := GenS5{F0: 100, F1: true}
                    switch ("foo bar" + "hello") {
                    case "foo bar":
                        var v44 int64 = v5
                        _ = v44
                    case "\n\t":
                        v3 += (^(v3))
                    case "a":
                        if true {
                            for fi45 := 0; fi45 < 13; fi45++ {
                                v35 -= fi45
                                if true {
                                    fmt.Println(v5)
                                }
                                fmt.Println(v9, fi45, v35)
                                _ = fi45
                            }
                        }
                    default:
                        v3 = 753
                    }
                    if false {
                        switch (v35 / (((v35 & (^(v35)))) | 1)) {
                        case -9223372036854775808:
                            var v46 uint32 = v9
                            _ = v46
                        case -664:
                            var v47 float64 = (((-2.25 + float64(float32(0.1))) - 1e10) + 0.1)
                            _ = v47
                        }
                    }
                    _ = v43
                }
            }
            v48 := GenS6{F0: 0}
            var v49 uint8 = 2
            _ = v48
            _ = v49
        } else {
            for fi50 := 0; fi50 < 11; fi50++ {
                v6[792] = v9
                _ = fi50
            }
            v35 += v35
            if true {
                for fi51 := 0; fi51 < 2; fi51++ {
                    delete(v1, v9)
                    _ = fi51
                }
                v35 -= v35
                v35 -= (v35 * (v35 & v35))
            } else {
                v52 := GenS7{F0: v35, F1: 0}
                for fi53 := 0; fi53 < 15; fi53++ {
                    v54 := GenS5{F0: 255, F1: ("" > "a")}
                    v55 := GenS8{F0: 42, F1: 91}
                    _ = fi53
                    _ = v54
                    _ = v55
                }
                _ = v52
            }
        }
        var v56 uint8 = 0
        var v57 int64 = -742
        _ = v56
        _ = v57
    }
    v9 -= 4294967295
    v9 = (v8 / ((v9) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
