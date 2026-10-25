package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
type GenS2 struct { F0 uint64; F1 string; F2 int64 }
type GenS3 struct { F0 uint64; F1 int16; F2 uint64 }
type GenS4 struct { F0 int64; F1 int64 }
type GenS5 struct { F0 uint16; F1 uint16 }

func main() {
    var v1 bool = true
    v2 := make(map[uint32]string)
    var v3 uint = 44
    v4 := []uint32{2147483647, 184}
    var v5 string = "\n\t"
    var v6 float64 = (1.5 / 0.1)
    var v7 uint = 42
    if (42 < 1) {
        switch "foo bar" {
        case "hello":
            v7 += (v3 % (((^(v7))) | 1))
        case "foo bar":
            switch -2147483648 {
            case 4294967295:
                v8 := GenS1{F0: -1, F1: (^(v3))}
                _ = v8
            case -317:
                var v9 int16 = -32768
                _ = v9
            default:
                if v1 {
                    for fi10 := 0; fi10 < 13; fi10++ {
                        v5 = v5
                        fmt.Println(v5)
                        _ = fi10
                    }
                    v1 = ((v5 + "a") <= v5)
                    var v11 uint = (v7 + (v7 / ((18446744073709551615) | 1)))
                    _ = v11
                }
            }
        case "":
            var v12 float32 = 0.0
            _ = v12
        }
    } else {
        v3 -= v3
    }
    v6 = (3.14159 - 3.14159)
    v3 -= (^(v3))
    var v13 uint16 = 42
    var v14 uint8 = 2
    v15 := v2[uint32(v14)]
    if v1 {
        v3 = 65535
        for fi16 := 0; fi16 < 16; fi16++ {
            var v17 int8 = -128
            fmt.Println(v3)
            _ = fi16
            _ = v17
        }
        v18 := GenS2{F0: 4294967295, F1: ((v5 + ("a" + "hello")) + v15), F2: 765}
        _ = v18
    }
    var v19 uint16 = 42
    delete(v2, 2147483647)
    var v20 bool = (v5 < v5)
    if ((v14 ^ v14) != v14) {
        for fi21 := 0; fi21 < 15; fi21++ {
            v14 = v14
            switch 32768 {
            case 32768:
                v15 = v5
            }
            v22 := GenS1{F0: 32767, F1: v3}
            _ = fi21
            _ = v22
        }
    } else {
        var v23 uint16 = v13
        _ = v23
    }
    v24 := GenS2{F0: 607, F1: v15, F2: -142}
    var v25 uint32 = 4294967295
    if (42 >= 2) {
        if ((v15 + "\n\t") == v5) {
            for fi26 := 0; fi26 < 14; fi26++ {
                fmt.Println(v25, v25)
                fmt.Println(fi26, v7, fi26)
                _ = fi26
            }
        } else {
            if len(v4) > 0 {
                v4[0] = (v25 / (((v25 % ((v25) | 1))) | 1))
            }
            switch (v15 + ((v15 + ("a" + v15)) + v15)) {
            case "":
                switch 32767 {
                case 4294967295:
                    if true {
                        v27 := GenS3{F0: 1, F1: -32768, F2: 190}
                        fmt.Println(v15)
                        v19 = uint16(v3)
                        _ = v27
                    }
                case 959:
                    v15 = v15
                default:
                    _ = len(v2)
                }
            }
            switch -9223372036854775808 {
            case 4294967295:
                var v28 int = 128
                _ = v28
            default:
                v20 = ((v20 || true) || false)
            }
        }
        v5 = v5
        _ = len(v4)
    }
    for fi29 := 0; fi29 < 14; fi29++ {
        v2[v25] = (v5 + v15)
        if false {
            var v30 uint8 = 2
            _ = v30
        } else {
            v31 := GenS4{F0: -1, F1: 0}
            v32 := GenS5{F0: 32767, F1: (^(v13))}
            v7 = (v3 & v3)
            _ = v31
            _ = v32
        }
        _ = fi29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
