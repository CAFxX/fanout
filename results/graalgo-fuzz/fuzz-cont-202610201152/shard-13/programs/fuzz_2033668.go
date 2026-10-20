package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8 }
type GenS2 struct { F0 uint16; F1 int }
type GenS3 struct { F0 uint8; F1 int32 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int; F1 uint16 }

func main() {
    var v1 uint32 = 2147483647
    var v2 uint8 = 2
    var v3 float64 = (0.0 * 1.5)
    fmt.Println(v2, v3)
    if true {
        v4 := GenS1{F0: 42, F1: v2}
        fmt.Println(v2)
        if false {
            v5 := GenS2{F0: 884, F1: 32767}
            if false {
                v3 += v3
                var v6 uint64 = 0
                v7 := GenS3{F0: v2, F1: -1}
                _ = v6
                _ = v7
            }
            v2 = (v2 ^ v2)
            _ = v5
        } else {
            v1 -= v1
        }
        _ = v4
    }
    v2 = 42
    for fi8 := 0; fi8 < 20; fi8++ {
        var v9 uint8 = v2
        for fi10 := 0; fi10 < 3; fi10++ {
            if false {
                fmt.Println(v1, v3)
            } else {
                v11 := GenS1{F0: fi10, F1: (v9 + (v9 * (v9 >> ((2) & 63))))}
                switch "hello" {
                case "":
                    var v12 int32 = 42
                    _ = v12
                default:
                    v1 -= (v1 / (((v1 ^ v1)) | 1))
                }
                var v13 uint32 = v1
                _ = v11
                _ = v13
            }
            _ = fi10
        }
        v1 = v1
        _ = fi8
        _ = v9
    }
    fmt.Println(v3, v1)
    if (852 < 831) {
        var v14 bool = true
        v15 := []int{-110, 9223372036854775807}
        _ = v14
        _ = v15
    }
    var v16 float32 = (1e10 / 1.5)
    v17 := GenS4{F0: -1}
    if true {
        v2 = (v2 ^ 186)
    } else {
        if true {
            v2 -= 42
            v18 := []int{-932, 4294967295}
            _ = v18
        } else {
            var v19 uint16 = 0
            switch "hello" {
            case "":
                var v20 int = 0
                _ = v20
            default:
                var v21 bool = false
                _ = v21
            }
            if true {
                v19 = 65535
                if true {
                    if (!true) {
                        switch 1 {
                        case -2147483648:
                            var v22 int = -987
                            _ = v22
                        case -173:
                            v23 := GenS5{F0: 127, F1: 0}
                            _ = v23
                        }
                        v16 = v16
                    } else {
                        var v24 int64 = 573
                        v25 := GenS3{F0: 42, F1: -116}
                        v24 += v24
                        _ = v24
                        _ = v25
                    }
                    v26 := GenS1{F0: 662, F1: (v2 * 2)}
                    _ = v26
                }
            } else {
                if (v16 <= ((0.1 + float32(v3)) - v16)) {
                    if true {
                        if (-1 == -79) {
                            v27 := GenS5{F0: 0, F1: v19}
                            _ = v27
                        }
                        var v28 int16 = 42
                        v3 -= (0.1 - float64(v16))
                        _ = v28
                    }
                    v29 := []int{}
                    _ = v29
                } else {
                    if ((-1 <= 176) && true) {
                        v30 := GenS4{F0: 32767}
                        _ = v30
                    }
                    v31 := []int{-9223372036854775808}
                    _ = v31
                }
            }
            _ = v19
        }
    }
    var v32 uint8 = 255
    v2 = v2
    var v33 uint64 = 1
    var v34 uint32 = 564
    var v35 int16 = int16(v1)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
}
