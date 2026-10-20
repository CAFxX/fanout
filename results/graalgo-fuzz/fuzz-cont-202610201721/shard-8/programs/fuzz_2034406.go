package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }
type GenS2 struct { F0 int32; F1 int16; F2 int32 }
type GenS3 struct { F0 uint16; F1 bool }
func genFn0() uint16 {
    var v7 uint8 = 127
    v8 := GenS1{F0: 707, F1: 42}
    _ = v7
    _ = v8
    return 2
}

func main() {
    var v1 bool = ("foo bar" > ("a" + ("\n\t" + "")))
    v2 := make(map[uint32]bool)
    v3 := make(map[int]bool)
    v4 := []int{}
    var v5 float32 = ((((3.14159 - (3.14159 * 1.5)) * 0.1) / 3.14159) - -2.25)
    var v6 int16 = 0
    v9 := GenS2{F0: 1, F1: v6, F2: 42}
    v1 = (!((("foo bar" + "") + "") < "a"))
    v6 -= 32767
    _ = len(v4)
    fmt.Println(v5)
    var v10 uint8 = 95
    v1 = true
    if len(v4) > 0 {
        v4[0] = 525
    }
    if v1 {
        switch -9223372036854775808 {
        case -2147483648:
            delete(v3, 4294967295)
        case -1:
            v1 = false
        case 128:
            for fi11 := 0; fi11 < 15; fi11++ {
                if (true || (0 == 538)) {
                    v5 = ((v5 - 1.5) - v5)
                }
                if (float64(v5) >= 0.1) {
                    v6 -= v6
                    v6 = v6
                    var v12 int16 = (v6 << ((4294967295) & 63))
                    _ = v12
                }
                fmt.Println(v1, v5, v6)
                _ = fi11
            }
        default:
            fmt.Println(v1, v5)
        }
        for fi13 := 0; fi13 < 14; fi13++ {
            var v14 int16 = (v6 / (((v6 >> ((0) & 63))) | 1))
            fmt.Println(v6, fi13)
            if false {
                for fi15 := 0; fi15 < 9; fi15++ {
                    if (v1 || (0.1 >= (-2.25 + 1e10))) {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    }
                    _ = fi15
                }
            } else {
                var v16 int32 = -1
                switch (-(fi13)) {
                case -2147483648:
                    switch 32768 {
                    case 891:
                        for fi17 := 0; fi17 < 9; fi17++ {
                            v14 -= (^(v6))
                            v10 += (^(v10))
                            if (genFn0() != 795) {
                                var v18 int8 = -1
                                v19 := GenS3{F0: 65535, F1: v1}
                                v18 = v18
                                _ = v18
                                _ = v19
                            } else {
                                var v20 uint64 = 1
                                v21 := GenS3{F0: genFn0(), F1: (v1 && (((0 >= 9223372036854775807) || v1) || v1))}
                                var v22 uint64 = 1
                                _ = v20
                                _ = v21
                                _ = v22
                            }
                            _ = fi17
                        }
                    case 32767:
                        var v23 uint64 = 1
                        _ = v23
                    case 285:
                        if (!v1) {
                            var v24 uint64 = 89
                            v4 = append(v4, fi13)
                            _ = v24
                        } else {
                            v4 = append(v4, fi13)
                            for fi25 := 0; fi25 < 12; fi25++ {
                                v10 += v10
                                switch (("foo bar" + "") + "foo bar") {
                                case "foo bar":
                                    var v26 int = (fi25 - fi13)
                                    _ = v26
                                case "\n\t":
                                    var v27 int32 = (v16 & v16)
                                    _ = v27
                                default:
                                    if len(v4) > 0 {
                                        v4[0] = (fi25 + fi13)
                                    }
                                }
                                _ = fi25
                            }
                        }
                    default:
                        var v28 int8 = -128
                        _ = v28
                    }
                case 349:
                    v29 := GenS2{F0: (v16 >> ((0) & 63)), F1: (^(v14)), F2: v16}
                    _ = v29
                case 127:
                    if false {
                        var v30 int = (fi13 * (fi13 * fi13))
                        _ = v30
                    }
                default:
                    var v31 float32 = float32(1.5)
                    _ = v31
                }
                if v1 {
                    var v32 int64 = 9223372036854775807
                    _ = v32
                }
                _ = v16
            }
            _ = fi13
            _ = v14
        }
    } else {
        if v1 {
            switch -27 {
            case 0:
                v6 = 32767
            case 9223372036854775807:
                v33 := GenS1{F0: 255, F1: -128}
                _ = v33
            }
            v34 := GenS1{F0: 255, F1: -128}
            if v1 {
                fmt.Println(v1, v10, v10)
                fmt.Println(v5)
            }
            _ = v34
        }
    }
    v35 := v2[2147483647]
    var v36 uint64 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
