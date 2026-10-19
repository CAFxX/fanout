package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int16; F1 uint16 }

func main() {
    var v1 uint = 4294967295
    var v2 bool = (311 >= 255)
    v3 := make(map[uint32]int)
    var v4 int32 = 2147483647
    fmt.Println(v2, v2, v2)
    v1 += (v1 << ((v1) & 63))
    fmt.Println(v2, v1)
    var v5 uint = (v1 ^ v1)
    fmt.Println(v4, v1)
    var v6 uint64 = 42
    fmt.Println(v4, v1, v1)
    switch "a" {
    case "":
        v1 -= 4294967295
    }
    v3[859] = -9223372036854775808
    if ((v2 || (v2 && v2)) && true) {
        v7 := GenS1{F0: 128}
        v8 := GenS1{F0: 255}
        for fi9 := 0; fi9 < 3; fi9++ {
            if (uint64(v5) <= (^(v6))) {
                for fi10 := 0; fi10 < 9; fi10++ {
                    var v11 uint16 = 2
                    _ = fi10
                    _ = v11
                }
                fmt.Println(v6, v1, v2)
            } else {
                if (!false) {
                    v12 := GenS2{F0: -41, F1: 0}
                    switch "\n\t" {
                    case "hello":
                        v13 := GenS1{F0: 127}
                        _ = v13
                    case "\n\t":
                        var v14 float32 = ((0.1 * ((0.0 * (1e10 * 0.1)) * 1e10)) * 1e10)
                        _ = v14
                    case "a":
                        fmt.Println(v1, v2, v4)
                    }
                    _ = v12
                } else {
                    var v15 int64 = int64(v6)
                    fmt.Println(v2, v2, v1)
                    _ = v15
                }
            }
            _ = fi9
        }
        _ = v7
        _ = v8
    } else {
        var v16 float32 = (3.14159 + 1.5)
        for fi17 := 0; fi17 < 6; fi17++ {
            v18 := GenS2{F0: 726, F1: 168}
            switch (fi17 / (((fi17 & int(v6))) | 1)) {
            case 32768:
                if true {
                    v4 = 1
                    var v19 bool = ((v1 % (((v1 << ((v1) & 63))) | 1)) < v1)
                    var v20 uint64 = (^(v6))
                    _ = v19
                    _ = v20
                } else {
                    var v21 uint64 = (v6 + (v6 - 0))
                    var v22 string = "a"
                    switch ("" + v22) {
                    case "\n\t":
                        v22 = v22
                    case "":
                        switch fi17 {
                        case -907:
                            v21 += v6
                        }
                    }
                    _ = v21
                    _ = v22
                }
            case -9223372036854775808:
                for fi23 := 0; fi23 < 4; fi23++ {
                    if (!(!(127 >= 24))) {
                        v5 = v1
                    }
                    _ = fi23
                }
            case 1:
                if (v2 || true) {
                    var v24 int32 = 1
                    _ = v24
                }
            default:
                switch ("hello" + "\n\t") {
                case "a":
                    var v25 int64 = 0
                    _ = v25
                }
            }
            _ = fi17
            _ = v18
        }
        _ = v16
    }
    fmt.Println(v6, v4)
    for fi26 := 0; fi26 < 16; fi26++ {
        var v27 uint = (v1 - v1)
        _ = fi26
        _ = v27
    }
    var v28 int32 = (-(v4))
    if false {
        switch "hello" {
        case "":
            if v2 {
                var v29 int16 = 1000
                _ = v29
            }
        default:
            for fi30 := 0; fi30 < 15; fi30++ {
                v31 := []int{32767, (fi30 / ((128) | 1))}
                for fi32 := 0; fi32 < 3; fi32++ {
                    v4 += (v4 >> (((v1 / ((4294967295) | 1))) & 63))
                    v1 = 255
                    v3[2] = fi30
                    _ = fi32
                }
                v33 := GenS1{F0: 223}
                _ = fi30
                _ = v31
                _ = v33
            }
        }
        var v34 int64 = 9223372036854775807
        v2 = v2
        _ = v34
    } else {
        v5 = (^(v1))
        v35 := []int{32768, -9223372036854775808}
        v28 += v28
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
