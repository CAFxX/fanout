package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint8; F1 int8 }
type GenS4 struct { F0 uint64; F1 uint64; F2 uint32 }
type GenS5 struct { F0 uint64; F1 int64; F2 uint8 }
func genFn0(p0 int, p1 uint8, p2 int64) uint64 {
    fmt.Println(p0, p2)
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = p2
    return uint64(p1)
}
func genFn1() uint64 {
    if (127 < 127) {
        var v5 uint64 = 0
        _ = v5
    } else {
        switch "a" {
        case "hello":
            if ("hello" == "\n\t") {
                v6 := GenS1{F0: 100, F1: -385}
                _ = v6
            } else {
                var v7 uint64 = 2
                switch "\n\t" {
                case "":
                    v8 := GenS2{F0: -1}
                    _ = v8
                case "hello":
                    v7 = 18446744073709551615
                case "a":
                    for fi9 := 0; fi9 < 6; fi9++ {
                        fmt.Println(v7, fi9)
                        _ = fi9
                    }
                default:
                    v10 := GenS3{F0: 255, F1: 0}
                    _ = v10
                }
                v11 := GenS1{F0: 87, F1: -9223372036854775808}
                _ = v11
                _ = v7
            }
        case "":
            var v12 float32 = -2.25
            _ = v12
        default:
            v13 := GenS3{F0: 108, F1: 100}
            _ = v13
        }
        var v14 int32 = 42
        v15 := []int{-43, 128, 9223372036854775807}
        _ = v14
        _ = v15
    }
    v16 := GenS4{F0: 0, F1: 4294967295, F2: 330}
    _ = v16
    return 4294967295
}

func main() {
    v1 := make(map[int]string)
    var v2 int = 342
    var v3 uint64 = 832
    var v4 int32 = 0
    for fi17 := 0; fi17 < 4; fi17++ {
        for fi18 := 0; fi18 < 17; fi18++ {
            if (-915 <= 1) {
                v2 += 9223372036854775807
                fmt.Println(fi18, v3)
            } else {
                v2 += 863
                v1[-1] = "\n\t"
                fmt.Println(v2)
            }
            v19 := GenS4{F0: (v3 / ((v3) | 1)), F1: (v3 / ((v3) | 1)), F2: 497}
            _ = fi18
            _ = v19
        }
        if false {
            for fi20 := 0; fi20 < 1; fi20++ {
                v21 := []int{(v2 << ((4294967295) & 63)), -2147483648}
                _ = fi20
                _ = v21
            }
            for fi22 := 0; fi22 < 20; fi22++ {
                var v23 uint = 731
                var v24 int = 612
                var v25 int = (v24 * fi22)
                if ((1 != 517) || (44 >= 42)) {
                    v26 := []int{-2147483648}
                    if true {
                        if ((true && (false && true)) && true) {
                            v27 := v1[v25]
                            var v28 int64 = 190
                            _ = v27
                            _ = v28
                        } else {
                            var v29 int32 = 0
                            _ = v29
                        }
                        var v30 int = v24
                        var v31 uint = (v23 << ((v23) & 63))
                        _ = v30
                        _ = v31
                    }
                    _ = v26
                } else {
                    var v32 int = -786
                    _ = v32
                }
                _ = fi22
                _ = v23
                _ = v24
                _ = v25
            }
            if (((((3.14159 + 1.5) / float64(1.5)) - float64(((0.0 + 0.0) + 0.1))) * 3.14159) != 1e10) {
                for fi33 := 0; fi33 < 1; fi33++ {
                    v3 = (v3 | (v3 % ((genFn0((fi17 * v2), 0, 42)) | 1)))
                    var v34 uint = 65535
                    _ = fi33
                    _ = v34
                }
            }
        } else {
            fmt.Println(v3, fi17)
            switch fi17 {
            case 84:
                var v35 uint8 = 0
                _ = v35
            case -97:
                for fi36 := 0; fi36 < 7; fi36++ {
                    v2 = (fi17 % ((-875) | 1))
                    var v37 bool = true
                    var v38 int8 = 42
                    _ = fi36
                    _ = v37
                    _ = v38
                }
            case -1:
                _ = len(v1)
            default:
                fmt.Println(fi17, v4)
            }
        }
        _ = fi17
    }
    v39 := []int{(v2 ^ -646), (v2 * 32767)}
    v4 = v4
    _ = len(v1)
    var v40 uint8 = 42
    if ("" != "\n\t") {
        for fi41 := 0; fi41 < 13; fi41++ {
            if (!false) {
                v42 := GenS4{F0: genFn1(), F1: (v3 | 1), F2: 237}
                for fi43 := 0; fi43 < 1; fi43++ {
                    v4 = (v4 * v4)
                    v3 += 2
                    v40 = (v40 & (v40 / ((v40) | 1)))
                    _ = fi43
                }
                _ = v42
            } else {
                v3 -= (^(v3))
                fmt.Println(fi41, v40)
            }
            var v44 bool = true
            var v45 float32 = (0.1 / (((1.5 - 0.0) - 1e10) - 1e10))
            _ = fi41
            _ = v44
            _ = v45
        }
    }
    v46 := GenS5{F0: v3, F1: 0, F2: v40}
    for fi47 := 0; fi47 < 8; fi47++ {
        _ = len(v1)
        _ = fi47
    }
    var v48 int8 = -5
    v2 -= (v2 - int(v3))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v39:", len(v39), v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v48:", v48)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v46
    _ = v48
}
