package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint; F1 uint16; F2 string }
type GenS3 struct { F0 uint16; F1 int32; F2 int32 }

func main() {
    v1 := make(map[int]string)
    var v2 float32 = ((1.5 + 1e10) + (0.0 / 3.14159))
    v3 := []uint32{174, 0, 2147483648}
    v2 = v2
    v2 += v2
    if len(v3) > 0 {
        v3[0] = 731
    }
    v2 += v2
    var v4 bool = (true && ((!false) || (47 == 185)))
    var v5 int = 170
    v6 := GenS1{F0: 0}
    v7 := GenS1{F0: -128}
    var v8 uint64 = 42
    if true {
        var v9 float32 = v2
        v2 -= v2
        _ = v9
    } else {
        for fi10 := 0; fi10 < 8; fi10++ {
            var v11 uint = 65535
            _ = fi10
            _ = v11
        }
    }
    switch "hello" {
    case "foo bar":
        v12 := GenS1{F0: -125}
        _ = v12
    case "hello":
        var v13 uint8 = 255
        _ = v13
    default:
        var v14 uint32 = 2147483648
        _ = v14
    }
    switch v5 {
    case 1:
        v15 := GenS2{F0: 65535, F1: 0, F2: "a"}
        _ = v15
    case 880:
        switch ("hello" + "") {
        case "\n\t":
            v16 := GenS2{F0: 21, F1: 32767, F2: "hello"}
            _ = v16
        case "hello":
            v8 = v8
        case "a":
            var v17 string = (("foo bar" + "a") + "foo bar")
            _ = v17
        default:
            switch (-(v5)) {
            case 42:
                v8 -= (^(v8))
            case -9223372036854775808:
                v5 -= 817
            case 32768:
                v18 := GenS2{F0: 1, F1: 65535, F2: ("\n\t" + "hello")}
                _ = v18
            default:
                var v19 uint8 = 255
                _ = v19
            }
        }
    }
    v3 = append(v3, 391)
    v2 += 0.1
    v20 := GenS3{F0: 735, F1: 583, F2: -2147483648}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
