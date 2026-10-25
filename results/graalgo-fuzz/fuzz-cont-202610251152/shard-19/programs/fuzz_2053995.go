package main

import (
    "fmt"
)

type GenS1 struct { F0 int }

func main() {
    var v1 uint16 = 65535
    var v2 float32 = 0.0
    var v3 int32 = -619
    var v4 uint16 = (v1 + 65535)
    v5 := []uint32{1, 473, 2147483648, uint32(v4)}
    var v6 int = -9223372036854775808
    var v7 int64 = 904
    for fi8 := 0; fi8 < 1; fi8++ {
        for fi9 := 0; fi9 < 12; fi9++ {
            var v10 int64 = (^(v7))
            _ = fi9
            _ = v10
        }
        var v11 uint16 = v1
        _ = fi8
        _ = v11
    }
    if false {
        for fi12 := 0; fi12 < 20; fi12++ {
            if (false || false) {
                var v13 float64 = (0.1 * (((1e10 - 0.1) * 1.5) / ((0.0 - (float64(v2) - 3.14159)) / ((float64((float32((((float64(float32(0.1)) / ((3.14159 / 0.1) - 1e10)) * (1.5 + (1.5 - 1e10))) / -2.25)) + (v2 + (float32(1e10) + ((v2 + float32(float64(float32((3.14159 * (1e10 / ((1.5 / ((1e10 * 1e10) * (1e10 * (((0.1 - 1e10) * (1e10 / 1.5)) - 1.5)))) + -2.25))))))) / 3.14159))))) / 1e10) / ((0.1 - -2.25) - float64(float32(-2.25)))))))
                var v14 float32 = v2
                var v15 int = (fi12 & fi12)
                _ = v13
                _ = v14
                _ = v15
            }
            var v16 uint = 4294967295
            _ = fi12
            _ = v16
        }
    } else {
        if true {
            var v17 int = 0
            _ = v17
        } else {
            var v18 string = "hello"
            var v19 int16 = 42
            v4 = v4
            _ = v18
            _ = v19
        }
        if (!false) {
            var v20 string = ""
            _ = v20
        }
        _ = len(v5)
    }
    if true {
        fmt.Println(v6, v2, v3)
        v2 -= 1.5
    } else {
        fmt.Println(v3)
        fmt.Println(v3, v1)
        var v21 int8 = -115
        _ = v21
    }
    for fi22 := 0; fi22 < 10; fi22++ {
        var v23 uint = 370
        _ = fi22
        _ = v23
    }
    fmt.Println(v1)
    var v24 uint8 = 1
    var v25 int8 = 0
    fmt.Println(v24, v25)
    v26 := GenS1{F0: v6}
    if ((false && ((float64(0.0) + ((1e10 + (1.5 / (0.1 + 0.0))) + 0.0)) > ((-2.25 * 0.0) / ((float64(((v2 * 0.1) / ((((v2 - -2.25) / (v2 * -2.25)) * v2) + ((v2 / v2) - (3.14159 - ((v2 - float32(1e10)) * v2)))))) / -2.25) + (((0.0 * 0.0) + float64(0.0)) - float64(v2)))))) || true) {
        var v27 string = "\n\t"
        if (int8(v3) <= v25) {
            var v28 bool = false
            v6 = 32768
            _ = v28
        }
        var v29 int32 = v3
        _ = v27
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
