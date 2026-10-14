package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool }
func genFn0(p0 int32, p1 int8) uint64 {
    fmt.Println(p0)
    switch ("a" + "hello") {
    case "\n\t":
        fmt.Println(p1, p1)
    default:
        var v6 uint8 = 127
        _ = v6
    }
    p1 = -35
    switch 2147483647 {
    case 9223372036854775807:
        var v7 int64 = -437
        _ = v7
    case -394:
        p1 = p1
    case 0:
        p0 += 2147483647
    default:
        v8 := []int{2147483647, 4294967295, 42}
        _ = v8
    }
    p1 -= -7
    _ = p0
    _ = p1
    return 713
}
func genFn1() uint64 {
    var v9 uint32 = 2147483647
    for fi10 := 0; fi10 < 14; fi10++ {
        v9 -= (v9 % ((v9) | 1))
        v11 := []int{fi10, 1}
        var v12 int8 = 1
        _ = fi10
        _ = v11
        _ = v12
    }
    _ = v9
    return 18446744073709551615
}

func main() {
    v1 := []int{}
    v2 := make(map[uint32]int)
    v3 := []uint32{65535}
    var v4 int16 = 542
    var v5 float32 = (float32(1.5) + (0.0 - 0.1))
    var v13 int64 = -103
    for fi14 := 0; fi14 < 13; fi14++ {
        for fi15 := 0; fi15 < 14; fi15++ {
            if len(v3) > 0 {
                v3[0] = 4294967295
            }
            v16 := GenS1{F0: 102, F1: (!false)}
            _ = fi15
            _ = v16
        }
        _ = fi14
    }
    v5 -= 1.5
    if ("a" <= "") {
        if ((true || (false && false)) && false) {
            if true {
                v5 -= v5
                for fi17 := 0; fi17 < 14; fi17++ {
                    fmt.Println(v4)
                    var v18 string = "a"
                    fmt.Println(v5, v13, v5)
                    _ = fi17
                    _ = v18
                }
            } else {
                v19 := GenS1{F0: -128, F1: true}
                _ = v19
            }
            v13 += v13
        }
        v20 := GenS1{F0: -1, F1: true}
        _ = v20
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v21 uint64 = 678
    switch "\n\t" {
    case "":
        fmt.Println(v21, v4, v13)
    default:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    for fi22 := 0; fi22 < 19; fi22++ {
        v23 := GenS1{F0: 64, F1: false}
        var v24 int64 = (^(v13))
        _ = fi22
        _ = v23
        _ = v24
    }
    for fi25 := 0; fi25 < 4; fi25++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        if (43 > 0) {
            fmt.Println(fi25)
            var v26 bool = (false && true)
            _ = v26
        } else {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
        _ = fi25
    }
    for fi27 := 0; fi27 < 16; fi27++ {
        v13 -= v13
        _ = fi27
    }
    var v28 bool = (true && true)
    if (0 == 0) {
        v13 += v13
    }
    fmt.Println(v4, v13, v4)
    for fi29 := 0; fi29 < 11; fi29++ {
        for fi30 := 0; fi30 < 12; fi30++ {
            _ = len(v2)
            _ = fi30
        }
        _ = fi29
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v2
    _ = v21
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
