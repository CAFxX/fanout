package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int32; F1 string }

func main() {
    v1 := make(map[uint32]string)
    var v2 int32 = -2147483648
    v3 := []int{}
    var v4 float64 = 0.0
    var v5 int32 = -2147483648
    var v6 string = "hello"
    for fi7 := 0; fi7 < 7; fi7++ {
        if (1e10 < 0.1) {
            switch fi7 {
            case -2147483648:
                var v8 string = (v6 + "a")
                _ = v8
            case 124:
                for fi9 := 0; fi9 < 7; fi9++ {
                    var v10 int64 = -1
                    fmt.Println(fi9)
                    var v11 int64 = (v10 + 9223372036854775807)
                    _ = fi9
                    _ = v10
                    _ = v11
                }
            }
            v5 -= (v5 % ((v5) | 1))
            fmt.Println(v2, v4, v5)
        }
        switch fi7 {
        case 32768:
            delete(v1, 65535)
        case -457:
            v2 = -2147483648
        }
        _ = fi7
    }
    var v12 float64 = 0.1
    if (708 >= 255) {
        delete(v1, 2147483648)
        if false {
            _ = len(v1)
            var v13 float32 = (0.0 * (0.1 * (0.1 * ((float32(v4) * 1.5) + 1e10))))
            for fi14 := 0; fi14 < 16; fi14++ {
                v15 := GenS1{F0: 161}
                _ = fi14
                _ = v15
            }
            _ = v13
        }
        v4 = (-2.25 - 1e10)
    } else {
        v16 := GenS2{F0: 18446744073709551615}
        for fi17 := 0; fi17 < 5; fi17++ {
            var v18 uint8 = 2
            v19 := GenS1{F0: 42}
            _ = fi17
            _ = v18
            _ = v19
        }
        v12 -= (0.1 - (float64(0.1) / (v4 * v12)))
        _ = v16
    }
    if true {
        switch (v6 + "hello") {
        case "hello":
            v2 = (v2 << ((0) & 63))
        case "\n\t":
            var v20 float64 = float64((float32(v4) / 1.0))
            _ = v20
        default:
            delete(v1, 2147483647)
        }
        fmt.Println(v4, v6)
        for fi21 := 0; fi21 < 6; fi21++ {
            v2 = 0
            _ = fi21
        }
    }
    v22 := GenS3{F0: (v5 * 1), F1: v6}
    switch -9223372036854775808 {
    case 32768:
        switch 32767 {
        case 532:
            var v23 int64 = 42
            _ = v23
        case 9223372036854775807:
            var v24 int = 9223372036854775807
            _ = v24
        case 2147483647:
            v1[0] = v6
        }
    }
    v6 = v6
    var v25 float64 = v4
    v2 += v2
    var v26 int = 1
    switch ("foo bar" + v6) {
    case "hello":
        v4 = v12
    case "\n\t":
        v12 = v4
    case "":
        fmt.Println(v12, v26, v2)
    default:
        fmt.Println(v12, v25)
    }
    fmt.Println(v4, v5, v26)
    fmt.Println(v4, v12)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v22
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
