package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 uint32 }
type GenS2 struct { F0 int32; F1 int64 }
type GenS3 struct { F0 int32; F1 bool }

func main() {
    v1 := []uint32{993}
    var v2 float32 = 0.0
    var v3 uint = 18446744073709551615
    v4 := make(map[uint32]string)
    switch -2147483648 {
    case -211:
        switch -2147483648 {
        case 2147483647:
            v3 = 931
        case -429:
            v5 := GenS1{F0: (v3 + v3), F1: 42, F2: 2147483648}
            _ = v5
        default:
            if true {
                if (0 > 65535) {
                    v6 := GenS1{F0: 539, F1: 0, F2: 891}
                    if (("\n\t" + ((("hello" + ("hello" + "foo bar")) + "\n\t") + ("hello" + ("hello" + "\n\t")))) == "a") {
                        for fi7 := 0; fi7 < 4; fi7++ {
                            for fi8 := 0; fi8 < 7; fi8++ {
                                var v9 int64 = 468
                                v10 := GenS2{F0: 0, F1: v9}
                                var v11 int8 = 0
                                _ = fi8
                                _ = v10
                                _ = v11
                                _ = v9
                            }
                            for fi12 := 0; fi12 < 6; fi12++ {
                                v13 := GenS2{F0: -827, F1: 9223372036854775807}
                                var v14 int8 = -1
                                _ = fi12
                                _ = v13
                                _ = v14
                            }
                            var v15 int8 = 42
                            _ = fi7
                            _ = v15
                        }
                        v2 += (((v2 - (v2 * (0.0 * (0.1 / (v2 + (v2 * (float32(0.1) + -2.25))))))) / 0.1) - 3.14159)
                        if true {
                            var v16 uint64 = 102
                            v3 = 4294967295
                            _ = v16
                        }
                    } else {
                        if (42 > -2147483648) {
                            var v17 uint = (v3 << (((v3 + v3)) & 63))
                            var v18 int16 = -224
                            switch 127 {
                            case 128:
                                for fi19 := 0; fi19 < 6; fi19++ {
                                    switch fi19 {
                                    case 4294967295:
                                        if true {
                                            switch "foo bar" {
                                            case "hello":
                                                v2 = ((0.1 + v2) + v2)
                                            }
                                            _ = len(v4)
                                            v1 = append(v1, 65535)
                                        }
                                    case 32767:
                                        v3 -= 204
                                    case 128:
                                        var v20 float64 = 3.14159
                                        _ = v20
                                    }
                                    delete(v4, 957)
                                    _ = fi19
                                }
                            }
                            _ = v17
                            _ = v18
                        }
                    }
                    _ = v6
                }
                v21 := GenS3{F0: 2147483647, F1: true}
                fmt.Println(v3, v2)
                _ = v21
            } else {
                var v22 uint8 = 105
                v22 = 2
                _ = v22
            }
        }
    }
    v3 -= 0
    for fi23 := 0; fi23 < 3; fi23++ {
        v3 = (v3 << (((v3 / (((v3 ^ (^(v3)))) | 1))) & 63))
        _ = fi23
    }
    v2 = v2
    var v24 int64 = -9223372036854775808
    var v25 int64 = v24
    v25 = (v25 + v25)
    for fi26 := 0; fi26 < 1; fi26++ {
        var v27 float64 = (3.14159 / 1.0)
        _ = fi26
        _ = v27
    }
    fmt.Println(v25, v3)
    var v28 int = -9223372036854775808
    if false {
        var v29 bool = false
        var v30 string = ("hello" + "hello")
        _ = v29
        _ = v30
    } else {
        fmt.Println(v2)
        v28 += v28
    }
    var v31 int = v28
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v28
    _ = v3
    _ = v31
    _ = v4
}
