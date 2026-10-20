package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8 }
type GenS2 struct { F0 uint8; F1 uint16 }

func main() {
    var v1 float32 = 0.0
    var v2 uint = 42
    var v3 uint = 42
    var v4 int = 128
    if (255 != 195) {
        if false {
            if true {
                if false {
                    if (!false) {
                        v1 += v1
                        v3 = v2
                    }
                } else {
                    if false {
                        var v5 bool = (!true)
                        fmt.Println(v2)
                        _ = v5
                    }
                    var v6 int = (v4 | v4)
                    if false {
                        v7 := []int{(v6 + (v4 & v6)), v4}
                        switch "" {
                        case "":
                            if len(v7) > 0 {
                                v7[0] = (v6 | 0)
                            }
                        }
                        if (!true) {
                            v4 = (-(v6))
                        } else {
                            var v8 string = "foo bar"
                            _ = v8
                        }
                        _ = v7
                    } else {
                        for fi9 := 0; fi9 < 1; fi9++ {
                            for fi10 := 0; fi10 < 6; fi10++ {
                                var v11 int = (fi9 / ((32768) | 1))
                                v12 := []int{(v6 - 9223372036854775807), (fi9 ^ (-(fi9))), -1, (-(fi10))}
                                _ = fi10
                                _ = v11
                                _ = v12
                            }
                            v13 := []int{v4, (v4 - -9223372036854775808), v4, 0}
                            _ = fi9
                            _ = v13
                        }
                        var v14 int8 = -34
                        fmt.Println(v2)
                        _ = v14
                    }
                    _ = v6
                }
                for fi15 := 0; fi15 < 7; fi15++ {
                    v16 := []int{(^(v4)), (fi15 + v4), (v4 << (((v3 * v2)) & 63)), v4}
                    _ = fi15
                    _ = v16
                }
                fmt.Println(v1)
            } else {
                for fi17 := 0; fi17 < 20; fi17++ {
                    v3 = (v2 * (v2 + (v3 - v3)))
                    _ = fi17
                }
                fmt.Println(v2)
            }
            v4 += -2147483648
            var v18 int16 = 32767
            _ = v18
        } else {
            v3 = v2
            switch v4 {
            case -188:
                var v19 int16 = -1
                _ = v19
            case 42:
                v20 := GenS1{F0: true, F1: 241}
                _ = v20
            case -2147483648:
                switch "foo bar" {
                case "":
                    v4 -= (v4 - (v4 % ((-677) | 1)))
                default:
                    var v21 uint16 = 1
                    _ = v21
                }
            }
        }
    }
    for fi22 := 0; fi22 < 2; fi22++ {
        var v23 uint32 = 2
        v24 := GenS2{F0: 233, F1: 874}
        _ = fi22
        _ = v23
        _ = v24
    }
    v25 := GenS2{F0: 127, F1: 1}
    v1 = (3.14159 / (v1 / 1.0))
    var v26 int16 = 55
    fmt.Println(v2, v3)
    if false {
        if (v3 == v2) {
            for fi27 := 0; fi27 < 9; fi27++ {
                v28 := []int{fi27, (v4 | (-(fi27))), (fi27 | fi27), fi27}
                v29 := GenS2{F0: 0, F1: 0}
                _ = fi27
                _ = v28
                _ = v29
            }
            var v30 uint = 37
            _ = v30
        }
        v31 := []int{(v4 - (v4 ^ (v4 ^ v4))), int(v2), v4}
        _ = v31
    }
    if (true || false) {
        if false {
            v32 := GenS1{F0: false, F1: 0}
            _ = v32
        }
        var v33 int32 = 327
        _ = v33
    } else {
        v4 = v4
    }
    v26 = v26
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
}
