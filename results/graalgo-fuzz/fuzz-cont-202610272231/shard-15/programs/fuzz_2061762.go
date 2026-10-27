package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int; F2 int16 }
type GenS2 struct { F0 int }
func genFn0(p0 int, p1 bool, p2 int64) string {
    if true {
        p2 = 466
        v6 := []int{}
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        _ = v6
    }
    v7 := []int{(p0 / (((-(p0))) | 1)), 0}
    var v8 int32 = -1
    v8 = (v8 & v8)
    switch (p0 >> ((1) & 63)) {
    case -2147483648:
        if false {
            var v9 float64 = ((((3.14159 / -2.25) / 0.1) / (float64(1.5) / 0.1)) / 1.0)
            _ = v9
        }
    case 32767:
        switch "hello" {
        case "hello":
            switch 127 {
            case -9223372036854775808:
                var v10 uint32 = 4294967295
                _ = v10
            default:
                var v11 uint8 = 42
                _ = v11
            }
        case "foo bar":
            var v12 int32 = (v8 + 2147483647)
            _ = v12
        default:
            v13 := GenS1{F0: 4294967295, F1: p0, F2: 850}
            _ = v13
        }
    case 9223372036854775807:
        var v14 float32 = -2.25
        _ = v14
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return "foo bar"
}

func main() {
    var v1 uint = 800
    var v2 uint = 42
    v3 := make(map[uint32]int)
    var v4 uint64 = 182
    var v5 uint = 4294967295
    var v15 uint = (v5 + 0)
    v5 = (v5 << ((v1) & 63))
    var v16 string = "\n\t"
    var v17 float32 = (0.0 - (((3.14159 + 1e10) - (float32(float64(float32(1.5))) / (0.1 / (((1e10 / 1.0) - (1e10 - float32((0.1 + 0.1)))) - 0.0)))) * 0.0))
    fmt.Println(v4, v16)
    fmt.Println(v17)
    if false {
        v18 := GenS1{F0: 42, F1: 32768, F2: -564}
        switch 32767 {
        case -2147483648:
            var v19 int8 = -49
            _ = v19
        case -181:
            v2 += (v15 + v15)
        case 32767:
            fmt.Println(v2, v15, v16)
        }
        if false {
            v2 -= 0
            for fi20 := 0; fi20 < 3; fi20++ {
                if true {
                    if (uint16(v1) < 592) {
                        var v21 uint16 = 137
                        var v22 uint8 = 128
                        _ = v21
                        _ = v22
                    }
                } else {
                    v23 := GenS2{F0: fi20}
                    fmt.Println(v16)
                    var v24 uint64 = v4
                    _ = v23
                    _ = v24
                }
                fmt.Println(v4, v4, v2)
                _ = fi20
            }
            switch (v16 + ("a" + "\n\t")) {
            case "a":
                switch (v16 + v16) {
                case "":
                    v4 = 4294967295
                case "foo bar":
                    v1 -= 4294967295
                case "\n\t":
                    v25 := GenS2{F0: 1}
                    _ = v25
                default:
                    var v26 int8 = 1
                    _ = v26
                }
            case "\n\t":
                v27 := []int{-237, 127, 691}
                _ = v27
            case "foo bar":
                v5 -= (v5 ^ (^(v2)))
            }
        }
        _ = v18
    } else {
        for fi28 := 0; fi28 < 8; fi28++ {
            var v29 uint64 = v4
            for fi30 := 0; fi30 < 15; fi30++ {
                var v31 uint8 = 42
                _ = fi30
                _ = v31
            }
            _ = fi28
            _ = v29
        }
        var v32 bool = ((!((false && true) || false)) || true)
        _ = v32
    }
    fmt.Println(v2, v2, v4)
    fmt.Println(v16, v17, v16)
    v2 = (v15 / (((v15 & v5)) | 1))
    var v33 int = 127
    v15 += v15
    v34 := []int{0, (-(v33)), v33}
    v35 := GenS1{F0: 733, F1: v33, F2: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
}
