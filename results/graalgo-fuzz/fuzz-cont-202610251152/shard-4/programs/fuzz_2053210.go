package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string; F2 uint8 }

func main() {
    var v1 uint8 = 255
    v2 := make(map[string]uint32)
    var v3 int = 32767
    v4 := GenS1{F0: int8(v1), F1: "", F2: 94}
    v5 := GenS1{F0: 100, F1: "hello", F2: 189}
    var v6 uint16 = 1
    delete(v2, ("hello" + ("\n\t" + ("hello" + "a"))))
    v1 += 127
    v6 = 643
    fmt.Println(v3)
    for fi7 := 0; fi7 < 3; fi7++ {
        fmt.Println(v3, v6, v3)
        var v8 bool = (0.0 < -2.25)
        v6 = v6
        _ = fi7
        _ = v8
    }
    var v9 string = ("" + ("foo bar" + ("\n\t" + ("foo bar" + (("hello" + "") + "hello")))))
    fmt.Println(v3, v6, v6)
    if ((v1 | (v1 | 2)) < v1) {
        v9 = (v9 + v9)
        v9 = ""
    } else {
        var v10 int8 = -128
        switch -1 {
        case 732:
            for fi11 := 0; fi11 < 20; fi11++ {
                var v12 uint64 = 362
                for fi13 := 0; fi13 < 4; fi13++ {
                    v12 = v12
                    v1 += v1
                    _ = fi13
                }
                _ = fi11
                _ = v12
            }
        }
        _ = v10
    }
    fmt.Println(v1, v1, v6)
    if ((0.0 - 0.1) >= ((1.5 / -2.25) - float32(float64(0.1)))) {
        if (-1 <= 9223372036854775807) {
            fmt.Println(v3, v3)
        } else {
            fmt.Println(v3, v9, v1)
            var v14 int16 = -234
            _ = v14
        }
        v3 += v3
        for fi15 := 0; fi15 < 5; fi15++ {
            if (0.1 <= 1e10) {
                var v16 uint32 = 154
                switch "" {
                case "\n\t":
                    var v17 uint32 = v16
                    _ = v17
                case "hello":
                    v6 = v6
                }
                if true {
                    v16 = v16
                    var v18 string = ("foo bar" + "")
                    v19 := GenS1{F0: -1, F1: v9, F2: (v1 & v1)}
                    _ = v18
                    _ = v19
                } else {
                    v3 += (fi15 + fi15)
                    v20 := GenS1{F0: 102, F1: v9, F2: (v1 >> ((255) & 63))}
                    _ = v20
                }
                _ = v16
            } else {
                switch 4294967295 {
                case 1:
                    var v21 float64 = 3.14159
                    _ = v21
                case 127:
                    var v22 uint8 = (v1 >> ((2) & 63))
                    _ = v22
                case 585:
                    var v23 int32 = 2147483647
                    _ = v23
                default:
                    for fi24 := 0; fi24 < 16; fi24++ {
                        v25 := GenS1{F0: 1, F1: ("\n\t" + ""), F2: (^(v1))}
                        fmt.Println(v1, v6, v1)
                        var v26 uint64 = 1
                        _ = fi24
                        _ = v25
                        _ = v26
                    }
                }
                var v27 float64 = 3.14159
                if true {
                    v9 = (((v9 + "hello") + v9) + v9)
                    v3 = fi15
                }
                _ = v27
            }
            _ = fi15
        }
    }
    fmt.Println(v9)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
