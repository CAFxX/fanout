package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool; F2 uint32 }
type GenS2 struct { F0 int16; F1 int }
type GenS3 struct { F0 bool; F1 uint8; F2 uint32 }
type GenS4 struct { F0 uint64; F1 uint16; F2 uint16 }
func genFn0(p0 uint16, p1 int64, p2 int32) int8 {
    v5 := []int{32768, -1}
    var v6 int64 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 36
}

func main() {
    var v1 int8 = 100
    v2 := []uint32{}
    v3 := []uint32{1, 2, 779, 2147483647}
    v4 := make(map[uint32]bool)
    var v7 int16 = int16(v1)
    for fi8 := 0; fi8 < 2; fi8++ {
        v9 := GenS1{F0: -61, F1: (0.1 > 0.1), F2: 4294967295}
        _ = fi8
        _ = v9
    }
    delete(v4, 65535)
    for fi10 := 0; fi10 < 20; fi10++ {
        var v11 int = 344
        _ = len(v3)
        if (((("foo bar" + "") + ("a" + "a")) + "hello") > "foo bar") {
            v12 := v4[4294967295]
            var v13 int64 = -1
            v4[582] = v12
            _ = v12
            _ = v13
        }
        _ = fi10
        _ = v11
    }
    for fi14 := 0; fi14 < 5; fi14++ {
        v2 = append(v2, 847)
        _ = fi14
    }
    if true {
        fmt.Println(v1, v1)
    }
    v7 -= (v7 | -1)
    v7 += (^(v7))
    if true {
        switch ("a" + "a") {
        case "\n\t":
            v1 = (^(v1))
        default:
            switch "hello" {
            case "a":
                if ((^(v1)) <= 127) {
                    switch 421 {
                    case 42:
                        v7 = 695
                    case 469:
                        var v15 uint32 = 42
                        _ = v15
                    case 32767:
                        for fi16 := 0; fi16 < 20; fi16++ {
                            var v17 float64 = 1e10
                            _ = fi16
                            _ = v17
                        }
                    }
                    var v18 float32 = (1e10 - (1.5 * 0.1))
                    _ = v18
                }
            case "hello":
                if false {
                    var v19 bool = true
                    for fi20 := 0; fi20 < 15; fi20++ {
                        switch "\n\t" {
                        case "hello":
                            v21 := GenS1{F0: -1, F1: (944 > (^(fi20))), F2: 2147483647}
                            _ = v21
                        case "foo bar":
                            var v22 float32 = -2.25
                            _ = v22
                        case "a":
                            v19 = v19
                        default:
                            v3 = append(v3, 42)
                        }
                        v19 = true
                        _ = fi20
                    }
                    _ = v19
                }
            }
        }
    } else {
        var v23 bool = (!false)
        var v24 int = 2147483647
        _ = v23
        _ = v24
    }
    for fi25 := 0; fi25 < 1; fi25++ {
        for fi26 := 0; fi26 < 4; fi26++ {
            var v27 bool = true
            if v27 {
                if (42 < 549) {
                    if len(v2) > 0 {
                        v2[0] = 42
                    }
                    v28 := GenS2{F0: -1, F1: fi25}
                    v29 := GenS3{F0: v27, F1: 131, F2: 578}
                    _ = v28
                    _ = v29
                } else {
                    v30 := v4[0]
                    var v31 uint32 = 922
                    var v32 int64 = 1
                    var v33 float32 = 0.0
                    v27 = (!(32767 >= 2))
                    _ = v30
                    _ = v31
                    _ = v32
                    _ = v33
                }
                switch fi25 {
                case 0:
                    v1 = (v1 + v1)
                case 9223372036854775807:
                    v34 := GenS4{F0: 952, F1: 2, F2: 2}
                    _ = v34
                default:
                    var v35 bool = v27
                    _ = v35
                }
            } else {
                switch (fi25 % ((fi26) | 1)) {
                case 32768:
                    if (397 >= 42) {
                        var v36 uint32 = 2147483648
                        _ = v36
                    } else {
                        fmt.Println(v7)
                        v27 = (-529 <= fi26)
                    }
                case 0:
                    if true {
                        v1 = v1
                        v7 = (v7 << ((226) & 63))
                    }
                case 128:
                    v7 = v7
                }
                var v37 int16 = v7
                _ = v37
            }
            var v38 float32 = 3.14159
            _ = fi26
            _ = v27
            _ = v38
        }
        if false {
            var v39 uint16 = 0
            v7 -= int16(v1)
            v4[1] = true
            _ = v39
        }
        _ = fi25
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
