package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint16; F1 uint; F2 int8 }
type GenS4 struct { F0 int32 }
func genFn0(p0 uint32, p1 int) int32 {
    var v6 uint = 356
    var v7 int64 = 0
    v6 = 255
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 137
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint32 = 0
    v3 := []uint32{(^(v2)), 2, 79}
    v4 := make(map[uint32]int)
    var v5 int8 = 0
    var v8 uint = 210
    v2 = 645
    var v9 float64 = -2.25
    switch "" {
    case "foo bar":
        v10 := GenS1{F0: 811}
        _ = v10
    case "":
        var v11 int = 512
        _ = v11
    case "hello":
        if (!(353 > 1)) {
            var v12 int64 = 732
            _ = len(v3)
            _ = len(v3)
            _ = v12
        } else {
            v13 := v4[v2]
            fmt.Println(v13)
            _ = v13
        }
    }
    switch -9223372036854775808 {
    case 2147483647:
        fmt.Println(v2, v8)
    case 128:
        _ = len(v3)
    }
    if (151 > 2) {
        var v14 float32 = ((0.1 - 0.1) * (((((0.1 / (1e10 * float32(v9))) + float32(v9)) / ((float32(v9) * -2.25) - 1.5)) - 3.14159) + 0.1))
        for fi15 := 0; fi15 < 3; fi15++ {
            if false {
                _ = len(v3)
                switch "" {
                case "foo bar":
                    v8 += (v8 << ((437) & 63))
                case "hello":
                    fmt.Println(v2)
                default:
                    fmt.Println(v2, v14)
                }
                delete(v1, ("hello" + "foo bar"))
            }
            v16 := GenS2{F0: 1}
            _ = fi15
            _ = v16
        }
        _ = len(v3)
        _ = v14
    } else {
        _ = len(v3)
    }
    if false {
        fmt.Println(v5, v5)
        fmt.Println(v8, v5, v2)
        var v17 uint32 = (v2 + (v2 * (v2 | (v2 >> ((v8) & 63)))))
        _ = v17
    } else {
        v18 := GenS3{F0: 2, F1: (v8 * 0), F2: 0}
        v3 = append(v3, 0)
        if true {
            fmt.Println(v2, v2)
        } else {
            v19 := GenS4{F0: 42}
            if (726 != (v2 - v2)) {
                if (128 >= 128) {
                    switch "hello" {
                    case "hello":
                        var v20 uint16 = uint16(v8)
                        _ = v20
                    default:
                        v21 := v1["\n\t"]
                        _ = v21
                    }
                    var v22 uint8 = 128
                    v5 -= (v5 / ((-5) | 1))
                    _ = v22
                } else {
                    v5 = -37
                }
                var v23 int32 = -2147483648
                _ = v23
            }
            _ = v19
        }
        _ = v18
    }
    v24 := v4[1]
    v9 = v9
    var v25 string = "\n\t"
    for fi26 := 0; fi26 < 5; fi26++ {
        var v27 int8 = 127
        var v28 uint64 = 0
        _ = fi26
        _ = v27
        _ = v28
    }
    if true {
        _ = len(v3)
    } else {
        v29 := GenS2{F0: 42}
        _ = len(v4)
        if false {
            var v30 float32 = float32(v9)
            switch v25 {
            case "hello":
                delete(v1, ((v25 + v25) + ""))
            case "foo bar":
                for fi31 := 0; fi31 < 6; fi31++ {
                    var v32 int32 = 1
                    for fi33 := 0; fi33 < 16; fi33++ {
                        v34 := GenS4{F0: genFn0((v2 & (v2 + 1)), fi31)}
                        delete(v1, "foo bar")
                        var v35 bool = (!true)
                        _ = fi33
                        _ = v34
                        _ = v35
                    }
                    _ = fi31
                    _ = v32
                }
            case "\n\t":
                v36 := v4[v2]
                _ = v36
            default:
                var v37 int16 = 32767
                _ = v37
            }
            _ = v30
        } else {
            var v38 int16 = -54
            _ = v38
        }
        _ = v29
    }
    var v39 bool = (!true)
    var v40 int8 = (v5 % ((v5) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v8
    _ = v9
}
