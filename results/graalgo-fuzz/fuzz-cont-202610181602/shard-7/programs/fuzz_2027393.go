package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint32; F1 string }
type GenS3 struct { F0 int8; F1 int; F2 bool }
type GenS4 struct { F0 int64; F1 uint }

func main() {
    var v1 uint32 = 42
    var v2 uint32 = (v1 << ((1) & 63))
    var v3 int32 = -548
    var v4 uint16 = 255
    var v5 int8 = 127
    fmt.Println(v2)
    v6 := GenS1{F0: v5}
    v2 = (v2 / (((v2 << ((65535) & 63))) | 1))
    var v7 float64 = ((-2.25 / -2.25) + 3.14159)
    v8 := []int{752}
    v9 := GenS2{F0: 2147483648, F1: "a"}
    for fi10 := 0; fi10 < 1; fi10++ {
        if (1 < 1) {
            v4 += v4
            if len(v8) > 0 {
                v8[0] = (^(fi10))
            }
        } else {
            for fi11 := 0; fi11 < 10; fi11++ {
                var v12 uint32 = v1
                v4 = v4
                v13 := GenS1{F0: (v5 - (v5 / ((v5) | 1)))}
                _ = fi11
                _ = v12
                _ = v13
            }
            _ = len(v8)
        }
        if (42 >= -9223372036854775808) {
            if true {
                v14 := GenS3{F0: v5, F1: (fi10 ^ fi10), F2: false}
                _ = v14
            }
            v7 += v7
        } else {
            var v15 uint = 18446744073709551615
            _ = v15
        }
        _ = fi10
    }
    if (1 != 0) {
        v7 -= (-2.25 - -2.25)
        fmt.Println(v5, v7)
        if ((v3 >> ((2) & 63)) <= v3) {
            for fi16 := 0; fi16 < 2; fi16++ {
                var v17 uint32 = 2147483647
                switch "a" {
                case "hello":
                    if len(v8) > 0 {
                        v8[0] = 1
                    }
                }
                _ = fi16
                _ = v17
            }
            switch "foo bar" {
            case "hello":
                if (!(uint32(v4) < (v1 << ((0) & 63)))) {
                    var v18 uint8 = 42
                    v1 = (^(v1))
                    v2 = v2
                    _ = v18
                }
            }
            for fi19 := 0; fi19 < 20; fi19++ {
                if (4294967295 < 1) {
                    var v20 int16 = -32768
                    _ = v20
                }
                fmt.Println(fi19, v5)
                _ = fi19
            }
        }
    } else {
        if (42 > 1) {
            if (1e10 < 0.0) {
                var v21 float64 = float64(1.5)
                fmt.Println(v4, v7)
                fmt.Println(v7, v7)
                _ = v21
            }
            if true {
                var v22 int32 = (v3 ^ v3)
                var v23 uint = 1
                var v24 uint16 = (v4 >> ((uint(v4)) & 63))
                _ = v22
                _ = v23
                _ = v24
            }
            if (65535 >= 255) {
                if ((v7 / (((-2.25 - (((1e10 / v7) * ((-2.25 * v7) + v7)) * ((v7 * float64(0.1)) / v7))) + 1.5) * float64(((0.1 - (float32((float64((0.0 - (-2.25 + ((((1e10 - 0.0) + -2.25) / 1e10) / float32(float64((float32(((float64(1e10) * (float64(0.1) / 1.5)) * float64((0.0 - (float32(float64(((0.0 + ((0.0 * 3.14159) - (3.14159 / 0.1))) * (1e10 * (0.0 - ((-2.25 - float32(v7)) - 0.1)))))) + (0.0 + (0.1 - ((1.5 + (-2.25 / (1.5 + (0.1 / 0.1)))) + 3.14159)))))))) / -2.25))))))) / v7)) * (0.1 - -2.25))) / (1.5 + 1e10))))) < -2.25) {
                    var v25 int64 = -9223372036854775808
                    if true {
                        v26 := GenS4{F0: v25, F1: 4294967295}
                        _ = v26
                    } else {
                        if ((v25 - v25) != 1) {
                            fmt.Println(v7, v7)
                            var v27 float64 = v7
                            _ = v27
                        }
                        var v28 int8 = -19
                        _ = v28
                    }
                    var v29 uint32 = (v2 >> ((4294967295) & 63))
                    _ = v25
                    _ = v29
                } else {
                    v4 -= v4
                }
            } else {
                switch "hello" {
                case "a":
                    v30 := GenS3{F0: -128, F1: 42, F2: true}
                    _ = v30
                case "\n\t":
                    v8 = append(v8, 32768)
                case "foo bar":
                    var v31 int8 = (v5 / ((v5) | 1))
                    _ = v31
                default:
                    var v32 string = ("\n\t" + ("hello" + "foo bar"))
                    _ = v32
                }
                var v33 string = "foo bar"
                var v34 int16 = 32767
                _ = v33
                _ = v34
            }
        }
        for fi35 := 0; fi35 < 19; fi35++ {
            fmt.Println(v2, v3, v7)
            _ = fi35
        }
    }
    var v36 float32 = (-2.25 + 0.0)
    v37 := GenS3{F0: (v5 / (((v5 % (((-(v5))) | 1))) | 1)), F1: -1, F2: false}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
