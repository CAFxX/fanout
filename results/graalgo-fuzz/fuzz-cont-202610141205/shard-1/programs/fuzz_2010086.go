package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int; F2 int64 }

func main() {
    var v1 uint8 = 127
    v2 := make(map[string]bool)
    var v3 int = 609
    v4 := []string{"", "\n\t", ""}
    var v5 string = ""
    v6 := make(map[uint32]string)
    for fi7 := 0; fi7 < 20; fi7++ {
        var v8 uint64 = 244
        var v9 uint8 = 1
        delete(v2, "\n\t")
        _ = fi7
        _ = v8
        _ = v9
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    for fi10 := 0; fi10 < 5; fi10++ {
        v3 = (^(v3))
        _ = fi10
    }
    v5 = v5
    switch v5 {
    case "hello":
        v1 = 42
    case "\n\t":
        var v11 float32 = -2.25
        _ = v11
    default:
        for fi12 := 0; fi12 < 13; fi12++ {
            var v13 int = (fi12 / (((fi12 % (((fi12 ^ v3)) | 1))) | 1))
            _ = fi12
            _ = v13
        }
    }
    fmt.Println(v1)
    v14 := GenS1{F0: (v3 - -850), F1: 127, F2: -650}
    for fi15 := 0; fi15 < 10; fi15++ {
        fmt.Println(v3, v5, v3)
        v16 := GenS1{F0: (fi15 & (v3 % (((fi15 | 2147483647)) | 1))), F1: -9223372036854775808, F2: 70}
        if true {
            if false {
                switch fi15 {
                case 9223372036854775807:
                    for fi17 := 0; fi17 < 2; fi17++ {
                        v5 = ((v5 + v5) + ((("foo bar" + v5) + v5) + v5))
                        _ = fi17
                    }
                default:
                    v3 += (-(v3))
                }
                v4 = append(v4, "\n\t")
            } else {
                for fi18 := 0; fi18 < 2; fi18++ {
                    if (2147483647 >= -1) {
                        var v19 bool = true
                        v20 := v6[395]
                        _ = v19
                        _ = v20
                    }
                    if (!false) {
                        v21 := GenS1{F0: (v3 | fi15), F1: (-(fi15)), F2: -1}
                        _ = v21
                    } else {
                        var v22 int64 = 180
                        var v23 int32 = 2147483647
                        v22 += v22
                        v5 = ("foo bar" + "foo bar")
                        _ = v22
                        _ = v23
                    }
                    _ = fi18
                }
            }
            var v24 int64 = 9223372036854775807
            _ = v24
        }
        _ = fi15
        _ = v16
    }
    for fi25 := 0; fi25 < 11; fi25++ {
        var v26 float64 = 1.5
        _ = fi25
        _ = v26
    }
    v1 = (v1 & 115)
    var v27 uint8 = 255
    var v28 uint8 = (v27 * v1)
    var v29 int32 = -132
    for fi30 := 0; fi30 < 1; fi30++ {
        if ((4294967295 < 314) || true) {
            _ = len(v4)
        } else {
            var v31 int16 = int16(v29)
            if (!(42 <= 1)) {
                var v32 string = (("\n\t" + v5) + "\n\t")
                _ = v32
            } else {
                v1 -= 136
            }
            _ = v31
        }
        v27 += v28
        _ = fi30
    }
    var v33 float64 = ((float64(1.5) + 0.0) - float64(0.0))
    var v34 uint = 331
    v4 = append(v4, (v5 + v5))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
