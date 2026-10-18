package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint64; F1 uint32; F2 uint32 }

func main() {
    v1 := []uint32{4294967295, 0, 399, 2147483648}
    var v2 uint32 = 2147483647
    var v3 uint8 = 101
    var v4 float32 = 3.14159
    fmt.Println(v3)
    var v5 uint16 = 255
    _ = len(v1)
    for fi6 := 0; fi6 < 9; fi6++ {
        switch (fi6 << ((42) & 63)) {
        case 42:
            var v7 int = 4294967295
            _ = v7
        case 9223372036854775807:
            var v8 uint8 = (v3 & (v3 >> ((703) & 63)))
            _ = v8
        case 127:
            v2 = 1
        }
        for fi9 := 0; fi9 < 8; fi9++ {
            switch (-(fi9)) {
            case 417:
                var v10 int64 = 1
                _ = v10
            case 9223372036854775807:
                v4 += 1e10
            case -117:
                v11 := GenS1{F0: (fi9 + int(v2)), F1: 1}
                _ = v11
            default:
                if ((32767 <= 196) && false) {
                    fmt.Println(fi9, fi9, fi9)
                    if false {
                        if ((true || (-45 >= 1)) || (0 >= 4294967295)) {
                            fmt.Println(fi6)
                        }
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        if (v3 != (v3 << ((317) & 63))) {
                            var v12 uint32 = (^(v2))
                            _ = v12
                        } else {
                            v13 := GenS2{F0: "a"}
                            v5 = v5
                            if len(v1) > 0 {
                                v1[0] = uint32(v3)
                            }
                            _ = v13
                        }
                    } else {
                        var v14 uint16 = v5
                        _ = v14
                    }
                    if true {
                        fmt.Println(fi6, fi9)
                        var v15 bool = false
                        _ = v15
                    } else {
                        if ((fi6 - (fi6 % (((^(fi9))) | 1))) <= fi6) {
                            v5 = (v5 + 483)
                        }
                    }
                } else {
                    v16 := GenS2{F0: "\n\t"}
                    fmt.Println(v3, fi6)
                    if true {
                        v17 := GenS2{F0: "a"}
                        _ = v17
                    } else {
                        switch (fi9 >> ((0) & 63)) {
                        case 2147483647:
                            switch "hello" {
                            case "foo bar":
                                var v18 string = "a"
                                _ = v18
                            }
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        case 0:
                            switch (fi6 >> ((65535) & 63)) {
                            case 32767:
                                if true {
                                    if (float64(v4) >= -2.25) {
                                        _ = len(v1)
                                        var v19 string = "hello"
                                        _ = v19
                                    } else {
                                        v5 = v5
                                        var v20 uint64 = 703
                                        switch (("foo bar" + (("" + "a") + ("hello" + "foo bar"))) + (("hello" + ("" + "")) + (("foo bar" + (("" + (("" + "a") + "")) + "")) + ("a" + "hello")))) {
                                        case "foo bar":
                                            var v21 uint8 = v3
                                            _ = v21
                                        default:
                                            fmt.Println(v2, fi9)
                                        }
                                        _ = v20
                                    }
                                } else {
                                    if true {
                                        var v22 int = (fi6 + -9223372036854775808)
                                        v22 += -554
                                        _ = v22
                                    } else {
                                        v5 += v5
                                        if true {
                                            _ = len(v1)
                                        }
                                    }
                                }
                            }
                        case -428:
                            v5 -= v5
                        }
                        v23 := GenS3{F0: 2}
                        if ((true || false) && (v5 <= uint16(v2))) {
                            var v24 int8 = -1
                            v3 = (^(v3))
                            v2 -= 2
                            _ = v24
                        }
                        _ = v23
                    }
                    _ = v16
                }
            }
            v1 = append(v1, (v2 % ((v2) | 1)))
            switch "a" {
            case "foo bar":
                switch "" {
                case "\n\t":
                    var v25 int64 = 42
                    _ = v25
                default:
                    v3 += v3
                }
            default:
                switch (fi9 - (fi9 ^ fi6)) {
                case 127:
                    v26 := GenS4{F0: int8(v2)}
                    _ = v26
                default:
                    v2 += (v2 + (^(v2)))
                }
            }
            _ = fi9
        }
        _ = fi6
    }
    v2 += 42
    v3 -= (v3 - (v3 * 1))
    var v27 uint64 = 106
    v28 := GenS5{F0: -128}
    var v29 uint = 0
    var v30 uint64 = (^(v27))
    v1 = append(v1, (v2 + (v2 >> ((0) & 63))))
    if (v29 != 0) {
        var v31 int64 = 42
        var v32 uint = (^(v29))
        v32 = 4294967295
        _ = v31
        _ = v32
    }
    switch 127 {
    case -731:
        v3 = 127
    case 32768:
        v29 -= 0
    default:
        var v33 uint16 = v5
        _ = v33
    }
    v34 := GenS6{F0: v30, F1: (v2 % (((v2 * (v2 % ((763) | 1)))) | 1)), F2: (v2 | 65535)}
    v5 -= 255
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v34
    _ = v4
    _ = v5
}
