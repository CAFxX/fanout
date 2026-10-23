package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16; F2 uint32 }
type GenS2 struct { F0 string }
func genFn0() uint32 {
    var v7 int = 4294967295
    v8 := GenS1{F0: 758, F1: 42, F2: 4294967295}
    _ = v7
    _ = v8
    return 382
}

func main() {
    var v1 string = "\n\t"
    var v2 int64 = 42
    var v3 uint64 = 718
    var v4 uint64 = (v3 + v3)
    var v5 string = v1
    var v6 int = 42
    v1 = v1
    if ((4294967295 != 65535) && true) {
        v6 += (v6 >> ((263) & 63))
        fmt.Println(v5, v1, v6)
    } else {
        switch v1 {
        case "":
            v4 = 2
        default:
            v5 = "a"
        }
        for fi9 := 0; fi9 < 17; fi9++ {
            switch 32767 {
            case 127:
                var v10 float32 = (0.0 * float32(1.5))
                _ = v10
            default:
                var v11 int16 = 446
                _ = v11
            }
            _ = fi9
        }
        v12 := []int{-1, (^(v6)), v6, (^(v6))}
        _ = v12
    }
    v13 := []int{(v6 / ((v6) | 1)), (v6 % ((v6) | 1)), 32767, (^(v6))}
    v13 = append(v13, (v6 | v6))
    if true {
        _ = len(v13)
        var v14 uint = 65535
        fmt.Println(v14, v4)
        _ = v14
    } else {
        if len(v13) > 0 {
            v13[0] = -125
        }
    }
    v4 += (^(v3))
    var v15 int64 = 42
    var v16 string = "\n\t"
    v17 := GenS1{F0: (v15 >> ((255) & 63)), F1: 2, F2: 898}
    var v18 uint32 = genFn0()
    if len(v13) > 0 {
        v13[0] = (v6 ^ (-(v6)))
    }
    switch "" {
    case "":
        for fi19 := 0; fi19 < 3; fi19++ {
            if ((0 == 747) && true) {
                if false {
                    for fi20 := 0; fi20 < 4; fi20++ {
                        v5 = (v5 + v16)
                        _ = fi20
                    }
                } else {
                    for fi21 := 0; fi21 < 16; fi21++ {
                        var v22 float64 = -2.25
                        var v23 bool = false
                        if v23 {
                            if len(v13) >= 2 {
                                v13 = v13[1:len(v13)]
                            }
                            if true {
                                v13 = append(v13, (^(fi21)))
                                if len(v13) >= 2 {
                                    v13 = v13[1:len(v13)]
                                }
                                var v24 float64 = v22
                                _ = v24
                            }
                        } else {
                            fmt.Println(v5)
                            var v25 int8 = 0
                            if (702 < 254) {
                                v15 = (v2 | (v2 >> ((719) & 63)))
                            } else {
                                if v23 {
                                    v26 := GenS1{F0: (v2 - v2), F1: 255, F2: 2147483648}
                                    _ = v26
                                } else {
                                    if len(v13) > 0 {
                                        v13[0] = fi19
                                    }
                                    var v27 string = "a"
                                    v15 += 0
                                    _ = v27
                                }
                                _ = len(v13)
                                var v28 uint8 = 134
                                _ = v28
                            }
                            var v29 uint32 = v18
                            _ = v25
                            _ = v29
                        }
                        _ = fi21
                        _ = v22
                        _ = v23
                    }
                    v18 = (v18 + v18)
                    switch "" {
                    case "a":
                        var v30 uint8 = 127
                        _ = v30
                    case "foo bar":
                        if len(v13) >= 2 {
                            v13 = v13[1:len(v13)]
                        }
                    case "\n\t":
                        v18 = (v18 | 2147483648)
                    default:
                        var v31 bool = true
                        _ = v31
                    }
                }
            } else {
                v4 -= 18446744073709551615
            }
            _ = fi19
        }
    case "a":
        if (true || true) {
            v32 := GenS1{F0: -9223372036854775808, F1: 65535, F2: v18}
            for fi33 := 0; fi33 < 9; fi33++ {
                fmt.Println(fi33)
                _ = fi33
            }
            v34 := GenS2{F0: v1}
            _ = v32
            _ = v34
        } else {
            var v35 string = (("a" + ("foo bar" + "foo bar")) + ((v1 + (v16 + v5)) + "hello"))
            if (413 < v3) {
                var v36 float32 = 0.0
                v1 = "hello"
                for fi37 := 0; fi37 < 7; fi37++ {
                    v35 = ""
                    v36 += 0.0
                    _ = fi37
                }
                _ = v36
            }
            _ = v35
        }
    }
    v38 := GenS2{F0: v5}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
