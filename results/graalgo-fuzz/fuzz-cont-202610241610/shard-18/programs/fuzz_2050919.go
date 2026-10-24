package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 string; F1 uint; F2 int64 }

func main() {
    v1 := make(map[int]string)
    var v2 bool = (2 != 128)
    var v3 int16 = 136
    v4 := []string{"", "", ("\n\t" + ("" + "foo bar")), "\n\t"}
    if (123 >= 352) {
        for fi5 := 0; fi5 < 18; fi5++ {
            if true {
                switch fi5 {
                case 576:
                    for fi6 := 0; fi6 < 7; fi6++ {
                        switch 4294967295 {
                        case 32767:
                            v7 := GenS1{F0: 2147483648}
                            v8 := GenS1{F0: 279}
                            var v9 uint8 = 2
                            _ = v7
                            _ = v8
                            _ = v9
                        case 127:
                            v3 -= v3
                        }
                        _ = fi6
                    }
                case 9223372036854775807:
                    v10 := GenS2{F0: (("foo bar" + "a") + "a"), F1: 65535, F2: 9223372036854775807}
                    _ = v10
                }
                v4 = append(v4, "foo bar")
            } else {
                if (517 != (v3 ^ v3)) {
                    var v11 int8 = int8(v3)
                    if false {
                        var v12 float32 = (((1.5 * (1.5 - float32((0.1 + 0.1)))) + float32(1.5)) + 3.14159)
                        _ = v12
                    }
                    _ = v11
                }
                v2 = v2
            }
            _ = fi5
        }
        v3 -= 0
    } else {
        v2 = v2
        switch 9223372036854775807 {
        case -679:
            var v13 string = ("hello" + ("hello" + "a"))
            _ = v13
        case 128:
            fmt.Println(v2, v2)
        case 127:
            var v14 uint64 = 119
            _ = v14
        default:
            v3 -= 0
        }
        switch ("a" + "a") {
        case "hello":
            fmt.Println(v3, v3)
        case "foo bar":
            for fi15 := 0; fi15 < 5; fi15++ {
                if (true || v2) {
                    if ((!false) && v2) {
                        _ = len(v1)
                        v2 = (!((("hello" + ("hello" + ("hello" + ("foo bar" + "")))) <= ("\n\t" + ("" + "foo bar"))) || (68 != 255)))
                    }
                    var v16 int16 = (v3 / ((-1) | 1))
                    v3 = -32768
                    _ = v16
                }
                _ = fi15
            }
        case "a":
            var v17 int32 = -2147483648
            _ = v17
        default:
            v2 = ((767 == 65535) && v2)
        }
    }
    if v2 {
        var v18 uint32 = 986
        _ = v18
    } else {
        if v2 {
            v19 := v1[2147483647]
            _ = v19
        } else {
            var v20 float64 = (float64(3.14159) * (0.0 + float64(float32((3.14159 + 1e10)))))
            switch "hello" {
            case "":
                v21 := GenS1{F0: 2147483647}
                _ = v21
            default:
                v20 += float64((1e10 + (1e10 - 0.0)))
            }
            _ = len(v1)
            _ = v20
        }
    }
    var v22 int32 = -81
    v3 = (v3 ^ (^(v3)))
    var v23 bool = v2
    switch "hello" {
    case "foo bar":
        v3 = -32768
    case "hello":
        fmt.Println(v22, v2, v2)
    }
    for fi24 := 0; fi24 < 11; fi24++ {
        var v25 uint64 = 816
        var v26 int64 = 718
        _ = fi24
        _ = v25
        _ = v26
    }
    var v27 uint = 44
    v2 = true
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v27
    _ = v3
    _ = v4
}
