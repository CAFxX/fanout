package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64 }

func main() {
    v1 := []int{}
    var v2 float32 = 0.0
    v3 := make(map[uint32]bool)
    v4 := []string{("foo bar" + ("foo bar" + "hello")), ("\n\t" + "hello"), "\n\t"}
    v5 := make(map[int]bool)
    var v6 int8 = 127
    if len(v1) > 0 {
        v1[0] = 9223372036854775807
    }
    var v7 bool = true
    for fi8 := 0; fi8 < 15; fi8++ {
        v9 := GenS1{F0: -9223372036854775808, F1: 2}
        _ = fi8
        _ = v9
    }
    delete(v5, -9223372036854775808)
    for fi10 := 0; fi10 < 19; fi10++ {
        var v11 uint8 = 255
        _ = fi10
        _ = v11
    }
    for fi12 := 0; fi12 < 10; fi12++ {
        v7 = (32767 != -1)
        var v13 float32 = v2
        v6 += (^(v6))
        _ = fi12
        _ = v13
    }
    var v14 int8 = (v6 << ((910) & 63))
    switch 42 {
    case 32768:
        v15 := GenS1{F0: -1, F1: 1}
        _ = v15
    case -1:
        v14 = (v6 & (v6 + (v6 / ((1) | 1))))
    case -270:
        for fi16 := 0; fi16 < 10; fi16++ {
            for fi17 := 0; fi17 < 5; fi17++ {
                switch (fi17 >> ((2) & 63)) {
                case -766:
                    v18 := GenS1{F0: 0, F1: 858}
                    _ = v18
                case 958:
                    fmt.Println(v6)
                    v7 = (826 < 2)
                }
                _ = fi17
            }
            if (308 != 123) {
                switch ("\n\t" + "") {
                case "":
                    var v19 uint32 = 395
                    _ = v19
                case "a":
                    var v20 float64 = 1e10
                    _ = v20
                default:
                    v14 = (v6 | -1)
                }
            }
            for fi21 := 0; fi21 < 6; fi21++ {
                delete(v5, 9223372036854775807)
                _ = fi21
            }
            _ = fi16
        }
    }
    v2 -= v2
    v14 += 0
    for fi22 := 0; fi22 < 9; fi22++ {
        var v23 float64 = 1.5
        _ = fi22
        _ = v23
    }
    v24 := GenS1{F0: 42, F1: 18446744073709551615}
    for fi25 := 0; fi25 < 1; fi25++ {
        var v26 uint = 2
        v26 += v26
        var v27 bool = (!(v7 && v7))
        _ = fi25
        _ = v26
        _ = v27
    }
    v7 = true
    v6 += v14
    fmt.Println(v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
