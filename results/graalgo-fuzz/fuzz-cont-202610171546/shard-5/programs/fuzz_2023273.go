package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32; F2 int32 }
type GenS2 struct { F0 int16; F1 int8 }
type GenS3 struct { F0 int32; F1 uint; F2 uint16 }

func main() {
    var v1 int16 = -28
    var v2 int32 = 185
    v3 := make(map[string]uint32)
    v4 := make(map[uint32]string)
    var v5 uint16 = 32767
    v6 := make(map[uint32]bool)
    v1 += v1
    v7 := []int{32767}
    var v8 uint32 = 4294967295
    _ = len(v3)
    switch (("foo bar" + "") + "hello") {
    case "hello":
        if false {
            var v9 int8 = 42
            if (!(!true)) {
                if (!false) {
                    v5 = v5
                    var v10 int8 = -81
                    _ = v10
                } else {
                    if (-2.25 > (float32(0.1) * 3.14159)) {
                        v11 := GenS1{F0: 1, F1: (v2 << ((65535) & 63)), F2: (v2 / (((v2 >> ((273) & 63))) | 1))}
                        var v12 uint16 = 645
                        var v13 int16 = 0
                        _ = v11
                        _ = v12
                        _ = v13
                    } else {
                        v8 = v8
                        for fi14 := 0; fi14 < 6; fi14++ {
                            v15 := GenS1{F0: 1, F1: (-(v2)), F2: v2}
                            var v16 uint = 1
                            _ = fi14
                            _ = v15
                            _ = v16
                        }
                    }
                }
            } else {
                v8 += 2
            }
            _ = v9
        }
    case "\n\t":
        v2 = v2
    }
    v5 += uint16(v8)
    fmt.Println(v8, v1)
    v17 := GenS1{F0: 42, F1: (v2 / (((v2 + (^(v2)))) | 1)), F2: (v2 % (((v2 - v2)) | 1))}
    v2 = 0
    var v18 uint64 = 539
    if true {
        switch 0 {
        case -949:
            v1 = v1
        case 634:
            var v19 bool = true
            _ = v19
        default:
            for fi20 := 0; fi20 < 5; fi20++ {
                v1 = (v1 ^ v1)
                if true {
                    var v21 int8 = 109
                    v22 := GenS2{F0: v1, F1: 127}
                    var v23 int32 = (v2 / ((v2) | 1))
                    _ = v21
                    _ = v22
                    _ = v23
                }
                v24 := GenS3{F0: 1, F1: 4294967295, F2: 255}
                _ = fi20
                _ = v24
            }
        }
        var v25 int32 = 0
        _ = v25
    }
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    var v26 uint16 = v5
    _ = len(v6)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
