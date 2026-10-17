package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8; F2 bool }
type GenS2 struct { F0 bool; F1 int16; F2 uint8 }
func genFn0(p0 uint64) int64 {
    p0 = p0
    fmt.Println(p0)
    p0 = (p0 + 42)
    p0 = (p0 >> ((1) & 63))
    p0 -= (^(p0))
    _ = p0
    return 1
}
func genFn1(p0 uint) bool {
    switch -903 {
    case 106:
        switch 292 {
        case -1:
            fmt.Println(p0)
        case 102:
            for fi5 := 0; fi5 < 11; fi5++ {
                v6 := GenS1{F0: -132, F1: 1, F2: true}
                _ = fi5
                _ = v6
            }
        default:
            fmt.Println(p0)
        }
    }
    var v7 uint32 = 2147483647
    v8 := []int{}
    _ = p0
    _ = v7
    _ = v8
    return true
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[int]bool)
    var v3 int16 = -32768
    v4 := []uint32{65535, 1, 181}
    fmt.Println(v3)
    delete(v2, -9223372036854775808)
    fmt.Println(v3)
    var v9 float32 = (float32((3.14159 / 1.0)) / (float32(float64((1e10 * (-2.25 - (1e10 * -2.25))))) / -2.25))
    fmt.Println(v3, v3)
    v3 = v3
    var v10 uint8 = 68
    if (true && (!false)) {
        for fi11 := 0; fi11 < 9; fi11++ {
            if true {
                if len(v4) > 0 {
                    v4[0] = 4294967295
                }
            }
            fmt.Println(v10, v10, v9)
            _ = fi11
        }
        v1[4294967295] = 32767
    } else {
        if (!(!true)) {
            var v12 string = "foo bar"
            switch v12 {
            case "a":
                for fi13 := 0; fi13 < 4; fi13++ {
                    var v14 float32 = 3.14159
                    var v15 int16 = (v3 ^ (v3 ^ v3))
                    v3 += v15
                    _ = fi13
                    _ = v14
                    _ = v15
                }
            case "":
                for fi16 := 0; fi16 < 12; fi16++ {
                    for fi17 := 0; fi17 < 12; fi17++ {
                        if false {
                            switch -9223372036854775808 {
                            case -856:
                                v9 = (v9 - v9)
                            case 1:
                                switch fi17 {
                                case -1:
                                    var v18 uint8 = v10
                                    _ = v18
                                case 32768:
                                    switch (fi17 & 2147483647) {
                                    case -915:
                                        var v19 int8 = 0
                                        _ = v19
                                    }
                                default:
                                    if (false && (!(false || (v12 != "")))) {
                                        fmt.Println(v3, v12, fi16)
                                        if len(v4) > 0 {
                                            v4[0] = 65535
                                        }
                                        v20 := v1[318]
                                        v21 := GenS1{F0: 0, F1: 42, F2: true}
                                        _ = v20
                                        _ = v21
                                    }
                                }
                            default:
                                switch v12 {
                                case "foo bar":
                                    var v22 int8 = 26
                                    _ = v22
                                case "hello":
                                    var v23 int32 = 0
                                    _ = v23
                                default:
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                }
                            }
                            fmt.Println(v3, v10)
                            v24 := v1[1]
                            var v25 int16 = v3
                            var v26 string = ("" + "hello")
                            _ = v24
                            _ = v25
                            _ = v26
                        }
                        v10 -= 127
                        _ = fi17
                    }
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    for fi27 := 0; fi27 < 17; fi27++ {
                        if len(v4) > 0 {
                            v4[0] = 42
                        }
                        if ((v3 & (v3 % (((v3 & (v3 ^ v3))) | 1))) == v3) {
                            v12 = ("\n\t" + "\n\t")
                            v9 += 1.5
                            fmt.Println(v10, v12)
                            v2[4294967295] = false
                        } else {
                            v9 -= v9
                            v9 += (float32(1.5) - 3.14159)
                            if len(v4) > 0 {
                                v4[0] = 2147483647
                            }
                        }
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        _ = fi27
                    }
                    _ = fi16
                }
            case "hello":
                var v28 int16 = v3
                _ = v28
            default:
                var v29 string = v12
                _ = v29
            }
            v10 -= (v10 / ((v10) | 1))
            _ = v12
        }
        for fi30 := 0; fi30 < 7; fi30++ {
            fmt.Println(fi30, v3, v10)
            _ = fi30
        }
    }
    fmt.Println(v9)
    for fi31 := 0; fi31 < 6; fi31++ {
        v32 := GenS2{F0: (712 != 282), F1: v3, F2: (v10 >> ((763) & 63))}
        if false {
            var v33 uint16 = 188
            _ = v33
        }
        for fi34 := 0; fi34 < 14; fi34++ {
            _ = len(v4)
            _ = fi34
        }
        _ = fi31
        _ = v32
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
