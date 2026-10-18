package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32 }
type GenS2 struct { F0 int8; F1 uint8; F2 int16 }
type GenS3 struct { F0 string; F1 uint32; F2 int32 }

func main() {
    v1 := []int{32768, 127, 2147483647}
    var v2 int32 = 34
    var v3 float32 = (float32(float64(1.5)) / 1.5)
    v4 := []string{"\n\t"}
    v5 := make(map[uint32]uint32)
    v6 := make(map[int]int)
    v2 += (v2 * (v2 / ((112) | 1)))
    if ("" > "a") {
        if true {
            if (153 <= 42) {
                if (782 >= 32767) {
                    v2 = 989
                } else {
                    v7 := v6[-2147483648]
                    _ = v7
                }
                switch 401 {
                case 128:
                    var v8 int16 = 14
                    _ = v8
                case 218:
                    var v9 bool = (1 == -728)
                    _ = v9
                case 1:
                    fmt.Println(v3, v2)
                default:
                    if false {
                        var v10 string = "foo bar"
                        _ = v10
                    } else {
                        v11 := GenS1{F0: -678, F1: 907}
                        _ = v11
                    }
                }
            } else {
                switch 32767 {
                case 127:
                    fmt.Println(v2, v3)
                default:
                    v3 += (v3 * v3)
                }
            }
            delete(v5, 2147483648)
        } else {
            var v12 float32 = v3
            for fi13 := 0; fi13 < 19; fi13++ {
                var v14 int16 = -32768
                v15 := GenS2{F0: 127, F1: 87, F2: (v14 >> ((42) & 63))}
                v14 += 1
                _ = fi13
                _ = v14
                _ = v15
            }
            _ = v12
        }
    }
    var v16 uint32 = 1
    var v17 uint8 = 255
    var v18 uint8 = (v17 / (((v17 >> ((2) & 63))) | 1))
    fmt.Println(v17, v2, v2)
    v3 += 3.14159
    var v19 float64 = 0.1
    fmt.Println(v17, v3)
    var v20 float32 = v3
    switch "a" {
    case "foo bar":
        v21 := GenS2{F0: 127, F1: (v18 - 255), F2: 1}
        _ = v21
    case "":
        _ = len(v1)
    case "\n\t":
        v19 -= v19
    default:
        if (((1 < 0) && (551 >= 0)) || false) {
            var v22 uint64 = 277
            fmt.Println(v3, v20, v16)
            _ = v22
        } else {
            v23 := GenS2{F0: -2, F1: v18, F2: 1}
            for fi24 := 0; fi24 < 20; fi24++ {
                for fi25 := 0; fi25 < 16; fi25++ {
                    switch (("a" + "hello") + ("" + "foo bar")) {
                    case "\n\t":
                        var v26 uint8 = v17
                        _ = v26
                    case "hello":
                        v20 = 0.0
                    case "":
                        v27 := GenS3{F0: "", F1: v16, F2: (v2 % (((-(v2))) | 1))}
                        _ = v27
                    }
                    var v28 bool = false
                    _ = fi25
                    _ = v28
                }
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = fi24
            }
            _ = v23
        }
    }
    v29 := GenS1{F0: -47, F1: v16}
    v5[(^(v16))] = (v16 << ((235) & 63))
    v30 := GenS1{F0: 40, F1: (v16 << ((255) & 63))}
    v19 = float64(((float32(v19) * ((v20 - v3) / 0.1)) * float32(v19)))
    for fi31 := 0; fi31 < 19; fi31++ {
        var v32 uint8 = v18
        _ = fi31
        _ = v32
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
