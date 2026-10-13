package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint; F2 int16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int8; F1 string; F2 int64 }

func main() {
    var v1 uint32 = 2147483647
    var v2 uint8 = 128
    v3 := make(map[int]bool)
    var v4 float32 = (3.14159 + (3.14159 + (0.0 + ((float32((float64((0.0 + (-2.25 + ((((1e10 * -2.25) + ((float32((-2.25 * 0.0)) + 0.1) - (float32(-2.25) + 3.14159))) / ((0.1 - 3.14159) * (0.0 + (float32((0.1 * 3.14159)) / 1.0)))) / (float32(1.5) + 1e10))))) - 0.0)) - 0.0) / float32((-2.25 - 0.1))))))
    var v5 int16 = 486
    fmt.Println(v4, v4)
    var v6 float64 = (float64(v4) - (float64(v4) - 0.0))
    for fi7 := 0; fi7 < 7; fi7++ {
        for fi8 := 0; fi8 < 1; fi8++ {
            var v9 int64 = -9223372036854775808
            v5 += 1000
            if (!(v1 == 945)) {
                v4 = v4
            }
            if (((v1 * 0) >= (^(v1))) || false) {
                if false {
                    v10 := GenS1{F0: uint16(v2), F1: 255, F2: v5}
                    v11 := []int{int(v2), (^(fi7)), fi8, 9223372036854775807}
                    v3[fi7] = false
                    v5 += int16(v1)
                    _ = v10
                    _ = v11
                }
                var v12 bool = ((fi7 & -700) > fi8)
                switch "\n\t" {
                case "foo bar":
                    delete(v3, fi7)
                case "":
                    v13 := GenS2{F0: true}
                    _ = v13
                }
                _ = v12
            }
            v14 := []int{int(v5), fi7, fi7}
            _ = fi8
            _ = v14
            _ = v9
        }
        _ = fi7
    }
    var v15 uint64 = uint64(v2)
    var v16 uint16 = uint16(v15)
    v4 -= 1e10
    var v17 int = -936
    v1 = v1
    v18 := GenS1{F0: v16, F1: 0, F2: 1000}
    v19 := GenS1{F0: (^(v16)), F1: 42, F2: (-(v5))}
    for fi20 := 0; fi20 < 11; fi20++ {
        if (-1 != -860) {
            for fi21 := 0; fi21 < 3; fi21++ {
                if false {
                    var v22 float32 = v4
                    v15 = (v15 % ((v15) | 1))
                    v16 = v16
                    var v23 uint64 = 0
                    _ = v22
                    _ = v23
                } else {
                    v24 := GenS1{F0: 125, F1: 455, F2: v5}
                    _ = v24
                }
                _ = fi21
            }
            if true {
                var v25 string = ""
                if (!true) {
                    var v26 uint16 = (^(v16))
                    v17 = (v17 - -740)
                    if true {
                        fmt.Println(v15)
                        v27 := GenS3{F0: fi20}
                        _ = v27
                    } else {
                        v28 := []int{(fi20 - fi20), v17, 0}
                        v2 = (v2 >> ((42) & 63))
                        _ = v28
                    }
                    _ = v26
                } else {
                    v16 = v16
                }
                _ = v25
            } else {
                var v29 uint8 = 255
                v15 = (v15 ^ (^(v15)))
                fmt.Println(fi20, v1, v5)
                _ = v29
            }
            v30 := v3[fi20]
            _ = v30
        }
        v31 := []int{fi20, -1}
        _ = fi20
        _ = v31
    }
    v32 := []int{}
    switch 42 {
    case -515:
        var v33 uint16 = 296
        _ = v33
    }
    v34 := GenS4{F0: 100, F1: ("a" + "hello"), F2: -1}
    for fi35 := 0; fi35 < 4; fi35++ {
        v36 := GenS4{F0: 42, F1: "\n\t", F2: -1}
        v2 = (^(v2))
        var v37 int32 = 2147483647
        _ = fi35
        _ = v36
        _ = v37
    }
    var v38 bool = ((true && (!false)) || true)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v32
    _ = v34
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
