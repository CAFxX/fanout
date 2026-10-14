package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 int32 }
type GenS2 struct { F0 int32; F1 int; F2 int64 }
type GenS3 struct { F0 uint8; F1 uint32; F2 int32 }
type GenS4 struct { F0 uint64 }

func main() {
    var v1 int = -1
    var v2 float32 = 1e10
    v3 := make(map[string]uint32)
    var v4 string = ("a" + ("a" + "foo bar"))
    v5 := GenS1{F0: 66, F1: -1, F2: -2147483648}
    v3[v4] = 2
    for fi6 := 0; fi6 < 17; fi6++ {
        if true {
            var v7 uint16 = 255
            _ = v7
        } else {
            fmt.Println(fi6)
            for fi8 := 0; fi8 < 9; fi8++ {
                switch "foo bar" {
                case "a":
                    v9 := GenS2{F0: 1, F1: -2147483648, F2: -9223372036854775808}
                    _ = v9
                case "foo bar":
                    delete(v3, (v4 + v4))
                    v10 := GenS1{F0: 2, F1: 0, F2: 651}
                    if (65535 != 137) {
                        fmt.Println(fi8, fi8, fi6)
                    }
                    _ = v10
                case "hello":
                    switch fi6 {
                    case -2147483648:
                        v2 += (0.0 / (float32(1e10) / (0.1 * -2.25)))
                        v11 := GenS2{F0: 42, F1: v1, F2: -9223372036854775808}
                        fmt.Println(v2, v1, v1)
                        v12 := GenS1{F0: 213, F1: 361, F2: -1}
                        _ = v11
                        _ = v12
                    case -97:
                        v3[(("foo bar" + "") + "\n\t")] = 1
                    case 1:
                        fmt.Println(v2)
                    }
                }
                _ = fi8
            }
            v1 -= 1
        }
        _ = fi6
    }
    switch 1 {
    case 1:
        for fi13 := 0; fi13 < 8; fi13++ {
            var v14 int16 = 684
            v2 -= float32(1.5)
            _ = fi13
            _ = v14
        }
    }
    var v15 uint64 = 894
    switch (v1 ^ (v1 - v1)) {
    case -2147483648:
        if false {
            var v16 float32 = v2
            for fi17 := 0; fi17 < 16; fi17++ {
                switch "" {
                case "\n\t":
                    for fi18 := 0; fi18 < 19; fi18++ {
                        v2 = 1e10
                        v19 := v3[v4]
                        _ = fi18
                        _ = v19
                    }
                case "hello":
                    if false {
                        for fi20 := 0; fi20 < 1; fi20++ {
                            v2 = (v2 * 1.5)
                            _ = fi20
                        }
                        var v21 float64 = 1.5
                        switch fi17 {
                        case 128:
                            var v22 int32 = -1
                            _ = v22
                        case 507:
                            fmt.Println(fi17, v4, v4)
                        case -9223372036854775808:
                            v15 = v15
                        }
                        _ = v21
                    } else {
                        v2 -= 0.1
                        var v23 string = (v4 + ("foo bar" + v4))
                        _ = v23
                    }
                case "foo bar":
                    v1 = (-(v1))
                default:
                    v24 := GenS2{F0: -452, F1: (v1 & (fi17 + fi17)), F2: 42}
                    _ = v24
                }
                for fi25 := 0; fi25 < 7; fi25++ {
                    if ((1.5 > 1.5) || (true || (fi17 <= (fi25 ^ 9223372036854775807)))) {
                        v4 = "\n\t"
                        v26 := v3[""]
                        v1 += fi25
                        v16 -= v16
                        _ = v26
                    } else {
                        v2 += 3.14159
                        if (!(uint8(v15) >= 1)) {
                            var v27 string = ""
                            var v28 int64 = 9223372036854775807
                            v16 = (v2 / (v2 - 0.1))
                            v4 = (v27 + v4)
                            _ = v27
                            _ = v28
                        } else {
                            v2 = (1.5 + float32(float64(0.1)))
                            v29 := GenS2{F0: 348, F1: fi17, F2: 42}
                            _ = v29
                        }
                        var v30 int64 = 9223372036854775807
                        var v31 uint64 = (v15 >> ((255) & 63))
                        var v32 bool = (!(!(true || true)))
                        _ = v30
                        _ = v31
                        _ = v32
                    }
                    _ = fi25
                }
                v15 = (v15 - (v15 >> ((893) & 63)))
                _ = fi17
            }
            _ = v16
        } else {
            var v33 uint64 = 240
            fmt.Println(v1, v2, v33)
            _ = len(v3)
            _ = v33
        }
    default:
        var v34 uint8 = 0
        _ = v34
    }
    var v35 uint64 = 4294967295
    switch "foo bar" {
    case "\n\t":
        v1 = v1
    default:
        for fi36 := 0; fi36 < 16; fi36++ {
            fmt.Println(v2)
            var v37 int8 = 0
            v38 := []int{(v1 - fi36), fi36, (v1 & int(v37))}
            _ = fi36
            _ = v37
            _ = v38
        }
    }
    v39 := []int{32768, v1, (-(v1)), v1}
    for fi40 := 0; fi40 < 14; fi40++ {
        v41 := GenS3{F0: 42, F1: 42, F2: -631}
        _ = fi40
        _ = v41
    }
    v35 = 18446744073709551615
    v1 = (v1 * (v1 / (((v1 * (v1 & (^(v1))))) | 1)))
    var v42 float64 = 0.1
    switch v4 {
    case "\n\t":
        v43 := GenS1{F0: 42, F1: 581, F2: -276}
        _ = v43
    case "":
        if (-764 >= 929) {
            v44 := GenS1{F0: 176, F1: 2147483647, F2: 1}
            _ = v44
        } else {
            fmt.Println(v2, v35)
            var v45 float64 = v42
            var v46 uint64 = v35
            _ = v45
            _ = v46
        }
    case "foo bar":
        if len(v39) >= 2 {
            v39 = v39[1:len(v39)]
        }
    }
    if len(v39) > 0 {
        v39[0] = (v1 / (((v1 & 426)) | 1))
    }
    fmt.Println(v35, v42, v2)
    v47 := GenS4{F0: (v15 - v35)}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v39:", len(v39), v39)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v35
    _ = v39
    _ = v4
    _ = v42
    _ = v47
    _ = v5
}
