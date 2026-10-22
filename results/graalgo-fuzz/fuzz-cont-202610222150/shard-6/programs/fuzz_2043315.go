package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int }
type GenS2 struct { F0 int8; F1 int64 }
type GenS3 struct { F0 string; F1 uint32 }
type GenS4 struct { F0 uint16; F1 int64; F2 int16 }
func genFn0(p0 string, p1 int64) uint32 {
    var v6 int16 = 0
    var v7 bool = (1.5 != 0.1)
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 1
}

func main() {
    var v1 uint8 = 198
    var v2 int32 = 1
    var v3 int8 = -22
    v4 := []int{4294967295, 733, -91}
    v5 := make(map[int]string)
    var v8 int16 = 1000
    if len(v4) > 0 {
        v4[0] = -433
    }
    fmt.Println(v2)
    v9 := GenS1{F0: v3, F1: 989}
    switch 935 {
    case 42:
        var v10 int16 = (v8 * -32768)
        _ = v10
    case 2147483647:
        if (true && false) {
            switch 4294967295 {
            case 9223372036854775807:
                v11 := GenS1{F0: v3, F1: -2147483648}
                _ = v11
            default:
                v3 = (-(v3))
            }
            var v12 int16 = (^(v8))
            v13 := GenS1{F0: int8(v1), F1: -573}
            _ = v12
            _ = v13
        }
    case 128:
        if false {
            v14 := GenS2{F0: -31, F1: 9223372036854775807}
            var v15 uint32 = uint32(v1)
            v16 := GenS3{F0: "a", F1: v15}
            _ = v14
            _ = v15
            _ = v16
        }
    default:
        _ = len(v4)
    }
    for fi17 := 0; fi17 < 17; fi17++ {
        for fi18 := 0; fi18 < 14; fi18++ {
            v19 := GenS3{F0: "foo bar", F1: 65535}
            v3 += 100
            var v20 bool = false
            _ = fi18
            _ = v19
            _ = v20
        }
        var v21 uint = 0
        v22 := GenS4{F0: 995, F1: -946, F2: v8}
        _ = fi17
        _ = v21
        _ = v22
    }
    for fi23 := 0; fi23 < 15; fi23++ {
        fmt.Println(v1, v2)
        _ = fi23
    }
    switch -763 {
    case -9223372036854775808:
        for fi24 := 0; fi24 < 17; fi24++ {
            switch fi24 {
            case -1:
                switch "\n\t" {
                case "hello":
                    if (-2.25 >= 3.14159) {
                        v25 := v5[fi24]
                        _ = v25
                    } else {
                        v3 += (v3 / (((v3 % ((0) | 1))) | 1))
                        delete(v5, fi24)
                        fmt.Println(v2, v1)
                    }
                case "":
                    for fi26 := 0; fi26 < 6; fi26++ {
                        fmt.Println(v3, v3, fi24)
                        switch "a" {
                        case "\n\t":
                            var v27 uint32 = 65535
                            _ = v27
                        case "a":
                            delete(v5, fi24)
                        }
                        _ = fi26
                    }
                }
            }
            fmt.Println(v3)
            fmt.Println(v2)
            _ = fi24
        }
    case 553:
        switch ("\n\t" + "a") {
        case "hello":
            var v28 int = 32767
            _ = v28
        case "foo bar":
            var v29 uint = 774
            _ = v29
        default:
            var v30 float64 = 1.5
            _ = v30
        }
    case 32768:
        v31 := GenS4{F0: 714, F1: 1, F2: (v8 - (v8 % (((v8 << ((2) & 63))) | 1)))}
        _ = v31
    default:
        if true {
            v2 -= 1
            v32 := GenS4{F0: 2, F1: 42, F2: (v8 & v8)}
            _ = v32
        } else {
            var v33 uint8 = (v1 + (v1 | (v1 % ((v1) | 1))))
            var v34 float32 = 1.5
            v33 = (^(v33))
            _ = v33
            _ = v34
        }
    }
    var v35 int16 = v8
    v36 := GenS1{F0: 1, F1: 353}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
