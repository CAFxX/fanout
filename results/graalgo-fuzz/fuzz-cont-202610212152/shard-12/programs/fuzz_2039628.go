package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool }
type GenS2 struct { F0 int; F1 int16; F2 string }
func genFn0(p0 int16, p1 uint16) int {
    v7 := []int{}
    var v8 uint32 = 2147483647
    var v9 int32 = -936
    _ = len(v7)
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return 2147483647
}

func main() {
    var v1 string = ("" + ("hello" + ""))
    var v2 int64 = -505
    v3 := make(map[uint32]bool)
    var v4 uint = 18446744073709551615
    v5 := []uint32{}
    var v6 uint = (v4 << (((v4 | v4)) & 63))
    for fi10 := 0; fi10 < 2; fi10++ {
        switch (fi10 * -302) {
        case 127:
            if (false || (2 != 1)) {
                v6 += v6
            }
        case -407:
            var v11 int = (fi10 + genFn0(329, 42))
            _ = v11
        case 715:
            for fi12 := 0; fi12 < 10; fi12++ {
                fmt.Println(v2, v1, fi12)
                v1 = v1
                fmt.Println(v4)
                if (!(0.1 == 0.0)) {
                    var v13 int16 = 42
                    _ = v13
                }
                _ = fi12
            }
        }
        if ((2147483647 <= 875) && true) {
            for fi14 := 0; fi14 < 4; fi14++ {
                var v15 uint = (v6 << (((v4 + (v6 * v4))) & 63))
                var v16 int32 = 2147483647
                var v17 uint = v6
                var v18 uint8 = 139
                var v19 string = v1
                _ = fi14
                _ = v15
                _ = v16
                _ = v17
                _ = v18
                _ = v19
            }
            v4 = v6
        }
        _ = fi10
    }
    var v20 float64 = -2.25
    v20 = ((1e10 + float64(float32((3.14159 - ((-2.25 + (v20 / v20)) / v20))))) / (((1e10 + (((v20 - (float64(0.0) - 1e10)) * v20) * ((float64(0.0) + 1.5) / v20))) / float64((((0.1 + 3.14159) + (float32(v20) / 1.0)) / 0.1))) - v20))
    if (127 == -42) {
        var v21 uint16 = 1
        v22 := GenS1{F0: false, F1: (121 >= -1)}
        if (100 <= 108) {
            var v23 uint8 = uint8(v4)
            var v24 uint8 = (v23 / ((29) | 1))
            var v25 int32 = 571
            _ = v23
            _ = v24
            _ = v25
        } else {
            _ = len(v3)
            for fi26 := 0; fi26 < 12; fi26++ {
                if (255 > (v21 - 1)) {
                    switch "a" {
                    case "a":
                        var v27 uint16 = 255
                        _ = v27
                    }
                    switch 4294967295 {
                    case 9223372036854775807:
                        for fi28 := 0; fi28 < 8; fi28++ {
                            var v29 float32 = 3.14159
                            v30 := GenS1{F0: false, F1: true}
                            _ = fi28
                            _ = v29
                            _ = v30
                        }
                    }
                }
                var v31 int64 = 1
                _ = fi26
                _ = v31
            }
        }
        _ = v21
        _ = v22
    } else {
        v32 := GenS1{F0: false, F1: (v6 > (v6 | v6))}
        _ = v32
    }
    if (false || (180 != -1)) {
        v4 = v6
    } else {
        v33 := GenS2{F0: genFn0(-1, 2), F1: -509, F2: v1}
        _ = v33
    }
    v34 := GenS2{F0: 9223372036854775807, F1: 0, F2: v1}
    var v35 int8 = -1
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    fmt.Println(v20, v4)
    delete(v3, 0)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
