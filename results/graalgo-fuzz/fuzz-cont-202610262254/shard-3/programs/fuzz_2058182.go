package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int; F1 int }
type GenS4 struct { F0 int }

func main() {
    var v1 uint32 = 42
    v2 := make(map[string]bool)
    var v3 float64 = (3.14159 / 1.5)
    fmt.Println(v3)
    v4 := v2["a"]
    if (!false) {
        for fi5 := 0; fi5 < 19; fi5++ {
            for fi6 := 0; fi6 < 4; fi6++ {
                var v7 uint = 175
                switch ("foo bar" + "a") {
                case "a":
                    fmt.Println(fi5, fi6, fi5)
                default:
                    var v8 uint64 = 4294967295
                    fmt.Println(fi5)
                    switch ("foo bar" + "foo bar") {
                    case "hello":
                        v9 := GenS1{F0: (-513 > -464)}
                        _ = v9
                    default:
                        v10 := []int{(fi5 - (fi6 / (((fi5 / ((42) | 1))) | 1))), (fi6 | 2147483647)}
                        _ = v10
                    }
                    _ = v8
                }
                var v11 float32 = 0.1
                if ("a" <= ("hello" + "foo bar")) {
                    fmt.Println(fi5, v3)
                } else {
                    var v12 uint8 = 128
                    _ = v12
                }
                v13 := GenS2{F0: 2}
                var v14 int16 = int16(v1)
                fmt.Println(fi6, fi5)
                var v15 float32 = 1.5
                var v16 int16 = v14
                if true {
                    v11 = (v15 + 3.14159)
                }
                _ = fi6
                _ = v11
                _ = v13
                _ = v14
                _ = v15
                _ = v16
                _ = v7
            }
            var v17 int32 = 842
            if (!true) {
                v18 := GenS1{F0: (0 > 904)}
                v19 := GenS3{F0: 620, F1: fi5}
                _ = v18
                _ = v19
            } else {
                fmt.Println(v4, v1, v17)
            }
            _ = fi5
            _ = v17
        }
        v20 := GenS4{F0: 0}
        for fi21 := 0; fi21 < 10; fi21++ {
            v1 += (v1 ^ 42)
            v4 = v4
            v1 += v1
            _ = fi21
        }
        _ = v20
    } else {
        var v22 int32 = -1
        for fi23 := 0; fi23 < 19; fi23++ {
            var v24 uint32 = v1
            switch 128 {
            case 32767:
                v3 -= v3
            case -401:
                var v25 int8 = 42
                _ = v25
            default:
                switch 127 {
                case 1:
                    v22 -= (v22 << ((42) & 63))
                }
            }
            var v26 int8 = 100
            _ = fi23
            _ = v24
            _ = v26
        }
        var v27 uint8 = 128
        _ = v22
        _ = v27
    }
    var v28 uint64 = 1
    fmt.Println(v3)
    v3 -= ((((1e10 - float64((0.1 + ((0.1 * (3.14159 + 1.5)) + 0.1)))) / v3) * (((v3 - (1e10 - ((v3 / (v3 + 0.1)) + float64((1.5 - (-2.25 * (1.5 - (1e10 + -2.25)))))))) - float64(1e10)) - v3)) / v3)
    v29 := []int{}
    var v30 float64 = -2.25
    v4 = (-1 == -99)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}
