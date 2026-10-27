package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 string; F1 uint32; F2 bool }
func genFn0(p0 bool, p1 int32) uint32 {
    fmt.Println(p1)
    p1 = p1
    for fi5 := 0; fi5 < 7; fi5++ {
        for fi6 := 0; fi6 < 20; fi6++ {
            fmt.Println(fi5, fi5, fi6)
            v7 := GenS1{F0: 360, F1: 1}
            var v8 uint16 = 32767
            _ = fi6
            _ = v7
            _ = v8
        }
        v9 := []int{}
        v10 := GenS2{F0: ((127 < 240) || p0)}
        _ = fi5
        _ = v10
        _ = v9
    }
    v11 := []int{-2147483648, 31, 73}
    switch "a" {
    case "":
        p0 = (-10 <= 1)
    case "hello":
        for fi12 := 0; fi12 < 7; fi12++ {
            var v13 int32 = (p1 ^ (p1 << ((42) & 63)))
            _ = fi12
            _ = v13
        }
    case "foo bar":
        v14 := GenS3{F0: ((("\n\t" + "") + "a") + "\n\t"), F1: 31, F2: (1 > 42)}
        _ = v14
    }
    _ = p0
    _ = p1
    _ = v11
    return 302
}

func main() {
    var v1 int16 = -1
    var v2 string = "a"
    v3 := make(map[uint32]int)
    v4 := []int{}
    fmt.Println(v1, v1)
    var v15 int16 = (v1 - v1)
    v16 := GenS2{F0: (!false)}
    var v17 int8 = -104
    v15 -= (v1 * (v15 * (v1 | v1)))
    _ = len(v4)
    var v18 uint64 = 2
    if false {
        fmt.Println(v18)
        v1 -= (v15 * v15)
        v4 = append(v4, 9223372036854775807)
    } else {
        v17 = (v17 / ((v17) | 1))
    }
    switch 1 {
    case 32767:
        v18 = v18
    case -435:
        delete(v3, 42)
    }
    switch 32767 {
    case 829:
        v3[42] = 127
    case -1:
        if (-1 < -1) {
            v18 = v18
        }
    default:
        v2 = v2
    }
    v17 += (v17 * 100)
    v4 = append(v4, 127)
    v15 = (v1 | (v15 << ((4294967295) & 63)))
    if false {
        switch v2 {
        case "a":
            switch -9223372036854775808 {
            case -73:
                var v19 float32 = 0.1
                _ = v19
            case -451:
                fmt.Println(v1, v18, v2)
            }
        case "\n\t":
            v20 := GenS1{F0: 1, F1: 10}
            _ = v20
        default:
            var v21 int32 = 42
            _ = v21
        }
    } else {
        if (!true) {
            for fi22 := 0; fi22 < 4; fi22++ {
                if true {
                    v1 -= (v1 - (^(v1)))
                    v3[415] = (fi22 / ((9223372036854775807) | 1))
                    if (!false) {
                        v2 = v2
                        switch ("\n\t" + (v2 + v2)) {
                        case "hello":
                            var v23 int32 = -101
                            _ = v23
                        case "a":
                            v1 -= (-(v15))
                        }
                    } else {
                        switch (fi22 | (^(fi22))) {
                        case 2147483647:
                            var v24 float32 = 3.14159
                            _ = v24
                        }
                        var v25 int16 = (v15 / ((v1) | 1))
                        _ = v25
                    }
                }
                v1 = 782
                var v26 int16 = 42
                _ = fi22
                _ = v26
            }
            v3[0] = -537
        }
        var v27 uint32 = 228
        fmt.Println(v17)
        _ = v27
    }
    v15 = (-(v15))
    for fi28 := 0; fi28 < 2; fi28++ {
        var v29 uint16 = 65535
        if false {
            v17 = -55
            switch -742 {
            case 42:
                v18 = v18
            case -1:
                _ = len(v3)
            }
            v30 := GenS1{F0: -9223372036854775808, F1: 42}
            _ = v30
        }
        _ = fi28
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
