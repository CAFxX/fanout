package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }
type GenS2 struct { F0 int; F1 int32; F2 int8 }
type GenS3 struct { F0 uint16; F1 int64; F2 uint8 }
type GenS4 struct { F0 string; F1 uint }
func genFn0(p0 uint32) uint {
    v7 := GenS1{F0: 883, F1: 42}
    if (!(p0 < (p0 >> ((255) & 63)))) {
        p0 += (^(p0))
    }
    _ = p0
    _ = v7
    return 255
}
func genFn1(p0 string) uint {
    switch 9223372036854775807 {
    case 9223372036854775807:
        switch 1 {
        case 543:
            p0 = p0
        case 128:
            p0 = ("" + p0)
        default:
            for fi8 := 0; fi8 < 7; fi8++ {
                switch p0 {
                case "hello":
                    v9 := []int{}
                    _ = v9
                case "":
                    for fi10 := 0; fi10 < 18; fi10++ {
                        v11 := []int{fi8, 779}
                        var v12 uint32 = 361
                        var v13 uint32 = v12
                        _ = fi10
                        _ = v11
                        _ = v12
                        _ = v13
                    }
                case "a":
                    switch 128 {
                    case -2147483648:
                        var v14 uint16 = 255
                        _ = v14
                    }
                default:
                    if true {
                        var v15 uint8 = 1
                        _ = v15
                    } else {
                        switch fi8 {
                        case 127:
                            var v16 uint64 = 42
                            _ = v16
                        }
                        if true {
                            v17 := GenS1{F0: 4294967295, F1: 0}
                            _ = v17
                        }
                        p0 = (p0 + "a")
                    }
                }
                _ = fi8
            }
        }
    }
    for fi18 := 0; fi18 < 4; fi18++ {
        p0 = ("\n\t" + "")
        var v19 uint = 18446744073709551615
        _ = fi18
        _ = v19
    }
    for fi20 := 0; fi20 < 9; fi20++ {
        v21 := GenS2{F0: fi20, F1: -1, F2: -128}
        v22 := GenS3{F0: 42, F1: 1, F2: 85}
        _ = fi20
        _ = v21
        _ = v22
    }
    _ = p0
    return 0
}

func main() {
    v1 := make(map[string]int)
    v2 := []int{}
    v3 := make(map[string]uint32)
    var v4 uint = 65535
    v5 := []int{}
    v6 := []int{4294967295, -436}
    delete(v1, "\n\t")
    delete(v1, "\n\t")
    switch 9223372036854775807 {
    case -395:
        v4 -= v4
    default:
        v23 := GenS2{F0: 951, F1: 42, F2: 19}
        _ = v23
    }
    var v24 int32 = -2147483648
    v25 := GenS3{F0: 255, F1: -9223372036854775808, F2: 127}
    v26 := GenS4{F0: ("\n\t" + "hello"), F1: v4}
    v27 := GenS4{F0: "a", F1: 1}
    fmt.Println(v24, v4)
    var v28 bool = (!(241 == 116))
    for fi29 := 0; fi29 < 16; fi29++ {
        fmt.Println(v28, v24)
        v30 := GenS2{F0: (fi29 * 1), F1: -1, F2: -1}
        v3[(("\n\t" + "\n\t") + "")] = 541
        _ = fi29
        _ = v30
    }
    v1["hello"] = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
