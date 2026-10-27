package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool }
type GenS2 struct { F0 uint8; F1 int }

func main() {
    var v1 int16 = 1
    var v2 string = "hello"
    var v3 uint32 = 2147483647
    v4 := make(map[int]string)
    var v5 int8 = 127
    fmt.Println(v3)
    for fi6 := 0; fi6 < 17; fi6++ {
        var v7 int32 = -1
        _ = fi6
        _ = v7
    }
    v8 := GenS1{F0: -1, F1: (529 != 4294967295)}
    fmt.Println(v2, v5, v3)
    fmt.Println(v3, v2, v3)
    for fi9 := 0; fi9 < 20; fi9++ {
        if ((true || ((1.5 - 3.14159) == (((((0.1 * 1.5) / (0.0 + float32((3.14159 + -2.25)))) - 1e10) + float32(((0.0 * 0.1) * -2.25))) + (1.5 + (0.0 * (1e10 - 1e10)))))) && (false || false)) {
            v3 += v3
            var v10 float32 = 0.0
            switch v2 {
            case "":
                switch (fi9 + (fi9 - 289)) {
                case 32767:
                    var v11 uint8 = 2
                    _ = v11
                case 42:
                    v5 -= v5
                default:
                    fmt.Println(v5, v3, v3)
                }
            case "a":
                for fi12 := 0; fi12 < 7; fi12++ {
                    switch ("" + v2) {
                    case "foo bar":
                        v2 = v2
                    case "hello":
                        v13 := []int{}
                        _ = v13
                    case "\n\t":
                        fmt.Println(v5, v1, fi12)
                        var v14 string = (v2 + v2)
                        _ = v14
                    default:
                        v1 += 881
                    }
                    if false {
                        v1 -= 42
                    }
                    _ = fi12
                }
            }
            _ = v10
        } else {
            var v15 float64 = (-2.25 / -2.25)
            if ((!(511 != 42)) || true) {
                var v16 int8 = (v5 >> ((0) & 63))
                _ = v16
            }
            for fi17 := 0; fi17 < 19; fi17++ {
                var v18 uint16 = 1
                v18 = v18
                v19 := v4[fi17]
                _ = fi17
                _ = v18
                _ = v19
            }
            _ = v15
        }
        _ = fi9
    }
    v20 := GenS1{F0: 2147483647, F1: false}
    for fi21 := 0; fi21 < 3; fi21++ {
        for fi22 := 0; fi22 < 4; fi22++ {
            if true {
                fmt.Println(v1, fi22)
            } else {
                v1 += (v1 * v1)
                fmt.Println(fi21)
            }
            v4[fi22] = v2
            switch v2 {
            case "foo bar":
                v23 := GenS1{F0: -2147483648, F1: (1 >= -202)}
                _ = v23
            case "a":
                var v24 bool = true
                _ = v24
            }
            var v25 uint32 = v3
            switch (fi21 ^ (-(fi22))) {
            case 1:
                var v26 uint32 = v3
                _ = v26
            }
            v27 := GenS2{F0: 127, F1: fi22}
            _ = fi22
            _ = v25
            _ = v27
        }
        for fi28 := 0; fi28 < 3; fi28++ {
            fmt.Println(v1, v3, v2)
            v1 = (v1 * v1)
            var v29 int32 = -1
            v30 := []int{}
            _ = fi28
            _ = v29
            _ = v30
        }
        v2 = v2
        _ = fi21
    }
    v31 := GenS2{F0: 1, F1: 1}
    var v32 bool = (383 < -1)
    v3 = (^(v3))
    if ((v3 & (v3 >> ((991) & 63))) < (^(v3))) {
        var v33 float32 = 0.0
        v5 = -29
        v1 = v1
        _ = v33
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v8
}
