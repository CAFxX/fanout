package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int32; F1 int32 }
type GenS4 struct { F0 uint32 }
func genFn0() string {
    var v7 uint = 255
    var v8 uint = (^(v7))
    v9 := GenS1{F0: -2147483648}
    var v10 int16 = 42
    var v11 uint32 = 42
    _ = v10
    _ = v11
    _ = v7
    _ = v8
    _ = v9
    return "hello"
}

func main() {
    var v1 int16 = -32768
    v2 := make(map[uint32]string)
    var v3 uint8 = 1
    var v4 int64 = -1
    v5 := make(map[string]int)
    v6 := []int{2147483647, 4294967295, 173}
    if false {
        v1 = v1
    }
    delete(v2, 0)
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    v4 += v4
    switch "\n\t" {
    case "\n\t":
        v12 := GenS2{F0: 206}
        _ = v12
    case "hello":
        if false {
            switch 32767 {
            case 32768:
                var v13 int = -1
                _ = v13
            case 42:
                v14 := v5["a"]
                _ = v14
            default:
                var v15 uint32 = 4294967295
                _ = v15
            }
            fmt.Println(v3)
        } else {
            for fi16 := 0; fi16 < 11; fi16++ {
                if false {
                    v17 := GenS2{F0: 4294967295}
                    v3 -= v3
                    var v18 int = (fi16 * 2147483647)
                    _ = v17
                    _ = v18
                }
                fmt.Println(fi16)
                var v19 uint = 255
                _ = fi16
                _ = v19
            }
            if true {
                fmt.Println(v3, v4)
            }
        }
    }
    v20 := GenS3{F0: 2147483647, F1: 0}
    for fi21 := 0; fi21 < 2; fi21++ {
        v22 := v2[42]
        _ = fi21
        _ = v22
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    if false {
        v6 = append(v6, -2147483648)
        var v23 int16 = 682
        switch "\n\t" {
        case "a":
            v4 += 1
        default:
            v24 := GenS4{F0: 2147483647}
            _ = v24
        }
        _ = v23
    }
    v4 += (v4 | (-(v4)))
    var v25 float64 = (((-2.25 / 3.14159) * -2.25) - 3.14159)
    switch ("hello" + "hello") {
    case "hello":
        switch "a" {
        case "hello":
            if (0 >= -84) {
                for fi26 := 0; fi26 < 6; fi26++ {
                    v6 = append(v6, (fi26 >> ((255) & 63)))
                    if (!true) {
                        var v27 uint8 = 127
                        v1 += -1
                        v1 = 32767
                        _ = v27
                    }
                    fmt.Println(v4)
                    _ = fi26
                }
                v1 += v1
                if (32767 >= 159) {
                    if false {
                        var v28 bool = false
                        _ = v28
                    } else {
                        v3 += 186
                        fmt.Println(v4, v4, v1)
                    }
                    var v29 int64 = v4
                    _ = v29
                }
            }
        case "\n\t":
            v30 := GenS4{F0: 425}
            _ = v30
        case "":
            v3 += (^(v3))
        default:
            fmt.Println(v25)
        }
    default:
        v1 += (v1 ^ v1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
