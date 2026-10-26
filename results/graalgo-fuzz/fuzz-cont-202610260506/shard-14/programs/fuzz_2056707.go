package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int; F2 uint64 }
type GenS2 struct { F0 uint8; F1 uint32 }
type GenS3 struct { F0 int32; F1 int16 }
func genFn0() int16 {
    var v7 int32 = 1
    v7 = v7
    v8 := []int{-2147483648, 0}
    var v9 string = ""
    if len(v8) > 0 {
        v8[0] = 1
    }
    _ = v7
    _ = v8
    _ = v9
    return -754
}
func genFn1(p0 int8, p1 string) uint32 {
    p0 += (^(p0))
    var v10 uint = 0
    v10 = v10
    switch 127 {
    case -9223372036854775808:
        v11 := []int{418}
        _ = v11
    case -778:
        v12 := []int{-520, 269}
        _ = v12
    case 1:
        v13 := GenS1{F0: -252, F1: 42, F2: 684}
        _ = v13
    default:
        switch 128 {
        case 0:
            p0 = (^(p0))
        case 42:
            p0 -= p0
        }
    }
    v10 = (v10 << (((v10 - 42)) & 63))
    _ = p0
    _ = p1
    _ = v10
    return 116
}

func main() {
    v1 := make(map[uint32]int)
    v2 := []uint32{2147483648}
    v3 := make(map[int]int)
    var v4 float32 = 0.0
    var v5 uint = 655
    var v6 int16 = int16(v5)
    var v14 string = "hello"
    v2 = append(v2, 1)
    v4 = 3.14159
    fmt.Println(v14, v4)
    var v15 int16 = (v6 | (v6 << ((v5) & 63)))
    v4 += (float32(3.14159) * (v4 - v4))
    if len(v2) > 0 {
        v2[0] = 2
    }
    var v16 int64 = 1
    v14 = v14
    if false {
        var v17 int32 = 1
        for fi18 := 0; fi18 < 9; fi18++ {
            switch "" {
            case "a":
                var v19 bool = true
                _ = v19
            case "hello":
                var v20 float32 = (-2.25 - float32(1.5))
                _ = v20
            case "\n\t":
                var v21 float64 = (3.14159 + (1.5 / float64(v4)))
                _ = v21
            default:
                v16 += -1
            }
            _ = fi18
        }
        _ = v17
    }
    var v22 uint16 = 838
    switch v14 {
    case "foo bar":
        v4 = (-2.25 + (float32(((float64(((v4 * 3.14159) - (1e10 / (float32(((-2.25 - 1e10) + 0.1)) - (v4 + (0.0 * v4)))))) + (1e10 - 0.0)) * (((((float64(v4) - (((-2.25 - 1.5) / -2.25) * (3.14159 - 3.14159))) + float64((0.1 * (v4 / float32(-2.25))))) - 1e10) * (0.0 * 0.1)) + (float64(0.0) + -2.25)))) - float32((float64(v4) / float64(3.14159)))))
    case "a":
        var v23 int64 = (v16 - (v16 % (((^(v16))) | 1)))
        _ = v23
    case "\n\t":
        v24 := GenS2{F0: 232, F1: 572}
        _ = v24
    }
    switch v14 {
    case "foo bar":
        var v25 float64 = 0.0
        _ = v25
    case "":
        var v26 uint16 = v22
        _ = v26
    case "\n\t":
        for fi27 := 0; fi27 < 4; fi27++ {
            for fi28 := 0; fi28 < 4; fi28++ {
                v29 := GenS3{F0: 2147483647, F1: 1000}
                v30 := GenS3{F0: 1, F1: -32768}
                v22 -= 65535
                var v31 int = (fi28 % (((fi27 << (((v5 & 65535)) & 63))) | 1))
                _ = fi28
                _ = v29
                _ = v30
                _ = v31
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            var v32 string = "a"
            _ = fi27
            _ = v32
        }
    }
    fmt.Println(v6, v22)
    for fi33 := 0; fi33 < 16; fi33++ {
        switch (fi33 + 307) {
        case 9223372036854775807:
            v4 -= (1e10 / v4)
        case 128:
            if ((-2.25 <= 0.1) || (((((("hello" + "") + v14) + v14) + v14) + v14) >= v14)) {
                var v34 uint8 = 42
                v22 += v22
                _ = v34
            } else {
                for fi35 := 0; fi35 < 1; fi35++ {
                    v36 := v3[fi35]
                    _ = fi35
                    _ = v36
                }
                fmt.Println(v6)
                var v37 int32 = 0
                _ = v37
            }
        }
        _ = fi33
    }
    var v38 string = "hello"
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v22
    _ = v3
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
