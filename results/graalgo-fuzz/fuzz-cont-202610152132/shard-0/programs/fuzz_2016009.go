package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int64; F1 int32 }

func main() {
    var v1 int16 = 32767
    v2 := []uint32{42, 65535, 873, 863}
    var v3 uint8 = 42
    v4 := make(map[uint32]uint32)
    var v5 float32 = 3.14159
    v6 := GenS1{F0: 0}
    v1 = v1
    fmt.Println(v1, v5)
    if (!false) {
        fmt.Println(v5)
    } else {
        if (772 != 753) {
            var v7 uint32 = 65535
            _ = v7
        } else {
            _ = len(v2)
            var v8 int64 = -9223372036854775808
            for fi9 := 0; fi9 < 9; fi9++ {
                if (false || false) {
                    for fi10 := 0; fi10 < 13; fi10++ {
                        fmt.Println(v1, fi10)
                        _ = fi10
                    }
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    var v11 bool = (!true)
                    _ = v11
                }
                var v12 uint = 65535
                _ = fi9
                _ = v12
            }
            _ = v8
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    v5 -= v5
    v5 -= (v5 - 0.0)
    v5 -= ((-2.25 / float32(1.5)) / v5)
    v3 -= v3
    fmt.Println(v1, v5, v3)
    fmt.Println(v3, v1, v5)
    if false {
        if (!false) {
            _ = len(v4)
        } else {
            if true {
                v1 -= v1
            } else {
                for fi13 := 0; fi13 < 15; fi13++ {
                    v5 = v5
                    v14 := GenS1{F0: -100}
                    v15 := v4[65535]
                    _ = fi13
                    _ = v14
                    _ = v15
                }
            }
        }
        v16 := GenS2{F0: 239, F1: 1}
        v5 += 0.1
        _ = v16
    }
    if (!(true || (762 > 4294967295))) {
        for fi17 := 0; fi17 < 20; fi17++ {
            for fi18 := 0; fi18 < 9; fi18++ {
                v5 = (((((float32(1.5) * v5) + v5) + 1e10) - v5) + v5)
                _ = fi18
            }
            _ = fi17
        }
    } else {
        v19 := GenS1{F0: -9223372036854775808}
        v5 = (float32(3.14159) * 1e10)
        for fi20 := 0; fi20 < 13; fi20++ {
            delete(v4, 2)
            if (916 <= 42) {
                fmt.Println(v5, v5, v5)
            } else {
                if false {
                    if false {
                        switch "hello" {
                        case "\n\t":
                            v3 = (v3 + (v3 * v3))
                        case "hello":
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        default:
                            fmt.Println(fi20)
                        }
                        var v21 int8 = 0
                        _ = v21
                    }
                    fmt.Println(v5)
                } else {
                    for fi22 := 0; fi22 < 13; fi22++ {
                        var v23 float64 = 1.5
                        v5 = 1e10
                        var v24 bool = (65535 > 42)
                        _ = fi22
                        _ = v23
                        _ = v24
                    }
                    fmt.Println(v5)
                    _ = len(v2)
                }
            }
            _ = fi20
        }
        _ = v19
    }
    for fi25 := 0; fi25 < 12; fi25++ {
        var v26 int16 = (v1 + int16(fi25))
        v27 := GenS2{F0: 0, F1: -131}
        _ = fi25
        _ = v26
        _ = v27
    }
    var v28 float64 = 0.1
    v3 = v3
    fmt.Println(v5, v5)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
