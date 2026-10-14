package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8 }
type GenS2 struct { F0 uint16; F1 int8 }
func genFn0(p0 bool) int32 {
    if p0 {
        for fi4 := 0; fi4 < 15; fi4++ {
            fmt.Println(p0)
            p0 = (("a" != "\n\t") && true)
            if p0 {
                v5 := []int{(fi4 << ((1) & 63)), fi4, -1}
                fmt.Println(fi4)
                switch fi4 {
                case 818:
                    var v6 int32 = 0
                    _ = v6
                case 2147483647:
                    if (910 <= 4294967295) {
                        p0 = ((false && (128 == 127)) && p0)
                    } else {
                        p0 = (648 == 42)
                        if p0 {
                            var v7 string = "foo bar"
                            _ = v7
                        } else {
                            fmt.Println(fi4, p0)
                            for fi8 := 0; fi8 < 3; fi8++ {
                                p0 = true
                                fmt.Println(fi8, p0, fi8)
                                var v9 float32 = 1e10
                                _ = fi8
                                _ = v9
                            }
                        }
                    }
                case -1:
                    for fi10 := 0; fi10 < 12; fi10++ {
                        switch fi10 {
                        case 0:
                            p0 = p0
                        }
                        _ = fi10
                    }
                }
                _ = v5
            }
            _ = fi4
        }
        p0 = (595 < 557)
        fmt.Println(p0)
    } else {
        var v11 float32 = 1e10
        _ = v11
    }
    if ("hello" != "") {
        p0 = p0
    } else {
        var v12 float32 = ((0.0 * 1.5) + (3.14159 / ((3.14159 - 0.0) * 1e10)))
        fmt.Println(p0, p0)
        _ = v12
    }
    for fi13 := 0; fi13 < 20; fi13++ {
        var v14 int8 = -128
        v14 = (v14 % ((v14) | 1))
        v14 = v14
        _ = fi13
        _ = v14
    }
    v15 := []int{2147483647, 4294967295}
    _ = p0
    _ = v15
    return 772
}
func genFn1(p0 bool, p1 uint16) uint {
    if p0 {
        var v16 int8 = 35
        _ = v16
    }
    for fi17 := 0; fi17 < 2; fi17++ {
        v18 := []int{fi17}
        if len(v18) >= 2 {
            v18 = v18[1:len(v18)]
        }
        _ = fi17
        _ = v18
    }
    _ = p0
    _ = p1
    return 42
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint8 = 216
    var v3 uint16 = 1
    var v19 int = -785
    for fi20 := 0; fi20 < 9; fi20++ {
        fmt.Println(v19, v2, v3)
        v2 += uint8(v3)
        for fi21 := 0; fi21 < 20; fi21++ {
            var v22 uint32 = 806
            var v23 bool = (78 == uint(v22))
            v24 := GenS1{F0: v23, F1: 1}
            _ = fi21
            _ = v22
            _ = v23
            _ = v24
        }
        _ = fi20
    }
    var v25 uint16 = (v3 / ((v3) | 1))
    for fi26 := 0; fi26 < 6; fi26++ {
        fmt.Println(v19)
        v25 = (v3 / ((255) | 1))
        _ = fi26
    }
    v3 = (v25 >> ((65535) & 63))
    v3 = (v25 | (v3 >> ((42) & 63)))
    var v27 int16 = -1
    for fi28 := 0; fi28 < 6; fi28++ {
        for fi29 := 0; fi29 < 8; fi29++ {
            v30 := GenS2{F0: 255, F1: 29}
            var v31 string = "\n\t"
            if true {
                fmt.Println(fi28, v27)
                var v32 int16 = v27
                v1[2147483648] = false
                _ = v32
            }
            _ = fi29
            _ = v30
            _ = v31
        }
        _ = fi28
    }
    var v33 uint32 = 737
    v34 := GenS1{F0: false, F1: -128}
    if (true || ((float32((0.0 * float64((-2.25 - 3.14159)))) * 0.1) < float32(0.1))) {
        v33 = (v33 >> ((443) & 63))
        var v35 bool = true
        v35 = (!true)
        _ = v35
    }
    _ = len(v1)
    v27 += v27
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    _ = v1
    _ = v19
    _ = v2
    _ = v25
    _ = v27
    _ = v3
    _ = v33
    _ = v34
}
