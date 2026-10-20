package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint64; F1 int16; F2 int }
type GenS3 struct { F0 uint16; F1 int }

func main() {
    var v1 int16 = -719
    var v2 uint8 = 30
    var v3 int16 = v1
    if (true && true) {
        v4 := []int{4294967295, 0, 4294967295, -894}
        for fi5 := 0; fi5 < 16; fi5++ {
            var v6 float32 = ((3.14159 * 3.14159) + 1.5)
            _ = fi5
            _ = v6
        }
        for fi7 := 0; fi7 < 20; fi7++ {
            fmt.Println(fi7, v1, v1)
            _ = fi7
        }
        _ = v4
    }
    for fi8 := 0; fi8 < 7; fi8++ {
        v1 = v3
        _ = fi8
    }
    v1 += v1
    for fi9 := 0; fi9 < 19; fi9++ {
        v10 := GenS1{F0: (v1 * v1)}
        if false {
            v2 += (v2 << ((65535) & 63))
        } else {
            v11 := GenS2{F0: 4294967295, F1: (-(v1)), F2: (fi9 ^ fi9)}
            _ = v11
        }
        if (true || false) {
            var v12 uint = 255
            _ = v12
        } else {
            switch "foo bar" {
            case "a":
                var v13 float32 = float32((0.0 * 0.0))
                _ = v13
            case "foo bar":
                var v14 string = ("\n\t" + (((("hello" + "") + "a") + ("" + (("foo bar" + ("a" + ("hello" + "\n\t"))) + "hello"))) + "foo bar"))
                _ = v14
            default:
                v15 := []int{0}
                _ = v15
            }
            var v16 int8 = -128
            _ = v16
        }
        _ = fi9
        _ = v10
    }
    for fi17 := 0; fi17 < 6; fi17++ {
        fmt.Println(v1, v3, fi17)
        _ = fi17
    }
    v2 = v2
    switch 128 {
    case 128:
        v18 := GenS3{F0: 0, F1: 128}
        _ = v18
    case -736:
        if true {
            var v19 uint8 = v2
            var v20 float32 = -2.25
            fmt.Println(v19, v2, v3)
            _ = v19
            _ = v20
        }
    }
    v21 := []int{-562}
    if (300 == 506) {
        _ = len(v21)
        fmt.Println(v1)
        var v22 int32 = -573
        _ = v22
    }
    var v23 int8 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v21
    _ = v23
    _ = v3
}
