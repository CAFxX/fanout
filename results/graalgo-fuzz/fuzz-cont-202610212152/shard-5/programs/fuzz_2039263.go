package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 uint }
type GenS2 struct { F0 int64; F1 bool; F2 uint64 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int8 }
func genFn0(p0 bool, p1 uint16) bool {
    v4 := GenS1{F0: 2, F1: 552, F2: uint(p1)}
    var v5 int8 = -128
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return (0 < 2147483647)
}
func genFn1(p0 uint, p1 uint16) int32 {
    var v6 uint16 = 107
    if (!false) {
        var v7 string = "hello"
        switch 127 {
        case 719:
            if false {
                for fi8 := 0; fi8 < 14; fi8++ {
                    v9 := GenS1{F0: 128, F1: 779, F2: (p0 + 18446744073709551615)}
                    if (42 != 1000) {
                        v10 := GenS2{F0: 478, F1: (!false), F2: 602}
                        _ = v10
                    } else {
                        v11 := GenS1{F0: 2, F1: 2147483647, F2: (^(p0))}
                        _ = v11
                    }
                    _ = fi8
                    _ = v9
                }
                for fi12 := 0; fi12 < 5; fi12++ {
                    var v13 bool = true
                    _ = fi12
                    _ = v13
                }
            }
        case -1:
            if (0 != -1) {
                p1 = (p1 >> ((p0) & 63))
                for fi14 := 0; fi14 < 9; fi14++ {
                    v15 := GenS3{F0: 42}
                    _ = fi14
                    _ = v15
                }
            } else {
                if (-448 == -2147483648) {
                    var v16 string = ((v7 + ("\n\t" + "\n\t")) + "\n\t")
                    _ = v16
                }
                switch 215 {
                case 1:
                    v17 := GenS1{F0: 255, F1: 199, F2: 65535}
                    _ = v17
                case -9223372036854775808:
                    p1 = 564
                case 128:
                    var v18 uint16 = p1
                    _ = v18
                default:
                    v7 = "a"
                }
                var v19 string = ("" + (v7 + ((("\n\t" + v7) + "hello") + v7)))
                _ = v19
            }
        case 0:
            var v20 int = 0
            _ = v20
        }
        _ = v7
    }
    _ = p0
    _ = p1
    _ = v6
    return -2147483648
}

func main() {
    v1 := make(map[uint32]string)
    v2 := []uint32{1, 0}
    var v3 float32 = ((-2.25 + -2.25) / 0.1)
    delete(v1, 23)
    v21 := GenS4{F0: -128}
    if (0 < 4294967295) {
        switch "\n\t" {
        case "":
            var v22 int8 = 127
            _ = v22
        case "hello":
            v23 := v1[969]
            _ = v23
        }
    } else {
        var v24 uint8 = 2
        for fi25 := 0; fi25 < 8; fi25++ {
            switch fi25 {
            case 32767:
                var v26 uint16 = 716
                _ = v26
            case -2147483648:
                _ = len(v1)
            default:
                for fi27 := 0; fi27 < 18; fi27++ {
                    var v28 int32 = -2147483648
                    fmt.Println(v24)
                    v24 += 148
                    _ = fi27
                    _ = v28
                }
            }
            _ = fi25
        }
        fmt.Println(v24, v3)
        _ = v24
    }
    v29 := v1[0]
    var v30 uint64 = 42
    var v31 uint32 = 2147483648
    if (255 > 1) {
        var v32 uint64 = v30
        v2 = append(v2, 249)
        fmt.Println(v31)
        _ = v32
    }
    if true {
        v1[(v31 | 200)] = ""
        for fi33 := 0; fi33 < 4; fi33++ {
            var v34 uint32 = v31
            _ = fi33
            _ = v34
        }
    }
    v31 = (^(v31))
    var v35 float32 = v3
    v36 := GenS1{F0: 42, F1: 2147483647, F2: 330}
    fmt.Println(v29, v30, v31)
    _ = len(v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v2
    _ = v21
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v35
    _ = v36
}
