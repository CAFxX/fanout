package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64 }
type GenS2 struct { F0 string; F1 uint16 }
type GenS3 struct { F0 int32; F1 int64; F2 int32 }
func genFn0() uint8 {
    var v5 uint64 = 1
    v5 = v5
    switch 42 {
    case 9223372036854775807:
        v6 := []int{-9223372036854775808, -9223372036854775808}
        _ = v6
    case 544:
        v7 := GenS1{F0: -2147483648, F1: 1}
        _ = v7
    case 1:
        switch 9223372036854775807 {
        case 1:
            switch "a" {
            case "a":
                v5 = v5
            }
        default:
            v5 = 4294967295
        }
    default:
        if false {
            v8 := GenS2{F0: "foo bar", F1: 65535}
            v5 = (v5 % (((v5 | v5)) | 1))
            v9 := GenS2{F0: "hello", F1: 1}
            _ = v8
            _ = v9
        } else {
            var v10 int = 156
            _ = v10
        }
    }
    _ = v5
    return 2
}

func main() {
    v1 := []uint32{2, 1, 2147483647, 84}
    v2 := []int{9223372036854775807}
    var v3 float32 = float32((1.5 / -2.25))
    var v4 int64 = 976
    v3 = v3
    v1 = append(v1, 674)
    var v11 float32 = (3.14159 / v3)
    v4 += v4
    _ = len(v2)
    if len(v1) > 0 {
        v1[0] = 774
    }
    v4 += (v4 + v4)
    if false {
        _ = len(v2)
        var v12 int8 = -128
        _ = v12
    } else {
        switch -9223372036854775808 {
        case 32767:
            for fi13 := 0; fi13 < 14; fi13++ {
                var v14 int8 = -1
                if true {
                    switch "hello" {
                    case "":
                        v3 += (v3 + float32(-2.25))
                    case "hello":
                        var v15 bool = true
                        _ = v15
                    }
                } else {
                    v14 = -1
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
                _ = fi13
                _ = v14
            }
        default:
            for fi16 := 0; fi16 < 9; fi16++ {
                for fi17 := 0; fi17 < 10; fi17++ {
                    var v18 int64 = v4
                    var v19 float64 = 1.5
                    if false {
                        if (917 != 616) {
                            v3 = v3
                            v20 := GenS2{F0: "", F1: 32767}
                            _ = v20
                        }
                    }
                    v11 = -2.25
                    _ = fi17
                    _ = v18
                    _ = v19
                }
                v3 = v11
                _ = fi16
            }
        }
        v11 -= v3
        v2 = append(v2, -687)
    }
    v11 = 0.1
    if true {
        var v21 int64 = -1
        _ = v21
    } else {
        v22 := GenS3{F0: -1, F1: -1, F2: -559}
        _ = v22
    }
    v3 += v11
    for fi23 := 0; fi23 < 4; fi23++ {
        v2 = append(v2, 128)
        _ = fi23
    }
    var v24 int = 4294967295
    var v25 int8 = 0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
}
