package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32; F2 uint8 }

func main() {
    var v1 string = "hello"
    var v2 uint32 = 975
    var v3 uint16 = 32767
    v4 := make(map[int]uint32)
    var v5 uint = 815
    if false {
        if false {
            for fi6 := 0; fi6 < 12; fi6++ {
                v3 = v3
                _ = fi6
            }
        } else {
            var v7 float32 = 3.14159
            v3 = v3
            var v8 float32 = (((v7 - float32(float64((-2.25 - (v7 - 0.1))))) / (((float32((-2.25 * (0.0 - ((-2.25 * 3.14159) * ((3.14159 - 0.0) - 3.14159))))) + float32((3.14159 * 1e10))) - v7) * float32(1.5))) - v7)
            _ = v7
            _ = v8
        }
        for fi9 := 0; fi9 < 8; fi9++ {
            for fi10 := 0; fi10 < 3; fi10++ {
                v3 -= v3
                v11 := []int{-373}
                _ = fi10
                _ = v11
            }
            v2 = 2147483648
            _ = fi9
        }
    }
    v3 -= (v3 / (((^(v3))) | 1))
    var v12 uint32 = (v2 >> ((380) & 63))
    fmt.Println(v3, v12, v12)
    switch -2147483648 {
    case 938:
        v2 = 896
    case 2147483647:
        v4[1] = 276
    default:
        v13 := []int{-9223372036854775808, 32768, 922}
        _ = v13
    }
    if (!(!true)) {
        for fi14 := 0; fi14 < 3; fi14++ {
            v12 += v12
            switch (fi14 % (((fi14 + fi14)) | 1)) {
            case 128:
                v15 := []int{(fi14 / ((fi14) | 1)), (^(fi14)), 9223372036854775807}
                _ = v15
            }
            for fi16 := 0; fi16 < 3; fi16++ {
                if (255 >= v3) {
                    switch (fi14 + fi16) {
                    case 0:
                        v2 += v2
                    default:
                        var v17 bool = true
                        _ = v17
                    }
                    switch (v1 + v1) {
                    case "\n\t":
                        var v18 uint32 = uint32(v3)
                        _ = v18
                    }
                }
                _ = fi16
            }
            _ = fi14
        }
        v4[-1] = (v12 << ((v5) & 63))
    }
    if true {
        delete(v4, 127)
        if (!(false && (float32(1.5) <= ((1e10 * 0.1) * 1e10)))) {
            var v19 int16 = 0
            v20 := v4[32768]
            _ = v19
            _ = v20
        }
    }
    v21 := []int{1, -749}
    _ = len(v21)
    if ((v12 + 65535) != (v2 | v12)) {
        if (318 <= 18446744073709551615) {
            for fi22 := 0; fi22 < 1; fi22++ {
                var v23 uint = v5
                var v24 int64 = -9223372036854775808
                _ = fi22
                _ = v23
                _ = v24
            }
        }
        v25 := GenS1{F0: 1, F1: 2147483647, F2: 0}
        v3 -= (v3 * v3)
        _ = v25
    } else {
        _ = len(v4)
        v12 = (v2 & v12)
    }
    v3 += (v3 % (((v3 * v3)) | 1))
    v1 = (v1 + (((v1 + (((v1 + ("hello" + "a")) + "\n\t") + ("" + "hello"))) + "") + v1))
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
