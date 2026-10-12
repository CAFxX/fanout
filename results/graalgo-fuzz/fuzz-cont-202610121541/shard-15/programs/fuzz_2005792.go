package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16 }

func main() {
    var v1 uint = 75
    v2 := make(map[int]uint32)
    v3 := []uint32{224, 65535}
    var v4 float64 = ((float64(((float32((0.1 - (1.5 * -2.25))) + 1.5) - 1e10)) + 0.1) + -2.25)
    v5 := make(map[int]uint32)
    v6 := make(map[uint32]bool)
    fmt.Println(v4, v1)
    v7 := GenS1{F0: 80, F1: 41}
    var v8 int64 = 1
    for fi9 := 0; fi9 < 8; fi9++ {
        var v10 int = 468
        var v11 uint64 = 0
        switch 42 {
        case 1:
            v8 = v8
        }
        _ = fi9
        _ = v10
        _ = v11
    }
    fmt.Println(v1, v4)
    v1 = v1
    if (864 >= 0) {
        var v12 uint16 = 94
        switch "\n\t" {
        case "":
            v13 := GenS1{F0: -2147483648, F1: 32767}
            _ = v13
        case "hello":
            for fi14 := 0; fi14 < 16; fi14++ {
                v1 += 0
                _ = fi14
            }
        case "\n\t":
            if (!(v4 < -2.25)) {
                v8 = v8
                var v15 string = "hello"
                switch ("a" + "") {
                case "\n\t":
                    fmt.Println(v8, v1)
                case "a":
                    v6[2] = (true || true)
                }
                _ = v15
            }
        }
        var v16 float32 = 1e10
        _ = v12
        _ = v16
    } else {
        for fi17 := 0; fi17 < 19; fi17++ {
            v1 = 917
            _ = fi17
        }
        var v18 float64 = ((-2.25 + ((v4 + ((v4 + v4) - 0.0)) / (0.0 + (1e10 / (v4 - 0.0))))) * (v4 + v4))
        _ = v18
    }
    if len(v3) > 0 {
        v3[0] = 0
    }
    for fi19 := 0; fi19 < 11; fi19++ {
        v4 += ((1e10 + 1.5) * (((v4 * ((-2.25 * v4) * 0.1)) - v4) * (float64((-2.25 / (1.5 + (((0.0 + ((3.14159 + (-2.25 - (3.14159 / 1.0))) / ((((1e10 + 3.14159) - (float32((v4 / (1.5 / (-2.25 / v4)))) * 1e10)) / -2.25) + 0.1))) * 1e10) * 1e10)))) + v4)))
        _ = fi19
    }
    v4 -= 3.14159
    switch 38 {
    case 4294967295:
        if false {
            switch "" {
            case "\n\t":
                fmt.Println(v4)
            }
            v1 += 21
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        } else {
            for fi20 := 0; fi20 < 5; fi20++ {
                for fi21 := 0; fi21 < 15; fi21++ {
                    v22 := GenS1{F0: -2147483648, F1: 44}
                    _ = fi21
                    _ = v22
                }
                _ = fi20
            }
            v4 = 0.0
            v23 := GenS1{F0: 688, F1: 32767}
            _ = v23
        }
    }
    if (58 != 0) {
        v1 += v1
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
