package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint64; F1 bool; F2 int8 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint }
func genFn0() uint {
    var v6 int8 = -1
    for fi7 := 0; fi7 < 4; fi7++ {
        v8 := []int{fi7, 0, (-(fi7))}
        _ = len(v8)
        for fi9 := 0; fi9 < 9; fi9++ {
            if false {
                var v10 int16 = -32768
                switch 4294967295 {
                case 9223372036854775807:
                    if len(v8) > 0 {
                        v8[0] = (^(fi9))
                    }
                case 42:
                    var v11 int64 = 9223372036854775807
                    _ = v11
                default:
                    if false {
                        if len(v8) >= 2 {
                            v8 = v8[1:len(v8)]
                        }
                    } else {
                        if len(v8) > 0 {
                            v8[0] = 127
                        }
                        if false {
                            v12 := GenS1{F0: 42}
                            fmt.Println(fi9, fi9, v6)
                            switch "foo bar" {
                            case "":
                                v10 -= (v10 % (((v10 - (v10 * (v10 & 1000)))) | 1))
                            case "a":
                                v10 += v10
                            case "hello":
                                var v13 int32 = -2147483648
                                if len(v8) > 0 {
                                    v8[0] = (-(fi9))
                                }
                                switch fi9 {
                                case -9223372036854775808:
                                    v14 := GenS2{F0: 1, F1: true, F2: v6}
                                    _ = v14
                                case 275:
                                    fmt.Println(v13)
                                }
                                _ = v13
                            }
                            _ = v12
                        }
                        var v15 uint16 = 65535
                        var v16 int64 = -9223372036854775808
                        v17 := GenS3{F0: 65535}
                        _ = v15
                        _ = v16
                        _ = v17
                    }
                }
                _ = len(v8)
                _ = v10
            } else {
                fmt.Println(fi7, v6, v6)
                var v18 float64 = (float64(3.14159) + 0.1)
                _ = v18
            }
            if (3.14159 != (-2.25 / -2.25)) {
                fmt.Println(fi9, v6, fi9)
            }
            var v19 int32 = 0
            _ = fi9
            _ = v19
        }
        _ = fi7
        _ = v8
    }
    var v20 uint8 = 42
    v21 := []int{-519, 128, 9223372036854775807, -611}
    switch (("foo bar" + "hello") + ("hello" + (("\n\t" + ("\n\t" + "")) + (((("" + "") + "") + "foo bar") + "\n\t")))) {
    case "\n\t":
        switch 0 {
        case 127:
            for fi22 := 0; fi22 < 19; fi22++ {
                var v23 uint16 = 2
                _ = fi22
                _ = v23
            }
        }
    case "":
        var v24 uint8 = 127
        _ = v24
    case "a":
        if (3.14159 >= 1.5) {
            switch "hello" {
            case "\n\t":
                var v25 int8 = v6
                _ = v25
            }
        }
    default:
        if len(v21) >= 2 {
            v21 = v21[1:len(v21)]
        }
    }
    _ = v20
    _ = v21
    _ = v6
    return 82
}

func main() {
    var v1 uint64 = 1
    v2 := make(map[string]uint32)
    var v3 string = ("" + ("a" + "\n\t"))
    v4 := make(map[string]string)
    v5 := make(map[uint32]uint32)
    var v26 string = "hello"
    fmt.Println(v3, v26, v3)
    v26 = "\n\t"
    var v27 uint64 = (v1 % ((v1) | 1))
    v28 := []int{32768, -2147483648, -2147483648}
    switch (("a" + ((("a" + "hello") + v3) + "hello")) + "") {
    case "hello":
        fmt.Println(v1)
    case "foo bar":
        var v29 int16 = -972
        _ = v29
    case "\n\t":
        for fi30 := 0; fi30 < 11; fi30++ {
            if len(v28) > 0 {
                v28[0] = fi30
            }
            _ = fi30
        }
    }
    if len(v28) > 0 {
        v28[0] = 2147483647
    }
    v31 := GenS4{F0: 65535}
    fmt.Println(v27, v3)
    delete(v4, (v26 + v26))
    if true {
        if true {
            if (1 < (^(v27))) {
                switch "hello" {
                case "\n\t":
                    if true {
                        var v32 bool = true
                        _ = v32
                    } else {
                        v26 = v3
                    }
                default:
                    if len(v28) >= 2 {
                        v28 = v28[1:len(v28)]
                    }
                }
            }
            v27 = (^(v27))
            v28 = append(v28, -677)
        }
        var v33 float32 = 0.0
        _ = v33
    }
    if (((0.0 - 1e10) < -2.25) && false) {
        var v34 uint = 0
        v3 = ("a" + v3)
        v35 := GenS1{F0: 128}
        _ = v34
        _ = v35
    } else {
        var v36 int = 128
        v37 := GenS4{F0: 255}
        _ = v36
        _ = v37
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v31
    _ = v4
    _ = v5
}
