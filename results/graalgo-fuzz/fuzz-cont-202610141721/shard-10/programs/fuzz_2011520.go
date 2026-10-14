package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint64 }

func main() {
    var v1 float32 = ((3.14159 - (3.14159 + (-2.25 + -2.25))) + ((float32((3.14159 * 0.1)) - 3.14159) + 0.1))
    var v2 int32 = -2147483648
    var v3 int = -9223372036854775808
    v4 := []uint32{42, 42, 188, 2}
    var v5 int = (v3 - v3)
    for fi6 := 0; fi6 < 6; fi6++ {
        _ = len(v4)
        _ = fi6
    }
    for fi7 := 0; fi7 < 16; fi7++ {
        v2 -= (v2 + 0)
        _ = fi7
    }
    v2 -= (-(v2))
    var v8 uint64 = 145
    if ((false || (true || (65535 < 0))) && (!true)) {
        v8 += 2
    } else {
        var v9 uint8 = 1
        if len(v4) > 0 {
            v4[0] = 4294967295
        }
        v9 = v9
        _ = v9
    }
    v1 = float32(float64(-2.25))
    v5 -= (v3 ^ v5)
    switch (v3 % ((0) | 1)) {
    case 42:
        if (false && (9223372036854775807 != 1)) {
            if len(v4) > 0 {
                v4[0] = 4294967295
            }
        } else {
            var v10 uint32 = 1
            switch 128 {
            case 0:
                for fi11 := 0; fi11 < 4; fi11++ {
                    v1 -= float32(-2.25)
                    switch "a" {
                    case "foo bar":
                        v12 := GenS1{F0: (^(fi11))}
                        _ = v12
                    default:
                        if ((2 >= uint(v10)) || (1.5 <= -2.25)) {
                            v5 -= fi11
                            for fi13 := 0; fi13 < 13; fi13++ {
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                _ = fi13
                            }
                        } else {
                            v4 = append(v4, v10)
                            var v14 uint64 = v8
                            var v15 int16 = 1
                            _ = v14
                            _ = v15
                        }
                    }
                    _ = fi11
                }
            case 2147483647:
                var v16 int64 = 42
                _ = v16
            case 32767:
                for fi17 := 0; fi17 < 1; fi17++ {
                    if (1e10 >= -2.25) {
                        v18 := GenS2{F0: (v8 | (v8 & v8))}
                        v3 = (^(fi17))
                        for fi19 := 0; fi19 < 2; fi19++ {
                            var v20 float32 = v1
                            v8 = 1
                            _ = fi19
                            _ = v20
                        }
                        _ = v18
                    } else {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    }
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    _ = fi17
                }
            default:
                _ = len(v4)
            }
            if ((float32(1e10) + float32(1.5)) <= float32(0.0)) {
                if ((((v2 & v2) > -1) || (!true)) || false) {
                    fmt.Println(v10, v1, v1)
                }
                var v21 uint32 = 2
                _ = len(v4)
                _ = v21
            } else {
                var v22 int = (v3 * v3)
                var v23 float32 = 0.1
                switch ("" + "hello") {
                case "\n\t":
                    fmt.Println(v3)
                }
                _ = v22
                _ = v23
            }
            _ = v10
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
