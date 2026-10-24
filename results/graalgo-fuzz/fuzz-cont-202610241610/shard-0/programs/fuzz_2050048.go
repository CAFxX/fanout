package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int; F2 int64 }
type GenS2 struct { F0 int32 }

func main() {
    var v1 uint64 = 312
    var v2 uint = 2
    var v3 int64 = 530
    var v4 int32 = -1
    var v5 float64 = 3.14159
    var v6 uint32 = 774
    v7 := []int{-620, 127, 128, 1}
    for fi8 := 0; fi8 < 1; fi8++ {
        for fi9 := 0; fi9 < 6; fi9++ {
            var v10 int = fi8
            var v11 float32 = -2.25
            _ = fi9
            _ = v10
            _ = v11
        }
        var v12 bool = false
        _ = fi8
        _ = v12
    }
    if (true || (127 < 127)) {
        v13 := GenS1{F0: v6, F1: 900, F2: v3}
        _ = v13
    }
    fmt.Println(v3, v6, v5)
    if len(v7) > 0 {
        v7[0] = -875
    }
    fmt.Println(v2)
    if true {
        if len(v7) > 0 {
            v7[0] = 32767
        }
    } else {
        for fi14 := 0; fi14 < 6; fi14++ {
            v15 := GenS2{F0: 154}
            _ = fi14
            _ = v15
        }
        switch "hello" {
        case "a":
            if true {
                var v16 uint8 = 42
                if len(v7) > 0 {
                    v7[0] = -2147483648
                }
                _ = v16
            }
        case "foo bar":
            if false {
                if (1 >= (^(v3))) {
                    for fi17 := 0; fi17 < 16; fi17++ {
                        switch "hello" {
                        case "\n\t":
                            v4 += 1
                        case "":
                            _ = len(v7)
                        }
                        _ = fi17
                    }
                }
                v3 = (v3 >> ((1) & 63))
            }
        case "hello":
            v2 = 4294967295
        }
    }
    v3 = 44
    var v18 string = "\n\t"
    v19 := GenS2{F0: (v4 * (v4 & int32(v1)))}
    for fi20 := 0; fi20 < 12; fi20++ {
        var v21 uint8 = 127
        switch fi20 {
        case -1:
            fmt.Println(v3)
        case 128:
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
        case -825:
            var v22 string = v18
            _ = v22
        default:
            _ = len(v7)
        }
        var v23 uint32 = 421
        _ = fi20
        _ = v21
        _ = v23
    }
    v1 += 42
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
