package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }
type GenS2 struct { F0 int }
func genFn0(p0 int) uint32 {
    for fi7 := 0; fi7 < 9; fi7++ {
        p0 -= (-(fi7))
        _ = fi7
    }
    for fi8 := 0; fi8 < 18; fi8++ {
        p0 += -703
        if false {
            var v9 uint8 = 1
            _ = v9
        }
        _ = fi8
    }
    v10 := []int{-918, 128}
    _ = p0
    _ = v10
    return 562
}
func genFn1() string {
    if (true && true) {
        var v11 bool = false
        _ = v11
    } else {
        for fi12 := 0; fi12 < 15; fi12++ {
            for fi13 := 0; fi13 < 13; fi13++ {
                v14 := []int{(fi13 & (fi13 << ((1) & 63)))}
                var v15 string = "\n\t"
                _ = fi13
                _ = v14
                _ = v15
            }
            var v16 uint64 = 42
            v16 -= v16
            _ = fi12
            _ = v16
        }
        v17 := GenS1{F0: 877, F1: -1}
        _ = v17
    }
    v18 := []int{892, -139, 0, 32768}
    for fi19 := 0; fi19 < 4; fi19++ {
        var v20 string = ("\n\t" + "hello")
        var v21 float64 = -2.25
        var v22 float64 = v21
        _ = fi19
        _ = v20
        _ = v21
        _ = v22
    }
    var v23 bool = false
    fmt.Println(v23)
    _ = v18
    _ = v23
    return ""
}

func main() {
    v1 := []string{"\n\t", "", ("" + ("foo bar" + "\n\t")), "\n\t"}
    v2 := make(map[int]int)
    var v3 uint = 18446744073709551615
    var v4 uint16 = uint16(v3)
    v5 := make(map[uint32]uint32)
    var v6 uint32 = 1
    var v24 int16 = 1000
    v25 := GenS1{F0: 966, F1: 42}
    var v26 int16 = (v24 << (((v3 << (((v3 % ((v3) | 1))) & 63))) & 63))
    v27 := GenS1{F0: 837, F1: 1}
    fmt.Println(v26)
    var v28 uint64 = 2
    v1 = append(v1, "\n\t")
    for fi29 := 0; fi29 < 5; fi29++ {
        for fi30 := 0; fi30 < 2; fi30++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            var v31 int32 = int32(v28)
            _ = fi30
            _ = v31
        }
        v28 = uint64(v3)
        var v32 int32 = -807
        _ = fi29
        _ = v32
    }
    v28 = (v28 & (v28 % ((1) | 1)))
    v3 += (v3 >> (((^(v3))) & 63))
    var v33 uint8 = 0
    _ = len(v5)
    v34 := GenS2{F0: 1}
    var v35 uint = (v3 * (v3 >> ((uint(v33)) & 63)))
    for fi36 := 0; fi36 < 11; fi36++ {
        var v37 int16 = (v26 | v24)
        var v38 int64 = 1
        _ = fi36
        _ = v37
        _ = v38
    }
    _ = len(v5)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
