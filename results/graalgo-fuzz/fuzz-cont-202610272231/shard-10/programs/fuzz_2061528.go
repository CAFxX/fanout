package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint }
type GenS2 struct { F0 uint64; F1 uint }
type GenS3 struct { F0 int64 }
func genFn0() int8 {
    var v6 uint = 4294967295
    var v7 uint16 = 0
    var v8 int = -76
    if (371 != 365) {
        v7 += (^(v7))
        var v9 float64 = -2.25
        v8 -= (v8 % ((v8) | 1))
        _ = v9
    }
    _ = v6
    _ = v7
    _ = v8
    return -65
}

func main() {
    var v1 uint16 = 42
    var v2 int8 = 51
    v3 := make(map[uint32]int)
    v4 := []int{-1, 780}
    v5 := []string{"foo bar", "a", ("a" + "hello"), ("foo bar" + "")}
    v2 += genFn0()
    v1 -= v1
    v1 = (v1 * 255)
    v1 = v1
    for fi10 := 0; fi10 < 5; fi10++ {
        for fi11 := 0; fi11 < 4; fi11++ {
            v1 = (v1 ^ v1)
            _ = fi11
        }
        switch "\n\t" {
        case "hello":
            v2 = (v2 ^ (v2 % (((v2 & v2)) | 1)))
        case "foo bar":
            var v12 float32 = 1.5
            _ = v12
        case "\n\t":
            v13 := GenS1{F0: v2, F1: 364}
            _ = v13
        }
        v14 := GenS2{F0: 392, F1: 684}
        _ = fi10
        _ = v14
    }
    v2 = genFn0()
    v15 := GenS3{F0: 832}
    var v16 int = 2147483647
    if false {
        if true {
            if true {
                var v17 uint8 = 0
                _ = v17
            }
            fmt.Println(v2, v1, v16)
            v16 -= 251
        } else {
            delete(v3, 4294967295)
        }
    } else {
        if ((v1 & v1) < (v1 | 727)) {
            v3[2147483647] = (v16 * v16)
        }
    }
    if false {
        v2 -= v2
        for fi18 := 0; fi18 < 3; fi18++ {
            v1 += v1
            var v19 float32 = 1e10
            var v20 int16 = 42
            _ = fi18
            _ = v19
            _ = v20
        }
    } else {
        v2 = v2
    }
    v3[4294967295] = 1
    v21 := GenS1{F0: 0, F1: 1}
    var v22 uint32 = 692
    switch "a" {
    case "hello":
        v2 = v2
    case "a":
        var v23 uint32 = uint32(v1)
        _ = v23
    case "foo bar":
        if (!true) {
            v4 = append(v4, v16)
        }
    }
    var v24 int8 = genFn0()
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v21
    _ = v22
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
