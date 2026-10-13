package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }

func main() {
    v1 := []string{"foo bar"}
    var v2 uint = 91
    var v3 int64 = 1
    v4 := []uint32{0, 42, 522, 0}
    var v5 bool = false
    v1 = append(v1, "a")
    for fi6 := 0; fi6 < 9; fi6++ {
        var v7 bool = v5
        _ = fi6
        _ = v7
    }
    var v8 int8 = -113
    if len(v4) > 0 {
        v4[0] = 42
    }
    if (255 == 505) {
        if len(v4) > 0 {
            v4[0] = 1
        }
        if (float64(float32((((1.5 / 1.0) / 1.0) - 3.14159))) <= -2.25) {
            var v9 float64 = float64((float32(-2.25) / 0.1))
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = v9
        }
        var v10 string = ("" + "hello")
        _ = v10
    } else {
        v4 = append(v4, 65535)
        if v5 {
            var v11 int64 = 42
            for fi12 := 0; fi12 < 9; fi12++ {
                for fi13 := 0; fi13 < 3; fi13++ {
                    var v14 int16 = int16(fi13)
                    _ = fi13
                    _ = v14
                }
                var v15 bool = (32767 >= int16(v8))
                if v5 {
                    if v5 {
                        var v16 int8 = int8(fi12)
                        v8 += v16
                        if false {
                            var v17 int8 = 71
                            _ = v17
                        } else {
                            switch "\n\t" {
                            case "hello":
                                fmt.Println(v15, v8)
                            case "a":
                                var v18 float32 = 0.1
                                _ = v18
                            case "\n\t":
                                v4 = append(v4, 42)
                            default:
                                for fi19 := 0; fi19 < 16; fi19++ {
                                    _ = len(v1)
                                    _ = fi19
                                }
                            }
                            switch "" {
                            case "hello":
                                v11 -= 42
                            default:
                                for fi20 := 0; fi20 < 1; fi20++ {
                                    v21 := GenS1{F0: 18446744073709551615, F1: 71}
                                    var v22 uint8 = 1
                                    _ = fi20
                                    _ = v21
                                    _ = v22
                                }
                            }
                        }
                        _ = v16
                    }
                }
                _ = fi12
                _ = v15
            }
            _ = v11
        }
    }
    var v23 uint64 = 1
    if false {
        for fi24 := 0; fi24 < 3; fi24++ {
            v5 = false
            switch 127 {
            case -191:
                v25 := GenS1{F0: v23, F1: 255}
                _ = v25
            case 716:
                var v26 float64 = 3.14159
                _ = v26
            }
            switch (^(fi24)) {
            case -1:
                switch 610 {
                case -9223372036854775808:
                    if len(v4) > 0 {
                        v4[0] = 0
                    }
                case 127:
                    v27 := GenS1{F0: v23, F1: 42}
                    _ = v27
                }
            default:
                var v28 float32 = 3.14159
                _ = v28
            }
            _ = fi24
        }
        v23 -= (v23 | (v23 << ((v2) & 63)))
        _ = len(v4)
    }
    var v29 uint64 = 526
    var v30 int64 = v3
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v23
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v8
}
