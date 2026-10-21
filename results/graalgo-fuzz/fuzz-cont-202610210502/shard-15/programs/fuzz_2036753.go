package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 uint8 }
type GenS2 struct { F0 string; F1 int16 }
type GenS3 struct { F0 int; F1 string }
type GenS4 struct { F0 uint64; F1 uint }
func genFn0(p0 int) int32 {
    p0 -= p0
    var v6 float64 = 0.0
    v7 := []int{1}
    v6 = v6
    v8 := GenS1{F0: 255, F1: 0, F2: 255}
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return -44
}

func main() {
    var v1 string = "hello"
    v2 := make(map[string]string)
    v3 := []uint32{}
    v4 := make(map[int]int)
    v5 := make(map[string]int)
    switch 32767 {
    case 9223372036854775807:
        for fi9 := 0; fi9 < 9; fi9++ {
            v1 = ("" + "a")
            v1 = ("hello" + v1)
            _ = fi9
        }
    case -462:
        var v10 uint8 = 247
        _ = v10
    case 127:
        v1 = (v1 + v1)
    default:
        var v11 uint16 = 65535
        _ = v11
    }
    switch 4294967295 {
    case 128:
        var v12 uint64 = 1
        _ = v12
    case 1:
        v1 = "a"
    case 2147483647:
        switch 9223372036854775807 {
        case -9223372036854775808:
            if false {
                fmt.Println(v1)
                delete(v2, v1)
            } else {
                var v13 float32 = 1.5
                _ = v13
            }
        default:
            v14 := GenS2{F0: ("\n\t" + v1), F1: 1000}
            _ = v14
        }
    }
    v1 = (v1 + v1)
    var v15 int32 = -974
    var v16 uint16 = 111
    switch "\n\t" {
    case "a":
        var v17 int8 = 42
        _ = v17
    case "":
        if (!true) {
            v18 := v4[-9223372036854775808]
            v1 = (v1 + v1)
            v19 := GenS3{F0: (v18 >> ((uint(v16)) & 63)), F1: "hello"}
            _ = v18
            _ = v19
        } else {
            v20 := GenS1{F0: 42, F1: 2, F2: 255}
            if true {
                v1 = "foo bar"
                v15 = (v15 | (^(v15)))
            }
            if true {
                for fi21 := 0; fi21 < 8; fi21++ {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    v1 = ("" + (v1 + (v1 + v1)))
                    var v22 uint8 = 0
                    _ = fi21
                    _ = v22
                }
                v15 += v15
                var v23 int16 = 1000
                _ = v23
            } else {
                var v24 string = "\n\t"
                if false {
                    for fi25 := 0; fi25 < 20; fi25++ {
                        if ((127 != 128) && (!false)) {
                            v15 = (^(v15))
                        } else {
                            fmt.Println(v16, fi25, v16)
                            v26 := GenS1{F0: (^(v16)), F1: v16, F2: 128}
                            _ = v26
                        }
                        for fi27 := 0; fi27 < 15; fi27++ {
                            var v28 bool = false
                            v16 = (v16 / (((v16 + v16)) | 1))
                            _ = fi27
                            _ = v28
                        }
                        _ = len(v4)
                        _ = fi25
                    }
                    switch 128 {
                    case 128:
                        for fi29 := 0; fi29 < 9; fi29++ {
                            for fi30 := 0; fi30 < 3; fi30++ {
                                var v31 uint32 = 42
                                _ = fi30
                                _ = v31
                            }
                            v24 = v24
                            if true {
                                v1 = v24
                                var v32 uint32 = 4294967295
                                fmt.Println(v1)
                                _ = v32
                            } else {
                                v33 := GenS3{F0: -2147483648, F1: v24}
                                for fi34 := 0; fi34 < 7; fi34++ {
                                    v35 := GenS4{F0: uint64(v16), F1: 0}
                                    var v36 uint8 = 2
                                    _ = fi34
                                    _ = v35
                                    _ = v36
                                }
                                _ = v33
                            }
                            _ = fi29
                        }
                    case 127:
                        var v37 uint16 = v16
                        _ = v37
                    }
                } else {
                    var v38 int = 42
                    _ = v38
                }
                _ = v24
            }
            _ = v20
        }
    default:
        var v39 uint64 = 2
        _ = v39
    }
    v1 = (((v1 + v1) + v1) + (v1 + ((("foo bar" + ((v1 + (v1 + v1)) + "foo bar")) + "\n\t") + v1)))
    fmt.Println(v15, v15)
    if ("\n\t" <= v1) {
        if ((-(v15)) != (v15 >> ((18446744073709551615) & 63))) {
            v16 = (v16 - (v16 * v16))
        }
        v3 = append(v3, 1)
    } else {
        delete(v2, "")
    }
    _ = len(v3)
    var v40 float64 = (-2.25 + 0.0)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v5
}
