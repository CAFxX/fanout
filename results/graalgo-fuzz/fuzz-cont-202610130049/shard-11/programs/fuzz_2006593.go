package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64; F2 int32 }

func main() {
    var v1 int64 = -1
    var v2 int16 = 548
    v3 := make(map[uint32]uint32)
    var v4 uint16 = 65535
    var v5 uint = 255
    var v6 uint16 = v4
    var v7 float32 = (((((1.5 * 3.14159) - 1.5) - 0.1) + 1e10) * (0.0 + (0.0 - 1e10)))
    if true {
        v8 := []int{-246, 1, 127}
        for fi9 := 0; fi9 < 9; fi9++ {
            for fi10 := 0; fi10 < 18; fi10++ {
                v2 += (v2 >> ((v5) & 63))
                fmt.Println(v7)
                switch "foo bar" {
                case "\n\t":
                    var v11 uint32 = 2
                    fmt.Println(v5, v7, v5)
                    var v12 int8 = 100
                    _ = v11
                    _ = v12
                case "":
                    if ((v1 + 9223372036854775807) == (-(v1))) {
                        if false {
                            if ((!false) || (!true)) {
                                v5 += 737
                                switch "foo bar" {
                                case "hello":
                                    v5 = v5
                                case "a":
                                    v6 = v4
                                case "":
                                    var v13 uint64 = 0
                                    _ = v13
                                }
                                v14 := GenS1{F0: 0, F1: v1, F2: -279}
                                _ = v14
                            }
                        }
                    }
                }
                if (42 > 39) {
                    v2 -= 210
                }
                var v15 float64 = 1e10
                _ = fi10
                _ = v15
            }
            v16 := GenS1{F0: 4294967295, F1: v1, F2: -1}
            v4 += 304
            _ = fi9
            _ = v16
        }
        _ = v8
    } else {
        if true {
            fmt.Println(v6)
            v17 := []int{}
            var v18 uint64 = 233
            _ = v17
            _ = v18
        }
    }
    v19 := GenS1{F0: uint32(v5), F1: (v1 + v1), F2: 0}
    var v20 int8 = -34
    if (1 >= 611) {
        v21 := []int{-1}
        _ = v21
    }
    var v22 int = -599
    for fi23 := 0; fi23 < 2; fi23++ {
        if true {
            for fi24 := 0; fi24 < 11; fi24++ {
                v7 = ((v7 - float32(0.1)) + float32((1.5 - float64(((((v7 * ((1.5 / 1.5) - (v7 - v7))) * (v7 * (v7 * (v7 / v7)))) * 3.14159) + float32(float64(((v7 / v7) / (float32((1.5 - 3.14159)) * ((v7 / float32(1e10)) / float32((1e10 * 1.5))))))))))))
                _ = fi24
            }
            v25 := []int{}
            if false {
                v5 = v5
            } else {
                v26 := GenS1{F0: 2147483647, F1: int64(fi23), F2: 668}
                for fi27 := 0; fi27 < 5; fi27++ {
                    var v28 string = "\n\t"
                    v29 := GenS1{F0: 665, F1: (v1 % (((v1 / ((v1) | 1))) | 1)), F2: -123}
                    var v30 uint8 = 1
                    _ = fi27
                    _ = v28
                    _ = v29
                    _ = v30
                }
                var v31 int64 = (v1 + -1)
                _ = v26
                _ = v31
            }
            _ = v25
        }
        v32 := []int{fi23}
        _ = fi23
        _ = v32
    }
    var v33 bool = (true || (326 <= 345))
    switch "a" {
    case "hello":
        v34 := GenS1{F0: 155, F1: 0, F2: -52}
        _ = v34
    case "a":
        fmt.Println(v4, v7, v2)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v22
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
