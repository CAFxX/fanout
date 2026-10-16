package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 int64 }
type GenS2 struct { F0 int64; F1 uint32 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int64 }

func main() {
    v1 := []int{32768, -1, 9223372036854775807, 128}
    var v2 float64 = 1e10
    var v3 float64 = (0.0 + v2)
    var v4 float64 = (float64(0.1) / 1e10)
    var v5 uint32 = 258
    v6 := make(map[int]int)
    v7 := GenS1{F0: 413, F1: -219, F2: 1}
    if (-257 <= 1) {
        var v8 uint = 65535
        _ = v8
    }
    if true {
        v9 := GenS1{F0: 390, F1: -1, F2: -9223372036854775808}
        _ = v9
    }
    var v10 int32 = -2147483648
    var v11 uint64 = 4294967295
    if ((v10 << ((702) & 63)) != v10) {
        for fi12 := 0; fi12 < 10; fi12++ {
            var v13 int8 = 0
            _ = fi12
            _ = v13
        }
        if ((v11 + v11) < (v11 * (v11 * (^(v11))))) {
            switch ("hello" + (("foo bar" + "") + "")) {
            case "foo bar":
                if true {
                    v10 += (v10 / ((v10) | 1))
                } else {
                    v11 += v11
                    if (-2147483648 == v10) {
                        v3 -= (v3 - v2)
                    }
                    v3 += ((v4 * ((1e10 + float64(((0.1 - 0.1) * (0.1 * (1.5 - 3.14159))))) - (1e10 / 0.1))) / 3.14159)
                }
            case "":
                switch 128 {
                case 1:
                    for fi14 := 0; fi14 < 6; fi14++ {
                        var v15 int64 = -9223372036854775808
                        _ = fi14
                        _ = v15
                    }
                case 32767:
                    v10 -= (v10 & (-(v10)))
                case 9223372036854775807:
                    var v16 uint = 1
                    _ = v16
                }
            case "\n\t":
                switch -861 {
                case -9223372036854775808:
                    for fi17 := 0; fi17 < 12; fi17++ {
                        switch ("foo bar" + (("\n\t" + "foo bar") + (("hello" + "a") + ""))) {
                        case "":
                            for fi18 := 0; fi18 < 18; fi18++ {
                                var v19 int64 = 42
                                v2 += ((1.5 * (v4 * (float64((((1e10 + 1.5) / 3.14159) * 0.1)) * 1.5))) - float64(float32(float64((1.5 + (0.1 / -2.25))))))
                                v6[(fi17 + -9223372036854775808)] = fi18
                                _ = fi18
                                _ = v19
                            }
                        case "a":
                            var v20 bool = ((^(v11)) == v11)
                            _ = v20
                        }
                        _ = fi17
                    }
                case 32767:
                    v21 := GenS1{F0: 100, F1: -9223372036854775808, F2: 0}
                    _ = v21
                case 9223372036854775807:
                    if (-1 != 42) {
                        var v22 uint32 = v5
                        if true {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        } else {
                            v23 := GenS2{F0: 1, F1: v22}
                            if true {
                                _ = len(v1)
                            } else {
                                v1 = append(v1, 128)
                            }
                            for fi24 := 0; fi24 < 15; fi24++ {
                                v10 -= v10
                                switch "a" {
                                case "foo bar":
                                    for fi25 := 0; fi25 < 8; fi25++ {
                                        v22 = (v22 % (((^(v5))) | 1))
                                        _ = fi25
                                    }
                                default:
                                    if false {
                                        var v26 float64 = (float64(float32(1.5)) - (v2 / 0.1))
                                        for fi27 := 0; fi27 < 12; fi27++ {
                                            fmt.Println(v4, v3, v10)
                                            _ = fi27
                                        }
                                        var v28 int16 = -1
                                        _ = v26
                                        _ = v28
                                    } else {
                                        var v29 int16 = -32768
                                        _ = v29
                                    }
                                }
                                v3 = v2
                                _ = fi24
                            }
                            _ = v23
                        }
                        var v30 int = 4294967295
                        _ = v22
                        _ = v30
                    }
                default:
                    var v31 uint16 = 0
                    _ = v31
                }
            default:
                v32 := GenS1{F0: 42, F1: int64(v5), F2: 9223372036854775807}
                _ = v32
            }
            v33 := GenS3{F0: 18446744073709551615}
            _ = v33
        } else {
            if ((v5 * 2) == v5) {
                for fi34 := 0; fi34 < 10; fi34++ {
                    if true {
                        if true {
                            fmt.Println(v4, v10)
                            if false {
                                fmt.Println(v4, v2)
                            }
                            if (!true) {
                                v2 += v3
                                v35 := GenS2{F0: 607, F1: 65535}
                                var v36 string = "a"
                                _ = v35
                                _ = v36
                            }
                        } else {
                            var v37 bool = true
                            v38 := GenS4{F0: 255}
                            _ = v37
                            _ = v38
                        }
                    }
                    _ = fi34
                }
                fmt.Println(v3)
                switch -1 {
                case -9223372036854775808:
                    v4 -= v2
                default:
                    v39 := GenS5{F0: 710}
                    _ = v39
                }
            } else {
                fmt.Println(v4, v4, v5)
                v4 = (v4 + (v3 + v3))
            }
        }
        var v40 int8 = -1
        _ = v40
    } else {
        for fi41 := 0; fi41 < 4; fi41++ {
            if len(v1) > 0 {
                v1[0] = fi41
            }
            fmt.Println(v3)
            var v42 int = fi41
            _ = fi41
            _ = v42
        }
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println(v3, v2)
    fmt.Println(v10, v11, v2)
    var v43 int8 = 127
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v5
    _ = v6
    _ = v7
}
