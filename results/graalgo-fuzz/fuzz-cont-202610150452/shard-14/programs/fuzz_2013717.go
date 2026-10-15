package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8; F2 uint64 }
type GenS2 struct { F0 uint; F1 uint32 }
type GenS3 struct { F0 uint16; F1 int; F2 uint32 }
func genFn0() int64 {
    var v6 float64 = 1e10
    var v7 uint64 = 2
    fmt.Println(v6, v6)
    _ = v6
    _ = v7
    return -9223372036854775808
}

func main() {
    v1 := make(map[int]int)
    var v2 int = -427
    var v3 uint16 = 32767
    var v4 uint32 = 2
    var v5 float64 = 0.1
    fmt.Println(v3, v2, v2)
    switch ("hello" + ("\n\t" + "")) {
    case "a":
        switch v2 {
        case 381:
            if (false && true) {
                switch (v2 / (((v2 % ((v2) | 1))) | 1)) {
                case 380:
                    fmt.Println(v4)
                case 4294967295:
                    v8 := GenS1{F0: (v4 % ((v4) | 1)), F1: 28, F2: 851}
                    _ = v8
                }
                v9 := []int{(^(v2)), v2, v2, (-(v2))}
                _ = v9
            } else {
                var v10 int32 = 1
                var v11 uint = 42
                if (!(false && true)) {
                    v12 := []int{v2, -229, v2, (v2 - v2)}
                    _ = v12
                }
                _ = v10
                _ = v11
            }
        case -1:
            var v13 int16 = 1000
            _ = v13
        default:
            var v14 float64 = (v5 * (((v5 / 1.5) + v5) * -2.25))
            _ = v14
        }
    case "":
        if ("foo bar" <= ("a" + "hello")) {
            v3 -= v3
            v15 := []int{773}
            fmt.Println(v3)
            _ = v15
        } else {
            switch (v2 >> ((4294967295) & 63)) {
            case 42:
                var v16 int16 = -32768
                _ = v16
            default:
                v17 := GenS1{F0: v4, F1: -1, F2: 0}
                _ = v17
            }
            var v18 int8 = -48
            _ = v18
        }
    case "hello":
        var v19 float32 = (1e10 - float32(((v5 + (v5 * 3.14159)) / 1.0)))
        _ = v19
    default:
        var v20 uint32 = 557
        _ = v20
    }
    if (1000 <= 1000) {
        v4 = (v4 << ((18446744073709551615) & 63))
        var v21 string = "hello"
        var v22 uint = 144
        _ = v21
        _ = v22
    } else {
        v23 := GenS2{F0: 28, F1: uint32(v3)}
        _ = v23
    }
    if (0 != 42) {
        var v24 int = 32767
        var v25 int = (-(v2))
        _ = v24
        _ = v25
    }
    if false {
        v2 = -529
        var v26 bool = ((1 <= -32768) || (42 <= 42))
        for fi27 := 0; fi27 < 14; fi27++ {
            if true {
                v26 = false
            } else {
                fmt.Println(v4)
                for fi28 := 0; fi28 < 9; fi28++ {
                    if v26 {
                        var v29 uint8 = 128
                        _ = v29
                    }
                    var v30 int32 = -1
                    _ = fi28
                    _ = v30
                }
                delete(v1, (fi27 ^ (v2 + v2)))
            }
            v31 := GenS3{F0: (v3 >> ((345) & 63)), F1: fi27, F2: (v4 - v4)}
            _ = fi27
            _ = v31
        }
        _ = v26
    }
    _ = len(v1)
    v1[v2] = (-(v2))
    var v32 uint64 = 18446744073709551615
    v32 = (v32 >> ((4294967295) & 63))
    v33 := GenS3{F0: 65535, F1: (v2 / (((v2 >> ((4294967295) & 63))) | 1)), F2: uint32(v3)}
    var v34 int32 = int32(v2)
    v2 += 9223372036854775807
    v34 = v34
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
}
