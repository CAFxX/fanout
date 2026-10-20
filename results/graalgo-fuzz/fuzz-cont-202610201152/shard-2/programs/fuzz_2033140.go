package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8; F2 int16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint32; F1 uint }
type GenS4 struct { F0 string; F1 uint16; F2 int32 }
type GenS5 struct { F0 int16; F1 uint16 }

func main() {
    v1 := make(map[int]string)
    var v2 bool = false
    var v3 float32 = float32(((3.14159 / (3.14159 * 1.5)) / 1.5))
    v4 := make(map[int]string)
    var v5 int16 = -1
    var v6 string = "\n\t"
    v2 = (v2 && (255 > 2))
    v1[2147483647] = (("a" + "\n\t") + "a")
    v5 += v5
    var v7 uint16 = 0
    var v8 uint32 = 767
    for fi9 := 0; fi9 < 9; fi9++ {
        switch (fi9 / (((fi9 * (fi9 * fi9))) | 1)) {
        case -912:
            for fi10 := 0; fi10 < 20; fi10++ {
                var v11 uint32 = 1
                v12 := GenS1{F0: v2, F1: 209, F2: (v5 * 0)}
                _ = fi10
                _ = v11
                _ = v12
            }
        case -2147483648:
            for fi13 := 0; fi13 < 12; fi13++ {
                if v2 {
                    _ = len(v4)
                    v14 := GenS1{F0: (42 == 153), F1: 1, F2: -1}
                    if false {
                        v8 = v8
                        var v15 int16 = -131
                        _ = v15
                    }
                    _ = v14
                } else {
                    v16 := GenS2{F0: (v6 + "a")}
                    v7 = (v7 % (((^(v7))) | 1))
                    _ = v16
                }
                delete(v1, 691)
                _ = fi13
            }
        case 9223372036854775807:
            var v17 uint64 = 1
            _ = v17
        default:
            v18 := GenS3{F0: (v8 - v8), F1: 762}
            _ = v18
        }
        _ = fi9
    }
    switch 9223372036854775807 {
    case -2147483648:
        var v19 uint = 18446744073709551615
        _ = v19
    case 127:
        v20 := GenS4{F0: v6, F1: 32767, F2: -454}
        _ = v20
    default:
        v2 = false
    }
    v1[-968] = (v6 + v6)
    v21 := []int{-430, 42, -1}
    switch v6 {
    case "\n\t":
        switch -990 {
        case 9223372036854775807:
            if v2 {
                v3 += (0.0 / 1.0)
                v1[1] = "foo bar"
                switch 32767 {
                case 127:
                    v2 = (-9223372036854775808 <= 2147483647)
                default:
                    if false {
                        v6 = v6
                        v1[860] = "foo bar"
                    } else {
                        v2 = v2
                        if (1 != 0) {
                            switch 127 {
                            case -364:
                                for fi22 := 0; fi22 < 5; fi22++ {
                                    v1[-385] = (v6 + "a")
                                    _ = fi22
                                }
                            case 2147483647:
                                v23 := GenS1{F0: true, F1: 255, F2: v5}
                                _ = v23
                            default:
                                if (v6 < v6) {
                                    var v24 uint8 = 163
                                    v24 = (v24 / ((uint8(v7)) | 1))
                                    _ = v24
                                }
                            }
                        }
                    }
                }
            } else {
                delete(v4, -9223372036854775808)
            }
        case 42:
            v4[9223372036854775807] = ("foo bar" + (v6 + ""))
        }
    case "hello":
        switch v6 {
        case "a":
            var v25 uint64 = 18446744073709551615
            _ = v25
        default:
            v26 := GenS5{F0: v5, F1: (^(v7))}
            _ = v26
        }
    case "foo bar":
        var v27 float32 = v3
        _ = v27
    default:
        var v28 int = 127
        _ = v28
    }
    fmt.Println(v2)
    v8 -= 42
    var v29 float32 = (v3 - v3)
    var v30 int32 = 1
    if v2 {
        for fi31 := 0; fi31 < 3; fi31++ {
            v30 += 807
            v29 = v29
            _ = fi31
        }
    } else {
        v32 := GenS2{F0: "\n\t"}
        if v2 {
            v30 = v30
            fmt.Println(v8, v29)
            fmt.Println(v8, v7, v7)
        } else {
            for fi33 := 0; fi33 < 2; fi33++ {
                var v34 uint64 = 0
                _ = fi33
                _ = v34
            }
        }
        var v35 uint16 = 255
        _ = v32
        _ = v35
    }
    v3 = (((3.14159 - 1e10) * ((v3 / 1.0) / ((float32(0.0) * (0.1 + ((v3 + (v3 - v29)) - (1.5 + -2.25)))) - (0.1 + 0.1)))) * (v29 * (v3 * v3)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v21
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
