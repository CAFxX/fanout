package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64; F2 uint32 }
type GenS2 struct { F0 uint16; F1 int64; F2 int8 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint8 = 0
    var v3 uint64 = 1
    v2 = 243
    var v4 int16 = 794
    fmt.Println(v3)
    switch 42 {
    case 880:
        _ = len(v1)
    default:
        switch "a" {
        case "hello":
            v3 = 42
        case "a":
            _ = len(v1)
        case "":
            fmt.Println(v4)
        }
    }
    v5 := v1[2147483648]
    var v6 int16 = (^(v4))
    v5 += (v5 ^ v5)
    if false {
        v7 := GenS1{F0: v3, F1: 1, F2: v5}
        var v8 int16 = v6
        _ = v7
        _ = v8
    } else {
        for fi9 := 0; fi9 < 4; fi9++ {
            v10 := GenS2{F0: 2, F1: -1, F2: -1}
            v11 := GenS1{F0: (v3 ^ v3), F1: -1, F2: (v5 >> ((416) & 63))}
            _ = fi9
            _ = v10
            _ = v11
        }
    }
    var v12 int16 = -256
    if (1e10 == float64(1.5)) {
        for fi13 := 0; fi13 < 20; fi13++ {
            for fi14 := 0; fi14 < 7; fi14++ {
                v6 += v4
                if (true && true) {
                    v15 := []int{(fi14 ^ 133)}
                    _ = v15
                }
                _ = fi14
            }
            for fi16 := 0; fi16 < 7; fi16++ {
                v17 := v1[65535]
                v4 = v4
                _ = fi16
                _ = v17
            }
            if false {
                var v18 uint64 = uint64(v2)
                if (!(false || ((-(fi13)) == (-(fi13))))) {
                    var v19 float32 = 0.1
                    v20 := GenS2{F0: uint16(v5), F1: 1, F2: 113}
                    _ = v19
                    _ = v20
                } else {
                    v12 = v12
                }
                _ = v18
            } else {
                if (65535 <= 350) {
                    var v21 bool = (!true)
                    delete(v1, uint32(v2))
                    _ = v21
                } else {
                    for fi22 := 0; fi22 < 19; fi22++ {
                        if ((v3 * (v3 / (((v3 >> ((2) & 63))) | 1))) == (v3 / (((v3 - v3)) | 1))) {
                            var v23 uint32 = (v5 + 4294967295)
                            _ = v23
                        } else {
                            fmt.Println(v6, v4)
                            if ("a" != "foo bar") {
                                switch "" {
                                case "hello":
                                    v3 += v3
                                }
                                v24 := []int{(^(fi22)), fi13, (^(fi13))}
                                var v25 int64 = 9223372036854775807
                                _ = v24
                                _ = v25
                            }
                            _ = len(v1)
                        }
                        _ = fi22
                    }
                }
            }
            _ = fi13
        }
        if (true || (-952 < 4294967295)) {
            v26 := []int{354}
            v5 = v5
            _ = v26
        } else {
            for fi27 := 0; fi27 < 9; fi27++ {
                v28 := []int{0, 2147483647, fi27, (-(fi27))}
                var v29 uint16 = 0
                _ = fi27
                _ = v28
                _ = v29
            }
        }
    }
    var v30 string = ("\n\t" + "hello")
    v3 = (v3 % ((v3) | 1))
    if (true && true) {
        v31 := GenS1{F0: (v3 | v3), F1: 246, F2: v5}
        _ = v31
    } else {
        v12 = -365
        fmt.Println(v12, v2)
        var v32 bool = ((314 < int(v2)) || false)
        _ = v32
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
