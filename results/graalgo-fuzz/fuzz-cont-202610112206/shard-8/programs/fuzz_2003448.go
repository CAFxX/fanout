package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint }
func genFn0(p0 int16, p1 uint64, p2 int32) string {
    p1 = p1
    v4 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return "a"
}
func genFn1(p0 uint32, p1 int8) uint16 {
    var v5 uint64 = 42
    for fi6 := 0; fi6 < 20; fi6++ {
        p1 = int8(v5)
        _ = fi6
    }
    fmt.Println(v5)
    _ = p0
    _ = p1
    _ = v5
    return 65535
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[string]bool)
    v3 := []uint32{}
    switch 484 {
    case -2147483648:
        v7 := GenS1{F0: 127, F1: 32767}
        _ = v7
    default:
        for fi8 := 0; fi8 < 16; fi8++ {
            v9 := GenS2{F0: false}
            _ = fi8
            _ = v9
        }
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    switch 127 {
    case -220:
        for fi10 := 0; fi10 < 4; fi10++ {
            var v11 int8 = 5
            v11 -= v11
            _ = fi10
            _ = v11
        }
    case 32767:
        var v12 float32 = 3.14159
        _ = v12
    default:
        var v13 int8 = 42
        _ = v13
    }
    for fi14 := 0; fi14 < 8; fi14++ {
        var v15 float32 = ((float32(0.0) * -2.25) * float32(1.5))
        v3 = append(v3, 1)
        _ = fi14
        _ = v15
    }
    v3 = append(v3, 1)
    if ((1.5 <= 3.14159) && false) {
        var v16 int16 = 32767
        v16 += (^(v16))
        for fi17 := 0; fi17 < 16; fi17++ {
            for fi18 := 0; fi18 < 12; fi18++ {
                v19 := GenS1{F0: 255, F1: fi18}
                _ = fi18
                _ = v19
            }
            _ = fi17
        }
        _ = v16
    }
    for fi20 := 0; fi20 < 1; fi20++ {
        var v21 string = (((("\n\t" + "hello") + "hello") + "hello") + "foo bar")
        _ = fi20
        _ = v21
    }
    var v22 uint32 = 2147483648
    for fi23 := 0; fi23 < 10; fi23++ {
        for fi24 := 0; fi24 < 11; fi24++ {
            v22 += 75
            _ = fi24
        }
        _ = fi23
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi25 := 0; fi25 < 1; fi25++ {
        v26 := GenS1{F0: 245, F1: (fi25 & (^(fi25)))}
        var v27 uint32 = (v22 >> ((682) & 63))
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi25
        _ = v26
        _ = v27
    }
    for fi28 := 0; fi28 < 18; fi28++ {
        if (1000 > -857) {
            v22 = v22
        } else {
            v29 := GenS3{F0: 18446744073709551615}
            for fi30 := 0; fi30 < 13; fi30++ {
                var v31 int64 = 41
                var v32 int16 = -1
                if (-120 >= -128) {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
                var v33 uint32 = v22
                _ = fi30
                _ = v31
                _ = v32
                _ = v33
            }
            _ = v29
        }
        var v34 float64 = 0.1
        _ = fi28
        _ = v34
    }
    v35 := GenS1{F0: 42, F1: 127}
    if true {
        fmt.Println(v22)
        if false {
            v36 := GenS3{F0: 255}
            v22 = (v22 | 490)
            var v37 float64 = 1e10
            _ = v36
            _ = v37
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v35
}
