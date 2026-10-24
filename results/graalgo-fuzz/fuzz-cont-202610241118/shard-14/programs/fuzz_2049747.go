package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool; F2 int32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 bool; F1 int64; F2 string }
type GenS4 struct { F0 int8; F1 string }

func main() {
    var v1 string = ("foo bar" + ("hello" + "hello"))
    v2 := make(map[uint32]string)
    var v3 uint64 = 42
    v4 := make(map[string]uint32)
    switch 9223372036854775807 {
    case 736:
        if false {
            v3 = 482
        } else {
            _ = len(v2)
            switch (v1 + v1) {
            case "\n\t":
                var v5 string = v1
                _ = v5
            case "hello":
                switch -2147483648 {
                case 2147483647:
                    for fi6 := 0; fi6 < 11; fi6++ {
                        v1 = ((v1 + v1) + v1)
                        if (false && (128 < fi6)) {
                            v7 := GenS1{F0: 65535, F1: false, F2: -1}
                            var v8 uint32 = 1
                            switch (fi6 & (-(fi6))) {
                            case -2147483648:
                                v9 := GenS1{F0: 42, F1: (!true), F2: 1}
                                _ = v9
                            case -395:
                                var v10 int8 = -1
                                _ = v10
                            }
                            _ = v7
                            _ = v8
                        }
                        var v11 string = ("a" + (v1 + v1))
                        _ = fi6
                        _ = v11
                    }
                case -2147483648:
                    delete(v2, 0)
                case 128:
                    var v12 float32 = 3.14159
                    _ = v12
                }
            default:
                if (((9223372036854775807 == 42) && (true || true)) || true) {
                    var v13 uint16 = 32767
                    var v14 int32 = 61
                    _ = v13
                    _ = v14
                } else {
                    for fi15 := 0; fi15 < 14; fi15++ {
                        fmt.Println(v1, fi15, v3)
                        switch v1 {
                        case "hello":
                            v16 := []int{(fi15 / ((fi15) | 1)), (fi15 ^ (-(fi15))), 128}
                            _ = v16
                        }
                        for fi17 := 0; fi17 < 8; fi17++ {
                            fmt.Println(v1, fi17)
                            switch v1 {
                            case "\n\t":
                                v3 = v3
                            }
                            _ = fi17
                        }
                        _ = fi15
                    }
                    if (42 > 0) {
                        v1 = v1
                    } else {
                        var v18 uint16 = 590
                        v19 := GenS1{F0: 255, F1: false, F2: 1}
                        v20 := v4[(v1 + v1)]
                        _ = v18
                        _ = v19
                        _ = v20
                    }
                }
            }
            v2[1] = "\n\t"
        }
    case 0:
        if true {
            for fi21 := 0; fi21 < 5; fi21++ {
                v3 += 967
                v22 := []int{(fi21 | (fi21 >> ((18446744073709551615) & 63)))}
                _ = fi21
                _ = v22
            }
            v23 := GenS2{F0: 100}
            v1 = "hello"
            _ = v23
        }
    case -610:
        v24 := GenS3{F0: true, F1: 9223372036854775807, F2: v1}
        _ = v24
    default:
        fmt.Println(v1)
    }
    v1 = v1
    fmt.Println(v3, v1)
    var v25 int = 2147483647
    switch "" {
    case "hello":
        v26 := []int{2147483647}
        _ = v26
    case "":
        if (42 < int64(v3)) {
            switch v1 {
            case "\n\t":
                fmt.Println(v3)
            case "a":
                v27 := []int{}
                _ = v27
            }
            switch v25 {
            case 9223372036854775807:
                v28 := GenS4{F0: 100, F1: ("foo bar" + (v1 + "a"))}
                _ = v28
            default:
                var v29 int16 = int16(v25)
                _ = v29
            }
            for fi30 := 0; fi30 < 11; fi30++ {
                if false {
                    for fi31 := 0; fi31 < 16; fi31++ {
                        v3 -= 4294967295
                        _ = fi31
                    }
                } else {
                    v3 -= v3
                    v32 := []int{1, (fi30 & (fi30 << ((0) & 63)))}
                    v33 := GenS2{F0: -110}
                    _ = v32
                    _ = v33
                }
                _ = fi30
            }
        }
    }
    v34 := []int{(v25 - (v25 - (^(v25)))), v25, (^(v25)), (v25 & v25)}
    v25 -= v25
    _ = len(v34)
    _ = len(v34)
    var v35 bool = true
    v3 -= v3
    if true {
        var v36 bool = (0 > 2)
        v2[2147483648] = v1
        v37 := GenS1{F0: 255, F1: (!(v36 || v36)), F2: 1}
        _ = v36
        _ = v37
    }
    v38 := GenS3{F0: false, F1: -9223372036854775808, F2: v1}
    _ = len(v34)
    v39 := GenS1{F0: 992, F1: (false && false), F2: -2147483648}
    v40 := GenS2{F0: 100}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v34
    _ = v35
    _ = v38
    _ = v39
    _ = v4
    _ = v40
}
