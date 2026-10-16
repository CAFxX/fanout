package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 int8 }
type GenS2 struct { F0 uint32; F1 string; F2 uint32 }

func main() {
    var v1 uint = 859
    var v2 float32 = 1.5
    var v3 uint64 = 984
    var v4 uint16 = 32767
    if true {
        fmt.Println(v1, v2)
        v4 -= 42
        v2 = v2
    }
    if ((!true) && (1e10 >= (float64((v2 + (3.14159 + ((((v2 - v2) * 3.14159) / v2) / v2)))) - (-2.25 + float64((v2 / (-2.25 + (0.1 / 1e10)))))))) {
        if true {
            fmt.Println(v4, v1, v2)
            switch 42 {
            case 2147483647:
                var v5 uint = 4294967295
                _ = v5
            default:
                v1 += 401
            }
            if true {
                fmt.Println(v4, v3)
                for fi6 := 0; fi6 < 16; fi6++ {
                    v7 := []int{}
                    v1 = v1
                    _ = fi6
                    _ = v7
                }
            } else {
                v2 = v2
                var v8 uint64 = uint64(v1)
                _ = v8
            }
        } else {
            if true {
                v1 = 439
                for fi9 := 0; fi9 < 3; fi9++ {
                    if (-2147483648 <= 2147483647) {
                        if (true && (42 < 0)) {
                            var v10 uint16 = 32767
                            v11 := []int{-9223372036854775808}
                            _ = v10
                            _ = v11
                        } else {
                            for fi12 := 0; fi12 < 12; fi12++ {
                                v13 := GenS1{F0: -1, F1: -1, F2: 127}
                                _ = fi12
                                _ = v13
                            }
                            for fi14 := 0; fi14 < 15; fi14++ {
                                var v15 uint8 = 42
                                var v16 int32 = 2147483647
                                _ = fi14
                                _ = v15
                                _ = v16
                            }
                            var v17 uint16 = v4
                            _ = v17
                        }
                        fmt.Println(v3, v3, v1)
                        switch fi9 {
                        case 170:
                            for fi18 := 0; fi18 < 2; fi18++ {
                                var v19 int32 = -657
                                _ = fi18
                                _ = v19
                            }
                        default:
                            switch -498 {
                            case 9223372036854775807:
                                var v20 uint64 = 2
                                _ = v20
                            case -1:
                                v21 := []int{fi9}
                                _ = v21
                            case 42:
                                switch -9223372036854775808 {
                                case -482:
                                    var v22 uint64 = v3
                                    _ = v22
                                case 4294967295:
                                    var v23 bool = false
                                    _ = v23
                                }
                            }
                        }
                    }
                    _ = fi9
                }
                fmt.Println(v2)
            } else {
                v3 -= (v3 / ((v3) | 1))
            }
            v4 -= 0
        }
    }
    if (653 > -2147483648) {
        v1 = (^(v1))
        v1 += (v1 / ((301) | 1))
    } else {
        v24 := []int{}
        _ = v24
    }
    if (0 != 2147483647) {
        var v25 float32 = 0.0
        if false {
            v4 -= v4
            for fi26 := 0; fi26 < 4; fi26++ {
                var v27 int = fi26
                v25 -= v25
                _ = fi26
                _ = v27
            }
            var v28 uint64 = 0
            _ = v28
        }
        _ = v25
    }
    if (2 <= v1) {
        var v29 int8 = -19
        v3 = (v3 ^ (^(v3)))
        _ = v29
    } else {
        var v30 uint8 = 42
        v1 = (v1 / (((v1 | v1)) | 1))
        var v31 uint16 = v4
        _ = v30
        _ = v31
    }
    v32 := GenS2{F0: 2147483648, F1: ((("hello" + "") + "hello") + "foo bar"), F2: 2147483648}
    var v33 uint16 = (v4 - (v4 - (v4 & v4)))
    v4 += 348
    v34 := GenS1{F0: 42, F1: 42, F2: 70}
    switch 32768 {
    case 128:
        var v35 uint64 = (v3 & (v3 / ((v3) | 1)))
        _ = v35
    default:
        for fi36 := 0; fi36 < 5; fi36++ {
            v2 -= -2.25
            _ = fi36
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
}
