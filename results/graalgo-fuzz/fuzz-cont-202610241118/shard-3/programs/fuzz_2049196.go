package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint }
type GenS2 struct { F0 uint8; F1 int16; F2 uint8 }
type GenS3 struct { F0 uint32; F1 int32; F2 uint16 }
func genFn0(p0 uint32) uint32 {
    var v5 int16 = -32768
    v6 := []int{-9223372036854775808, -557, 32768}
    for fi7 := 0; fi7 < 4; fi7++ {
        p0 -= (p0 >> ((1) & 63))
        _ = fi7
    }
    _ = p0
    _ = v5
    _ = v6
    return (p0 ^ (p0 / (((^(p0))) | 1)))
}
func genFn1(p0 int) uint16 {
    for fi8 := 0; fi8 < 10; fi8++ {
        p0 += (fi8 * (p0 % ((2147483647) | 1)))
        var v9 int16 = -32768
        v10 := []int{}
        _ = fi8
        _ = v10
        _ = v9
    }
    v11 := []int{p0, p0, (p0 + p0)}
    for fi12 := 0; fi12 < 15; fi12++ {
        p0 = (-(fi12))
        _ = fi12
    }
    switch (("" + ((((("\n\t" + "\n\t") + "foo bar") + "foo bar") + (("" + (("" + "\n\t") + "a")) + "a")) + ("" + "\n\t"))) + ("" + ("" + "a"))) {
    case "hello":
        var v13 int8 = 64
        _ = v13
    case "":
        for fi14 := 0; fi14 < 8; fi14++ {
            switch p0 {
            case -1:
                if ((1 > -847) && (4294967295 >= 33)) {
                    v15 := GenS1{F0: (((1 >= 1) && (42 == 42)) && true), F1: 65535}
                    _ = v15
                }
            default:
                switch (p0 / ((p0) | 1)) {
                case -667:
                    p0 = 515
                case 4294967295:
                    switch fi14 {
                    case 32767:
                        var v16 int = fi14
                        _ = v16
                    case 128:
                        if true {
                            if (-128 != 96) {
                                v17 := GenS2{F0: 231, F1: 42, F2: 255}
                                p0 = 933
                                _ = v17
                            }
                        }
                    case 252:
                        var v18 int8 = 42
                        _ = v18
                    }
                }
            }
            p0 += -480
            _ = fi14
        }
    }
    if len(v11) > 0 {
        v11[0] = 47
    }
    _ = p0
    _ = v11
    return 255
}

func main() {
    v1 := []uint32{1}
    v2 := []uint32{2, 992, 65535}
    v3 := make(map[uint32]int)
    v4 := make(map[int]uint32)
    var v19 uint64 = 303
    v3[42] = -572
    fmt.Println(v19)
    var v20 uint16 = 255
    var v21 int = 1
    v22 := GenS3{F0: genFn0(0), F1: 1, F2: (v20 | 1)}
    if (!false) {
        v23 := GenS3{F0: 2147483648, F1: -975, F2: 2}
        _ = v23
    }
    v21 += v21
    var v24 int = (v21 | v21)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v24
    _ = v3
    _ = v4
}
