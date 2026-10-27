package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool }
type GenS2 struct { F0 int8; F1 int32; F2 uint32 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 string }

func main() {
    var v1 float32 = float32(1e10)
    var v2 string = "hello"
    var v3 string = ("foo bar" + "")
    var v4 float32 = (0.1 + 1.5)
    var v5 int = 736
    var v6 bool = (42 != 127)
    v7 := []int{(v5 - (-(v5))), v5, (v5 & 9223372036854775807)}
    var v8 int16 = -192
    for fi9 := 0; fi9 < 19; fi9++ {
        fmt.Println(v3, v8)
        var v10 float32 = v1
        _ = fi9
        _ = v10
    }
    for fi11 := 0; fi11 < 19; fi11++ {
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
        v3 = "a"
        _ = fi11
    }
    v12 := GenS1{F0: 312, F1: v6}
    for fi13 := 0; fi13 < 14; fi13++ {
        if v6 {
            var v14 int8 = -128
            _ = v14
        } else {
            var v15 float64 = (1.5 / ((((3.14159 * 1.5) + 1.5) * 3.14159) / 1.0))
            switch fi13 {
            case -1:
                v16 := GenS1{F0: 2, F1: (32767 <= 2)}
                _ = v16
            case 1:
                fmt.Println(v4)
            case 740:
                var v17 int8 = 127
                _ = v17
            }
            _ = v15
        }
        v6 = ((18446744073709551615 != 141) && true)
        switch v3 {
        case "a":
            if v6 {
                if v6 {
                    switch (v5 / (((fi13 ^ (v5 / ((-23) | 1)))) | 1)) {
                    case 4294967295:
                        v18 := GenS2{F0: 97, F1: -1, F2: 2147483647}
                        _ = v18
                    case -509:
                        for fi19 := 0; fi19 < 12; fi19++ {
                            v20 := GenS2{F0: 42, F1: -2147483648, F2: 42}
                            var v21 uint16 = 65535
                            fmt.Println(v5, v3)
                            _ = fi19
                            _ = v20
                            _ = v21
                        }
                    case -688:
                        if v6 {
                            for fi22 := 0; fi22 < 2; fi22++ {
                                switch 32767 {
                                case 0:
                                    v23 := GenS3{F0: 32767}
                                    _ = v23
                                case 9223372036854775807:
                                    fmt.Println(v6, v1)
                                    if len(v7) >= 2 {
                                        v7 = v7[1:len(v7)]
                                    }
                                    var v24 int64 = -683
                                    _ = v24
                                case -1:
                                    var v25 uint16 = 65535
                                    _ = v25
                                }
                                _ = fi22
                            }
                        } else {
                            if (!v6) {
                                v8 += 1
                                if len(v7) >= 2 {
                                    v7 = v7[1:len(v7)]
                                }
                            } else {
                                if true {
                                    switch 32767 {
                                    case -259:
                                        v7 = append(v7, (-(fi13)))
                                    default:
                                        var v26 int = (-(fi13))
                                        _ = v26
                                    }
                                    v5 = (^(fi13))
                                    v7 = append(v7, (v5 >> ((440) & 63)))
                                } else {
                                    v27 := GenS4{F0: v2}
                                    for fi28 := 0; fi28 < 8; fi28++ {
                                        if (87 != 32) {
                                            var v29 bool = (false || v6)
                                            var v30 uint64 = 2
                                            switch (fi13 | fi28) {
                                            case 1:
                                                v5 = (v5 >> ((255) & 63))
                                            case -125:
                                                v29 = v29
                                                _ = len(v7)
                                                v5 -= fi13
                                                v3 = (v2 + v3)
                                                _ = len(v7)
                                                if v6 {
                                                    var v31 int16 = v8
                                                    _ = v31
                                                }
                                                var v32 int8 = -128
                                                _ = v32
                                            case -1:
                                                var v33 int = (v5 / ((v5) | 1))
                                                _ = v33
                                            default:
                                                v1 += 0.0
                                            }
                                            _ = v29
                                            _ = v30
                                        }
                                        v7 = append(v7, fi28)
                                        _ = fi28
                                    }
                                    _ = v27
                                }
                                fmt.Println(v6)
                                if false {
                                    _ = len(v7)
                                    switch (fi13 / (((fi13 << ((42) & 63))) | 1)) {
                                    case 300:
                                        var v34 bool = (((v2 + "\n\t") + v2) >= "hello")
                                        _ = v34
                                    case 32768:
                                        v35 := GenS4{F0: ("hello" + v3)}
                                        _ = v35
                                    }
                                }
                            }
                            for fi36 := 0; fi36 < 10; fi36++ {
                                var v37 float32 = 1e10
                                _ = fi36
                                _ = v37
                            }
                        }
                    }
                } else {
                    var v38 int64 = 0
                    _ = v38
                }
                v7 = append(v7, fi13)
                v2 = v3
            }
        case "foo bar":
            fmt.Println(v5, v1, v3)
        }
        _ = fi13
    }
    var v39 uint64 = 4294967295
    v1 -= v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
