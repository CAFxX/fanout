package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64; F2 uint16 }
func genFn0(p0 uint32, p1 uint8, p2 int) uint64 {
    var v7 int16 = -753
    v8 := []int{(p2 % (((p2 & int(p1))) | 1)), (^(p2)), (p2 % (((p2 / ((p2) | 1))) | 1))}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 42
}
func genFn1() uint8 {
    for fi9 := 0; fi9 < 18; fi9++ {
        var v10 float32 = (1e10 + ((((0.1 - ((float32(1e10) * float32(-2.25)) + 1.5)) / -2.25) + 0.1) * 1e10))
        for fi11 := 0; fi11 < 11; fi11++ {
            v12 := []int{127}
            _ = fi11
            _ = v12
        }
        v10 += 0.0
        _ = fi9
        _ = v10
    }
    v13 := []int{}
    switch "\n\t" {
    case "\n\t":
        var v14 int8 = 96
        _ = v14
    case "a":
        var v15 uint64 = 2
        _ = v15
    case "":
        var v16 int32 = 0
        _ = v16
    }
    _ = v13
    return 7
}

func main() {
    var v1 int32 = 2147483647
    var v2 int16 = -1
    v3 := []uint32{0}
    var v4 uint32 = 65535
    v5 := make(map[uint32]uint32)
    var v6 int = 0
    switch v6 {
    case -586:
        v6 += 32768
    case 0:
        var v17 uint16 = 590
        _ = v17
    case -854:
        var v18 uint = 796
        _ = v18
    default:
        switch -593 {
        case 4294967295:
            fmt.Println(v2, v6)
        }
    }
    if true {
        for fi19 := 0; fi19 < 5; fi19++ {
            for fi20 := 0; fi20 < 13; fi20++ {
                v21 := GenS1{F0: ((2 == 0) || false), F1: -1, F2: 103}
                var v22 float64 = 0.1
                _ = fi20
                _ = v21
                _ = v22
            }
            for fi23 := 0; fi23 < 20; fi23++ {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                fmt.Println(fi23)
                _ = fi23
            }
            if (!("hello" > "hello")) {
                v24 := GenS1{F0: (!false), F1: 42, F2: 0}
                var v25 int32 = v1
                _ = v24
                _ = v25
            } else {
                var v26 bool = true
                var v27 int64 = -810
                _ = v26
                _ = v27
            }
            _ = fi19
        }
    }
    switch (v6 % (((v6 & (v6 >> ((0) & 63)))) | 1)) {
    case 2147483647:
        var v28 uint32 = (v4 << ((1) & 63))
        _ = v28
    }
    for fi29 := 0; fi29 < 13; fi29++ {
        for fi30 := 0; fi30 < 1; fi30++ {
            v1 -= (v1 ^ v1)
            _ = fi30
        }
        for fi31 := 0; fi31 < 9; fi31++ {
            var v32 uint8 = 42
            v33 := GenS1{F0: (9223372036854775807 < -1), F1: 778, F2: 255}
            _ = fi31
            _ = v32
            _ = v33
        }
        _ = fi29
    }
    for fi34 := 0; fi34 < 2; fi34++ {
        fmt.Println(v4)
        _ = fi34
    }
    var v35 string = ("hello" + ((((("\n\t" + "hello") + "hello") + ("\n\t" + "a")) + ("a" + "a")) + (("" + ("" + "")) + "hello")))
    v36 := GenS1{F0: false, F1: int64(v6), F2: 32767}
    for fi37 := 0; fi37 < 15; fi37++ {
        v3 = append(v3, 126)
        for fi38 := 0; fi38 < 20; fi38++ {
            v39 := v5[1]
            v4 = (v4 << ((18446744073709551615) & 63))
            _ = fi38
            _ = v39
        }
        _ = fi37
    }
    var v40 uint = 416
    var v41 int16 = (v2 * 625)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
}
