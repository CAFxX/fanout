package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
func genFn0(p0 int64) int16 {
    v7 := []int{9223372036854775807, -1, 4294967295, 32768}
    p0 = p0
    _ = p0
    _ = v7
    return 32767
}
func genFn1() int8 {
    if false {
        var v8 string = "hello"
        _ = v8
    }
    if (1 == 334) {
        var v9 int32 = 178
        if false {
            for fi10 := 0; fi10 < 16; fi10++ {
                v11 := []int{fi10}
                _ = fi10
                _ = v11
            }
            fmt.Println(v9)
        }
        var v12 float64 = (((0.1 * (1.5 * 0.1)) - 3.14159) - (3.14159 * (3.14159 * ((0.0 / 0.1) + (1e10 * (-2.25 - float64(-2.25)))))))
        _ = v12
        _ = v9
    }
    var v13 int16 = -32768
    _ = v13
    return 0
}

func main() {
    var v1 int32 = -2147483648
    v2 := make(map[string]uint32)
    var v3 uint16 = 313
    v4 := make(map[uint32]bool)
    v5 := make(map[string]int)
    var v6 int8 = 42
    var v14 int8 = v6
    v15 := GenS1{F0: v3}
    v16 := []int{42, 1}
    var v17 float32 = 0.0
    for fi18 := 0; fi18 < 3; fi18++ {
        for fi19 := 0; fi19 < 13; fi19++ {
            var v20 uint8 = 255
            var v21 uint8 = 197
            v14 -= v6
            _ = fi19
            _ = v20
            _ = v21
        }
        _ = fi18
    }
    fmt.Println(v3)
    var v22 uint16 = 0
    v1 -= (v1 ^ 180)
    if (127 > 0) {
        var v23 int64 = 42
        if (0 <= 1) {
            var v24 uint16 = v3
            for fi25 := 0; fi25 < 20; fi25++ {
                v5[""] = fi25
                v3 += v22
                _ = fi25
            }
            if len(v16) > 0 {
                v16[0] = int(v1)
            }
            _ = v24
        }
        _ = v23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
