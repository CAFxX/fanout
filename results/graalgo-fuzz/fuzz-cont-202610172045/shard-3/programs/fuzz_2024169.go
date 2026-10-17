package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint16; F1 uint16 }
func genFn0(p0 int32, p1 uint8) int32 {
    switch ((("" + ("hello" + "hello")) + "foo bar") + "foo bar") {
    case "foo bar":
        p1 = p1
    case "a":
        for fi4 := 0; fi4 < 7; fi4++ {
            v5 := GenS1{F0: (fi4 * 127)}
            var v6 int32 = (p0 ^ 1)
            var v7 int32 = v6
            _ = fi4
            _ = v5
            _ = v6
            _ = v7
        }
    }
    v8 := []int{1, 2147483647}
    _ = p0
    _ = p1
    _ = v8
    return 42
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int8 = -110
    v3 := make(map[int]uint32)
    var v9 int = 127
    for fi10 := 0; fi10 < 4; fi10++ {
        fmt.Println(v9, fi10)
        _ = fi10
    }
    switch v9 {
    case -846:
        v11 := GenS1{F0: 2147483647}
        _ = v11
    }
    var v12 float64 = (0.0 / 0.1)
    for fi13 := 0; fi13 < 20; fi13++ {
        var v14 uint64 = 1
        v15 := []int{}
        switch (v9 ^ 32768) {
        case 32767:
            var v16 uint32 = 513
            _ = v16
        case 128:
            switch fi13 {
            case 0:
                for fi17 := 0; fi17 < 11; fi17++ {
                    v18 := GenS1{F0: fi17}
                    fmt.Println(fi13, fi17)
                    var v19 uint8 = 127
                    _ = fi17
                    _ = v18
                    _ = v19
                }
            case 762:
                for fi20 := 0; fi20 < 11; fi20++ {
                    var v21 int32 = 0
                    _ = fi20
                    _ = v21
                }
            case 128:
                v12 -= v12
            default:
                if (32767 != 65535) {
                    var v22 uint = 65535
                    v23 := GenS2{F0: 255, F1: 42}
                    var v24 bool = (!(!false))
                    _ = v22
                    _ = v23
                    _ = v24
                }
            }
        }
        _ = fi13
        _ = v14
        _ = v15
    }
    v2 -= -1
    switch 781 {
    case 336:
        v1[384] = false
    case 128:
        fmt.Println(v2, v2, v2)
    default:
        var v25 bool = (18446744073709551615 != 18446744073709551615)
        _ = v25
    }
    v26 := GenS2{F0: 255, F1: 255}
    fmt.Println(v2)
    v9 -= v9
    v27 := GenS2{F0: 1, F1: 0}
    v28 := []int{(v9 / ((v9) | 1)), v9}
    if false {
        if (659 >= 1) {
            v9 += (v9 & 42)
            v2 = -71
            for fi29 := 0; fi29 < 6; fi29++ {
                fmt.Println(v12, fi29, v2)
                v2 += (v2 << ((255) & 63))
                _ = fi29
            }
        } else {
            for fi30 := 0; fi30 < 16; fi30++ {
                v2 += v2
                _ = fi30
            }
            v12 = (float64(3.14159) * 0.1)
        }
    } else {
        v9 = v9
        v9 = v9
        v31 := v3[-2147483648]
        _ = v31
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v9
}
