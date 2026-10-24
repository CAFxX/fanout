package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int32; F1 int8 }
type GenS3 struct { F0 uint64; F1 int64 }
func genFn0(p0 string, p1 int32, p2 int64) int32 {
    v7 := []int{513, -721, 127, 2147483647}
    var v8 uint16 = 645
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 1
}
func genFn1(p0 uint, p1 uint16) int8 {
    var v9 int32 = 0
    p1 += 574
    var v10 uint16 = 255
    v11 := []int{-2147483648}
    p0 = (^(p0))
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v9
    return -128
}

func main() {
    var v1 uint64 = 1
    v2 := []uint32{2, 30, 2, 2147483648}
    var v3 int16 = -32768
    var v4 float64 = float64(1e10)
    v5 := make(map[int]string)
    var v6 int16 = 587
    var v12 uint32 = 2147483648
    v13 := GenS1{F0: 0}
    var v14 uint64 = v1
    var v15 int = 128
    v5[(v15 & (v15 >> ((4294967295) & 63)))] = ""
    v6 += v6
    fmt.Println(v3, v1)
    var v16 int32 = 0
    v3 = v6
    _ = len(v5)
    v6 += 42
    if (!(!true)) {
        v17 := GenS2{F0: v16, F1: genFn1(769, 15)}
        switch "foo bar" {
        case "foo bar":
            v3 = v3
        case "a":
            if false {
                if len(v2) > 0 {
                    v2[0] = v12
                }
                v18 := GenS3{F0: 1, F1: -1}
                v5[(v15 % ((v15) | 1))] = "\n\t"
                _ = v18
            } else {
                v1 += v14
            }
        case "\n\t":
            var v19 int16 = v3
            _ = v19
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = v17
    }
    if (((-2.25 - (1.5 * float32(v4))) != -2.25) && false) {
        var v20 int64 = 0
        if (v3 != (v6 * v6)) {
            switch v15 {
            case 636:
                fmt.Println(v20, v15)
            case 42:
                if false {
                    var v21 int8 = int8(v14)
                    fmt.Println(v1, v20, v21)
                    for fi22 := 0; fi22 < 1; fi22++ {
                        for fi23 := 0; fi23 < 19; fi23++ {
                            var v24 bool = true
                            _ = fi23
                            _ = v24
                        }
                        _ = fi22
                    }
                    _ = v21
                } else {
                    var v25 uint16 = 2
                    v26 := GenS3{F0: uint64(v25), F1: (v20 % ((v20) | 1))}
                    fmt.Println(v6, v25)
                    _ = v25
                    _ = v26
                }
            default:
                fmt.Println(v15, v15, v20)
            }
            var v27 int32 = v16
            fmt.Println(v12)
            _ = v27
        } else {
            switch "a" {
            case "":
                var v28 uint8 = 0
                _ = v28
            case "foo bar":
                var v29 uint8 = 23
                _ = v29
            case "hello":
                for fi30 := 0; fi30 < 5; fi30++ {
                    for fi31 := 0; fi31 < 8; fi31++ {
                        v1 = v14
                        _ = fi31
                    }
                    _ = fi30
                }
            default:
                _ = len(v2)
            }
        }
        var v32 uint8 = 29
        _ = v20
        _ = v32
    } else {
        if true {
            v33 := v5[(v15 ^ v15)]
            var v34 bool = false
            _ = v33
            _ = v34
        } else {
            v35 := GenS1{F0: 127}
            v12 += 2147483647
            _ = v35
        }
        var v36 int16 = (v3 >> ((996) & 63))
        _ = v36
    }
    switch "hello" {
    case "a":
        v6 += -1
    case "":
        v37 := v5[(-(v15))]
        _ = v37
    case "\n\t":
        var v38 bool = (false && (!(true || false)))
        _ = v38
    default:
        v39 := GenS3{F0: 4294967295, F1: 0}
        _ = v39
    }
    v3 -= -332
    var v40 int64 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
