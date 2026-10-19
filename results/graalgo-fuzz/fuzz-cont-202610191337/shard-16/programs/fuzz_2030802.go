package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint32; F1 int64; F2 uint32 }
type GenS4 struct { F0 int64; F1 int8; F2 uint8 }

func main() {
    var v1 int32 = 555
    v2 := make(map[uint32]int)
    v3 := make(map[string]string)
    v4 := make(map[uint32]int)
    v5 := make(map[string]uint32)
    var v6 uint = 752
    if true {
        v7 := v4[805]
        _ = v7
    }
    if (42 != 0) {
        if (1 <= 187) {
            var v8 int32 = (v1 & (-(v1)))
            v9 := GenS1{F0: 42}
            _ = v8
            _ = v9
        } else {
            fmt.Println(v6)
        }
        var v10 float64 = float64((0.1 * ((float32((1e10 + 1e10)) / 1e10) + (float32(-2.25) - ((((float32(0.1) * -2.25) + (1.5 * (float32((3.14159 - -2.25)) + 1e10))) * (3.14159 * -2.25)) + 0.0)))))
        fmt.Println(v6)
        _ = v10
    }
    v11 := v4[2]
    if false {
        if true {
            switch -691 {
            case 32768:
                if (65535 <= 863) {
                    for fi12 := 0; fi12 < 11; fi12++ {
                        if false {
                            for fi13 := 0; fi13 < 15; fi13++ {
                                v6 += (v6 % ((v6) | 1))
                                v1 += v1
                                _ = fi13
                            }
                        }
                        _ = fi12
                    }
                    v6 = v6
                }
            case 42:
                v14 := []int{4294967295, v11, (v11 << ((v6) & 63))}
                _ = v14
            case -838:
                var v15 uint = 255
                _ = v15
            }
            if false {
                v1 = (v1 ^ (v1 ^ (-(v1))))
                if ((1e10 == 1e10) || (true && true)) {
                    fmt.Println(v6, v11)
                }
            }
            switch v11 {
            case 32767:
                v16 := GenS2{F0: 42}
                _ = v16
            }
        } else {
            var v17 float64 = 1e10
            if false {
                var v18 int8 = -128
                fmt.Println(v6)
                _ = v18
            } else {
                for fi19 := 0; fi19 < 12; fi19++ {
                    v20 := []int{}
                    var v21 float32 = 0.1
                    var v22 uint = v6
                    _ = fi19
                    _ = v20
                    _ = v21
                    _ = v22
                }
                var v23 int8 = 42
                _ = v23
            }
            _ = len(v2)
            _ = v17
        }
    }
    v6 += v6
    v24 := []int{v11, (v11 & (v11 >> ((v6) & 63))), (-(v11)), v11}
    var v25 float32 = (float32((1.5 + -2.25)) + 0.1)
    switch 127 {
    case -176:
        v26 := GenS3{F0: 671, F1: 9223372036854775807, F2: 0}
        _ = v26
    default:
        v27 := GenS4{F0: 0, F1: 100, F2: 127}
        _ = v27
    }
    fmt.Println(v1, v1, v1)
    var v28 float64 = 1.5
    fmt.Println(v25)
    if len(v24) > 0 {
        v24[0] = v11
    }
    if false {
        switch v11 {
        case 42:
            var v29 uint = (v6 << (((v6 & v6)) & 63))
            _ = v29
        case 4294967295:
            var v30 uint8 = 1
            _ = v30
        default:
            delete(v3, "hello")
        }
    }
    for fi31 := 0; fi31 < 8; fi31++ {
        if true {
            v6 += 65
            if (42 != -9223372036854775808) {
                for fi32 := 0; fi32 < 10; fi32++ {
                    v11 = fi32
                    _ = fi32
                }
                if (v1 <= 822) {
                    v25 = ((v25 + (((v25 / float32(v28)) * float32(v28)) * 0.0)) - v25)
                }
                switch "\n\t" {
                case "a":
                    var v33 uint32 = uint32(v6)
                    _ = v33
                }
            }
        } else {
            for fi34 := 0; fi34 < 3; fi34++ {
                var v35 uint16 = 0
                _ = len(v3)
                v6 -= 2
                _ = fi34
                _ = v35
            }
            v1 = v1
        }
        _ = fi31
    }
    switch (v11 * (v11 * (v11 + (v11 << (((v6 << (((v6 | v6)) & 63))) & 63))))) {
    case -9223372036854775808:
        v11 -= (v11 | 32768)
    case 128:
        fmt.Println(v11, v28, v25)
    case 32767:
        if false {
            var v36 float64 = v28
            delete(v4, 843)
            _ = v36
        }
    default:
        var v37 string = "hello"
        _ = v37
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v25:", v25)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v24
    _ = v25
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
