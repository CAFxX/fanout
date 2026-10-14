package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int }

func main() {
    var v1 int = -2147483648
    v2 := make(map[string]bool)
    var v3 uint32 = 0
    v4 := make(map[int]bool)
    v5 := make(map[int]uint32)
    v6 := make(map[string]int)
    v7 := []int{(-(v1))}
    var v8 float32 = -2.25
    for fi9 := 0; fi9 < 15; fi9++ {
        var v10 uint16 = 214
        if false {
            var v11 string = "\n\t"
            for fi12 := 0; fi12 < 1; fi12++ {
                v11 = "hello"
                var v13 uint16 = (^(v10))
                var v14 float32 = v8
                var v15 int64 = -1
                if (110 > 127) {
                    if true {
                        v8 = ((v8 * v14) + float32(-2.25))
                    } else {
                        switch (fi9 & (v1 << ((777) & 63))) {
                        case 670:
                            switch v1 {
                            case -369:
                                var v16 int8 = -128
                                _ = v16
                            case -753:
                                fmt.Println(v15, fi12, v1)
                                var v17 uint64 = 2
                                switch v1 {
                                case 42:
                                    if true {
                                        if false {
                                            v18 := GenS1{F0: -331, F1: 1}
                                            var v19 uint16 = v10
                                            if len(v7) >= 2 {
                                                v7 = v7[1:len(v7)]
                                            }
                                            _ = v18
                                            _ = v19
                                        } else {
                                            var v20 int8 = 1
                                            fmt.Println(v20, v13)
                                            var v21 uint16 = (v10 / ((v10) | 1))
                                            _ = v20
                                            _ = v21
                                        }
                                    }
                                    var v22 uint8 = uint8(v3)
                                    var v23 uint = 0
                                    _ = v22
                                    _ = v23
                                }
                                v11 = v11
                                fmt.Println(v8, v8)
                                fmt.Println(fi12, v1, v15)
                                _ = v17
                            }
                        case 0:
                            v11 = v11
                        default:
                            v1 += (fi12 % (((fi9 * fi12)) | 1))
                        }
                        if true {
                            v2["a"] = false
                        }
                        v24 := GenS1{F0: -1, F1: 57}
                        var v25 float64 = 1.5
                        _ = v24
                        _ = v25
                    }
                }
                _ = fi12
                _ = v13
                _ = v14
                _ = v15
            }
            _ = v11
        }
        _ = fi9
        _ = v10
    }
    v26 := GenS2{F0: 100}
    var v27 int8 = 0
    v28 := GenS3{F0: v1}
    if len(v7) > 0 {
        v7[0] = v1
    }
    v27 = 12
    var v29 string = ""
    switch v29 {
    case "hello":
        for fi30 := 0; fi30 < 8; fi30++ {
            if (42 < -2147483648) {
                var v31 int32 = 0
                _ = v31
            }
            _ = fi30
        }
    }
    _ = len(v6)
    v1 = (v1 ^ 535)
    v32 := GenS2{F0: 1}
    v8 += v8
    var v33 int = (v1 >> ((150) & 63))
    if (true && (1 <= -1)) {
        v29 = v29
        v3 = 42
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
