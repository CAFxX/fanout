package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16; F2 int64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint16; F1 uint32; F2 uint8 }

func main() {
    var v1 int64 = 1
    var v2 uint8 = 19
    var v3 int16 = -1
    v4 := make(map[string]bool)
    var v5 uint32 = 2147483647
    v6 := make(map[uint32]string)
    v7 := GenS1{F0: v1, F1: 105, F2: (-(v1))}
    var v8 int = int(v5)
    if true {
        for fi9 := 0; fi9 < 6; fi9++ {
            v10 := GenS2{F0: 100}
            _ = fi9
            _ = v10
        }
        var v11 uint64 = 4294967295
        _ = v11
    }
    v2 = (v2 % ((v2) | 1))
    switch 32768 {
    case 2147483647:
        var v12 uint32 = (v5 | (v5 + v5))
        _ = v12
    case 720:
        var v13 bool = (uint(v2) > 255)
        _ = v13
    }
    if ((v5 >= (v5 * 479)) && (771 >= 65535)) {
        switch "" {
        case "foo bar":
            var v14 int = v8
            _ = v14
        case "a":
            v15 := GenS3{F0: -697}
            _ = v15
        case "":
            switch ("\n\t" + "hello") {
            case "hello":
                for fi16 := 0; fi16 < 4; fi16++ {
                    var v17 uint64 = 18446744073709551615
                    if (!(false && false)) {
                        for fi18 := 0; fi18 < 15; fi18++ {
                            v19 := GenS4{F0: v1}
                            _ = len(v6)
                            v4["\n\t"] = true
                            v20 := GenS5{F0: 840, F1: v5, F2: (v2 * 2)}
                            _ = fi18
                            _ = v19
                            _ = v20
                        }
                        if true {
                            fmt.Println(v2, v17)
                        } else {
                            v21 := GenS5{F0: 220, F1: (v5 / (((v5 * (v5 + (^(v5))))) | 1)), F2: v2}
                            _ = v21
                        }
                    }
                    fmt.Println(v8, v5, v5)
                    _ = fi16
                    _ = v17
                }
            default:
                fmt.Println(v8, v3)
            }
        default:
            v22 := []int{v8}
            _ = v22
        }
        var v23 bool = (true && (2 >= 2))
        _ = v23
    }
    switch 4294967295 {
    case -804:
        v1 = (v1 / ((v1) | 1))
    case 2147483647:
        v24 := []int{int(v5), 0, 4294967295}
        _ = v24
    default:
        var v25 int32 = int32(v2)
        _ = v25
    }
    if ("\n\t" == "\n\t") {
        fmt.Println(v1, v1)
        fmt.Println(v5)
        v5 = (^(v5))
    }
    v26 := []int{v8}
    var v27 uint = 255
    var v28 float64 = (1.5 + float64(float32(0.0)))
    for fi29 := 0; fi29 < 19; fi29++ {
        delete(v4, "hello")
        v27 = 18446744073709551615
        _ = fi29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
