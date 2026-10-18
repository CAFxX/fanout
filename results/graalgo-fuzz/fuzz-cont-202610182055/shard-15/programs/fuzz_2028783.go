package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool }
type GenS2 struct { F0 string; F1 string; F2 string }
type GenS3 struct { F0 int }
func genFn0(p0 int64, p1 uint8, p2 uint16) int16 {
    p0 -= p0
    var v7 uint64 = 528
    fmt.Println(p2, v7, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 822
}
func genFn1(p0 int, p1 int16) int16 {
    if ((genFn0(-891, 255, 1) >= (^(p1))) && true) {
        var v8 int16 = 776
        v9 := []int{}
        _ = v8
        _ = v9
    }
    p0 -= p0
    for fi10 := 0; fi10 < 13; fi10++ {
        v11 := GenS1{F0: p0, F1: (0 == 18446744073709551615)}
        _ = fi10
        _ = v11
    }
    p0 += (p0 << ((42) & 63))
    _ = p0
    _ = p1
    return p1
}

func main() {
    v1 := []uint32{42, 250, 426}
    var v2 uint = 0
    v3 := make(map[uint32]string)
    v4 := make(map[string]string)
    var v5 float64 = ((1.5 / 3.14159) + (1e10 - 1e10))
    var v6 uint64 = 18446744073709551615
    switch 32768 {
    case 9223372036854775807:
        _ = len(v1)
    case 2147483647:
        var v12 uint16 = 764
        _ = v12
    case 128:
        switch 32767 {
        case 32767:
            var v13 int32 = -2147483648
            _ = v13
        case 9223372036854775807:
            var v14 float64 = ((float64(1e10) / v5) + (v5 / v5))
            _ = v14
        }
    default:
        var v15 uint16 = 767
        _ = v15
    }
    _ = len(v4)
    v6 -= v6
    switch 128 {
    case -676:
        fmt.Println(v5)
    case -9223372036854775808:
        v16 := GenS2{F0: ("hello" + "\n\t"), F1: "a", F2: "foo bar"}
        _ = v16
    case 1:
        var v17 int64 = -9223372036854775808
        _ = v17
    }
    v5 += (v5 / v5)
    var v18 int64 = 1
    v2 = v2
    for fi19 := 0; fi19 < 5; fi19++ {
        if ((^(v6)) >= v6) {
            v20 := GenS1{F0: (^(fi19)), F1: false}
            var v21 float64 = v5
            _ = v20
            _ = v21
        } else {
            for fi22 := 0; fi22 < 14; fi22++ {
                v23 := GenS3{F0: fi19}
                _ = fi22
                _ = v23
            }
            for fi24 := 0; fi24 < 20; fi24++ {
                var v25 bool = (2147483648 != 2147483647)
                _ = fi24
                _ = v25
            }
        }
        for fi26 := 0; fi26 < 10; fi26++ {
            v27 := GenS1{F0: fi26, F1: (65535 >= 638)}
            _ = fi26
            _ = v27
        }
        if false {
            for fi28 := 0; fi28 < 19; fi28++ {
                fmt.Println(fi28, v18, v6)
                var v29 string = "foo bar"
                _ = fi28
                _ = v29
            }
        }
        _ = fi19
    }
    var v30 uint = 363
    if true {
        fmt.Println(v18)
    }
    v2 -= 42
    v6 = 0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
