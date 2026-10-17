package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint }

func main() {
    v1 := make(map[uint32]int)
    var v2 uint64 = 477
    var v3 int64 = 42
    var v4 float32 = 0.1
    v5 := GenS1{F0: -665, F1: (3.14159 != (1.5 + -2.25))}
    switch 1 {
    case 4294967295:
        if (42 >= 613) {
            var v6 int32 = -460
            _ = v6
        } else {
            v7 := []int{0, 4294967295}
            v1[2147483648] = 816
            _ = v7
        }
    case -214:
        switch "" {
        case "\n\t":
            for fi8 := 0; fi8 < 16; fi8++ {
                var v9 int64 = (^(v3))
                _ = fi8
                _ = v9
            }
        case "foo bar":
            if false {
                var v10 float64 = float64((float32((0.0 - -2.25)) * v4))
                for fi11 := 0; fi11 < 5; fi11++ {
                    v12 := []int{fi11}
                    var v13 bool = (555 < 32767)
                    _ = fi11
                    _ = v12
                    _ = v13
                }
                v10 = (float64(v4) + v10)
                _ = v10
            }
        }
    case 32767:
        switch -1 {
        case -17:
            for fi14 := 0; fi14 < 4; fi14++ {
                var v15 int16 = 0
                switch fi14 {
                case 42:
                    v15 = v15
                case 4294967295:
                    v16 := GenS1{F0: -1, F1: false}
                    _ = v16
                default:
                    v2 = (v2 & 4294967295)
                }
                var v17 bool = ((-(v15)) >= (v15 / ((v15) | 1)))
                _ = fi14
                _ = v15
                _ = v17
            }
        case -798:
            v18 := GenS2{F0: 42}
            _ = v18
        case 33:
            fmt.Println(v2, v2)
        }
    }
    v19 := GenS2{F0: 2}
    var v20 uint8 = 255
    fmt.Println(v4, v4)
    var v21 int = 127
    delete(v1, 2)
    _ = len(v1)
    v2 = v2
    v22 := GenS1{F0: 0, F1: false}
    v3 -= v3
    v3 -= 9223372036854775807
    switch ("\n\t" + "") {
    case "hello":
        v1[65535] = (v21 | v21)
    case "a":
        v23 := GenS1{F0: 522, F1: (663 <= -32768)}
        _ = v23
    case "\n\t":
        v2 -= v2
    default:
        var v24 uint64 = (v2 | 743)
        _ = v24
    }
    var v25 string = ""
    for fi26 := 0; fi26 < 3; fi26++ {
        var v27 float32 = float32(0.0)
        var v28 string = v25
        _ = fi26
        _ = v27
        _ = v28
    }
    v21 = (^(v21))
    var v29 uint = 972
    v30 := GenS3{F0: 18446744073709551615}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v25:", v25)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v25
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
