package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32; F2 uint64 }

func main() {
    var v1 int16 = 385
    var v2 int = 42
    var v3 float32 = 0.1
    v4 := make(map[int]string)
    v3 = (v3 - (float32((1e10 + (3.14159 / (3.14159 + 0.1)))) / (v3 - v3)))
    var v5 uint = 65535
    fmt.Println(v5, v1, v3)
    v1 += (v1 << (((v5 / (((v5 << ((4294967295) & 63))) | 1))) & 63))
    for fi6 := 0; fi6 < 16; fi6++ {
        v7 := GenS1{F0: (v1 / (((v1 - (v1 | (-(v1))))) | 1)), F1: 65535, F2: 42}
        v8 := GenS1{F0: (v1 | v1), F1: 65535, F2: 18446744073709551615}
        _ = fi6
        _ = v7
        _ = v8
    }
    _ = len(v4)
    var v9 int16 = 1000
    for fi10 := 0; fi10 < 16; fi10++ {
        fmt.Println(v1, v3, fi10)
        if ("" == ("hello" + "\n\t")) {
            v9 = (v1 | -32768)
            if true {
                v5 -= 1
                switch ("" + (("foo bar" + ("" + "\n\t")) + ("a" + ("\n\t" + "\n\t")))) {
                case "hello":
                    v1 += (v9 >> ((v5) & 63))
                case "foo bar":
                    v11 := GenS1{F0: (v9 + (v1 ^ 42)), F1: 1, F2: 750}
                    _ = v11
                default:
                    if (!(761 <= 4294967295)) {
                        fmt.Println(v9)
                    } else {
                        switch "foo bar" {
                        case "":
                            var v12 uint8 = 128
                            _ = v12
                        case "hello":
                            switch "foo bar" {
                            case "a":
                                var v13 uint8 = 255
                                _ = v13
                            }
                        default:
                            v9 += (v1 + 909)
                        }
                        var v14 int16 = v9
                        _ = v14
                    }
                }
            }
            var v15 int32 = 750
            _ = v15
        } else {
            var v16 int8 = -47
            _ = v16
        }
        _ = fi10
    }
    var v17 bool = false
    var v18 uint64 = 1
    var v19 string = ((("foo bar" + "") + "a") + ("foo bar" + "foo bar"))
    var v20 int64 = -427
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
