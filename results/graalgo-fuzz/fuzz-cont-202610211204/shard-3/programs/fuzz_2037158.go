package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 int8 }
type GenS2 struct { F0 int16; F1 int; F2 uint8 }
type GenS3 struct { F0 int; F1 int8 }
func genFn0(p0 int64, p1 uint16, p2 int8) uint8 {
    if (!false) {
        p0 -= p0
        var v7 float32 = float32(1e10)
        for fi8 := 0; fi8 < 5; fi8++ {
            var v9 int = 815
            v10 := []int{(fi8 * -360), (v9 % ((fi8) | 1))}
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = v7
    }
    fmt.Println(p1, p2)
    v11 := []int{-530, 796, 42, 2147483647}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return 110
}
func genFn1(p0 uint8, p1 int, p2 uint16) int64 {
    v12 := []int{}
    switch (p1 / (((p1 << ((4294967295) & 63))) | 1)) {
    case 127:
        p1 -= p1
    case 9223372036854775807:
        p1 = p1
    case 128:
        if (false || true) {
            if len(v12) >= 2 {
                v12 = v12[1:len(v12)]
            }
            if (1 == 2147483647) {
                for fi13 := 0; fi13 < 20; fi13++ {
                    var v14 int16 = 42
                    switch fi13 {
                    case 32767:
                        v14 += int16(p0)
                    case 1:
                        for fi15 := 0; fi15 < 3; fi15++ {
                            switch "\n\t" {
                            case "foo bar":
                                if len(v12) >= 2 {
                                    v12 = v12[1:len(v12)]
                                }
                            }
                            var v16 uint64 = 1
                            v17 := GenS1{F0: p2, F1: 1, F2: int8(p2)}
                            _ = fi15
                            _ = v16
                            _ = v17
                        }
                    case 0:
                        v18 := GenS2{F0: (-(v14)), F1: (fi13 << ((42) & 63)), F2: p0}
                        _ = v18
                    default:
                        var v19 uint8 = (p0 << ((612) & 63))
                        _ = v19
                    }
                    _ = fi13
                    _ = v14
                }
            } else {
                p2 = p2
                fmt.Println(p0, p1)
                var v20 int32 = 411
                _ = v20
            }
        }
    }
    p0 = genFn0(0, p2, 127)
    var v21 uint32 = 2147483648
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v21
    return 1
}

func main() {
    v1 := make(map[int]uint32)
    var v2 uint8 = 128
    v3 := []uint32{1, 1, 2}
    var v4 int = -619
    var v5 int32 = 738
    v6 := make(map[uint32]string)
    for fi22 := 0; fi22 < 13; fi22++ {
        for fi23 := 0; fi23 < 11; fi23++ {
            fmt.Println(fi23)
            var v24 string = ("" + "a")
            _ = fi23
            _ = v24
        }
        _ = fi22
    }
    _ = len(v3)
    if true {
        if false {
            v25 := GenS3{F0: v4, F1: 127}
            _ = v25
        }
        if false {
            v4 = -9223372036854775808
            v5 = 2147483647
        } else {
            v4 = v4
            fmt.Println(v4, v5, v2)
        }
    }
    v2 += (^(v2))
    _ = len(v3)
    v2 -= v2
    v2 += v2
    var v26 uint8 = v2
    var v27 float64 = -2.25
    var v28 uint = 901
    v2 += (^(v26))
    v26 = genFn0(471, 978, 1)
    v27 += 1.5
    v29 := GenS1{F0: 43, F1: 775, F2: -1}
    var v30 int16 = 261
    v5 = -587
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
