package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }
type GenS2 struct { F0 uint16; F1 uint64; F2 bool }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int16 }

func main() {
    var v1 int8 = 100
    v2 := []uint32{42, 520, 2147483647, 2147483647}
    var v3 uint16 = 2
    var v4 uint16 = (v3 & v3)
    v1 -= (-(v1))
    if false {
        if true {
            v3 = v3
            var v5 int16 = 347
            if true {
                v6 := GenS1{F0: 42, F1: int64(v4)}
                _ = v6
            }
            _ = v5
        } else {
            var v7 uint64 = 0
            if false {
                switch "" {
                case "\n\t":
                    if false {
                        var v8 int16 = -1
                        switch ("foo bar" + "") {
                        case "":
                            var v9 string = "\n\t"
                            _ = v9
                        default:
                            for fi10 := 0; fi10 < 12; fi10++ {
                                fmt.Println(v8)
                                _ = fi10
                            }
                        }
                        _ = v8
                    } else {
                        v11 := GenS1{F0: 4294967295, F1: 9223372036854775807}
                        _ = v11
                    }
                default:
                    for fi12 := 0; fi12 < 19; fi12++ {
                        if false {
                            switch "\n\t" {
                            case "hello":
                                for fi13 := 0; fi13 < 12; fi13++ {
                                    v14 := GenS1{F0: 42, F1: -9223372036854775808}
                                    _ = fi13
                                    _ = v14
                                }
                            case "\n\t":
                                fmt.Println(v4, v3, fi12)
                            default:
                                var v15 int32 = 2147483647
                                _ = v15
                            }
                        }
                        var v16 int32 = 42
                        if ((!(((float64(0.0) + -2.25) / 1.0) <= 1e10)) && (2 <= 2147483648)) {
                            var v17 float32 = 0.1
                            if false {
                                v3 = v3
                            } else {
                                if len(v2) > 0 {
                                    v2[0] = 65535
                                }
                            }
                            _ = v17
                        }
                        _ = fi12
                        _ = v16
                    }
                }
            } else {
                v18 := GenS2{F0: (v3 | v4), F1: v7, F2: false}
                for fi19 := 0; fi19 < 13; fi19++ {
                    var v20 uint16 = (^(v4))
                    _ = fi19
                    _ = v20
                }
                _ = len(v2)
                _ = v18
            }
            v4 = 0
            _ = v7
        }
    }
    v3 -= v3
    v21 := GenS3{F0: -2147483648}
    if true {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        var v22 int64 = 0
        _ = v22
    }
    var v23 float64 = (float64(3.14159) * float64(((((3.14159 / 1.5) / -2.25) + (3.14159 - -2.25)) + -2.25)))
    v24 := GenS4{F0: 163}
    v4 = v3
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}
