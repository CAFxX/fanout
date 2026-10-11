package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16; F2 bool }
type GenS2 struct { F0 uint16; F1 uint16 }
type GenS3 struct { F0 uint16 }

func main() {
    v1 := []uint32{285}
    v2 := make(map[string]uint32)
    v3 := make(map[string]string)
    var v4 uint64 = 254
    var v5 int16 = 1
    v5 -= (^(v5))
    var v6 uint64 = 4294967295
    for fi7 := 0; fi7 < 10; fi7++ {
        fmt.Println(v4, v4, fi7)
        fmt.Println(v4)
        v8 := GenS1{F0: 2, F1: (v5 - -704), F2: true}
        _ = fi7
        _ = v8
    }
    var v9 uint64 = v4
    if (!false) {
        v10 := GenS2{F0: 347, F1: 1}
        var v11 bool = (42 <= 65535)
        fmt.Println(v6)
        _ = v10
        _ = v11
    }
    if true {
        for fi12 := 0; fi12 < 6; fi12++ {
            for fi13 := 0; fi13 < 4; fi13++ {
                v5 -= v5
                v14 := GenS1{F0: 1, F1: v5, F2: ((((1e10 / 0.1) - -2.25) * 1e10) >= -2.25)}
                _ = fi13
                _ = v14
            }
            switch ("foo bar" + "a") {
            case "\n\t":
                switch "" {
                case "":
                    var v15 float64 = 0.1
                    _ = v15
                }
            case "hello":
                var v16 uint32 = 207
                _ = v16
            case "a":
                _ = len(v1)
            }
            delete(v3, "\n\t")
            _ = fi12
        }
        switch "" {
        case "\n\t":
            v4 = (^(v9))
        }
    } else {
        var v17 float32 = 0.1
        _ = v17
    }
    var v18 string = ("a" + "")
    if true {
        v19 := GenS2{F0: 883, F1: 2}
        _ = v19
    }
    fmt.Println(v5, v5)
    if false {
        v20 := GenS2{F0: 2, F1: 1}
        _ = v20
    }
    var v21 int16 = v5
    var v22 int64 = 42
    fmt.Println(v5, v6)
    if (!true) {
        var v23 uint32 = 2
        var v24 string = v18
        if false {
            v24 = ("a" + v24)
            _ = len(v1)
            switch -810 {
            case 32767:
                var v25 uint8 = 245
                _ = v25
            case 964:
                v21 -= 32767
            case 42:
                fmt.Println(v18, v4, v23)
            }
        } else {
            var v26 float64 = 0.0
            fmt.Println(v21)
            _ = v26
        }
        _ = v23
        _ = v24
    } else {
        v27 := GenS3{F0: 42}
        _ = v27
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
