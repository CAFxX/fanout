package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint32; F1 int32 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 float64 = (1e10 - -2.25)
    v3 := make(map[string]bool)
    switch "\n\t" {
    case "":
        for fi4 := 0; fi4 < 7; fi4++ {
            switch "\n\t" {
            case "a":
                v2 -= (v2 / (float64((float32(float64(float32(v2))) + (0.0 / 3.14159))) - (v2 / 1.0)))
            }
            for fi5 := 0; fi5 < 9; fi5++ {
                v3["\n\t"] = (false && ((-128 > -1) || true))
                _ = fi5
            }
            if false {
                if (false && true) {
                    if (2147483647 < 17) {
                        v2 = -2.25
                        var v6 string = "hello"
                        _ = v6
                    }
                    v2 = v2
                } else {
                    for fi7 := 0; fi7 < 16; fi7++ {
                        v8 := []int{-9223372036854775808, (-(fi7)), 111}
                        _ = fi7
                        _ = v8
                    }
                }
                v9 := GenS1{F0: 2}
                _ = v9
            }
            _ = fi4
        }
    case "\n\t":
        for fi10 := 0; fi10 < 6; fi10++ {
            var v11 uint64 = 4294967295
            var v12 bool = true
            _ = fi10
            _ = v11
            _ = v12
        }
    case "hello":
        var v13 uint32 = 2
        _ = v13
    default:
        v1[0] = 65535
    }
    for fi14 := 0; fi14 < 19; fi14++ {
        fmt.Println(v2, v2)
        var v15 float64 = v2
        v16 := GenS1{F0: 42}
        _ = fi14
        _ = v15
        _ = v16
    }
    for fi17 := 0; fi17 < 12; fi17++ {
        v18 := GenS2{F0: 100}
        _ = fi17
        _ = v18
    }
    var v19 int8 = -108
    var v20 uint64 = 1
    v21 := []int{1}
    var v22 uint16 = 80
    _ = len(v3)
    if false {
        v19 = (v19 - -128)
        v2 += (v2 * v2)
    } else {
        for fi23 := 0; fi23 < 8; fi23++ {
            var v24 int16 = 42
            var v25 int8 = v19
            _ = fi23
            _ = v24
            _ = v25
        }
    }
    _ = len(v21)
    switch -777 {
    case 455:
        v26 := GenS2{F0: 101}
        _ = v26
    default:
        for fi27 := 0; fi27 < 12; fi27++ {
            if false {
                v22 = v22
            } else {
                v20 += 18446744073709551615
                if (false || true) {
                    v19 += 29
                }
                var v28 string = ""
                _ = v28
            }
            v19 += (-(v19))
            _ = fi27
        }
    }
    for fi29 := 0; fi29 < 8; fi29++ {
        var v30 uint16 = v22
        _ = fi29
        _ = v30
    }
    v20 -= 222
    v31 := GenS3{F0: 2147483647, F1: 958}
    v20 = (v20 | (v20 * 1))
    v22 = v22
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v31
}
