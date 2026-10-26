package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint }
type GenS2 struct { F0 int8; F1 uint16; F2 bool }
type GenS3 struct { F0 string; F1 int64; F2 int32 }
type GenS4 struct { F0 uint8; F1 int16 }
type GenS5 struct { F0 int32; F1 int }

func main() {
    var v1 uint64 = 2
    v2 := make(map[int]uint32)
    var v3 float32 = (1.5 + 1.5)
    v4 := GenS1{F0: 65535, F1: 2}
    var v5 string = ""
    v1 -= (^(v1))
    for fi6 := 0; fi6 < 15; fi6++ {
        v1 = (v1 | v1)
        v7 := GenS1{F0: 65535, F1: 0}
        _ = fi6
        _ = v7
    }
    delete(v2, -353)
    fmt.Println(v3)
    switch -1 {
    case 0:
        if false {
            if ((v3 + v3) != v3) {
                fmt.Println(v1)
                v3 = v3
                for fi8 := 0; fi8 < 15; fi8++ {
                    for fi9 := 0; fi9 < 7; fi9++ {
                        v3 = 0.0
                        fmt.Println(v3, v1)
                        _ = fi9
                    }
                    v10 := v2[9223372036854775807]
                    _ = fi8
                    _ = v10
                }
            }
            for fi11 := 0; fi11 < 13; fi11++ {
                _ = len(v2)
                switch (-(fi11)) {
                case -2147483648:
                    for fi12 := 0; fi12 < 9; fi12++ {
                        var v13 int64 = -1
                        _ = fi12
                        _ = v13
                    }
                case -935:
                    v2[fi11] = 606
                default:
                    v14 := []int{(^(fi11)), fi11, fi11, fi11}
                    _ = v14
                }
                if ((float32((3.14159 * 3.14159)) + 0.1) < 0.0) {
                    fmt.Println(v3, fi11, v1)
                    if false {
                        switch v5 {
                        case "hello":
                            var v15 uint32 = 1
                            _ = v15
                        default:
                            var v16 int64 = 382
                            _ = v16
                        }
                        v2[2147483647] = 131
                    } else {
                        v17 := []int{fi11}
                        v3 = (-2.25 + (1e10 + v3))
                        if false {
                            for fi18 := 0; fi18 < 12; fi18++ {
                                v1 -= (v1 % ((v1) | 1))
                                fmt.Println(fi18, fi11)
                                v19 := GenS2{F0: 57, F1: 42, F2: false}
                                _ = fi18
                                _ = v19
                            }
                            fmt.Println(fi11)
                            switch fi11 {
                            case 1:
                                switch (v5 + ((v5 + "hello") + v5)) {
                                case "":
                                    v20 := GenS2{F0: 127, F1: 973, F2: false}
                                    _ = v20
                                }
                            }
                        }
                        _ = v17
                    }
                    v5 = (("a" + v5) + "\n\t")
                } else {
                    fmt.Println(v3)
                    v21 := GenS1{F0: 65535, F1: 65535}
                    switch fi11 {
                    case 42:
                        v1 = (v1 ^ 2)
                    }
                    _ = v21
                }
                _ = fi11
            }
        }
    }
    v22 := GenS1{F0: 1, F1: 42}
    for fi23 := 0; fi23 < 17; fi23++ {
        for fi24 := 0; fi24 < 17; fi24++ {
            var v25 float64 = 1.5
            var v26 uint = 18446744073709551615
            v27 := GenS3{F0: v5, F1: 9223372036854775807, F2: 0}
            _ = fi24
            _ = v25
            _ = v26
            _ = v27
        }
        _ = fi23
    }
    var v28 float32 = v3
    for fi29 := 0; fi29 < 19; fi29++ {
        v30 := []int{fi29, (^(fi29))}
        var v31 bool = false
        _ = fi29
        _ = v30
        _ = v31
    }
    switch 1 {
    case 355:
        v32 := v2[1]
        _ = v32
    case 661:
        for fi33 := 0; fi33 < 12; fi33++ {
            v3 += ((v3 * ((float32(0.1) * v3) + 1e10)) / ((v28 - (v28 * float32(-2.25))) * 1e10))
            _ = fi33
        }
    default:
        for fi34 := 0; fi34 < 7; fi34++ {
            v35 := []int{-695}
            _ = fi34
            _ = v35
        }
    }
    v36 := GenS4{F0: 42, F1: 0}
    var v37 uint16 = 1
    v38 := GenS5{F0: 2147483647, F1: 42}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v28
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v5
}
