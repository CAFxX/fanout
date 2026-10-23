package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint32; F1 bool; F2 int8 }
type GenS4 struct { F0 int8; F1 uint16; F2 string }

func main() {
    var v1 float64 = 1e10
    v2 := make(map[int]int)
    v3 := make(map[int]int)
    var v4 uint32 = 2
    for fi5 := 0; fi5 < 1; fi5++ {
        for fi6 := 0; fi6 < 14; fi6++ {
            v7 := []int{(-(fi6)), (fi5 & (fi5 << ((65535) & 63))), 4294967295, fi5}
            _ = fi6
            _ = v7
        }
        var v8 int32 = 2147483647
        fmt.Println(fi5)
        _ = fi5
        _ = v8
    }
    v1 -= (0.1 / (v1 * v1))
    v9 := []int{-57, 32768}
    if false {
        v4 = 42
        if false {
            for fi10 := 0; fi10 < 18; fi10++ {
                v11 := GenS1{F0: 231, F1: (834 > 330)}
                if len(v9) >= 2 {
                    v9 = v9[1:len(v9)]
                }
                v1 -= float64(0.0)
                _ = fi10
                _ = v11
            }
        } else {
            fmt.Println(v1)
        }
        for fi12 := 0; fi12 < 6; fi12++ {
            var v13 uint64 = 845
            _ = fi12
            _ = v13
        }
    } else {
        var v14 int = 418
        var v15 int16 = -774
        v16 := v2[799]
        _ = v14
        _ = v15
        _ = v16
    }
    v17 := GenS2{F0: 0}
    if true {
        _ = len(v9)
        v4 += (v4 + 42)
    }
    v18 := GenS3{F0: v4, F1: true, F2: 127}
    v4 -= v4
    var v19 int32 = -1
    if (!(240 >= 1)) {
        var v20 float32 = (0.1 - 1e10)
        _ = v20
    } else {
        for fi21 := 0; fi21 < 14; fi21++ {
            v19 = (v19 << ((1) & 63))
            var v22 uint8 = 127
            if (true && false) {
                for fi23 := 0; fi23 < 5; fi23++ {
                    switch 709 {
                    case 133:
                        v4 += (^(v4))
                    case 127:
                        if (535 != 2) {
                            if (65535 <= 255) {
                                var v24 string = ("" + "a")
                                var v25 uint16 = 1
                                if ((false || (v19 <= (v19 >> ((42) & 63)))) || ((!(0 != 724)) || false)) {
                                    var v26 uint16 = uint16(v22)
                                    _ = v26
                                } else {
                                    v3[fi21] = (fi21 >> ((225) & 63))
                                    v9 = append(v9, fi23)
                                }
                                _ = v24
                                _ = v25
                            }
                            v22 += v22
                            v4 = 627
                        } else {
                            v27 := GenS3{F0: v4, F1: (!false), F2: 127}
                            _ = v27
                        }
                    case 32767:
                        var v28 bool = false
                        _ = v28
                    default:
                        if ((v4 + v4) > v4) {
                            v29 := GenS1{F0: v22, F1: (true || ((0.0 + (0.1 / v1)) > (0.0 * (v1 / (float64(((-2.25 + 3.14159) / (float32((float64((float32(float64(float32(3.14159))) * float32(-2.25))) - (-2.25 * (v1 + (v1 / 1.0))))) / float32(3.14159)))) / v1)))))}
                            var v30 uint16 = 255
                            _ = v29
                            _ = v30
                        } else {
                            var v31 int64 = 52
                            switch "hello" {
                            case "":
                                if false {
                                    var v32 uint = 593
                                    v2[(fi23 >> ((v32) & 63))] = -1
                                    _ = v32
                                } else {
                                    var v33 float32 = (-2.25 * 1.5)
                                    _ = v33
                                }
                            }
                            _ = v31
                        }
                    }
                    _ = fi23
                }
                var v34 bool = false
                var v35 int16 = 42
                _ = v34
                _ = v35
            } else {
                switch "foo bar" {
                case "":
                    switch fi21 {
                    case 128:
                        v4 -= v4
                    case 1:
                        _ = len(v9)
                    case 42:
                        v36 := GenS4{F0: -1, F1: 668, F2: ""}
                        _ = v36
                    default:
                        var v37 bool = true
                        _ = v37
                    }
                default:
                    for fi38 := 0; fi38 < 15; fi38++ {
                        var v39 int16 = 1000
                        _ = fi38
                        _ = v39
                    }
                }
            }
            _ = fi21
            _ = v22
        }
    }
    var v40 string = ("\n\t" + "\n\t")
    v1 -= v1
    var v41 uint64 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v9
}
