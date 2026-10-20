package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 bool; F1 uint64; F2 int64 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint16; F1 int8 }
func genFn0() int {
    var v4 uint16 = 2
    for fi5 := 0; fi5 < 2; fi5++ {
        if true {
            var v6 uint64 = uint64(v4)
            _ = v6
        } else {
            var v7 int32 = 42
            switch (fi5 / ((fi5) | 1)) {
            case 32767:
                for fi8 := 0; fi8 < 11; fi8++ {
                    v9 := GenS1{F0: -293}
                    v10 := GenS2{F0: false, F1: 1, F2: 1}
                    _ = fi8
                    _ = v10
                    _ = v9
                }
            case 4294967295:
                fmt.Println(v4, v7)
            case 2147483647:
                switch "a" {
                case "a":
                    var v11 int32 = v7
                    _ = v11
                }
            }
            _ = v7
        }
        v12 := GenS3{F0: 982}
        _ = fi5
        _ = v12
    }
    if (!true) {
        switch 337 {
        case 9223372036854775807:
            var v13 int8 = 45
            _ = v13
        case 128:
            v14 := GenS1{F0: 0}
            _ = v14
        case 32768:
            for fi15 := 0; fi15 < 9; fi15++ {
                var v16 uint16 = 65535
                _ = fi15
                _ = v16
            }
        }
    } else {
        v17 := GenS1{F0: -32768}
        if false {
            v4 = (v4 + (v4 % ((255) | 1)))
            v18 := []int{492, 9223372036854775807, 32767}
            v4 = v4
            _ = v18
        }
        _ = v17
    }
    fmt.Println(v4)
    _ = v4
    return -9223372036854775808
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint = 255
    v3 := make(map[int]int)
    fmt.Println(v2)
    for fi19 := 0; fi19 < 11; fi19++ {
        var v20 int64 = -1
        v2 += (^(v2))
        v21 := GenS1{F0: 32767}
        _ = fi19
        _ = v20
        _ = v21
    }
    var v22 uint8 = 127
    v23 := []int{0}
    fmt.Println(v2, v2)
    for fi24 := 0; fi24 < 5; fi24++ {
        var v25 float64 = 1e10
        switch genFn0() {
        case 127:
            v26 := GenS4{F0: 2, F1: 0}
            _ = v26
        case 0:
            v22 += v22
        case 1:
            switch (("" + ("hello" + ("foo bar" + "hello"))) + "hello") {
            case "":
                switch "\n\t" {
                case "":
                    v27 := GenS2{F0: true, F1: 416, F2: -1}
                    _ = v27
                case "a":
                    v22 -= 42
                default:
                    v28 := GenS4{F0: 42, F1: 112}
                    _ = v28
                }
            }
        }
        _ = fi24
        _ = v25
    }
    var v29 uint16 = 1
    switch 4294967295 {
    case 807:
        if (2 != 152) {
            var v30 uint8 = (v22 + uint8(v29))
            _ = v30
        }
    default:
        v22 = (v22 / (((v22 / (((v22 >> (((^(v2))) & 63))) | 1))) | 1))
    }
    v22 -= uint8(v29)
    v1[42] = false
    var v31 uint64 = 1
    fmt.Println(v29)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v29
    _ = v3
    _ = v31
}
