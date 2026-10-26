package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool; F2 uint64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int16; F1 int; F2 uint8 }
type GenS5 struct { F0 uint64; F1 int; F2 uint32 }

func main() {
    var v1 int64 = 0
    var v2 float64 = (-2.25 / 1e10)
    var v3 int64 = v1
    var v4 int64 = v3
    switch ("foo bar" + "\n\t") {
    case "\n\t":
        for fi5 := 0; fi5 < 11; fi5++ {
            v6 := GenS1{F0: 0, F1: (!true), F2: 4294967295}
            if true {
                v4 += (v4 / ((9223372036854775807) | 1))
            } else {
                switch "hello" {
                case "\n\t":
                    var v7 int = (fi5 ^ 9223372036854775807)
                    _ = v7
                case "hello":
                    var v8 int64 = (v1 + v3)
                    _ = v8
                }
            }
            v1 = v4
            _ = fi5
            _ = v6
        }
    case "hello":
        var v9 int16 = -32768
        _ = v9
    }
    for fi10 := 0; fi10 < 18; fi10++ {
        v11 := GenS2{F0: -1}
        fmt.Println(v4)
        v12 := []int{}
        _ = fi10
        _ = v11
        _ = v12
    }
    if (1 > 2147483647) {
        v13 := []int{264, 127}
        _ = v13
    } else {
        var v14 string = "hello"
        _ = v14
    }
    v1 += v3
    v1 = 0
    switch 489 {
    case -2147483648:
        v15 := GenS1{F0: -1, F1: false, F2: 1}
        _ = v15
    case 32768:
        if (true && true) {
            v16 := GenS3{F0: 127}
            v17 := GenS2{F0: 100}
            _ = v16
            _ = v17
        }
    }
    v18 := GenS3{F0: 1}
    for fi19 := 0; fi19 < 9; fi19++ {
        var v20 uint8 = 255
        v1 = 219
        _ = fi19
        _ = v20
    }
    v21 := GenS4{F0: -32768, F1: -9223372036854775808, F2: 61}
    if false {
        v22 := []int{639, 1, 128}
        var v23 uint = 678
        for fi24 := 0; fi24 < 11; fi24++ {
            for fi25 := 0; fi25 < 12; fi25++ {
                v26 := GenS5{F0: 4294967295, F1: 497, F2: 124}
                _ = fi25
                _ = v26
            }
            v27 := GenS4{F0: -619, F1: 2147483647, F2: 1}
            _ = fi24
            _ = v27
        }
        _ = v22
        _ = v23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v2
    _ = v21
    _ = v3
    _ = v4
}
