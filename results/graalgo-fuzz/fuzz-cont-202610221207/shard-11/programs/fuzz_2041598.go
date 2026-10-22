package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string }
type GenS2 struct { F0 int8; F1 int64; F2 bool }
func genFn0(p0 int8) uint {
    fmt.Println(p0)
    for fi4 := 0; fi4 < 1; fi4++ {
        if false {
            p0 += (p0 % ((127) | 1))
            if true {
                v5 := GenS1{F0: 2, F1: ""}
                _ = v5
            }
        } else {
            v6 := GenS2{F0: (p0 % (((p0 / ((93) | 1))) | 1)), F1: -1, F2: true}
            for fi7 := 0; fi7 < 11; fi7++ {
                var v8 int64 = -668
                _ = fi7
                _ = v8
            }
            if true {
                var v9 float32 = 3.14159
                v10 := GenS2{F0: (^(p0)), F1: 0, F2: (!false)}
                var v11 uint64 = 1
                _ = v10
                _ = v11
                _ = v9
            } else {
                fmt.Println(fi4, fi4)
            }
            _ = v6
        }
        p0 -= -1
        v12 := []int{}
        _ = fi4
        _ = v12
    }
    fmt.Println(p0)
    _ = p0
    return 2
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int32 = 327
    v3 := make(map[uint32]string)
    v2 = -1
    switch -1 {
    case 2147483647:
        v2 = (v2 % (((v2 * -233)) | 1))
    default:
        for fi13 := 0; fi13 < 18; fi13++ {
            v2 = (v2 - (v2 / (((v2 | v2)) | 1)))
            _ = fi13
        }
    }
    switch -108 {
    case 700:
        for fi14 := 0; fi14 < 7; fi14++ {
            if (true || true) {
                v2 = v2
            } else {
                var v15 float64 = float64(0.1)
                _ = v15
            }
            fmt.Println(fi14)
            _ = fi14
        }
    case -740:
        fmt.Println(v2)
    case 4294967295:
        if true {
            if true {
                for fi16 := 0; fi16 < 20; fi16++ {
                    v2 -= (v2 & -732)
                    v2 -= v2
                    _ = fi16
                }
            }
            var v17 int16 = -1
            _ = v17
        } else {
            for fi18 := 0; fi18 < 2; fi18++ {
                for fi19 := 0; fi19 < 2; fi19++ {
                    v2 += v2
                    _ = fi19
                }
                v2 -= (v2 * -2147483648)
                _ = fi18
            }
            for fi20 := 0; fi20 < 19; fi20++ {
                v21 := GenS2{F0: 127, F1: 9223372036854775807, F2: (!(((-119 == -1) && true) || (0.0 < 1.5)))}
                _ = fi20
                _ = v21
            }
            v2 = 42
        }
    default:
        switch ((("" + "a") + "a") + ("a" + "hello")) {
        case "hello":
            for fi22 := 0; fi22 < 14; fi22++ {
                fmt.Println(fi22, v2)
                var v23 int = (^(fi22))
                _ = fi22
                _ = v23
            }
        default:
            v2 = (^(v2))
        }
    }
    v2 += (^(v2))
    var v24 uint = 42
    fmt.Println(v2, v2)
    switch 0 {
    case 4294967295:
        var v25 uint64 = 18446744073709551615
        _ = v25
    }
    v2 += v2
    for fi26 := 0; fi26 < 6; fi26++ {
        if true {
            var v27 string = "a"
            _ = v27
        }
        v28 := []int{fi26, (fi26 << ((v24) & 63)), fi26, (fi26 - fi26)}
        v2 -= 1
        _ = fi26
        _ = v28
    }
    v24 -= (v24 >> (((v24 * v24)) & 63))
    v2 = (^(v2))
    v29 := v1[919]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v24
    _ = v29
    _ = v3
}
