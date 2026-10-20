package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8 }
type GenS2 struct { F0 uint64; F1 int8 }
func genFn0(p0 int16, p1 uint32, p2 uint16) uint {
    v7 := []int{}
    var v8 uint32 = 916
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 430
}

func main() {
    v1 := []string{"foo bar"}
    v2 := []uint32{}
    v3 := []int{-728}
    var v4 float64 = 3.14159
    v5 := []uint32{833}
    v6 := make(map[string]string)
    for fi9 := 0; fi9 < 19; fi9++ {
        var v10 uint64 = 1
        var v11 string = ("" + ("hello" + "a"))
        v10 = v10
        _ = fi9
        _ = v10
        _ = v11
    }
    var v12 float32 = (0.1 + 0.0)
    v5 = append(v5, 0)
    for fi13 := 0; fi13 < 5; fi13++ {
        for fi14 := 0; fi14 < 6; fi14++ {
            var v15 int = (fi13 % ((fi13) | 1))
            var v16 int32 = -666
            _ = fi14
            _ = v15
            _ = v16
        }
        for fi17 := 0; fi17 < 7; fi17++ {
            if len(v3) > 0 {
                v3[0] = fi13
            }
            _ = fi17
        }
        v12 = (0.0 - float32(v4))
        _ = fi13
    }
    if (2147483647 <= 747) {
        var v18 int16 = -32768
        _ = v18
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi19 := 0; fi19 < 10; fi19++ {
        switch "" {
        case "hello":
            for fi20 := 0; fi20 < 7; fi20++ {
                if true {
                    fmt.Println(v12, fi19)
                }
                _ = fi20
            }
        case "":
            fmt.Println(fi19, v4)
        case "a":
            v21 := GenS1{F0: 2147483647, F1: 149}
            _ = v21
        default:
            for fi22 := 0; fi22 < 7; fi22++ {
                v12 = 1e10
                if ((1 == -1) && (4294967295 == 42)) {
                    var v23 float32 = (v12 * v12)
                    v24 := GenS2{F0: 4294967295, F1: -82}
                    if true {
                        v12 += (float32(v4) / (((((0.0 / 0.1) * (0.0 - float32(v4))) * v23) + (v23 / v23)) - ((float32(v4) + (v23 * 1e10)) * 0.0)))
                        v12 = (v23 + (v23 * 1.5))
                        v12 = 0.0
                    } else {
                        if true {
                            switch "hello" {
                            case "\n\t":
                                v4 += (v4 + 0.0)
                            }
                            v25 := GenS2{F0: 2, F1: 25}
                            var v26 uint = 4294967295
                            fmt.Println(v4, fi19)
                            _ = v25
                            _ = v26
                        }
                        var v27 int32 = 0
                        if true {
                            switch fi22 {
                            case 2147483647:
                                v12 = (3.14159 * ((3.14159 - v12) + v12))
                            case 9223372036854775807:
                                v27 = v27
                            case 0:
                                _ = len(v1)
                            }
                            fmt.Println(fi22, fi22, fi22)
                            switch 128 {
                            case -136:
                                if (false && false) {
                                    var v28 uint16 = 42
                                    _ = v28
                                }
                            case 90:
                                fmt.Println(v12)
                            }
                        }
                        _ = v27
                    }
                    _ = v23
                    _ = v24
                }
                _ = fi22
            }
        }
        var v29 int = (fi19 >> ((810) & 63))
        _ = len(v1)
        _ = fi19
        _ = v29
    }
    fmt.Println(v4, v4)
    if true {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v6[("\n\t" + ("hello" + "hello"))] = ""
    } else {
        fmt.Println(v4)
    }
    switch "\n\t" {
    case "a":
        var v30 uint64 = 2
        _ = v30
    case "hello":
        fmt.Println(v4, v12)
    case "":
        var v31 uint8 = 255
        _ = v31
    }
    if len(v3) > 0 {
        v3[0] = 32768
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
