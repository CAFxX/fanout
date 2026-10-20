package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string }
type GenS2 struct { F0 uint32; F1 uint64 }

func main() {
    var v1 uint = 65535
    var v2 float32 = float32(((1.5 * 1e10) + 3.14159))
    var v3 uint = v1
    var v4 string = "\n\t"
    var v5 float64 = 3.14159
    switch 1 {
    case -2147483648:
        var v6 int16 = 42
        _ = v6
    case 4294967295:
        switch 32768 {
        case -9223372036854775808:
            for fi7 := 0; fi7 < 5; fi7++ {
                v8 := []int{fi7, -852}
                if true {
                    fmt.Println(v3, v1)
                }
                var v9 int = fi7
                _ = fi7
                _ = v8
                _ = v9
            }
        default:
            if (128 == 1) {
                v10 := GenS1{F0: 1, F1: ""}
                _ = v10
            }
        }
    }
    var v11 int64 = 9223372036854775807
    for fi12 := 0; fi12 < 6; fi12++ {
        var v13 uint8 = 255
        fmt.Println(v5)
        for fi14 := 0; fi14 < 17; fi14++ {
            v15 := GenS1{F0: 516, F1: v4}
            _ = fi14
            _ = v15
        }
        _ = fi12
        _ = v13
    }
    v16 := []int{-913, 127, 53, -976}
    v16 = append(v16, -1)
    for fi17 := 0; fi17 < 15; fi17++ {
        v5 -= v5
        if ((0 < 96) && (true || (32767 != 582))) {
            var v18 int32 = -2147483648
            var v19 int64 = v11
            var v20 uint32 = 124
            _ = v18
            _ = v19
            _ = v20
        }
        for fi21 := 0; fi21 < 6; fi21++ {
            v5 = (v5 / (3.14159 * v5))
            v4 = ((v4 + v4) + "hello")
            if true {
                var v22 int8 = -1
                _ = v22
            } else {
                fmt.Println(v2)
                v23 := GenS1{F0: 1, F1: "\n\t"}
                switch "\n\t" {
                case "foo bar":
                    fmt.Println(v1, v3, v4)
                }
                _ = v23
            }
            fmt.Println(fi21)
            switch ("\n\t" + (v4 + ("hello" + (v4 + ("a" + "\n\t"))))) {
            case "\n\t":
                var v24 int = (fi21 & fi21)
                _ = v24
            case "a":
                var v25 string = v4
                _ = v25
            default:
                fmt.Println(v2, fi17)
            }
            switch "" {
            case "":
                if len(v16) > 0 {
                    v16[0] = fi17
                }
            case "foo bar":
                switch 864 {
                case 42:
                    var v26 float64 = 1.5
                    _ = v26
                case 682:
                    v4 = ""
                case 0:
                    v27 := GenS1{F0: 4294967295, F1: v4}
                    _ = v27
                default:
                    var v28 int8 = -1
                    _ = v28
                }
            }
            var v29 float32 = v2
            v29 = 3.14159
            fmt.Println(v11)
            _ = fi21
            _ = v29
        }
        _ = fi17
    }
    v30 := GenS2{F0: 4294967295, F1: 1}
    v31 := GenS2{F0: 979, F1: 1}
    v16 = append(v16, -2147483648)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v16
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}
