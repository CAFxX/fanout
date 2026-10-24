package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 int8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint16; F1 int32; F2 uint8 }
type GenS6 struct { F0 int }
func genFn0() int16 {
    var v4 int = 0
    v5 := []int{(v4 & v4)}
    v6 := GenS1{F0: -1, F1: ("foo bar" + ((("hello" + (("\n\t" + "foo bar") + "foo bar")) + "a") + ("\n\t" + ""))), F2: -86}
    _ = v4
    _ = v5
    _ = v6
    return int16(v4)
}

func main() {
    v1 := make(map[string]bool)
    var v2 int8 = 1
    v3 := make(map[int]string)
    v2 = v2
    v7 := GenS2{F0: 1}
    v2 = v2
    v8 := GenS1{F0: -1, F1: "hello", F2: (^(v2))}
    switch -1 {
    case -953:
        switch ("" + ("foo bar" + "")) {
        case "foo bar":
            var v9 int64 = 0
            _ = v9
        case "hello":
            v10 := []int{32767, 0, -214}
            _ = v10
        case "":
            switch -625 {
            case 0:
                var v11 uint32 = 0
                _ = v11
            case -2147483648:
                if ((3.14159 + ((0.1 * 0.0) + (((((1.5 + float32((((((0.0 * 0.1) + 0.0) - 1e10) - 1e10) * 0.0))) - 1.5) + -2.25) * 1e10) - ((3.14159 / -2.25) / 3.14159)))) != ((3.14159 - -2.25) - (0.1 / -2.25))) {
                    v2 = v2
                    v12 := GenS3{F0: 1}
                    _ = v12
                } else {
                    switch 1 {
                    case 871:
                        v13 := GenS1{F0: -1, F1: "a", F2: v2}
                        _ = v13
                    }
                }
            case 448:
                v2 += (v2 / ((v2) | 1))
            }
        default:
            var v14 uint64 = 18446744073709551615
            _ = v14
        }
    case -851:
        if ((-2.25 - 3.14159) > 3.14159) {
            fmt.Println(v2)
            v15 := GenS1{F0: 315, F1: "\n\t", F2: -128}
            _ = v15
        } else {
            fmt.Println(v2)
        }
    case 210:
        if (true || (!false)) {
            var v16 bool = (81 >= 128)
            _ = v16
        }
    default:
        if true {
            v17 := GenS4{F0: 42}
            for fi18 := 0; fi18 < 4; fi18++ {
                fmt.Println(v2)
                _ = fi18
            }
            v19 := GenS5{F0: 0, F1: 0, F2: 1}
            _ = v17
            _ = v19
        }
    }
    var v20 int64 = -206
    v20 = 192
    v2 = (^(v2))
    for fi21 := 0; fi21 < 2; fi21++ {
        v22 := v1[("\n\t" + "")]
        v22 = v22
        if (861 == 1) {
            var v23 uint32 = 65535
            for fi24 := 0; fi24 < 5; fi24++ {
                var v25 float64 = (1.5 + 0.1)
                _ = fi24
                _ = v25
            }
            _ = v23
        }
        _ = fi21
        _ = v22
    }
    v20 = -1
    v26 := GenS3{F0: 9223372036854775807}
    v27 := GenS6{F0: -348}
    var v28 bool = (("\n\t" + "") < "hello")
    var v29 bool = v28
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v20
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v7
    _ = v8
}
