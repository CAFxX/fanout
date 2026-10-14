package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string }
type GenS2 struct { F0 uint8; F1 uint }
type GenS3 struct { F0 uint32; F1 uint8; F2 string }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int8 }

func main() {
    var v1 uint16 = 84
    var v2 int64 = 1
    var v3 uint = 2
    var v4 string = "hello"
    var v5 uint8 = 42
    var v6 int32 = -650
    v1 -= v1
    var v7 uint64 = 1
    v8 := GenS1{F0: "", F1: (v4 + v4)}
    if false {
        v2 += v2
        v3 += (v3 & (v3 - v3))
        var v9 int8 = 42
        _ = v9
    } else {
        v10 := GenS2{F0: 127, F1: v3}
        _ = v10
    }
    v11 := []int{-9223372036854775808, 42, 0, 1}
    switch v4 {
    case "":
        for fi12 := 0; fi12 < 17; fi12++ {
            if len(v11) >= 2 {
                v11 = v11[1:len(v11)]
            }
            v1 = (v1 % (((v1 % ((42) | 1))) | 1))
            _ = fi12
        }
    case "a":
        for fi13 := 0; fi13 < 9; fi13++ {
            v14 := GenS1{F0: "a", F1: ("foo bar" + "")}
            switch (fi13 + fi13) {
            case -9223372036854775808:
                _ = len(v11)
            case 2147483647:
                fmt.Println(v6)
            default:
                switch (fi13 << (((v3 & v3)) & 63)) {
                case -2147483648:
                    _ = len(v11)
                case 9223372036854775807:
                    switch (fi13 - (fi13 + 1)) {
                    case 2147483647:
                        v11 = append(v11, fi13)
                    case 650:
                        fmt.Println(v7, v6, v7)
                    default:
                        switch (fi13 & (fi13 & 4294967295)) {
                        case 594:
                            if false {
                                var v15 int32 = (v6 / ((v6) | 1))
                                v6 = 0
                                fmt.Println(v1, v3)
                                _ = v15
                            }
                        case 127:
                            for fi16 := 0; fi16 < 9; fi16++ {
                                v17 := GenS2{F0: 255, F1: v3}
                                v5 = uint8(v3)
                                if false {
                                    v5 -= 224
                                    var v18 int8 = 42
                                    _ = v18
                                } else {
                                    if false {
                                        switch v4 {
                                        case "a":
                                            var v19 float32 = 0.1
                                            _ = v19
                                        case "\n\t":
                                            var v20 float32 = 1.5
                                            switch 1 {
                                            case 555:
                                                _ = len(v11)
                                            default:
                                                fmt.Println(v20, fi13)
                                            }
                                            var v21 int64 = v2
                                            _ = v20
                                            _ = v21
                                        }
                                        if false {
                                            v1 += 2
                                            if false {
                                                var v22 uint8 = v5
                                                v5 = v22
                                                v23 := GenS3{F0: 2, F1: 42, F2: "foo bar"}
                                                _ = v22
                                                _ = v23
                                            } else {
                                                if len(v11) >= 2 {
                                                    v11 = v11[1:len(v11)]
                                                }
                                                _ = len(v11)
                                                var v24 int64 = (v2 >> ((uint(v7)) & 63))
                                                _ = v24
                                            }
                                        }
                                        _ = len(v11)
                                    }
                                    v3 = (v3 % ((v3) | 1))
                                }
                                _ = fi16
                                _ = v17
                            }
                        case -2147483648:
                            if true {
                                _ = len(v11)
                                if (!(("foo bar" + (v4 + v4)) == v4)) {
                                    var v25 string = v4
                                    fmt.Println(v2)
                                    _ = v25
                                }
                                v26 := GenS4{F0: fi13}
                                _ = v26
                            } else {
                                v1 -= v1
                                v27 := GenS5{F0: 127}
                                _ = v27
                            }
                        }
                    }
                default:
                    var v28 float64 = float64((0.1 / float32((((float64((0.0 - (float32((((0.0 / 1.0) + 0.1) * 1.5)) - (float32(1e10) + 1e10)))) + 1e10) + 3.14159) * -2.25))))
                    _ = v28
                }
            }
            v7 -= v7
            _ = fi13
            _ = v14
        }
    case "hello":
        fmt.Println(v3, v7, v4)
    }
    var v29 bool = (0 >= 691)
    fmt.Println(v2)
    fmt.Println(v1, v5)
    var v30 int64 = (v2 / ((v2) | 1))
    var v31 int32 = (v6 / ((992) | 1))
    var v32 uint16 = 710
    if true {
        var v33 string = v4
        for fi34 := 0; fi34 < 19; fi34++ {
            var v35 string = (v4 + ((v33 + (((v33 + v4) + (("a" + "\n\t") + v33)) + v4)) + v4))
            _ = fi34
            _ = v35
        }
        _ = v33
    }
    fmt.Println(v30)
    v4 = "foo bar"
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
