package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint; F2 int8 }
func genFn0() string {
    if (0 != 9223372036854775807) {
        v7 := GenS1{F0: 128, F1: 4294967295, F2: -1}
        _ = v7
    } else {
        var v8 uint = 255
        _ = v8
    }
    var v9 uint64 = 107
    v10 := GenS1{F0: 128, F1: 42, F2: 127}
    v9 += v9
    v11 := GenS1{F0: 0, F1: 293, F2: 42}
    _ = v10
    _ = v11
    _ = v9
    return ""
}

func main() {
    var v1 int32 = -1
    var v2 uint8 = 255
    v3 := make(map[uint32]int)
    var v4 uint64 = 769
    v5 := make(map[uint32]uint32)
    var v6 float64 = 1.5
    fmt.Println(v1)
    v1 -= (v1 / (((-(v1))) | 1))
    v6 = v6
    v12 := GenS1{F0: 122, F1: 223, F2: 100}
    switch "a" {
    case "hello":
        if false {
            for fi13 := 0; fi13 < 11; fi13++ {
                v4 += 4294967295
                v14 := GenS1{F0: v2, F1: 0, F2: 42}
                _ = fi13
                _ = v14
            }
            v1 = (v1 % (((^(v1))) | 1))
        } else {
            var v15 string = "a"
            fmt.Println(v15)
            _ = v15
        }
    case "\n\t":
        for fi16 := 0; fi16 < 13; fi16++ {
            if false {
                v2 -= v2
            } else {
                v6 += (v6 - v6)
            }
            _ = fi16
        }
    }
    if false {
        for fi17 := 0; fi17 < 18; fi17++ {
            switch fi17 {
            case 4294967295:
                v2 += (v2 % ((uint8(v4)) | 1))
            }
            var v18 int = 4294967295
            v18 = 798
            _ = fi17
            _ = v18
        }
        var v19 float32 = ((0.1 / (0.1 / 1e10)) / float32(1e10))
        _ = v19
    }
    v20 := GenS1{F0: v2, F1: 0, F2: -109}
    if false {
        v5[42] = 2147483648
    } else {
        var v21 int32 = 0
        _ = v21
    }
    var v22 int32 = v1
    if true {
        v23 := GenS1{F0: (v2 | (v2 % (((v2 - v2)) | 1))), F1: 609, F2: -128}
        _ = v23
    }
    if true {
        var v24 uint16 = 65535
        _ = v24
    }
    var v25 int32 = 0
    if false {
        if false {
            if true {
                v6 = v6
            }
            v22 = (v1 - 0)
        }
        v1 = v22
        v4 -= v4
    }
    switch 0 {
    case 4294967295:
        for fi26 := 0; fi26 < 14; fi26++ {
            v2 = uint8(v4)
            if (!false) {
                v4 = v4
            }
            _ = fi26
        }
    default:
        for fi27 := 0; fi27 < 9; fi27++ {
            v6 += 0.0
            var v28 uint8 = 0
            switch 0 {
            case 42:
                v29 := GenS1{F0: (v28 + v2), F1: 255, F2: 42}
                _ = v29
            default:
                switch "\n\t" {
                case "":
                    var v30 string = "hello"
                    _ = v30
                }
            }
            _ = fi27
            _ = v28
        }
    }
    switch 0 {
    case -854:
        for fi31 := 0; fi31 < 1; fi31++ {
            if true {
                v32 := GenS1{F0: (^(v2)), F1: 255, F2: 127}
                _ = v32
            }
            _ = fi31
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v20
    _ = v22
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
