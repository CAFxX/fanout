package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 int }
type GenS2 struct { F0 uint32; F1 uint; F2 int8 }
type GenS3 struct { F0 int64; F1 int64; F2 uint64 }
type GenS4 struct { F0 int16; F1 uint }
type GenS5 struct { F0 int8; F1 int8; F2 uint }
type GenS6 struct { F0 uint }
func genFn0(p0 int16, p1 uint8, p2 uint8) string {
    p0 -= p0
    if (1 != 42) {
        for fi6 := 0; fi6 < 3; fi6++ {
            v7 := GenS1{F0: 2, F1: 255, F2: (fi6 * (fi6 | (fi6 & 2147483647)))}
            _ = fi6
            _ = v7
        }
    } else {
        v8 := GenS2{F0: 0, F1: 407, F2: 1}
        switch ("foo bar" + "") {
        case "hello":
            v9 := []int{473, -9223372036854775808, -2147483648}
            _ = v9
        case "":
            for fi10 := 0; fi10 < 10; fi10++ {
                v11 := []int{208, (fi10 ^ 42)}
                _ = fi10
                _ = v11
            }
        default:
            for fi12 := 0; fi12 < 13; fi12++ {
                p1 = 82
                v13 := GenS3{F0: -527, F1: 42, F2: 464}
                _ = fi12
                _ = v13
            }
        }
        for fi14 := 0; fi14 < 8; fi14++ {
            if (false && ((1.5 <= (1.5 - -2.25)) || (((-(p0)) <= 0) && (584 == -1)))) {
                if true {
                    fmt.Println(p1)
                    var v15 int64 = 0
                    _ = v15
                }
                var v16 int = 32767
                _ = v16
            }
            switch "hello" {
            case "foo bar":
                fmt.Println(p2)
            case "\n\t":
                p1 = 0
            }
            for fi17 := 0; fi17 < 3; fi17++ {
                fmt.Println(p2, p1)
                v18 := []int{fi17}
                var v19 bool = (!(true || false))
                _ = fi17
                _ = v18
                _ = v19
            }
            _ = fi14
        }
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    return "foo bar"
}
func genFn1() int16 {
    var v20 bool = (42 <= -35)
    fmt.Println(v20)
    switch 4294967295 {
    case 128:
        v21 := GenS4{F0: -586, F1: 709}
        _ = v21
    case 32767:
        var v22 uint = 485
        _ = v22
    case 894:
        var v23 bool = (100 == 0)
        _ = v23
    }
    _ = v20
    return -641
}

func main() {
    v1 := []uint32{}
    v2 := make(map[uint32]string)
    v3 := make(map[int]int)
    v4 := make(map[string]uint32)
    v5 := []uint32{4294967295, 2147483647, 2147483648}
    v24 := GenS5{F0: 127, F1: 0, F2: 65535}
    var v25 uint = 4294967295
    var v26 bool = (-826 < 569)
    if true {
        for fi27 := 0; fi27 < 6; fi27++ {
            _ = len(v3)
            _ = fi27
        }
    } else {
        var v28 float64 = 1.5
        _ = v28
    }
    fmt.Println(v26)
    v26 = true
    var v29 int = 292
    fmt.Println(v26, v29)
    v29 = (v29 + 32768)
    v3[(^(v29))] = v29
    fmt.Println(v29, v29)
    v29 += (v29 / ((9223372036854775807) | 1))
    _ = len(v3)
    v30 := GenS2{F0: 908, F1: (v25 * 410), F2: -83}
    v31 := v2[2147483648]
    v32 := v3[(^(v29))]
    v33 := GenS6{F0: (v25 / ((4294967295) | 1))}
    v32 = 127
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
}
