package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 uint8 }
type GenS2 struct { F0 int64; F1 uint32 }
type GenS3 struct { F0 uint16; F1 uint32 }
func genFn0(p0 uint8, p1 uint32, p2 uint8) int32 {
    v4 := GenS1{F0: p0, F1: 65535, F2: (p2 - (p2 | 1))}
    p1 -= (p1 << ((uint(p1)) & 63))
    p2 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return -152
}

func main() {
    var v1 bool = false
    v2 := make(map[uint32]uint32)
    v3 := make(map[uint32]bool)
    var v5 uint32 = 417
    switch -9223372036854775808 {
    case 32767:
        if v1 {
            switch -9223372036854775808 {
            case 9223372036854775807:
                var v6 int8 = -33
                _ = v6
            }
            v1 = (42 > 56)
            if v1 {
                v7 := GenS1{F0: 2, F1: 65535, F2: 1}
                var v8 uint = 4294967295
                _ = v7
                _ = v8
            }
        } else {
            fmt.Println(v1, v5)
        }
    case 1:
        v1 = false
    case -733:
        if (false && v1) {
            var v9 int8 = -31
            _ = v9
        }
    default:
        if (((v5 + 289) <= (v5 ^ v5)) && v1) {
            switch -1 {
            case -9223372036854775808:
                if true {
                    for fi10 := 0; fi10 < 9; fi10++ {
                        var v11 int64 = 9223372036854775807
                        v5 -= 1
                        _ = fi10
                        _ = v11
                    }
                    v12 := GenS1{F0: 130, F1: uint16(v5), F2: 128}
                    _ = v12
                } else {
                    for fi13 := 0; fi13 < 2; fi13++ {
                        for fi14 := 0; fi14 < 7; fi14++ {
                            if ((-2.25 < (-2.25 / 0.1)) && (0 >= -965)) {
                                var v15 float64 = 1e10
                                _ = v15
                            }
                            _ = fi14
                        }
                        _ = fi13
                    }
                }
            case -2147483648:
                for fi16 := 0; fi16 < 14; fi16++ {
                    if (v1 && (436 >= 2)) {
                        v17 := []int{(fi16 - (fi16 >> ((2) & 63))), -248}
                        _ = v17
                    }
                    v5 -= (v5 / (((v5 - 0)) | 1))
                    _ = fi16
                }
            case 127:
                v5 = 923
            default:
                var v18 int16 = -32768
                _ = v18
            }
            var v19 int = 127
            var v20 float64 = 1e10
            _ = v19
            _ = v20
        }
    }
    v21 := GenS2{F0: 9223372036854775807, F1: 42}
    var v22 int64 = 0
    v23 := GenS3{F0: 255, F1: v5}
    for fi24 := 0; fi24 < 18; fi24++ {
        v25 := GenS2{F0: (v22 ^ v22), F1: v5}
        if false {
            v1 = v1
        } else {
            delete(v3, v5)
            v22 = (^(v22))
        }
        v26 := GenS2{F0: (v22 | 0), F1: v5}
        _ = fi24
        _ = v25
        _ = v26
    }
    var v27 int8 = 123
    v28 := GenS1{F0: 42, F1: 2, F2: 29}
    if (155 > 0) {
        v29 := GenS2{F0: (v22 - 1), F1: 2147483648}
        for fi30 := 0; fi30 < 4; fi30++ {
            v31 := GenS1{F0: 255, F1: 765, F2: 62}
            for fi32 := 0; fi32 < 5; fi32++ {
                v5 -= 2147483648
                switch "" {
                case "hello":
                    v1 = (("a" + ("hello" + "a")) > ((("a" + "foo bar") + "\n\t") + "hello"))
                case "a":
                    v27 = 127
                case "":
                    var v33 int32 = -209
                    _ = v33
                default:
                    fmt.Println(v22, v22, v22)
                    switch "" {
                    case "":
                        v22 += v22
                    case "foo bar":
                        var v34 uint32 = v5
                        _ = v34
                    }
                }
                switch "a" {
                case "\n\t":
                    if (!v1) {
                        var v35 uint64 = 18446744073709551615
                        switch fi30 {
                        case -712:
                            var v36 int64 = v22
                            _ = v36
                        case -1:
                            var v37 uint8 = 1
                            _ = v37
                        case 32768:
                            v38 := []int{(fi30 >> ((574) & 63)), (fi30 << ((2) & 63))}
                            _ = v38
                        }
                        v39 := GenS1{F0: 188, F1: 296, F2: 127}
                        _ = v35
                        _ = v39
                    }
                case "a":
                    switch ("\n\t" + "\n\t") {
                    case "":
                        v1 = ((fi30 < fi30) || (((((-2.25 - (((1e10 * 0.0) + 3.14159) / 1.0)) == -2.25) || (3.14159 != float32(1e10))) || (1 >= 2)) && false))
                    case "foo bar":
                        var v40 uint32 = v5
                        _ = v40
                    case "a":
                        var v41 uint32 = 1
                        _ = v41
                    }
                case "":
                    v1 = v1
                default:
                    if false {
                        _ = len(v3)
                        v1 = false
                        v42 := GenS3{F0: 42, F1: (v5 & (v5 + 2))}
                        _ = v42
                    }
                }
                _ = fi32
            }
            fmt.Println(v22)
            _ = fi30
            _ = v31
        }
        v1 = false
        _ = v29
    } else {
        for fi43 := 0; fi43 < 4; fi43++ {
            v44 := []int{32767, 83}
            var v45 float64 = (((0.1 / 1.0) - (0.0 / 1.5)) / 1.0)
            _ = fi43
            _ = v44
            _ = v45
        }
        fmt.Println(v22, v27)
    }
    var v46 uint = 2
    v47 := v2[(v5 - v5)]
    fmt.Println(v47, v47, v1)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v27
    _ = v28
    _ = v3
    _ = v46
    _ = v47
    _ = v5
}
