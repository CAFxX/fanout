package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int64; F1 int64; F2 uint8 }
type GenS3 struct { F0 string; F1 uint64 }
type GenS4 struct { F0 int32; F1 uint }

func main() {
    var v1 int8 = 33
    var v2 int16 = 1000
    var v3 uint32 = 238
    v4 := []string{}
    var v5 int16 = (v2 | 1)
    v6 := []uint32{v3}
    for fi7 := 0; fi7 < 20; fi7++ {
        v8 := GenS1{F0: -1}
        v3 = (^(v3))
        _ = fi7
        _ = v8
    }
    v6 = append(v6, v3)
    fmt.Println(v1, v1)
    v9 := GenS1{F0: -446}
    v10 := GenS2{F0: -1, F1: 0, F2: uint8(v3)}
    var v11 uint16 = 1
    switch 2147483647 {
    case -881:
        if len(v4) > 0 {
            v4[0] = ""
        }
    default:
        var v12 uint32 = (v3 % (((v3 * (v3 >> ((2) & 63)))) | 1))
        _ = v12
    }
    v5 = v5
    for fi13 := 0; fi13 < 10; fi13++ {
        var v14 uint16 = (v11 / ((255) | 1))
        if false {
            v11 = (v11 & 2)
            for fi15 := 0; fi15 < 11; fi15++ {
                v4 = append(v4, ("" + "\n\t"))
                fmt.Println(v11, v11, v5)
                _ = fi15
            }
        }
        switch (fi13 + fi13) {
        case 2147483647:
            var v16 uint64 = 1
            _ = v16
        case 319:
            v17 := GenS3{F0: ("" + ""), F1: 726}
            _ = v17
        default:
            for fi18 := 0; fi18 < 18; fi18++ {
                if (104 <= 164) {
                    var v19 int64 = 0
                    _ = v19
                } else {
                    if false {
                        v14 = 255
                    } else {
                        if (false && true) {
                            v20 := GenS4{F0: -2147483648, F1: 338}
                            if true {
                                v21 := GenS3{F0: ("foo bar" + "a"), F1: 0}
                                _ = v21
                            }
                            _ = v20
                        }
                        var v22 int8 = v1
                        _ = v22
                    }
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    v14 -= (v14 >> ((351) & 63))
                }
                if true {
                    if len(v6) > 0 {
                        v6[0] = (v3 ^ (v3 / ((v3) | 1)))
                    }
                    v6 = append(v6, v3)
                }
                var v23 string = "a"
                var v24 int32 = 0
                v1 += v1
                v14 += (v11 ^ (v14 | (v14 - v14)))
                var v25 uint = 65535
                _ = fi18
                _ = v23
                _ = v24
                _ = v25
            }
        }
        _ = fi13
        _ = v14
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
