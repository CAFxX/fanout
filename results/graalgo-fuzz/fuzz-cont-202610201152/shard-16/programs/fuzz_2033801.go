package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint32; F1 uint }
func genFn0(p0 uint8) bool {
    p0 = (p0 << ((255) & 63))
    switch -9223372036854775808 {
    case 32767:
        var v7 bool = false
        _ = v7
    case -527:
        fmt.Println(p0)
    }
    v8 := GenS1{F0: p0}
    _ = p0
    _ = v8
    return ((2 == 42) || false)
}
func genFn1() int64 {
    for fi9 := 0; fi9 < 11; fi9++ {
        v10 := GenS1{F0: 42}
        _ = fi9
        _ = v10
    }
    for fi11 := 0; fi11 < 6; fi11++ {
        for fi12 := 0; fi12 < 20; fi12++ {
            v13 := GenS2{F0: 2, F1: 592}
            var v14 int64 = 9223372036854775807
            switch "\n\t" {
            case "\n\t":
                v15 := []int{fi12, (fi11 << ((1) & 63)), (^(fi12)), (fi11 | fi12)}
                _ = v15
            default:
                var v16 float32 = -2.25
                _ = v16
            }
            var v17 string = ("a" + "foo bar")
            v18 := GenS2{F0: 42, F1: 18446744073709551615}
            _ = fi12
            _ = v13
            _ = v14
            _ = v17
            _ = v18
        }
        _ = fi11
    }
    var v19 uint = 65535
    v19 += v19
    _ = v19
    return -9223372036854775808
}

func main() {
    var v1 int64 = 1
    var v2 string = ("\n\t" + "hello")
    v3 := make(map[uint32]uint32)
    var v4 uint = 968
    var v5 uint8 = 0
    v6 := make(map[int]string)
    v20 := []int{32768}
    for fi21 := 0; fi21 < 16; fi21++ {
        if (!true) {
            var v22 int32 = 42
            _ = len(v20)
            var v23 uint = 528
            _ = v22
            _ = v23
        }
        v4 = (v4 >> ((0) & 63))
        _ = fi21
    }
    _ = len(v20)
    v4 = (^(v4))
    for fi24 := 0; fi24 < 8; fi24++ {
        var v25 int64 = (v1 - v1)
        switch (v2 + "hello") {
        case "foo bar":
            switch (fi24 / (((fi24 | (fi24 >> (((^(v4))) & 63)))) | 1)) {
            case 32767:
                v3[uint32(v5)] = 2147483647
            case 9223372036854775807:
                v4 += (^(v4))
            default:
                var v26 bool = false
                _ = v26
            }
        default:
            var v27 int8 = 0
            _ = v27
        }
        _ = fi24
        _ = v25
    }
    if true {
        v1 += (v1 % (((^(v1))) | 1))
        var v28 uint64 = 4294967295
        _ = v28
    } else {
        if true {
            v5 = (v5 * v5)
        }
    }
    delete(v3, 2)
    if (42 > 1) {
        v29 := GenS2{F0: 65535, F1: 402}
        var v30 uint8 = (v5 >> (((v4 + v4)) & 63))
        if true {
            v30 = 125
            delete(v3, 290)
            for fi31 := 0; fi31 < 12; fi31++ {
                fmt.Println(v5, v5, fi31)
                _ = fi31
            }
        }
        _ = v29
        _ = v30
    }
    var v32 uint64 = 18446744073709551615
    var v33 uint16 = 1
    var v34 float32 = 1.5
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
