package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8 }
type GenS2 struct { F0 int32; F1 int16; F2 uint8 }
type GenS3 struct { F0 int; F1 int64; F2 int64 }

func main() {
    v1 := make(map[int]string)
    var v2 uint8 = 42
    var v3 uint16 = 2
    var v4 uint16 = (^(v3))
    var v5 uint = 65535
    var v6 uint8 = (v2 ^ (v2 << (((v5 - 0)) & 63)))
    var v7 bool = true
    v6 = 143
    fmt.Println(v3, v4)
    switch -196 {
    case 4294967295:
        switch 2147483647 {
        case 32768:
            v8 := GenS1{F0: "hello", F1: 42}
            _ = v8
        case 4294967295:
            v9 := []int{-1, 0, 2147483647}
            _ = v9
        default:
            for fi10 := 0; fi10 < 8; fi10++ {
                var v11 uint32 = 0
                _ = fi10
                _ = v11
            }
        }
    default:
        delete(v1, 32768)
    }
    var v12 uint16 = (^(v4))
    v13 := GenS1{F0: "foo bar", F1: 42}
    v14 := GenS1{F0: "foo bar", F1: -27}
    v4 = (v3 ^ v4)
    v4 = (v12 | (v3 - 42))
    v3 = v4
    v2 += (v2 / (((v2 - v2)) | 1))
    var v15 float32 = -2.25
    var v16 uint = (^(v5))
    if v7 {
        var v17 uint = (v16 & v16)
        var v18 uint16 = 255
        _ = v17
        _ = v18
    } else {
        if v7 {
            v19 := GenS2{F0: 42, F1: 397, F2: (v6 ^ 250)}
            fmt.Println(v7)
            var v20 uint16 = v12
            _ = v19
            _ = v20
        } else {
            v21 := GenS1{F0: ("\n\t" + "foo bar"), F1: -128}
            if true {
                v7 = false
                for fi22 := 0; fi22 < 11; fi22++ {
                    if v7 {
                        v15 += 0.0
                    } else {
                        v15 += v15
                        var v23 uint8 = (v6 + (v6 % ((v6) | 1)))
                        _ = v23
                    }
                    _ = fi22
                }
                v7 = true
            } else {
                for fi24 := 0; fi24 < 5; fi24++ {
                    v15 = -2.25
                    v25 := GenS1{F0: "hello", F1: -128}
                    _ = fi24
                    _ = v25
                }
                var v26 string = ""
                v12 = 50
                _ = v26
            }
            _ = v21
        }
    }
    var v27 uint16 = v3
    switch -718 {
    case 476:
        if (127 == 100) {
            if v7 {
                v28 := GenS3{F0: -9223372036854775808, F1: 0, F2: 470}
                var v29 uint8 = (v6 & uint8(v5))
                fmt.Println(v5, v12, v3)
                _ = v28
                _ = v29
            }
            var v30 int64 = -727
            _ = v30
        } else {
            switch "" {
            case "hello":
                fmt.Println(v27, v27, v12)
            default:
                v4 += 65535
            }
        }
    case -31:
        for fi31 := 0; fi31 < 8; fi31++ {
            v16 = 1
            var v32 int = (fi31 & (fi31 ^ fi31))
            _ = fi31
            _ = v32
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
