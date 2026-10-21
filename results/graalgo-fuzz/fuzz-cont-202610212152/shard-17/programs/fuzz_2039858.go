package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 uint }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int8; F1 uint8 }
func genFn0() uint8 {
    for fi4 := 0; fi4 < 9; fi4++ {
        var v5 int32 = 130
        _ = fi4
        _ = v5
    }
    switch -234 {
    case 4294967295:
        var v6 int32 = -890
        _ = v6
    case 42:
        if true {
            v7 := GenS1{F0: 65535, F1: 228, F2: 255}
            var v8 uint32 = 2
            _ = v7
            _ = v8
        } else {
            var v9 int32 = -2147483648
            var v10 uint = 65535
            fmt.Println(v10, v10)
            _ = v10
            _ = v9
        }
    case -2147483648:
        for fi11 := 0; fi11 < 19; fi11++ {
            v12 := []int{}
            fmt.Println(fi11)
            for fi13 := 0; fi13 < 3; fi13++ {
                var v14 bool = (42 < 42)
                _ = len(v12)
                _ = fi13
                _ = v14
            }
            _ = fi11
            _ = v12
        }
    }
    return 230
}
func genFn1(p0 int64) int32 {
    if (94 < -936) {
        var v15 int64 = (p0 >> ((563) & 63))
        _ = v15
    } else {
        var v16 float64 = ((3.14159 * ((1e10 * 1.5) / ((0.0 - 3.14159) * (1e10 / -2.25)))) + 1e10)
        _ = v16
    }
    var v17 float64 = -2.25
    _ = p0
    _ = v17
    return 1
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int = 2147483647
    var v3 float32 = ((float32(0.0) + 3.14159) / (0.0 - float32(((3.14159 - 1e10) / (((float64(3.14159) * (0.0 - float64(1.5))) - (0.0 - 0.1)) + ((1e10 + -2.25) / -2.25))))))
    v18 := GenS2{F0: 42}
    fmt.Println(v3, v3)
    v3 = (v3 * v3)
    v3 -= ((((0.0 - (v3 + float32(((float64(v3) + float64(float32(float64(v3)))) + -2.25)))) * (float32(0.0) * ((v3 + 1e10) - 0.1))) / v3) * 1e10)
    switch v2 {
    case 535:
        var v19 float32 = 1.5
        _ = v19
    case 127:
        if (-9223372036854775808 <= -9223372036854775808) {
            for fi20 := 0; fi20 < 19; fi20++ {
                var v21 uint64 = 978
                var v22 uint16 = 1
                v21 = v21
                _ = fi20
                _ = v21
                _ = v22
            }
            if (2 <= 711) {
                v23 := GenS3{F0: -1, F1: genFn0()}
                v3 = ((v3 / ((1.5 * v3) / (float32((1.5 - -2.25)) + (float32((3.14159 / 0.1)) * 0.0)))) * ((0.1 / -2.25) * (v3 / v3)))
                _ = v23
            } else {
                var v24 string = ("a" + "foo bar")
                fmt.Println(v3)
                _ = v24
            }
            fmt.Println(v3, v3)
        }
    case -611:
        v25 := []int{(-(v2)), (v2 + -237), (v2 / ((9223372036854775807) | 1)), 908}
        _ = v25
    }
    v3 -= 0.1
    switch ("a" + "\n\t") {
    case "":
        v26 := GenS3{F0: -1, F1: 42}
        _ = v26
    }
    v3 = 1.5
    delete(v1, 2)
    var v27 int16 = 237
    v27 += -387
    v27 = (-(v27))
    v28 := []int{32767, 32768, 2147483647, v2}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v18
    _ = v2
    _ = v27
    _ = v28
    _ = v3
}
