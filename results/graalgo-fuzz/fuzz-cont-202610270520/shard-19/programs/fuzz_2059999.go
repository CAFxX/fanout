package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint16; F1 uint32 }

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[string]string)
    var v3 string = ("foo bar" + (((((("hello" + ("a" + (("" + "a") + (("\n\t" + ("a" + "")) + ("foo bar" + ("\n\t" + "a")))))) + "foo bar") + "hello") + ((("foo bar" + "a") + "a") + "a")) + ("\n\t" + "")) + "foo bar"))
    v4 := []string{}
    var v5 int = 127
    v3 = v3
    var v6 float32 = 1.5
    switch "" {
    case "\n\t":
        var v7 int32 = -2147483648
        _ = v7
    case "a":
        v3 = v3
    default:
        v8 := GenS1{F0: 1000}
        _ = v8
    }
    for fi9 := 0; fi9 < 13; fi9++ {
        var v10 uint = 2
        v10 = (^(v10))
        if (false || ((v5 / (((fi9 & (v5 % (((v5 ^ v5)) | 1)))) | 1)) < (v5 % (((v5 % ((1) | 1))) | 1)))) {
            if false {
                var v11 uint32 = 132
                _ = v11
            } else {
                v12 := GenS2{F0: 0, F1: 0}
                v5 -= 32767
                _ = v12
            }
        }
        _ = fi9
        _ = v10
    }
    if len(v4) > 0 {
        v4[0] = "hello"
    }
    v5 = 641
    for fi13 := 0; fi13 < 16; fi13++ {
        v14 := v1[1]
        for fi15 := 0; fi15 < 9; fi15++ {
            _ = len(v4)
            if true {
                var v16 uint32 = 2147483648
                v5 += v5
                if ((fi15 % (((-(fi15))) | 1)) > fi13) {
                    v6 = v6
                } else {
                    switch v5 {
                    case 42:
                        v17 := GenS1{F0: 712}
                        _ = v17
                    case 32768:
                        var v18 uint = 18446744073709551615
                        _ = v18
                    default:
                        var v19 float64 = 0.1
                        _ = v19
                    }
                }
                _ = v16
            }
            v20 := GenS2{F0: 255, F1: 450}
            v21 := GenS2{F0: 32767, F1: 4294967295}
            v6 -= float32((1.5 + (float64(((v6 * v6) / v6)) + (0.1 + (1e10 - 1e10)))))
            _ = fi15
            _ = v20
            _ = v21
        }
        _ = fi13
        _ = v14
    }
    var v22 int64 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
