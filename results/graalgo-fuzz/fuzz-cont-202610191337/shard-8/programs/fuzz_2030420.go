package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }
type GenS2 struct { F0 int64; F1 int32 }
type GenS3 struct { F0 int8 }

func main() {
    var v1 float64 = 1.5
    var v2 int = -932
    var v3 int8 = -128
    v2 += v2
    switch (v2 % ((9223372036854775807) | 1)) {
    case -589:
        if (255 == 32767) {
            v3 += (v3 * (v3 ^ v3))
            var v4 float32 = -2.25
            _ = v4
        }
    case 36:
        v5 := []int{(v2 * (v2 / ((-1) | 1)))}
        _ = v5
    case 127:
        v1 -= (v1 / v1)
    default:
        for fi6 := 0; fi6 < 18; fi6++ {
            var v7 int16 = 0
            v3 += v3
            var v8 uint8 = 2
            _ = fi6
            _ = v7
            _ = v8
        }
    }
    fmt.Println(v3, v1, v2)
    var v9 int = -2147483648
    if true {
        v1 = v1
        v10 := GenS1{F0: 127, F1: 42}
        _ = v10
    }
    v11 := GenS1{F0: 114, F1: -2147483648}
    if false {
        for fi12 := 0; fi12 < 15; fi12++ {
            switch ("foo bar" + ("\n\t" + ("foo bar" + ""))) {
            case "\n\t":
                v13 := GenS2{F0: 1, F1: 0}
                _ = v13
            }
            v1 -= ((v1 * v1) - (v1 / -2.25))
            if (858 != 772) {
                fmt.Println(v3, v3, v3)
                fmt.Println(v9, v2, v3)
            }
            _ = fi12
        }
        v2 = v9
        v14 := []int{v2, v9, v2}
        _ = v14
    } else {
        var v15 uint32 = 298
        v16 := []int{}
        _ = v15
        _ = v16
    }
    v9 += (v2 + (-(v9)))
    switch "foo bar" {
    case "hello":
        v17 := GenS1{F0: 2, F1: 0}
        _ = v17
    case "foo bar":
        var v18 uint = 614
        _ = v18
    case "":
        v9 += v2
    }
    v9 -= 769
    v9 += v2
    for fi19 := 0; fi19 < 4; fi19++ {
        switch (("" + "foo bar") + ("a" + "a")) {
        case "hello":
            switch "hello" {
            case "foo bar":
                var v20 int32 = -1
                _ = v20
            }
        }
        _ = fi19
    }
    v9 = (-(v2))
    v21 := GenS3{F0: (v3 / ((-62) | 1))}
    var v22 float32 = 1.5
    var v23 float64 = 0.0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v9
}
