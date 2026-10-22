package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int8; F1 int8; F2 int16 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int; F1 int16 }
func genFn0() int8 {
    var v5 int16 = 42
    if false {
        v6 := GenS1{F0: 0}
        _ = v6
    }
    var v7 bool = false
    _ = v5
    _ = v7
    return -67
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[uint32]bool)
    var v3 int64 = 135
    v4 := []uint32{800, 42}
    _ = len(v4)
    if false {
        var v8 uint = 905
        var v9 int64 = v3
        _ = v8
        _ = v9
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v3 -= v3
    switch "foo bar" {
    case "foo bar":
        v3 += -301
    case "":
        switch 32768 {
        case 4294967295:
            v10 := GenS1{F0: v3}
            _ = v10
        }
    }
    v11 := GenS1{F0: (v3 + 42)}
    fmt.Println(v3)
    var v12 float64 = float64(-2.25)
    switch "\n\t" {
    case "\n\t":
        v3 += v3
    default:
        v3 = v3
    }
    fmt.Println(v12)
    if true {
        if true {
            switch "" {
            case "":
                if (!true) {
                    v3 -= v3
                    v13 := GenS2{F0: 61, F1: -1, F2: 1000}
                    if (!true) {
                        v14 := GenS1{F0: v3}
                        var v15 uint16 = 0
                        v16 := GenS1{F0: 42}
                        _ = v14
                        _ = v15
                        _ = v16
                    }
                    _ = v13
                }
            default:
                switch ((("\n\t" + "\n\t") + "\n\t") + "foo bar") {
                case "\n\t":
                    fmt.Println(v12, v3)
                case "a":
                    var v17 int16 = 32767
                    _ = v17
                case "":
                    v12 = (-2.25 + 1.5)
                default:
                    switch ("foo bar" + "\n\t") {
                    case "":
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    }
                }
            }
            var v18 int8 = 88
            v4 = append(v4, 2147483647)
            _ = v18
        } else {
            v19 := GenS3{F0: 255}
            var v20 int8 = genFn0()
            _ = v19
            _ = v20
        }
    } else {
        var v21 int = -9223372036854775808
        _ = v21
    }
    var v22 uint = 42
    var v23 uint8 = 0
    v24 := GenS4{F0: 32767, F1: -32768}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}
