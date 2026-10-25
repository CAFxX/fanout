package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint16; F1 int8 }

func main() {
    v1 := make(map[int]uint32)
    var v2 int8 = 100
    var v3 bool = ("foo bar" > "a")
    v4 := []uint32{451, 951}
    v5 := make(map[int]uint32)
    _ = len(v1)
    v1[-639] = 0
    v3 = v3
    v6 := GenS1{F0: 18446744073709551615}
    v7 := GenS1{F0: 2}
    var v8 int8 = -1
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v9 := GenS1{F0: 42}
    for fi10 := 0; fi10 < 9; fi10++ {
        var v11 int8 = v8
        _ = fi10
        _ = v11
    }
    for fi12 := 0; fi12 < 14; fi12++ {
        v13 := GenS2{F0: 220, F1: (v8 ^ -1)}
        _ = fi12
        _ = v13
    }
    for fi14 := 0; fi14 < 15; fi14++ {
        v2 = v8
        v15 := GenS2{F0: 573, F1: 1}
        _ = fi14
        _ = v15
    }
    var v16 uint = 255
    fmt.Println(v16, v8)
    for fi17 := 0; fi17 < 3; fi17++ {
        for fi18 := 0; fi18 < 7; fi18++ {
            if (42 == 1) {
                var v19 int32 = 2147483647
                _ = v19
            } else {
                switch "\n\t" {
                case "foo bar":
                    var v20 uint64 = 4294967295
                    _ = v20
                case "a":
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    switch fi18 {
                    case 9223372036854775807:
                        _ = len(v5)
                    case 127:
                        fmt.Println(v2)
                    default:
                        if (0 != 0) {
                            v21 := GenS1{F0: 762}
                            switch "a" {
                            case "":
                                switch fi17 {
                                case 0:
                                    v3 = (v8 <= v8)
                                case -9223372036854775808:
                                    var v22 bool = true
                                    _ = v22
                                default:
                                    var v23 string = ((("" + "a") + "\n\t") + (("\n\t" + "") + ""))
                                    _ = v23
                                }
                            case "\n\t":
                                switch (-(fi17)) {
                                case 128:
                                    fmt.Println(v8)
                                case 779:
                                    v3 = (194 != 127)
                                case -1:
                                    v24 := GenS1{F0: 586}
                                    _ = v24
                                }
                            case "foo bar":
                                var v25 int32 = 2147483647
                                _ = v25
                            }
                            _ = v21
                        } else {
                            fmt.Println(v8, v16, v8)
                            if false {
                                var v26 int16 = 0
                                var v27 uint8 = 0
                                _ = v26
                                _ = v27
                            }
                            var v28 uint = 0
                            _ = v28
                        }
                    }
                case "\n\t":
                    v2 -= v8
                }
                v4 = append(v4, 1)
                var v29 bool = v3
                _ = v29
            }
            _ = fi18
        }
        _ = fi17
    }
    v30 := GenS2{F0: 839, F1: v2}
    v31 := GenS2{F0: 0, F1: (v2 - v8)}
    _ = len(v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
