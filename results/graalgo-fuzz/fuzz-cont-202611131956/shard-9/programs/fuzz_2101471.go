package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32 }
type GenS2 struct { F0 uint64; F1 int8; F2 int64 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int64; F1 int32; F2 int16 }

func main() {
    v1 := make(map[int]string)
    v2 := make(map[int]bool)
    v3 := []string{"hello", (("hello" + "hello") + "a")}
    var v4 float64 = float64(((3.14159 + 1e10) * 3.14159))
    v5 := []int{127, -2147483648, 2147483647}
    v6 := make(map[uint32]string)
    if false {
        var v7 uint16 = 0
        _ = v7
    } else {
        var v8 bool = true
        _ = v8
    }
    v4 += v4
    var v9 int32 = 410
    switch "" {
    case "hello":
        v9 = (v9 % (((-(v9))) | 1))
    default:
        v10 := GenS1{F0: false, F1: 2147483648}
        _ = v10
    }
    var v11 bool = true
    for fi12 := 0; fi12 < 8; fi12++ {
        delete(v2, fi12)
        _ = fi12
    }
    switch (("foo bar" + "") + ("hello" + "hello")) {
    case "foo bar":
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    case "hello":
        v4 -= (float64((1.5 / -2.25)) / 1.0)
    case "\n\t":
        if len(v5) > 0 {
            v5[0] = -673
        }
    default:
        fmt.Println(v9, v11, v11)
    }
    v11 = v11
    v5 = append(v5, 32768)
    if (128 < 1) {
        v11 = (true || v11)
        var v13 int = 9223372036854775807
        _ = v13
    }
    v14 := GenS1{F0: true, F1: 42}
    fmt.Println(v11)
    v15 := GenS2{F0: 1, F1: 0, F2: 9223372036854775807}
    if v11 {
        v16 := GenS2{F0: 301, F1: -128, F2: 9223372036854775807}
        v17 := GenS3{F0: -788}
        var v18 int = 2147483647
        _ = v16
        _ = v17
        _ = v18
    } else {
        v4 += v4
    }
    v19 := v6[2147483648]
    switch -1 {
    case 42:
        for fi20 := 0; fi20 < 20; fi20++ {
            if len(v3) > 0 {
                v3[0] = v19
            }
            for fi21 := 0; fi21 < 13; fi21++ {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                v9 -= (-(v9))
                switch ("" + "\n\t") {
                case "foo bar":
                    fmt.Println(v9, v11)
                case "hello":
                    switch "foo bar" {
                    case "foo bar":
                        var v22 uint16 = 255
                        _ = len(v5)
                        _ = v22
                    }
                    v4 = 1.5
                case "a":
                    var v23 bool = (0 >= 1)
                    _ = v23
                default:
                    v2[int(v9)] = v11
                }
                _ = fi21
            }
            switch ((((v19 + v19) + (v19 + v19)) + v19) + v19) {
            case "\n\t":
                if v11 {
                    if (-1 >= v9) {
                        var v24 int32 = v9
                        _ = v24
                    } else {
                        fmt.Println(fi20)
                    }
                    for fi25 := 0; fi25 < 3; fi25++ {
                        var v26 string = v19
                        fmt.Println(v9, v11, v26)
                        fmt.Println(fi25)
                        _ = fi25
                        _ = v26
                    }
                    for fi27 := 0; fi27 < 18; fi27++ {
                        var v28 uint8 = 2
                        switch v19 {
                        case "":
                            v29 := GenS3{F0: 9223372036854775807}
                            _ = v29
                        case "foo bar":
                            var v30 uint32 = 42
                            _ = v30
                        default:
                            var v31 uint8 = 128
                            _ = v31
                        }
                        var v32 int8 = -1
                        _ = fi27
                        _ = v28
                        _ = v32
                    }
                } else {
                    switch "foo bar" {
                    case "hello":
                        v33 := GenS1{F0: (false && true), F1: 2147483648}
                        _ = v33
                    default:
                        _ = len(v3)
                    }
                    v9 += v9
                }
            case "":
                v34 := GenS2{F0: 2, F1: int8(fi20), F2: 1}
                _ = v34
            case "foo bar":
                var v35 uint16 = 32767
                _ = v35
            default:
                for fi36 := 0; fi36 < 3; fi36++ {
                    var v37 bool = (-1 >= 42)
                    if false {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        fmt.Println(v9, fi20, fi20)
                        v1[-22] = v19
                    } else {
                        var v38 uint64 = 1
                        v39 := GenS4{F0: 1, F1: v9, F2: 1000}
                        _ = v38
                        _ = v39
                    }
                    _ = fi36
                    _ = v37
                }
            }
            _ = fi20
        }
    case -617:
        var v40 uint32 = 2147483647
        _ = v40
    default:
        var v41 uint8 = 127
        _ = v41
    }
    if len(v3) > 0 {
        v3[0] = "foo bar"
    }
    var v42 uint = 845
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v5
    _ = v6
    _ = v9
}
