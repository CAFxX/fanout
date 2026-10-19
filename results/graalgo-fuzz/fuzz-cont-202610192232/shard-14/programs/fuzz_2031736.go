package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int; F2 string }
type GenS2 struct { F0 int; F1 uint64 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 string; F1 int8 }
type GenS5 struct { F0 string; F1 uint8; F2 uint16 }

func main() {
    v1 := make(map[int]string)
    var v2 bool = (218 > 662)
    var v3 uint8 = 0
    var v4 uint32 = 2
    var v5 uint64 = 740
    v6 := GenS1{F0: uint64(v4), F1: 0, F2: "\n\t"}
    if false {
        fmt.Println(v2, v3, v5)
    } else {
        switch -794 {
        case 683:
            for fi7 := 0; fi7 < 14; fi7++ {
                v8 := GenS2{F0: (fi7 ^ 9223372036854775807), F1: 42}
                _ = fi7
                _ = v8
            }
        case -9223372036854775808:
            var v9 uint = 2
            _ = v9
        case -50:
            if true {
                switch -451 {
                case 325:
                    var v10 uint16 = 161
                    _ = v10
                }
                var v11 string = ("foo bar" + "hello")
                if (0.1 != (-2.25 - -2.25)) {
                    v12 := GenS2{F0: 9223372036854775807, F1: uint64(v3)}
                    _ = v12
                }
                _ = v11
            }
        }
        v13 := GenS3{F0: 2}
        _ = v13
    }
    delete(v1, 317)
    var v14 int64 = 1
    fmt.Println(v2)
    v4 += 2
    v15 := []int{128, 32767, -1, 857}
    switch 127 {
    case -2147483648:
        switch 0 {
        case 128:
            if (42 == (v3 * 0)) {
                var v16 string = "a"
                v16 = "a"
                for fi17 := 0; fi17 < 20; fi17++ {
                    switch "a" {
                    case "foo bar":
                        if v2 {
                            fmt.Println(v2, v4)
                        }
                    case "hello":
                        if (!false) {
                            v5 -= v5
                        } else {
                            v14 = (v14 + -9223372036854775808)
                            v15 = append(v15, fi17)
                        }
                    case "a":
                        if (v2 && true) {
                            switch (-(fi17)) {
                            case -148:
                                if len(v15) > 0 {
                                    v15[0] = fi17
                                }
                            case 796:
                                v16 = "\n\t"
                            case 2147483647:
                                switch "hello" {
                                case "":
                                    var v18 int32 = 0
                                    _ = v18
                                case "hello":
                                    for fi19 := 0; fi19 < 11; fi19++ {
                                        var v20 uint = 557
                                        _ = fi19
                                        _ = v20
                                    }
                                case "\n\t":
                                    v2 = true
                                }
                            default:
                                if v2 {
                                    v21 := GenS4{F0: (v16 + ("foo bar" + "\n\t")), F1: int8(v3)}
                                    _ = v21
                                } else {
                                    v15 = append(v15, 2147483647)
                                    v3 = 51
                                }
                            }
                        } else {
                            var v22 bool = v2
                            v4 = v4
                            _ = v22
                        }
                    }
                    _ = fi17
                }
                _ = v16
            }
        case -49:
            if (255 >= uint(v3)) {
                delete(v1, 127)
                var v23 int = 9223372036854775807
                var v24 uint32 = (v4 & v4)
                _ = v23
                _ = v24
            } else {
                v1[2147483647] = "foo bar"
                var v25 uint = 65535
                var v26 uint32 = v4
                _ = v25
                _ = v26
            }
        case 42:
            v4 -= (^(v4))
        default:
            v27 := GenS4{F0: (("" + "") + (("a" + (("foo bar" + "") + "")) + (("a" + "\n\t") + ("a" + (("" + "\n\t") + ""))))), F1: -44}
            _ = v27
        }
    case 888:
        v5 -= v5
    }
    v28 := v1[128]
    var v29 uint64 = (v5 / (((v5 / ((v5) | 1))) | 1))
    for fi30 := 0; fi30 < 11; fi30++ {
        for fi31 := 0; fi31 < 13; fi31++ {
            v32 := GenS5{F0: (v28 + (v28 + v28)), F1: v3, F2: 2}
            var v33 float32 = 1.5
            var v34 uint = 684
            _ = fi31
            _ = v32
            _ = v33
            _ = v34
        }
        v35 := GenS5{F0: (v28 + ""), F1: 105, F2: 42}
        _ = fi30
        _ = v35
    }
    fmt.Println(v2, v14)
    var v36 int8 = -52
    _ = len(v15)
    fmt.Println(v4, v4)
    v15 = append(v15, 32767)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
