package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 bool }

func main() {
    var v1 uint32 = 4294967295
    v2 := []uint32{v1}
    var v3 uint = 2
    var v4 int = -818
    var v5 uint = (v3 % (((v3 % ((773) | 1))) | 1))
    var v6 int8 = 42
    var v7 float64 = (1.5 - (0.0 - (((1.5 - 1.5) + 1e10) - 3.14159)))
    for fi8 := 0; fi8 < 13; fi8++ {
        v6 = 100
        _ = fi8
    }
    v6 = v6
    v7 = 3.14159
    switch v4 {
    case 2147483647:
        if ((!(((v7 + 0.0) / (v7 / v7)) >= (v7 * (float64(1e10) + (v7 - v7))))) && false) {
            v9 := GenS1{F0: 1}
            _ = v9
        }
    case 1:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    case -208:
        if len(v2) > 0 {
            v2[0] = 264
        }
    default:
        for fi10 := 0; fi10 < 2; fi10++ {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            v11 := GenS1{F0: 1}
            var v12 float32 = (3.14159 / 3.14159)
            _ = fi10
            _ = v11
            _ = v12
        }
    }
    v3 = v5
    if (false && (v6 > (v6 & v6))) {
        var v13 float32 = ((3.14159 + (((0.0 * ((3.14159 * -2.25) / (1.5 + 1.5))) / 0.1) * 1.5)) - 1e10)
        _ = v13
    }
    v1 += 2147483648
    v4 = v4
    for fi14 := 0; fi14 < 15; fi14++ {
        if (!false) {
            v15 := GenS1{F0: -32768}
            _ = v15
        } else {
            v16 := GenS2{F0: (v4 * fi14)}
            for fi17 := 0; fi17 < 1; fi17++ {
                v18 := GenS2{F0: v4}
                v2 = append(v2, 2147483647)
                _ = fi17
                _ = v18
            }
            for fi19 := 0; fi19 < 15; fi19++ {
                var v20 uint32 = v1
                if ((v6 ^ (v6 % (((v6 >> ((v3) & 63))) | 1))) < (^(v6))) {
                    var v21 uint8 = 0
                    _ = v21
                }
                var v22 uint32 = uint32(v3)
                switch fi19 {
                case -1:
                    var v23 string = "foo bar"
                    _ = v23
                case -545:
                    if (v6 > (v6 + (v6 - v6))) {
                        _ = len(v2)
                    } else {
                        if false {
                            switch fi19 {
                            case 128:
                                switch "" {
                                case "foo bar":
                                    fmt.Println(v3, fi14, v1)
                                    var v24 float64 = v7
                                    fmt.Println(v4)
                                    _ = v24
                                case "hello":
                                    var v25 int = (fi14 * (v4 >> (((v5 % ((v5) | 1))) & 63)))
                                    _ = v25
                                case "":
                                    v22 = (v20 - 42)
                                }
                            }
                            if len(v2) > 0 {
                                v2[0] = 65535
                            }
                            if (!(v20 != (v20 % ((0) | 1)))) {
                                v1 = v1
                                var v26 bool = true
                                _ = v26
                            }
                        } else {
                            var v27 uint = uint(v20)
                            v20 -= (v20 + v22)
                            _ = v27
                        }
                        switch ("" + "foo bar") {
                        case "hello":
                            v22 = (^(v22))
                        case "":
                            var v28 uint32 = (v20 & v20)
                            _ = v28
                        }
                    }
                case 42:
                    v29 := GenS3{F0: ((((float32(v7) * 1e10) - 3.14159) > ((0.0 - 0.0) + (((float32(v7) / 1.5) * 3.14159) / float32(((v7 * ((0.0 / v7) / (float64(float32(((v7 * v7) + 0.0))) + v7))) - -2.25))))) && true)}
                    _ = v29
                default:
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
                var v30 int = (^(v4))
                v4 -= fi19
                _ = fi19
                _ = v20
                _ = v22
                _ = v30
            }
            _ = v16
        }
        v7 = v7
        _ = fi14
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
