package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool }
type GenS2 struct { F0 int; F1 int16; F2 int64 }
type GenS3 struct { F0 uint64; F1 uint8; F2 int }
func genFn0() bool {
    var v4 uint = 517
    var v5 int16 = -969
    _ = v4
    _ = v5
    return true
}
func genFn1(p0 bool, p1 uint32, p2 int32) int8 {
    var v6 uint = 997
    if p0 {
        if false {
            if (4294967295 >= 719) {
                var v7 int8 = 127
                v8 := []int{0, 4294967295, 32768, 32767}
                _ = v7
                _ = v8
            } else {
                var v9 int = -737
                var v10 int32 = p2
                p2 += (p2 >> ((65535) & 63))
                _ = v10
                _ = v9
            }
        } else {
            fmt.Println(v6, v6, v6)
        }
        v6 = (v6 / (((v6 - v6)) | 1))
        p1 = (p1 * p1)
    }
    var v11 uint = v6
    var v12 bool = (-629 != 1)
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    _ = v6
    return -128
}

func main() {
    v1 := make(map[uint32]string)
    v2 := []int{128, -2147483648, -9223372036854775808}
    var v3 int32 = 879
    var v13 string = ""
    v3 = v3
    var v14 uint = 1
    v14 = (v14 >> ((v14) & 63))
    for fi15 := 0; fi15 < 4; fi15++ {
        for fi16 := 0; fi16 < 10; fi16++ {
            var v17 uint16 = 255
            var v18 float32 = 1.5
            _ = fi16
            _ = v17
            _ = v18
        }
        _ = fi15
    }
    v14 = (v14 / (((v14 % (((v14 * v14)) | 1))) | 1))
    if len(v2) > 0 {
        v2[0] = -9223372036854775808
    }
    v14 = (v14 + 1)
    var v19 bool = false
    fmt.Println(v3, v13, v14)
    _ = len(v2)
    var v20 uint32 = 923
    _ = len(v1)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    switch "foo bar" {
    case "":
        var v21 uint16 = 65535
        _ = v21
    case "a":
        fmt.Println(v19, v14, v14)
    default:
        for fi22 := 0; fi22 < 11; fi22++ {
            switch (fi22 | (fi22 % (((fi22 & fi22)) | 1))) {
            case -952:
                var v23 uint16 = 110
                _ = v23
            default:
                for fi24 := 0; fi24 < 12; fi24++ {
                    switch 127 {
                    case 128:
                        v25 := GenS1{F0: 21, F1: (4294967295 >= (v20 ^ v20))}
                        _ = v25
                    case -9223372036854775808:
                        _ = len(v1)
                    default:
                        if true {
                            var v26 int32 = v3
                            _ = v26
                        } else {
                            switch 127 {
                            case 127:
                                var v27 uint64 = 18446744073709551615
                                _ = v27
                            case 655:
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            case 4294967295:
                                v2 = append(v2, -443)
                            default:
                                v28 := GenS2{F0: (fi22 + fi22), F1: -280, F2: 12}
                                _ = v28
                            }
                            switch fi22 {
                            case 615:
                                if v19 {
                                    v3 += 329
                                    if false {
                                        if v19 {
                                            var v29 uint32 = v20
                                            _ = v29
                                        }
                                        var v30 float64 = 1.5
                                        v31 := GenS2{F0: (fi22 + (fi22 / ((fi24) | 1))), F1: -1, F2: 542}
                                        var v32 string = v13
                                        _ = v30
                                        _ = v31
                                        _ = v32
                                    } else {
                                        fmt.Println(v3)
                                        v14 = v14
                                        switch (v13 + (v13 + "hello")) {
                                        case "\n\t":
                                            if v19 {
                                                var v33 uint64 = 0
                                                _ = v33
                                            }
                                        }
                                        fmt.Println(v19)
                                    }
                                } else {
                                    var v34 int64 = 1
                                    switch fi22 {
                                    case -717:
                                        var v35 uint = (v14 * v14)
                                        _ = v35
                                    default:
                                        var v36 uint16 = 265
                                        _ = v36
                                    }
                                    _ = v34
                                }
                            }
                            var v37 int8 = -1
                            _ = v37
                        }
                    }
                    delete(v1, (v20 & 65535))
                    if v19 {
                        v38 := GenS1{F0: 42, F1: v19}
                        v13 = v13
                        v39 := GenS3{F0: 18446744073709551615, F1: 211, F2: fi22}
                        _ = v38
                        _ = v39
                    }
                    _ = fi24
                }
            }
            _ = fi22
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v14
    _ = v19
    _ = v2
    _ = v20
    _ = v3
}
