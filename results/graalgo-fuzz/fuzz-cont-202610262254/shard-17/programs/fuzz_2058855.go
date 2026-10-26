package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int64; F1 uint32; F2 uint8 }
type GenS3 struct { F0 int32; F1 uint32; F2 uint }
func genFn0(p0 bool) bool {
    var v4 int = 128
    var v5 int16 = 42
    v4 += v4
    v6 := GenS1{F0: 255}
    v7 := []int{1, (-(v4)), (^(v4)), v4}
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return p0
}
func genFn1(p0 uint8) bool {
    switch 64 {
    case -244:
        var v8 float64 = -2.25
        _ = v8
    case 128:
        fmt.Println(p0)
    default:
        var v9 uint64 = 18446744073709551615
        _ = v9
    }
    var v10 bool = (32767 <= -683)
    fmt.Println(v10)
    _ = p0
    _ = v10
    return v10
}

func main() {
    var v1 int32 = 2147483647
    v2 := make(map[string]string)
    var v3 uint64 = 1
    v11 := GenS1{F0: 2}
    delete(v2, "foo bar")
    switch 127 {
    case -1:
        var v12 uint8 = 128
        _ = v12
    case 9223372036854775807:
        switch "\n\t" {
        case "hello":
            v3 -= 1
        case "\n\t":
            if true {
                if false {
                    if false {
                        if true {
                            if false {
                                v13 := []int{32767, 2147483647, -607, 1}
                                if (24 == 42) {
                                    var v14 float64 = (0.1 - (3.14159 + 0.0))
                                    fmt.Println(v1)
                                    v15 := GenS1{F0: 0}
                                    _ = v14
                                    _ = v15
                                } else {
                                    v3 += 4294967295
                                    var v16 uint = 42
                                    _ = v16
                                }
                                _ = v13
                            }
                        }
                        for fi17 := 0; fi17 < 11; fi17++ {
                            if true {
                                v1 += v1
                                switch "a" {
                                case "hello":
                                    if false {
                                        v2["foo bar"] = "hello"
                                        switch fi17 {
                                        case 4294967295:
                                            if false {
                                                v3 = 941
                                                fmt.Println(fi17)
                                            } else {
                                                v1 += v1
                                                v18 := GenS2{F0: -9223372036854775808, F1: 4294967295, F2: 128}
                                                v3 = (v3 | v3)
                                                _ = v18
                                            }
                                        case 991:
                                            var v19 uint16 = 304
                                            _ = v19
                                        case 128:
                                            for fi20 := 0; fi20 < 17; fi20++ {
                                                v1 -= v1
                                                v3 = (^(v3))
                                                var v21 uint8 = 128
                                                var v22 int32 = -586
                                                var v23 uint64 = 0
                                                _ = fi20
                                                _ = v21
                                                _ = v22
                                                _ = v23
                                            }
                                        default:
                                            var v24 string = "a"
                                            _ = v24
                                        }
                                        for fi25 := 0; fi25 < 18; fi25++ {
                                            v2["hello"] = "a"
                                            fmt.Println(fi17)
                                            var v26 uint16 = 2
                                            _ = fi25
                                            _ = v26
                                        }
                                    }
                                case "foo bar":
                                    v2["hello"] = ""
                                case "\n\t":
                                    for fi27 := 0; fi27 < 7; fi27++ {
                                        var v28 float32 = 1e10
                                        v29 := GenS2{F0: 0, F1: 5, F2: 255}
                                        var v30 int16 = -345
                                        _ = fi27
                                        _ = v28
                                        _ = v29
                                        _ = v30
                                    }
                                default:
                                    v3 += 4294967295
                                }
                            }
                            v31 := []int{32767, fi17, (fi17 ^ fi17), (fi17 >> ((18446744073709551615) & 63))}
                            v3 = v3
                            _ = fi17
                            _ = v31
                        }
                        v2[("\n\t" + "hello")] = ""
                    }
                }
                for fi32 := 0; fi32 < 19; fi32++ {
                    var v33 int64 = -1
                    _ = fi32
                    _ = v33
                }
            } else {
                v1 += 314
            }
        default:
            for fi34 := 0; fi34 < 8; fi34++ {
                if (2 != (v3 % ((v3) | 1))) {
                    v35 := GenS3{F0: (v1 / (((v1 / ((v1) | 1))) | 1)), F1: 65535, F2: 65535}
                    var v36 uint32 = 0
                    _ = v35
                    _ = v36
                } else {
                    if true {
                        if (8 == uint32(v3)) {
                            var v37 int8 = 42
                            v1 -= -2147483648
                            _ = v37
                        }
                    } else {
                        var v38 int32 = (-(v1))
                        _ = v38
                    }
                    fmt.Println(v1)
                    v39 := []int{fi34, fi34}
                    _ = v39
                }
                fmt.Println(v1, v3, fi34)
                fmt.Println(fi34)
                _ = fi34
            }
        }
    case 1:
        v40 := GenS1{F0: 255}
        _ = v40
    }
    for fi41 := 0; fi41 < 1; fi41++ {
        v3 -= v3
        _ = fi41
    }
    var v42 int64 = 881
    for fi43 := 0; fi43 < 20; fi43++ {
        v1 = 0
        v44 := []int{}
        _ = fi43
        _ = v44
    }
    fmt.Println(v42)
    switch 128 {
    case 1:
        v1 = v1
    default:
        if false {
            v1 += v1
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v42
}
