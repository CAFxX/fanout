package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 bool; F1 int; F2 int }
func genFn0() bool {
    if true {
        var v6 float32 = float32(1e10)
        v7 := GenS1{F0: true}
        _ = v6
        _ = v7
    }
    var v8 uint8 = 1
    var v9 uint8 = (v8 * 42)
    _ = v8
    _ = v9
    return true
}

func main() {
    var v1 uint64 = 2
    v2 := []uint32{2, 4294967295}
    v3 := make(map[uint32]uint32)
    var v4 string = ""
    var v5 int8 = 100
    v1 = (^(v1))
    var v10 uint8 = 1
    if ((v1 ^ v1) == (^(v1))) {
        if (0 > -2147483648) {
            var v11 int = -619
            var v12 int = (v11 | (v11 | (v11 ^ v11)))
            _ = v11
            _ = v12
        }
        v13 := GenS2{F0: true, F1: 4294967295, F2: -793}
        v5 = (v5 / (((v5 & v5)) | 1))
        _ = v13
    }
    v4 = ""
    var v14 int16 = 1000
    if true {
        var v15 uint32 = 501
        switch 4294967295 {
        case 2147483647:
            if false {
                var v16 int8 = (v5 ^ (-(v5)))
                if false {
                    v17 := GenS1{F0: (!(true || true))}
                    for fi18 := 0; fi18 < 17; fi18++ {
                        v14 -= (v14 >> ((1) & 63))
                        for fi19 := 0; fi19 < 6; fi19++ {
                            if (true || true) {
                                if len(v2) > 0 {
                                    v2[0] = 1
                                }
                                v5 -= (v5 & (^(v5)))
                                if true {
                                    v5 = (v16 ^ int8(v15))
                                    var v20 string = (v4 + v4)
                                    _ = v20
                                }
                                if len(v2) > 0 {
                                    v2[0] = (v15 >> ((255) & 63))
                                }
                                fmt.Println(v15)
                            }
                            v3[(v15 - 815)] = v15
                            var v21 int8 = -1
                            _ = fi19
                            _ = v21
                        }
                        _ = fi18
                    }
                    if false {
                        var v22 uint16 = 42
                        _ = v22
                    }
                    _ = v17
                }
                fmt.Println(v10, v15, v5)
                _ = v16
            }
        case -898:
            v4 = (v4 + v4)
        default:
            var v23 int = 127
            _ = v23
        }
        v24 := GenS2{F0: (130 >= 255), F1: -9223372036854775808, F2: int(v5)}
        _ = v15
        _ = v24
    }
    var v25 float32 = (-2.25 / (-2.25 - (1.5 + 0.0)))
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
