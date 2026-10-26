package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 uint }
type GenS2 struct { F0 string; F1 bool }
type GenS3 struct { F0 uint64; F1 int32; F2 bool }
func genFn0(p0 int32, p1 uint64, p2 int8) int32 {
    switch 42 {
    case 1:
        var v6 uint8 = 127
        _ = v6
    case 864:
        v7 := GenS1{F0: 255, F1: 42, F2: 1}
        _ = v7
    case 128:
        p1 = (p1 / ((2) | 1))
    default:
        var v8 uint16 = 255
        _ = v8
    }
    if false {
        p2 += p2
    }
    for fi9 := 0; fi9 < 18; fi9++ {
        v10 := GenS2{F0: "hello", F1: ("a" > ((("" + (("hello" + "a") + "\n\t")) + "foo bar") + (("a" + "a") + "")))}
        _ = fi9
        _ = v10
    }
    v11 := []int{-272}
    for fi12 := 0; fi12 < 16; fi12++ {
        var v13 int16 = -515
        if true {
            v13 -= (v13 - v13)
            p1 = (p1 << ((42) & 63))
        } else {
            var v14 float64 = 0.1
            if false {
                if false {
                    for fi15 := 0; fi15 < 13; fi15++ {
                        var v16 int8 = p2
                        _ = fi15
                        _ = v16
                    }
                }
                fmt.Println(v14)
            } else {
                switch (fi12 / (((fi12 & (fi12 << ((65535) & 63)))) | 1)) {
                case 948:
                    if true {
                        var v17 string = "hello"
                        _ = v17
                    } else {
                        _ = len(v11)
                        p0 -= p0
                    }
                case 4294967295:
                    var v18 int64 = 9223372036854775807
                    _ = v18
                case 128:
                    fmt.Println(p1)
                }
                v11 = append(v11, fi12)
            }
            _ = v14
        }
        _ = fi12
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return -1
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 float32 = -2.25
    var v3 float32 = (float32(float64(v2)) / 1e10)
    var v4 uint32 = 381
    var v5 int32 = -18
    var v19 int16 = -1
    var v20 float32 = ((((-2.25 - ((float32(1.5) + v3) - (0.0 + -2.25))) / v2) * v3) * (v3 * 1.5))
    var v21 int64 = -1
    v22 := []int{127, 127, 1}
    for fi23 := 0; fi23 < 14; fi23++ {
        v2 = v3
        if len(v22) > 0 {
            v22[0] = (fi23 ^ 42)
        }
        v24 := GenS2{F0: "", F1: (!false)}
        _ = fi23
        _ = v24
    }
    v21 -= v21
    var v25 uint8 = 1
    v26 := GenS1{F0: v25, F1: (v25 | 1), F2: 0}
    if false {
        v20 = (0.0 + v20)
        v5 = (v5 * (v5 ^ genFn0(-2147483648, 42, 0)))
        if ((v4 ^ v4) >= v4) {
            fmt.Println(v21, v5, v5)
        } else {
            _ = len(v1)
            v27 := GenS3{F0: 18446744073709551615, F1: (-(v5)), F2: false}
            if len(v22) > 0 {
                v22[0] = 32767
            }
            _ = v27
        }
    } else {
        delete(v1, v4)
    }
    if (!false) {
        for fi28 := 0; fi28 < 13; fi28++ {
            _ = len(v1)
            _ = len(v1)
            _ = fi28
        }
        var v29 bool = ((176 > (v25 / (((^(v25))) | 1))) || (true && (v4 <= (v4 + v4))))
        if true {
            var v30 int64 = v21
            switch "a" {
            case "foo bar":
                if len(v22) > 0 {
                    v22[0] = 0
                }
            }
            _ = v30
        }
        _ = v29
    } else {
        fmt.Println(v19, v19, v19)
        v3 = v2
        v5 = (v5 & v5)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
