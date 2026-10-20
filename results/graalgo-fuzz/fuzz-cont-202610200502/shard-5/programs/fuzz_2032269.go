package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint; F2 uint16 }
type GenS2 struct { F0 uint32 }
func genFn0(p0 uint64, p1 int64, p2 uint32) uint {
    var v4 int64 = (p1 / ((p1) | 1))
    var v5 uint16 = 255
    p1 = 0
    for fi6 := 0; fi6 < 12; fi6++ {
        var v7 int64 = 479
        _ = fi6
        _ = v7
    }
    v8 := []int{2147483647, 128, 317, -673}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v8
    return 188
}
func genFn1(p0 int32, p1 int64) int16 {
    var v9 int32 = (p0 + (p0 ^ (^(p0))))
    p0 = 42
    if ((0 != (v9 * 0)) && true) {
        var v10 int8 = 100
        v11 := []int{42, -2147483648}
        for fi12 := 0; fi12 < 10; fi12++ {
            var v13 uint8 = 2
            _ = fi12
            _ = v13
        }
        _ = v10
        _ = v11
    } else {
        for fi14 := 0; fi14 < 5; fi14++ {
            var v15 uint16 = 457
            _ = fi14
            _ = v15
        }
        v16 := []int{}
        if len(v16) > 0 {
            v16[0] = 32768
        }
        _ = v16
    }
    var v17 float32 = 3.14159
    _ = p0
    _ = p1
    _ = v17
    _ = v9
    return -1
}

func main() {
    var v1 uint16 = 0
    v2 := make(map[uint32]uint32)
    v3 := make(map[string]int)
    v18 := GenS1{F0: false, F1: uint(v1), F2: (v1 & (v1 % ((2) | 1)))}
    fmt.Println(v1)
    v1 = 42
    var v19 uint64 = 18446744073709551615
    fmt.Println(v19)
    v19 -= (v19 << ((2) & 63))
    v19 -= 656
    var v20 int64 = 42
    v21 := GenS1{F0: (-483 == -54), F1: 2, F2: 1}
    fmt.Println(v20, v1, v1)
    var v22 float32 = 0.1
    _ = len(v3)
    v23 := GenS1{F0: (-128 < 1), F1: 42, F2: (v1 / (((^(v1))) | 1))}
    switch "" {
    case "\n\t":
        var v24 uint8 = 222
        _ = v24
    case "":
        switch "\n\t" {
        case "foo bar":
            v22 = v22
        }
    case "hello":
        v25 := GenS2{F0: 2147483647}
        _ = v25
    default:
        v1 = 2
    }
    v26 := GenS2{F0: 4294967295}
    if false {
        fmt.Println(v19, v19)
    } else {
        v20 = (^(v20))
        for fi27 := 0; fi27 < 20; fi27++ {
            v28 := GenS1{F0: (1 < -133), F1: genFn0((v19 / ((v19) | 1)), (v20 / ((int64(v1)) | 1)), 2147483648), F2: v1}
            switch ("a" + "foo bar") {
            case "foo bar":
                switch fi27 {
                case -1:
                    if false {
                        if true {
                            v29 := GenS1{F0: true, F1: genFn0(v19, (v20 >> ((18446744073709551615) & 63)), 2147483648), F2: 65535}
                            _ = v29
                        } else {
                            v30 := GenS2{F0: 4294967295}
                            _ = v30
                        }
                    }
                }
            default:
                v31 := []int{}
                _ = v31
            }
            v1 += 65535
            _ = fi27
            _ = v28
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v26
    _ = v3
}
