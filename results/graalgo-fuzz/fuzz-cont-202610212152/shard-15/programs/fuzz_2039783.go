package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64; F2 uint64 }
type GenS2 struct { F0 int64; F1 uint; F2 int64 }
func genFn0(p0 string, p1 uint) int32 {
    v6 := GenS1{F0: 32767, F1: 784, F2: 365}
    var v7 uint8 = 0
    if (-9223372036854775808 >= 42) {
        switch p0 {
        case "a":
            var v8 uint16 = 65535
            _ = v8
        }
        if false {
            fmt.Println(p1, p0, v7)
            v9 := []int{42, 2147483647, -824, 4294967295}
            if len(v9) >= 2 {
                v9 = v9[1:len(v9)]
            }
            _ = v9
        }
        for fi10 := 0; fi10 < 14; fi10++ {
            v7 += v7
            v11 := []int{1}
            var v12 bool = (p1 <= p1)
            _ = fi10
            _ = v11
            _ = v12
        }
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 68
}
func genFn1() int8 {
    var v13 uint32 = 2147483647
    var v14 int32 = 42
    _ = v13
    _ = v14
    return 42
}

func main() {
    var v1 int16 = 0
    var v2 uint8 = 127
    v3 := make(map[uint32]bool)
    var v4 uint = 2
    var v5 int32 = int32(v4)
    v4 += (v4 & (v4 ^ v4))
    var v15 uint64 = 4294967295
    if true {
        v15 = v15
        v4 = (v4 & 42)
    } else {
        v5 = (-(v5))
    }
    var v16 float32 = float32((1.5 / ((1.5 - float64(((3.14159 - (1e10 / (3.14159 * float32(0.1)))) * (float32(float64(0.1)) - (3.14159 / 1.5))))) + 1e10)))
    delete(v3, 42)
    for fi17 := 0; fi17 < 20; fi17++ {
        v16 = 0.1
        var v18 bool = (388 >= 589)
        _ = fi17
        _ = v18
    }
    if false {
        var v19 int8 = 1
        var v20 uint64 = uint64(v2)
        _ = v19
        _ = v20
    }
    v21 := []int{}
    var v22 int32 = v5
    if true {
        fmt.Println(v2, v2)
        var v23 bool = ((-591 > 9223372036854775807) || (((0.1 + 1.5) + 0.0) < ((-2.25 / ((1.5 + 0.1) / float64(float32(1.5)))) * (0.1 + (3.14159 / float64(((v16 * 1.5) / (v16 * (float32(-2.25) - ((-2.25 * v16) + (v16 * 0.0)))))))))))
        v5 += (v22 | int32(v2))
        _ = v23
    }
    delete(v3, 125)
    for fi24 := 0; fi24 < 20; fi24++ {
        if false {
            v4 += 255
            if len(v21) > 0 {
                v21[0] = (fi24 | fi24)
            }
            var v25 uint32 = 2
            _ = v25
        } else {
            var v26 uint32 = 42
            v21 = append(v21, (fi24 / (((fi24 | int(v15))) | 1)))
            fmt.Println(v16)
            _ = v26
        }
        if (true && true) {
            v27 := GenS2{F0: 659, F1: v4, F2: -484}
            for fi28 := 0; fi28 < 12; fi28++ {
                v29 := GenS2{F0: 1, F1: v4, F2: -9223372036854775808}
                v22 -= (v5 + v22)
                _ = fi28
                _ = v29
            }
            _ = v27
        }
        var v30 uint8 = v2
        _ = fi24
        _ = v30
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
