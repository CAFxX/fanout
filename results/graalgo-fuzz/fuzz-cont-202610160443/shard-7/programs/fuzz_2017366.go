package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int32; F1 int16 }

func main() {
    var v1 float64 = 0.0
    v2 := make(map[uint32]bool)
    v3 := make(map[int]bool)
    v4 := make(map[int]int)
    var v5 uint32 = 11
    for fi6 := 0; fi6 < 12; fi6++ {
        v7 := []int{604, (fi6 | (fi6 ^ -1))}
        switch "\n\t" {
        case "":
            v8 := GenS1{F0: 1000}
            _ = v8
        default:
            v5 = (v5 - v5)
        }
        var v9 uint32 = v5
        _ = fi6
        _ = v7
        _ = v9
    }
    fmt.Println(v1, v1)
    v10 := GenS2{F0: 814}
    v5 += (v5 * v5)
    v1 = v1
    if true {
        switch "hello" {
        case "foo bar":
            var v11 uint = 65535
            _ = v11
        }
        if ("" > "hello") {
            fmt.Println(v5, v1)
            var v12 float32 = (-2.25 * (0.1 * (-2.25 / 3.14159)))
            if (((((-2.25 + (((-2.25 / (0.0 + 1.5)) / float32(1e10)) + v12)) * -2.25) + v12) / ((((((1e10 + v12) * float32((v1 - 0.1))) - float32(v1)) / v12) - ((((v12 / ((v12 - ((3.14159 - 0.0) / v12)) - ((v12 - ((3.14159 - ((1.5 * v12) - v12)) * (1.5 - v12))) - float32(v1)))) + (0.0 / 3.14159)) - v12) - float32(0.0))) + v12)) >= float32(0.1)) {
                v13 := GenS1{F0: 133}
                if false {
                    if false {
                        v14 := []int{32768, 187}
                        _ = v14
                    } else {
                        switch 829 {
                        case 341:
                            v1 = 0.0
                        case 333:
                            v15 := v3[-9223372036854775808]
                            _ = v15
                        }
                        if true {
                            v1 = float64(v12)
                            fmt.Println(v1)
                        }
                        v1 = 0.1
                    }
                } else {
                    v1 += v1
                    v1 = 1.5
                }
                _ = v13
            }
            _ = v12
        }
    }
    v5 = (v5 % ((v5) | 1))
    v16 := []int{}
    var v17 int8 = -1
    v18 := GenS1{F0: 32767}
    v19 := GenS3{F0: -576, F1: int16(v5)}
    for fi20 := 0; fi20 < 6; fi20++ {
        if len(v16) > 0 {
            v16[0] = fi20
        }
        _ = fi20
    }
    fmt.Println(v1)
    switch "\n\t" {
    case "foo bar":
        v1 = v1
    case "a":
        switch ("hello" + "a") {
        case "":
            if (false && (-474 >= int32(v17))) {
                fmt.Println(v5, v17)
            }
        case "hello":
            var v21 int8 = v17
            _ = v21
        case "\n\t":
            v5 -= (v5 >> ((431) & 63))
        }
    case "\n\t":
        if false {
            v5 = v5
        } else {
            v17 += (^(v17))
        }
    default:
        for fi22 := 0; fi22 < 7; fi22++ {
            var v23 float32 = 3.14159
            switch "\n\t" {
            case "a":
                fmt.Println(fi22, fi22)
            case "hello":
                for fi24 := 0; fi24 < 10; fi24++ {
                    v17 -= (-(v17))
                    var v25 uint64 = 409
                    if len(v16) > 0 {
                        v16[0] = -9223372036854775808
                    }
                    _ = fi24
                    _ = v25
                }
            case "foo bar":
                v26 := GenS3{F0: -664, F1: 1}
                _ = v26
            default:
                for fi27 := 0; fi27 < 6; fi27++ {
                    var v28 string = "a"
                    v4[0] = (-(fi22))
                    _ = fi27
                    _ = v28
                }
            }
            v23 -= v23
            _ = fi22
            _ = v23
        }
    }
    var v29 int32 = -14
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
