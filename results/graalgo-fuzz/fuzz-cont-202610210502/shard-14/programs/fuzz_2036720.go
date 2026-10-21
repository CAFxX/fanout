package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 uint32 }
type GenS2 struct { F0 int; F1 uint16; F2 string }

func main() {
    var v1 uint32 = 2147483648
    var v2 uint64 = 921
    v3 := []int{}
    var v4 int8 = -21
    v5 := make(map[int]string)
    v6 := GenS1{F0: (^(v1)), F1: v1, F2: v1}
    switch 42 {
    case 9223372036854775807:
        var v7 uint = 42
        _ = v7
    case -9223372036854775808:
        var v8 uint32 = 65535
        _ = v8
    case 964:
        for fi9 := 0; fi9 < 3; fi9++ {
            var v10 int = (fi9 & -855)
            _ = fi9
            _ = v10
        }
    default:
        var v11 uint32 = 42
        _ = v11
    }
    fmt.Println(v1, v2, v1)
    v4 += (v4 + v4)
    if false {
        if true {
            v1 = v1
            fmt.Println(v2)
            delete(v5, -220)
        } else {
            var v12 uint = uint(v1)
            for fi13 := 0; fi13 < 16; fi13++ {
                for fi14 := 0; fi14 < 19; fi14++ {
                    v1 -= v1
                    switch "\n\t" {
                    case "":
                        var v15 bool = (-1 > -101)
                        _ = v15
                    case "foo bar":
                        _ = len(v5)
                    }
                    if (-489 != 1000) {
                        v2 += v2
                        v1 = v1
                        v2 += v2
                    } else {
                        v1 += (^(v1))
                        var v16 uint16 = 32767
                        _ = v16
                    }
                    _ = fi14
                }
                var v17 uint64 = (v2 >> (((v12 - v12)) & 63))
                v2 += uint64(v1)
                _ = fi13
                _ = v17
            }
            for fi18 := 0; fi18 < 19; fi18++ {
                switch fi18 {
                case 9223372036854775807:
                    for fi19 := 0; fi19 < 1; fi19++ {
                        v2 = 4294967295
                        v12 = (v12 * 710)
                        delete(v5, (fi19 / ((32768) | 1)))
                        var v20 float64 = (3.14159 * 0.0)
                        _ = fi19
                        _ = v20
                    }
                case 32768:
                    var v21 int16 = 42
                    _ = v21
                }
                _ = fi18
            }
            _ = v12
        }
    } else {
        fmt.Println(v4)
        for fi22 := 0; fi22 < 13; fi22++ {
            v3 = append(v3, fi22)
            if (1 != 255) {
                var v23 int8 = (v4 >> ((0) & 63))
                v24 := GenS1{F0: (^(v1)), F1: (v1 * uint32(v2)), F2: (v1 * (v1 % ((v1) | 1)))}
                _ = v23
                _ = v24
            } else {
                var v25 string = ""
                _ = v25
            }
            _ = fi22
        }
    }
    for fi26 := 0; fi26 < 17; fi26++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        if false {
            var v27 float64 = float64((float32(0.0) * (0.1 - float32(0.0))))
            _ = v27
        } else {
            fmt.Println(v2, v1)
        }
        v28 := GenS1{F0: v1, F1: (v1 & (v1 | (^(v1)))), F2: uint32(v2)}
        _ = fi26
        _ = v28
    }
    fmt.Println(v4)
    v29 := GenS1{F0: (v1 & 143), F1: v1, F2: (^(v1))}
    for fi30 := 0; fi30 < 20; fi30++ {
        for fi31 := 0; fi31 < 20; fi31++ {
            v32 := GenS1{F0: v1, F1: v1, F2: v1}
            v2 -= (v2 ^ v2)
            v33 := GenS1{F0: (v1 / ((uint32(v2)) | 1)), F1: (v1 | 2147483647), F2: 531}
            _ = fi31
            _ = v32
            _ = v33
        }
        switch "hello" {
        case "":
            v34 := v5[-1]
            _ = v34
        default:
            v35 := GenS2{F0: fi30, F1: 65535, F2: "hello"}
            _ = v35
        }
        if (!(!false)) {
            switch ("foo bar" + ((("\n\t" + "") + ((("hello" + "a") + "a") + "a")) + "foo bar")) {
            case "\n\t":
                if false {
                    var v36 uint64 = v2
                    _ = v36
                }
            case "foo bar":
                v4 = (^(v4))
            default:
                var v37 uint = 42
                _ = v37
            }
        }
        _ = fi30
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi38 := 0; fi38 < 6; fi38++ {
        _ = len(v3)
        var v39 uint32 = v1
        _ = fi38
        _ = v39
    }
    v2 += v2
    var v40 uint8 = 0
    var v41 int16 = -1
    fmt.Println(v41, v41)
    if (3.14159 > 1e10) {
        if len(v3) > 0 {
            v3[0] = -636
        }
    }
    v1 = 19
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
}
