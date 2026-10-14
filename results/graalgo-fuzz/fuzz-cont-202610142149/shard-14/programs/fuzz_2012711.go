package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool }
type GenS2 struct { F0 bool; F1 uint32 }
type GenS3 struct { F0 int8; F1 uint; F2 int8 }
func genFn0(p0 int, p1 uint) uint {
    v7 := []int{p0, p0}
    fmt.Println(p1, p0)
    fmt.Println(p0, p1)
    _ = p0
    _ = p1
    _ = v7
    return (p1 & (p1 | 4294967295))
}

func main() {
    v1 := make(map[string]int)
    var v2 string = ("hello" + "")
    var v3 int64 = -9223372036854775808
    v4 := make(map[string]bool)
    var v5 float32 = 0.0
    var v6 int32 = 2147483647
    if ((((3.14159 - -2.25) + 0.1) + (((-2.25 + 1e10) + (1.5 * float64((v5 + (float32(((1e10 / float64((v5 * v5))) - 0.0)) * 1e10))))) / -2.25)) != ((3.14159 - 0.0) / 1.0)) {
        v8 := GenS1{F0: (true && (!(32767 < 32767))), F1: (100 <= -76)}
        v9 := GenS2{F0: ((v2 == (v2 + v2)) && (!(true || (-76 >= 127)))), F1: 944}
        _ = v8
        _ = v9
    } else {
        var v10 bool = (true || false)
        _ = v10
    }
    if (42 == 65) {
        switch 9223372036854775807 {
        case 2147483647:
            v11 := GenS3{F0: 1, F1: 255, F2: 0}
            _ = v11
        case -548:
            v3 -= (v3 % (((v3 | (-(v3)))) | 1))
        default:
            if true {
                v12 := []int{-2147483648}
                _ = v12
            } else {
                v5 = v5
                v3 = (v3 ^ v3)
            }
        }
        var v13 float64 = ((0.0 - (0.1 - 1.5)) - 1e10)
        switch "a" {
        case "foo bar":
            var v14 int32 = 1
            _ = v14
        default:
            v5 += 1.5
        }
        _ = v13
    }
    switch -2147483648 {
    case -484:
        fmt.Println(v6)
    case -2147483648:
        v2 = "foo bar"
    case 211:
        v2 = "hello"
    default:
        var v15 string = v2
        _ = v15
    }
    if false {
        for fi16 := 0; fi16 < 5; fi16++ {
            v17 := GenS3{F0: 100, F1: 42, F2: -128}
            var v18 int32 = v6
            switch (^(fi16)) {
            case 1:
                var v19 bool = (4294967295 < 2147483648)
                _ = v19
            case -285:
                var v20 uint64 = 540
                _ = v20
            }
            _ = fi16
            _ = v17
            _ = v18
        }
        var v21 uint64 = 291
        var v22 uint8 = 0
        _ = v21
        _ = v22
    } else {
        var v23 int32 = (v6 ^ (v6 ^ v6))
        _ = v23
    }
    var v24 float64 = (0.1 + 3.14159)
    v6 -= (v6 ^ (v6 * (v6 + 652)))
    fmt.Println(v5, v6)
    for fi25 := 0; fi25 < 9; fi25++ {
        var v26 uint64 = 42
        var v27 float64 = (v24 * 0.0)
        _ = fi25
        _ = v26
        _ = v27
    }
    var v28 bool = (v6 == v6)
    for fi29 := 0; fi29 < 20; fi29++ {
        v24 = 0.0
        var v30 int8 = 90
        var v31 float64 = 1.5
        _ = fi29
        _ = v30
        _ = v31
    }
    v32 := []int{2147483647, -825}
    v3 = int64(v6)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v28
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
