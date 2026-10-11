package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 string }
type GenS2 struct { F0 int64; F1 string }

func main() {
    var v1 uint = 54
    v2 := []uint32{4294967295, 2147483647, 2147483647}
    var v3 uint8 = 44
    var v4 uint16 = 255
    v3 = 102
    var v5 string = "a"
    v3 = (v3 << ((735) & 63))
    var v6 int16 = -15
    var v7 uint32 = 4294967295
    fmt.Println(v1, v7)
    var v8 uint64 = 709
    var v9 float32 = (float32(1.5) * (1.5 - float32((0.0 - 3.14159))))
    v10 := GenS1{F0: 42, F1: v8, F2: (v5 + "\n\t")}
    v1 = 305
    if false {
        switch int(v8) {
        case 2147483647:
            fmt.Println(v9, v6)
        case 127:
            if ((1e10 + (0.1 - -2.25)) == ((0.1 - 0.0) + ((0.0 * (((float64(v9) * (((1e10 - (0.1 * 1e10)) + (3.14159 / 1e10)) + (0.0 / (-2.25 * 3.14159)))) * -2.25) - 1e10)) + 0.1))) {
                if len(v2) > 0 {
                    v2[0] = (v7 + v7)
                }
            }
        case 32767:
            var v11 int8 = 1
            _ = v11
        default:
            if false {
                switch ((("a" + v5) + "\n\t") + ("hello" + (v5 + "\n\t"))) {
                case "a":
                    _ = len(v2)
                default:
                    v12 := GenS1{F0: (v7 - (v7 * (^(v7)))), F1: v8, F2: v5}
                    _ = v12
                }
            } else {
                _ = len(v2)
                v4 = uint16(v3)
            }
        }
    } else {
        if true {
            fmt.Println(v9, v8)
            v7 = uint32(v1)
        }
        v5 = v5
        v6 = (v6 / ((v6) | 1))
    }
    _ = len(v2)
    fmt.Println(v8, v7)
    for fi13 := 0; fi13 < 10; fi13++ {
        v14 := GenS1{F0: (v7 << ((v1) & 63)), F1: uint64(v1), F2: v5}
        v2 = append(v2, (^(v7)))
        for fi15 := 0; fi15 < 18; fi15++ {
            if (v9 == (v9 / ((-2.25 * (-2.25 + v9)) * 1.5))) {
                fmt.Println(fi15, v1)
                if false {
                    var v16 int8 = -128
                    var v17 float32 = ((v9 + ((v9 - v9) - v9)) - v9)
                    _ = v16
                    _ = v17
                } else {
                    switch v5 {
                    case "hello":
                        v18 := GenS1{F0: (^(v7)), F1: v8, F2: v5}
                        _ = v18
                    case "\n\t":
                        if true {
                            if len(v2) > 0 {
                                v2[0] = v7
                            }
                            var v19 uint16 = 833
                            if len(v2) > 0 {
                                v2[0] = uint32(v8)
                            }
                            _ = v19
                        } else {
                            var v20 float32 = -2.25
                            _ = v20
                        }
                    case "":
                        v21 := GenS1{F0: 0, F1: v8, F2: "a"}
                        _ = v21
                    default:
                        v22 := GenS2{F0: -1, F1: v5}
                        _ = v22
                    }
                    if len(v2) > 0 {
                        v2[0] = 881
                    }
                    var v23 uint16 = (^(v4))
                    _ = v23
                }
                v1 = 4294967295
                if (!false) {
                    var v24 uint8 = 127
                    if true {
                        _ = len(v2)
                        var v25 int16 = int16(v7)
                        _ = v25
                    }
                    v9 = (3.14159 * (-2.25 / (v9 - v9)))
                    _ = v24
                }
            }
            v7 = v7
            var v26 uint = uint(v8)
            _ = fi15
            _ = v26
        }
        _ = fi13
        _ = v14
    }
    var v27 string = v5
    if ((v8 % (((v8 * v8)) | 1)) >= 2) {
        var v28 string = v27
        v3 -= (v3 / ((v3) | 1))
        switch 9223372036854775807 {
        case 597:
            var v29 int = 32767
            _ = v29
        case 276:
            for fi30 := 0; fi30 < 6; fi30++ {
                fmt.Println(v7, v1, v6)
                _ = fi30
            }
        case -795:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
