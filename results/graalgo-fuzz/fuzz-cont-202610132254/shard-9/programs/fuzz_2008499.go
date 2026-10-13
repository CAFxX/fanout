package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64 }
type GenS2 struct { F0 string; F1 uint; F2 uint16 }
type GenS3 struct { F0 uint8; F1 int8 }
func genFn0() int {
    switch 4294967295 {
    case 431:
        for fi5 := 0; fi5 < 19; fi5++ {
            var v6 int32 = 1
            var v7 float64 = (0.0 - -2.25)
            _ = fi5
            _ = v6
            _ = v7
        }
    case 46:
        switch ("" + (("foo bar" + ("" + "")) + "\n\t")) {
        case "a":
            var v8 int32 = -175
            _ = v8
        case "\n\t":
            if (!false) {
                var v9 int64 = 9223372036854775807
                var v10 string = "hello"
                var v11 int32 = -198
                _ = v10
                _ = v11
                _ = v9
            }
        case "":
            v12 := GenS1{F0: 0, F1: 42}
            _ = v12
        }
    default:
        v13 := []int{-1, -685, 0, 0}
        _ = v13
    }
    for fi14 := 0; fi14 < 15; fi14++ {
        var v15 float32 = (0.0 - 0.1)
        _ = fi14
        _ = v15
    }
    var v16 uint = 42
    var v17 float64 = 1e10
    _ = v16
    _ = v17
    return 127
}
func genFn1(p0 string, p1 int32) int8 {
    switch ("\n\t" + ("hello" + p0)) {
    case "":
        if true {
            p0 = (p0 + p0)
        } else {
            for fi18 := 0; fi18 < 5; fi18++ {
                v19 := GenS2{F0: "", F1: 2, F2: 0}
                fmt.Println(fi18)
                var v20 uint32 = 4294967295
                _ = fi18
                _ = v19
                _ = v20
            }
        }
    case "foo bar":
        for fi21 := 0; fi21 < 18; fi21++ {
            v22 := GenS1{F0: 81, F1: 66}
            p1 -= p1
            v23 := GenS3{F0: 1, F1: -1}
            _ = fi21
            _ = v22
            _ = v23
        }
    }
    if (!(((2 <= 63) || false) && true)) {
        v24 := []int{-9223372036854775808}
        var v25 int = 42
        switch p0 {
        case "":
            v24 = append(v24, (v25 + 42))
        case "a":
            fmt.Println(p0)
        }
        _ = v24
        _ = v25
    }
    p1 += p1
    for fi26 := 0; fi26 < 1; fi26++ {
        for fi27 := 0; fi27 < 5; fi27++ {
            v28 := []int{fi26, (fi26 / (((fi27 - (fi27 ^ 429))) | 1)), 42, (fi26 | (fi26 ^ fi26))}
            fmt.Println(fi27, p1, fi27)
            _ = fi27
            _ = v28
        }
        for fi29 := 0; fi29 < 8; fi29++ {
            var v30 int16 = -620
            _ = fi29
            _ = v30
        }
        _ = fi26
    }
    v31 := []int{}
    _ = p0
    _ = p1
    _ = v31
    return 100
}

func main() {
    var v1 int16 = -1
    v2 := make(map[string]string)
    var v3 uint32 = 252
    v4 := []uint32{(v3 + v3), v3, v3}
    v3 = 0
    v3 = 2147483647
    fmt.Println(v1)
    if true {
        v3 += v3
    }
    v3 += 2147483647
    v1 = (v1 - v1)
    v1 += v1
    v3 -= v3
    for fi32 := 0; fi32 < 16; fi32++ {
        v3 = (v3 >> ((255) & 63))
        if true {
            v4 = append(v4, v3)
        }
        v1 = 157
        _ = fi32
    }
    v3 -= 2147483648
    v3 -= 65535
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
