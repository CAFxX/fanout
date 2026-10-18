package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 bool; F1 int; F2 uint16 }
type GenS3 struct { F0 int8; F1 int16; F2 uint16 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int8; F1 int16 }
type GenS6 struct { F0 uint16; F1 bool; F2 uint32 }
type GenS7 struct { F0 int8; F1 bool; F2 uint16 }

func main() {
    var v1 string = "hello"
    v2 := make(map[int]bool)
    var v3 uint16 = 255
    v4 := make(map[string]uint32)
    fmt.Println(v3)
    v5 := GenS1{F0: 838}
    switch -581 {
    case 128:
        v6 := GenS2{F0: (!(572 > 995)), F1: 2147483647, F2: (v3 / ((v3) | 1))}
        _ = v6
    case 205:
        var v7 float64 = float64(1e10)
        _ = v7
    case -371:
        v1 = v1
    default:
        var v8 int64 = 1
        _ = v8
    }
    var v9 uint16 = v3
    v10 := GenS3{F0: -128, F1: -705, F2: v9}
    for fi11 := 0; fi11 < 6; fi11++ {
        if (249 >= 1) {
            fmt.Println(fi11)
            v12 := GenS4{F0: 0}
            _ = v12
        }
        _ = fi11
    }
    for fi13 := 0; fi13 < 2; fi13++ {
        v3 -= (v3 % ((v9) | 1))
        _ = fi13
    }
    var v14 float32 = (1e10 + float32(0.1))
    var v15 uint = 201
    _ = len(v4)
    switch (("foo bar" + v1) + (v1 + "hello")) {
    case "":
        v9 += v3
    case "foo bar":
        if (42 > 127) {
            v16 := GenS2{F0: false, F1: 4294967295, F2: v9}
            v17 := GenS5{F0: -1, F1: 423}
            _ = v16
            _ = v17
        } else {
            fmt.Println(v15)
            v18 := []int{42, -9223372036854775808}
            _ = v18
        }
    default:
        for fi19 := 0; fi19 < 11; fi19++ {
            v14 = ((0.1 + float32((float64((v14 / (v14 + ((3.14159 + -2.25) * ((0.1 / 1.5) + 1.5))))) - 1e10))) * ((0.0 - v14) / float32((1e10 / (1e10 / float64(((v14 + (3.14159 / (v14 / 1.5))) - v14)))))))
            v14 = v14
            if (fi19 != (fi19 | -2147483648)) {
                v20 := GenS6{F0: v9, F1: true, F2: 730}
                _ = v20
            }
            _ = fi19
        }
    }
    if (159 <= 24) {
        switch v1 {
        case "foo bar":
            v3 = (^(v3))
        case "hello":
            if true {
                v3 = (v3 ^ v3)
            } else {
                if (v14 < -2.25) {
                    v15 = v15
                }
                v21 := GenS7{F0: -82, F1: true, F2: (v3 - (v3 + v3))}
                _ = v21
            }
        default:
            v9 -= 2
        }
        var v22 int16 = 739
        var v23 float32 = (float32((0.0 + 0.0)) / -2.25)
        _ = v22
        _ = v23
    } else {
        fmt.Println(v3, v9, v9)
    }
    v2[-1] = (42 >= 2147483648)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
