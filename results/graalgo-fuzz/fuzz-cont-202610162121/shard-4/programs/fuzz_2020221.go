package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }
type GenS2 struct { F0 int; F1 int64 }
func genFn0(p0 int32) uint8 {
    for fi5 := 0; fi5 < 5; fi5++ {
        var v6 int64 = 854
        _ = fi5
        _ = v6
    }
    var v7 float32 = -2.25
    for fi8 := 0; fi8 < 12; fi8++ {
        switch "hello" {
        case "":
            p0 += (p0 + p0)
        case "foo bar":
            if true {
                v9 := []int{fi8, fi8, (^(fi8))}
                fmt.Println(fi8, v7)
                _ = v9
            }
        case "a":
            if true {
                var v10 uint32 = 42
                v11 := GenS1{F0: 2, F1: -99}
                _ = v10
                _ = v11
            }
        default:
            v12 := []int{296}
            _ = v12
        }
        _ = fi8
    }
    _ = p0
    _ = v7
    return 255
}
func genFn1(p0 int64, p1 uint32) uint16 {
    if true {
        p1 = p1
        if (128 == 42) {
            v13 := []int{32768, -9223372036854775808, 2147483647, 128}
            _ = v13
        }
        for fi14 := 0; fi14 < 11; fi14++ {
            switch (^(fi14)) {
            case -1:
                switch fi14 {
                case 921:
                    p0 = (p0 % ((9223372036854775807) | 1))
                }
            default:
                v15 := GenS1{F0: 535, F1: 0}
                _ = v15
            }
            v16 := []int{fi14, 127}
            _ = fi14
            _ = v16
        }
    } else {
        var v17 uint32 = (p1 << ((18446744073709551615) & 63))
        p0 += p0
        for fi18 := 0; fi18 < 5; fi18++ {
            p1 = 940
            _ = fi18
        }
        _ = v17
    }
    var v19 uint16 = 869
    fmt.Println(p0)
    for fi20 := 0; fi20 < 11; fi20++ {
        v19 = (v19 << ((619) & 63))
        _ = fi20
    }
    _ = p0
    _ = p1
    _ = v19
    return v19
}

func main() {
    v1 := []string{"", ("foo bar" + ("\n\t" + "")), "a"}
    v2 := make(map[uint32]uint32)
    v3 := make(map[uint32]uint32)
    var v4 float32 = float32((0.0 - 0.1))
    for fi21 := 0; fi21 < 13; fi21++ {
        v22 := GenS1{F0: genFn1(-9223372036854775808, 17), F1: 0}
        for fi23 := 0; fi23 < 16; fi23++ {
            v24 := GenS2{F0: fi21, F1: -956}
            _ = fi23
            _ = v24
        }
        v1 = append(v1, "a")
        _ = fi21
        _ = v22
    }
    _ = len(v1)
    var v25 bool = true
    fmt.Println(v25)
    var v26 int8 = 100
    v26 -= (v26 * (v26 >> ((255) & 63)))
    v1 = append(v1, ("foo bar" + ("\n\t" + "foo bar")))
    fmt.Println(v26)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
}
