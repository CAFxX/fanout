package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }
type GenS2 struct { F0 int16; F1 uint16 }
func genFn0(p0 uint8, p1 uint8, p2 int32) uint {
    var v5 uint64 = 1
    var v6 uint8 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 65535
}
func genFn1() uint {
    var v7 uint32 = 65535
    for fi8 := 0; fi8 < 20; fi8++ {
        v9 := []int{fi8, 4294967295, -9223372036854775808}
        v7 = (v7 % ((v7) | 1))
        switch "a" {
        case "foo bar":
            if (255 != 1) {
                var v10 int8 = -128
                _ = v10
            } else {
                _ = len(v9)
                for fi11 := 0; fi11 < 5; fi11++ {
                    fmt.Println(fi11)
                    v7 = v7
                    _ = fi11
                }
                v7 = (v7 / (((v7 & (v7 * v7))) | 1))
            }
        }
        _ = fi8
        _ = v9
    }
    var v12 uint64 = 191
    _ = v12
    _ = v7
    return 1
}

func main() {
    var v1 int64 = -737
    var v2 int16 = -823
    v3 := make(map[uint32]string)
    var v4 uint64 = 263
    v13 := GenS1{F0: 2, F1: -55}
    v4 = v4
    switch 128 {
    case -547:
        var v14 float32 = ((1e10 + (0.0 * 1.5)) + ((-2.25 - 0.0) + 3.14159))
        _ = v14
    }
    var v15 int32 = 42
    v16 := GenS1{F0: 255, F1: 7}
    var v17 int64 = v1
    if ((939 == 0) && (120 >= 1)) {
        for fi18 := 0; fi18 < 19; fi18++ {
            v15 = -2147483648
            _ = fi18
        }
    }
    fmt.Println(v17)
    if (true && true) {
        var v19 float32 = (-2.25 - 1e10)
        v17 -= (v17 ^ (v17 + v17))
        _ = v19
    } else {
        if (!(v17 == -9223372036854775808)) {
            var v20 float64 = (0.1 * 0.1)
            _ = v20
        } else {
            v21 := GenS2{F0: 1000, F1: 700}
            switch "foo bar" {
            case "foo bar":
                var v22 string = ("hello" + ("\n\t" + "hello"))
                _ = v22
            case "a":
                for fi23 := 0; fi23 < 7; fi23++ {
                    v24 := GenS2{F0: v2, F1: 2}
                    _ = fi23
                    _ = v24
                }
            case "\n\t":
                switch "a" {
                case "foo bar":
                    for fi25 := 0; fi25 < 19; fi25++ {
                        if ((v15 + v15) < 1) {
                            if ((((!false) || true) || (false || true)) || false) {
                                v4 -= (v4 ^ (v4 % ((2) | 1)))
                                var v26 int8 = -72
                                v1 = v17
                                _ = v26
                            } else {
                                if (-2.25 == float32(0.1)) {
                                    if (1 < 859) {
                                        if (42 == 42) {
                                            v27 := []int{32768, (fi25 + (fi25 - (fi25 | fi25))), (^(fi25)), 32768}
                                            v2 -= (v2 % ((v2) | 1))
                                            fmt.Println(v2, v17)
                                            _ = v27
                                        } else {
                                            v28 := GenS2{F0: v2, F1: 0}
                                            v4 = v4
                                            _ = v28
                                        }
                                    } else {
                                        var v29 uint16 = 32767
                                        var v30 float32 = (0.0 + 0.1)
                                        for fi31 := 0; fi31 < 1; fi31++ {
                                            fmt.Println(v29, v4, v4)
                                            _ = fi31
                                        }
                                        _ = v29
                                        _ = v30
                                    }
                                    for fi32 := 0; fi32 < 20; fi32++ {
                                        var v33 float64 = 1.5
                                        _ = fi32
                                        _ = v33
                                    }
                                    fmt.Println(fi25, v17)
                                } else {
                                    v1 = v17
                                    v34 := GenS2{F0: v2, F1: 0}
                                    fmt.Println(v17, v1)
                                    _ = v34
                                }
                                for fi35 := 0; fi35 < 12; fi35++ {
                                    var v36 int16 = v2
                                    var v37 uint16 = 1
                                    _ = fi35
                                    _ = v36
                                    _ = v37
                                }
                            }
                        } else {
                            v2 -= -1
                        }
                        fmt.Println(v17, v1)
                        _ = fi25
                    }
                case "a":
                    var v38 float64 = 3.14159
                    _ = v38
                }
            }
            _ = v21
        }
        var v39 bool = false
        _ = v39
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
