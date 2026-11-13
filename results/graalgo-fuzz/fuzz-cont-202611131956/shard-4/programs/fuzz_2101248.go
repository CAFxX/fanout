package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint64; F1 uint64 }
type GenS3 struct { F0 int32; F1 uint16 }
func genFn0(p0 int32, p1 int16) int16 {
    p0 = (p0 ^ p0)
    v5 := GenS1{F0: 0}
    if (!(!false)) {
        for fi6 := 0; fi6 < 13; fi6++ {
            v7 := GenS1{F0: 2}
            v8 := []int{}
            _ = fi6
            _ = v7
            _ = v8
        }
    }
    _ = p0
    _ = p1
    _ = v5
    return (^(p1))
}
func genFn1(p0 int32) int8 {
    p0 += p0
    var v9 float32 = ((float32(((-2.25 / (float64(-2.25) + -2.25)) * 3.14159)) - 1e10) / ((3.14159 + 0.1) + 3.14159))
    p0 = p0
    _ = p0
    _ = v9
    return 0
}

func main() {
    var v1 uint32 = 42
    v2 := []string{"\n\t"}
    v3 := make(map[uint32]int)
    var v4 bool = (true || (42 == 0))
    v1 -= 2
    for fi10 := 0; fi10 < 7; fi10++ {
        v11 := GenS2{F0: 4294967295, F1: 0}
        var v12 uint32 = (v1 >> ((283) & 63))
        _ = fi10
        _ = v11
        _ = v12
    }
    if len(v2) > 0 {
        v2[0] = "foo bar"
    }
    switch -50 {
    case 2147483647:
        switch ("\n\t" + "") {
        case "foo bar":
            switch "foo bar" {
            case "a":
                var v13 float32 = (3.14159 + 3.14159)
                _ = v13
            case "foo bar":
                if v4 {
                    for fi14 := 0; fi14 < 5; fi14++ {
                        v4 = v4
                        v3[(v1 << ((688) & 63))] = 79
                        _ = fi14
                    }
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
            case "hello":
                v1 = (v1 ^ v1)
            default:
                v1 += v1
            }
        case "a":
            v3[v1] = -1
        case "hello":
            v4 = v4
        }
    }
    fmt.Println(v1, v4)
    v1 = (v1 & (v1 * v1))
    for fi15 := 0; fi15 < 19; fi15++ {
        v16 := v3[(v1 >> ((0) & 63))]
        var v17 uint = 255
        v1 = (^(v1))
        _ = fi15
        _ = v16
        _ = v17
    }
    switch "" {
    case "a":
        var v18 uint8 = 127
        _ = v18
    case "foo bar":
        v19 := GenS3{F0: 42, F1: 1}
        _ = v19
    }
    v4 = v4
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
