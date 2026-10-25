package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint; F1 bool; F2 int64 }

func main() {
    var v1 int32 = 0
    var v2 float64 = 1e10
    var v3 uint8 = 255
    v4 := make(map[int]string)
    var v5 int32 = -2147483648
    delete(v4, 486)
    v5 += (v5 << ((4294967295) & 63))
    fmt.Println(v5)
    var v6 uint64 = 1
    v3 += v3
    var v7 uint32 = 2147483648
    v8 := []int{128}
    v1 -= (v1 % ((v1) | 1))
    fmt.Println(v5, v6)
    v7 -= 1
    v7 += (v7 ^ (v7 + uint32(v6)))
    fmt.Println(v2, v6, v1)
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    for fi9 := 0; fi9 < 4; fi9++ {
        if true {
            switch ("hello" + "foo bar") {
            case "foo bar":
                v10 := v4[fi9]
                _ = v10
            case "a":
                var v11 int32 = (v5 ^ (^(v5)))
                _ = v11
            default:
                v1 += (v1 * v5)
            }
            if false {
                var v12 int = fi9
                v13 := GenS1{F0: 0}
                var v14 int32 = -698
                _ = v12
                _ = v13
                _ = v14
            }
            v2 = v2
        }
        var v15 bool = (391 == 2)
        var v16 float32 = -2.25
        _ = fi9
        _ = v15
        _ = v16
    }
    for fi17 := 0; fi17 < 3; fi17++ {
        v5 = (v1 | v5)
        for fi18 := 0; fi18 < 14; fi18++ {
            if true {
                if (42 < uint16(v7)) {
                    switch "foo bar" {
                    case "hello":
                        var v19 uint = 65535
                        _ = v19
                    case "\n\t":
                        switch ("" + ("" + "foo bar")) {
                        case "a":
                            var v20 int32 = 1
                            _ = v20
                        case "":
                            var v21 int32 = (v5 * v5)
                            _ = v21
                        case "foo bar":
                            switch "foo bar" {
                            case "foo bar":
                                v22 := GenS2{F0: 0, F1: (v3 >= 41), F2: -9223372036854775808}
                                _ = v22
                            case "hello":
                                v3 = v3
                            case "\n\t":
                                v7 = 4294967295
                            default:
                                delete(v4, fi18)
                            }
                        default:
                            var v23 uint64 = (v6 << ((580) & 63))
                            _ = v23
                        }
                    default:
                        _ = len(v4)
                        v24 := GenS1{F0: -2147483648}
                        _ = v24
                    }
                }
                if len(v8) > 0 {
                    v8[0] = 42
                }
                v7 = (v7 + uint32(v3))
                v25 := v4[fi17]
                var v26 float32 = ((3.14159 + 0.1) + ((-2.25 + 3.14159) * -2.25))
                _ = v25
                _ = v26
            } else {
                switch "foo bar" {
                case "\n\t":
                    var v27 uint8 = (^(v3))
                    _ = v27
                case "":
                    var v28 float32 = -2.25
                    _ = v28
                case "hello":
                    v29 := v4[(fi17 * fi18)]
                    _ = v29
                default:
                    v4[(fi17 << ((0) & 63))] = "a"
                }
            }
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
            fmt.Println(fi17)
            var v30 int8 = 0
            _ = fi18
            _ = v30
        }
        _ = fi17
    }
    for fi31 := 0; fi31 < 20; fi31++ {
        v3 = 42
        var v32 uint64 = 4294967295
        var v33 uint32 = (v7 % ((2147483647) | 1))
        _ = fi31
        _ = v32
        _ = v33
    }
    var v34 uint8 = v3
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
