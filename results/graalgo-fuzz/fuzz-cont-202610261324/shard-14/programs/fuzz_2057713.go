package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 uint16 }
type GenS2 struct { F0 int16; F1 uint32; F2 uint16 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int8; F1 uint32 }
func genFn0(p0 bool, p1 int8) uint16 {
    var v7 float32 = ((0.1 - 1e10) / ((0.1 / 1.0) + (float32(1.5) - 0.1)))
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = v7
    return 518
}

func main() {
    var v1 int64 = -9223372036854775808
    var v2 string = "a"
    v3 := []string{v2, v2}
    v4 := make(map[uint32]bool)
    var v5 uint32 = 65535
    v6 := []uint32{1000, (v5 ^ 339)}
    if ((128 < 84) || false) {
        for fi8 := 0; fi8 < 16; fi8++ {
            var v9 uint = 65535
            if (v2 <= v2) {
                switch (((((v2 + v2) + "a") + (v2 + "hello")) + v2) + v2) {
                case "foo bar":
                    v9 -= (v9 * v9)
                }
                var v10 uint16 = 32767
                _ = v10
            }
            _ = fi8
            _ = v9
        }
        fmt.Println(v1, v1, v1)
    } else {
        v11 := GenS1{F0: 65535, F1: 4294967295, F2: 42}
        _ = v11
    }
    var v12 uint32 = (^(v5))
    switch -637 {
    case -9223372036854775808:
        for fi13 := 0; fi13 < 20; fi13++ {
            switch (fi13 / ((fi13) | 1)) {
            case -2147483648:
                var v14 int = (fi13 % ((-411) | 1))
                _ = v14
            case 2147483647:
                v15 := GenS2{F0: 1, F1: 286, F2: 255}
                _ = v15
            case -1:
                v1 = (-(v1))
            default:
                if (32767 == 0) {
                    for fi16 := 0; fi16 < 7; fi16++ {
                        switch fi13 {
                        case 9223372036854775807:
                            switch fi16 {
                            case 32767:
                                var v17 uint8 = 127
                                _ = v17
                            default:
                                switch "hello" {
                                case "":
                                    v18 := v4[42]
                                    _ = v18
                                case "hello":
                                    var v19 bool = true
                                    _ = v19
                                case "foo bar":
                                    var v20 uint32 = v12
                                    var v21 int16 = 32767
                                    _ = v20
                                    _ = v21
                                default:
                                    var v22 int8 = 0
                                    _ = v22
                                }
                            }
                        case 42:
                            if (("hello" + v2) != v2) {
                                v23 := GenS3{F0: fi13}
                                _ = v23
                            }
                        default:
                            var v24 int = fi13
                            _ = v24
                        }
                        _ = len(v4)
                        switch ("foo bar" + "") {
                        case "":
                            fmt.Println(v12, v1)
                        case "a":
                            var v25 float32 = float32(((float64(1.5) * 3.14159) - 3.14159))
                            _ = v25
                        }
                        switch ("\n\t" + (v2 + v2)) {
                        case "":
                            v12 -= 346
                        }
                        _ = fi16
                    }
                    v26 := GenS4{F0: 127}
                    var v27 uint16 = 0
                    _ = v26
                    _ = v27
                }
            }
            _ = fi13
        }
    default:
        fmt.Println(v12)
    }
    v5 = (v5 / (((^(v12))) | 1))
    v28 := GenS3{F0: 128}
    fmt.Println(v2, v2, v5)
    var v29 uint32 = 833
    v30 := GenS1{F0: 2, F1: -9223372036854775808, F2: 530}
    var v31 float32 = (0.0 - -2.25)
    for fi32 := 0; fi32 < 19; fi32++ {
        v33 := GenS5{F0: -128, F1: (v5 / (((v12 * 0)) | 1))}
        if len(v3) > 0 {
            v3[0] = v2
        }
        var v34 uint8 = 255
        _ = fi32
        _ = v33
        _ = v34
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
