package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 bool; F1 uint8 }
type GenS3 struct { F0 uint32; F1 uint; F2 bool }

func main() {
    var v1 int = 128
    var v2 int32 = 560
    v3 := make(map[string]bool)
    var v4 float64 = 0.0
    switch (v1 + v1) {
    case 42:
        var v5 uint32 = 2
        _ = v5
    case -9223372036854775808:
        switch 0 {
        case 32768:
            v1 = (v1 ^ v1)
        }
    default:
        fmt.Println(v4, v1)
    }
    v6 := GenS1{F0: 255}
    fmt.Println(v1, v1, v2)
    v1 = (v1 ^ (v1 >> ((763) & 63)))
    for fi7 := 0; fi7 < 19; fi7++ {
        for fi8 := 0; fi8 < 10; fi8++ {
            if (false || (9223372036854775807 >= -255)) {
                v2 = v2
            }
            switch "\n\t" {
            case "\n\t":
                if false {
                    v9 := v3[(("foo bar" + ("" + ("a" + "hello"))) + ("a" + ("" + "hello")))]
                    _ = v9
                } else {
                    switch (("hello" + "\n\t") + "a") {
                    case "hello":
                        var v10 uint32 = 242
                        _ = v10
                    case "":
                        v4 += (1e10 / v4)
                    case "foo bar":
                        v11 := GenS2{F0: ((((("" + "\n\t") + "\n\t") + "foo bar") + ("\n\t" + "a")) > "foo bar"), F1: 42}
                        _ = v11
                    }
                }
            case "a":
                v2 = int32(v1)
            }
            if true {
                var v12 int8 = 100
                v13 := GenS1{F0: 255}
                var v14 int8 = (v12 >> ((834) & 63))
                _ = v12
                _ = v13
                _ = v14
            } else {
                v15 := v3["\n\t"]
                var v16 uint = 314
                delete(v3, ("" + "a"))
                _ = v15
                _ = v16
            }
            _ = fi8
        }
        switch "a" {
        case "\n\t":
            if false {
                for fi17 := 0; fi17 < 20; fi17++ {
                    v18 := []int{fi17, (^(fi17))}
                    _ = fi17
                    _ = v18
                }
                v4 = (1e10 * 0.0)
                var v19 int32 = (v2 & v2)
                _ = v19
            }
        }
        _ = fi7
    }
    if true {
        fmt.Println(v4)
    } else {
        v1 += (v1 | v1)
        v2 = (v2 >> ((4294967295) & 63))
    }
    if (((((1.5 - ((v4 / (((v4 - v4) / ((v4 * v4) / v4)) / float64(float32(v4)))) - v4)) - 3.14159) <= v4) && (false && true)) || (42 <= 1)) {
        if true {
            v20 := []int{}
            fmt.Println(v2)
            if (((float64(1.5) * 1e10) <= v4) || ((((0.1 - ((((0.1 + 3.14159) * 0.1) + (1e10 * float32(float64(0.0)))) * (3.14159 * 0.0))) + 3.14159) <= (-2.25 * 1e10)) || false)) {
                if len(v20) > 0 {
                    v20[0] = (v1 >> ((0) & 63))
                }
                fmt.Println(v4, v2)
            } else {
                fmt.Println(v4, v1, v4)
                for fi21 := 0; fi21 < 9; fi21++ {
                    var v22 int32 = (v2 - (v2 ^ v2))
                    _ = len(v3)
                    _ = fi21
                    _ = v22
                }
            }
            _ = v20
        } else {
            v4 += v4
            if (false || ((1.5 + -2.25) <= ((((v4 + (v4 * v4)) + 0.1) - v4) / 1.0))) {
                v23 := []int{-152, (v1 / ((v1) | 1))}
                _ = v23
            }
        }
    }
    v4 = 0.1
    switch "" {
    case "foo bar":
        v2 += -1
    case "hello":
        var v24 string = "\n\t"
        _ = v24
    case "":
        v25 := GenS3{F0: 369, F1: 18446744073709551615, F2: false}
        _ = v25
    }
    v26 := []int{(-(v1)), v1, v1}
    var v27 int32 = -1
    for fi28 := 0; fi28 < 12; fi28++ {
        v29 := GenS2{F0: true, F1: 201}
        v2 -= (v2 | (v27 * -293))
        _ = fi28
        _ = v29
    }
    if false {
        _ = len(v3)
    } else {
        fmt.Println(v1, v1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v6
}
