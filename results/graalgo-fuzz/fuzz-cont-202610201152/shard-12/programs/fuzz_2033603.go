package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }
type GenS2 struct { F0 uint16; F1 int64 }
type GenS3 struct { F0 string; F1 uint16; F2 uint }

func main() {
    var v1 uint = 1
    var v2 uint16 = 66
    v3 := make(map[string]bool)
    if false {
        v2 -= (v2 ^ 42)
    } else {
        v2 += v2
    }
    v4 := GenS1{F0: 32767, F1: 32767}
    var v5 uint16 = (v2 & v2)
    var v6 float32 = -2.25
    var v7 float32 = v6
    var v8 uint32 = 868
    delete(v3, "a")
    if true {
        var v9 int8 = 100
        switch -444 {
        case 32767:
            var v10 uint8 = 255
            _ = v10
        case 2147483647:
            var v11 int8 = 42
            _ = v11
        }
        _ = v9
    }
    var v12 string = ("a" + "")
    if false {
        v13 := v3["\n\t"]
        if true {
            for fi14 := 0; fi14 < 13; fi14++ {
                for fi15 := 0; fi15 < 1; fi15++ {
                    v16 := GenS2{F0: (^(v2)), F1: 0}
                    _ = fi15
                    _ = v16
                }
                v17 := []int{-674, fi14, fi14, (fi14 << (((v1 * (v1 << (((v1 - v1)) & 63)))) & 63))}
                v1 = v1
                _ = fi14
                _ = v17
            }
            var v18 float32 = (3.14159 * (0.1 * v7))
            _ = v18
        }
        var v19 int64 = 9223372036854775807
        _ = v13
        _ = v19
    }
    for fi20 := 0; fi20 < 1; fi20++ {
        var v21 bool = false
        v22 := GenS1{F0: 32767, F1: 1000}
        v8 = v8
        _ = fi20
        _ = v21
        _ = v22
    }
    var v23 int32 = 42
    switch -2147483648 {
    case 128:
        v24 := GenS2{F0: uint16(v8), F1: 1}
        _ = v24
    case 2147483647:
        if true {
            switch v12 {
            case "\n\t":
                switch v12 {
                case "":
                    for fi25 := 0; fi25 < 20; fi25++ {
                        fmt.Println(v8, v8, v12)
                        _ = fi25
                    }
                }
            case "foo bar":
                fmt.Println(v2, v8)
            case "":
                var v26 float32 = (((v6 * v7) - v7) * 0.0)
                _ = v26
            }
            switch v12 {
            case "":
                var v27 uint = v1
                _ = v27
            case "\n\t":
                var v28 int = -1
                _ = v28
            }
        }
    case 32767:
        var v29 uint16 = 852
        _ = v29
    default:
        switch 32768 {
        case 1:
            var v30 uint32 = (v8 - v8)
            _ = v30
        case 0:
            v12 = ""
        case 127:
            for fi31 := 0; fi31 < 15; fi31++ {
                for fi32 := 0; fi32 < 4; fi32++ {
                    if true {
                        v33 := GenS3{F0: v12, F1: v5, F2: v1}
                        _ = v33
                    }
                    _ = fi32
                }
                v34 := GenS1{F0: -32768, F1: -818}
                _ = fi31
                _ = v34
            }
        default:
            v2 = uint16(v1)
        }
    }
    v35 := GenS2{F0: (v2 ^ uint16(v1)), F1: 590}
    var v36 string = v12
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v2
    _ = v23
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
