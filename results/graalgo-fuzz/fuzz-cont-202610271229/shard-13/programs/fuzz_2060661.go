package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool }
type GenS2 struct { F0 uint64; F1 uint32; F2 uint64 }
func genFn0(p0 int, p1 int16, p2 int8) int8 {
    var v6 uint = 1
    v7 := []int{p0, 128, 1, 32767}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return p2
}
func genFn1(p0 int32, p1 uint8, p2 bool) uint32 {
    if false {
        for fi8 := 0; fi8 < 5; fi8++ {
            v9 := []int{}
            _ = fi8
            _ = v9
        }
        var v10 uint16 = 42
        _ = v10
    } else {
        v11 := []int{9223372036854775807, 32767, -744, 711}
        p0 = (p0 * 0)
        _ = v11
    }
    if p2 {
        var v12 float64 = (3.14159 / 1.0)
        _ = v12
    } else {
        for fi13 := 0; fi13 < 11; fi13++ {
            v14 := []int{fi13, (-(fi13)), fi13}
            _ = fi13
            _ = v14
        }
        fmt.Println(p2)
        p1 -= p1
    }
    v15 := []int{0, 4294967295, 32767, -9223372036854775808}
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    return uint32(p1)
}

func main() {
    var v1 int64 = -489
    v2 := make(map[uint32]uint32)
    var v3 uint = 0
    v4 := make(map[string]uint32)
    var v5 string = "foo bar"
    v16 := GenS1{F0: 1, F1: true}
    var v17 uint64 = uint64(v3)
    v18 := GenS1{F0: -4, F1: ((!(32767 < 32767)) || false)}
    fmt.Println(v1, v3, v3)
    var v19 uint = (v3 ^ 853)
    var v20 int32 = 243
    switch 32767 {
    case -2147483648:
        v5 = v5
    case 689:
        fmt.Println(v17, v19)
    case 1:
        if (!false) {
            if false {
                v19 = 953
                for fi21 := 0; fi21 < 2; fi21++ {
                    var v22 int64 = v1
                    _ = fi21
                    _ = v22
                }
            }
        } else {
            v23 := []int{-515, 120, 4294967295}
            switch (((("a" + ((("\n\t" + ((v5 + ("\n\t" + ("a" + "\n\t"))) + ((v5 + v5) + v5))) + "\n\t") + v5)) + ("" + v5)) + v5) + "a") {
            case "hello":
                v23 = append(v23, -304)
            case "":
                var v24 int32 = (v20 ^ -1)
                _ = v24
            case "foo bar":
                var v25 uint8 = 0
                _ = v25
            }
            switch "hello" {
            case "a":
                for fi26 := 0; fi26 < 17; fi26++ {
                    for fi27 := 0; fi27 < 13; fi27++ {
                        fmt.Println(v5, v20)
                        var v28 uint = (v3 ^ v3)
                        switch "a" {
                        case "":
                            v2[65535] = 902
                        case "foo bar":
                            var v29 uint64 = (v17 - 4294967295)
                            _ = v29
                        default:
                            fmt.Println(fi26, v28)
                        }
                        _ = len(v4)
                        _ = fi27
                        _ = v28
                    }
                    var v30 uint64 = (v17 & uint64(v3))
                    _ = fi26
                    _ = v30
                }
            case "foo bar":
                v17 -= v17
            default:
                v31 := GenS1{F0: 32767, F1: false}
                _ = v31
            }
            _ = v23
        }
    }
    v1 += v1
    var v32 uint8 = 52
    v33 := GenS2{F0: v17, F1: genFn1(-278, (^(v32)), (true && false)), F2: (v17 + (v17 << ((v19) & 63)))}
    if false {
        fmt.Println(v3, v17, v32)
        v2[42] = 835
        v34 := GenS1{F0: -1, F1: false}
        _ = v34
    }
    for fi35 := 0; fi35 < 20; fi35++ {
        var v36 int64 = (v1 >> (((v19 & (v3 | 65535))) & 63))
        var v37 float64 = 1e10
        _ = fi35
        _ = v36
        _ = v37
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
}
