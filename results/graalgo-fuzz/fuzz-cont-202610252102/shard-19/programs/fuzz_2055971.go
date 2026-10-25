package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32 }
type GenS2 struct { F0 int; F1 uint64 }
type GenS3 struct { F0 int8; F1 uint16 }
type GenS4 struct { F0 int32 }
func genFn0() int16 {
    for fi4 := 0; fi4 < 15; fi4++ {
        v5 := []int{718, fi4}
        switch (("hello" + "a") + (("" + ("" + (("\n\t" + ("hello" + "hello")) + ""))) + ("foo bar" + "foo bar"))) {
        case "\n\t":
            var v6 bool = false
            _ = v6
        case "hello":
            var v7 int8 = 42
            _ = v7
        }
        var v8 uint32 = 0
        _ = fi4
        _ = v5
        _ = v8
    }
    v9 := GenS1{F0: 65535, F1: 1}
    _ = v9
    return 10
}

func main() {
    v1 := []uint32{373, 272, 81}
    v2 := make(map[uint32]uint32)
    var v3 uint64 = 590
    v3 = v3
    v10 := GenS2{F0: -684, F1: (v3 * 0)}
    if true {
        for fi11 := 0; fi11 < 14; fi11++ {
            v12 := GenS2{F0: (fi11 % ((fi11) | 1)), F1: 1}
            var v13 uint64 = (v3 / ((v3) | 1))
            for fi14 := 0; fi14 < 19; fi14++ {
                if (false || true) {
                    if true {
                        fmt.Println(v3, v13, fi11)
                    }
                } else {
                    v13 = v3
                }
                _ = fi14
            }
            _ = fi11
            _ = v12
            _ = v13
        }
        v3 = v3
    } else {
        v3 += (v3 | (v3 * v3))
    }
    v3 -= v3
    for fi15 := 0; fi15 < 16; fi15++ {
        for fi16 := 0; fi16 < 19; fi16++ {
            if (0 <= 825) {
                var v17 uint16 = 255
                v3 = (v3 % ((v3) | 1))
                _ = v17
            }
            switch (fi16 << ((166) & 63)) {
            case 42:
                if (!true) {
                    v3 -= (v3 >> ((42) & 63))
                }
            default:
                v18 := GenS3{F0: 42, F1: 42}
                _ = v18
            }
            v19 := GenS1{F0: uint32(v3), F1: 2147483647}
            _ = fi16
            _ = v19
        }
        _ = len(v1)
        v20 := GenS4{F0: -2147483648}
        _ = fi15
        _ = v20
    }
    if len(v1) > 0 {
        v1[0] = 4294967295
    }
    fmt.Println(v3)
    v1 = append(v1, 691)
    fmt.Println(v3)
    v3 += (v3 << ((255) & 63))
    if false {
        fmt.Println(v3)
    } else {
        v21 := v2[400]
        _ = v21
    }
    v3 -= 543
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
}
