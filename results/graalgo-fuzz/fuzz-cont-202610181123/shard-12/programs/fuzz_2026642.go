package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int; F1 bool; F2 string }
type GenS4 struct { F0 int64; F1 uint8 }
type GenS5 struct { F0 int32 }

func main() {
    var v1 uint = 602
    var v2 int64 = 0
    var v3 uint32 = 189
    v4 := make(map[string]bool)
    v5 := make(map[int]string)
    var v6 float64 = 3.14159
    v2 -= v2
    var v7 uint8 = 255
    switch 498 {
    case 149:
        if true {
            v7 -= (v7 >> (((^(v1))) & 63))
            var v8 uint16 = 32767
            for fi9 := 0; fi9 < 12; fi9++ {
                v6 -= 0.0
                _ = fi9
            }
            _ = v8
        } else {
            var v10 int = -2147483648
            v11 := []int{v10}
            v1 = (v1 * (^(v1)))
            _ = v10
            _ = v11
        }
    case -452:
        fmt.Println(v3)
    default:
        var v12 int32 = 42
        _ = v12
    }
    v13 := []int{2147483647}
    fmt.Println(v3, v2, v7)
    var v14 float64 = (float64(0.1) / v6)
    v4[("\n\t" + "")] = (!(!false))
    switch ("" + "hello") {
    case "foo bar":
        if (false || (32767 > 72)) {
            fmt.Println(v1, v2)
        }
    default:
        var v15 int = 42
        _ = v15
    }
    v14 -= float64(3.14159)
    for fi16 := 0; fi16 < 6; fi16++ {
        fmt.Println(v2, v1)
        if (v3 == (v3 + uint32(v1))) {
            if (-128 < -1) {
                var v17 int32 = 977
                _ = v17
            } else {
                v3 -= (v3 + v3)
                if (false || (true && false)) {
                    var v18 float32 = 0.1
                    v19 := GenS1{F0: 2, F1: 100}
                    v13 = append(v13, (fi16 ^ -9223372036854775808))
                    _ = v18
                    _ = v19
                }
                for fi20 := 0; fi20 < 7; fi20++ {
                    v21 := GenS2{F0: (-(fi20))}
                    _ = fi20
                    _ = v21
                }
            }
        } else {
            for fi22 := 0; fi22 < 19; fi22++ {
                if true {
                    delete(v5, (^(fi16)))
                    switch "" {
                    case "":
                        v3 = (v3 - 2)
                    default:
                        if true {
                            v2 += (-(v2))
                            var v23 int16 = -32768
                            _ = v23
                        } else {
                            if len(v13) > 0 {
                                v13[0] = fi22
                            }
                            switch "\n\t" {
                            case "\n\t":
                                v3 += 1
                            case "":
                                switch -2147483648 {
                                case 32768:
                                    var v24 int8 = -77
                                    _ = v24
                                }
                            case "foo bar":
                                if true {
                                    fmt.Println(fi22, v1)
                                    if false {
                                        v25 := GenS2{F0: fi16}
                                        switch (("hello" + "\n\t") + "\n\t") {
                                        case "":
                                            v13 = append(v13, int(v2))
                                        }
                                        if len(v13) > 0 {
                                            v13[0] = -2147483648
                                        }
                                        _ = len(v4)
                                        _ = v25
                                    }
                                    if (-1 <= 0) {
                                        if false {
                                            v26 := GenS3{F0: (fi22 + (fi16 * -9223372036854775808)), F1: false, F2: ""}
                                            if len(v13) > 0 {
                                                v13[0] = (fi16 << (((v1 << (((v1 << ((v1) & 63))) & 63))) & 63))
                                            }
                                            _ = v26
                                        }
                                        v14 = ((-2.25 * float64(float32(v6))) * float64(1.5))
                                    } else {
                                        var v27 bool = true
                                        fmt.Println(v7)
                                        switch ("foo bar" + (("hello" + "hello") + "\n\t")) {
                                        case "\n\t":
                                            v3 -= v3
                                        case "hello":
                                            v2 = v2
                                        }
                                        _ = v27
                                    }
                                }
                            }
                        }
                        var v28 uint = uint(v7)
                        if (false || (true && false)) {
                            var v29 float32 = 3.14159
                            v30 := GenS2{F0: (fi22 & fi22)}
                            fmt.Println(v14, v1)
                            _ = v29
                            _ = v30
                        } else {
                            var v31 float64 = v14
                            var v32 uint64 = 18446744073709551615
                            var v33 uint64 = (^(v32))
                            _ = v31
                            _ = v32
                            _ = v33
                        }
                        _ = v28
                    }
                }
                v1 = (^(v1))
                if false {
                    delete(v4, ("\n\t" + ""))
                    v14 = v6
                    fmt.Println(v3)
                } else {
                    var v34 uint32 = v3
                    v35 := GenS4{F0: 9223372036854775807, F1: 255}
                    _ = v34
                    _ = v35
                }
                _ = fi22
            }
            v2 -= (v2 * v2)
            v7 = v7
        }
        _ = fi16
    }
    switch 2147483647 {
    case -52:
        var v36 int = 0
        _ = v36
    case -815:
        var v37 uint64 = 42
        _ = v37
    case 2147483647:
        delete(v5, -930)
    }
    for fi38 := 0; fi38 < 5; fi38++ {
        v3 += (v3 >> ((4294967295) & 63))
        _ = fi38
    }
    v39 := GenS5{F0: -618}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
