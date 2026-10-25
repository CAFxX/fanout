package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
type GenS2 struct { F0 uint16; F1 int }
type GenS3 struct { F0 uint32; F1 uint32; F2 int }
type GenS4 struct { F0 int64 }

func main() {
    var v1 bool = true
    v2 := []int{-691, 127}
    v3 := []uint32{2, 2147483648}
    if v1 {
        switch 32768 {
        case -753:
            var v4 uint8 = 42
            _ = v4
        default:
            fmt.Println(v1)
        }
    }
    if v1 {
        v1 = v1
        v1 = false
    } else {
        v1 = (!v1)
    }
    v1 = v1
    v1 = (3.14159 != 3.14159)
    fmt.Println(v1)
    for fi5 := 0; fi5 < 6; fi5++ {
        v1 = v1
        _ = fi5
    }
    for fi6 := 0; fi6 < 12; fi6++ {
        if false {
            var v7 float64 = (0.0 * 3.14159)
            switch (fi6 % ((fi6) | 1)) {
            case 42:
                switch fi6 {
                case -616:
                    for fi8 := 0; fi8 < 15; fi8++ {
                        var v9 int8 = -1
                        _ = fi8
                        _ = v9
                    }
                default:
                    v7 = v7
                }
            }
            var v10 int = fi6
            _ = v10
            _ = v7
        } else {
            if ("hello" != "") {
                fmt.Println(fi6)
            }
        }
        v11 := GenS1{F0: 18446744073709551615, F1: 598}
        _ = fi6
        _ = v11
    }
    _ = len(v3)
    v1 = (18446744073709551615 <= 307)
    v1 = true
    for fi12 := 0; fi12 < 20; fi12++ {
        var v13 string = ""
        for fi14 := 0; fi14 < 5; fi14++ {
            var v15 string = v13
            _ = fi14
            _ = v15
        }
        _ = fi12
        _ = v13
    }
    v1 = true
    if (111 > 418) {
        var v16 uint = 490
        if false {
            v16 = (v16 & v16)
            switch "foo bar" {
            case "a":
                fmt.Println(v16, v16)
            }
            fmt.Println(v16, v16)
        } else {
            v17 := GenS2{F0: 1, F1: 32768}
            if false {
                v1 = ((1.5 * 0.1) == 1.5)
                v16 -= 65535
            }
            v18 := GenS1{F0: v16, F1: -1}
            _ = v17
            _ = v18
        }
        _ = v16
    } else {
        v19 := GenS3{F0: 0, F1: 0, F2: -1}
        v20 := GenS2{F0: 284, F1: 2147483647}
        if (-2.25 > ((3.14159 / ((1.5 / (-2.25 + float64(1.5))) - -2.25)) * (3.14159 + (-2.25 + (((float64(float32(0.0)) - 3.14159) - 3.14159) / (1e10 * float64(0.1))))))) {
            var v21 uint32 = 65535
            if (!v1) {
                v22 := GenS4{F0: -9223372036854775808}
                _ = v22
            } else {
                v21 = (v21 * v21)
                switch "hello" {
                case "hello":
                    v2 = append(v2, -467)
                }
                v1 = true
            }
            var v23 uint64 = 0
            _ = v21
            _ = v23
        }
        _ = v19
        _ = v20
    }
    var v24 uint16 = 65535
    _ = len(v3)
    v25 := GenS4{F0: -1}
    for fi26 := 0; fi26 < 19; fi26++ {
        fmt.Println(v1)
        _ = fi26
    }
    var v27 string = "\n\t"
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v27
    _ = v3
}
