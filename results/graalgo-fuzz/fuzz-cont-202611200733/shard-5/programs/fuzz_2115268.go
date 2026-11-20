package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64 }
type GenS2 struct { F0 uint64; F1 bool }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int16 }
func genFn0() int32 {
    var v4 uint8 = 2
    v4 += (^(v4))
    _ = v4
    return 0
}
func genFn1() uint8 {
    switch 32767 {
    case 4294967295:
        var v5 bool = false
        _ = v5
    case -2147483648:
        for fi6 := 0; fi6 < 18; fi6++ {
            v7 := []int{}
            _ = fi6
            _ = v7
        }
    default:
        var v8 int16 = -32768
        _ = v8
    }
    v9 := []int{-9223372036854775808, -2147483648, 128}
    if len(v9) > 0 {
        v9[0] = -2147483648
    }
    _ = v9
    return 127
}

func main() {
    var v1 uint32 = 1
    var v2 int32 = 42
    v3 := make(map[uint32]bool)
    v10 := GenS1{F0: 255, F1: 42}
    v11 := GenS2{F0: 580, F1: true}
    v2 = v2
    if true {
        var v12 int8 = 85
        v2 = v2
        v2 = 1
        _ = v12
    }
    switch 945 {
    case 9223372036854775807:
        if false {
            for fi13 := 0; fi13 < 7; fi13++ {
                switch ("a" + "\n\t") {
                case "hello":
                    var v14 uint64 = 42
                    _ = v14
                default:
                    var v15 uint8 = 181
                    _ = v15
                }
                switch "a" {
                case "\n\t":
                    for fi16 := 0; fi16 < 5; fi16++ {
                        var v17 float64 = 1e10
                        v2 -= (-(v2))
                        if (true && (769 != 848)) {
                            v18 := GenS3{F0: 4294967295}
                            v19 := GenS2{F0: 18446744073709551615, F1: ((fi16 ^ (fi13 / ((fi16) | 1))) >= -1)}
                            _ = v18
                            _ = v19
                        }
                        _ = fi16
                        _ = v17
                    }
                case "":
                    v2 += 0
                case "a":
                    v20 := GenS2{F0: 362, F1: false}
                    _ = v20
                default:
                    var v21 uint32 = (v1 ^ v1)
                    _ = v21
                }
                var v22 int = -1
                _ = fi13
                _ = v22
            }
        }
    case 15:
        v23 := GenS2{F0: 1, F1: true}
        _ = v23
    default:
        for fi24 := 0; fi24 < 8; fi24++ {
            v25 := []int{fi24, 127, fi24, (fi24 - fi24)}
            if true {
                v26 := GenS2{F0: 249, F1: true}
                delete(v3, v1)
                _ = v26
            }
            switch ("\n\t" + "a") {
            case "a":
                v2 -= genFn0()
            }
            _ = fi24
            _ = v25
        }
    }
    var v27 float64 = -2.25
    var v28 uint16 = uint16(v1)
    v29 := GenS4{F0: 776}
    v30 := GenS2{F0: 2, F1: true}
    for fi31 := 0; fi31 < 13; fi31++ {
        fmt.Println(v28, v28)
        var v32 int16 = int16(v1)
        _ = fi31
        _ = v32
    }
    var v33 uint = 255
    v27 = v27
    v34 := GenS1{F0: 151, F1: 1}
    fmt.Println(v33, v2, v1)
    v28 = v28
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v33
    _ = v34
}
