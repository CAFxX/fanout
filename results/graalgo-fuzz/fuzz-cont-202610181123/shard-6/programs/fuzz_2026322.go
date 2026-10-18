package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }

func main() {
    var v1 string = ""
    var v2 uint64 = 22
    var v3 float64 = 3.14159
    v4 := make(map[string]uint32)
    v5 := []int{4294967295, 127}
    var v6 uint64 = 2
    _ = len(v5)
    if ((v3 != 0.1) && false) {
        v7 := GenS1{F0: -647, F1: 1}
        var v8 uint64 = 738
        _ = v7
        _ = v8
    } else {
        var v9 uint64 = (v2 | v6)
        _ = v9
    }
    v5 = append(v5, -948)
    v1 = ("" + ("" + v1))
    if true {
        if (9223372036854775807 != 127) {
            for fi10 := 0; fi10 < 3; fi10++ {
                var v11 float64 = 3.14159
                _ = fi10
                _ = v11
            }
        }
        v6 -= v6
    }
    _ = len(v4)
    for fi12 := 0; fi12 < 4; fi12++ {
        for fi13 := 0; fi13 < 6; fi13++ {
            switch fi13 {
            case 9223372036854775807:
                var v14 uint8 = 56
                _ = v14
            }
            _ = fi13
        }
        for fi15 := 0; fi15 < 2; fi15++ {
            v3 -= v3
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = fi15
        }
        for fi16 := 0; fi16 < 20; fi16++ {
            var v17 int32 = 0
            if ((-32768 > 42) || (42 == 87)) {
                v18 := GenS1{F0: (v17 % ((30) | 1)), F1: 73}
                switch (fi12 + 128) {
                case -1:
                    var v19 float64 = 1e10
                    _ = v19
                case 855:
                    if false {
                        v5 = append(v5, fi16)
                        if (0 > 258) {
                            v20 := GenS1{F0: (v17 | (v17 % ((-1) | 1))), F1: 42}
                            fmt.Println(v1, fi16, fi16)
                            _ = v20
                        } else {
                            if (!false) {
                                v6 -= (v2 ^ v2)
                                if len(v5) > 0 {
                                    v5[0] = (fi16 << ((4294967295) & 63))
                                }
                                var v21 float32 = (float32((v3 * v3)) - ((float32(((0.0 + v3) * v3)) - float32(v3)) + 1.5))
                                _ = v21
                            } else {
                                var v22 uint16 = 32767
                                v2 += v2
                                _ = v22
                            }
                            v2 += v6
                            var v23 int16 = 1000
                            var v24 uint = 0
                            _ = v23
                            _ = v24
                        }
                        delete(v4, (((v1 + v1) + (v1 + ((v1 + "hello") + (v1 + "foo bar")))) + ""))
                    }
                case 127:
                    v1 = "hello"
                    var v25 bool = (true || ((2147483647 < 692) && false))
                    _ = v25
                default:
                    if true {
                        v2 = (v2 << ((65535) & 63))
                    } else {
                        var v26 uint16 = 32767
                        v27 := GenS1{F0: 2147483647, F1: 127}
                        _ = v26
                        _ = v27
                    }
                }
                if (65535 > 65535) {
                    if false {
                        var v28 uint8 = 255
                        v2 = v6
                        fmt.Println(v17, v3, v28)
                        if false {
                            v29 := GenS1{F0: (v17 - v17), F1: 42}
                            var v30 string = ""
                            var v31 uint64 = 988
                            _ = v29
                            _ = v30
                            _ = v31
                        }
                        _ = v28
                    } else {
                        var v32 uint16 = 32767
                        _ = v32
                    }
                    fmt.Println(fi12, fi16, v17)
                }
                v33 := GenS1{F0: v17, F1: 100}
                _ = v18
                _ = v33
            } else {
                var v34 float32 = float32(1.5)
                _ = v34
            }
            _ = fi16
            _ = v17
        }
        _ = fi12
    }
    v6 -= (v6 * (v2 * 433))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
