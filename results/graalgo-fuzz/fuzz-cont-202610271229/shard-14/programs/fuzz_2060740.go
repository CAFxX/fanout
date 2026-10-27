package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint32; F1 bool; F2 int64 }
type GenS3 struct { F0 int64; F1 int }
type GenS4 struct { F0 int32; F1 string }
type GenS5 struct { F0 int32; F1 uint8 }
func genFn0(p0 int) uint32 {
    switch (p0 - p0) {
    case 2147483647:
        for fi5 := 0; fi5 < 10; fi5++ {
            p0 = (^(fi5))
            v6 := []int{-772, fi5, 688, (fi5 ^ 0)}
            _ = fi5
            _ = v6
        }
    default:
        v7 := GenS1{F0: 9223372036854775807}
        _ = v7
    }
    p0 = p0
    _ = p0
    return 2147483647
}

func main() {
    var v1 int = 128
    var v2 float32 = (1e10 / 3.14159)
    var v3 uint8 = 255
    var v4 uint = 255
    v8 := GenS2{F0: genFn0(v1), F1: (true && false), F2: 0}
    v9 := GenS2{F0: 701, F1: true, F2: 9223372036854775807}
    var v10 uint16 = 0
    v11 := []int{v1, v1, v1, v1}
    v10 -= v10
    fmt.Println(v1, v4, v10)
    var v12 float64 = 3.14159
    var v13 uint16 = (^(v10))
    v10 = (v10 | v10)
    for fi14 := 0; fi14 < 6; fi14++ {
        var v15 uint16 = (v13 >> ((v4) & 63))
        v16 := GenS1{F0: -181}
        _ = fi14
        _ = v15
        _ = v16
    }
    v17 := GenS3{F0: -9223372036854775808, F1: v1}
    v18 := GenS4{F0: -2147483648, F1: ("foo bar" + ("\n\t" + "\n\t"))}
    if false {
        var v19 float64 = 0.1
        _ = v19
    }
    if false {
        v13 -= (v13 << (((v4 & v4)) & 63))
        switch v1 {
        case -2147483648:
            var v20 uint32 = 518
            _ = v20
        case -9223372036854775808:
            v21 := GenS4{F0: -163, F1: ""}
            _ = v21
        }
        if true {
            if false {
                var v22 uint64 = 2
                for fi23 := 0; fi23 < 2; fi23++ {
                    var v24 float32 = v2
                    _ = fi23
                    _ = v24
                }
                switch "hello" {
                case "":
                    v25 := GenS5{F0: 549, F1: v3}
                    _ = v25
                case "foo bar":
                    v1 -= v1
                case "hello":
                    for fi26 := 0; fi26 < 14; fi26++ {
                        v22 = uint64(v13)
                        _ = fi26
                    }
                }
                _ = v22
            }
        }
    }
    for fi27 := 0; fi27 < 13; fi27++ {
        if len(v11) > 0 {
            v11[0] = 32767
        }
        var v28 int64 = -9223372036854775808
        v28 -= (v28 * (-(v28)))
        _ = fi27
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
