package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16 }
type GenS2 struct { F0 int; F1 uint64 }
type GenS3 struct { F0 int; F1 int }
type GenS4 struct { F0 string; F1 uint32; F2 uint64 }
type GenS5 struct { F0 int8; F1 uint8 }

func main() {
    v1 := make(map[string]uint32)
    var v2 uint64 = 1
    var v3 bool = false
    var v4 int16 = -1
    v5 := make(map[string]bool)
    var v6 uint = 18446744073709551615
    v7 := []int{127, -1}
    fmt.Println(v3, v6)
    var v8 int8 = 0
    if (1 != 1) {
        switch 57 {
        case -1:
            fmt.Println(v8, v3, v3)
        case 422:
            var v9 float32 = 1.5
            _ = v9
        }
        v10 := v5["a"]
        _ = v10
    } else {
        if v3 {
            var v11 uint8 = 1
            v12 := GenS1{F0: 127, F1: 255}
            _ = v11
            _ = v12
        } else {
            fmt.Println(v6, v8, v8)
            if v3 {
                if v3 {
                    v7 = append(v7, -2147483648)
                    v2 = v2
                    v13 := GenS1{F0: 32768, F1: 255}
                    _ = v13
                } else {
                    v8 += (v8 << (((v6 << (((v6 * v6)) & 63))) & 63))
                }
                v2 = 705
            }
        }
    }
    for fi14 := 0; fi14 < 14; fi14++ {
        v15 := GenS2{F0: (fi14 * (fi14 | fi14)), F1: 1}
        switch "\n\t" {
        case "a":
            v3 = ((v4 >> (((v6 & v6)) & 63)) <= (v4 ^ v4))
        }
        _ = fi14
        _ = v15
    }
    v1[("\n\t" + "a")] = 65535
    v7 = append(v7, 128)
    var v16 uint = v6
    v17 := GenS3{F0: -9223372036854775808, F1: 2147483647}
    if (v4 < -32768) {
        if (!(!v3)) {
            switch "hello" {
            case "\n\t":
                delete(v5, "foo bar")
            case "":
                fmt.Println(v8)
            case "a":
                v16 -= (v16 >> ((v6) & 63))
            }
            v3 = (0.0 < 3.14159)
        } else {
            v18 := GenS2{F0: -1, F1: (v2 << ((2) & 63))}
            v3 = (0 > 347)
            delete(v1, ("\n\t" + "\n\t"))
            _ = v18
        }
        v1["foo bar"] = 42
        var v19 float32 = (float32(-2.25) * 0.0)
        _ = v19
    } else {
        for fi20 := 0; fi20 < 20; fi20++ {
            fmt.Println(v4, v4, v6)
            fmt.Println(v8, v2, v4)
            var v21 bool = (!true)
            _ = fi20
            _ = v21
        }
        v7 = append(v7, 125)
        var v22 uint32 = 65535
        _ = v22
    }
    if false {
        switch "" {
        case "\n\t":
            if v3 {
                fmt.Println(v4)
                if v3 {
                    switch 300 {
                    case 127:
                        var v23 string = "hello"
                        _ = v23
                    case 1:
                        for fi24 := 0; fi24 < 10; fi24++ {
                            v25 := GenS4{F0: ((("hello" + ("foo bar" + "")) + "") + ("foo bar" + ("foo bar" + ("a" + "")))), F1: 273, F2: 0}
                            if ((int(v2) > (^(fi24))) || (!(uint64(v6) != v2))) {
                                for fi26 := 0; fi26 < 11; fi26++ {
                                    v4 += v4
                                    var v27 int16 = (v4 * v4)
                                    _ = fi26
                                    _ = v27
                                }
                            } else {
                                if true {
                                    v2 = (v2 - v2)
                                } else {
                                    v3 = false
                                    fmt.Println(v16, fi24)
                                    v4 -= int16(v8)
                                }
                                for fi28 := 0; fi28 < 6; fi28++ {
                                    if len(v7) >= 2 {
                                        v7 = v7[1:len(v7)]
                                    }
                                    var v29 uint64 = v2
                                    _ = fi28
                                    _ = v29
                                }
                                v30 := GenS3{F0: 42, F1: fi24}
                                _ = v30
                            }
                            _ = fi24
                            _ = v25
                        }
                    default:
                        var v31 uint32 = 42
                        _ = v31
                    }
                    v32 := GenS5{F0: (-(v8)), F1: 255}
                    v4 = v4
                    _ = v32
                } else {
                    v3 = (("foo bar" + "\n\t") < ((("\n\t" + "") + "\n\t") + "hello"))
                }
                for fi33 := 0; fi33 < 10; fi33++ {
                    fmt.Println(v3, v8, v3)
                    fmt.Println(v6)
                    _ = fi33
                }
            } else {
                v2 = v2
                if len(v7) > 0 {
                    v7[0] = -299
                }
                for fi34 := 0; fi34 < 5; fi34++ {
                    v35 := GenS4{F0: "\n\t", F1: 2147483648, F2: 2}
                    switch (fi34 ^ fi34) {
                    case 42:
                        var v36 int8 = 42
                        _ = v36
                    case 4294967295:
                        switch ("foo bar" + "\n\t") {
                        case "foo bar":
                            for fi37 := 0; fi37 < 15; fi37++ {
                                v16 -= 548
                                fmt.Println(v2, v8, fi37)
                                _ = fi37
                            }
                        case "\n\t":
                            delete(v1, ("" + "hello"))
                        case "hello":
                            v38 := GenS2{F0: (fi34 % (((fi34 << (((v6 >> (((v16 | 925)) & 63))) & 63))) | 1)), F1: v2}
                            _ = v38
                        default:
                            v4 = (v4 * 206)
                        }
                    case 9223372036854775807:
                        if len(v7) > 0 {
                            v7[0] = (fi34 ^ 638)
                        }
                    default:
                        var v39 int64 = -450
                        _ = v39
                    }
                    _ = fi34
                    _ = v35
                }
            }
        case "":
            v6 = v16
        default:
            if len(v7) > 0 {
                v7[0] = 9223372036854775807
            }
        }
    }
    var v40 int32 = -1
    if ((4294967295 == v2) || v3) {
        switch "foo bar" {
        case "foo bar":
            v41 := v1["hello"]
            _ = v41
        default:
            v42 := GenS4{F0: "hello", F1: 65535, F2: v2}
            _ = v42
        }
        var v43 int8 = 98
        _ = v43
    }
    v3 = false
    for fi44 := 0; fi44 < 15; fi44++ {
        v16 = 255
        _ = fi44
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
