package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64 }
type GenS2 struct { F0 uint8; F1 uint }
type GenS3 struct { F0 int32; F1 uint32 }

func main() {
    v1 := make(map[uint32]string)
    var v2 int8 = 52
    var v3 int32 = 42
    if (0 >= 0) {
        for fi4 := 0; fi4 < 8; fi4++ {
            var v5 int8 = (v2 << ((2) & 63))
            _ = fi4
            _ = v5
        }
    }
    for fi6 := 0; fi6 < 4; fi6++ {
        v1[2147483647] = ""
        var v7 uint64 = 42
        _ = fi6
        _ = v7
    }
    var v8 bool = (394 <= 50)
    for fi9 := 0; fi9 < 9; fi9++ {
        delete(v1, 2147483648)
        v10 := GenS1{F0: 255, F1: 1}
        var v11 uint16 = 821
        _ = fi9
        _ = v10
        _ = v11
    }
    var v12 int8 = (v2 / (((v2 << ((686) & 63))) | 1))
    v12 = 0
    for fi13 := 0; fi13 < 3; fi13++ {
        if (!v8) {
            for fi14 := 0; fi14 < 17; fi14++ {
                var v15 uint8 = 2
                _ = fi14
                _ = v15
            }
        } else {
            for fi16 := 0; fi16 < 3; fi16++ {
                fmt.Println(v2, v12)
                var v17 float64 = (0.0 - 0.1)
                _ = fi16
                _ = v17
            }
            switch fi13 {
            case 2147483647:
                v18 := GenS2{F0: 255, F1: 236}
                _ = v18
            case 4294967295:
                for fi19 := 0; fi19 < 3; fi19++ {
                    if (!(!(65535 < 42))) {
                        var v20 bool = v8
                        v21 := []int{(fi19 % (((fi13 & 32768)) | 1)), fi19}
                        _ = v20
                        _ = v21
                    }
                    fmt.Println(v8, fi13, fi13)
                    _ = fi19
                }
            }
            for fi22 := 0; fi22 < 20; fi22++ {
                v8 = v8
                _ = fi22
            }
        }
        v2 -= int8(v3)
        _ = fi13
    }
    for fi23 := 0; fi23 < 3; fi23++ {
        v2 = v2
        for fi24 := 0; fi24 < 16; fi24++ {
            switch fi24 {
            case 42:
                v25 := GenS2{F0: 255, F1: 0}
                if (9223372036854775807 < 42) {
                    var v26 int8 = (v12 | (v2 | (v12 % ((v2) | 1))))
                    v26 = (v26 << ((2) & 63))
                    _ = v26
                } else {
                    if (204 != 1) {
                        v27 := v1[0]
                        _ = v27
                    } else {
                        v28 := v1[0]
                        _ = v28
                    }
                    switch "a" {
                    case "hello":
                        v12 -= (v12 / ((100) | 1))
                    }
                    switch int(v12) {
                    case 308:
                        fmt.Println(fi23, fi24)
                    case -5:
                        switch "" {
                        case "\n\t":
                            var v29 string = ((("\n\t" + "") + "") + ("" + "\n\t"))
                            _ = v29
                        case "":
                            _ = len(v1)
                        default:
                            v30 := []int{-876, 2147483647, fi24, fi23}
                            _ = v30
                        }
                    default:
                        v8 = v8
                    }
                }
                v1[636] = ""
                v31 := GenS2{F0: 0, F1: 968}
                _ = v25
                _ = v31
            case 1:
                v12 = v12
            case 2147483647:
                fmt.Println(fi23, v3)
            default:
                fmt.Println(v3)
            }
            if v8 {
                switch "foo bar" {
                case "foo bar":
                    switch 4294967295 {
                    case 30:
                        fmt.Println(v3)
                    }
                }
            }
            v12 += 127
            fmt.Println(v2)
            _ = fi24
        }
        _ = fi23
    }
    v32 := GenS3{F0: v3, F1: 4294967295}
    var v33 uint = 1
    v34 := GenS1{F0: 15, F1: 1}
    var v35 float32 = 1e10
    v33 = 222
    var v36 uint64 = 18446744073709551615
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v8
}
