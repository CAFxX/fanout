package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint64; F1 int; F2 uint16 }
type GenS3 struct { F0 bool; F1 int8; F2 uint16 }
func genFn0() uint64 {
    v6 := []int{0}
    v7 := GenS1{F0: 42}
    _ = v6
    _ = v7
    return 0
}
func genFn1(p0 uint8) uint32 {
    for fi8 := 0; fi8 < 14; fi8++ {
        p0 -= (p0 >> ((255) & 63))
        p0 = (p0 + p0)
        _ = fi8
    }
    switch "" {
    case "":
        v9 := GenS1{F0: -2147483648}
        _ = v9
    case "hello":
        var v10 int32 = 1
        _ = v10
    }
    p0 -= (p0 * (p0 + p0))
    p0 -= (p0 >> ((4294967295) & 63))
    _ = p0
    return 2147483648
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[string]bool)
    v3 := []string{(("" + "hello") + "foo bar"), "", ("" + ("" + (("foo bar" + "hello") + "\n\t"))), "hello"}
    v4 := make(map[string]string)
    var v5 string = "a"
    var v11 string = (v5 + v5)
    v5 = (v5 + "a")
    var v12 bool = (-636 < 42)
    switch 697 {
    case 744:
        fmt.Println(v12)
    case -2147483648:
        _ = len(v4)
    default:
        v11 = ((v11 + (v11 + v5)) + "\n\t")
    }
    var v13 float64 = (-2.25 * float64(1e10))
    for fi14 := 0; fi14 < 20; fi14++ {
        for fi15 := 0; fi15 < 18; fi15++ {
            var v16 bool = v12
            v17 := GenS2{F0: 0, F1: fi14, F2: 377}
            _ = fi15
            _ = v16
            _ = v17
        }
        var v18 int8 = -128
        switch v11 {
        case "hello":
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
        _ = fi14
        _ = v18
    }
    switch 128 {
    case 289:
        v12 = v12
    case 1:
        fmt.Println(v13, v5)
    case 4294967295:
        v19 := v2[(v11 + v5)]
        _ = v19
    }
    v20 := GenS3{F0: v12, F1: -1, F2: 0}
    fmt.Println(v5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
