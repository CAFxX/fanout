package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int8; F1 bool; F2 int32 }
func genFn0(p0 int32) uint8 {
    var v6 int8 = 42
    v7 := []int{983, -1, -9223372036854775808, -1}
    v6 -= 80
    switch "a" {
    case "a":
        var v8 bool = false
        _ = v8
    case "hello":
        v7 = append(v7, 9223372036854775807)
    case "\n\t":
        for fi9 := 0; fi9 < 12; fi9++ {
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
            if false {
                var v10 uint64 = 0
                v11 := GenS1{F0: (p0 * p0)}
                _ = v10
                _ = v11
            } else {
                v6 -= v6
                v12 := GenS1{F0: (-(p0))}
                _ = v12
            }
            var v13 string = "\n\t"
            _ = fi9
            _ = v13
        }
    default:
        for fi14 := 0; fi14 < 14; fi14++ {
            switch 0 {
            case 42:
                var v15 uint16 = 853
                _ = v15
            }
            _ = len(v7)
            _ = fi14
        }
    }
    v6 -= (v6 ^ v6)
    _ = p0
    _ = v6
    _ = v7
    return 239
}
func genFn1() int8 {
    if false {
        for fi16 := 0; fi16 < 6; fi16++ {
            var v17 float64 = 1.5
            fmt.Println(v17, v17)
            fmt.Println(v17)
            _ = fi16
            _ = v17
        }
        var v18 float64 = (-2.25 / (3.14159 + 1.5))
        switch ("\n\t" + "a") {
        case "\n\t":
            for fi19 := 0; fi19 < 18; fi19++ {
                v20 := []int{-909, fi19, fi19}
                _ = len(v20)
                _ = fi19
                _ = v20
            }
        case "foo bar":
            v21 := GenS2{F0: 42}
            _ = v21
        default:
            if (0 < 1) {
                switch "a" {
                case "\n\t":
                    var v22 int64 = -481
                    _ = v22
                }
            }
        }
        _ = v18
    }
    var v23 int = -9223372036854775808
    var v24 uint = 4294967295
    _ = v23
    _ = v24
    return -1
}

func main() {
    var v1 int = 32768
    v2 := []uint32{4294967295, 4294967295}
    v3 := make(map[string]bool)
    var v4 int32 = 0
    v5 := []int{v1, (v1 * v1), (v1 >> ((255) & 63))}
    var v25 bool = (!(!true))
    v4 -= 1
    v1 = v1
    for fi26 := 0; fi26 < 13; fi26++ {
        v27 := GenS3{F0: 127, F1: false, F2: (v4 * v4)}
        switch 9223372036854775807 {
        case 127:
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        default:
            var v28 bool = ((v1 ^ v1) < int(v4))
            _ = v28
        }
        _ = fi26
        _ = v27
    }
    _ = len(v3)
    if len(v5) > 0 {
        v5[0] = (v1 >> ((18446744073709551615) & 63))
    }
    var v29 uint = 4294967295
    for fi30 := 0; fi30 < 13; fi30++ {
        v31 := GenS2{F0: 0}
        _ = fi30
        _ = v31
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
