package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int64; F1 int32 }
func genFn0() bool {
    var v7 int = 9223372036854775807
    var v8 int32 = -23
    var v9 uint64 = 791
    for fi10 := 0; fi10 < 4; fi10++ {
        switch (v7 / (((v7 >> ((224) & 63))) | 1)) {
        case 383:
            var v11 uint16 = 65535
            _ = v11
        case 484:
            switch "foo bar" {
            case "":
                v12 := GenS1{F0: false, F1: 151}
                _ = v12
            case "\n\t":
                fmt.Println(v8)
            }
        default:
            v13 := []int{(v7 << ((65535) & 63))}
            _ = v13
        }
        var v14 uint16 = 65535
        _ = fi10
        _ = v14
    }
    if (((false || true) || true) && (1 >= 9223372036854775807)) {
        v7 = v7
    } else {
        var v15 float64 = (3.14159 - float64((1e10 / (3.14159 / 1.5))))
        switch (^(v7)) {
        case 42:
            var v16 int = v7
            _ = v16
        case -2147483648:
            v17 := GenS1{F0: ((((3.14159 + (((float32(3.14159) - (0.0 * (float32(1e10) - 0.1))) / -2.25) * 1.5)) - (1e10 - float32(v15))) * 1.5) == (3.14159 / 0.1)), F1: -443}
            _ = v17
        default:
            switch ("" + "a") {
            case "a":
                fmt.Println(v7, v15)
            case "\n\t":
                v9 = 4294967295
            case "":
                v18 := []int{v7, (v7 + (v7 ^ 4294967295)), 0, (v7 / (((v7 >> ((42) & 63))) | 1))}
                _ = v18
            default:
                if true {
                    v7 += (v7 ^ 9223372036854775807)
                } else {
                    var v19 uint64 = (v9 ^ (v9 * 0))
                    _ = v19
                }
            }
        }
        var v20 uint64 = 1
        _ = v15
        _ = v20
    }
    _ = v7
    _ = v8
    _ = v9
    return (false || true)
}

func main() {
    v1 := make(map[int]uint32)
    var v2 float32 = 3.14159
    v3 := make(map[uint32]int)
    v4 := []string{(("hello" + ("foo bar" + "")) + "hello"), "\n\t"}
    var v5 uint32 = 42
    var v6 float64 = 0.1
    for fi21 := 0; fi21 < 3; fi21++ {
        v2 -= (((((v2 - float32(1e10)) * v2) - float32((((1e10 + 1e10) * v6) / v6))) - float32(v6)) + v2)
        _ = fi21
    }
    for fi22 := 0; fi22 < 16; fi22++ {
        v6 = 3.14159
        _ = fi22
    }
    v23 := GenS2{F0: 534}
    for fi24 := 0; fi24 < 11; fi24++ {
        fmt.Println(v2, fi24, v6)
        for fi25 := 0; fi25 < 5; fi25++ {
            v2 = float32(v6)
            _ = fi25
        }
        for fi26 := 0; fi26 < 6; fi26++ {
            v27 := GenS3{F0: 9223372036854775807, F1: 251}
            _ = fi26
            _ = v27
        }
        _ = fi24
    }
    v2 -= v2
    v5 -= 865
    v2 = ((0.0 - (float32(0.0) * (v2 * (1e10 - ((float32(-2.25) - -2.25) - (0.1 - (((float32(1.5) - v2) + v2) * float32(1e10)))))))) + v2)
    var v28 uint8 = 0
    fmt.Println(v28, v5, v5)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v29 := GenS1{F0: false, F1: 39}
    for fi30 := 0; fi30 < 20; fi30++ {
        for fi31 := 0; fi31 < 19; fi31++ {
            var v32 uint16 = 515
            if len(v4) > 0 {
                v4[0] = "hello"
            }
            v32 -= 1
            _ = fi31
            _ = v32
        }
        _ = fi30
    }
    var v33 int8 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v28
    _ = v29
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
