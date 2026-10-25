package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8; F2 uint }
type GenS2 struct { F0 uint16; F1 uint32 }
type GenS3 struct { F0 int16; F1 string; F2 int16 }
func genFn0(p0 int32, p1 uint8) uint16 {
    var v7 uint64 = 1
    for fi8 := 0; fi8 < 11; fi8++ {
        switch fi8 {
        case 127:
            if ((0 <= 1) && true) {
                v7 = v7
                var v9 int32 = p0
                for fi10 := 0; fi10 < 4; fi10++ {
                    v11 := []int{(fi8 | (fi10 | (fi8 * fi10))), fi8, (fi8 | fi10)}
                    _ = fi10
                    _ = v11
                }
                _ = v9
            }
        case 0:
            p1 = uint8(v7)
        default:
            v7 += 18446744073709551615
        }
        if (40 >= 493) {
            v12 := []int{(fi8 * int(p1))}
            _ = v12
        } else {
            v7 += v7
            for fi13 := 0; fi13 < 18; fi13++ {
                fmt.Println(p1, v7)
                _ = fi13
            }
        }
        _ = fi8
    }
    if (uint16(v7) >= 228) {
        p1 = (p1 % (((p1 >> ((18446744073709551615) & 63))) | 1))
    } else {
        if true {
            if false {
                p1 = p1
                v14 := GenS1{F0: 591, F1: -47, F2: 65535}
                var v15 bool = false
                _ = v14
                _ = v15
            } else {
                var v16 uint8 = (p1 >> ((255) & 63))
                _ = v16
            }
            switch 105 {
            case 717:
                fmt.Println(v7, v7, p1)
            case -1:
                if (65535 != 921) {
                    if false {
                        v17 := []int{0}
                        var v18 int16 = -1
                        v19 := GenS2{F0: 1, F1: 807}
                        _ = v17
                        _ = v18
                        _ = v19
                    }
                    var v20 int = 127
                    _ = v20
                }
            }
        }
    }
    _ = p0
    _ = p1
    _ = v7
    return 2
}
func genFn1(p0 uint, p1 int16) uint {
    p0 += (p0 | 0)
    v21 := GenS2{F0: 347, F1: 4294967295}
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = v21
    return 385
}

func main() {
    var v1 float32 = 1e10
    var v2 int = -765
    var v3 uint = 0
    var v4 int16 = 680
    var v5 uint32 = 65535
    var v6 int8 = 1
    for fi22 := 0; fi22 < 7; fi22++ {
        switch "hello" {
        case "":
            fmt.Println(v5, v1)
        }
        v1 = v1
        fmt.Println(fi22, v1, fi22)
        _ = fi22
    }
    fmt.Println(v1)
    if true {
        fmt.Println(v6)
        var v23 uint16 = 155
        _ = v23
    }
    v5 = 0
    v24 := GenS3{F0: v4, F1: (("\n\t" + "foo bar") + ("\n\t" + "hello")), F2: -32768}
    v1 = 3.14159
    var v25 float64 = 0.0
    fmt.Println(v5, v6)
    for fi26 := 0; fi26 < 17; fi26++ {
        fmt.Println(v6, v5, v2)
        fmt.Println(v4)
        _ = fi26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
