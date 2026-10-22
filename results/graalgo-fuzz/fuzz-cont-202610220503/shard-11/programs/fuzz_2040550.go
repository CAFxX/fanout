package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int16; F1 int }
type GenS4 struct { F0 bool; F1 bool }
func genFn0(p0 uint32, p1 int8) int32 {
    switch "foo bar" {
    case "\n\t":
        if (!((p0 - 2147483647) >= (^(p0)))) {
            var v5 int = -946
            var v6 uint8 = 127
            _ = v5
            _ = v6
        } else {
            if (true || (42 >= 1000)) {
                v7 := []int{-2147483648, 4294967295, -9223372036854775808}
                _ = v7
            }
            var v8 bool = (-622 > 644)
            v9 := GenS1{F0: v8}
            _ = v8
            _ = v9
        }
    }
    var v10 float64 = 0.1
    for fi11 := 0; fi11 < 18; fi11++ {
        switch "" {
        case "hello":
            if (false && true) {
                v10 = (0.0 / (float64(0.1) + v10))
                p1 += (p1 & -44)
                var v12 int32 = -24
                _ = v12
            } else {
                v13 := GenS2{F0: 2}
                p0 = (p0 * (p0 + p0))
                _ = v13
            }
        }
        _ = fi11
    }
    _ = p0
    _ = p1
    _ = v10
    return -638
}

func main() {
    var v1 float64 = ((0.1 / 3.14159) + 0.1)
    v2 := []int{}
    var v3 int64 = 42
    v4 := make(map[uint32]bool)
    var v14 float32 = float32(float64(((1.5 + 1e10) * 1e10)))
    fmt.Println(v1)
    if ((float32(float64(0.1)) / (float32((0.0 + v1)) + 0.0)) <= (v14 - float32(((v1 / ((v1 + v1) + v1)) - v1)))) {
        for fi15 := 0; fi15 < 18; fi15++ {
            var v16 float32 = v14
            _ = fi15
            _ = v16
        }
        v1 += float64(v14)
    } else {
        _ = len(v4)
        var v17 int32 = 0
        for fi18 := 0; fi18 < 5; fi18++ {
            v3 -= v3
            switch 48 {
            case -628:
                var v19 bool = true
                _ = v19
            default:
                v3 = 54
            }
            _ = fi18
        }
        _ = v17
    }
    if (!true) {
        switch ("\n\t" + "a") {
        case "":
            for fi20 := 0; fi20 < 10; fi20++ {
                v14 = 3.14159
                _ = fi20
            }
        case "foo bar":
            v1 -= -2.25
        case "\n\t":
            switch "" {
            case "":
                var v21 uint16 = 0
                _ = v21
            case "hello":
                v2 = append(v2, -2147483648)
            case "\n\t":
                v14 = v14
            default:
                fmt.Println(v1, v14)
            }
        }
        var v22 int64 = 42
        _ = v22
    }
    fmt.Println(v3)
    fmt.Println(v3, v1)
    var v23 int = 32767
    for fi24 := 0; fi24 < 18; fi24++ {
        v14 = ((3.14159 + ((v14 + v14) / 0.1)) - 0.1)
        _ = fi24
    }
    if false {
        for fi25 := 0; fi25 < 15; fi25++ {
            v26 := GenS3{F0: -1, F1: fi25}
            v3 = (v3 & -9223372036854775808)
            _ = fi25
            _ = v26
        }
        v27 := GenS2{F0: 2}
        if (652 != 42) {
            var v28 uint32 = 462
            v29 := GenS4{F0: true, F1: (!(!(false && false)))}
            _ = v28
            _ = v29
        }
        _ = v27
    } else {
        v3 = (-(v3))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v14
    _ = v2
    _ = v23
    _ = v3
    _ = v4
}
