package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 uint32 }
type GenS2 struct { F0 uint; F1 bool; F2 uint16 }
type GenS3 struct { F0 int64; F1 bool }
type GenS4 struct { F0 uint8; F1 string; F2 int }
func genFn0(p0 int16, p1 uint, p2 uint16) int32 {
    var v7 bool = false
    p2 -= 338
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 0
}

func main() {
    v1 := make(map[int]string)
    var v2 string = "hello"
    var v3 int8 = 42
    v4 := make(map[uint32]int)
    var v5 string = ""
    var v6 uint32 = 42
    var v8 int32 = genFn0(-457, 255, 65535)
    var v9 uint16 = 781
    v10 := GenS1{F0: int16(v6), F1: (v9 | 440), F2: (^(v6))}
    v11 := GenS1{F0: 0, F1: v9, F2: (v6 >> ((255) & 63))}
    if (v9 >= 688) {
        for fi12 := 0; fi12 < 20; fi12++ {
            v13 := GenS2{F0: 65535, F1: (1 == 9223372036854775807), F2: v9}
            switch (v5 + ("hello" + (v5 + v5))) {
            case "hello":
                v5 = ("a" + v2)
            }
            _ = fi12
            _ = v13
        }
    }
    for fi14 := 0; fi14 < 16; fi14++ {
        var v15 uint = 1
        for fi16 := 0; fi16 < 1; fi16++ {
            v17 := []int{-9223372036854775808}
            switch v2 {
            case "foo bar":
                switch (v5 + v5) {
                case "a":
                    if true {
                        v18 := GenS3{F0: -9223372036854775808, F1: false}
                        if len(v17) >= 2 {
                            v17 = v17[1:len(v17)]
                        }
                        var v19 uint32 = (^(v6))
                        var v20 int8 = (^(v3))
                        v3 += (v20 + v20)
                        _ = v18
                        _ = v19
                        _ = v20
                    } else {
                        fmt.Println(v5)
                    }
                case "\n\t":
                    if false {
                        var v21 float64 = 0.1
                        fmt.Println(v8, fi16)
                        var v22 int16 = 1000
                        _ = v21
                        _ = v22
                    } else {
                        var v23 int16 = 1000
                        _ = v23
                    }
                case "foo bar":
                    v24 := GenS3{F0: -1, F1: true}
                    _ = v24
                default:
                    v4[(^(v6))] = (fi14 >> (((v15 | v15)) & 63))
                }
            case "hello":
                fmt.Println(v15, v5)
            default:
                var v25 float32 = float32((0.1 / (float64(float32(3.14159)) - 0.1)))
                _ = v25
            }
            _ = fi16
            _ = v17
        }
        fmt.Println(fi14, v3)
        _ = fi14
        _ = v15
    }
    fmt.Println(v2)
    fmt.Println(v5)
    v26 := GenS4{F0: 187, F1: "hello", F2: 703}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
