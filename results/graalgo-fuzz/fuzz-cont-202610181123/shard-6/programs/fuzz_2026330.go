package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8 }
type GenS2 struct { F0 uint32 }

func main() {
    v1 := make(map[int]uint32)
    var v2 int = -650
    v3 := make(map[string]int)
    v4 := GenS1{F0: 760, F1: -16}
    if (2147483648 != 1) {
        v2 -= (v2 >> ((2) & 63))
    } else {
        fmt.Println(v2)
        v2 -= -1
    }
    for fi5 := 0; fi5 < 1; fi5++ {
        v6 := GenS1{F0: 42, F1: 7}
        var v7 int64 = 397
        _ = fi5
        _ = v6
        _ = v7
    }
    v2 -= v2
    for fi8 := 0; fi8 < 11; fi8++ {
        switch fi8 {
        case 2147483647:
            var v9 float64 = 0.1
            _ = v9
        case 4294967295:
            v10 := []int{1, fi8}
            _ = v10
        default:
            var v11 uint = 42
            _ = v11
        }
        for fi12 := 0; fi12 < 20; fi12++ {
            if ((!false) && (-1 != 0)) {
                var v13 int16 = -956
                _ = v13
            }
            v14 := []int{v2, fi8, -156}
            _ = fi12
            _ = v14
        }
        for fi15 := 0; fi15 < 5; fi15++ {
            switch "foo bar" {
            case "foo bar":
                fmt.Println(v2, fi8, fi15)
            default:
                v2 = (fi8 - fi15)
            }
            _ = len(v1)
            var v16 int = fi15
            var v17 uint8 = 167
            v17 -= v17
            _ = fi15
            _ = v16
            _ = v17
        }
        _ = fi8
    }
    switch v2 {
    case 9223372036854775807:
        v2 += (v2 % ((1) | 1))
    case 249:
        switch (v2 - v2) {
        case -9223372036854775808:
            v18 := v3[("a" + "foo bar")]
            _ = v18
        case -829:
            v2 = -753
        }
    }
    var v19 int16 = -1
    v20 := GenS1{F0: 2147483647, F1: -1}
    v21 := GenS1{F0: 42, F1: 100}
    for fi22 := 0; fi22 < 4; fi22++ {
        fmt.Println(v19, v19)
        _ = fi22
    }
    var v23 int8 = -1
    for fi24 := 0; fi24 < 9; fi24++ {
        v25 := []int{(v2 % ((fi24) | 1)), fi24, 128}
        if (229 >= 2) {
            if false {
                v26 := GenS2{F0: 4294967295}
                v23 -= (v23 ^ 124)
                switch (v2 | (v2 % ((1) | 1))) {
                case 2147483647:
                    for fi27 := 0; fi27 < 17; fi27++ {
                        var v28 int8 = 127
                        v29 := v3["foo bar"]
                        _ = fi27
                        _ = v28
                        _ = v29
                    }
                }
                _ = v26
            } else {
                var v30 float32 = (float32(3.14159) - float32(1e10))
                v31 := GenS2{F0: 1}
                _ = v30
                _ = v31
            }
        }
        v23 = v23
        _ = fi24
        _ = v25
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v23
    _ = v3
    _ = v4
}
