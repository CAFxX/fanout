package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string }
type GenS2 struct { F0 bool; F1 string }
func genFn0() uint16 {
    v4 := GenS1{F0: 42, F1: "\n\t"}
    var v5 uint8 = 128
    v5 = v5
    switch "" {
    case "":
        v5 -= (v5 ^ 255)
    case "hello":
        v5 = 0
    case "\n\t":
        v5 += v5
    default:
        v6 := []int{42, 0, 32768}
        _ = v6
    }
    for fi7 := 0; fi7 < 4; fi7++ {
        v8 := []int{fi7, fi7}
        _ = fi7
        _ = v8
    }
    _ = v4
    _ = v5
    return 705
}

func main() {
    var v1 int = -552
    v2 := []int{}
    var v3 float32 = (-2.25 * (0.0 - 1e10))
    for fi9 := 0; fi9 < 5; fi9++ {
        var v10 int32 = 42
        v11 := GenS2{F0: (255 <= 145), F1: (("foo bar" + (("foo bar" + "hello") + (("foo bar" + "\n\t") + "a"))) + "hello")}
        fmt.Println(v10)
        _ = fi9
        _ = v10
        _ = v11
    }
    for fi12 := 0; fi12 < 17; fi12++ {
        switch (fi12 ^ v1) {
        case -2147483648:
            fmt.Println(fi12)
        case 0:
            var v13 float64 = ((((1e10 * 1e10) + 0.1) * float64(-2.25)) + 1e10)
            _ = v13
        case 127:
            v3 = (1.5 + 3.14159)
        }
        if ("foo bar" > "hello") {
            v3 = 1.5
        } else {
            var v14 float32 = float32(-2.25)
            _ = v14
        }
        var v15 string = "foo bar"
        _ = fi12
        _ = v15
    }
    if true {
        if (42 < 1) {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        } else {
            v3 -= 1e10
            if true {
                var v16 int = -9223372036854775808
                _ = v16
            }
        }
        v2 = append(v2, v1)
        if (255 >= 220) {
            var v17 string = ("a" + "\n\t")
            switch 32767 {
            case 4294967295:
                if (!(0 != 42)) {
                    if true {
                        v1 = v1
                        fmt.Println(v3, v1, v17)
                        var v18 float64 = 3.14159
                        _ = v18
                    }
                } else {
                    switch 1 {
                    case -9223372036854775808:
                        fmt.Println(v1, v17, v1)
                    default:
                        v3 = (float32(1.5) / float32(1e10))
                    }
                }
            case 127:
                v2 = append(v2, v1)
            case -9223372036854775808:
                v17 = v17
            }
            v1 = 897
            _ = v17
        }
    }
    var v19 float64 = (-2.25 * 3.14159)
    v20 := GenS2{F0: false, F1: ("foo bar" + "")}
    fmt.Println(v19)
    v2 = append(v2, (v1 - v1))
    var v21 uint = 2
    switch v1 {
    case 180:
        fmt.Println(v3, v19)
    case 1:
        v1 += (^(v1))
    case -9223372036854775808:
        fmt.Println(v21)
    default:
        for fi22 := 0; fi22 < 11; fi22++ {
            for fi23 := 0; fi23 < 13; fi23++ {
                var v24 string = ("hello" + "hello")
                v3 -= v3
                _ = fi23
                _ = v24
            }
            _ = fi22
        }
    }
    v1 -= (^(v1))
    for fi25 := 0; fi25 < 17; fi25++ {
        var v26 int8 = 127
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        var v27 string = ("foo bar" + (("foo bar" + "hello") + ("\n\t" + "hello")))
        _ = fi25
        _ = v26
        _ = v27
    }
    var v28 int16 = -32768
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v28
    _ = v3
}
