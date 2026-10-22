package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int16; F1 int64 }

func main() {
    var v1 float64 = 3.14159
    v2 := []uint32{2147483647, 4294967295, 1}
    var v3 int8 = -30
    v4 := GenS1{F0: 42}
    var v5 uint32 = 42
    switch 9223372036854775807 {
    case -9223372036854775808:
        v5 = v5
    case 9223372036854775807:
        for fi6 := 0; fi6 < 17; fi6++ {
            var v7 string = ("" + "foo bar")
            _ = fi6
            _ = v7
        }
    case 0:
        if true {
            v8 := GenS1{F0: -9223372036854775808}
            for fi9 := 0; fi9 < 11; fi9++ {
                v3 -= 0
                _ = fi9
            }
            _ = v8
        } else {
            v1 -= (1e10 * 0.0)
            switch 42 {
            case -1:
                var v10 int32 = -309
                _ = v10
            default:
                switch ("hello" + "a") {
                case "hello":
                    fmt.Println(v1, v3, v1)
                default:
                    switch 23 {
                    case 9223372036854775807:
                        v11 := GenS2{F0: 1, F1: 42}
                        _ = v11
                    case -875:
                        switch 42 {
                        case -2147483648:
                            v5 = (v5 - 2)
                        }
                    case 1:
                        v2 = append(v2, (v5 >> ((65535) & 63)))
                    default:
                        var v12 uint64 = 480
                        _ = v12
                    }
                }
            }
        }
    default:
        v13 := GenS2{F0: 0, F1: 0}
        _ = v13
    }
    var v14 float32 = (0.0 / (-2.25 + ((((1e10 - 1.5) + 0.0) + ((float32((v1 - v1)) / float32(v1)) + (3.14159 / 0.1))) - float32((1.5 / (v1 / v1))))))
    _ = len(v2)
    v15 := GenS1{F0: 42}
    if (false || false) {
        v14 = (3.14159 / (v14 * ((3.14159 / ((v14 / float32(v1)) - ((float32(float64(v14)) * (0.0 - v14)) / (0.0 - ((1.5 / (v14 - float32(v1))) - v14))))) - (v14 + 3.14159))))
    } else {
        v14 = 0.0
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    v5 = v5
    fmt.Println(v14)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
