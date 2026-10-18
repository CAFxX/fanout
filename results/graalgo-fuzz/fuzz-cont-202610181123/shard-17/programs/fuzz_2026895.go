package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint64; F1 string; F2 uint64 }
func genFn0(p0 uint32, p1 uint32) uint {
    v4 := GenS1{F0: (82 != 128)}
    var v5 uint16 = 42
    v5 = v5
    var v6 int32 = 42
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return 125
}
func genFn1(p0 uint8) int16 {
    var v7 uint8 = 255
    var v8 uint64 = 2
    var v9 uint16 = 255
    if (true || false) {
        v10 := []int{}
        _ = v10
    }
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return -1
}

func main() {
    var v1 int32 = 42
    v2 := make(map[uint32]uint32)
    var v3 int32 = v1
    switch -510 {
    case 278:
        if (4294967295 != 535) {
            v11 := GenS2{F0: 42, F1: "hello", F2: 18446744073709551615}
            _ = v11
        } else {
            v12 := GenS1{F0: false}
            fmt.Println(v1)
            if true {
                v1 = (v1 << ((2) & 63))
                for fi13 := 0; fi13 < 14; fi13++ {
                    v14 := []int{fi13, fi13, (fi13 | fi13), fi13}
                    fmt.Println(fi13)
                    _ = fi13
                    _ = v14
                }
            }
            _ = v12
        }
    case -166:
        v1 = (v1 | (^(v3)))
    }
    v1 = 0
    fmt.Println(v3, v1)
    if (v1 == (^(v3))) {
        delete(v2, 2147483647)
    }
    for fi15 := 0; fi15 < 8; fi15++ {
        fmt.Println(v3, v3, v1)
        _ = fi15
    }
    v1 = 42
    v16 := []int{-950, 0}
    if false {
        fmt.Println(v3, v3)
    }
    switch ("" + "hello") {
    case "hello":
        switch 2147483647 {
        case -2147483648:
            switch "hello" {
            case "a":
                if (!(210 == 0)) {
                    for fi17 := 0; fi17 < 12; fi17++ {
                        var v18 float32 = (1e10 - float32(0.1))
                        _ = fi17
                        _ = v18
                    }
                }
            case "\n\t":
                if len(v16) > 0 {
                    v16[0] = 0
                }
            }
        case 4294967295:
            for fi19 := 0; fi19 < 5; fi19++ {
                _ = len(v2)
                _ = fi19
            }
        }
    case "\n\t":
        var v20 int8 = 42
        _ = v20
    }
    for fi21 := 0; fi21 < 19; fi21++ {
        v22 := v2[2147483648]
        var v23 float32 = (((((1.5 + ((((0.0 / 1.0) + (0.1 * (0.0 + 0.1))) * 3.14159) / (1e10 - float32(3.14159)))) * 1e10) - 0.1) * 0.0) / (1.5 * 1.5))
        if true {
            v24 := GenS2{F0: 764, F1: "", F2: 977}
            v25 := v2[v22]
            v2[(v25 / (((v22 % ((v25) | 1))) | 1))] = v22
            _ = v24
            _ = v25
        }
        _ = fi21
        _ = v22
        _ = v23
    }
    if len(v16) > 0 {
        v16[0] = 240
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
}
