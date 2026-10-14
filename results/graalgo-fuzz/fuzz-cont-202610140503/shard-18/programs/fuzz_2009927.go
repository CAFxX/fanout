package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int8; F1 int16; F2 uint32 }
type GenS3 struct { F0 bool; F1 int32 }
type GenS4 struct { F0 uint; F1 uint; F2 string }
type GenS5 struct { F0 uint16; F1 int32 }
func genFn0() int32 {
    v5 := []int{32768, -1}
    for fi6 := 0; fi6 < 19; fi6++ {
        if (-1 < 42) {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            var v7 float32 = (1e10 - 0.1)
            _ = v7
        } else {
            var v8 int = fi6
            v8 = fi6
            if true {
                var v9 int64 = 1
                fmt.Println(v8)
                v10 := GenS1{F0: v8}
                _ = v10
                _ = v9
            }
            _ = v8
        }
        v11 := GenS2{F0: -1, F1: -1, F2: 483}
        _ = fi6
        _ = v11
    }
    if true {
        var v12 int = 0
        _ = v12
    } else {
        var v13 bool = false
        var v14 int16 = -631
        _ = v13
        _ = v14
    }
    v15 := GenS2{F0: -128, F1: -738, F2: 405}
    for fi16 := 0; fi16 < 19; fi16++ {
        var v17 bool = false
        _ = fi16
        _ = v17
    }
    _ = v15
    _ = v5
    return -49
}
func genFn1(p0 int) int64 {
    for fi18 := 0; fi18 < 4; fi18++ {
        switch "\n\t" {
        case "":
            for fi19 := 0; fi19 < 12; fi19++ {
                if true {
                    if true {
                        v20 := []int{(fi19 + 2147483647)}
                        fmt.Println(p0, fi18, p0)
                        p0 -= (p0 | (fi19 & 42))
                        if (!(-947 != -32768)) {
                            var v21 string = ""
                            p0 = fi19
                            fmt.Println(v21)
                            _ = v21
                        } else {
                            v22 := GenS3{F0: (!(!false)), F1: -1}
                            v23 := GenS3{F0: ((!(!(148 != 47))) || (1 != 4294967295)), F1: genFn0()}
                            _ = v22
                            _ = v23
                        }
                        _ = v20
                    } else {
                        v24 := GenS3{F0: true, F1: 235}
                        switch p0 {
                        case 4294967295:
                            p0 = fi19
                        case 1:
                            v25 := []int{-1, fi19, 128, fi18}
                            _ = v25
                        default:
                            v26 := []int{(p0 + -667), p0}
                            _ = v26
                        }
                        _ = v24
                    }
                    var v27 string = (("a" + "\n\t") + "a")
                    _ = v27
                }
                var v28 uint8 = 0
                v29 := GenS3{F0: true, F1: 2147483647}
                _ = fi19
                _ = v28
                _ = v29
            }
        default:
            p0 = p0
        }
        var v30 int8 = 1
        _ = fi18
        _ = v30
    }
    p0 -= (^(p0))
    var v31 uint16 = 542
    var v32 int = p0
    p0 -= p0
    _ = p0
    _ = v31
    _ = v32
    return 9223372036854775807
}

func main() {
    v1 := []uint32{0, 42}
    var v2 int16 = 1
    var v3 uint16 = 255
    var v4 float32 = 3.14159
    if true {
        v3 = (v3 - (v3 & v3))
    }
    var v33 uint64 = 18446744073709551615
    switch 2147483647 {
    case -648:
        for fi34 := 0; fi34 < 3; fi34++ {
            var v35 uint16 = v3
            _ = fi34
            _ = v35
        }
    default:
        v33 = 18446744073709551615
    }
    var v36 float32 = (((-2.25 * (float32(1e10) - v4)) * -2.25) * v4)
    v4 = v36
    v36 = -2.25
    if len(v1) > 0 {
        v1[0] = 65535
    }
    fmt.Println(v33)
    v37 := GenS4{F0: 0, F1: 2, F2: "\n\t"}
    v38 := GenS1{F0: 128}
    var v39 int32 = 662
    v40 := GenS5{F0: v3, F1: (v39 % (((v39 % ((v39) | 1))) | 1))}
    v41 := GenS3{F0: (false && (999 <= -1)), F1: (v39 / (((v39 | v39)) | 1))}
    switch "a" {
    case "":
        v2 = v2
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v36:", v36)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
}
