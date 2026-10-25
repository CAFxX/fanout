package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 int8; F1 uint8; F2 uint64 }
type GenS3 struct { F0 uint; F1 int16; F2 uint }
type GenS4 struct { F0 bool; F1 int8 }
type GenS5 struct { F0 uint; F1 int8 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 int; F1 int16; F2 int16 }
type GenS8 struct { F0 string }
func genFn0(p0 int8) uint8 {
    v4 := []int{4294967295, 32768}
    var v5 uint32 = 65535
    v5 = (v5 & v5)
    var v6 uint = 255
    var v7 int32 = 1
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return 128
}

func main() {
    var v1 float32 = 0.0
    v2 := []int{}
    var v3 float32 = (v1 - float32(0.1))
    var v8 uint64 = 0
    if false {
        v8 = v8
    } else {
        if false {
            var v9 float32 = v3
            _ = v9
        }
    }
    _ = len(v2)
    switch 1 {
    case 4294967295:
        var v10 int64 = -198
        _ = v10
    case -778:
        if true {
            for fi11 := 0; fi11 < 5; fi11++ {
                v3 = v1
                _ = fi11
            }
        }
    case -670:
        if false {
            var v12 uint64 = (v8 + (v8 * v8))
            fmt.Println(v3, v12)
            _ = v12
        } else {
            v13 := GenS1{F0: 777, F1: 2147483647}
            switch 128 {
            case -1:
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            case 127:
                v1 = 1.5
            }
            var v14 uint64 = (v8 % (((v8 / (((v8 / ((v8) | 1))) | 1))) | 1))
            _ = v13
            _ = v14
        }
    default:
        for fi15 := 0; fi15 < 7; fi15++ {
            for fi16 := 0; fi16 < 19; fi16++ {
                v17 := GenS1{F0: 205, F1: 132}
                switch (fi16 << ((65535) & 63)) {
                case 501:
                    fmt.Println(v8)
                case -459:
                    if ((v8 / ((v8) | 1)) >= (v8 / (((^(v8))) | 1))) {
                        var v18 uint64 = 18446744073709551615
                        if false {
                            if true {
                                if true {
                                    v19 := GenS1{F0: 42, F1: 42}
                                    if true {
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        v18 += (^(v8))
                                    }
                                    _ = v19
                                }
                                switch "foo bar" {
                                case "hello":
                                    v20 := GenS1{F0: 725, F1: 622}
                                    _ = v20
                                default:
                                    v21 := GenS1{F0: 874, F1: 2147483647}
                                    _ = v21
                                }
                            } else {
                                fmt.Println(fi15, v18)
                            }
                            v22 := GenS2{F0: 79, F1: 1, F2: v8}
                            v23 := GenS3{F0: 18446744073709551615, F1: 42, F2: 65535}
                            v18 = 18446744073709551615
                            if (!false) {
                                var v24 int = fi15
                                _ = v24
                            }
                            _ = v22
                            _ = v23
                        }
                        var v25 int32 = 0
                        var v26 uint16 = 2
                        var v27 int32 = (-(v25))
                        var v28 int32 = 1
                        _ = v18
                        _ = v25
                        _ = v26
                        _ = v27
                        _ = v28
                    }
                case -576:
                    v29 := GenS3{F0: 133, F1: 0, F2: 1}
                    _ = v29
                default:
                    switch 42 {
                    case 32767:
                        var v30 string = ("\n\t" + "hello")
                        _ = v30
                    case -780:
                        v3 = ((v3 - float32(1.5)) / 0.1)
                    }
                }
                v1 += (v1 * v3)
                _ = fi16
                _ = v17
            }
            if len(v2) > 0 {
                v2[0] = (fi15 / (((fi15 * -7)) | 1))
            }
            fmt.Println(v1)
            _ = fi15
        }
    }
    fmt.Println(v8, v8, v3)
    v2 = append(v2, 138)
    var v31 int64 = -327
    var v32 uint8 = 128
    var v33 int8 = 42
    var v34 string = "foo bar"
    var v35 int = -885
    if false {
        var v36 uint64 = (v8 | v8)
        v37 := GenS4{F0: true, F1: -1}
        if false {
            if (((((v1 - (0.0 / v3)) * v1) != v1) || (400 == 42)) && (v32 < v32)) {
                var v38 int64 = (v31 & -733)
                if (("" + v34) == (v34 + (((v34 + v34) + "foo bar") + v34))) {
                    for fi39 := 0; fi39 < 18; fi39++ {
                        var v40 uint = 65535
                        v41 := GenS5{F0: (v40 ^ (v40 | (v40 & v40))), F1: (v33 << (((v40 | 42)) & 63))}
                        _ = fi39
                        _ = v40
                        _ = v41
                    }
                } else {
                    if len(v2) > 0 {
                        v2[0] = 32767
                    }
                    if ((float64((((1.5 - (v3 + v3)) + v3) / 1.5)) * 1e10) != 1e10) {
                        v42 := GenS6{F0: 42}
                        _ = v42
                    } else {
                        if true {
                            var v43 uint64 = uint64(v32)
                            var v44 float32 = v3
                            var v45 int = (v35 << ((1) & 63))
                            _ = v43
                            _ = v44
                            _ = v45
                        }
                        var v46 int64 = 42
                        _ = v46
                    }
                }
                if ((("foo bar" + ("\n\t" + v34)) != v34) && false) {
                    var v47 uint16 = uint16(v8)
                    _ = v47
                } else {
                    v31 -= (-(v38))
                    var v48 float64 = 1.5
                    _ = v48
                }
                _ = v38
            }
            if ((-102 < v31) || (false || (uint32(v8) > 2147483647))) {
                v2 = append(v2, 166)
            }
        } else {
            v35 = (v35 >> ((4294967295) & 63))
            v3 -= float32(1e10)
            v8 = 4294967295
        }
        _ = v36
        _ = v37
    } else {
        v3 = (float32(0.0) / ((0.0 / ((float32(((1.5 - 1e10) - 1e10)) - ((((0.0 * (3.14159 - v1)) - -2.25) - (1.5 + v1)) + v1)) / v1)) + v3))
        for fi49 := 0; fi49 < 4; fi49++ {
            _ = len(v2)
            switch v35 {
            case -645:
                var v50 int32 = 42
                _ = v50
            case 32767:
                var v51 bool = ("foo bar" <= v34)
                _ = v51
            case 127:
                fmt.Println(v3, v8)
            }
            for fi52 := 0; fi52 < 15; fi52++ {
                var v53 int = (fi52 & v35)
                _ = fi52
                _ = v53
            }
            _ = fi49
        }
        v54 := GenS7{F0: v35, F1: 1000, F2: 1}
        _ = v54
    }
    v3 = 0.1
    var v55 uint8 = genFn0(-128)
    switch v35 {
    case 42:
        var v56 uint16 = 843
        _ = v56
    default:
        if true {
            v35 += 678
            var v57 float32 = v1
            var v58 uint16 = 716
            _ = v57
            _ = v58
        }
    }
    v59 := GenS8{F0: v34}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v55:", v55)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v55
    _ = v59
    _ = v8
}
