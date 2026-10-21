package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint8; F1 int64 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int }
func genFn0(p0 int32, p1 bool, p2 uint64) uint32 {
    v4 := []int{2147483647, 127, 42}
    v5 := GenS1{F0: 2147483648}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 2147483647
}
func genFn1() uint32 {
    var v6 uint16 = 930
    if true {
        v7 := []int{-273, 269}
        for fi8 := 0; fi8 < 7; fi8++ {
            v7 = append(v7, (fi8 / ((381) | 1)))
            v9 := GenS2{F0: 127, F1: -704}
            var v10 int = fi8
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = v7
    } else {
        v6 = v6
    }
    v11 := []int{}
    _ = v11
    _ = v6
    return 1
}

func main() {
    v1 := []uint32{65535, 42}
    v2 := make(map[uint32]int)
    v3 := []int{2147483647}
    switch 757 {
    case 927:
        var v12 bool = (65535 < 0)
        _ = v12
    }
    for fi13 := 0; fi13 < 12; fi13++ {
        for fi14 := 0; fi14 < 2; fi14++ {
            v15 := GenS1{F0: 2147483648}
            var v16 uint = 255
            _ = fi14
            _ = v15
            _ = v16
        }
        _ = fi13
    }
    if (!false) {
        for fi17 := 0; fi17 < 1; fi17++ {
            v18 := v2[1]
            _ = fi17
            _ = v18
        }
        v19 := GenS3{F0: "foo bar"}
        v20 := GenS1{F0: 329}
        _ = v19
        _ = v20
    }
    var v21 int = -455
    var v22 float32 = (float32((0.0 - ((0.1 - 1.5) / 1e10))) - 0.0)
    _ = len(v2)
    var v23 int16 = -704
    if (true || true) {
        if (!((!((2 <= 1) && (32767 <= 0))) || false)) {
            v23 -= 600
        }
        for fi24 := 0; fi24 < 8; fi24++ {
            fmt.Println(fi24, fi24, v23)
            if (245 <= 1) {
                for fi25 := 0; fi25 < 4; fi25++ {
                    var v26 int32 = -1
                    _ = fi25
                    _ = v26
                }
            } else {
                for fi27 := 0; fi27 < 9; fi27++ {
                    v21 = (fi24 << ((544) & 63))
                    var v28 uint = 724
                    _ = fi27
                    _ = v28
                }
            }
            _ = fi24
        }
    }
    for fi29 := 0; fi29 < 4; fi29++ {
        for fi30 := 0; fi30 < 10; fi30++ {
            v31 := GenS1{F0: genFn0(42, (0 != 128), 0)}
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            v32 := GenS4{F0: fi30}
            v23 = v23
            _ = fi30
            _ = v31
            _ = v32
        }
        var v33 int16 = 32767
        _ = fi29
        _ = v33
    }
    var v34 uint = 255
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v34
}
