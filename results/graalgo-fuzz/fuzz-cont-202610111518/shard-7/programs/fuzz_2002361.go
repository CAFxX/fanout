package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool; F2 uint32 }
type GenS2 struct { F0 int8; F1 uint16 }
func genFn0() int64 {
    v4 := GenS1{F0: 180, F1: (-704 != 414), F2: 718}
    var v5 int8 = 101
    var v6 int8 = (v5 | (^(v5)))
    v7 := []int{}
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return 1
}

func main() {
    var v1 int = 1
    var v2 uint64 = 2
    var v3 uint = 401
    for fi8 := 0; fi8 < 17; fi8++ {
        var v9 int16 = int16(v3)
        _ = fi8
        _ = v9
    }
    for fi10 := 0; fi10 < 19; fi10++ {
        var v11 uint32 = 357
        _ = fi10
        _ = v11
    }
    if (4294967295 > 65535) {
        var v12 uint32 = 920
        _ = v12
    } else {
        for fi13 := 0; fi13 < 18; fi13++ {
            v3 = 0
            _ = fi13
        }
        var v14 uint32 = 0
        _ = v14
    }
    fmt.Println(v3)
    if false {
        v1 = -44
        v15 := []int{}
        _ = v15
    } else {
        v16 := []int{(v1 / (((v1 ^ v1)) | 1)), (v1 | -38)}
        var v17 int16 = 0
        _ = v16
        _ = v17
    }
    var v18 uint8 = 42
    for fi19 := 0; fi19 < 11; fi19++ {
        switch fi19 {
        case 42:
            v20 := []int{(fi19 ^ int(v3)), (fi19 >> (((v3 + (v3 | v3))) & 63)), v1, (v1 - (v1 % (((v1 - fi19)) | 1)))}
            _ = v20
        case 439:
            v21 := GenS2{F0: 1, F1: 42}
            _ = v21
        case 1:
            v1 = int(v18)
        default:
            var v22 int8 = -1
            _ = v22
        }
        if true {
            fmt.Println(v18, v18)
            fmt.Println(fi19, v3)
            v18 -= (^(v18))
        } else {
            for fi23 := 0; fi23 < 5; fi23++ {
                fmt.Println(v2)
                _ = fi23
            }
            switch v1 {
            case 4294967295:
                if (!true) {
                    for fi24 := 0; fi24 < 12; fi24++ {
                        v2 = 435
                        _ = fi24
                    }
                    var v25 uint64 = v2
                    var v26 bool = true
                    _ = v25
                    _ = v26
                } else {
                    var v27 int16 = -1
                    _ = v27
                }
            case 9223372036854775807:
                v28 := []int{(v1 - v1), v1}
                _ = v28
            case 370:
                switch "hello" {
                case "":
                    fmt.Println(v1, v2, v3)
                case "hello":
                    fmt.Println(v1)
                default:
                    fmt.Println(fi19, v3)
                }
            }
        }
        _ = fi19
    }
    if (!(!true)) {
        v1 += (v1 | v1)
        var v29 string = ("foo bar" + "\n\t")
        var v30 bool = true
        _ = v29
        _ = v30
    } else {
        var v31 float32 = 1.5
        _ = v31
    }
    var v32 int32 = 2147483647
    if ((0.1 * (0.0 + float32(((0.1 * (1e10 - 0.0)) / 1.0)))) <= float32((3.14159 / 3.14159))) {
        if true {
            v1 -= (v1 + v1)
            var v33 uint8 = 42
            v34 := GenS1{F0: -9223372036854775808, F1: true, F2: 4294967295}
            _ = v33
            _ = v34
        }
        fmt.Println(v3, v32, v1)
    } else {
        v35 := GenS2{F0: -95, F1: 0}
        _ = v35
    }
    fmt.Println(v1, v3, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v32
}
