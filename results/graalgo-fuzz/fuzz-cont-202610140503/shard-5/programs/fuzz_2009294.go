package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16 }
type GenS2 struct { F0 int32 }
func genFn0() int {
    var v7 uint64 = 87
    var v8 float64 = float64(float32((0.1 - (((3.14159 / 0.1) / 3.14159) + -2.25))))
    fmt.Println(v7, v7)
    var v9 uint32 = 42
    v10 := []int{128, -455, -343}
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return int(v9)
}
func genFn1(p0 uint32) bool {
    var v11 float64 = 0.0
    v12 := GenS1{F0: (!(1 <= -391)), F1: 65535}
    var v13 uint8 = 8
    _ = p0
    _ = v11
    _ = v12
    _ = v13
    return ((genFn0() == 127) || true)
}

func main() {
    var v1 int8 = -15
    var v2 uint8 = 179
    var v3 int16 = 1
    var v4 string = "hello"
    var v5 uint32 = 779
    v6 := make(map[uint32]bool)
    v14 := v6[(v5 - (v5 ^ v5))]
    v3 -= 831
    v15 := []int{}
    _ = len(v15)
    switch 42 {
    case 70:
        if (-9223372036854775808 < -9223372036854775808) {
            switch "hello" {
            case "\n\t":
                v16 := GenS2{F0: 2147483647}
                _ = v16
            case "hello":
                var v17 float32 = float32(0.1)
                _ = v17
            case "a":
                v5 -= (v5 / ((v5) | 1))
            }
            for fi18 := 0; fi18 < 5; fi18++ {
                v4 = ""
                var v19 float64 = -2.25
                _ = fi18
                _ = v19
            }
        }
    }
    for fi20 := 0; fi20 < 9; fi20++ {
        v1 = v1
        for fi21 := 0; fi21 < 17; fi21++ {
            fmt.Println(fi21, v14, v3)
            _ = fi21
        }
        v5 = v5
        _ = fi20
    }
    fmt.Println(v14, v14)
    if (42 != 65535) {
        fmt.Println(v4, v3, v3)
    }
    var v22 uint64 = 2
    var v23 uint16 = uint16(v22)
    if v14 {
        switch v4 {
        case "hello":
            fmt.Println(v1, v4)
        case "a":
            v24 := GenS1{F0: false, F1: 111}
            _ = v24
        case "\n\t":
            fmt.Println(v2)
        }
        var v25 uint16 = v23
        _ = v25
    } else {
        for fi26 := 0; fi26 < 3; fi26++ {
            v4 = "a"
            var v27 uint8 = 2
            _ = fi26
            _ = v27
        }
    }
    for fi28 := 0; fi28 < 13; fi28++ {
        v14 = v14
        _ = fi28
    }
    var v29 int = 9223372036854775807
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v22
    _ = v23
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
