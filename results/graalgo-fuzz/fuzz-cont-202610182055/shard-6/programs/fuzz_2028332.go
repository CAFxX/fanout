package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string }

func main() {
    var v1 uint8 = 220
    var v2 uint16 = 255
    var v3 int16 = 42
    v4 := make(map[int]uint32)
    var v5 uint = 849
    var v6 uint64 = 4294967295
    for fi7 := 0; fi7 < 20; fi7++ {
        var v8 int = fi7
        v8 = (v8 << (((v5 % ((v5) | 1))) & 63))
        _ = fi7
        _ = v8
    }
    var v9 int64 = 1
    v2 = 42
    switch ("a" + ("hello" + ("foo bar" + "foo bar"))) {
    case "a":
        v10 := []int{}
        _ = v10
    }
    switch "foo bar" {
    case "":
        switch 32767 {
        case -9223372036854775808:
            var v11 uint8 = v1
            _ = v11
        default:
            fmt.Println(v5, v5, v2)
        }
    case "foo bar":
        var v12 int = 4294967295
        _ = v12
    case "a":
        for fi13 := 0; fi13 < 17; fi13++ {
            for fi14 := 0; fi14 < 14; fi14++ {
                switch (^(fi14)) {
                case 9223372036854775807:
                    var v15 int8 = -1
                    _ = v15
                case 0:
                    v16 := GenS1{F0: (^(fi13)), F1: "\n\t"}
                    _ = v16
                }
                _ = fi14
            }
            var v17 uint64 = v6
            _ = len(v4)
            _ = fi13
            _ = v17
        }
    default:
        fmt.Println(v6, v9)
    }
    var v18 uint = (v5 * (^(v5)))
    var v19 uint = (^(v5))
    v20 := []int{0, 4294967295, -58}
    for fi21 := 0; fi21 < 4; fi21++ {
        var v22 float32 = 3.14159
        switch int(v2) {
        case -577:
            var v23 uint = (v5 << ((uint(v1)) & 63))
            _ = v23
        default:
            for fi24 := 0; fi24 < 13; fi24++ {
                var v25 uint32 = 4294967295
                v26 := GenS1{F0: fi21, F1: "hello"}
                _ = fi24
                _ = v25
                _ = v26
            }
        }
        switch fi21 {
        case -1:
            v27 := GenS1{F0: fi21, F1: "a"}
            _ = v27
        default:
            delete(v4, (fi21 | (fi21 - -2147483648)))
        }
        _ = fi21
        _ = v22
    }
    var v28 uint8 = v1
    if true {
        var v29 uint8 = v1
        _ = v29
    } else {
        var v30 uint8 = 128
        _ = v30
    }
    v5 = uint(v6)
    v28 += (v28 >> ((243) & 63))
    v31 := GenS1{F0: -2147483648, F1: "foo bar"}
    for fi32 := 0; fi32 < 15; fi32++ {
        var v33 float64 = (-2.25 * -2.25)
        _ = fi32
        _ = v33
    }
    var v34 int16 = (^(v3))
    delete(v4, 380)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v28
    _ = v3
    _ = v31
    _ = v34
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
