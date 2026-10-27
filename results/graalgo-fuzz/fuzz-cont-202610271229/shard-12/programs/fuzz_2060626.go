package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32 }
type GenS2 struct { F0 uint32; F1 int8 }
type GenS3 struct { F0 bool; F1 uint }
type GenS4 struct { F0 int8; F1 uint32; F2 bool }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int8; F1 int16; F2 bool }
type GenS7 struct { F0 string }
type GenS8 struct { F0 int16; F1 uint; F2 int }
type GenS9 struct { F0 int }

func main() {
    var v1 float64 = float64(0.1)
    var v2 bool = (0 <= 4294967295)
    var v3 uint32 = 42
    v4 := make(map[uint32]uint32)
    for fi5 := 0; fi5 < 9; fi5++ {
        fmt.Println(v2, v1, v1)
        v1 = 1.5
        _ = fi5
    }
    for fi6 := 0; fi6 < 11; fi6++ {
        v1 = v1
        if true {
            v2 = v2
        }
        _ = fi6
    }
    v7 := GenS1{F0: -1, F1: 556}
    var v8 uint64 = 4294967295
    v9 := []int{9223372036854775807, 777, 925, 4294967295}
    v10 := GenS1{F0: 0, F1: -684}
    fmt.Println(v1)
    for fi11 := 0; fi11 < 20; fi11++ {
        v12 := GenS2{F0: 2, F1: 127}
        var v13 int8 = 29
        _ = fi11
        _ = v12
        _ = v13
    }
    v14 := GenS3{F0: (467 == 2147483647), F1: 1}
    v1 += 0.1
    v15 := GenS3{F0: true, F1: 2}
    if v2 {
        v3 -= v3
    } else {
        if false {
            v16 := GenS4{F0: 29, F1: v3, F2: (-32768 <= 1000)}
            _ = v16
        }
    }
    fmt.Println(v2)
    for fi17 := 0; fi17 < 5; fi17++ {
        var v18 string = ((("" + "a") + ("\n\t" + "foo bar")) + "")
        _ = fi17
        _ = v18
    }
    if (v2 && (161 != 0)) {
        if (!("\n\t" >= "hello")) {
            fmt.Println(v8, v2)
            for fi19 := 0; fi19 < 14; fi19++ {
                v2 = ((fi19 + fi19) >= fi19)
                for fi20 := 0; fi20 < 2; fi20++ {
                    switch ("a" + "foo bar") {
                    case "":
                        v21 := GenS3{F0: v2, F1: 4294967295}
                        _ = v21
                    case "hello":
                        v22 := GenS5{F0: "\n\t"}
                        _ = v22
                    case "foo bar":
                        if v2 {
                            v23 := v4[v3]
                            _ = v23
                        }
                    default:
                        if ((1.5 - (-2.25 * 0.1)) != (1.5 + (((0.1 - float32(v1)) - (-2.25 / 1.0)) / 0.1))) {
                            if len(v9) > 0 {
                                v9[0] = 32767
                            }
                        } else {
                            v24 := GenS6{F0: -128, F1: 0, F2: v2}
                            v8 += 552
                            _ = v24
                        }
                    }
                    delete(v4, 42)
                    _ = fi20
                }
                _ = fi19
            }
            if true {
                for fi25 := 0; fi25 < 12; fi25++ {
                    v26 := GenS7{F0: ("a" + "foo bar")}
                    var v27 int64 = -892
                    v28 := GenS8{F0: -32768, F1: 907, F2: (fi25 >> ((65535) & 63))}
                    _ = fi25
                    _ = v26
                    _ = v27
                    _ = v28
                }
                v8 += (^(v8))
            } else {
                v3 = v3
                var v29 uint = uint(v8)
                _ = v29
            }
        } else {
            fmt.Println(v2, v8, v1)
        }
    } else {
        var v30 uint8 = 255
        for fi31 := 0; fi31 < 9; fi31++ {
            _ = len(v4)
            v3 += 2147483647
            _ = fi31
        }
        var v32 string = "a"
        _ = v30
        _ = v32
    }
    switch 4294967295 {
    case 9223372036854775807:
        v1 -= v1
    case 128:
        if (1 < uint8(v8)) {
            _ = len(v9)
            if len(v9) >= 2 {
                v9 = v9[1:len(v9)]
            }
            v2 = v2
        }
    default:
        for fi33 := 0; fi33 < 1; fi33++ {
            switch "hello" {
            case "foo bar":
                v34 := GenS9{F0: 32768}
                _ = v34
            case "":
                if v2 {
                    v1 = v1
                    for fi35 := 0; fi35 < 5; fi35++ {
                        v36 := GenS2{F0: 65535, F1: 1}
                        _ = fi35
                        _ = v36
                    }
                }
            }
            _ = fi33
        }
    }
    var v37 uint32 = v3
    _ = len(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
