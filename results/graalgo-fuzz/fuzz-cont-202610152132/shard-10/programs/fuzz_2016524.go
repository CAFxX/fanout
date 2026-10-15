package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint }

func main() {
    var v1 uint32 = 2147483647
    v2 := make(map[int]int)
    var v3 uint8 = 2
    v4 := make(map[string]string)
    var v5 int8 = 1
    var v6 float64 = 0.1
    var v7 float32 = (1.5 - 0.0)
    var v8 uint16 = 32767
    fmt.Println(v8)
    for fi9 := 0; fi9 < 9; fi9++ {
        v1 = v1
        var v10 uint16 = v8
        if (float32(float64(v7)) < (v7 / 1e10)) {
            var v11 bool = true
            if (v1 != 65535) {
                for fi12 := 0; fi12 < 20; fi12++ {
                    if (v7 >= 1.5) {
                        v13 := GenS1{F0: v11}
                        if ((!false) && v11) {
                            v8 += uint16(v1)
                            v7 -= float32(((v6 + (v6 / v6)) - v6))
                            v14 := GenS2{F0: fi12}
                            _ = v14
                        }
                        v3 = v3
                        _ = v13
                    } else {
                        var v15 uint8 = (v3 - 255)
                        var v16 int16 = int16(v1)
                        switch "" {
                        case "\n\t":
                            _ = len(v2)
                        }
                        _ = v15
                        _ = v16
                    }
                    fmt.Println(v7)
                    _ = fi12
                }
                for fi17 := 0; fi17 < 17; fi17++ {
                    if false {
                        fmt.Println(fi9)
                        v18 := GenS3{F0: 42}
                        v19 := GenS4{F0: uint(v10)}
                        var v20 int8 = -1
                        var v21 bool = (-32768 < 42)
                        _ = v18
                        _ = v19
                        _ = v20
                        _ = v21
                    }
                    v22 := GenS2{F0: (fi17 * fi17)}
                    if v11 {
                        v10 = (v8 << ((2) & 63))
                        switch ("foo bar" + "a") {
                        case "hello":
                            var v23 int16 = 1000
                            _ = v23
                        case "\n\t":
                            v5 -= (v5 / ((v5) | 1))
                        }
                        v7 = float32((1e10 / (v6 + 0.0)))
                    }
                    v24 := []int{(fi9 / (((fi17 & 128)) | 1)), -285, -154}
                    switch (("foo bar" + "") + "a") {
                    case "\n\t":
                        v3 += v3
                    default:
                        var v25 uint32 = (v1 | 65535)
                        _ = v25
                    }
                    _ = fi17
                    _ = v22
                    _ = v24
                }
                var v26 int32 = -1
                _ = v26
            }
            _ = v11
        }
        _ = fi9
        _ = v10
    }
    var v27 int64 = 9223372036854775807
    var v28 int8 = 42
    var v29 float64 = float64(v7)
    v5 -= (v5 * v5)
    for fi30 := 0; fi30 < 8; fi30++ {
        switch 0 {
        case 0:
            var v31 uint = 0
            _ = v31
        case 1:
            v8 = (v8 << ((4294967295) & 63))
        case 42:
            v5 = (v28 ^ (v5 / (((^(v5))) | 1)))
        }
        _ = fi30
    }
    v32 := GenS3{F0: 224}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
