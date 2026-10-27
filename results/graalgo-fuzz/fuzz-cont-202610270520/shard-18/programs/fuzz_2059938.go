package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }
type GenS2 struct { F0 uint; F1 int64; F2 uint8 }
type GenS3 struct { F0 int; F1 int8 }
type GenS4 struct { F0 string; F1 uint; F2 uint64 }
func genFn0(p0 uint16, p1 int16) int16 {
    v5 := []int{}
    switch 1 {
    case -2147483648:
        if (127 < 0) {
            p0 += (p0 % ((2) | 1))
            for fi6 := 0; fi6 < 8; fi6++ {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                if false {
                    v7 := GenS1{F0: -33, F1: -1}
                    _ = v7
                } else {
                    p0 = 873
                    var v8 uint64 = 2
                    _ = v8
                }
                _ = fi6
            }
            for fi9 := 0; fi9 < 15; fi9++ {
                p1 += p1
                _ = fi9
            }
        } else {
            fmt.Println(p1)
            if ((p0 + (^(p0))) != (p0 | p0)) {
                p1 -= 0
                var v10 uint64 = 187
                p1 = (p1 & -1)
                _ = v10
            } else {
                v11 := GenS2{F0: 42, F1: 828, F2: 2}
                var v12 uint = 583
                if false {
                    fmt.Println(p0, v12)
                }
                _ = v11
                _ = v12
            }
            var v13 string = ""
            _ = v13
        }
    }
    _ = p0
    _ = p1
    _ = v5
    return p1
}
func genFn1(p0 uint8) int8 {
    var v14 uint8 = (^(p0))
    switch "foo bar" {
    case "a":
        v15 := GenS3{F0: 221, F1: 42}
        _ = v15
    case "foo bar":
        v16 := GenS4{F0: (((("hello" + ("" + "\n\t")) + ("hello" + "\n\t")) + "\n\t") + "\n\t"), F1: 255, F2: 2}
        _ = v16
    }
    _ = p0
    _ = v14
    return 127
}

func main() {
    var v1 int64 = 9223372036854775807
    var v2 uint16 = 2
    v3 := make(map[uint32]uint32)
    v4 := make(map[int]int)
    var v17 int8 = -51
    switch 32767 {
    case -1:
        v1 = v1
    case 32767:
        var v18 int16 = -1
        _ = v18
    }
    if (("hello" + "a") == "hello") {
        fmt.Println(v2, v1)
        for fi19 := 0; fi19 < 5; fi19++ {
            var v20 float32 = (1e10 + (3.14159 - 1.5))
            _ = fi19
            _ = v20
        }
    }
    fmt.Println(v1)
    if (!((-1 <= 548) && true)) {
        var v21 float64 = (1.5 * (0.0 * -2.25))
        _ = v21
    } else {
        v1 = 33
        v1 -= -1
    }
    var v22 uint8 = 2
    for fi23 := 0; fi23 < 10; fi23++ {
        v24 := v3[0]
        var v25 int32 = -59
        fmt.Println(v1)
        _ = fi23
        _ = v24
        _ = v25
    }
    v17 = (v17 % (((-(v17))) | 1))
    v2 = uint16(v22)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v17
    _ = v2
    _ = v22
    _ = v3
    _ = v4
}
