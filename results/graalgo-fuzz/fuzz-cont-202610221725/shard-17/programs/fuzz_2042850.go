package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint64; F1 int }
type GenS4 struct { F0 uint16; F1 int }

func main() {
    var v1 float32 = float32((0.0 / float64(((0.1 - 0.0) - 1e10))))
    var v2 int64 = -1
    v3 := make(map[int]bool)
    if (!false) {
        fmt.Println(v2, v2)
        v4 := []int{-966, -276, -747}
        var v5 float32 = 1e10
        _ = v4
        _ = v5
    } else {
        var v6 bool = (false || (1 == 1000))
        v7 := GenS1{F0: 4294967295, F1: 100}
        if (!v6) {
            if v6 {
                fmt.Println(v1, v1, v6)
                v8 := []int{9223372036854775807, 2147483647, 250, 1}
                if len(v8) >= 2 {
                    v8 = v8[1:len(v8)]
                }
                _ = v8
            }
        } else {
            v1 -= v1
        }
        _ = v6
        _ = v7
    }
    for fi9 := 0; fi9 < 2; fi9++ {
        var v10 uint = 780
        if true {
            v1 = v1
            v11 := GenS1{F0: (v10 * (v10 / ((0) | 1))), F1: 100}
            for fi12 := 0; fi12 < 3; fi12++ {
                switch ("hello" + "a") {
                case "hello":
                    if true {
                        v13 := GenS1{F0: 42, F1: -74}
                        _ = v13
                    }
                case "\n\t":
                    if (!((v2 >> ((v10) & 63)) != (v2 * (v2 >> ((v10) & 63))))) {
                        var v14 string = "hello"
                        var v15 uint = 255
                        _ = v14
                        _ = v15
                    }
                case "":
                    var v16 uint64 = 1
                    _ = v16
                default:
                    fmt.Println(fi12, fi9, v1)
                }
                v17 := []int{fi9}
                _ = fi12
                _ = v17
            }
            _ = v11
        }
        _ = fi9
        _ = v10
    }
    var v18 float64 = (((-2.25 - 0.1) / 3.14159) - 0.1)
    v19 := []int{-314, -9223372036854775808, 2147483647, 0}
    var v20 float64 = v18
    v19 = append(v19, 32768)
    if ("foo bar" >= (("foo bar" + ("hello" + "\n\t")) + "a")) {
        _ = len(v3)
        v21 := GenS2{F0: 344}
        var v22 int32 = 911
        _ = v21
        _ = v22
    }
    v1 = float32(v18)
    v20 = (v18 * (0.0 * v20))
    var v23 uint16 = 589
    v24 := GenS3{F0: 2, F1: 1}
    if len(v19) >= 2 {
        v19 = v19[1:len(v19)]
    }
    var v25 string = "hello"
    v20 = v18
    for fi26 := 0; fi26 < 2; fi26++ {
        var v27 float64 = float64((1e10 + v1))
        v28 := GenS1{F0: 0, F1: 1}
        _ = fi26
        _ = v27
        _ = v28
    }
    for fi29 := 0; fi29 < 12; fi29++ {
        if (!false) {
            for fi30 := 0; fi30 < 9; fi30++ {
                switch 1 {
                case -85:
                    switch fi30 {
                    case 513:
                        var v31 float32 = v1
                        _ = v31
                    case 32768:
                        v2 += v2
                    case 569:
                        _ = len(v19)
                    default:
                        v32 := GenS4{F0: (v23 ^ v23), F1: (^(fi29))}
                        _ = v32
                    }
                default:
                    v33 := GenS4{F0: (v23 & (v23 ^ (v23 >> ((255) & 63)))), F1: (fi30 | (fi30 % ((fi30) | 1)))}
                    _ = v33
                }
                _ = fi30
            }
            switch 464 {
            case -29:
                v2 = -9223372036854775808
            }
            switch "\n\t" {
            case "foo bar":
                var v34 uint8 = 127
                _ = v34
            default:
                var v35 int64 = (v2 >> ((535) & 63))
                _ = v35
            }
        } else {
            if len(v19) >= 2 {
                v19 = v19[1:len(v19)]
            }
            for fi36 := 0; fi36 < 14; fi36++ {
                v18 = ((((3.14159 - v20) * (1e10 + v18)) * 0.1) / (v20 * (float64((float32(0.0) + (v1 + ((1e10 - v1) * (float32(3.14159) / 1.0))))) - (v20 - ((((v20 / float64(-2.25)) + 0.1) - (1e10 / (v18 - v18))) / -2.25)))))
                _ = fi36
            }
            var v37 uint = 65535
            _ = v37
        }
        _ = fi29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v24
    _ = v25
    _ = v3
}
