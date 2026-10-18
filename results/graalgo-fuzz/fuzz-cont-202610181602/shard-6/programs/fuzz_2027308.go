package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool; F2 int }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 bool; F1 bool }
type GenS4 struct { F0 uint16 }

func main() {
    var v1 float32 = -2.25
    var v2 uint64 = 4294967295
    v3 := []uint32{276, 42}
    var v4 float32 = v1
    var v5 uint8 = 0
    v6 := make(map[int]string)
    for fi7 := 0; fi7 < 16; fi7++ {
        fmt.Println(v2, v5)
        _ = fi7
    }
    fmt.Println(v2, v5, v2)
    v3 = append(v3, 0)
    v8 := GenS1{F0: 4294967295, F1: (false && (100 > 42)), F2: 42}
    if (32768 <= 127) {
        fmt.Println(v2)
    }
    v6[32768] = "foo bar"
    for fi9 := 0; fi9 < 10; fi9++ {
        var v10 int16 = 804
        var v11 uint8 = 173
        var v12 bool = true
        _ = fi9
        _ = v10
        _ = v11
        _ = v12
    }
    if (0 == 2147483647) {
        var v13 uint8 = v5
        _ = v13
    }
    for fi14 := 0; fi14 < 8; fi14++ {
        if (381 == uint16(v5)) {
            v2 += (v2 / (((v2 * (v2 % ((v2) | 1)))) | 1))
        } else {
            if (1 >= 42) {
                v15 := GenS1{F0: 4294967295, F1: (-716 <= 1), F2: fi14}
                switch (("\n\t" + "foo bar") + "") {
                case "\n\t":
                    if true {
                        _ = len(v6)
                        v2 = uint64(v5)
                    }
                case "":
                    var v16 int8 = 42
                    _ = v16
                case "foo bar":
                    v17 := GenS2{F0: -118}
                    _ = v17
                default:
                    if false {
                        for fi18 := 0; fi18 < 3; fi18++ {
                            v19 := GenS3{F0: false, F1: true}
                            _ = fi18
                            _ = v19
                        }
                        for fi20 := 0; fi20 < 8; fi20++ {
                            v21 := GenS1{F0: 0, F1: true, F2: fi20}
                            var v22 uint8 = 42
                            if (false && (!(v4 != v4))) {
                                v23 := GenS4{F0: 32767}
                                _ = v23
                            }
                            _ = fi20
                            _ = v21
                            _ = v22
                        }
                    }
                }
                _ = v15
            } else {
                v24 := GenS4{F0: 790}
                _ = v24
            }
        }
        _ = fi14
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
