package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint; F1 int64 }
type GenS3 struct { F0 uint16; F1 uint8; F2 int }
type GenS4 struct { F0 uint; F1 uint32; F2 int16 }

func main() {
    var v1 int = 654
    var v2 int16 = 32767
    var v3 int8 = 100
    v2 += v2
    v4 := []int{(-(v1)), v1, v1}
    for fi5 := 0; fi5 < 6; fi5++ {
        v1 = int(v2)
        v1 -= (fi5 + -9223372036854775808)
        _ = fi5
    }
    fmt.Println(v3, v2, v2)
    for fi6 := 0; fi6 < 7; fi6++ {
        var v7 uint8 = 128
        for fi8 := 0; fi8 < 1; fi8++ {
            if len(v4) > 0 {
                v4[0] = 32767
            }
            _ = fi8
        }
        _ = fi6
        _ = v7
    }
    if false {
        v3 = -98
        if (255 > 4) {
            if true {
                var v9 int8 = v3
                _ = v9
            }
            v2 -= int16(v3)
            var v10 string = ("\n\t" + ("" + "\n\t"))
            _ = v10
        } else {
            var v11 int16 = 42
            v12 := GenS1{F0: 42}
            fmt.Println(v11)
            _ = v11
            _ = v12
        }
    }
    var v13 int8 = v3
    switch ("\n\t" + (("hello" + "") + "hello")) {
    case "":
        if ((!false) && false) {
            v1 = 938
            for fi14 := 0; fi14 < 7; fi14++ {
                v15 := GenS2{F0: 459, F1: 9223372036854775807}
                _ = fi14
                _ = v15
            }
            v16 := GenS3{F0: 2, F1: 127, F2: 703}
            _ = v16
        } else {
            v4 = append(v4, v1)
        }
    default:
        _ = len(v4)
    }
    _ = len(v4)
    switch 0 {
    case -1:
        var v17 string = ""
        _ = v17
    }
    var v18 string = ("a" + "a")
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v19 := GenS4{F0: 136, F1: 980, F2: (v2 - int16(v1))}
    v20 := GenS2{F0: 1, F1: 9223372036854775807}
    switch v18 {
    case "":
        if true {
            switch (v18 + "foo bar") {
            case "foo bar":
                v21 := GenS1{F0: 111}
                _ = v21
            case "a":
                var v22 float32 = 0.1
                _ = v22
            default:
                var v23 uint = 42
                _ = v23
            }
            var v24 float32 = (float32(1.5) / -2.25)
            v25 := GenS4{F0: 2, F1: 990, F2: v2}
            _ = v24
            _ = v25
        }
    default:
        var v26 uint64 = 1
        _ = v26
    }
    v27 := GenS4{F0: 1, F1: 399, F2: 1000}
    v2 += v2
    v4 = append(v4, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v13
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v27
    _ = v3
    _ = v4
}
