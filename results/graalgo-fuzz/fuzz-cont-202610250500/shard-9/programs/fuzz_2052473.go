package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 string; F1 int32; F2 bool }
type GenS3 struct { F0 string; F1 uint16; F2 int }
type GenS4 struct { F0 uint16 }

func main() {
    var v1 int8 = 127
    v2 := make(map[uint32]string)
    v3 := make(map[string]bool)
    var v4 bool = (0.0 >= 0.0)
    v5 := make(map[uint32]string)
    var v6 uint16 = 2
    v7 := []int{559, -1, 9223372036854775807}
    if (-1 != -528) {
        if ((0 <= 905) && (-1 != 0)) {
            var v8 float32 = 1.5
            for fi9 := 0; fi9 < 11; fi9++ {
                v4 = v4
                v8 -= ((v8 * ((((float32((1.5 - 3.14159)) + 3.14159) + v8) + 1.5) / ((v8 * ((((v8 * float32((-2.25 - 3.14159))) + 0.0) / 3.14159) - (v8 * v8))) / v8))) * v8)
                v10 := GenS1{F0: 825}
                _ = fi9
                _ = v10
            }
            _ = v8
        } else {
            v6 -= v6
            var v11 string = "\n\t"
            var v12 uint8 = 2
            _ = v11
            _ = v12
        }
        v4 = true
        for fi13 := 0; fi13 < 10; fi13++ {
            switch ((((("" + ("foo bar" + "foo bar")) + ("hello" + "hello")) + "hello") + ("a" + "foo bar")) + "") {
            case "foo bar":
                v1 -= -77
            case "":
                for fi14 := 0; fi14 < 18; fi14++ {
                    if ((9223372036854775807 > 42) && v4) {
                        v6 += (v6 & (v6 | 255))
                        var v15 int64 = -9223372036854775808
                        if false {
                            if v4 {
                                v16 := GenS2{F0: "hello", F1: 0, F2: v4}
                                switch 932 {
                                case -2147483648:
                                    var v17 int32 = -1
                                    _ = v17
                                }
                                _ = v16
                            }
                            if v4 {
                                if (0.0 != ((((0.1 / 3.14159) * float64(3.14159)) - float64(float32(float64(float32(3.14159))))) / ((-2.25 / 3.14159) * 0.1))) {
                                    if v4 {
                                        if false {
                                            if len(v7) > 0 {
                                                v7[0] = 2147483647
                                            }
                                        }
                                    }
                                    v18 := GenS2{F0: "foo bar", F1: -343, F2: v4}
                                    switch "\n\t" {
                                    case "":
                                        _ = len(v7)
                                    default:
                                        var v19 string = "foo bar"
                                        _ = v19
                                    }
                                    _ = v18
                                }
                            }
                        } else {
                            var v20 uint16 = v6
                            v21 := GenS3{F0: "\n\t", F1: v20, F2: fi14}
                            _ = v20
                            _ = v21
                        }
                        _ = v15
                    }
                    _ = fi14
                }
            case "hello":
                var v22 uint64 = 1
                _ = v22
            }
            _ = fi13
        }
    }
    var v23 float32 = float32((3.14159 - (((0.0 * (0.1 / (-2.25 / 1e10))) + 1e10) + ((1e10 / 1.5) + 0.0))))
    v24 := GenS4{F0: (v6 << ((255) & 63))}
    var v25 uint = 398
    for fi26 := 0; fi26 < 18; fi26++ {
        var v27 uint = 1
        _ = fi26
        _ = v27
    }
    fmt.Println(v6)
    var v28 int64 = 9223372036854775807
    v25 = (v25 << (((v25 & v25)) & 63))
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    v1 = (v1 ^ v1)
    var v29 uint16 = v6
    v28 -= v28
    v6 = 255
    delete(v3, "foo bar")
    v1 -= (v1 >> ((v25) & 63))
    v30 := GenS3{F0: "\n\t", F1: (v6 + v29), F2: -467}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
