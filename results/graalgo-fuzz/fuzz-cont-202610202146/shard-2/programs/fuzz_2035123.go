package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint; F1 uint64; F2 string }
type GenS4 struct { F0 uint8 }
func genFn0(p0 bool) uint {
    v5 := GenS1{F0: 766, F1: 1}
    for fi6 := 0; fi6 < 14; fi6++ {
        p0 = p0
        _ = fi6
    }
    _ = p0
    _ = v5
    return 0
}

func main() {
    var v1 uint64 = 1
    var v2 bool = (-2.25 >= 0.1)
    v3 := make(map[uint32]bool)
    var v4 int8 = -58
    v7 := []int{-809, 4294967295}
    v8 := GenS1{F0: -217, F1: v4}
    if v2 {
        _ = len(v7)
        v1 = v1
    }
    switch -911 {
    case 2147483647:
        for fi9 := 0; fi9 < 15; fi9++ {
            switch "\n\t" {
            case "foo bar":
                v1 = v1
            case "hello":
                if v2 {
                    delete(v3, 1)
                    var v10 uint8 = 2
                    v10 -= v10
                    _ = v10
                } else {
                    if (v2 && ((v1 | v1) == v1)) {
                        fmt.Println(fi9, v2)
                        var v11 uint64 = v1
                        _ = v11
                    } else {
                        var v12 float32 = 0.1
                        for fi13 := 0; fi13 < 4; fi13++ {
                            v14 := GenS1{F0: -816, F1: -116}
                            delete(v3, 1)
                            _ = fi13
                            _ = v14
                        }
                        if (-1 < 609) {
                            var v15 uint32 = 139
                            v3[v15] = false
                            _ = v15
                        } else {
                            if true {
                                var v16 float32 = v12
                                _ = v16
                            }
                            var v17 int8 = 100
                            _ = v17
                        }
                        _ = v12
                    }
                }
            case "":
                v18 := GenS2{F0: 1}
                _ = v18
            }
            v19 := GenS3{F0: 766, F1: 0, F2: ("" + "foo bar")}
            _ = fi9
            _ = v19
        }
    case 4294967295:
        var v20 float32 = 1e10
        _ = v20
    case -890:
        switch ("foo bar" + ("\n\t" + "\n\t")) {
        case "foo bar":
            fmt.Println(v1, v2, v4)
        case "a":
            v21 := GenS3{F0: 423, F1: (v1 + (v1 << ((18446744073709551615) & 63))), F2: "foo bar"}
            _ = v21
        case "hello":
            v2 = false
        default:
            if (true && (1 == 255)) {
                for fi22 := 0; fi22 < 14; fi22++ {
                    if v2 {
                        v23 := v3[4294967295]
                        _ = v23
                    } else {
                        for fi24 := 0; fi24 < 2; fi24++ {
                            v2 = v2
                            v25 := GenS2{F0: 1}
                            v1 = (v1 * 2)
                            _ = fi24
                            _ = v25
                        }
                    }
                    v26 := v3[77]
                    _ = len(v3)
                    _ = fi22
                    _ = v26
                }
                var v27 uint64 = v1
                _ = v27
            }
        }
    default:
        if false {
            v28 := GenS2{F0: 1}
            _ = v28
        }
    }
    for fi29 := 0; fi29 < 17; fi29++ {
        v30 := GenS4{F0: 93}
        var v31 uint8 = 255
        _ = fi29
        _ = v30
        _ = v31
    }
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    var v32 int16 = -1
    fmt.Println(v2)
    var v33 int16 = (^(v32))
    if len(v7) > 0 {
        v7[0] = 128
    }
    v32 = v32
    if len(v7) > 0 {
        v7[0] = 2147483647
    }
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v7
    _ = v8
}
