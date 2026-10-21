package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint16; F1 int32 }

func main() {
    var v1 uint8 = 22
    v2 := []string{"a"}
    var v3 string = ("" + "\n\t")
    var v4 int = 791
    var v5 int16 = 557
    for fi6 := 0; fi6 < 4; fi6++ {
        if false {
            if (true || ((v5 & 0) <= v5)) {
                if len(v2) > 0 {
                    v2[0] = v3
                }
            } else {
                v3 = ""
                if (float32(1.5) < 1.5) {
                    if true {
                        switch v3 {
                        case "a":
                            var v7 string = v3
                            _ = v7
                        case "":
                            fmt.Println(v1, v5)
                        default:
                            if true {
                                for fi8 := 0; fi8 < 7; fi8++ {
                                    var v9 int = v4
                                    _ = fi8
                                    _ = v9
                                }
                                if true {
                                    var v10 int64 = 0
                                    _ = v10
                                } else {
                                    fmt.Println(v1, fi6, v1)
                                }
                            }
                        }
                        _ = len(v2)
                    }
                } else {
                    v4 += (fi6 % (((v4 / (((-(fi6))) | 1))) | 1))
                }
            }
            switch v4 {
            case 491:
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            case -682:
                v11 := GenS1{F0: 2147483648}
                _ = v11
            case 32767:
                v4 = (-(v4))
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
        var v12 int64 = -9223372036854775808
        v4 += (fi6 >> ((4294967295) & 63))
        _ = fi6
        _ = v12
    }
    switch v4 {
    case 1:
        v4 = v4
    }
    v1 += (^(v1))
    for fi13 := 0; fi13 < 17; fi13++ {
        if false {
            fmt.Println(v3, fi13)
        }
        _ = fi13
    }
    v4 = (v4 * v4)
    fmt.Println(v4, v5, v1)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    _ = len(v2)
    switch (v4 << ((308) & 63)) {
    case -2147483648:
        v4 = (v4 - v4)
    case 1:
        if ((-2.25 + ((0.0 * -2.25) / float64(1.5))) <= -2.25) {
            var v14 uint8 = 157
            for fi15 := 0; fi15 < 1; fi15++ {
                var v16 int8 = 0
                v17 := GenS1{F0: 1}
                _ = fi15
                _ = v16
                _ = v17
            }
            _ = v14
        } else {
            var v18 int = int(v1)
            v18 += 544
            v19 := GenS1{F0: 2}
            _ = v18
            _ = v19
        }
    }
    v20 := GenS2{F0: 682, F1: 2147483647}
    var v21 float64 = float64(0.0)
    v5 = (v5 << ((0) & 63))
    fmt.Println(v3)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
