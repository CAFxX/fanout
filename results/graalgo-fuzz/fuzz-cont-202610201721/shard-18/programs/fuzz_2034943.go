package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int64; F1 uint32; F2 uint64 }
type GenS3 struct { F0 uint16; F1 string; F2 uint64 }
type GenS4 struct { F0 uint32; F1 uint }
func genFn0(p0 int8, p1 bool, p2 int16) uint16 {
    if (p0 > p0) {
        v6 := []int{}
        if false {
            var v7 float64 = (0.1 / (-2.25 * 0.1))
            _ = v7
        } else {
            switch 715 {
            case 42:
                var v8 float64 = float64(3.14159)
                _ = v8
            case 2147483647:
                var v9 float32 = 3.14159
                _ = v9
            default:
                var v10 uint16 = 255
                _ = v10
            }
        }
        if len(v6) > 0 {
            v6[0] = 127
        }
        _ = v6
    }
    if (!(537 < 32767)) {
        var v11 float64 = -2.25
        v12 := []int{-110, 4294967295, 0, -743}
        if len(v12) > 0 {
            v12[0] = -2147483648
        }
        _ = v11
        _ = v12
    }
    p0 = -83
    for fi13 := 0; fi13 < 16; fi13++ {
        fmt.Println(p2)
        var v14 string = "hello"
        v15 := GenS1{F0: (fi13 >> ((18446744073709551615) & 63))}
        _ = fi13
        _ = v14
        _ = v15
    }
    _ = p0
    _ = p1
    _ = p2
    return 0
}

func main() {
    var v1 uint16 = 2
    var v2 uint = 761
    var v3 int = 32767
    v4 := []int{(v3 % ((v3) | 1))}
    var v5 uint = v2
    var v16 string = ""
    fmt.Println(v3, v1, v16)
    v17 := GenS2{F0: -993, F1: 2147483647, F2: 4294967295}
    if len(v4) > 0 {
        v4[0] = v3
    }
    _ = len(v4)
    var v18 uint64 = 18446744073709551615
    var v19 int32 = int32(v3)
    if false {
        if true {
            fmt.Println(v18, v18)
            var v20 int32 = v19
            v21 := GenS1{F0: v3}
            _ = v20
            _ = v21
        } else {
            if true {
                for fi22 := 0; fi22 < 13; fi22++ {
                    if ((((!true) && true) || (255 > 127)) || false) {
                        fmt.Println(v18, v19)
                        if (true || ((!true) && false)) {
                            var v23 string = (v16 + "hello")
                            v24 := GenS1{F0: 0}
                            _ = v23
                            _ = v24
                        }
                        switch fi22 {
                        case 1:
                            if (false && (!false)) {
                                v18 += v18
                                switch "\n\t" {
                                case "":
                                    if (v1 != 255) {
                                        var v25 int32 = (^(v19))
                                        _ = v25
                                    }
                                }
                                v26 := GenS2{F0: 42, F1: 640, F2: 1}
                                _ = v26
                            } else {
                                v3 += (fi22 + (fi22 * 32768))
                            }
                        case 32768:
                            v19 = v19
                        case 232:
                            v27 := GenS3{F0: 0, F1: ("a" + "a"), F2: (v18 - (^(v18)))}
                            _ = v27
                        default:
                            var v28 uint = v5
                            _ = v28
                        }
                    } else {
                        fmt.Println(v19, v5, v1)
                    }
                    v4 = append(v4, 4294967295)
                    var v29 uint32 = uint32(v1)
                    _ = fi22
                    _ = v29
                }
                var v30 int64 = 0
                _ = v30
            }
        }
        var v31 uint = 65535
        switch v3 {
        case -942:
            var v32 float32 = 3.14159
            _ = v32
        case 4294967295:
            v18 = 1
        case -1:
            for fi33 := 0; fi33 < 14; fi33++ {
                var v34 float32 = 0.1
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                v35 := GenS1{F0: fi33}
                _ = fi33
                _ = v34
                _ = v35
            }
        default:
            v3 += v3
        }
        _ = v31
    }
    v4 = append(v4, (v3 % ((v3) | 1)))
    if ((((1e10 + float32((1e10 + -2.25))) + -2.25) > float32(float64(((0.1 / (1.5 + (1.5 - 0.1))) / 1e10)))) && false) {
        var v36 uint = v5
        _ = v36
    } else {
        v37 := GenS4{F0: 2, F1: 4294967295}
        v5 += (v5 & v5)
        _ = v37
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
