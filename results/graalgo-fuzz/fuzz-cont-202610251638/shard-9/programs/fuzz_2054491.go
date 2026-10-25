package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint32; F1 int8 }
type GenS3 struct { F0 string; F1 bool }
type GenS4 struct { F0 string; F1 int32 }

func main() {
    v1 := make(map[uint32]string)
    var v2 uint8 = 42
    v3 := make(map[uint32]int)
    var v4 int16 = 42
    v5 := make(map[uint32]string)
    if true {
        if false {
            if (255 >= 0) {
                v6 := GenS1{F0: 452}
                _ = v6
            }
            v7 := []int{25, -722, 32767}
            for fi8 := 0; fi8 < 18; fi8++ {
                if true {
                    var v9 uint32 = 2147483647
                    _ = v9
                } else {
                    fmt.Println(fi8, v2, v4)
                    var v10 uint8 = v2
                    _ = v10
                }
                var v11 string = ((((((("a" + "a") + "hello") + "a") + "foo bar") + "foo bar") + "\n\t") + "a")
                if (false || false) {
                    v11 = "foo bar"
                    var v12 uint64 = 0
                    _ = v12
                }
                _ = fi8
                _ = v11
            }
            _ = v7
        } else {
            for fi13 := 0; fi13 < 9; fi13++ {
                v14 := v3[410]
                fmt.Println(fi13, v4)
                _ = fi13
                _ = v14
            }
            if ((42 != -1) && (!false)) {
                v2 += (v2 - 83)
            }
        }
        v2 = v2
        _ = len(v1)
    }
    v15 := GenS2{F0: 2, F1: 1}
    fmt.Println(v4, v4)
    v16 := []int{127, 321, -2147483648}
    var v17 int = -475
    v18 := GenS2{F0: 0, F1: -53}
    if (int8(v17) < -128) {
        _ = len(v16)
        if (float64(((1e10 / 3.14159) * 0.0)) < 0.1) {
            switch "hello" {
            case "\n\t":
                v2 = v2
            case "":
                var v19 int64 = 465
                _ = v19
            case "hello":
                var v20 float32 = (3.14159 / float32(float64((1.5 / 1e10))))
                _ = v20
            }
            _ = len(v3)
            var v21 uint16 = 42
            _ = v21
        } else {
            var v22 int16 = v4
            v23 := GenS3{F0: "foo bar", F1: (!true)}
            _ = v22
            _ = v23
        }
    } else {
        var v24 int16 = (v4 & (v4 | (v4 ^ v4)))
        _ = v24
    }
    var v25 uint = 19
    for fi26 := 0; fi26 < 16; fi26++ {
        _ = len(v5)
        v3[527] = (v17 >> (((v25 & (v25 % ((v25) | 1)))) & 63))
        _ = fi26
    }
    v2 += (v2 + (v2 ^ v2))
    if false {
        var v27 uint16 = 730
        var v28 bool = (false && false)
        for fi29 := 0; fi29 < 9; fi29++ {
            if v28 {
                fmt.Println(v4, v4)
                var v30 int16 = 1
                v31 := GenS4{F0: "foo bar", F1: -840}
                _ = v30
                _ = v31
            }
            var v32 uint8 = 128
            var v33 float32 = 1e10
            _ = fi29
            _ = v32
            _ = v33
        }
        _ = v27
        _ = v28
    }
    if len(v16) >= 2 {
        v16 = v16[1:len(v16)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
