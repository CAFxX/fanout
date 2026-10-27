package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int64; F1 uint16; F2 uint64 }

func main() {
    var v1 int = -576
    var v2 uint32 = 4294967295
    v3 := []uint32{v2}
    var v4 string = ("" + "foo bar")
    if false {
        var v5 uint8 = 1
        if false {
            fmt.Println(v5, v2, v2)
            v5 = v5
        }
        _ = v5
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v6 bool = false
    for fi7 := 0; fi7 < 12; fi7++ {
        var v8 uint64 = 42
        switch fi7 {
        case 664:
            fmt.Println(v8)
        default:
            v9 := GenS1{F0: 42, F1: 1}
            _ = v9
        }
        for fi10 := 0; fi10 < 15; fi10++ {
            var v11 int8 = 27
            switch 42 {
            case -2147483648:
                switch v4 {
                case "a":
                    switch 42 {
                    case 2147483647:
                        v12 := GenS2{F0: (v11 % (((v11 * v11)) | 1))}
                        _ = v12
                    case -954:
                        var v13 float64 = (3.14159 + -2.25)
                        _ = v13
                    case 32767:
                        var v14 uint32 = v2
                        _ = v14
                    }
                }
            }
            _ = fi10
            _ = v11
        }
        _ = fi7
        _ = v8
    }
    if true {
        v15 := GenS3{F0: 9223372036854775807, F1: 65535, F2: 18446744073709551615}
        fmt.Println(v2)
        switch (^(v1)) {
        case 128:
            if true {
                v2 = (v2 * v2)
            }
        case 1:
            var v16 int64 = -1
            _ = v16
        case 127:
            var v17 string = v4
            _ = v17
        }
        _ = v15
    }
    var v18 float64 = (-2.25 - 1.5)
    var v19 float64 = (0.1 + v18)
    fmt.Println(v6)
    var v20 uint64 = 0
    fmt.Println(v2)
    fmt.Println(v18, v18)
    v21 := GenS3{F0: 1, F1: 2, F2: v20}
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v6
}
