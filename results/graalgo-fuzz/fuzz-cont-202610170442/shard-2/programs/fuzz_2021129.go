package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint; F2 int }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint8 }
func genFn0() uint8 {
    switch 128 {
    case -963:
        v5 := []int{-1, -948, 831, -707}
        _ = v5
    case 32768:
        var v6 uint = 18446744073709551615
        _ = v6
    case 1:
        var v7 bool = false
        _ = v7
    }
    v8 := []int{127, 380}
    v9 := GenS1{F0: 42, F1: 2, F2: 625}
    v10 := GenS2{F0: 1}
    for fi11 := 0; fi11 < 20; fi11++ {
        switch "\n\t" {
        case "a":
            var v12 uint16 = 1
            _ = v12
        case "\n\t":
            v8 = append(v8, (fi11 / ((4294967295) | 1)))
        default:
            var v13 uint = 1
            _ = v13
        }
        v14 := GenS1{F0: -1, F1: 1, F2: fi11}
        _ = fi11
        _ = v14
    }
    _ = v10
    _ = v8
    _ = v9
    return 1
}
func genFn1(p0 uint16, p1 string, p2 uint16) uint8 {
    p2 = (p2 | p2)
    for fi15 := 0; fi15 < 3; fi15++ {
        for fi16 := 0; fi16 < 6; fi16++ {
            fmt.Println(p1, fi16)
            _ = fi16
        }
        var v17 string = ((p1 + "hello") + ("foo bar" + "hello"))
        _ = fi15
        _ = v17
    }
    p0 -= p0
    p0 -= 42
    p2 = p0
    _ = p0
    _ = p1
    _ = p2
    return 2
}

func main() {
    var v1 float64 = -2.25
    var v2 uint64 = 1
    var v3 uint16 = 255
    var v4 uint16 = v3
    var v18 int = -573
    if (!false) {
        v18 += 135
        v2 = (v2 >> ((4294967295) & 63))
    }
    v18 += v18
    if (1 <= (v4 - v4)) {
        switch v18 {
        case -197:
            for fi19 := 0; fi19 < 14; fi19++ {
                v2 += 506
                _ = fi19
            }
        case 1:
            var v20 float32 = (((float32((v1 - 0.0)) / 1e10) - 0.0) - 1.5)
            _ = v20
        }
    }
    if true {
        var v21 int = (v18 * (v18 - int(v3)))
        v3 -= 32767
        _ = v21
    }
    switch "hello" {
    case "hello":
        switch "" {
        case "":
            var v22 uint = 1
            _ = v22
        }
    case "":
        fmt.Println(v2, v4, v3)
    case "foo bar":
        switch v18 {
        case 0:
            v23 := GenS3{F0: 2}
            _ = v23
        case -1:
            for fi24 := 0; fi24 < 9; fi24++ {
                v25 := []int{v18, (-(fi24)), fi24}
                var v26 int = v18
                v25 = append(v25, v18)
                _ = fi24
                _ = v25
                _ = v26
            }
        case -9223372036854775808:
            switch -814 {
            case 32768:
                fmt.Println(v18, v3, v3)
            case 950:
                for fi27 := 0; fi27 < 18; fi27++ {
                    switch ("hello" + "foo bar") {
                    case "":
                        var v28 int = fi27
                        _ = v28
                    default:
                        v29 := []int{(v18 | (v18 | 128)), (^(v18))}
                        _ = v29
                    }
                    v18 += (v18 * fi27)
                    _ = fi27
                }
            }
        }
    }
    if (((v1 - (1.5 - -2.25)) - v1) >= ((float64(1e10) + v1) - v1)) {
        v30 := GenS4{F0: 2}
        _ = v30
    } else {
        var v31 string = "a"
        _ = v31
    }
    v2 -= 0
    v32 := []int{1}
    var v33 string = "foo bar"
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
}
