package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int; F2 uint8 }
type GenS2 struct { F0 uint32; F1 uint16; F2 int }
type GenS3 struct { F0 bool; F1 uint8; F2 uint8 }

func main() {
    var v1 float64 = (1e10 / float64(float32((float64((-2.25 * float32(0.1))) * (1.5 / 1.0)))))
    var v2 uint32 = 0
    var v3 float64 = 3.14159
    v4 := make(map[int]bool)
    var v5 uint = 460
    var v6 uint32 = 42
    v3 = (float64(0.0) - -2.25)
    fmt.Println(v3, v6)
    switch "hello" {
    case "a":
        v5 += v5
    case "\n\t":
        var v7 bool = (true && (float64((float32(1e10) * (((((((float32(0.0) / 1e10) + 1.5) / 3.14159) / 3.14159) * float32((1.5 * 3.14159))) + 1e10) + float32(float64(0.1))))) <= 0.0))
        _ = v7
    }
    v1 = (v3 * v1)
    if (-44 >= 1) {
        v8 := GenS1{F0: 1, F1: 42, F2: 0}
        _ = v8
    }
    var v9 int32 = -2147483648
    v10 := GenS2{F0: 4294967295, F1: 56, F2: 4294967295}
    switch 936 {
    case 127:
        if (!true) {
            v6 = v2
            v1 += float64((3.14159 * (3.14159 - 1.5)))
        }
    case 1:
        var v11 uint32 = (v2 - (v6 - v6))
        _ = v11
    case 32768:
        var v12 bool = false
        _ = v12
    }
    var v13 float64 = (1.5 - 1.5)
    v13 -= v13
    if true {
        for fi14 := 0; fi14 < 10; fi14++ {
            v15 := GenS2{F0: 4294967295, F1: 1, F2: (fi14 % ((9223372036854775807) | 1))}
            switch ("" + "foo bar") {
            case "foo bar":
                if false {
                    if (!(!((!(true && (false && true))) && true))) {
                        for fi16 := 0; fi16 < 7; fi16++ {
                            if true {
                                var v17 uint32 = 42
                                _ = v17
                            }
                            v18 := []int{(fi16 & (fi14 + (fi16 + fi14))), fi14, (fi16 / ((fi14) | 1))}
                            _ = fi16
                            _ = v18
                        }
                    }
                }
            default:
                if true {
                    v19 := GenS3{F0: ("" == "hello"), F1: 0, F2: 38}
                    _ = v19
                }
            }
            var v20 int8 = 68
            _ = fi14
            _ = v15
            _ = v20
        }
        var v21 string = "\n\t"
        _ = v21
    }
    for fi22 := 0; fi22 < 16; fi22++ {
        var v23 uint32 = (^(v6))
        var v24 int64 = 9223372036854775807
        _ = fi22
        _ = v23
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
