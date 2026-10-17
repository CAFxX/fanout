package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint8; F1 int32 }
type GenS3 struct { F0 int8; F1 int16; F2 int16 }
type GenS4 struct { F0 int16; F1 uint; F2 int8 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 uint32; F1 int32 }
type GenS8 struct { F0 int32 }
func genFn0(p0 int64, p1 int8, p2 string) uint16 {
    v7 := GenS1{F0: 2}
    v8 := GenS2{F0: 127, F1: -2147483648}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 0
}

func main() {
    var v1 int8 = 3
    v2 := make(map[int]int)
    v3 := make(map[string]int)
    var v4 float64 = 0.1
    v5 := []int{32767}
    var v6 uint16 = 255
    v9 := GenS2{F0: 246, F1: 42}
    fmt.Println(v1, v6, v4)
    v6 += genFn0(-1, (v1 % ((v1) | 1)), (("hello" + ("\n\t" + (("foo bar" + "\n\t") + ("" + "foo bar")))) + "foo bar"))
    for fi10 := 0; fi10 < 20; fi10++ {
        v6 = (v6 | (v6 >> ((uint(v6)) & 63)))
        v4 += float64(0.1)
        _ = fi10
    }
    if (float32(v4) != float32(float64(1.5))) {
        if false {
            v11 := GenS2{F0: 1, F1: -1}
            if false {
                switch 0 {
                case 128:
                    v1 = (v1 ^ (v1 >> ((65535) & 63)))
                case 9223372036854775807:
                    var v12 int8 = v1
                    _ = v12
                default:
                    for fi13 := 0; fi13 < 16; fi13++ {
                        var v14 uint16 = (v6 + v6)
                        v14 = 255
                        for fi15 := 0; fi15 < 11; fi15++ {
                            v16 := GenS3{F0: v1, F1: 32767, F2: 1000}
                            _ = len(v5)
                            _ = fi15
                            _ = v16
                        }
                        _ = fi13
                        _ = v14
                    }
                }
                var v17 int16 = -127
                switch "foo bar" {
                case "":
                    v2[918] = 128
                default:
                    for fi18 := 0; fi18 < 10; fi18++ {
                        v4 = 1e10
                        v6 -= 65535
                        var v19 int16 = (v17 & int16(v1))
                        _ = fi18
                        _ = v19
                    }
                }
                _ = v17
            } else {
                v20 := GenS4{F0: -32768, F1: 745, F2: v1}
                _ = v20
            }
            v4 = v4
            _ = v11
        }
        delete(v2, 1)
    }
    if (-2147483648 > -2147483648) {
        switch 128 {
        case 32768:
            switch 1 {
            case 32767:
                switch 0 {
                case 1:
                    v4 = v4
                case 2147483647:
                    if (-9223372036854775808 > -1) {
                        fmt.Println(v6, v4, v1)
                    } else {
                        if true {
                            fmt.Println(v4)
                            if false {
                                fmt.Println(v6)
                                if ((v6 ^ (v6 + v6)) > v6) {
                                    if len(v5) > 0 {
                                        v5[0] = 1
                                    }
                                    if (((1e10 / 1.5) + (-2.25 / 1.0)) <= 0.0) {
                                        v21 := GenS5{F0: v6}
                                        if (42 <= 0) {
                                            v6 -= (v6 >> ((4294967295) & 63))
                                            v22 := GenS4{F0: -32768, F1: 0, F2: 1}
                                            _ = v22
                                        }
                                        _ = v21
                                    } else {
                                        if ((true || false) || true) {
                                            _ = len(v3)
                                        }
                                        if len(v5) > 0 {
                                            v5[0] = 4294967295
                                        }
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                    }
                                    if true {
                                        v4 += v4
                                        v23 := v2[0]
                                        _ = v23
                                    } else {
                                        var v24 int = -633
                                        _ = v24
                                    }
                                }
                            }
                        }
                        v4 += float64((1e10 - (float32((v4 * v4)) * 1e10)))
                    }
                }
            case 2147483647:
                if len(v5) > 0 {
                    v5[0] = -988
                }
            default:
                v25 := GenS2{F0: 1, F1: -1}
                _ = v25
            }
        default:
            v1 = v1
        }
        var v26 uint8 = 128
        _ = v26
    }
    for fi27 := 0; fi27 < 17; fi27++ {
        v28 := v2[fi27]
        if true {
            for fi29 := 0; fi29 < 5; fi29++ {
                v28 -= fi27
                v4 = -2.25
                if (false && true) {
                    var v30 int64 = -1
                    var v31 int8 = 127
                    _ = v30
                    _ = v31
                } else {
                    switch fi29 {
                    case 9223372036854775807:
                        fmt.Println(fi27, v6, fi29)
                    case -589:
                        v3["a"] = fi29
                    case 982:
                        v1 -= -30
                    default:
                        v28 += v28
                    }
                }
                _ = fi29
            }
            v32 := GenS6{F0: -714}
            v6 -= 445
            _ = v32
        }
        _ = fi27
        _ = v28
    }
    switch 32768 {
    case 1:
        v33 := GenS7{F0: 378, F1: 0}
        _ = v33
    case 2147483647:
        for fi34 := 0; fi34 < 6; fi34++ {
            v6 = 80
            _ = fi34
        }
    default:
        if ((float32((0.0 / v4)) * -2.25) < (0.1 + float32(v4))) {
            if ((v4 + v4) == ((v4 / -2.25) / v4)) {
                fmt.Println(v1)
            } else {
                v35 := GenS8{F0: -1}
                _ = v35
            }
        }
    }
    for fi36 := 0; fi36 < 5; fi36++ {
        if false {
            v37 := GenS2{F0: 128, F1: 1}
            _ = len(v5)
            _ = v37
        }
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        fmt.Println(fi36, v6)
        _ = fi36
    }
    if ("\n\t" != "hello") {
        v38 := GenS7{F0: uint32(v6), F1: 239}
        _ = v38
    }
    if (178 < -114) {
        v1 = -1
        fmt.Println(v4, v6)
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
