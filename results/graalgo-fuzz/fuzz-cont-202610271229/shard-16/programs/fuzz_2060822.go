package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint16; F1 int8 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint16; F1 int8 }
func genFn0(p0 int16) uint64 {
    for fi4 := 0; fi4 < 8; fi4++ {
        v5 := []int{(fi4 << ((30) & 63)), fi4, -1}
        _ = fi4
        _ = v5
    }
    p0 -= 212
    var v6 int = 127
    for fi7 := 0; fi7 < 12; fi7++ {
        if false {
            p0 += -994
            fmt.Println(p0)
            fmt.Println(v6)
        } else {
            var v8 int32 = -2147483648
            if false {
                v9 := GenS1{F0: 42}
                if (94 < 527) {
                    v6 = fi7
                }
                _ = v9
            } else {
                v10 := GenS2{F0: 65535, F1: 11}
                fmt.Println(v8)
                switch "a" {
                case "hello":
                    v11 := GenS1{F0: 32767}
                    _ = v11
                }
                _ = v10
            }
            v6 -= -438
            _ = v8
        }
        v12 := []int{fi7, fi7}
        _ = fi7
        _ = v12
    }
    _ = p0
    _ = v6
    return 42
}

func main() {
    var v1 int8 = -128
    var v2 float32 = float32(((float64(1e10) * (float64(1e10) - 1e10)) / (float64(-2.25) + ((float64(3.14159) * -2.25) / float64(-2.25)))))
    var v3 uint8 = 42
    v2 = 0.0
    v13 := []int{-960}
    v2 -= -2.25
    for fi14 := 0; fi14 < 20; fi14++ {
        var v15 uint = 0
        v16 := GenS3{F0: false}
        var v17 uint16 = 255
        _ = fi14
        _ = v15
        _ = v16
        _ = v17
    }
    switch 32767 {
    case 0:
        v3 = 2
    case 4294967295:
        v18 := GenS2{F0: 255, F1: (v1 >> ((2) & 63))}
        _ = v18
    case 127:
        var v19 uint8 = (v3 + (v3 << ((255) & 63)))
        _ = v19
    default:
        for fi20 := 0; fi20 < 1; fi20++ {
            if true {
                var v21 float32 = v2
                v13 = append(v13, 455)
                _ = v21
            } else {
                var v22 uint16 = 32767
                var v23 uint = 42
                _ = v22
                _ = v23
            }
            _ = len(v13)
            var v24 uint32 = uint32(v3)
            _ = fi20
            _ = v24
        }
    }
    var v25 int16 = 32767
    switch 9223372036854775807 {
    case 193:
        v1 += (v1 % ((1) | 1))
    case 128:
        v1 += 127
    case 363:
        v25 += (v25 >> ((0) & 63))
    default:
        v3 -= (^(v3))
    }
    _ = len(v13)
    v26 := GenS4{F0: 255, F1: -89}
    var v27 float32 = (v2 + v2)
    if (!true) {
        var v28 uint32 = 2147483648
        _ = len(v13)
        _ = v28
    } else {
        var v29 float32 = (0.1 - v27)
        v3 = v3
        var v30 bool = (true && true)
        _ = v29
        _ = v30
    }
    for fi31 := 0; fi31 < 10; fi31++ {
        if ("" >= (("a" + "a") + ("" + ("\n\t" + "")))) {
            fmt.Println(fi31)
            var v32 uint = 65535
            v27 += v27
            _ = v32
        } else {
            _ = len(v13)
            fmt.Println(v25, v25)
        }
        _ = fi31
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
}
