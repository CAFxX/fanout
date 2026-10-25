package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 int8 }

func main() {
    v1 := make(map[string]int)
    v2 := make(map[string]uint32)
    var v3 int16 = -1
    v3 = (v3 % (((v3 * (^(v3)))) | 1))
    fmt.Println(v3)
    if false {
        var v4 int8 = -104
        _ = v4
    }
    v5 := GenS1{F0: 8, F1: 4294967295, F2: -128}
    fmt.Println(v3)
    v6 := GenS1{F0: 42, F1: 429, F2: -128}
    fmt.Println(v3)
    var v7 int32 = 525
    var v8 uint8 = 66
    var v9 int64 = 843
    for fi10 := 0; fi10 < 15; fi10++ {
        var v11 float32 = 0.0
        if true {
            fmt.Println(fi10)
        } else {
            var v12 int64 = 640
            var v13 uint16 = 76
            _ = v12
            _ = v13
        }
        _ = fi10
        _ = v11
    }
    delete(v2, "")
    switch ("foo bar" + "a") {
    case "foo bar":
        for fi14 := 0; fi14 < 9; fi14++ {
            v15 := []int{(fi14 / (((-(fi14))) | 1)), 142, fi14, fi14}
            v3 += 0
            var v16 float32 = 3.14159
            _ = fi14
            _ = v15
            _ = v16
        }
    default:
        fmt.Println(v3, v3)
    }
    fmt.Println(v8, v3, v8)
    if ("a" != ("foo bar" + "foo bar")) {
        for fi17 := 0; fi17 < 2; fi17++ {
            var v18 int64 = (v9 + v9)
            v3 -= int16(v7)
            var v19 uint16 = 1
            _ = fi17
            _ = v18
            _ = v19
        }
    } else {
        v3 = -1
        var v20 int32 = (v7 - v7)
        switch 9223372036854775807 {
        case 128:
            if (81 >= int(v8)) {
                v3 = v3
                v7 += (v7 % ((v20) | 1))
            } else {
                switch ("a" + ("foo bar" + ("hello" + "a"))) {
                case "foo bar":
                    if (637 == 42) {
                        v7 -= v20
                        var v21 float64 = (1.5 + (0.1 - 3.14159))
                        _ = v21
                    } else {
                        v22 := []int{32767, 128, 32768}
                        var v23 float32 = (float32((0.1 * 1.5)) - ((3.14159 + (-2.25 * (0.0 / 3.14159))) + 1.5))
                        var v24 int64 = -665
                        _ = v22
                        _ = v23
                        _ = v24
                    }
                case "hello":
                    v25 := GenS1{F0: 42, F1: 1, F2: -1}
                    _ = v25
                default:
                    v1["a"] = 0
                }
            }
        case 32768:
            if (!false) {
                v1["foo bar"] = -783
                v26 := GenS1{F0: 128, F1: 673, F2: 23}
                if false {
                    v27 := GenS1{F0: 103, F1: 1, F2: 0}
                    _ = v27
                } else {
                    var v28 int32 = (v20 % ((-1) | 1))
                    _ = v28
                }
                _ = v26
            }
        case -2147483648:
            v3 = v3
        }
        _ = v20
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
