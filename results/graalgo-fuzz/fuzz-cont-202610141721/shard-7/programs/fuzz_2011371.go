package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64 }
type GenS2 struct { F0 int32; F1 int16; F2 uint }
type GenS3 struct { F0 int16; F1 string }
type GenS4 struct { F0 string; F1 uint64; F2 uint32 }
type GenS5 struct { F0 uint16 }

func main() {
    v1 := make(map[string]int)
    var v2 uint = 18446744073709551615
    v3 := make(map[int]bool)
    var v4 float64 = -2.25
    var v5 uint8 = 128
    var v6 int8 = -128
    v7 := GenS1{F0: -9223372036854775808, F1: -9223372036854775808}
    v8 := GenS1{F0: int64(v6), F1: 840}
    var v9 uint32 = 67
    if false {
        if true {
            var v10 uint = (v2 % ((v2) | 1))
            _ = v10
        } else {
            switch 962 {
            case -872:
                var v11 int32 = -2147483648
                _ = v11
            }
            for fi12 := 0; fi12 < 3; fi12++ {
                v2 = v2
                v13 := GenS2{F0: 2147483647, F1: 1, F2: v2}
                switch "hello" {
                case "a":
                    v6 = 1
                case "hello":
                    v3[int(v9)] = false
                }
                _ = fi12
                _ = v13
            }
        }
        for fi14 := 0; fi14 < 12; fi14++ {
            v15 := GenS2{F0: -51, F1: int16(v5), F2: v2}
            var v16 int64 = 9223372036854775807
            for fi17 := 0; fi17 < 19; fi17++ {
                v18 := GenS1{F0: (^(v16)), F1: 1}
                var v19 string = ""
                v20 := GenS3{F0: 0, F1: ("a" + ("a" + v19))}
                v19 = "foo bar"
                v21 := []int{(fi14 ^ fi17), fi17, fi14}
                _ = fi17
                _ = v18
                _ = v19
                _ = v20
                _ = v21
            }
            _ = fi14
            _ = v15
            _ = v16
        }
    }
    v22 := []int{575, -467, -721}
    var v23 string = "a"
    fmt.Println(v9)
    switch v23 {
    case "a":
        v24 := GenS4{F0: "hello", F1: 42, F2: 2147483647}
        _ = v24
    case "":
        v5 = v5
    default:
        v23 = v23
    }
    for fi25 := 0; fi25 < 3; fi25++ {
        var v26 bool = false
        for fi27 := 0; fi27 < 14; fi27++ {
            var v28 string = v23
            var v29 float32 = -2.25
            _ = fi27
            _ = v28
            _ = v29
        }
        if false {
            v6 = (v6 >> ((2) & 63))
            v4 = float64(0.0)
        } else {
            switch fi25 {
            case 2147483647:
                _ = len(v22)
            case 145:
                v30 := GenS1{F0: -118, F1: 0}
                _ = v30
            default:
                var v31 float64 = ((float64(0.1) / (-2.25 - ((v4 * (((v4 - ((3.14159 * v4) - v4)) + v4) * 1.5)) - (v4 * float64(1.5))))) - float64(3.14159))
                _ = v31
            }
            v22 = append(v22, 1)
        }
        _ = fi25
        _ = v26
    }
    var v32 int32 = 779
    v2 -= v2
    v33 := GenS5{F0: 255}
    var v34 uint16 = 42
    v32 = 2147483647
    if false {
        var v35 int32 = (v32 + v32)
        var v36 uint32 = (v9 << ((v2) & 63))
        _ = v35
        _ = v36
    } else {
        if ((float64(((((1e10 - ((float32(1e10) / 3.14159) - (0.0 - 0.0))) / (((0.1 / 0.1) + (((-2.25 + -2.25) - 0.1) + (float32(1e10) - 0.1))) * 1.5)) * 3.14159) + 1e10)) == v4) || false) {
            if len(v22) >= 2 {
                v22 = v22[1:len(v22)]
            }
            var v37 uint16 = 255
            _ = v37
        } else {
            v5 = 127
            if false {
                v38 := GenS3{F0: 966, F1: ((((v23 + "") + v23) + v23) + "hello")}
                if ((v32 < v32) && ((v2 | (v2 % (((v2 >> ((408) & 63))) | 1))) >= 636)) {
                    v39 := GenS4{F0: v23, F1: 2, F2: v9}
                    _ = v39
                }
                _ = v38
            }
            var v40 string = (("hello" + (v23 + (v23 + "a"))) + v23)
            _ = v40
        }
    }
    for fi41 := 0; fi41 < 11; fi41++ {
        fmt.Println(v23, v34)
        for fi42 := 0; fi42 < 20; fi42++ {
            if len(v22) >= 2 {
                v22 = v22[1:len(v22)]
            }
            if true {
                var v43 uint32 = (v9 / ((v9) | 1))
                var v44 uint8 = v5
                fmt.Println(v32, v43)
                _ = v43
                _ = v44
            }
            var v45 int8 = (^(v6))
            _ = fi42
            _ = v45
        }
        v34 -= 0
        _ = fi41
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
